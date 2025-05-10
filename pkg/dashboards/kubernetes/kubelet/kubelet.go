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
		),
	).Component("kubernetes")
}
