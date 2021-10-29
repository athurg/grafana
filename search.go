package grafana

import (
	"net/url"
)

type SearchItem struct {
	Id        int
	Uid       string
	Title     string
	Type      string
	IsStarred bool
	SortMeta  int
	Slug      string
	Uri       string
	Url       string
	Tags      []string
}

func (cli *Grafana) SearchAll() ([]SearchItem, error) {
	return cli.search("")
}

func (cli *Grafana) SearchFolders() ([]SearchItem, error) {
	return cli.search("dash-folder")
}

func (cli *Grafana) SearchDashboards() ([]SearchItem, error) {
	return cli.search("dash-db")
}

func (cli *Grafana) search(searchType string) ([]SearchItem, error) {
	var items []SearchItem
	q := url.Values{
		"type": {searchType},
	}

	err := cli.request("GET", "/api/search?"+q.Encode(), &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}
