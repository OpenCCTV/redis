package metricsRedis

import (
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

func Gets(addr, password string) (*map[string]interface{}, error) {
	result := map[string]interface{}{}

	// TCP connection status test
	conn, err := net.DialTimeout("tcp", addr, time.Duration(2)*time.Second)
	if err != nil {
		result["ping"] = PingDown
		return &result, err
	}
	defer conn.Close()

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})
	defer client.Close()

	// redis ping-pong test
	pong, err := client.Ping().Result()
	if err != nil {
		errStr := strings.ToLower(err.Error())
		if strings.Index(errStr, "but no password is set") != -1 {
			// walk around issue redis raise error if you send auth password but it not required
			client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
			})
			defer client.Close()

			pong, err = client.Ping().Result()
		} else {
			result["ping"] = PingDown
			return &result, err
		}
	}

	if err != nil {
		result["ping"] = PingUp
		return &result, err
	}

	if pong != "PONG" {
		err = errors.New("redis ping got unexpected response")
		result["ping"] = PingUnknown
		return &result, err
	}

	reply, err := client.Info().Result()
	if err != nil {
		log.Println("get info failed", addr, err)
		result["ping"] = PingUnknown
		return &result, err
	}

	m := map[string]interface{}{}
	m["ping"] = PingUp
	m["cfg_maxmemory"] = GetFailed
	m["used_memory_rate"] = GetFailed
	m["hit_rate"] = GetFailed
	ParseReplyInfo(&reply, &m)

	maxmemory, err := getMaxMemory(client)
	if err == nil {
		m["cfg_maxmemory"] = maxmemory
	}

	usedRate, err := ComputingUsedMemoryRate(&m)
	if err == nil {
		m["used_memory_rate"] = usedRate
	}

	hitrate, err := computingHitRate(&m)
	if err == nil {
		m["hit_rate"] = hitrate
	}

	return &m, nil
}

func getMaxMemory(client *redis.Client) (maxMemory float64, err error) {
	resultMaxMemory, err := client.ConfigGet("maxmemory").Result()
	if err != nil {
		return
	}

	if len(resultMaxMemory) != 2 {
		return
	}

	maxMemory, err = strconv.ParseFloat(resultMaxMemory[1].(string), 64)
	if err != nil {
		return
	}
	return
}
