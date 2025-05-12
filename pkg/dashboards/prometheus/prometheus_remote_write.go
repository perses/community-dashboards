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

func withPrometheusRwTimestamps(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Timestamps",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusRemoteStorageTimestampLag(datasource, labelMatcher),
		panels.PrometheusRemoteStorageRateLag(datasource, labelMatcher),
	)
}

func withPrometheusRwSamples(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Samples",
		panelgroup.PanelsPerLine(1),
		panels.PrometheusRemoteStorageSampleRate(datasource, labelMatcher),
	)
}

func withPrometheusRwShard(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Shards",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusRemoteStorageCurrentShards(datasource, labelMatcher),
		panels.PrometheusRemoteStorageDesiredShards(datasource, labelMatcher),
		panels.PrometheusRemoteStorageMaxShards(datasource, labelMatcher),
		panels.PrometheusRemoteStorageMinShards(datasource, labelMatcher),
	)
}

func withPrometheusRwShardDetails(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Shard Details",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusRemoteStorageShardCapacity(datasource, labelMatcher),
		panels.PrometheusRemoteStoragePendingSamples(datasource, labelMatcher),
	)
}

func withPrometheusRwSegments(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Segments",
		panelgroup.PanelsPerLine(2),
		panels.PrometheusTSDBCurrentSegment(datasource, labelMatcher),
		panels.PrometheusRemoteWriteCurrentSegment(datasource, labelMatcher),
	)
}

func withPrometheusRwMiscRates(datasource string, labelMatcher promql.LabelMatcher) dashboard.Option {
	return dashboard.AddPanelGroup("Misc. Rates",
		panelgroup.PanelsPerLine(4),
		panels.PrometheusRemoteStorageDroppedSamplesRate(datasource, labelMatcher),
		panels.PrometheusRemoteStorageFailedSamplesRate(datasource, labelMatcher),
		panels.PrometheusRemoteStorageRetriedSamplesRate(datasource, labelMatcher),
		panels.PrometheusRemoteStorageEnqueueRetriesRate(datasource, labelMatcher),
	)
}

func BuildPrometheusRemoteWrite(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcher(clusterLabelName)
	return dashboards.NewDashboardResult(
		dashboard.New("prometheus-remote-write",
			dashboard.Name("Prometheus / Remote Write"),
			dashboard.ProjectName(project),
			dashboards.AddClusterVariable(datasource, clusterLabelName, "prometheus_remote_storage_shards"),
			dashboard.AddVariable("instance",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("instance",
						labelValuesVar.Matchers(
							promql.SetLabelMatchers(
								"prometheus_remote_storage_shards",
								[]promql.LabelMatcher{clusterLabelMatcher},
							),
						),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("instance"),
				),
			),
			dashboard.AddVariable("url",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("url",
						labelValuesVar.Matchers(
							promql.SetLabelMatchers(
								"prometheus_remote_storage_shards{instance='$instance'}",
								[]promql.LabelMatcher{clusterLabelMatcher},
							),
						),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("url"),
				),
			),
			withPrometheusRwTimestamps(datasource, clusterLabelMatcher),
			withPrometheusRwSamples(datasource, clusterLabelMatcher),
			withPrometheusRwShard(datasource, clusterLabelMatcher),
			withPrometheusRwShardDetails(datasource, clusterLabelMatcher),
			withPrometheusRwSegments(datasource, clusterLabelMatcher),
			withPrometheusRwMiscRates(datasource, clusterLabelMatcher),
		),
	).Component("prometheus")
}
