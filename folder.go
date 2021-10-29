package grafana

type Folder struct {
	Id    int
	Uid   string
	Title string
}

func (cli *Grafana) Folders() ([]Folder, error) {
	var folders []Folder
	err := cli.request("GET", "/api/folders", &folders)
	if err != nil {
		return nil, err
	}
	return folders, nil
}
