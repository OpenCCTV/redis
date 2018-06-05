# Redis metrics for monitoring.

We do not parse values in `int` or `float*` type and keep them original, end-user should take care type by itself.
For example, Open-falcon will parses all values into JSON `float64` finally.


## Metrics

Following metrics get from redis issue command `info` output.

Section `clients`

| KEY | TYPE | NOTES |
|-----|------|-------|
| connected_clients | GUAGE |
| client_longest_output_list | GUAGE |
| client_biggest_input_buf | GUAGE |
| blocked_clients | GUAGE |

Section `memory`

| KEY | TYPE | NOTES |
|-----|------|-------|
| used_memory | GAUGE  |
| used_memory_rss | GAUGE |
| used_memory_peak | GAUGE |
| used_memory_lua | GAUGE |
| mem_fragmentation_ratio | GAUGE | in percent |


Section `Persistence`

| KEY | TYPE | NOTES |
|-----|------|-------|
| loading | GUAGE |
| rdb_changes_since_last_save | GUAGE |
| rdb_bgsave_in_progress | GUAGE |
| rdb_last_save_time  | GUAGE |
| rdb_last_bgsave_time_sec | GUAGE |
| rdb_current_bgsave_time_sec | GUAGE |
| aof_enabled | GUAGE |
| aof_rewrite_in_progress | GUAGE |
| aof_rewrite_scheduled | GUAGE |
| aof_last_rewrite_time_sec | GUAGE |
| aof_current_rewrite_time_sec | GUAGE |


Section `Stats`

| KEY | TYPE | NOTES |
|-----|------|-------|
| total_connections_received | GUAGE |
| total_commands_processed | GUAGE |
| instantaneous_ops_per_sec | GUAGE |
| total_net_input_bytes | GUAGE |
| total_net_output_bytes | GUAGE |
| instantaneous_input_kbps  | GUAGE |
| instantaneous_output_kbps | GUAGE |
| rejected_connections | GUAGE |
| sync_full | GUAGE |
| sync_partial_ok | GUAGE |
| sync_partial_err | GUAGE |
| expired_keys | GUAGE |
| evicted_keys | GUAGE |
| keyspace_hits | GUAGE |
| keyspace_misses | GUAGE |
| pubsub_channels | GUAGE |
| pubsub_patterns | GUAGE |
| latest_fork_usec | GUAGE |
| migrate_cached_sockets | GUAGE |


Section `Replication`

| KEY | TYPE | NOTES |
|-----|------|-------|
| connected_slaves | GUAGE |
| master_repl_offset | GUAGE |
| repl_backlog_active | GUAGE |
| repl_backlog_size | GUAGE |
| repl_backlog_first_byte_offset | GUAGE |
| repl_backlog_histlen | GUAGE |

Section `CPU`

| KEY | TYPE | NOTES |
|-----|------|-------|
| used_cpu_sys | GUAGE |
| used_cpu_user | GUAGE |
| used_cpu_sys_children | GUAGE |
| used_cpu_user_children | GUAGE |

Section `Cluster`

| KEY | TYPE | NOTES |
|-----|------|-------|
| cluster_enabled | GUAGE |


Section `Keyspace`


All keys  will transform into following formats

	"db0:keys=1,expires=0,avg_ttl=0" => "db*_keys" "db*_expires" "db*_avg_ttl"


| KEY | TYPE | NOTES |
|-----|------|-------|
| db0_keys | GAUGE |
| db0_expires | GAUGE |
| db0_avg_ttl | GAUGE |


### Extra Metrics

| KEY | TYPE | NOTES |
|-----|------|-------|
| ping | GAUGE | 1 for up, 0 for down, -1 for unknown |
| cfg_maxmemory | GAUGE | get from `config get maxmemory` |
| hit_rate | GAUGE | `keyspace_hits / (keyspace_hits + keyspace_misses) * 100` |
| used_memory_rate | GAUGE | 0 if cfg_maxmemory=0, or `used_memory/cfg_maxmemory*100`  |


## Build

Build example

	go get -v github.com/MonitorMetrics/redis
	cd $GOPATH/src/github.com/MonitorMetrics/redis/examples/json.redis
	go build -o json.redis.bin


Start it

	./json.redis.bin -addr 127.0.0.1:6379

For more detail, see source code.




