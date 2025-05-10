package kubernetes

import (
	"github.com/perses/community-dashboards/pkg/dashboards"
	"github.com/perses/community-dashboards/pkg/promql"
	"github.com/perses/perses/go-sdk/panel"
	panelgroup "github.com/perses/perses/go-sdk/panel-group"
	statPanel "github.com/perses/perses/go-sdk/panel/stat"
	"github.com/perses/perses/go-sdk/prometheus/query"

	commonSdk "github.com/perses/perses/go-sdk/common"
)

func RunningKubeletStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Running Kubelets",
		panel.Description("Number of Running Kubelets Instances"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(kubelet_node_name{cluster=~'$cluster',"+GetKubeletMatcher()+"})",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func RunningPodStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Running Pods",
		panel.Description("Total Number of Running Pods"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(kubelet_running_pods{cluster=~'$cluster',"+GetKubeletMatcher()+", instance=~'$instance'})",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func RunningContainersStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Running Containers",
		panel.Description("Total Number of Running Containers"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(kubelet_running_containers{cluster=~'$cluster',"+GetKubeletMatcher()+", instance=~'$instance'})",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func ActVolumeCountStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Actual Volume Count",
		panel.Description("Total Number of Volumes Currently Mounted"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(kubelet_running_containers{cluster=~'$cluster',"+GetKubeletMatcher()+", instance=~'$instance', state='actual_state_of_world'})",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func DesiredVolumeCountStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Desired Volume Count",
		panel.Description("Total Number of Desired Volume Mounts"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(kubelet_running_containers{cluster=~'$cluster',"+GetKubeletMatcher()+", instance=~'$instance', state='desired_state_of_world'})",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}

func ConfigErrorCountStat(datasourceName string, labelMatchers ...promql.LabelMatcher) panelgroup.Option {
	return panelgroup.AddPanel("Config Error Count",
		panel.Description("Node Config Error Count Per Second"),
		statPanel.Chart(
			statPanel.Calculation(commonSdk.LastCalculation),
			statPanel.Format(commonSdk.Format{
				Unit:          string(commonSdk.DecimalUnit),
				DecimalPlaces: 0,
			}),
		),
		panel.AddQuery(
			query.PromQL(
				promql.SetLabelMatchers(
					"sum(rate(kubelet_node_config_error{cluster=~'$cluster',"+GetKubeletMatcher()+", instance=~'$instance'}[$__rate_interval]))",
					labelMatchers,
				),
				dashboards.AddQueryDataSource(datasourceName),
			),
		),
	)
}
