package snowflake_golang_client

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strings"
)

type Config struct {
	Host         string
	Port         int
	ResponseType string
}

const (
	Json     = "json"
	Protobuf = "protobuf"
)

type Client struct {
	host         string
	port         int
	responseType string
}

func NewJsonClient(host string, port int) *Client {
	return NewClient(&Config{
		Host:         host,
		Port:         port,
		ResponseType: Json,
	})
}

func NewProtobufClient(host string, port int) *Client {
	return NewClient(&Config{
		Host:         host,
		Port:         port,
		ResponseType: Protobuf,
	})
}

func NewClient(config *Config) *Client {

	if !strings.EqualFold(Json, config.ResponseType) && !strings.EqualFold(Protobuf, config.ResponseType) {
		panic(errors.New("unsupported response type '" + config.ResponseType + "'"))
	}

	return &Client{
		host:         config.Host,
		port:         config.Port,
		responseType: config.ResponseType,
	}
}

func (cli *Client) NextId() int64 {
	return cli.NextIds(1)[0]
}

func (cli *Client) NextIds(n int) []int64 {

	url := fmt.Sprintf("http://%s:%d/id?n=%d", cli.host, cli.port, n)

	switch {
	case strings.EqualFold(Json, cli.responseType):
		return cli.doNextJsonIds(url)
	case strings.EqualFold(Protobuf, cli.responseType):
		return cli.doNextProtobufIds(url)
	default:
		panic(errors.New("unsupported response type '" + cli.responseType + "'")) // 不会运行到此处
	}
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

func (cli *Client) doNextProtobufIds(url string) []int64 {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var idLst IdList
	if err = proto.Unmarshal(body, &idLst); err == nil {
		return idLst.Ids
	} else {
		panic(err)
	}
}
