package kubernetes

import (
	"github.com/perses/community-mixins/pkg/dashboards"
	"github.com/perses/community-mixins/pkg/promql"
	"github.com/perses/perses/go-sdk/panel"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	"github.com/perses/plugins/prometheus/sdk/go/query"

	commonSdk "github.com/perses/perses/go-sdk/common"
	tablePanel "github.com/perses/plugins/table/sdk/go"
)

func WorkloadNamespaceCPUUsageQuota(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("CPU Quota",
		panel.Description("Shows the CPU requests, limits, and usage of workloads in a namespace in tabular format."),
		tablePanel.Table(
			tablePanel.Transform([]commonSdk.Transform{
				{
					Kind: commonSdk.MergeSeriesKind,
					Spec: commonSdk.MergeSeriesSpec{
						Disabled: false,
					},
				},
			}),
			tablePanel.WithColumnSettings([]tablePanel.ColumnSettings{
				{
					Name:   "workload",
					Header: "Workload",
					Align:  tablePanel.LeftAlign,
				},
				{
					Name:   "workload_type",
					Header: "Type",
					Align:  tablePanel.LeftAlign,
				},
				{
					Name:   "value #1",
					Header: "Running Pods",
					Align:  tablePanel.RightAlign,
				},
				{
					Name:   "value #2",
					Header: "CPU Usage",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.DecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name:   "value #3",
					Header: "CPU Requests",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.DecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name:   "value #4",
					Header: "CPU Requests %",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.PercentDecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name:   "value #5",
					Header: "CPU Limits",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.DecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name:   "value #6",
					Header: "CPU Limits %",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.PercentDecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name: "timestamp",
					Hide: true,
				},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"count(namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload, workload_type)",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  "+GetNodeNSCPUSecondsRecordingRule()+"{cluster=\"$cluster\", namespace=\"$namespace\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  kube_pod_container_resource_requests{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"cpu\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  "+GetNodeNSCPUSecondsRecordingRule()+"{cluster=\"$cluster\", namespace=\"$namespace\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n/sum(\n  kube_pod_container_resource_requests{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"cpu\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  kube_pod_container_resource_limits{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"cpu\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  "+GetNodeNSCPUSecondsRecordingRule()+"{cluster=\"$cluster\", namespace=\"$namespace\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n/sum(\n  kube_pod_container_resource_limits{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"cpu\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func WorkloadNamespaceMemoryUsageQuota(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Memory Quota",
		panel.Description("Shows the memory requests, limits, and usage of workloads in a namespace in tabular format."),
		tablePanel.Table(
			tablePanel.Transform([]commonSdk.Transform{
				{
					Kind: commonSdk.MergeSeriesKind,
					Spec: commonSdk.MergeSeriesSpec{
						Disabled: false,
					},
				},
			}),
			tablePanel.WithColumnSettings([]tablePanel.ColumnSettings{
				{
					Name:   "workload",
					Header: "Workload",
					Align:  tablePanel.LeftAlign,
				},
				{
					Name:   "workload_type",
					Header: "Type",
					Align:  tablePanel.LeftAlign,
				},
				{
					Name:   "value #1",
					Header: "Running Pods",
					Align:  tablePanel.RightAlign,
				},
				{
					Name:   "value #2",
					Header: "Memory Usage",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.BytesUnit,
					},
				},
				{
					Name:   "value #3",
					Header: "Memory Requests",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.BytesUnit,
					},
				},
				{
					Name:   "value #4",
					Header: "Memory Requests %",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.PercentDecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name:   "value #5",
					Header: "Memory Limits",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.BytesUnit,
					},
				},
				{
					Name:   "value #6",
					Header: "Memory Limits %",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit:          &dashboards.PercentDecimalUnit,
						DecimalPlaces: 4,
					},
				},
				{
					Name: "timestamp",
					Hide: true,
				},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"count(namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload, workload_type)",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n    container_memory_working_set_bytes{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", container!=\"\", image!=\"\"}\n  * on(namespace,pod)\n    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  kube_pod_container_resource_requests{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"memory\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n    container_memory_working_set_bytes{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", container!=\"\", image!=\"\"}\n  * on(namespace,pod)\n    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n/sum(\n  kube_pod_container_resource_requests{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"memory\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n  kube_pod_container_resource_limits{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"memory\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(\n    container_memory_working_set_bytes{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", container!=\"\", image!=\"\"}\n  * on(namespace,pod)\n    group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n/sum(\n  kube_pod_container_resource_limits{"+GetKubeStateMetricsMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\", resource=\"memory\"}\n* on(namespace,pod)\n  group_left(workload, workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}\n) by (workload, workload_type)\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func WorkloadNamespaceCurrentNetworkUsage(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Current Network Usage",
		panel.Description("Shows the current network usage of the namespace by workloads."),
		tablePanel.Table(
			tablePanel.Transform([]commonSdk.Transform{
				{
					Kind: commonSdk.MergeSeriesKind,
					Spec: commonSdk.MergeSeriesSpec{
						Disabled: false,
					},
				},
			}),
			tablePanel.WithColumnSettings([]tablePanel.ColumnSettings{
				{
					Name:   "workload",
					Header: "Workload",
					Align:  tablePanel.LeftAlign,
				},
				{
					Name:   "value #1",
					Header: "Current Receive Bandwidth",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.BytesPerSecondsUnit,
					},
				},
				{
					Name:   "value #2",
					Header: "Current Transmit Bandwidth",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.BytesPerSecondsUnit,
					},
				},
				{
					Name:   "value #3",
					Header: "Rate of Received Packets",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.PacketsPerSecondsUnit,
					},
				},
				{
					Name:   "value #4",
					Header: "Rate of Transmitted Packets",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.PacketsPerSecondsUnit,
					},
				},
				{
					Name:   "value #5",
					Header: "Rate of Received Packets Dropped",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.PacketsPerSecondsUnit,
					},
				},
				{
					Name:   "value #6",
					Header: "Rate of Transmitted Packets Dropped",
					Align:  tablePanel.RightAlign,
					Format: &commonSdk.Format{
						Unit: &dashboards.PacketsPerSecondsUnit,
					},
				},
				{
					Name: "timestamp",
					Hide: true,
				},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_receive_bytes_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_transmit_bytes_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_receive_packets_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_transmit_packets_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_receive_packets_dropped_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"(sum(rate(container_network_transmit_packets_dropped_total{"+GetCAdvisorMatcher()+", cluster=\"$cluster\", namespace=\"$namespace\"}[$__rate_interval])\n* on (namespace,pod)\ngroup_left(workload,workload_type) namespace_workload_pod:kube_pod_owner:relabel{cluster=\"$cluster\", namespace=\"$namespace\", workload_type=~\"$type\"}) by (workload))\n",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}
