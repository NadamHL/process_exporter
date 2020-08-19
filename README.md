# process_exporter
旨在收集操作系统进程的cpu、mem、io读写、文件打开数等信息

#使用指南
默认端口19124，默认metrics，可以自行修改
```
# ./process_exporter -h
usage: process_exporter [<flags>]

Flags:
  -h, --help  Show context-sensitive help (also try --help-long and --help-man).
      --web.listen-address=":19124"  
              Address on which to expose metrics and web interface.
      --web.telemetry-path="/metrics"  
              Path under which to expose metrics.
```

#获取metrics
```
#curl http://x.x.x.x:19124/metrics
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 3.3159e-05
go_gc_duration_seconds{quantile="0.25"} 5.7363e-05
go_gc_duration_seconds{quantile="0.5"} 6.9539e-05
go_gc_duration_seconds{quantile="0.75"} 8.5687e-05
go_gc_duration_seconds{quantile="1"} 0.000142635
go_gc_duration_seconds_sum 16.820908414
go_gc_duration_seconds_count 232007
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 42
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.14.6"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 4.490936e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 5.32542721368e+11
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.722972e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 5.797465657e+09
# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.
# TYPE go_memstats_gc_cpu_fraction gauge
go_memstats_gc_cpu_fraction 0.043888202200900533
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.860744e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 4.490936e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 5.67296e+07
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 7.200768e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 46625
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 4.927488e+07
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 6.3930368e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.5978080955831974e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 5.797512282e+09
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 69440
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 81920
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 482936
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 606208
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 6.500864e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 8.228252e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 2.916352e+06
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 2.916352e+06
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 8.1346816e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 52
# HELP process_cpu_percent Get the cpu usage rate of a single process
# TYPE process_cpu_percent gauge
process_cpu_percent{ProCreateTime="1590397599000",ProName="kthreadd",ProPid="2",ProUserName="root",proCwd="/"} 0.020480813529649835
process_cpu_percent{ProCreateTime="1590397599000",ProName="systemd",ProPid="1",ProUserName="root",proCwd="/"} 0.01767007079573879
process_cpu_percent{ProCreateTime="1590397606000",ProName="lvmetad",ProPid="5841",ProUserName="root",proCwd="/"} 0
process_cpu_percent{ProCreateTime="1590397606000",ProName="systemd-journald",ProPid="5816",ProUserName="root",proCwd="/"} 0.0035471340130076878
process_cpu_percent{ProCreateTime="1590397606000",ProName="systemd-udevd",ProPid="5850",ProUserName="root",proCwd="/"} 7.826743192355279e-06
process_cpu_percent{ProCreateTime="1590397608000",ProName="abrt-watch-log",ProPid="11526",ProUserName="root",proCwd="/"} 0.00023061909597330807
process_cpu_percent{ProCreateTime="1590397608000",ProName="abrtd",ProPid="11501",ProUserName="root",proCwd="/"} 6.477306438750956e-06
process_cpu_percent{ProCreateTime="1590397608000",ProName="dbus-daemon",ProPid="11584",ProUserName="dbus",proCwd="/"} 0.008587423792622987
process_cpu_percent{ProCreateTime="1590397608000",ProName="gssproxy",ProPid="11533",ProUserName="root",proCwd="/"} 0
process_cpu_percent{ProCreateTime="1590397608000",ProName="irqbalance",ProPid="11504",ProUserName="root",proCwd="/"} 0.041643547545575636
process_cpu_percent{ProCreateTime="1590397608000",ProName="polkitd",ProPid="11566",ProUserName="polkitd",proCwd="/"} 0.00381405387765729
process_cpu_percent{ProCreateTime="1590397608000",ProName="rpcbind",ProPid="11513",ProUserName="rpc",proCwd="/"} 0.00014101635809906577
process_cpu_percent{ProCreateTime="1590397608000",ProName="systemd-logind",ProPid="11564",ProUserName="root",proCwd="/"} 0.0035982786167737947
process_cpu_percent{ProCreateTime="1590397613000",ProName="login",ProPid="11959",ProUserName="root",proCwd="/"} 2.428991463390325e-06
process_cpu_percent{ProCreateTime="1590397613000",ProName="master",ProPid="12391",ProUserName="root",proCwd="/var/spool/postfix"} 0.00028338233656584593
process_cpu_percent{ProCreateTime="1590397613000",ProName="qmgr",ProPid="12411",ProUserName="postfix",proCwd="/var/spool/postfix"} 7.259985547177053e-05
process_cpu_percent{ProCreateTime="1590397613000",ProName="sshd",ProPid="11863",ProUserName="root",proCwd="/"} 3.238655331877452e-06
process_cpu_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11877",ProUserName="zabbix",proCwd="/"} 0
process_cpu_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11887",ProUserName="zabbix",proCwd="/"} 0.1612950202393421
process_cpu_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11888",ProUserName="zabbix",proCwd="/"} 0.11810431279746096
process_cpu_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11889",ProUserName="zabbix",proCwd="/"} 0.11816490238863821
process_cpu_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11890",ProUserName="zabbix",proCwd="/"} 0.003995556020753323
process_cpu_percent{ProCreateTime="1590397634000",ProName="bash",ProPid="21198",ProUserName="root",proCwd="/root"} 2.5639426755444953e-06
process_cpu_percent{ProCreateTime="1590399253000",ProName="rsyslogd",ProPid="22525",ProUserName="root",proCwd="/"} 0.004449817535199111
process_cpu_percent{ProCreateTime="1590399383000",ProName="virtlogd",ProPid="22700",ProUserName="root",proCwd="/"} 1.3497620745611164e-06
process_cpu_percent{ProCreateTime="1590560304000",ProName="auditd",ProPid="173542",ProUserName="root",proCwd="/"} 0.0006704110227570324
process_cpu_percent{ProCreateTime="1590560330000",ProName="containerd",ProPid="173638",ProUserName="root",proCwd="/"} 1.101434345830524
process_cpu_percent{ProCreateTime="1590560330000",ProName="dockerd",ProPid="173643",ProUserName="root",proCwd="/"} 0.07924317022190132
process_cpu_percent{ProCreateTime="1590560569000",ProName="containerd-shim",ProPid="175001",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/6fed7cc0f26af9f20adefd34aa51bce91d220852bd896bbf0023f18318903e52"} 0.00428352974243338
process_cpu_percent{ProCreateTime="1590560569000",ProName="sh",ProPid="175020",ProUserName="root",proCwd="/"} 6.898904385307638e-07
process_cpu_percent{ProCreateTime="1590565963000",ProName="bash",ProPid="194392",ProUserName="root",proCwd="/"} 6.904042612548123e-07
process_cpu_percent{ProCreateTime="1590565963000",ProName="containerd-shim",ProPid="194362",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/4f4390d2ec436320b4bfec2dbab2e199af067afb8049b5ec768ada052f252a32"} 0.004223617133259783
process_cpu_percent{ProCreateTime="1590569813000",ProName="containerd-shim",ProPid="58380",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/98f4825dfb5f8c71e0b4339634b251dc8c82687cfefca5ed37bf83d26e2c6108"} 0.004313868266446309
process_cpu_percent{ProCreateTime="1590569813000",ProName="sh",ProPid="58402",ProUserName="root",proCwd="/"} 4.144629229832814e-07
process_cpu_percent{ProCreateTime="1590569825000",ProName="sh",ProPid="58983",ProUserName="root",proCwd="/"} 8.289272135018644e-07
process_cpu_percent{ProCreateTime="1590569988000",ProName="bash",ProPid="61746",ProUserName="root",proCwd="/"} 5.526305827029192e-07
process_cpu_percent{ProCreateTime="1590569988000",ProName="containerd-shim",ProPid="61727",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/db17674cdab1a8c5434305c13c625de6f08915364ba890d56461fbf30a12a05c"} 0.004496202440392491
process_cpu_percent{ProCreateTime="1590570002000",ProName="sh",ProPid="61829",ProUserName="root",proCwd="/"} 4.144737377884647e-07
process_cpu_percent{ProCreateTime="1590574576000",ProName="dnstest",ProPid="65631",ProUserName="root",proCwd=""} 0.0927374169695384
process_cpu_percent{ProCreateTime="1590574658000",ProName="dnstest",ProPid="65729",ProUserName="root",proCwd=""} 0.09139927078733581
process_cpu_percent{ProCreateTime="1592042508000",ProName="crond",ProPid="4673",ProUserName="root",proCwd="/"} 0.02154403111312633
process_cpu_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63636",ProUserName="root",proCwd="/usr/server/nginx/conf"} 0
process_cpu_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63637",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 7.628573687801391e-06
process_cpu_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63638",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 9.808166128440362e-06
process_cpu_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63639",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 1.2169391266401677e-05
process_cpu_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63640",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 1.5075514493671684e-05
process_cpu_percent{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171477",ProUserName="root",proCwd="/root"} 1.8081896473494233e-05
process_cpu_percent{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171482",ProUserName="root",proCwd="/root"} 25.260992193291557
process_cpu_percent{ProCreateTime="1596974755000",ProName="node_exporter",ProPid="58868",ProUserName="iflyweb",proCwd="/"} 4.10834481968035
process_cpu_percent{ProCreateTime="1597806547000",ProName="sshd",ProPid="194110",ProUserName="root",proCwd="/"} 0.007098495511585657
process_cpu_percent{ProCreateTime="1597806550000",ProName="bash",ProPid="194114",ProUserName="root",proCwd="/root/press"} 0.0058190310241015
process_cpu_percent{ProCreateTime="1597806584000",ProName="process_exporter",ProPid="194177",ProUserName="root",proCwd="/usr/server/process"} 951.2386216656873
process_cpu_percent{ProCreateTime="1597806704000",ProName="python3",ProPid="194391",ProUserName="root",proCwd="/root/press"} 1.9515607442083533
process_cpu_percent{ProCreateTime="1597808069000",ProName="pickup",ProPid="195669",ProUserName="postfix",proCwd="/var/spool/postfix"} 0
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 14377.03
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 65536
# HELP process_mem_percent Get the mem usage rate of a single process
# TYPE process_mem_percent gauge
process_mem_percent{ProCreateTime="1590397599000",ProName="kthreadd",ProPid="2",ProUserName="root",proCwd="/"} 0
process_mem_percent{ProCreateTime="1590397599000",ProName="systemd",ProPid="1",ProUserName="root",proCwd="/"} 0.0735170990228653
process_mem_percent{ProCreateTime="1590397606000",ProName="lvmetad",ProPid="5841",ProUserName="root",proCwd="/"} 0
process_mem_percent{ProCreateTime="1590397606000",ProName="systemd-journald",ProPid="5816",ProUserName="root",proCwd="/"} 0.060036078095436096
process_mem_percent{ProCreateTime="1590397606000",ProName="systemd-udevd",ProPid="5850",ProUserName="root",proCwd="/"} 0.00015788580640219152
process_mem_percent{ProCreateTime="1590397608000",ProName="abrt-watch-log",ProPid="11526",ProUserName="root",proCwd="/"} 7.59066388127394e-05
process_mem_percent{ProCreateTime="1590397608000",ProName="abrtd",ProPid="11501",ProUserName="root",proCwd="/"} 0.0003491705283522606
process_mem_percent{ProCreateTime="1590397608000",ProName="dbus-daemon",ProPid="11584",ProUserName="dbus",proCwd="/"} 0.0009017708362080157
process_mem_percent{ProCreateTime="1590397608000",ProName="gssproxy",ProPid="11533",ProUserName="root",proCwd="/"} 0
process_mem_percent{ProCreateTime="1590397608000",ProName="irqbalance",ProPid="11504",ProUserName="root",proCwd="/"} 0.00040685958811081946
process_mem_percent{ProCreateTime="1590397608000",ProName="polkitd",ProPid="11566",ProUserName="polkitd",proCwd="/"} 0.002504918957129121
process_mem_percent{ProCreateTime="1590397608000",ProName="rpcbind",ProPid="11513",ProUserName="rpc",proCwd="/"} 3.947145160054788e-05
process_mem_percent{ProCreateTime="1590397608000",ProName="systemd-logind",ProPid="11564",ProUserName="root",proCwd="/"} 0.0009807137539610267
process_mem_percent{ProCreateTime="1590397613000",ProName="login",ProPid="11959",ProUserName="root",proCwd="/"} 3.339891918585636e-05
process_mem_percent{ProCreateTime="1590397613000",ProName="master",ProPid="12391",ProUserName="root",proCwd="/var/spool/postfix"} 0.0008349730051122606
process_mem_percent{ProCreateTime="1590397613000",ProName="qmgr",ProPid="12411",ProUserName="postfix",proCwd="/var/spool/postfix"} 0.0007469213451258838
process_mem_percent{ProCreateTime="1590397613000",ProName="sshd",ProPid="11863",ProUserName="root",proCwd="/"} 0.0009260610095225275
process_mem_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11877",ProUserName="zabbix",proCwd="/"} 6.07253105044947e-06
process_mem_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11887",ProUserName="zabbix",proCwd="/"} 0.0024775925558060408
process_mem_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11888",ProUserName="zabbix",proCwd="/"} 0.0009473148384131491
process_mem_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11889",ProUserName="zabbix",proCwd="/"} 0.0009503511246293783
process_mem_percent{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11890",ProUserName="zabbix",proCwd="/"} 0.0005374189931899309
process_mem_percent{ProCreateTime="1590397634000",ProName="bash",ProPid="21198",ProUserName="root",proCwd="/root"} 4.55439840152394e-05
process_mem_percent{ProCreateTime="1590399253000",ProName="rsyslogd",ProPid="22525",ProUserName="root",proCwd="/"} 0.03648984059691429
process_mem_percent{ProCreateTime="1590399383000",ProName="virtlogd",ProPid="22700",ProUserName="root",proCwd="/"} 0.0008380092913284898
process_mem_percent{ProCreateTime="1590560304000",ProName="auditd",ProPid="173542",ProUserName="root",proCwd="/"} 0.00034613427123986185
process_mem_percent{ProCreateTime="1590560330000",ProName="containerd",ProPid="173638",ProUserName="root",proCwd="/"} 0.023370135575532913
process_mem_percent{ProCreateTime="1590560330000",ProName="dockerd",ProPid="173643",ProUserName="root",proCwd="/"} 0.06451153010129929
process_mem_percent{ProCreateTime="1590560569000",ProName="containerd-shim",ProPid="175001",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/6fed7cc0f26af9f20adefd34aa51bce91d220852bd896bbf0023f18318903e52"} 0.005650490056723356
process_mem_percent{ProCreateTime="1590560569000",ProName="sh",ProPid="175020",ProUserName="root",proCwd="/"} 0
process_mem_percent{ProCreateTime="1590565963000",ProName="bash",ProPid="194392",ProUserName="root",proCwd="/"} 0.00020342979405540973
process_mem_percent{ProCreateTime="1590565963000",ProName="containerd-shim",ProPid="194362",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/4f4390d2ec436320b4bfec2dbab2e199af067afb8049b5ec768ada052f252a32"} 0.00378015055321157
process_mem_percent{ProCreateTime="1590569813000",ProName="containerd-shim",ProPid="58380",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/98f4825dfb5f8c71e0b4339634b251dc8c82687cfefca5ed37bf83d26e2c6108"} 0.004627268761396408
process_mem_percent{ProCreateTime="1590569813000",ProName="sh",ProPid="58402",ProUserName="root",proCwd="/"} 8.197916758945212e-05
process_mem_percent{ProCreateTime="1590569825000",ProName="sh",ProPid="58983",ProUserName="root",proCwd="/"} 0.0001062692899722606
process_mem_percent{ProCreateTime="1590569988000",ProName="bash",ProPid="61746",ProUserName="root",proCwd="/"} 0.0003248804132454097
process_mem_percent{ProCreateTime="1590569988000",ProName="containerd-shim",ProPid="61727",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/db17674cdab1a8c5434305c13c625de6f08915364ba890d56461fbf30a12a05c"} 0.004924822598695755
process_mem_percent{ProCreateTime="1590570002000",ProName="sh",ProPid="61829",ProUserName="root",proCwd="/"} 6.679783837171271e-05
process_mem_percent{ProCreateTime="1590574576000",ProName="dnstest",ProPid="65631",ProUserName="root",proCwd=""} 0
process_mem_percent{ProCreateTime="1590574658000",ProName="dnstest",ProPid="65729",ProUserName="root",proCwd=""} 0
process_mem_percent{ProCreateTime="1592042508000",ProName="crond",ProPid="4673",ProUserName="root",proCwd="/"} 0.0007438850589096546
process_mem_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63636",ProUserName="root",proCwd="/usr/server/nginx/conf"} 3.036265525224735e-06
process_mem_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63637",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 0.002772110514342785
process_mem_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63638",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 0.0014148997142910957
process_mem_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63639",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 0.0015120601747184992
process_mem_percent{ProCreateTime="1592302480000",ProName="nginx",ProPid="63640",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 0.005522966850548983
process_mem_percent{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171477",ProUserName="root",proCwd="/root"} 0.007287037093192339
process_mem_percent{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171482",ProUserName="root",proCwd="/root"} 0.014203649945557117
process_mem_percent{ProCreateTime="1596974755000",ProName="node_exporter",ProPid="58868",ProUserName="iflyweb",proCwd="/"} 0.03277648612856865
process_mem_percent{ProCreateTime="1597806547000",ProName="sshd",ProPid="194110",ProUserName="root",proCwd="/"} 0.004612087272107601
process_mem_percent{ProCreateTime="1597806550000",ProName="bash",ProPid="194114",ProUserName="root",proCwd="/root/press"} 0.001694236183539033
process_mem_percent{ProCreateTime="1597806584000",ProName="process_exporter",ProPid="194177",ProUserName="root",proCwd="/usr/server/process"} 0.044769734144210815
process_mem_percent{ProCreateTime="1597806704000",ProName="python3",ProPid="194391",ProUserName="root",proCwd="/root/press"} 0.015293669886887074
process_mem_percent{ProCreateTime="1597808069000",ProName="pickup",ProPid="195669",ProUserName="postfix",proCwd="/var/spool/postfix"} 0.003154679900035262
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 29
# HELP process_open_files_num Number of open files per process
# TYPE process_open_files_num gauge
process_open_files_num{ProCreateTime="1590397599000",ProName="kthreadd",ProPid="2",ProUserName="root",proCwd="/"} 0
process_open_files_num{ProCreateTime="1590397599000",ProName="systemd",ProPid="1",ProUserName="root",proCwd="/"} 57
process_open_files_num{ProCreateTime="1590397606000",ProName="lvmetad",ProPid="5841",ProUserName="root",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397606000",ProName="systemd-journald",ProPid="5816",ProUserName="root",proCwd="/"} 27
process_open_files_num{ProCreateTime="1590397606000",ProName="systemd-udevd",ProPid="5850",ProUserName="root",proCwd="/"} 12
process_open_files_num{ProCreateTime="1590397608000",ProName="abrt-watch-log",ProPid="11526",ProUserName="root",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397608000",ProName="abrtd",ProPid="11501",ProUserName="root",proCwd="/"} 10
process_open_files_num{ProCreateTime="1590397608000",ProName="dbus-daemon",ProPid="11584",ProUserName="dbus",proCwd="/"} 13
process_open_files_num{ProCreateTime="1590397608000",ProName="gssproxy",ProPid="11533",ProUserName="root",proCwd="/"} 11
process_open_files_num{ProCreateTime="1590397608000",ProName="irqbalance",ProPid="11504",ProUserName="root",proCwd="/"} 3
process_open_files_num{ProCreateTime="1590397608000",ProName="polkitd",ProPid="11566",ProUserName="polkitd",proCwd="/"} 12
process_open_files_num{ProCreateTime="1590397608000",ProName="rpcbind",ProPid="11513",ProUserName="rpc",proCwd="/"} 9
process_open_files_num{ProCreateTime="1590397608000",ProName="systemd-logind",ProPid="11564",ProUserName="root",proCwd="/"} 18
process_open_files_num{ProCreateTime="1590397613000",ProName="login",ProPid="11959",ProUserName="root",proCwd="/"} 1
process_open_files_num{ProCreateTime="1590397613000",ProName="master",ProPid="12391",ProUserName="root",proCwd="/var/spool/postfix"} 90
process_open_files_num{ProCreateTime="1590397613000",ProName="qmgr",ProPid="12411",ProUserName="postfix",proCwd="/var/spool/postfix"} 12
process_open_files_num{ProCreateTime="1590397613000",ProName="sshd",ProPid="11863",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11877",ProUserName="zabbix",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11887",ProUserName="zabbix",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11888",ProUserName="zabbix",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11889",ProUserName="zabbix",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11890",ProUserName="zabbix",proCwd="/"} 5
process_open_files_num{ProCreateTime="1590397634000",ProName="bash",ProPid="21198",ProUserName="root",proCwd="/root"} 4
process_open_files_num{ProCreateTime="1590399253000",ProName="rsyslogd",ProPid="22525",ProUserName="root",proCwd="/"} 17
process_open_files_num{ProCreateTime="1590399383000",ProName="virtlogd",ProPid="22700",ProUserName="root",proCwd="/"} 11
process_open_files_num{ProCreateTime="1590560304000",ProName="auditd",ProPid="173542",ProUserName="root",proCwd="/"} 10
process_open_files_num{ProCreateTime="1590560330000",ProName="containerd",ProPid="173638",ProUserName="root",proCwd="/"} 17
process_open_files_num{ProCreateTime="1590560330000",ProName="dockerd",ProPid="173643",ProUserName="root",proCwd="/"} 51
process_open_files_num{ProCreateTime="1590560569000",ProName="containerd-shim",ProPid="175001",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/6fed7cc0f26af9f20adefd34aa51bce91d220852bd896bbf0023f18318903e52"} 20
process_open_files_num{ProCreateTime="1590560569000",ProName="sh",ProPid="175020",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590565963000",ProName="bash",ProPid="194392",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590565963000",ProName="containerd-shim",ProPid="194362",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/4f4390d2ec436320b4bfec2dbab2e199af067afb8049b5ec768ada052f252a32"} 20
process_open_files_num{ProCreateTime="1590569813000",ProName="containerd-shim",ProPid="58380",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/98f4825dfb5f8c71e0b4339634b251dc8c82687cfefca5ed37bf83d26e2c6108"} 29
process_open_files_num{ProCreateTime="1590569813000",ProName="sh",ProPid="58402",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590569825000",ProName="sh",ProPid="58983",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590569988000",ProName="bash",ProPid="61746",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590569988000",ProName="containerd-shim",ProPid="61727",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/db17674cdab1a8c5434305c13c625de6f08915364ba890d56461fbf30a12a05c"} 29
process_open_files_num{ProCreateTime="1590570002000",ProName="sh",ProPid="61829",ProUserName="root",proCwd="/"} 4
process_open_files_num{ProCreateTime="1590574576000",ProName="dnstest",ProPid="65631",ProUserName="root",proCwd=""} 0
process_open_files_num{ProCreateTime="1590574658000",ProName="dnstest",ProPid="65729",ProUserName="root",proCwd=""} 0
process_open_files_num{ProCreateTime="1592042508000",ProName="crond",ProPid="4673",ProUserName="root",proCwd="/"} 6
process_open_files_num{ProCreateTime="1592302480000",ProName="nginx",ProPid="63636",ProUserName="root",proCwd="/usr/server/nginx/conf"} 15
process_open_files_num{ProCreateTime="1592302480000",ProName="nginx",ProPid="63637",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 13
process_open_files_num{ProCreateTime="1592302480000",ProName="nginx",ProPid="63638",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 13
process_open_files_num{ProCreateTime="1592302480000",ProName="nginx",ProPid="63639",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 13
process_open_files_num{ProCreateTime="1592302480000",ProName="nginx",ProPid="63640",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 13
process_open_files_num{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171477",ProUserName="root",proCwd="/root"} 7
process_open_files_num{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171482",ProUserName="root",proCwd="/root"} 8
process_open_files_num{ProCreateTime="1596974755000",ProName="node_exporter",ProPid="58868",ProUserName="iflyweb",proCwd="/"} 6
process_open_files_num{ProCreateTime="1597806547000",ProName="sshd",ProPid="194110",ProUserName="root",proCwd="/"} 11
process_open_files_num{ProCreateTime="1597806550000",ProName="bash",ProPid="194114",ProUserName="root",proCwd="/root/press"} 4
process_open_files_num{ProCreateTime="1597806584000",ProName="process_exporter",ProPid="194177",ProUserName="root",proCwd="/usr/server/process"} 27
process_open_files_num{ProCreateTime="1597806704000",ProName="python3",ProPid="194391",ProUserName="root",proCwd="/root/press"} 11
process_open_files_num{ProCreateTime="1597808069000",ProName="pickup",ProPid="195669",ProUserName="postfix",proCwd="/var/spool/postfix"} 12
# HELP process_read_bytes Number of bytes read by the process
# TYPE process_read_bytes gauge
process_read_bytes{ProCreateTime="1590397599000",ProName="kthreadd",ProPid="2",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590397599000",ProName="systemd",ProPid="1",ProUserName="root",proCwd="/"} 4.219614976e+11
process_read_bytes{ProCreateTime="1590397606000",ProName="lvmetad",ProPid="5841",ProUserName="root",proCwd="/"} 90112
process_read_bytes{ProCreateTime="1590397606000",ProName="systemd-journald",ProPid="5816",ProUserName="root",proCwd="/"} 6.2406656e+07
process_read_bytes{ProCreateTime="1590397606000",ProName="systemd-udevd",ProPid="5850",ProUserName="root",proCwd="/"} 9.102336e+06
process_read_bytes{ProCreateTime="1590397608000",ProName="abrt-watch-log",ProPid="11526",ProUserName="root",proCwd="/"} 2.3785472e+07
process_read_bytes{ProCreateTime="1590397608000",ProName="abrtd",ProPid="11501",ProUserName="root",proCwd="/"} 1.73312e+06
process_read_bytes{ProCreateTime="1590397608000",ProName="dbus-daemon",ProPid="11584",ProUserName="dbus",proCwd="/"} 3.9161856e+07
process_read_bytes{ProCreateTime="1590397608000",ProName="gssproxy",ProPid="11533",ProUserName="root",proCwd="/"} 4096
process_read_bytes{ProCreateTime="1590397608000",ProName="irqbalance",ProPid="11504",ProUserName="root",proCwd="/"} 9.9999744e+08
process_read_bytes{ProCreateTime="1590397608000",ProName="polkitd",ProPid="11566",ProUserName="polkitd",proCwd="/"} 9.6858112e+07
process_read_bytes{ProCreateTime="1590397608000",ProName="rpcbind",ProPid="11513",ProUserName="rpc",proCwd="/"} 4.1201664e+07
process_read_bytes{ProCreateTime="1590397608000",ProName="systemd-logind",ProPid="11564",ProUserName="root",proCwd="/"} 5.791744e+07
process_read_bytes{ProCreateTime="1590397613000",ProName="login",ProPid="11959",ProUserName="root",proCwd="/"} 966656
process_read_bytes{ProCreateTime="1590397613000",ProName="master",ProPid="12391",ProUserName="root",proCwd="/var/spool/postfix"} 9.1021312e+07
process_read_bytes{ProCreateTime="1590397613000",ProName="qmgr",ProPid="12411",ProUserName="postfix",proCwd="/var/spool/postfix"} 1.2599296e+07
process_read_bytes{ProCreateTime="1590397613000",ProName="sshd",ProPid="11863",ProUserName="root",proCwd="/"} 1.5422202368e+10
process_read_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11877",ProUserName="zabbix",proCwd="/"} 4096
process_read_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11887",ProUserName="zabbix",proCwd="/"} 3.166212096e+09
process_read_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11888",ProUserName="zabbix",proCwd="/"} 1.0489856e+09
process_read_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11889",ProUserName="zabbix",proCwd="/"} 8.41756672e+08
process_read_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11890",ProUserName="zabbix",proCwd="/"} 1.2850176e+09
process_read_bytes{ProCreateTime="1590397634000",ProName="bash",ProPid="21198",ProUserName="root",proCwd="/root"} 557056
process_read_bytes{ProCreateTime="1590399253000",ProName="rsyslogd",ProPid="22525",ProUserName="root",proCwd="/"} 8.49131008e+08
process_read_bytes{ProCreateTime="1590399383000",ProName="virtlogd",ProPid="22700",ProUserName="root",proCwd="/"} 741376
process_read_bytes{ProCreateTime="1590560304000",ProName="auditd",ProPid="173542",ProUserName="root",proCwd="/"} 2.9786624e+07
process_read_bytes{ProCreateTime="1590560330000",ProName="containerd",ProPid="173638",ProUserName="root",proCwd="/"} 3.804659712e+09
process_read_bytes{ProCreateTime="1590560330000",ProName="dockerd",ProPid="173643",ProUserName="root",proCwd="/"} 3.12438784e+09
process_read_bytes{ProCreateTime="1590560569000",ProName="containerd-shim",ProPid="175001",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/6fed7cc0f26af9f20adefd34aa51bce91d220852bd896bbf0023f18318903e52"} 2.56651264e+08
process_read_bytes{ProCreateTime="1590560569000",ProName="sh",ProPid="175020",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590565963000",ProName="bash",ProPid="194392",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590565963000",ProName="containerd-shim",ProPid="194362",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/4f4390d2ec436320b4bfec2dbab2e199af067afb8049b5ec768ada052f252a32"} 5.9547648e+08
process_read_bytes{ProCreateTime="1590569813000",ProName="containerd-shim",ProPid="58380",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/98f4825dfb5f8c71e0b4339634b251dc8c82687cfefca5ed37bf83d26e2c6108"} 5.14945024e+08
process_read_bytes{ProCreateTime="1590569813000",ProName="sh",ProPid="58402",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590569825000",ProName="sh",ProPid="58983",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590569988000",ProName="bash",ProPid="61746",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590569988000",ProName="containerd-shim",ProPid="61727",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/db17674cdab1a8c5434305c13c625de6f08915364ba890d56461fbf30a12a05c"} 5.10615552e+08
process_read_bytes{ProCreateTime="1590570002000",ProName="sh",ProPid="61829",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1590574576000",ProName="dnstest",ProPid="65631",ProUserName="root",proCwd=""} 0
process_read_bytes{ProCreateTime="1590574658000",ProName="dnstest",ProPid="65729",ProUserName="root",proCwd=""} 0
process_read_bytes{ProCreateTime="1592042508000",ProName="crond",ProPid="4673",ProUserName="root",proCwd="/"} 2.2810624e+08
process_read_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63636",ProUserName="root",proCwd="/usr/server/nginx/conf"} 0
process_read_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63637",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 1.204224e+06
process_read_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63638",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 2.809856e+06
process_read_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63639",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 2.052096e+06
process_read_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63640",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 3.887104e+06
process_read_bytes{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171477",ProUserName="root",proCwd="/root"} 6.275817472e+09
process_read_bytes{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171482",ProUserName="root",proCwd="/root"} 8.79678464e+10
process_read_bytes{ProCreateTime="1596974755000",ProName="node_exporter",ProPid="58868",ProUserName="iflyweb",proCwd="/"} 1.1669504e+07
process_read_bytes{ProCreateTime="1597806547000",ProName="sshd",ProPid="194110",ProUserName="root",proCwd="/"} 0
process_read_bytes{ProCreateTime="1597806550000",ProName="bash",ProPid="194114",ProUserName="root",proCwd="/root/press"} 7.487488e+06
process_read_bytes{ProCreateTime="1597806584000",ProName="process_exporter",ProPid="194177",ProUserName="root",proCwd="/usr/server/process"} 0
process_read_bytes{ProCreateTime="1597806704000",ProName="python3",ProPid="194391",ProUserName="root",proCwd="/root/press"} 0
process_read_bytes{ProCreateTime="1597808069000",ProName="pickup",ProPid="195669",ProUserName="postfix",proCwd="/var/spool/postfix"} 0
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 6.039552e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.59780658466e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 4.45591552e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes -1
# HELP process_write_bytes Number of bytes write by the process
# TYPE process_write_bytes gauge
process_write_bytes{ProCreateTime="1590397599000",ProName="kthreadd",ProPid="2",ProUserName="root",proCwd="/"} 2.985984e+06
process_write_bytes{ProCreateTime="1590397599000",ProName="systemd",ProPid="1",ProUserName="root",proCwd="/"} 3.842825928704e+12
process_write_bytes{ProCreateTime="1590397606000",ProName="lvmetad",ProPid="5841",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397606000",ProName="systemd-journald",ProPid="5816",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397606000",ProName="systemd-udevd",ProPid="5850",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="abrt-watch-log",ProPid="11526",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="abrtd",ProPid="11501",ProUserName="root",proCwd="/"} 2.347008e+06
process_write_bytes{ProCreateTime="1590397608000",ProName="dbus-daemon",ProPid="11584",ProUserName="dbus",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="gssproxy",ProPid="11533",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="irqbalance",ProPid="11504",ProUserName="root",proCwd="/"} 5.6438784e+07
process_write_bytes{ProCreateTime="1590397608000",ProName="polkitd",ProPid="11566",ProUserName="polkitd",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="rpcbind",ProPid="11513",ProUserName="rpc",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397608000",ProName="systemd-logind",ProPid="11564",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590397613000",ProName="login",ProPid="11959",ProUserName="root",proCwd="/"} 16384
process_write_bytes{ProCreateTime="1590397613000",ProName="master",ProPid="12391",ProUserName="root",proCwd="/var/spool/postfix"} 598016
process_write_bytes{ProCreateTime="1590397613000",ProName="qmgr",ProPid="12411",ProUserName="postfix",proCwd="/var/spool/postfix"} 0
process_write_bytes{ProCreateTime="1590397613000",ProName="sshd",ProPid="11863",ProUserName="root",proCwd="/"} 2.53586989056e+11
process_write_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11877",ProUserName="zabbix",proCwd="/"} 4096
process_write_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11887",ProUserName="zabbix",proCwd="/"} 5.38505216e+08
process_write_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11888",ProUserName="zabbix",proCwd="/"} 7.7205504e+07
process_write_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11889",ProUserName="zabbix",proCwd="/"} 1.204224e+08
process_write_bytes{ProCreateTime="1590397613000",ProName="zabbix_agentd",ProPid="11890",ProUserName="zabbix",proCwd="/"} 16384
process_write_bytes{ProCreateTime="1590397634000",ProName="bash",ProPid="21198",ProUserName="root",proCwd="/root"} 0
process_write_bytes{ProCreateTime="1590399253000",ProName="rsyslogd",ProPid="22525",ProUserName="root",proCwd="/"} 1.143390208e+09
process_write_bytes{ProCreateTime="1590399383000",ProName="virtlogd",ProPid="22700",ProUserName="root",proCwd="/"} 8192
process_write_bytes{ProCreateTime="1590560304000",ProName="auditd",ProPid="173542",ProUserName="root",proCwd="/"} 6.57915904e+08
process_write_bytes{ProCreateTime="1590560330000",ProName="containerd",ProPid="173638",ProUserName="root",proCwd="/"} 1.11439872e+08
process_write_bytes{ProCreateTime="1590560330000",ProName="dockerd",ProPid="173643",ProUserName="root",proCwd="/"} 2.88292864e+08
process_write_bytes{ProCreateTime="1590560569000",ProName="containerd-shim",ProPid="175001",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/6fed7cc0f26af9f20adefd34aa51bce91d220852bd896bbf0023f18318903e52"} 1.175552e+06
process_write_bytes{ProCreateTime="1590560569000",ProName="sh",ProPid="175020",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590565963000",ProName="bash",ProPid="194392",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590565963000",ProName="containerd-shim",ProPid="194362",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/4f4390d2ec436320b4bfec2dbab2e199af067afb8049b5ec768ada052f252a32"} 32768
process_write_bytes{ProCreateTime="1590569813000",ProName="containerd-shim",ProPid="58380",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/98f4825dfb5f8c71e0b4339634b251dc8c82687cfefca5ed37bf83d26e2c6108"} 45056
process_write_bytes{ProCreateTime="1590569813000",ProName="sh",ProPid="58402",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590569825000",ProName="sh",ProPid="58983",ProUserName="root",proCwd="/"} 20480
process_write_bytes{ProCreateTime="1590569988000",ProName="bash",ProPid="61746",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590569988000",ProName="containerd-shim",ProPid="61727",ProUserName="root",proCwd="/run/containerd/io.containerd.runtime.v1.linux/moby/db17674cdab1a8c5434305c13c625de6f08915364ba890d56461fbf30a12a05c"} 45056
process_write_bytes{ProCreateTime="1590570002000",ProName="sh",ProPid="61829",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1590574576000",ProName="dnstest",ProPid="65631",ProUserName="root",proCwd=""} 1.22978304e+08
process_write_bytes{ProCreateTime="1590574658000",ProName="dnstest",ProPid="65729",ProUserName="root",proCwd=""} 1.23072512e+08
process_write_bytes{ProCreateTime="1592042508000",ProName="crond",ProPid="4673",ProUserName="root",proCwd="/"} 2.38239744e+08
process_write_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63636",ProUserName="root",proCwd="/usr/server/nginx/conf"} 4096
process_write_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63637",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 188416
process_write_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63638",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 225280
process_write_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63639",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 454656
process_write_bytes{ProCreateTime="1592302480000",ProName="nginx",ProPid="63640",ProUserName="iflyweb",proCwd="/usr/server/nginx/conf"} 417792
process_write_bytes{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171477",ProUserName="root",proCwd="/root"} 6.272937984e+09
process_write_bytes{ProCreateTime="1596536106000",ProName="python3.6",ProPid="171482",ProUserName="root",proCwd="/root"} 0
process_write_bytes{ProCreateTime="1596974755000",ProName="node_exporter",ProPid="58868",ProUserName="iflyweb",proCwd="/"} 0
process_write_bytes{ProCreateTime="1597806547000",ProName="sshd",ProPid="194110",ProUserName="root",proCwd="/"} 0
process_write_bytes{ProCreateTime="1597806550000",ProName="bash",ProPid="194114",ProUserName="root",proCwd="/root/press"} 12288
process_write_bytes{ProCreateTime="1597806584000",ProName="process_exporter",ProPid="194177",ProUserName="root",proCwd="/usr/server/process"} 81920
process_write_bytes{ProCreateTime="1597806704000",ProName="python3",ProPid="194391",ProUserName="root",proCwd="/root/press"} 454656
process_write_bytes{ProCreateTime="1597808069000",ProName="pickup",ProPid="195669",ProUserName="postfix",proCwd="/var/spool/postfix"} 0
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 9
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 10825
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0

```
