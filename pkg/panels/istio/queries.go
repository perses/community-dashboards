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
	// istio service additional
	"IncomingRequestDurationByClient50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.5,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.5,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClient90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.9,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.9,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClient95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClient99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClientNonmTLS50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.5,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClientNonmTLS90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClientNonmTLS95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByClientNonmTLS99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClient50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClient90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClient95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClient99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClientNonmTLS50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClientNonmTLS90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClientNonmTLS95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByClientNonmTLS99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_bytes_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("source_workload").EqualRegexp("$srcwl"),
									label.New("source_workload_namespace").EqualRegexp("$srcns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("source_workload", "source_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_bytes_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("source_workload").EqualRegexp("$srcwl"),
								label.New("source_workload_namespace").EqualRegexp("$srcns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("source_workload", "source_workload_namespace", "le"),
		),
	),
	"IstioResponseSizeByClient50": promqlbuilder.HistogramQuantile(0.5,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClient90": promqlbuilder.HistogramQuantile(0.9,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClient95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClient99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClientNonmTLS50": promqlbuilder.HistogramQuantile(0.5,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClientNonmTLS90": promqlbuilder.HistogramQuantile(0.9,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClientNonmTLS95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"IstioResponseSizeByClientNonmTLS99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"BytesReceivedFromTCPClient": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_received_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		0.001,
	),
	"BytesReceivedFromTCPClientNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_received_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		0.001,
	),
	"BytesSentToTCPClient": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_sent_bytes_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		0.001,
	),
	"BytesSentToTCPClientNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_sent_bytes_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		0.001,
	),
	// istio service workloads
	"IstioIncomingRequestsByService": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").Equal("destination"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "response_code"),
		0.001,
	),
	"IstioIncomingRequestsByServiceNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").Equal("destination"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "response_code"),
		0.001,
	),
	"IncomingSuccessRateByService": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqualRegexp("5.*"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
	),
	"IncomingSuccessRateByServiceNonmTLS": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqualRegexp("5.*"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
	),
	"IncomingRequestDurationByService50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByService90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByService95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByService99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").Equal("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").Equal("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByServiceNonmTLS50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByServiceNonmTLS90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByServiceNonmTLS95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestDurationByServiceNonmTLS99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("connection_security_policy").NotEqual("mutual_tls"),
									label.New("destination_service").EqualRegexp("$service"),
									label.New("destination_workload").EqualRegexp("$dstwl"),
									label.New("destination_workload_namespace").EqualRegexp("$dstns"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("destination_workload", "destination_workload_namespace", "le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("connection_security_policy").NotEqual("mutual_tls"),
								label.New("destination_service").EqualRegexp("$service"),
								label.New("destination_workload").EqualRegexp("$dstwl"),
								label.New("destination_workload_namespace").EqualRegexp("$dstns"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("destination_workload", "destination_workload_namespace", "le"),
		),
	),
	"IncomingRequestSizeByService50": promqlbuilder.HistogramQuantile(0.50,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByService90": promqlbuilder.HistogramQuantile(0.90,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByService95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByService99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByServiceNonmTLS50": promqlbuilder.HistogramQuantile(0.50,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByServiceNonmTLS90": promqlbuilder.HistogramQuantile(0.90,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByServiceNonmTLS95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"IncomingRequestSizeByServiceNonmTLS99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByService50": promqlbuilder.HistogramQuantile(0.50,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByService90": promqlbuilder.HistogramQuantile(0.90,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByService95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByService99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByServiceNonmTLS50": promqlbuilder.HistogramQuantile(0.50,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByServiceNonmTLS90": promqlbuilder.HistogramQuantile(0.90,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByServiceNonmTLS95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"ResponseSizeByServiceNonmTLS99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_response_bytes_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace", "le"),
	),
	"BytesReceivedFromTCPService": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_received_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"BytesReceivedFromTCPServiceNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_received_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"BytesSentToTCPService": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_sent_bytes_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"BytesSentToTCPServiceNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_sent_bytes_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	// istio service
	"ClientRequestVolume": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		0.001,
	),
	"ClientSuccessRate": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
							label.New("response_code").NotEqual("5.*"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
	),
	"ClientRequestDuration50": promqlbuilder.HistogramQuantile(0.50,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"ClientRequestDuration90": promqlbuilder.HistogramQuantile(0.90,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"ClientRequestDuration95": promqlbuilder.HistogramQuantile(0.95,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"ClientRequestDuration99": promqlbuilder.HistogramQuantile(0.99,
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("source_workload", "source_workload_namespace", "le"),
	),
	"ServerRequestVolume": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"ServiceTCPBytesReceived": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_received_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("1m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"ServiceTCPBytesSent": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_tcp_sent_bytes_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("destination_workload").EqualRegexp("$dstwl"),
							label.New("destination_workload_namespace").EqualRegexp("$dstns"),
						),
					),
					matrix.WithRangeAsVariable("$__rate_interval"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"ClientRequestVolumeStat": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("destination_workload", "destination_workload_namespace"),
		0.001,
	),
	"ClientSuccessRateStat": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqualRegexp("5.*"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		),
		promqlbuilder.Or(
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_requests_total"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("5m"),
					),
				),
			),
			promqlbuilder.Vector(1),
		).On(),
	),
	"ClientRequestDurationChart50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"ClientRequestDurationChart90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"ClientRequestDurationChart99": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.99,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").EqualRegexp("$qrep"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.99,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").EqualRegexp("$qrep"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"TCPReceivedBytesStat": promqlbuilder.Sum(
		promqlbuilder.IRate(
			matrix.New(
				vector.New(
					vector.WithMetricName("istio_tcp_received_bytes_total"),
					vector.WithLabelMatchers(
						label.New("reporter").EqualRegexp("$qrep"),
						label.New("destination_service").EqualRegexp("$service"),
					),
				),
				matrix.WithRangeAsVariable("1m"),
			),
		),
	),
	"ServerRequestVolumeStat": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		),
		0.001,
	),
	"ServerSuccessRateStat": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").Equal("destination"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqual("5.*"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		),
		promqlbuilder.Or(
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_requests_total"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("5m"),
					),
				),
			),
			promqlbuilder.Vector(1),
		).On(),
	),
	"ServerRequestDurationChart50": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.50,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.50,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"ServerRequestDurationChart90": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.90,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.90,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"ServerRequestDurationChart95": promqlbuilder.Or(
		promqlbuilder.Div(
			promqlbuilder.HistogramQuantile(0.95,
				promqlbuilder.Sum(
					promqlbuilder.IRate(
						matrix.New(
							vector.New(
								vector.WithMetricName("istio_request_duration_milliseconds_bucket"),
								vector.WithLabelMatchers(
									label.New("reporter").Equal("destination"),
									label.New("destination_service").EqualRegexp("$service"),
								),
							),
							matrix.WithRangeAsVariable("1m"),
						),
					),
				).By("le"),
			),
			&parser.NumberLiteral{Val: 1000},
		),
		promqlbuilder.HistogramQuantile(0.95,
			promqlbuilder.Sum(
				promqlbuilder.IRate(
					matrix.New(
						vector.New(
							vector.WithMetricName("istio_request_duration_seconds_bucket"),
							vector.WithLabelMatchers(
								label.New("reporter").Equal("destination"),
								label.New("destination_service").EqualRegexp("$service"),
							),
						),
						matrix.WithRangeAsVariable("1m"),
					),
				),
			).By("le"),
		),
	),
	"TCPSentBytesStat": promqlbuilder.Sum(
		promqlbuilder.IRate(
			matrix.New(
				vector.New(
					vector.WithMetricName("istio_tcp_sent_bytes_total"),
					vector.WithLabelMatchers(
						label.New("reporter").EqualRegexp("$qrep"),
						label.New("destination_service").EqualRegexp("$service"),
					),
				),
				matrix.WithRangeAsVariable("1m"),
			),
		),
	),
	"IncomingRequestsByClient": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "response_code"),
		0.001,
	),
	"IncomingRequestsByClientNonmTLS": promqlbuilder.Round(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace", "response_code"),
		0.001,
	),
	"IncomingSuccessRateByClient": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqual("5.*"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").Equal("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
	),
	"IncomingSuccessRateByClientNonmTLS": promqlbuilder.Div(
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("response_code").NotEqual("5.*"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
		promqlbuilder.Sum(
			promqlbuilder.IRate(
				matrix.New(
					vector.New(
						vector.WithMetricName("istio_requests_total"),
						vector.WithLabelMatchers(
							label.New("connection_security_policy").NotEqual("mutual_tls"),
							label.New("destination_service").EqualRegexp("$service"),
							label.New("reporter").EqualRegexp("$qrep"),
							label.New("source_workload").EqualRegexp("$srcwl"),
							label.New("source_workload_namespace").EqualRegexp("$srcns"),
						),
					),
					matrix.WithRangeAsVariable("5m"),
				),
			),
		).By("source_workload", "source_workload_namespace"),
	),
}

// OverrideIstionPanelQueries overrides the IstioCommonPanelQueries global.
// Refer to panel queries in the map, that you'd like to override.
// The convention of naming followed, is to use Panel function name (with _suffix, in case panel has multiple queries)
func OverrideIstionPanelQueries(queries map[string]parser.Expr) {
	maps.Copy(IstioCommonPanelQueries, queries)
}
