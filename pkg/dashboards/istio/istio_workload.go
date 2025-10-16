package istio

import (
	"github.com/perses/community-mixins/pkg/dashboards"
	panels "github.com/perses/community-mixins/pkg/panels/istio"
	"github.com/perses/perses/go-sdk/dashboard"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	listVar "github.com/perses/perses/go-sdk/variable/list-variable"
	markdownPanel "github.com/perses/plugins/markdown/sdk/go"
	labelValuesVar "github.com/perses/plugins/prometheus/sdk/go/variable/label-values"
	promqlVar "github.com/perses/plugins/prometheus/sdk/go/variable/promql"
	"github.com/prometheus/prometheus/model/labels"
)

// Helper function to create markdown header panels
func createWorkloadHeaderPanel(title string) panelgroup.Option {
	return panelgroup.AddPanel(title,
		markdownPanel.Markdown(title+" Header",
			markdownPanel.Text("<div class=\"dashboard-header text-center\">\n<span>"+title+"</span>\n</div>"),
		),
	)
}

func withWorkloadGeneralSection() dashboard.Option {
	return dashboard.AddPanelGroup("General",
		panelgroup.PanelsPerLine(1),
		panelgroup.PanelHeight(3),
		panelgroup.AddPanel("WORKLOAD",
			markdownPanel.Markdown("WORKLOAD Header",
				markdownPanel.Text("<div class=\"dashboard-header text-center\">\n<span>WORKLOAD: $workload.$namespace</span>\n</div>"),
			),
		),
	)
}

func withWorkloadGeneralIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("General (II)",
		panelgroup.PanelsPerLine(3),
		panelgroup.PanelHeight(4),
		panels.IncomingRequestVolumeStat(datasource, labelMatcher),
		panels.IncomingSuccessRateStat(datasource, labelMatcher),
		panels.RequestDurationChart(datasource, labelMatcher),
	)
}

func withWorkloadGeneralIIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("General (III)",
		panelgroup.PanelsPerLine(2),
		panelgroup.PanelHeight(4),
		panels.TCPServerTrafficStat(datasource, labelMatcher),
		panels.TCPClientTrafficStat(datasource, labelMatcher),
	)
}

func withWorkloadInboundWorkloadsSection() dashboard.Option {
	return dashboard.AddPanelGroup("Inbound Workloads",
		panelgroup.PanelsPerLine(1),
		panelgroup.PanelHeight(3),
		panelgroup.AddPanel("INBOUND WORKLOADS",
			markdownPanel.Markdown("INBOUND WORKLOADS Header",
				markdownPanel.Text("<div class=\"dashboard-header text-center\">\n<span>INBOUND WORKLOADS</span>\n</div>"),
			),
		),
	)
}

func withWorkloadInboundWorkloadsIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Inbound Workloads",
		panelgroup.PanelsPerLine(2),
		panelgroup.PanelHeight(6),
		panels.IncomingRequestVolume(datasource, labelMatcher), // "Incoming Requests By Source And Response Code"
		panels.IncomingSuccessRate(datasource, labelMatcher),   // "Incoming Success Rate (non-5xx responses) By Source"
	)
}

func withWorkloadInboundWorkloadsIIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Inbound Workloads (II)",
		panelgroup.PanelsPerLine(3),
		panelgroup.PanelHeight(6),
		panels.IncomingRequestDuration(datasource, labelMatcher),      // "Incoming Request Duration By Source"
		panels.IncomingRequestSizeBySource(datasource, labelMatcher),  // "Incoming Request Size By Source"
		panels.IncomingResponseSizeBySource(datasource, labelMatcher), // "Incoming Response Size By Source"
	)
}

func withWorkloadInboundWorkloadsIVSection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Inbound Workloads (III)",
		panelgroup.PanelsPerLine(2),
		panelgroup.PanelHeight(6),
		panels.InboundTCPBytesReceived(datasource, labelMatcher), // "Bytes Received from Incoming TCP Connection"
		panels.InboundTCPBytesSent(datasource, labelMatcher),     // "Bytes Sent to Incoming TCP Connection"
	)
}

func withWorkloadOutboundServicesSection() dashboard.Option {
	return dashboard.AddPanelGroup("Outbound Services",
		panelgroup.PanelsPerLine(1),
		panelgroup.PanelHeight(3),
		createWorkloadHeaderPanel("OUTBOUND SERVICES"),
	)
}

func withWorkloadOutboundServicesIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Outbound Services (II)",
		panelgroup.PanelsPerLine(2),
		panelgroup.PanelHeight(6),
		panels.OutgoingRequestVolume(datasource, labelMatcher),
		panels.OutgoingSuccessRate(datasource, labelMatcher),
	)
}

func withWorkloadOutboundServicesIIISection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Outbound Services (III)",
		panelgroup.PanelsPerLine(3),
		panelgroup.PanelHeight(6),
		panels.OutgoingRequestDuration(datasource, labelMatcher),
		panels.OutgoingRequestSize(datasource, labelMatcher),
		panels.OutgoingResponseSize(datasource, labelMatcher),
	)
}

func withWorkloadOutboundServicesIVSection(datasource string, labelMatcher *labels.Matcher) dashboard.Option {
	return dashboard.AddPanelGroup("Outbound Services (IV)",
		panelgroup.PanelsPerLine(2),
		panelgroup.PanelHeight(6),
		panels.TCPBytesSent(datasource, labelMatcher),
		panels.TCPBytesReceived(datasource, labelMatcher),
	)
}

func BuildIstioWorkload(project string, datasource string, clusterLabelName string) dashboards.DashboardResult {
	clusterLabelMatcher := dashboards.GetClusterLabelMatcherV2(clusterLabelName)
	return dashboards.NewDashboardResult(
		dashboard.New("istio-workload-dashboard",
			dashboard.ProjectName(project),
			dashboard.Name("Istio Workload Dashboard"),
			dashboards.AddClusterVariable(datasource, clusterLabelName, "istio_requests_total"),
			dashboard.AddVariable("namespace",
				listVar.List(
					promqlVar.PrometheusPromQL(
						"sum(istio_requests_total) by (destination_workload_namespace) or sum(istio_tcp_sent_bytes_total) by (destination_workload_namespace)",
						promqlVar.Datasource(datasource),
						promqlVar.LabelName("destination_workload_namespace"),
					),
					listVar.DisplayName("Namespace"),
					listVar.DefaultValue("bookinfo"),
				),
			),
			dashboard.AddVariable("workload",
				listVar.List(
					promqlVar.PrometheusPromQL(
						"(sum(istio_requests_total{destination_workload_namespace=~\"$namespace\"}) by (destination_workload) or sum(istio_requests_total{source_workload_namespace=~\"$namespace\"}) by (source_workload)) or (sum(istio_tcp_sent_bytes_total{destination_workload_namespace=~\"$namespace\"}) by (destination_workload) or sum(istio_tcp_sent_bytes_total{source_workload_namespace=~\"$namespace\"}) by (source_workload))",
						promqlVar.Datasource(datasource),
						promqlVar.LabelName("source_workload"),
					),
					listVar.DisplayName("Workload"),
					listVar.DefaultValue("details-v1"),
				),
			),
			dashboard.AddVariable("qrep",
				listVar.List(
					promqlVar.PrometheusPromQL(
						"sum(istio_requests_total) by (reporter)",
						promqlVar.Datasource(datasource),
						promqlVar.LabelName("reporter"),
					),
					listVar.DisplayName("Reporter"),
					listVar.DefaultValue("destination"),
					listVar.AllowMultiple(true),
				),
			),
			dashboard.AddVariable("srcns",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("source_workload_namespace",
						labelValuesVar.Matchers("istio_requests_total{reporter=~\"$qrep\", destination_workload=\"$workload\", destination_workload_namespace=~\"$namespace\"}"),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("Inbound Workload Namespace"),
					listVar.AllowAllValue(true),
					listVar.AllowMultiple(true),
				),
			),
			dashboard.AddVariable("srcwl",
				listVar.List(
					labelValuesVar.PrometheusLabelValues("source_workload",
						labelValuesVar.Matchers("istio_requests_total{reporter=~\"$qrep\", destination_workload=\"$workload\", destination_workload_namespace=~\"$namespace\", source_workload_namespace=~\"$srcns\"}"),
						dashboards.AddVariableDatasource(datasource),
					),
					listVar.DisplayName("Inbound Workload"),
					listVar.AllowAllValue(true),
					listVar.AllowMultiple(true),
				),
			),
			dashboard.AddVariable("dstsvc",
				listVar.List(
					promqlVar.PrometheusPromQL(
						"sum(istio_requests_total{reporter=\"source\", source_workload=~\"$workload\", source_workload_namespace=~\"$namespace\"}) by (destination_service) or sum(istio_tcp_sent_bytes_total{reporter=\"source\", source_workload=~\"$workload\", source_workload_namespace=~\"$namespace\"}) by (destination_service)",
						promqlVar.Datasource(datasource),
						promqlVar.LabelName("destination_service"),
					),
					listVar.DisplayName("Destination Service"),
					listVar.AllowAllValue(true),
					listVar.AllowMultiple(true),
				),
			),
			// Add all sections that match the JSON layout
			withWorkloadGeneralSection(),
			withWorkloadGeneralIISection(datasource, clusterLabelMatcher),
			withWorkloadGeneralIIISection(datasource, clusterLabelMatcher),
			withWorkloadInboundWorkloadsSection(),
			withWorkloadInboundWorkloadsIISection(datasource, clusterLabelMatcher),
			withWorkloadInboundWorkloadsIIISection(datasource, clusterLabelMatcher),
			withWorkloadInboundWorkloadsIVSection(datasource, clusterLabelMatcher),
			withWorkloadOutboundServicesSection(),
			withWorkloadOutboundServicesIISection(datasource, clusterLabelMatcher),
			withWorkloadOutboundServicesIIISection(datasource, clusterLabelMatcher),
			withWorkloadOutboundServicesIVSection(datasource, clusterLabelMatcher),
		),
	).Component("istio")
}
