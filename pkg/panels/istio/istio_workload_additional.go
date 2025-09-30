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

// Additional panels for Inbound Workloads
func IncomingRequestSizeBySource(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Request Size By Source",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{
					Unit: &dashboards.BytesUnit,
				},
				Min: 0,
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
				Palette:      timeSeriesPanel.Palette{Mode: timeSeriesPanel.AutoMode},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySource50"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySource90"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySource95"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySource99"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P99 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySource99"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P99 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySourceNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySourceNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySourceNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeBySourceNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99"),
			),
		),
	)
}

func IncomingResponseSizeBySource(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Response Size By Source",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{
					Unit: &dashboards.BytesUnit,
				},
				Min: 0,
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
				Palette:      timeSeriesPanel.Palette{Mode: timeSeriesPanel.AutoMode},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySource50"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySource90"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySource95"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySource99"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}}  P99 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySourceNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySourceNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySourceNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingResponseSizeBySourceNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload=~\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[1m])) by (source_workload, source_workload_namespace, le))",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{source_workload}}.{{source_workload_namespace}} P99"),
			),
		),
	)
}

func InboundTCPBytesReceived(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Bytes Received from Incoming TCP Connection",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{
					Unit: &dashboards.BytesPerSecondsUnit,
				},
				Min: 0,
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
				Palette:      timeSeriesPanel.Palette{Mode: timeSeriesPanel.AutoMode},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["InboundTCPBytesReceived"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"round(sum(irate(istio_tcp_received_bytes_total{reporter=~\"$qrep\", connection_security_policy=\"mutual_tls\", destination_workload_namespace=~\"$namespace\", destination_workload=~\"$workload\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[$__rate_interval])) by (source_workload, source_workload_namespace), 0.001)",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["InboundTCPBytesReceivedNonmTLS"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"round(sum(irate(istio_tcp_received_bytes_total{reporter=~\"$qrep\", connection_security_policy!=\"mutual_tls\", destination_workload_namespace=~\"$namespace\", destination_workload=~\"$workload\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[$__rate_interval])) by (source_workload, source_workload_namespace), 0.001)",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ source_workload }}.{{ source_workload_namespace}}"),
			),
		),
	)
}

func InboundTCPBytesSent(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Bytes Sent to Incoming TCP Connection",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{
					Unit: &dashboards.BytesPerSecondsUnit,
				},
				Min: 0,
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
				Palette:      timeSeriesPanel.Palette{Mode: timeSeriesPanel.AutoMode},
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["InboundTCPBytesSentTotal"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"round(sum(irate(istio_tcp_sent_bytes_total{connection_security_policy=\"mutual_tls\", reporter=~\"$qrep\", destination_workload_namespace=~\"$namespace\", destination_workload=~\"$workload\", source_workload=~\"$srcwl\", source_workload_namespace=~\"$srcns\"}[$__rate_interval])) by (source_workload, source_workload_namespace), 0.001)",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_service }} : {{ response_code }} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["InboundTCPBytesSentNonmTLS"],
					labelMatchers,
				).Pretty(0),
				// promql.SetLabelMatchers(
				// 	"round(sum(irate(istio_requests_total{destination_principal!~\"spiffe.*\", source_workload_namespace=~\"$namespace\", source_workload=~\"$workload\", reporter=\"source\", destination_service=~\"$dstsvc\"}[5m])) by (destination_service, response_code), 0.001)",
				// 	labelMatchers,
				// ),
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_service }} : {{ response_code }}"),
			),
		),
	)
}
