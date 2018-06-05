package metricsRedis

import (
	"testing"
)

func TestParseReplyInfoNotSetMaxmemory(t *testing.T) {
	sampleNotSetMaxmemory := `# Server
redis_version:3.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:687a2a319020fa42
redis_mode:standalone
os:Linux 4.4.0-124-generic x86_64
arch_bits:64
multiplexing_api:epoll
gcc_version:5.3.1
process_id:19714
run_id:507b0b1cefdc00139b47e83072e104a48133d4dc
tcp_port:6379
uptime_in_seconds:162735
uptime_in_days:1
hz:10
lru_clock:1098028
config_file:/etc/redis/redis.conf

# Clients
connected_clients:6
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:2

# Memory
used_memory:613992
used_memory_human:599.60K
used_memory_rss:3432448
used_memory_peak:745416
used_memory_peak_human:727.95K
used_memory_lua:36864
mem_fragmentation_ratio:5.59
mem_allocator:jemalloc-3.6.0

# Persistence
loading:0
rdb_changes_since_last_save:2
rdb_bgsave_in_progress:0
rdb_last_save_time:1527824168
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok

# Stats
total_connections_received:2794
total_commands_processed:4880192
instantaneous_ops_per_sec:29
total_net_input_bytes:92786569
total_net_output_bytes:34685064
instantaneous_input_kbps:0.56
instantaneous_output_kbps:0.17
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
evicted_keys:0
keyspace_hits:5
keyspace_misses:1
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:100
migrate_cached_sockets:0

# Replication
role:master
connected_slaves:0
master_repl_offset:0
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:363.70
used_cpu_user:180.52
used_cpu_sys_children:0.00
used_cpu_user_children:0.00

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=2,expires=1,avg_ttl=30260474352`

	m := map[string]interface{}{}
	ParseReplyInfo(&sampleNotSetMaxmemory, &m)

	connected_clients, ok := m["connected_clients"].(string)
	if !ok || connected_clients != "6" {
		t.Errorf("get connected_clients failed")
	}

	total_connections_received, ok := m["total_connections_received"].(string)
	if !ok || total_connections_received != "2794" {
		t.Errorf("get total_connections_received failed")
	}

	used_cpu_sys, ok := m["used_cpu_sys"].(string)
	if !ok || used_cpu_sys == "" {
		t.Errorf("get used_cpu_sys failed")
	}

	usedRate, _ := ComputingUsedMemoryRate(&m)
	if usedRate != 0 {
		t.Errorf("expected usedRate == 0, got %v", usedRate)
	}

	hitRate, _ := computingHitRate(&m)
	if hitRate != 83 {
		t.Errorf("expected hitRate == 83, got %v", hitRate)
	}
}

func TestParseReplyInfoSetMaxmemory(t *testing.T) {
	sampleSetMaxmemory := `# Server
redis_version:3.0.6
redis_git_sha1:00000000
redis_git_dirty:0
redis_build_id:687a2a319020fa42
redis_mode:standalone
os:Linux 4.4.0-124-generic x86_64
arch_bits:64
multiplexing_api:epoll
gcc_version:5.3.1
process_id:19714
run_id:507b0b1cefdc00139b47e83072e104a48133d4dc
tcp_port:6379
uptime_in_seconds:162735
uptime_in_days:1
hz:10
lru_clock:1098028
config_file:/etc/redis/redis.conf

# Clients
connected_clients:6
client_longest_output_list:0
client_biggest_input_buf:0
blocked_clients:2

# Memory
used_memory:613992
used_memory_human:599.60K
used_memory_rss:3432448
used_memory_peak:745416
used_memory_peak_human:727.95K
used_memory_lua:36864
mem_fragmentation_ratio:5.59
mem_allocator:jemalloc-3.6.0

# Persistence
loading:0
rdb_changes_since_last_save:2
rdb_bgsave_in_progress:0
rdb_last_save_time:1527824168
rdb_last_bgsave_status:ok
rdb_last_bgsave_time_sec:0
rdb_current_bgsave_time_sec:-1
aof_enabled:0
aof_rewrite_in_progress:0
aof_rewrite_scheduled:0
aof_last_rewrite_time_sec:-1
aof_current_rewrite_time_sec:-1
aof_last_bgrewrite_status:ok
aof_last_write_status:ok

# Stats
total_connections_received:2794
total_commands_processed:4880192
instantaneous_ops_per_sec:29
total_net_input_bytes:92786569
total_net_output_bytes:34685064
instantaneous_input_kbps:0.56
instantaneous_output_kbps:0.17
rejected_connections:0
sync_full:0
sync_partial_ok:0
sync_partial_err:0
expired_keys:0
evicted_keys:0
keyspace_hits:5
keyspace_misses:1
pubsub_channels:0
pubsub_patterns:0
latest_fork_usec:100
migrate_cached_sockets:0

# Replication
role:master
connected_slaves:0
master_repl_offset:0
repl_backlog_active:0
repl_backlog_size:1048576
repl_backlog_first_byte_offset:0
repl_backlog_histlen:0

# CPU
used_cpu_sys:363.70
used_cpu_user:180.52
used_cpu_sys_children:0.00
used_cpu_user_children:0.00

# Cluster
cluster_enabled:0

# Keyspace
db0:keys=2,expires=1,avg_ttl=30260474352`

	m := map[string]interface{}{}
	ParseReplyInfo(&sampleSetMaxmemory, &m)
	m["cfg_maxmemory"] = float64(1048576.0)

	connected_clients, ok := m["connected_clients"].(string)
	if !ok || connected_clients != "6" {
		t.Errorf("get connected_clients failed")
	}

	total_connections_received, ok := m["total_connections_received"].(string)
	if !ok || total_connections_received != "2794" {
		t.Errorf("get total_connections_received failed")
	}

	used_cpu_sys, ok := m["used_cpu_sys"].(string)
	if !ok || used_cpu_sys == "" {
		t.Errorf("get used_cpu_sys failed")
	}

	usedRate, _ := ComputingUsedMemoryRate(&m)
	if usedRate != 58 {
		t.Errorf("expected usedRate == 58, got %v", usedRate)
	}

	hitRate, _ := computingHitRate(&m)
	if hitRate != 83 {
		t.Errorf("expected hitRate == 83, got %v", hitRate)
	}
}
