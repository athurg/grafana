package grafana

type DashboardMeta struct {
	Type                  string
	CanSave               bool
	CanEdit               bool
	CanAdmin              bool
	CanStar               bool
	Slug                  string
	Url                   string
	Expires               string
	Created               string
	Updated               string
	UpdatedBy             string
	CreatedBy             string
	Version               int
	HasAcl                bool
	IsFolder              bool
	FolderId              int
	FolderUid             string
	FolderTitle           string
	FolderUrl             string
	Provisioned           bool
	ProvisionedExternalId string
}

type PanelTarget struct {
	Exemplar       bool
	Expr           string
	Format         string
	Interval       string
	IntervalFactor int
	LegendFormat   string
	RefId          string
}

type Panel struct {
	Id            int
	DataSource    string
	Type          string
	Title         string
	TimeFrom      interface{}
	TimeShift     interface{}
	PluginVersion string
	GridPos       struct{ H, W, X, Y int }
	Targets       []PanelTarget
}

type Dashboard struct {
	Id     int
	Uid    string
	Title  string
	Panels []Panel
}

func (cli *Grafana) Dashboard(uid string) (DashboardMeta, Dashboard, error) {
	var result struct {
		Meta      DashboardMeta
		Dashboard Dashboard
	}
	err := cli.request("GET", "/api/dashboards/uid/"+uid, &result)
	if err != nil {
		return DashboardMeta{}, Dashboard{}, err
	}

	return result.Meta, result.Dashboard, nil
}
