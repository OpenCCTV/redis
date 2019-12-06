package metricsRedis

var (
	MetricKeys = []string{
		// clients
		"connected_clients",          // 1
		"client_longest_output_list", // 0
		"client_biggest_input_buf",   // 0
		"blocked_clients",            // 0

		// memory
		"used_memory", // 508896
		//"used_memory_human",       // 496.97K
		"used_memory_rss",  // 6852608
		"used_memory_peak", // 508896
		//"used_memory_peak_human",  // 496.97K
		"used_memory_lua",         // 36864
		"mem_fragmentation_ratio", // 13.47

		// Persistence
		"loading",                      // 0
		"rdb_changes_since_last_save",  // 1
		"rdb_bgsave_in_progress",       // 0
		"rdb_last_save_time",           // 1524023227
		"rdb_last_bgsave_time_sec",     // -1
		"rdb_current_bgsave_time_sec",  // -1
		"aof_enabled",                  // 0
		"aof_rewrite_in_progress",      // 0
		"aof_rewrite_scheduled",        // 0
		"aof_last_rewrite_time_sec",    // -1
		"aof_current_rewrite_time_sec", // -1

		// Stats
		"total_connections_received", // 14
		"total_commands_processed",   // 12
		"instantaneous_ops_per_sec",  // 0
		"total_net_input_bytes",      // 201
		"total_net_output_bytes",     // 19191
		"instantaneous_input_kbps",   // 0.01
		"instantaneous_output_kbps",  // 0.00
		"rejected_connections",       // 0
		"sync_full",                  // 0
		"sync_partial_ok",            // 0
		"sync_partial_err",           // 0
		"expired_keys",               // 0
		"evicted_keys",               // 0
		"keyspace_hits",              // 1
		"keyspace_misses",            // 0
		"pubsub_channels",            // 0
		"pubsub_patterns",            // 0
		"latest_fork_usec",           // 0
		"migrate_cached_sockets",     // 0

		// Replication
		"connected_slaves",               // 0
		"master_repl_offset",             // 0
		"repl_backlog_active",            // 0
		"repl_backlog_size",              // 1048576
		"repl_backlog_first_byte_offset", // 0
		"repl_backlog_histlen",           // 0

		// CPU
		"used_cpu_sys",           // 0.23
		"used_cpu_user",          // 0.02
		"used_cpu_sys_children",  // 0.00
		"used_cpu_user_children", // 0.00

		// Cluster
		"cluster_enabled", // 0

		// Keyspace
		// "db0:keys=1,expires=0,avg_ttl=0" => "db*_keys" "db*_expires" "db*_avg_ttl"
	}

	MetricCounterTypes = map[string]string{
		"total_commands_processed":   "COUNTER",
		"total_connections_received": "COUNTER",
		"keyspace_hits":              "COUNTER",
		"keyspace_misses":            "COUNTER",
		"expired_keys":               "COUNTER",
		"used_cpu_sys":               "COUNTER",
		"used_cpu_sys_children":      "COUNTER",
		"total_net_input_bytes":      "COUNTER",
	}
	DefaultMetricCounterType = "GAUGE"
)

func GetCounterType(key string) string {
	t, ok := MetricCounterTypes[key]
	if ok {
		return t
	}
	return DefaultMetricCounterType
}
