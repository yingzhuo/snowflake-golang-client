package snowflake_golang_client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Config struct {
	Host         string
	Port         int
	ResponseType string
}

type Client struct {
	host string
	port int
}

func NewClient(config *Config) *Client {
	return &Client{
		host: config.Host,
		port: config.Port,
	}
}

func (cli *Client) NextId() int64 {
	return cli.NextIds(1)[0]
}

func (cli *Client) NextIds(n int) []int64 {
	url := fmt.Sprintf("http://%s:%d/id?n=%d", cli.host, cli.port, n)
	return cli.doNextJsonIds(url)
}

func (cli *Client) doNextJsonIds(url string) (ret []int64) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(body, &ret); err == nil {
		return
	} else {
		panic(err)
	}
}
