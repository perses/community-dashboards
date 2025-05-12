package nodeexporter

import (
	"github.com/perses/community-dashboards/pkg/dashboards"
	panels "github.com/perses/community-dashboards/pkg/panels/node_exporter"
	"github.com/perses/community-dashboards/pkg/promql"
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
	labelValuesVar "github.com/perses/plugins/prometheus/sdk/go/variable/label-values"
)

func withClusterCPU(datasource string, clusterLabelMatcher promql.LabelMatcher, instanceLabelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("CPU",
		panelgroup.PanelsPerLine(2),
		panels.ClusterNodeCPUUsagePercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
		panels.ClusterNodeCPUSaturationPercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
	)
}

func withClusterMemory(datasource string, clusterLabelMatcher promql.LabelMatcher, instanceLabelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Memory",
		panelgroup.PanelsPerLine(2),
		panels.ClusterNodeMemoryUsagePercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
		panels.ClusterNodeMemorySaturationPercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
	)
}

func withClusterNetwork(datasource string, clusterLabelMatcher promql.LabelMatcher, instanceLabelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Network",
		panelgroup.PanelsPerLine(2),
		panels.ClusterNodeNetworkUsageBytes(datasource, clusterLabelMatcher, instanceLabelMatcher),
		panels.ClusterNodeNetworkSaturationBytes(datasource, clusterLabelMatcher, instanceLabelMatcher),
	)
}

func withClusterDiskIO(datasource string, clusterLabelMatcher promql.LabelMatcher, instanceLabelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Disk IO",
		panelgroup.PanelsPerLine(2),
		panels.ClusterNodeDiskUsagePercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
		panels.ClusterNodeDiskSaturationPercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
	)
}

func withClusterDiskSpace(datasource string, clusterLabelMatcher promql.LabelMatcher, instanceLabelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Disk Space",
		panelgroup.PanelsPerLine(1),
		panels.ClusterNodeDiskSpacePercentage(datasource, clusterLabelMatcher, instanceLabelMatcher),
	)
}

func BuildNodeExporterClusterUseMethod(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcher(clusterLabelName)
	instanceLabelMatcher := promql.LabelMatcher{
		Name:  "instance",
		Value: "$instance",
		Type:  "=~",
	}
	return dashboards.NewDashboardResult(
		dashboard.New("node-exporter-cluster-use-method",
			dashboard.ProjectName(project),
			dashboard.Name("Node Exporter / USE Method / Cluster"),
			dashboards.AddClusterVariable(datasource, clusterLabelName, "node_uname_info{job='node', sysname!='Darwin'}"),
			dashboard.AddVariable("instance",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("instance",
						dashboards.AddVariableDatasource(datasource),
						labelValuesVar.Matchers(
							promql.SetLabelMatchers(
								"node_uname_info{job='node', sysname!='Darwin'}",
								[]promql.LabelMatcher{clusterLabelMatcher},
							)),
					),
					listVar.DisplayName("instance"),
					listVar.AllowAllValue(true),
					listVar.AllowMultiple(true),
				),
			),
			withClusterCPU(datasource, clusterLabelMatcher, instanceLabelMatcher),
			withClusterMemory(datasource, clusterLabelMatcher, instanceLabelMatcher),
			withClusterNetwork(datasource, clusterLabelMatcher, instanceLabelMatcher),
			withClusterDiskIO(datasource, clusterLabelMatcher, instanceLabelMatcher),
			withClusterDiskSpace(datasource, clusterLabelMatcher, instanceLabelMatcher),
		),
	).Component("node-exporter")
}
