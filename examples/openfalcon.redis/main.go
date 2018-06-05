// Example: collect specify redis instance metrics from `info` command and post them to open-falcon agent HTTP PUSH API.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/MonitorMetrics/falcon_helpers/agent"
	"github.com/MonitorMetrics/falcon_helpers/model"
	"github.com/MonitorMetrics/redis/redis"
)

var (
	Debug           = flag.Bool("debug", false, `enable debug`)
	FalconURL       = flag.String("url", "http://127.0.0.1:1988/v1/push", "falcon agnet PUSH URL")
	IntervalCollect = flag.Int("interval", 60, "collect stat interval in seconds")
	ShortenInterval = flag.Bool("s", false, "shorten interval to 5 seconds for debug")
	RedisAddr       = flag.String("addr", "127.0.0.1:6379", "redis instance address")
	RedisPassword   = flag.String("auth", "", "redis instance passsword")
)

func funcCallback(m *[]map[string]interface{}, args ...interface{}) {
	metrics := []*modelMetric.MetricItem{}

	redisAddr := args[0]

	now := time.Now().Unix()

	for _, item := range *m {
		for k, v := range item {
			tags := fmt.Sprintf("addr=%v", redisAddr)
			item := modelMetric.MetricItem{
				Endpoint:    "127.0.0.1",
				Metric:      metricsRedis.MetricPrefix + k,
				Value:       v,
				CounterType: "GAUGE",
				Tags:        tags,
				Timestamp:   now,
				Step:        60,
			}
			metrics = append(metrics, &item)
		}
	}

	out, err := json.Marshal(metrics)
	if err != nil {
		log.Println(err)
		return
	}

	respBody := helperAgent.SendToFalconAgent(*FalconURL, string(out))
	if *Debug {
		log.Println("falcon response", respBody)
	}
}

func collect(intervalCollect int, funcCallback helperAgent.FuncPush2Agent, redisAddr, password string) {
	for {
		metric, err := metricsRedis.Gets(redisAddr, password)
		if err != nil {
			log.Println("metricsRedis.Gets", err)
		}
		metrics := []map[string]interface{}{}
		metrics = append(metrics, *metric)
		funcCallback(&metrics, redisAddr)

		time.Sleep(time.Second * time.Duration(intervalCollect))
	}
}

func main() {
	flag.Parse()

	if *Debug {
		log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	}

	intervalInSeconds := *IntervalCollect
	if *ShortenInterval {
		intervalInSeconds = 5
	}

	redisAddr := *RedisAddr
	password := *RedisPassword

	if *Debug {
		log.Println("collect redis instance ", redisAddr, password)
	}

	go collect(intervalInSeconds, funcCallback, redisAddr, password)

	select {}
}
