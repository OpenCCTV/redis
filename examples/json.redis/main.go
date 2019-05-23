// Example: collect specify redis instance metrics from `info` command and dump to stdout in JSON format.
package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/OpenCCTV/redis/redis"
)

var (
	Addr     = flag.String("addr", "127.0.0.1:6379", "redis instance address")
	Password = flag.String("auth", "", "redis instance passsword")
)

func main() {
	flag.Parse()
	metrics, err := metricsRedis.Gets(*Addr, *Password)
	if err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
