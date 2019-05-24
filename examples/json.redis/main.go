// Example: collect specify redis instance metrics from `info` command and dump to stdout in JSON format.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/OpenCCTV/redis/redis"
)

var (
	Addr     string
	Password string
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	flag.StringVar(&Addr, "addr", "127.0.0.1:6379", "redis instance address")
	flag.StringVar(&Password, "auth", "", "redis instance passsword")

	flag.Parse()

	metrics, err := metricsRedis.Gets(Addr, Password)
	if err != nil {
		log.Fatalln(err)
	}

	out, err := json.MarshalIndent(metrics, "", "  ")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(out))
}
