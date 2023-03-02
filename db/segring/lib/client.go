package lib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

const (
	baseURL = "http://0.0.0.0:8080"
)

func main() {
	c := NewClient()

	c.Put("1", []byte("Alice"))
	c.Put("2", []byte("Bob"))

	a := c.Scan("1", 2)
	for _, v := range a {
		fmt.Println(string(v))
	}

	c.Delete("1")
	fmt.Println(string(c.Get("1")))
	fmt.Println(string(c.Get("2")))

}

type IClient interface {
	Put(k string, v []byte)
	Get(key string) []byte
	Scan(lKey string, count int) [][]byte
	Delete(k string)
	Close()
}

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{client: &http.Client{}}
}

func (c *Client) Put(k string, v []byte) {
	req, _ := http.NewRequest(http.MethodPost, baseURL+"/put/"+k, bytes.NewBuffer(v))
	readBytes(c, req)
}

func (c *Client) Get(key string) []byte {
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/get/"+key, nil)
	return readBytes(c, req)
}

func (c *Client) Scan(lKey string, count int) [][]byte {
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/scan/"+lKey+"/"+fmt.Sprint(count), nil)
	bodyBytes := readBytes(c, req)

	var res [][]byte
	for len(bodyBytes) > 0 {
		lenBytes := bodyBytes[:8]
		length := *(*int64)(unsafe.Pointer(&lenBytes[0]))
		bodyBytes = bodyBytes[8:]

		val := bodyBytes[:length]
		res = append(res, val)
		bodyBytes = bodyBytes[length:]
	}

	return res
}

func (c *Client) Delete(k string) {
	req, _ := http.NewRequest(http.MethodDelete, baseURL+"/delete/"+k, nil)
	readBytes(c, req)
}

func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func readBytes(c *Client, req *http.Request) []byte {
	resp, err := c.client.Do(req)
	if err != nil {
		panic(err)
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	return bodyBytes
}
