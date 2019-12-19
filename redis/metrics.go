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
		"total_net_output_bytes":     "COUNTER",
		"blocked_clients":            "COUNTER",
		"rejected_connections":       "COUNTER",
		
		"cmdstat_append_calls":   "COUNTER",
		"cmdstat_asking_calls":   "COUNTER",
		"cmdstat_auth_calls":   "COUNTER",
		"cmdstat_bgsave_calls":   "COUNTER",
		"cmdstat_blpop_calls":   "COUNTER",
		"cmdstat_brpoplpush_calls":   "COUNTER",
		"cmdstat_client_calls":   "COUNTER",
		"cmdstat_cluster_calls":   "COUNTER",
		"cmdstat_clusteradmin_calls":   "COUNTER",
		"cmdstat_command_calls":   "COUNTER",
		"cmdstat_config_calls":   "COUNTER",
		"cmdstat_dbsize_calls":   "COUNTER",
		"cmdstat_debug_calls":   "COUNTER",
		"cmdstat_decr_calls":   "COUNTER",
		"cmdstat_del_calls":   "COUNTER",
		"cmdstat_discard_calls":   "COUNTER",
		"cmdstat_dump_calls":   "COUNTER",
		"cmdstat_eval_calls":   "COUNTER",
		"cmdstat_evalsha_calls":   "COUNTER",
		"cmdstat_exec_calls":   "COUNTER",
		"cmdstat_exists_calls":   "COUNTER",
		"cmdstat_expire_calls":   "COUNTER",
		"cmdstat_expireat_calls":   "COUNTER",
		"cmdstat_flushall_calls":   "COUNTER",
		"cmdstat_flushdb_calls":   "COUNTER",
		"cmdstat_get_calls":   "COUNTER",
		"cmdstat_getset_calls":   "COUNTER",
		"cmdstat_hdel_calls":   "COUNTER",
		"cmdstat_hexists_calls":   "COUNTER",
		"cmdstat_hget_calls":   "COUNTER",
		"cmdstat_hgetall_calls":   "COUNTER",
		"cmdstat_hincrby_calls":   "COUNTER",
		"cmdstat_hincrbyfloat_calls":   "COUNTER",
		"cmdstat_hkeys_calls":   "COUNTER",
		"cmdstat_hlen_calls":   "COUNTER",
		"cmdstat_hmget_calls":   "COUNTER",
		"cmdstat_hmset_calls":   "COUNTER",
		"cmdstat_hscan_calls":   "COUNTER",
		"cmdstat_hset_calls":   "COUNTER",
		"cmdstat_hsetnx_calls":   "COUNTER",
		"cmdstat_hvals_calls":   "COUNTER",
		"cmdstat_incr_calls":   "COUNTER",
		"cmdstat_incrby_calls":   "COUNTER",
		"cmdstat_info_calls":   "COUNTER",
		"cmdstat_keys_calls":   "COUNTER",
		"cmdstat_latency_calls":   "COUNTER",
		"cmdstat_llen_calls":   "COUNTER",
		"cmdstat_lpop_calls":   "COUNTER",
		"cmdstat_lpush_calls":   "COUNTER",
		"cmdstat_lrange_calls":   "COUNTER",
		"cmdstat_lrem_calls":   "COUNTER",
		"cmdstat_lset_calls":   "COUNTER",
		"cmdstat_ltrim_calls":   "COUNTER",
		"cmdstat_memory_calls":   "COUNTER",
		"cmdstat_mget_calls":   "COUNTER",
		"cmdstat_migrate_calls":   "COUNTER",
		"cmdstat_monitor_calls":   "COUNTER",
		"cmdstat_mset_calls":   "COUNTER",
		"cmdstat_multi_calls":   "COUNTER",
		"cmdstat_object_calls":   "COUNTER",
		"cmdstat_pexpire_calls":   "COUNTER",
		"cmdstat_pexpireat_calls":   "COUNTER",
		"cmdstat_pfadd_calls":   "COUNTER",
		"cmdstat_pfcount_calls":   "COUNTER",
		"cmdstat_ping_calls":   "COUNTER",
		"cmdstat_psetex_calls":   "COUNTER",
		"cmdstat_psubscribe_calls":   "COUNTER",
		"cmdstat_psync_calls":   "COUNTER",
		"cmdstat_pttl_calls":   "COUNTER",
		"cmdstat_publish_calls":   "COUNTER",
		"cmdstat_punsubscribe_calls":   "COUNTER",
		"cmdstat_purgeslotasync_calls":   "COUNTER",
		"cmdstat_randomkey_calls":   "COUNTER",
		"cmdstat_readonly_calls":   "COUNTER",
		"cmdstat_rename_calls":   "COUNTER",
		"cmdstat_replconf_calls":   "COUNTER",
		"cmdstat_restore-asking_calls":   "COUNTER",
		"cmdstat_restore_calls":   "COUNTER",
		"cmdstat_rpop_calls":   "COUNTER",
		"cmdstat_rpoplpush_calls":   "COUNTER",
		"cmdstat_rpush_calls":   "COUNTER",
		"cmdstat_sadd_calls":   "COUNTER",
		"cmdstat_save_calls":   "COUNTER",
		"cmdstat_scan_calls":   "COUNTER",
		"cmdstat_scard_calls":   "COUNTER",
		"cmdstat_script_calls":   "COUNTER",
		"cmdstat_select_calls":   "COUNTER",
		"cmdstat_set_calls":   "COUNTER",
		"cmdstat_setbit_calls":   "COUNTER",
		"cmdstat_setex_calls":   "COUNTER",
		"cmdstat_setnx_calls":   "COUNTER",
		"cmdstat_sismember_calls":   "COUNTER",
		"cmdstat_slaveof_calls":   "COUNTER",
		"cmdstat_slowlog_calls":   "COUNTER",
		"cmdstat_bitcount_calls":   "COUNTER",
		"cmdstat_brpop_calls":   "COUNTER",
		"cmdstat_sdiff_calls":   "COUNTER",
		"cmdstat_smembers_calls":   "COUNTER",
		"cmdstat_spop_calls":   "COUNTER",
		"cmdstat_srandmember_calls":   "COUNTER",
		"cmdstat_srem_calls":   "COUNTER",
		"cmdstat_strlen_calls":   "COUNTER",
		"cmdstat_subscribe_calls":   "COUNTER",
		"cmdstat_sync_calls":   "COUNTER",
		"cmdstat_ttl_calls":   "COUNTER",
		"cmdstat_type_calls":   "COUNTER",
		"cmdstat_unsubscribe_calls":   "COUNTER",
		"cmdstat_zadd_calls":   "COUNTER",
		"cmdstat_zcard_calls":   "COUNTER",
		"cmdstat_zcount_calls":   "COUNTER",
		"cmdstat_zincrby_calls":   "COUNTER",
		"cmdstat_zinterstore_calls":   "COUNTER",
		"cmdstat_zrange_calls":   "COUNTER",
		"cmdstat_zrangebyscore_calls":   "COUNTER",
		"cmdstat_zrank_calls":   "COUNTER",
		"cmdstat_zrem_calls":   "COUNTER",
		"cmdstat_zremrangebyrank_calls":   "COUNTER",
		"cmdstat_zremrangebyscore_calls":   "COUNTER",
		"cmdstat_zrevrange_calls":   "COUNTER",
		"cmdstat_zrevrangebyscore_calls":   "COUNTER",
		"cmdstat_zrevrank_calls":   "COUNTER",
		"cmdstat_zscore_calls":   "COUNTER",
		"cmdstat_zunionstore_calls":   "COUNTER",
		"cmdstat_sunionstore_calls":   "COUNTER",
		"cmdstat_echo_calls":   "COUNTER",
		"cmdstat_getrange_calls":   "COUNTER",
		"cmdstat_lindex_calls":   "COUNTER",
		"cmdstat_pubsub_calls":   "COUNTER",
		"cmdstat_smove_calls":   "COUNTER",
		"cmdstat_sscan_calls":   "COUNTER",
		"cmdstat_time_calls":   "COUNTER",
		"cmdstat_watch_calls":   "COUNTER",
		"cmdstat_getbit_calls":   "COUNTER",
		"cmdstat_persist_calls":   "COUNTER",
		"cmdstat_zscan_calls":   "COUNTER",
		"cmdstat_unwatch_calls":   "COUNTER",
		"cmdstat_setrange_calls":   "COUNTER",
		"cmdstat_sinterstore_calls":   "COUNTER",
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
