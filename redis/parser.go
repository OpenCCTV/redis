package metricsRedis

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// SliceIndex a generic function in idiomatic.
// See https://stackoverflow.com/questions/8307478/how-to-find-out-element-position-in-slice
func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func ParseReplyInfo(reply *string, result *map[string]interface{}) {
	//keysCount := len(MetricKeys)

	dbsize := 0
	for _, line := range strings.Split(*reply, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// skip line not match pattern `key:value`
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		// parse `db0:keys=2,expires=1,avg_ttl=30265707228` into `db0_keys=... db0_expires=...`
		if key[0] == 'd' && key[1] == 'b' && len(key) > 2 {
			dbsize++

			pairs := strings.Split(value, ",")
			for _, pair := range pairs {
				parts = strings.Split(pair, "=")
				subkey := parts[0]
				subvalue := parts[1]

				subkey = fmt.Sprintf("%s_%s", key, subkey)
				(*result)[subkey] = subvalue
			}
			//} else if SliceIndex(keysCount, func(i int) bool { return MetricKeys[i] == key }) != -1 {
		} else {
			(*result)[key] = value
		}

		(*result)["dbsize"] = dbsize
	}
}

func ParseReplyClusterInfo(reply *string, result *map[string]interface{}) {
	for _, line := range strings.Split(*reply, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// skip line not match pattern `key:value`
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]
		(*result)[key] = value
	}
}

func ParseCommandStatsInfo(reply *string, result *map[string]interface{}) {
	regCommandStats := regexp.MustCompile(`(.*):calls=(.*),usec=(.*),usec_per_call=(.*)\r`)
	for _, line := range strings.Split(*reply, "\n") {
		if regCommandStats.MatchString(line) {
			subMatch := regCommandStats.FindAllStringSubmatch(line, -1)
			(*result)[subMatch[0][1]+"_calls"], _ = strconv.ParseFloat(subMatch[0][2], 64)
			(*result)[subMatch[0][1]+"_usec"], _ = strconv.ParseFloat(subMatch[0][3], 64)
			(*result)[subMatch[0][1]+"_usec_per_call"], _ = strconv.ParseFloat(subMatch[0][4], 64)
		}
	}
}

func computingHitRate(m *map[string]interface{}) (rate int, err error) {
	keyspaceHitsI, ok := (*m)["keyspace_hits"].(string)
	if !ok {
		return
	}
	keyspaceMissesI, ok := (*m)["keyspace_misses"].(string)
	if !ok {
		return
	}

	keyspaceHits, err := strconv.ParseFloat(keyspaceHitsI, 64)
	if err != nil {
		return
	}

	keyspaceMisses, err := strconv.ParseFloat(keyspaceMissesI, 64)
	if err != nil {
		return
	}
	total := keyspaceHits + keyspaceMisses
	if total > 0 {
		rate = int(keyspaceHits / total * 100)
	}
	return

}

func ComputingUsedMemoryRate(m *map[string]interface{}) (rate int, err error) {
	usedMemoryI, ok := (*m)["used_memory"].(string)
	if !ok {
		return -1, errors.New("used_memory invalid")
	}

	usedMemory, err := strconv.ParseFloat(usedMemoryI, 64)
	if err != nil {
		return -1, errors.New("used_memory invalid")
	}
	if err == nil {
		cfgMaxmemory, ok := (*m)["cfg_maxmemory"].(float64)
		if ok {
			if cfgMaxmemory > 0 {
				rate = int(usedMemory / cfgMaxmemory * 100)
			}
			return rate, nil
		} else {
			err = errors.New("cfg_maxmemory invalid")
		}
	}
	return
}
