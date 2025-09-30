package istio

import (
	"github.com/perses/community-dashboards/pkg/dashboards"
	"github.com/perses/community-dashboards/pkg/promql"
	commonSdk "github.com/perses/perses/go-sdk/common"
	"github.com/perses/perses/go-sdk/panel"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	"github.com/perses/plugins/prometheus/sdk/go/query"
	timeSeriesPanel "github.com/perses/plugins/timeserieschart/sdk/go"
	"github.com/prometheus/prometheus/model/labels"
)

// ========== CLIENT WORKLOAD PANELS (continued) ==========

func IncomingRequestDurationByClient(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Request Duration By Source",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.SecondsUnit},
				Min:    0,
			}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0.1,
			}),
		),
		// mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClient50"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.50, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.50, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClient90"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.90, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.90, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClient95"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.95, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.95, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClient99"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.99, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.99, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClientNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.50, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.50, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClientNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.90, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.90, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClientNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.95, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.95, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByClientNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				// "(histogram_quantile(0.99, sum(irate(istio_request_duration_milliseconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le)) / 1000) or histogram_quantile(0.99, sum(irate(istio_request_duration_seconds_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99"),
			),
		),
	)
}

func IncomingRequestSizeByClient(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Request Size By Source",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.BytesUnit},
				Min:    0,
			}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0.1,
			}),
		),
		// mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClient50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClient90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClient95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClient99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClientNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClientNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClientNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByClientNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99"),
			),
		),
	)
}

func ResponseSizeByClient(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Response Size By Source",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.BytesUnit},
				Min:    0,
			}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0.1,
			}),
		),
		// mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClient50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClient90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClient95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClient95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClientNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClientNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClientNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioResponseSizeByClientNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99"),
			),
		),
	)
}

func BytesReceivedFromTCPClient(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Bytes Received from Incoming TCP Connection",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.BytesPerSecondsUnit},
				Min:    0,
			}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0.1,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesReceivedFromTCPClient"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_received_bytes_total{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesReceivedFromTCPClientNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_received_bytes_total{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}}"),
			),
		),
	)
}

func BytesSentToTCPClient(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Bytes Sent to Incoming TCP Connection",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.BytesPerSecondsUnit},
				Min:    0,
			}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0.1,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesSentToTCPClient"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_sent_bytes_total{connection_security_policy=\"mutual_tls\", reporter=~\"$qrep\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesSentToTCPClientNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_sent_bytes_total{connection_security_policy!=\"mutual_tls\", reporter=~\"$qrep\", destination_service=~\"$service\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}}"),
			),
		),
	)
}
