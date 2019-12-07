package metricsRedis

import (
	"errors"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis"
)

var (
	ConnPools = make(map[string]*redis.Client)
	cpLock    = new(sync.RWMutex)
)

func GetConnPools() *map[string]*redis.Client {
	cpLock.RLock()
	defer cpLock.RUnlock()
	return &ConnPools
}

func GetInfoByPool(addr, password string) (*map[string]interface{}, error) {
	result := map[string]interface{}{}
	inscKey := addr
	// TCP connection status test
	conn, err := net.DialTimeout("tcp", addr, time.Duration(2)*time.Second)
	if err != nil {
		result["ping"] = PingDown
		return &result, err
	}
	defer conn.Close()
	result["ping"] = PingUp
	var client *redis.Client = nil

	cpLock.Lock()
	if _, ok := ConnPools[inscKey]; ok {
		client = ConnPools[inscKey]
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}

	// redis ping-pong test
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("net.DialTimeout", addr, password, err)

		errStr := strings.ToLower(err.Error())
		if strings.Index(errStr, "but no password is set") != -1 {
			// walk around issue redis raise error if you send auth password but it not required
			client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
			})
			pong, err = client.Ping().Result()
		} else {
			result["ping"] = PingDown
			cpLock.Unlock()
			return &result, err
		}
	}
	ConnPools[inscKey] = client

	if err != nil {
		result["ping"] = PingUp
		cpLock.Unlock()
		return &result, err
	}

	if pong != "PONG" {
		err = errors.New("redis ping got unexpected response")
		result["ping"] = PingUnknown
		cpLock.Unlock()
		return &result, err
	}
	cpLock.Unlock()

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

func GetClusterInfoByPool(addr, password string) (*map[string]interface{}, error) {
	result := map[string]interface{}{}
	inscKey := addr
	// TCP connection status test
	conn, err := net.DialTimeout("tcp", addr, time.Duration(2)*time.Second)
	if err != nil {
		result["ping"] = PingDown
		return &result, err
	}
	defer conn.Close()
	result["ping"] = PingUp
	var client *redis.Client = nil

	cpLock.Lock()
	if _, ok := ConnPools[inscKey]; ok {
		client = ConnPools[inscKey]
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}

	// redis ping-pong test
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("net.DialTimeout", addr, password, err)

		errStr := strings.ToLower(err.Error())
		if strings.Index(errStr, "but no password is set") != -1 {
			// walk around issue redis raise error if you send auth password but it not required
			client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
			})
			pong, err = client.Ping().Result()
		} else {
			result["ping"] = PingDown
			cpLock.Unlock()
			return &result, err
		}
	}
	ConnPools[inscKey] = client

	if err != nil {
		result["ping"] = PingUp
		cpLock.Unlock()
		return &result, err
	}

	if pong != "PONG" {
		err = errors.New("redis ping got unexpected response")
		result["ping"] = PingUnknown
		cpLock.Unlock()
		return &result, err
	}
	cpLock.Unlock()

	reply, err := client.ClusterInfo().Result()
	if err != nil {
		log.Println("get info failed", addr, err)
		result["ping"] = PingUnknown
		return &result, err
	}

	m := map[string]interface{}{}
	ParseReplyClusterInfo(&reply, &m)

	return &m, nil
}

func GetInfoCommandStatsByPool(addr, password string) (*map[string]interface{}, error) {
	result := map[string]interface{}{}
	inscKey := addr
	// TCP connection status test
	conn, err := net.DialTimeout("tcp", addr, time.Duration(2)*time.Second)
	if err != nil {
		result["ping"] = PingDown
		return &result, err
	}
	defer conn.Close()
	result["ping"] = PingUp
	var client *redis.Client = nil

	cpLock.Lock()
	if _, ok := ConnPools[inscKey]; ok {
		client = ConnPools[inscKey]
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		})
		ConnPools[inscKey] = client
	}

	// redis ping-pong test
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("net.DialTimeout", addr, password, err)

		errStr := strings.ToLower(err.Error())
		if strings.Index(errStr, "but no password is set") != -1 {
			// walk around issue redis raise error if you send auth password but it not required
			client = redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: "",
				DB:       0,
			})
			pong, err = client.Ping().Result()
		} else {
			result["ping"] = PingDown
			cpLock.Unlock()
			return &result, err
		}
	}
	ConnPools[inscKey] = client

	if err != nil {
		result["ping"] = PingUp
		cpLock.Unlock()
		return &result, err
	}

	if pong != "PONG" {
		err = errors.New("redis ping got unexpected response")
		result["ping"] = PingUnknown
		cpLock.Unlock()
		return &result, err
	}
	cpLock.Unlock()

	reply, err := client.Do("info", "commandstats").Result()
	if err != nil {
		log.Println("get info commandstats failed", addr, err)
		result["ping"] = PingUnknown
		return &result, err
	}
	replyString := reply.(string)
	m := map[string]interface{}{}
	ParseCommandStatsInfo(&replyString, &m)

	return &m, nil
}
