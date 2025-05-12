package prometheus

import (
	"github.com/perses/community-dashboards/pkg/dashboards"
	panels "github.com/perses/community-dashboards/pkg/panels/prometheus"
	"github.com/perses/community-dashboards/pkg/promql"
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"

	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
	labelValuesVar "github.com/perses/plugins/prometheus/sdk/go/variable/label-values"
)

func withPrometheusOverviewStatsGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Prometheus Stats",
		panelgroup.PanelsPerLine(1),
		panels.PrometheusStatsTable(datasource, labelMatcher),
	)
}

func withPrometheusOverviewDiscoveryGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Discovery",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusTargetSync(datasource, labelMatcher),
		panels.PrometheusTargets(datasource, labelMatcher),
	)
}

func withPrometheusRetrievalGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Retrieval",
		panelgroup.PanelsPerLine(3),
		panels.PrometheusAverageScrapeIntervalDuration(datasource, labelMatcher),
		panels.PrometheusScrapeFailures(datasource, labelMatcher),
		panels.PrometheusAppendedSamples(datasource, labelMatcher),
	)
}

func withPrometheusStorageGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Storage",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusHeadSeries(datasource, labelMatcher),
		panels.PrometheusHeadChunks(datasource, labelMatcher),
	)
}

func withPrometheusQueryGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Query",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusQueryRate(datasource, labelMatcher),
		panels.PrometheusQueryStateDuration(datasource, labelMatcher),
	)
}

func BuildPrometheusOverview(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcher(clusterLabelName)
	return dashboards.NewDashboardResult(
		dashboard.New("prometheus-overview",
			dashboard.ProjectName(project),
			dashboard.Name("Prometheus / Overview"),
			dashboard.AddVariable("job",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("job",
						labelValuesVar.Matchers("prometheus_build_info{}"),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("job"),
				),
			),
			dashboards.AddClusterVariable(datasource, clusterLabelName, "prometheus_build_info"),
			dashboard.AddVariable("instance",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("instance",
						labelValuesVar.Matchers(
							promql.SetLabelMatchers(
								"prometheus_build_info",
								[]promql.LabelMatcher{clusterLabelMatcher, {Name: "job", Type: "=", Value: "$job"}},
							),
						),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("instance"),
				),
			),
			withPrometheusOverviewStatsGroup(datasource, clusterLabelMatcher),
			withPrometheusOverviewDiscoveryGroup(datasource, clusterLabelMatcher),
			withPrometheusRetrievalGroup(datasource, clusterLabelMatcher),
			withPrometheusStorageGroup(datasource, clusterLabelMatcher),
			withPrometheusQueryGroup(datasource, clusterLabelMatcher),
		),
	).Component("prometheus")
}
