//检查Grafana上所有监控面板上的数据源
package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/athurg/grafana"
)

var (
	flagUrl  string
	flagKey  string
	flagSkip string
)

func init() {
	flag.StringVar(&flagUrl, "url", "https://grafana.com", "Grafana Server URL")
	flag.StringVar(&flagKey, "key", "", "Grafana API Key")
	flag.StringVar(&flagSkip, "skip", "", "Skipped datasources seperated by comma")
}

func main() {
	flag.Parse()

	grafanaCli := grafana.New(flagUrl, flagKey)

	dashboards, err := grafanaCli.SearchDashboards()
	if err != nil {
		fmt.Println(err)
		return
	}

	skippedDatasources := map[string]struct{}{}
	for _, s := range strings.Split(flagSkip, ",") {
		skippedDatasources[s] = struct{}{}
	}

	for _, dashboard := range dashboards {
		_, dashboard, err := grafanaCli.Dashboard(dashboard.Uid)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println()
		for _, panel := range dashboard.Panels {
			switch panel.Type {
			case "alertlist", "row":
				continue
			}

			if _, ok := skippedDatasources[panel.DataSource]; ok {
				continue
			}

			fmt.Printf("%-16s [%-10s] [%s] %s\n", panel.DataSource, dashboard.Uid, dashboard.Title, panel.Title)
		}
	}
}
