package lib

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"unsafe"
)

const (
	baseURL = "http://0.0.0.0:9090"
)

func main() {
	c := NewClient()

	c.Add("1", []byte("Arjun"))
	c.Add("2", []byte("Sangeetha"))

	a := c.Scan("1", 2)
	for _, v := range a {
		fmt.Println(string(v))
	}
	//c.Add("1", []byte("Arjun"))
	//fmt.Println(string(c.Read("2")))
	//
	////c.Delete("1")
	//fmt.Println(string(c.Read("1")))

}

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{client: &http.Client{}}
}

func (c *Client) Add(k string, v []byte) {
	req, _ := http.NewRequest(http.MethodPost, baseURL+"/api/v1/table/"+k, bytes.NewBuffer(v))
	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
}

func (c *Client) Read(key string) []byte {
	req, _ := http.NewRequest(http.MethodGet, baseURL+"/api/v1/table/"+key, nil)
	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)
	return bodyBytes

}

func (c *Client) Scan(lKey string, count int) [][]byte {
	req, _ := http.NewRequest(http.MethodPut, baseURL+"/api/v1/table/"+lKey+"/"+fmt.Sprint(count), nil)
	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
	bodyBytes, _ := io.ReadAll(resp.Body)

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
	req, _ := http.NewRequest(http.MethodDelete, baseURL+"/api/v1/table/"+k, nil)
	resp, _ := c.client.Do(req)
	defer resp.Body.Close()
}
