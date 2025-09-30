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

// ========== SERVICE WORKLOAD PANELS ==========

func IncomingRequestsByService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Requests By Destination Workload And Response Code",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{Min: 0}),
			timeSeriesPanel.WithLegend(timeSeriesPanel.Legend{
				Position: timeSeriesPanel.BottomPosition,
				Mode:     timeSeriesPanel.ListMode,
			}),
			timeSeriesPanel.WithVisual(timeSeriesPanel.Visual{
				Display:      timeSeriesPanel.LineDisplay,
				ConnectNulls: false,
				LineWidth:    1,
				AreaOpacity:  0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioIncomingRequestsByService"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_requests_total{connection_security_policy=\"mutual_tls\",destination_service=~\"$service\",reporter=\"destination\",destination_workload=~\"$dstwl\",destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace, response_code), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} : {{ response_code }} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioIncomingRequestsByServiceNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_requests_total{connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", reporter=\"destination\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace, response_code), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} : {{ response_code }}"),
			),
		),
	)
}

func IncomingSuccessRateByService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Success Rate (non-5xx responses) By Destination Workload",
		timeSeriesPanel.Chart(
			timeSeriesPanel.WithYAxis(timeSeriesPanel.YAxis{
				Format: &commonSdk.Format{Unit: &dashboards.PercentDecimalUnit},
				Min:    0,
				Max:    1.01,
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
					IstioCommonPanelQueries["IncomingSuccessRateByService"],
					labelMatchers,
				).Pretty(0),
				//"sum(irate(istio_requests_total{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\",response_code!~\"5.*\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace) / sum(irate(istio_requests_total{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IstioIncomingRequestsByServiceNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"sum(irate(istio_requests_total{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\",response_code!~\"5.*\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace) / sum(irate(istio_requests_total{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[5m])) by (destination_workload, destination_workload_namespace)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }}"),
			),
		),
	)
}

func IncomingRequestDurationByService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Request Duration By Service Workload",
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
					IstioCommonPanelQueries["IncomingRequestDurationByService50"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.50, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.50, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByService90"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.90, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.90, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByService95"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.95, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.95, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByService99"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.99, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.99, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByServiceNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.50, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.50, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByServiceNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.90, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.90, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByServiceNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.95, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.95, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestDurationByServiceNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				//"(histogram_quantile(0.99, sum(irate(istio_request_duration_milliseconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le)) / 1000) or histogram_quantile(0.99, sum(irate(istio_request_duration_seconds_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P99"),
			),
		),
	)
}

func IncomingRequestSizeByService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Incoming Request Size By Service Workload",
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
					IstioCommonPanelQueries["IncomingRequestSizeByService50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByService90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByService95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByService99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }}  P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByServiceNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByServiceNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByServiceNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["IncomingRequestSizeByServiceNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_request_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P99"),
			),
		),
	)
}

func ResponseSizeByService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
	return panelgroup.AddPanel("Response Size By Service Workload",
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
					IstioCommonPanelQueries["ResponseSizeByService50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByService90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }}  P90 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByService95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95 (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByService99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }}  P99 (🔐mTLS)"),
			),
		),
		// Non-mTLS P50, P90, P95, P99
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByServiceNonmTLS50"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.50, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P50"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByServiceNonmTLS90"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.90, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P90"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByServiceNonmTLS95"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.95, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P95"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["ResponseSizeByServiceNonmTLS99"],
					labelMatchers,
				).Pretty(0),
				//"histogram_quantile(0.99, sum(irate(istio_response_bytes_bucket{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[1m])) by (destination_workload, destination_workload_namespace, le))",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace }} P99"),
			),
		),
	)
}

func BytesReceivedFromTCPService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
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
					IstioCommonPanelQueries["BytesReceivedFromTCPService"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_received_bytes_total{reporter=\"destination\", connection_security_policy=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[$__rate_interval])) by (destination_workload, destination_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace}} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesReceivedFromTCPServiceNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_received_bytes_total{reporter=\"destination\", connection_security_policy!=\"mutual_tls\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[$__rate_interval])) by (destination_workload, destination_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{ destination_workload_namespace}}"),
			),
		),
	)
}

func BytesSentToTCPService(datasourceName string, labelMatchers ...*labels.Matcher) panelgroup.Option {
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
					IstioCommonPanelQueries["BytesSentToTCPService"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_sent_bytes_total{connection_security_policy=\"mutual_tls\", reporter=\"destination\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[$__rate_interval])) by (destination_workload, destination_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{destination_workload_namespace }} (🔐mTLS)"),
			),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchersV2(
					IstioCommonPanelQueries["BytesSentToTCPServiceNonmTLS"],
					labelMatchers,
				).Pretty(0),
				//"round(sum(irate(istio_tcp_sent_bytes_total{connection_security_policy!=\"mutual_tls\", reporter=\"destination\", destination_service=~\"$service\", destination_workload=~\"$dstwl\", destination_workload_namespace=~\"$dstns\"}[$__rate_interval])) by (destination_workload, destination_workload_namespace), 0.001)",
				dashboards.AddQueryDataSource(datasourceName),
				query.SeriesNameFormat("{{ destination_workload }}.{{destination_workload_namespace }}"),
			),
		),
	)
}
