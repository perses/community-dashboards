package thanos

import (
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
	labelValuesVar "github.com/perses/plugins/prometheus/sdk/go/variable/label-values"

	"github.com/perses/community-dashboards/pkg/dashboards"
	panels "github.com/perses/community-dashboards/pkg/panels/thanos"
	"github.com/perses/community-dashboards/pkg/promql"
)

func withThanosQueryFrontendRequestsGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Query Frontend API",
		panelgroup.PanelsPerLine(4),
		panelgroup.PanelHeight(8),
		panels.QueryFrontendRequestRate(datasource, labelMatcher),
		panels.QueryFrontendQueryRate(datasource, labelMatcher),
		panels.QueryFrontendErrors(datasource, labelMatcher),
		panels.QueryFrontendDurations(datasource, labelMatcher),
	)
}

func withThanosQueryFrontendCacheGroup(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Query Frontend Cache Operations",
		panelgroup.PanelsPerLine(4),
		panelgroup.PanelHeight(8),
		panels.QueryFrontendCacheRequestRate(datasource, labelMatcher),
		panels.QueryFrontendCacheHitRate(datasource, labelMatcher),
		panels.QueryFrontendCacheMissRate(datasource, labelMatcher),
		panels.QueryFrontendFetchedKeyRate(datasource, labelMatcher),
	)
}

func BuildThanosQueryFrontendOverview(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcher(clusterLabelName)
	return dashboards.NewDashboardResult(
		dashboard.New("thanos-query-frontend-overview",
			dashboard.ProjectName(project),
			dashboard.Name("Thanos / Query Frontend / Overview"),
			dashboard.AddVariable("job",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("job",
						labelValuesVar.Matchers("thanos_build_info{container=\"thanos-query-frontend\"}"),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("job"),
					listVar.AllowMultiple(true),
				),
			),
			dashboards.AddClusterVariable(datasource, clusterLabelName, "thanos_build_info"),
			dashboard.AddVariable("namespace",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("namespace",
						labelValuesVar.Matchers("thanos_status"),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("namespace"),
				),
			),
			withThanosQueryFrontendRequestsGroup(datasource, clusterLabelMatcher),
			withThanosQueryFrontendCacheGroup(datasource, clusterLabelMatcher),
			withThanosResourcesGroup(datasource, clusterLabelMatcher),
		),
	).Component("thanos")
}
