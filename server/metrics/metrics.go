package metrics

import (
	"github.com/kataras/iris"
	"his6/base/router"
)

func init() {
	router.RegisterGetHandler("/metrics", metricsHandle)
}

func metricsHandle(ctx iris.Context) {
	//s := "# HELP go_gc_duration_seconds A summary of the GC invocation durations.\n" +
	//	"# TYPE go_gc_duration_seconds summary\n" +
	//	"go_gc_duration_seconds{quantile=\"0\"} 0\n" +
	//	"go_gc_duration_seconds{quantile=\"0.25\"} 0\n" +
	//	"go_gc_duration_seconds{quantile=\"0.5\"} 0\n" +
	//	"go_gc_duration_seconds{quantile=\"0.75\"} 0\n" +
	//	"go_gc_duration_seconds{quantile=\"1\"} 0\n" +
	//	"go_gc_duration_seconds_sum 0\n" +
	//	"go_gc_duration_seconds_count 8\n" +
	//	"# HELP go_goroutines Number of goroutines that currently exist.\n" +
	//	"# TYPE go_goroutines gauge\n" +
	//	"go_goroutines 11\n" +
	//	"# HELP go_info Information about the Go environment.\n" +
	//	"# TYPE go_info gauge\n" +
	//	"go_info{version=\"go1.13.8\"} 1\n" +
	//	"# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.\n" +
	//	"# TYPE go_memstats_alloc_bytes gauge\n" +
	//	"go_memstats_alloc_bytes 1.0381424e+07\n" +
	//	"# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.\n" +
	//	"# TYPE go_memstats_alloc_bytes_total counter\n" +
	//	"go_memstats_alloc_bytes_total 2.62868e+07\n" +
	//	"# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.\n" +
	//	"# TYPE go_memstats_buck_hash_sys_bytes gauge\n" +
	//	"go_memstats_buck_hash_sys_bytes 1.447845e+06\n" +
	//	"# HELP go_memstats_frees_total Total number of frees.\n" +
	//	"# TYPE go_memstats_frees_total counter\n" +
	//	"go_memstats_frees_total 66865\n" +
	//	"# HELP go_memstats_gc_cpu_fraction The fraction of this program's available CPU time used by the GC since the program started.\n" +
	//	"# TYPE go_memstats_gc_cpu_fraction gauge\n" +
	//	"go_memstats_gc_cpu_fraction 2.133352958846288e-06\n" +
	//	"# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.\n" +
	//	"# TYPE go_memstats_gc_sys_bytes gauge\n" +
	//	"go_memstats_gc_sys_bytes 754176\n" +
	//	"# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.\n" +
	//	"# TYPE go_memstats_heap_alloc_bytes gauge\n" +
	//	"go_memstats_heap_alloc_bytes 1.0381424e+07\n" +
	//	"# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.\n" +
	//	"# TYPE go_memstats_heap_idle_bytes gauge\n" +
	//	"go_memstats_heap_idle_bytes 4.087808e+06\n" +
	//	"# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.\n" +
	//	"# TYPE go_memstats_heap_inuse_bytes gauge\n" +
	//	"go_memstats_heap_inuse_bytes 1.1771904e+07\n" +
	//	"# HELP go_memstats_heap_objects Number of allocated objects.\n" +
	//	"# TYPE go_memstats_heap_objects gauge\n" +
	//	"go_memstats_heap_objects 19458\n" +
	//	"# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.\n" +
	//	"# TYPE go_memstats_heap_released_bytes gauge\n" +
	//	"go_memstats_heap_released_bytes 3.76832e+06\n" +
	//	"# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.\n" +
	//	"# TYPE go_memstats_heap_sys_bytes gauge\n" +
	//	"go_memstats_heap_sys_bytes 1.5859712e+07\n" +
	//	"# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.\n" +
	//	"# TYPE go_memstats_last_gc_time_seconds gauge\n" +
	//	"go_memstats_last_gc_time_seconds 1.5822893242281368e+09\n" +
	//	"# HELP go_memstats_lookups_total Total number of pointer lookups.\n" +
	//	"# TYPE go_memstats_lookups_total counter\n" +
	//	"go_memstats_lookups_total 0\n" +
	//	"# HELP go_memstats_mallocs_total Total number of mallocs.\n" +
	//	"# TYPE go_memstats_mallocs_total counter\n" +
	//	"go_memstats_mallocs_total 86323\n" +
	//	"# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.\n" +
	//	"# TYPE go_memstats_mcache_inuse_bytes gauge\n" +
	//	"go_memstats_mcache_inuse_bytes 13632\n" +
	//	"# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.\n" +
	//	"# TYPE go_memstats_mcache_sys_bytes gauge\n" +
	//	"go_memstats_mcache_sys_bytes 16384\n" +
	//	"# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.\n" +
	//	"# TYPE go_memstats_mspan_inuse_bytes gauge\n" +
	//	"go_memstats_mspan_inuse_bytes 53176\n" +
	//	"# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.\n" +
	//	"# TYPE go_memstats_mspan_sys_bytes gauge\n" +
	//	"go_memstats_mspan_sys_bytes 65536\n" +
	//	"# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.\n" +
	//	"# TYPE go_memstats_next_gc_bytes gauge\n" +
	//	"go_memstats_next_gc_bytes 1.1361744e+07\n" +
	//	"# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.\n" +
	//	"# TYPE go_memstats_other_sys_bytes gauge\n" +
	//	"go_memstats_other_sys_bytes 1.532243e+06\n" +
	//	"# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.\n" +
	//	"# TYPE go_memstats_stack_inuse_bytes gauge\n" +
	//	"go_memstats_stack_inuse_bytes 917504\n" +
	//	"# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.\n" +
	//	"# TYPE go_memstats_stack_sys_bytes gauge\n" +
	//	"go_memstats_stack_sys_bytes 917504\n" +
	//	"# HELP go_memstats_sys_bytes Number of bytes obtained from system.\n" +
	//	"# TYPE go_memstats_sys_bytes gauge\n" +
	//	"go_memstats_sys_bytes 2.05934e+07\n" +
	//	"# HELP go_threads Number of OS threads created.\n" +
	//	"# TYPE go_threads gauge\n" +
	//	"go_threads 14\n" +
	//	"# HELP http_request_duration_seconds How long it took to process the request, partitioned by status code, method and HTTP path.\n" +
	//	"# TYPE http_request_duration_seconds histogram\n" +
	//	"http_request_duration_seconds_bucket{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\",le=\"0.3\"} 77\n" +
	//	"http_request_duration_seconds_bucket{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\",le=\"1.2\"} 77\n" +
	//	"http_request_duration_seconds_bucket{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\",le=\"5\"} 77\n" +
	//	"http_request_duration_seconds_bucket{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\",le=\"+Inf\"} 77\n" +
	//	"http_request_duration_seconds_sum{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\"} 0.05846669999999999\n" +
	//	"http_request_duration_seconds_count{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\"} 77\n" +
	//	"# HELP http_requests_total How many HTTP requests processed, partitioned by status code, method and HTTP path\n" +
	//	"# TYPE http_requests_total counter\n" +
	//	"http_requests_total{code=\"200\",method=\"GET\",path=\"/metrics\",service=\"serviceName\"} 77\n" +
	//	"# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.\n" +
	//	"# TYPE process_cpu_seconds_total counter\n" +
	//	"process_cpu_seconds_total 0.09375\n" +
	//	"# HELP process_max_fds Maximum number of open file descriptors.\n" +
	//	"# TYPE process_max_fds gauge\n" +
	//	"process_max_fds 1.6777216e+07\n" +
	//	"# HELP process_open_fds Number of open file descriptors.\n" +
	//	"# TYPE process_open_fds gauge\n" +
	//	"process_open_fds 154\n" +
	//	"# HELP process_resident_memory_bytes Resident memory size in bytes.\n" +
	//	"# TYPE process_resident_memory_bytes gauge\n" +
	//	"process_resident_memory_bytes 2.3097344e+07\n" +
	//	"# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.\n" +
	//	"# TYPE process_start_time_seconds gauge\n" +
	//	"process_start_time_seconds 1.582288625e+09\n" +
	//	"# HELP process_virtual_memory_bytes Virtual memory size in bytes.\n" +
	//	"# TYPE process_virtual_memory_bytes gauge\n" +
	//	"process_virtual_memory_bytes 2.7262976e+07\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.1\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.2\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.4\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"1\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"3\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"8\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"20\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"60\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"120\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"+Inf\"} 5\n" +
	//	"# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.\n" +
	//	"# TYPE promhttp_metric_handler_requests_in_flight gauge\n" +
	//	"promhttp_metric_handler_requests_in_flight 1\n" +
	//	"# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.\n" +
	//	"# TYPE promhttp_metric_handler_requests_total counter\n" +
	//	"promhttp_metric_handler_requests_total{code=\"200\"} 77\n" +
	//	"promhttp_metric_handler_requests_total{code=\"500\"} 0\n" +
	//	"promhttp_metric_handler_requests_total{code=\"503\"} 0\n"

	s := router.GetMetrics()
	//s := "prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.1\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.2\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"0.4\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"1\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"3\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"8\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"20\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"60\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"120\"} 5\n" +
	//	"prometheus_http_request_duration_seconds_bucket{handler=\"/api/v1/label/:name/values\",le=\"+Inf\"}5"

	ctx.Text(s)
}
