package kubelet

import (
	"github.com/perses/community-dashboards/pkg/dashboards"
	panels "github.com/perses/community-dashboards/pkg/panels/kubernetes"
	"github.com/perses/community-dashboards/pkg/promql"
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	labelValuesVar "github.com/perses/perses/go-sdk/prometheus/variable/label-values"
	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
)

func withKubeletStats(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Kubelet Stats",
		panelgroup.PanelsPerLine(6),
		panels.RunningKubeletStat(datasource, labelMatcher),
		panels.RunningPodStat(datasource, labelMatcher),
		panels.RunningContainersStat(datasource, labelMatcher),
		panels.ActVolumeCountStat(datasource, labelMatcher),
		panels.DesiredVolumeCountStat(datasource, labelMatcher),
		panels.ConfigErrorCountStat(datasource, labelMatcher),
	)
}

func withKubeletOperations(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Operation Rate and Errors",
		panelgroup.PanelsPerLine(2),
		panels.RunningKubeletStat(datasource, labelMatcher),
		panels.RunningPodStat(datasource, labelMatcher),
	)
}

func withKubeletOperationsQuantile(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Operation Duration 99th quantile",
		panelgroup.PanelsPerLine(1),
		panels.OperationDurationQuantile(datasource, labelMatcher),
	)
}

func withPodStartRateAndDuration(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Pod Start Rate and Duration",
		panelgroup.PanelsPerLine(2),
		panels.PodStartRate(datasource, labelMatcher),
		panels.PodStartDuration(datasource, labelMatcher),
	)
}

func withStorageOperationsAndErrors(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Storage Operations Rate and Errors",
		panelgroup.PanelsPerLine(2),
		panels.StorageOperationRate(datasource, labelMatcher),
		panels.StorageOperationErrorRate(datasource, labelMatcher),
	)
}

func withStorageOperationsQuantile(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Storage Operation Duration 99th quantile",
		panelgroup.PanelsPerLine(1),
		panels.StorageOperationDuration(datasource, labelMatcher),
	)
}

func BuildKubeletMixin(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcher(clusterLabelName)
	return dashboards.NewDashboardResult(
		dashboard.New("kubelet",
			dashboard.ProjectName(project),
			dashboard.Name("Kubernetes / Kubelet"),
			dashboard.AddVariable("cluster",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("cluster",
						labelValuesVar.Matchers("up{"+panels.GetKubeletMatcher()),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("cluster"),
				),
			),
			dashboard.AddVariable("instance",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("instance",
						labelValuesVar.Matchers(
							promql.SetLabelMatchers(
								"up{"+panels.GetKubeletMatcher(),
								[]promql.LabelMatcher{{Name: "cluster", Type: "=", Value: "$cluster"}},
							),
						),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("instance"),
				),
			),
			withKubeletStats(datasource, clusterLabelMatcher),
			withKubeletOperations(datasource, clusterLabelMatcher),
			withKubeletOperationsQuantile(datasource, clusterLabelMatcher),
			withPodStartRateAndDuration(datasource, clusterLabelMatcher),
			withStorageOperationsAndErrors(datasource, clusterLabelMatcher),
			withStorageOperationsQuantile(datasource, clusterLabelMatcher),
		),
	).Component("kubernetes")
}
