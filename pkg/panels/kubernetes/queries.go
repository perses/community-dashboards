package kubernetes

import (
	"maps"

	"github.com/perses/community-dashboards/pkg/promql"
	promqlbuilder "github.com/perses/promql-builder"
	"github.com/perses/promql-builder/label"
	"github.com/perses/promql-builder/vector"
	"github.com/prometheus/prometheus/promql/parser"
)

var KubernetesCommonPanelQueries = map[string]parser.Expr{
	// apiserver panels
	"APIServerAvailability": vector.New(
		vector.WithMetricName("apiserver_request:availability30d"),
		vector.WithLabelMatchers(
			label.New("verb").Equal("all"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerErrorBudget": promqlbuilder.Mul(
		&parser.NumberLiteral{Val: 100},
		promqlbuilder.Sub(
			vector.New(
				vector.WithMetricName("apiserver_request:availability30d"),
				vector.WithLabelMatchers(
					label.New("verb").Equal("all"),
					label.New("cluster").EqualRegexp("$cluster"),
				),
			),
			&parser.NumberLiteral{Val: 0.990000},
		),
	),
	"APIServerReadAvailability": vector.New(
		vector.WithMetricName("apiserver_request:availability30d"),
		vector.WithLabelMatchers(
			label.New("verb").Equal("read"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerReadSLIRequests": promql.SumBy(
		"code_resource:apiserver_request_total:rate5m",
		[]string{"code"},
		label.New("verb").Equal("read"),
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"APIServerReadSLIErrors": promqlbuilder.Div(
		promql.SumBy(
			"code_resource:apiserver_request_total:rate5m",
			[]string{"resource"},
			label.New("verb").Equal("read"),
			label.New("code").EqualRegexp("5.."),
			label.New("cluster").EqualRegexp("$cluster"),
		),
		promql.SumBy(
			"code_resource:apiserver_request_total:rate5m",
			[]string{"resource"},
			label.New("verb").Equal("read"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerReadSLIDuration": vector.New(
		vector.WithMetricName("cluster_quantile:apiserver_request_sli_duration_seconds:histogram_quantile"),
		vector.WithLabelMatchers(
			label.New("verb").Equal("read"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerWriteAvailability": vector.New(
		vector.WithMetricName("apiserver_request:availability30d"),
		vector.WithLabelMatchers(
			label.New("verb").Equal("write"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerWriteSLIRequests": promql.SumBy(
		"code_resource:apiserver_request_total:rate5m",
		[]string{"code"},
		label.New("verb").Equal("write"),
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"APIServerWriteSLIErrors": promqlbuilder.Div(
		promql.SumBy(
			"code_resource:apiserver_request_total:rate5m",
			[]string{"resource"},
			label.New("verb").Equal("write"),
			label.New("code").EqualRegexp("5.."),
			label.New("cluster").EqualRegexp("$cluster"),
		),
		promql.SumBy(
			"code_resource:apiserver_request_total:rate5m",
			[]string{"resource"},
			label.New("verb").Equal("write"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerWriteSLIDuration": vector.New(
		vector.WithMetricName("cluster_quantile:apiserver_request_sli_duration_seconds:histogram_quantile"),
		vector.WithLabelMatchers(
			label.New("verb").Equal("write"),
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"APIServerWorkQueueAddRate": promql.SumByRate(
		"workqueue_adds_total",
		[]string{"instance", "name"},
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("instance").EqualRegexp("$instance'"),
		label.New("job").Equal("kube-apiserver"),
	),
	"APIServerWorkQueueDepth": promql.SumByRate(
		"workqueue_depth",
		[]string{"instance", "name"},
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("instance").EqualRegexp("$instance'"),
		label.New("job").Equal("kube-apiserver"),
	),
	"APIServerWorkQueueLatency": promqlbuilder.HistogramQuantile(
		0.99,
		promql.SumByRate(
			"workqueue_queue_duration_seconds_bucket",
			[]string{"instance", "name", "le"},
			label.New("cluster").EqualRegexp("$cluster"),
			label.New("instance").EqualRegexp("$instance'"),
			label.New("job").Equal("kube-apiserver"),
		),
	),
	"ClusterCPUUsageQuotaPodOwn": promqlbuilder.Sum(
		vector.New(
			vector.WithMetricName("kube_pod_owner"),
			vector.WithLabelMatchers(
				label.New("job").Equal("kube-state-metrics'"),
				label.New("cluster").EqualRegexp("$cluster"),
			),
		),
	).By("namespace"),
	"ClusterCPUUsageQuotaNSwWorkload": promqlbuilder.Count(
		promqlbuilder.Avg(
			vector.New(
				vector.WithMetricName("namespace_workload_pod:kube_pod_owner:relabel"),
				vector.WithLabelMatchers(
					label.New("job").Equal("kube-state-metrics'"),
					label.New("cluster").EqualRegexp("$cluster"),
				),
			),
		).By("workload", "namespace"),
	).By("namespace"),
	"ClusterCPUUsageQuotaNodeNSCPU": promql.SumBy(
		"node_namespace_pod_container:container_cpu_usage_seconds_total:sum_irate",
		[]string{"namespace"},
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterCPUUsageQuotaNSKubePodContainer": promql.SumBy(
		"namespace_cpu:kube_pod_container_resource_requests:sum",
		[]string{"namespace"},
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterCPUUsageQuotaNSKubePodContainerDiv": promqlbuilder.Div(
		promql.SumBy(
			"namespace_cpu:kube_pod_container_resource_requests:sum",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
		promql.SumBy(
			"namespace_cpu:kube_pod_container_resource_requests:sum",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"ClusterCPUUsageQuotaNSCPUKubePodResources": promql.SumBy(
		"namespace_cpu:kube_pod_container_resource_limits:sum",
		[]string{"namespace"},
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterCPUUsageQuotaNSCPUKubePodResourcesDiv": promqlbuilder.Div(
		promql.SumBy(
			"node_namespace_pod_container:container_cpu_usage_seconds_total:sum_rate5m",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
		promql.SumBy(
			"namespace_cpu:kube_pod_container_resource_limits:sum",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"ClusterMemoryUsageQuotaPodOwner": promql.SumBy(
		"kube_pod_owner",
		[]string{"namespace"},
		label.New("job").Equal("kube-state-metrics"),
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterMemoryUsageQuotaNSWorkloadPodOwner": promqlbuilder.Count(
		promqlbuilder.Avg(
			vector.New(
				vector.WithMetricName("namespace_workload_pod:kube_pod_owner:relabel"),
				vector.WithLabelMatchers(
					label.New("cluster").EqualRegexp("$cluster"),
				),
			),
		).By("workload", "namespace"),
	).By("namespace"),
	"ClusterMemoryUsageQuotaContainerMem": promql.SumBy(
		"container_memory_rss",
		[]string{"namespace"},
		label.New("job").Equal("kube-state-metrics"),
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("cluster").NotEqualRegexp(""),
	),
	"ClusterMemoryUsageQuotaContainerResourceReqSum": promql.SumBy(
		"namespace_memory:kube_pod_container_resource_requests:sum",
		[]string{"namespace"},
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterMemoryUsageQuotaContainerResourceReqSumDiv": promqlbuilder.Div(
		promql.SumBy(
			"container_memory_rss",
			[]string{"namespace"},
			label.New("job").Equal("kube-state-metrics"),
			label.New("cluster").EqualRegexp("$cluster"),
			label.New("cluster").NotEqualRegexp(""),
		),
		promql.SumBy(
			"namespace_memory:kube_pod_container_resource_requests:sum",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"ClusterMemoryUsageQuotaContainerReqLimits": promql.SumBy(
		"namespace_memory:kube_pod_container_resource_limits:sum",
		[]string{"namespace"},
		label.New("cluster").EqualRegexp("$cluster"),
	),
	"ClusterMemoryUsageQuotaContainerReqLimitsDiv": promqlbuilder.Div(
		promql.SumBy(
			"container_memory_rss",
			[]string{"namespace"},
			label.New("job").Equal("kube-state-metrics"),
			label.New("cluster").EqualRegexp("$cluster"),
			label.New("cluster").NotEqualRegexp(""),
		),
		promql.SumBy(
			"namespace_memory:kube_pod_container_resource_limits:sum",
			[]string{"namespace"},
			label.New("cluster").EqualRegexp("$cluster"),
		),
	),
	"ClusterCurrentNetworkUsageBytesTotal": promql.SumByRate(
		"container_network_receive_bytes_total",
		[]string{"namespace"},
		label.New("job").Equal("kube-state-metrics"),
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("namespace").EqualRegexp(".+"),
	),
	"ClusterCurrentNetworkTrasmitBytesTotal": promql.SumByRate(
		"container_network_transmit_bytes_total",
		[]string{"namespace"},
		label.New("job").Equal("kube-state-metrics"),
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("namespace").EqualRegexp(".+"),
	),
	"ClusterCurrentNetworkReceivedTotal": promql.SumByRate(
		"container_network_receive_packets_total",
		[]string{"namespace"},
		label.New("job").Equal("kube-state-metrics"),
		label.New("cluster").EqualRegexp("$cluster"),
		label.New("namespace").EqualRegexp(".+"),
	),
}

// OverrideKubernetesPanelQueries overrides the KubernetesCommonPanelQueries global.
// Refer to panel queries in the map, that you'd like to override.
// The convention of naming followed, is to use Panel function name (with _suffix, in case panel has multiple queries)
func OverrideKubernetesPanelQueries(queries map[string]parser.Expr) {
	maps.Copy(KubernetesCommonPanelQueries, queries)
}
