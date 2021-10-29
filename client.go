package grafana

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Grafana struct {
	url string
	key string
}

func New(url, key string) *Grafana {
	return &Grafana{url: url, key: key}
}

func (cli *Grafana) request(method, path string, result interface{}) error {
	req, err := http.NewRequest(method, cli.url+path, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+cli.key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
