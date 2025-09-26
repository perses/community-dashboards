package istio

import (
	"maps"

	"github.com/perses/community-dashboards/pkg/promql"
	promqlbuilder "github.com/perses/promql-builder"
	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/matrix"
	"github.com/perses/promql-builder/vector"
	"github.com/prometheus/prometheus/promql/parser"
)

var IstioCommonPanelQueries = map[string]parser.Expr{
	// istio control plane
	"IstioPushSize": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_xds_config_size_bytes_bucket"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	).By("le"),
	"IstioPushTime": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_xds_push_time_bucket"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	).By("le"),
	"IstioConnectionsClientReported": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("envoy_cluster_upstream_cx_active"),
			vector.WithLabelMatchers(
				label.New("cluster_name").Equal("xds-grpc"),
			),
		),
	),
	"IstioConnections": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("pilot_xds"),
		),
	),
	"IstioCPUUsage": promqlbuilder.Sum(
		promqlbuilder.IRate(
			matrix.New(
				vector.New(
					vector.WithMetricName("container_cpu_usage_seconds_total"),
					vector.WithLabelMatchers(
						label.New("container").Equal("discovery"),
						label.New("pod").EqualRegexp("istiod-.*"),
					),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("pod"),
	"IstioEventsReg": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_k8s_reg_events"),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("type", "event"),
	"IstioEventsCfg": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_k8s_cfg_events"),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("type", "event"),
	"IstioEventsPilot": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_push_triggers"),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("type"),
	"IstioGoroutines": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("go_goroutines"),
					vector.WithLabelMatchers(
						label.New("app").Equal("istiod"),
					),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("pod"),
	"IstioMemoryAllocationsBytesTotal": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("go_memstats_alloc_bytes_total"),
					vector.WithLabelMatchers(
						label.New("app").Equal("istiod"),
					),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("pod"),
	"IstioMemoryAllocationsMallocsTotal": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("go_memstats_mallocs_total"),
					vector.WithLabelMatchers(
						label.New("app").Equal("istiod"),
					),
				),
				matrix.WithRangeAsString("$__rate_interval"),
			),
		),
	).By("pod"),
	"IstioMemoryUsageWorkingSetBytes": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("container_memory_working_set_bytes"),
			vector.WithLabelMatchers(
				label.New("container").Equal("discovery"),
				label.New("pod").EqualRegexp("istiod-.*"),
			),
		),
	).By("pod"),
	"IstioMemoryUsageInuseBytes": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("go_memstats_stack_inuse_bytes"),
			vector.WithLabelMatchers(
				label.New("app").Equal("istiod"),
			),
		),
	).By("pod"),
	"IstioMemoryUsageHeapInuseBytes": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("go_memstats_heap_inuse_bytes"),
			vector.WithLabelMatchers(
				label.New("app").Equal("istiod"),
			),
		),
	).By("pod"),
	"IstioMemoryUsageHeapAllocBytes": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("go_memstats_heap_alloc_bytes"),
			vector.WithLabelMatchers(
				label.New("app").Equal("istiod"),
			),
		),
	).By("pod"),
	"IstioPilotVersions": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("istio_build"),
			vector.WithLabelMatchers(
				label.New("component").Equal("pilot"),
			),
		),
	).By("tag"),
	"IstioPushErrorsRejects": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("pilot_total_xds_rejects"),
			vector.WithLabelMatchers(
				label.New("component").Equal("pilot"),
			),
		),
	).By("type"),
	"IstioPushErrors": vector.New(
		vector.WithMetricName("pilot_total_xds_internal_errors"),
	),
	"IstioInjectionSucessTotal": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("sidecar_injection_success_total"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	),
	"IstioInjectionFailureTotal": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("sidecar_injection_failure_total"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	),
	"IstioValidationPassed": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("galley_validation_passed"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	),
	"IstioValidationFailed": promqlbuilder.Sum(
		promqlbuilder.Rate(
			matrix.New(
				vector.New(
					vector.WithMetricName("galley_validation_failed"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	),
	"IstioXDSPushes": promqlbuilder.Sum(
		promqlbuilder.IRate(
			matrix.New(
				vector.New(
					vector.WithMetricName("pilot_xds_pushes"),
				),
				matrix.WithRangeAsString("1m"),
			),
		),
	).By("type"),
	// istio mesh
	"IstioHTTPGRPCWorkloads": promqlbuilder.LabelJoin(
		promqlbuilder.Sum(
			promqlbuilder.Rate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("source|waypoint"),
						),
					),
					matrix.WithRangeAsString("$__rate_interval"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "destination_service"),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioHTTPGRPCWorkloads50": promqlbuilder.LabelJoin(
		promqlbuilder.HistogramQuantile(0.5,
			promql.SumByRate(
				"istio_request_duration_milliseconds_bucket",
				[]string{"le", "destination_workload", "destination_workload_namespace"},
				label.New("reporter").EqualRegexp("source|waypoint"),
			),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioHTTPGRPCWorkloads90": promqlbuilder.LabelJoin(
		promqlbuilder.HistogramQuantile(0.9,
			promql.SumByRate(
				"istio_request_duration_milliseconds_bucket",
				[]string{"le", "destination_workload", "destination_workload_namespace"},
				label.New("reporter").EqualRegexp("source|waypoint"),
			),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioHTTPGRPCWorkloads99": promqlbuilder.LabelJoin(
		promqlbuilder.HistogramQuantile(0.99,
			promql.SumByRate(
				"istio_request_duration_milliseconds_bucket",
				[]string{"le", "destination_workload", "destination_workload_namespace"},
				label.New("reporter").EqualRegexp("source|waypoint"),
			),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioHTTPGRPCWorkloadsReqTotal": promqlbuilder.LabelJoin(
		promqlbuilder.Div(
			promql.SumByRate(
				"istio_requests_total",
				[]string{"destination_workload", "destination_workload_namespace"},
				label.New("reporter").EqualRegexp("source|waypoint"),
				label.New("response_code").NotEqualRegexp("5.."),
			),
			promql.SumByRate(
				"istio_requests_total",
				[]string{"destination_workload", "destination_workload_namespace"},
				label.New("reporter").EqualRegexp("source|waypoint"),
			),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioTCPServicesBytesRecv": promqlbuilder.LabelJoin(
		promql.SumByRate(
			"istio_tcp_received_bytes_total",
			[]string{"destination_workload", "destination_workload_namespace", "destination_service"},
			label.New("reporter").EqualRegexp("source|waypoint"),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioTCPServicesBytesSent": promqlbuilder.LabelJoin(
		promql.SumByRate(
			"istio_tcp_sent_bytes_total",
			[]string{"destination_workload", "destination_workload_namespace", "destination_service"},
			label.New("reporter").EqualRegexp("source|waypoint"),
		),
		"destination_workload_var",
		"destination_workload",
		"destination_workload_namespace",
	),
	"IstioGlobalRequestVolume": promqlbuilder.Round(
		promql.SumRate(
			"istio_requests_total",
			label.New("reporter").EqualRegexp("source|waypoint"),
		),
		0.01,
	),
	"IstionGlobalSuccessRate": promqlbuilder.Div(
		promql.SumRate(
			"istio_requests_total",
			label.New("reporter").EqualRegexp("source|waypoint"),
			label.New("response_code").NotEqualRegexp("5.."),
		),
		promql.SumRate(
			"istio_requests_total",
			label.New("reporter").EqualRegexp("source|waypoint"),
		),
	),
	"IstioGlobal4xxRate": promqlbuilder.Or(
		promqlbuilder.Round(
			promql.SumRate(
				"istio_requests_total",
				label.New("reporter").EqualRegexp("source|waypoint"),
				label.New("response_code").EqualRegexp("4.."),
			),
			0.01,
		),
		promqlbuilder.Vector(0),
	),
	"IstioGlobal5xxRate": promqlbuilder.Or(
		promqlbuilder.Round(
			promql.SumRate(
				"istio_requests_total",
				label.New("reporter").EqualRegexp("source|waypoint"),
				label.New("response_code").EqualRegexp("5.."),
			),
			0.01,
		),
		promqlbuilder.Vector(0),
	),
	"IstioComponentVersions": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("istio_build"),
		),
	).By("component", "tag"),
	// istio performance
	"IstioVCPUPer1kRPSIngressGateway": promqlbuilder.Div(
		promql.SumiRate(
			"container_cpu_usage_seconds_total",
			label.New("pod").EqualRegexp("istio-ingressgateway-.*"),
			label.New("container").Equal("istio-proxy"),
		),
		promqlbuilder.Div(
			promqlbuilder.Round(
				promql.SumiRate(
					"istio_requests_total",
					label.New("source_workload").Equal("istio-ingressgateway"),
					label.New("reporter").Equal("source"),
				),
				0.001,
			),
			&parser.NumberLiteral{Val: 1000},
		),
	),
	"IstioVCPUPer1kRPSProxy": promqlbuilder.Div(
		promqlbuilder.Div(
			promql.SumiRate(
				"container_cpu_usage_seconds_total",
				label.New("namespace").NotEqual("istio-system"),
				label.New("container").Equal("istio-proxy"),
			),
			promqlbuilder.Div(
				promqlbuilder.Round(
					promql.SumiRate(
						"istio_requests_total",
					),
					0.001,
				),
				&parser.NumberLiteral{Val: 1000},
			),
		),
		promqlbuilder.Gtr(
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_requests_total"),
							vector.WithLabelMatchers(
								label.New("source_workload").Equal("istio-ingressgateway"),
							),
						),
						matrix.WithRangeAsVariable("$__rate_interval"),
					),
				),
			),
			&parser.NumberLiteral{Val: 10},
		).Bool(),
	),
	"IstioVCPUIngressGateway": promql.SumRate(
		"container_cpu_usage_seconds_total",
		label.New("pod").EqualRegexp("istio-ingressgateway-.*"),
		label.New("container").Equal("istio-proxy"),
	),
	"IstioVCPUProxy": promql.SumRate(
		"container_cpu_usage_seconds_total",
		label.New("namespace").NotEqual("istio-system"),
		label.New("container").Equal("istio-proxy"),
	),
	"IstioPerformanceMemoryUsageIngressGateway": promqlbuilder.Div(
		promqlbuilder.Sum(
			vector.New(
				vector.WithMetricName("container_memory_working_set_bytes"),
				vector.WithLabelMatchers(
					label.New("pod").EqualRegexp("istio-ingressgateway-.*"),
				),
			),
		),
		promqlbuilder.Count(
			vector.New(
				vector.WithMetricName("container_memory_working_set_bytes"),
				vector.WithLabelMatchers(
					label.New("pod").EqualRegexp("istio-ingressgateway-.*"),
					label.New("container").NotEqual("POD"),
				),
			),
		),
	),
	"IstioPerformanceMemoryUsageProxy": promqlbuilder.Div(
		promqlbuilder.Sum(
			vector.New(
				vector.WithMetricName("container_memory_working_set_bytes"),
				vector.WithLabelMatchers(
					label.New("namespace").NotEqual("istio-system"),
					label.New("container").Equal("istio-proxy"),
				),
			),
		),
		promqlbuilder.Count(
			vector.New(
				vector.WithMetricName("container_memory_working_set_bytes"),
				vector.WithLabelMatchers(
					label.New("namespace").NotEqual("istio-system"),
					label.New("container").Equal("istio-proxy"),
				),
			),
		),
	),
	"IstioBytesTransferredIngressGateway": promql.SumiRate(
		"istio_response_bytes_sum",
		label.New("source_workload").Equal("istio-ingressgateway"),
		label.New("reporter").Equal("source"),
	),
	"IstioBytesTransferredProxy": promqlbuilder.Add(
		promql.SumiRate(
			"istio_response_bytes_sum",
			label.New("namespace").NotEqual("istio-system"),
			label.New("reporter").Equal("source"),
		),
		promql.SumiRate(
			"istio_request_bytes_sum",
			label.New("namespace").NotEqual("istio-system"),
			label.New("reporter").Equal("source"),
		),
	),
	"IstioComponentsByVersion": promql.SumBy(
		"istio_build",
		[]string{"component", "tag"},
	),
	"IstioProxyMemory": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("container_memory_working_set_bytes"),
			vector.WithLabelMatchers(
				label.New("container").Equal("istio-proxy"),
			),
		),
	),
	"IstioProxyVCPU": promql.SumRate(
		"container_cpu_usage_seconds_total",
		label.New("container").Equal("istio-proxy"),
	),
	"IstioProxyDisk": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("container_fs_usage_bytes"),
			vector.WithLabelMatchers(
				label.New("container").Equal("istio-proxy"),
			),
		),
	),
	"IstiodMemoryVirtual": vector.New(
		vector.WithMetricName("process_virtual_memory_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryResident": vector.New(
		vector.WithMetricName("process_resident_memory_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryGoMemStats": vector.New(
		vector.WithMetricName("go_memstats_heap_sys_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryHeapAllocBytes": vector.New(
		vector.WithMetricName("go_memstats_heap_alloc_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryAllocBytes": vector.New(
		vector.WithMetricName("go_memstats_alloc_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryInuseBytes": vector.New(
		vector.WithMetricName("go_memstats_heap_inuse_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryStackInuseBytes": vector.New(
		vector.WithMetricName("go_memstats_stack_inuse_bytes"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodMemoryContainerTotal": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("container_memory_working_set_bytes"),
			vector.WithLabelMatchers(
				label.New("container").EqualRegexp("discovery|istio-proxy"),
				label.New("pod").EqualRegexp("istiod-.*"),
			),
		),
	),
	"IstiodMemoryContainer": vector.New(
		vector.WithMetricName("container_memory_working_set_bytes"),
		vector.WithLabelMatchers(
			label.New("container").EqualRegexp("discovery|istio-proxy"),
			label.New("pod").EqualRegexp("istiod-.*"),
		),
	),
	"IstiodVCPUTotal": promql.SumRate(
		"container_cpu_usage_seconds_total",
		label.New("container").EqualRegexp("discovery|istio-proxy"),
		label.New("pod").EqualRegexp("istiod-.*"),
	),
	"IstiodVCPUContainer": promql.SumByRate(
		"container_cpu_usage_seconds_total",
		[]string{"container"},
		label.New("container").EqualRegexp("discovery|istio-proxy"),
		label.New("pod").EqualRegexp("istiod-.*"),
	),
	"IstiodVCPUPilot": promqlbuilder.IRate(
		matrix.New(
			vector.New(
				vector.WithMetricName("process_cpu_seconds_total"),
				vector.WithLabelMatchers(
					label.New("app").Equal("istiod"),
				),
			),
			matrix.WithRangeAsVariable("$__rate_interval"),
		),
	),
	"IstiodDiskOpenFDs": vector.New(
		vector.WithMetricName("process_open_fds"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
	"IstiodDiskContainerFSUsage": vector.New(
		vector.WithMetricName("container_fs_usage_bytes"),
		vector.WithLabelMatchers(
			label.New("container").EqualRegexp("discovery|istio-proxy"),
			label.New("pod").EqualRegexp("istiod-.*"),
		),
	),
	"IstiodGoroutines": vector.New(
		vector.WithMetricName("go_goroutines"),
		vector.WithLabelMatchers(
			label.New("app").Equal("istiod"),
		),
	),
}

// OverrideIstionPanelQueries overrides the IstioCommonPanelQueries global.
// Refer to panel queries in the map, that you'd like to override.
// The convention of naming followed, is to use Panel function name (with _suffix, in case panel has multiple queries)
func OverrideIstionPanelQueries(queries map[string]parser.Expr) {
	maps.Copy(IstioCommonPanelQueries, queries)
}
