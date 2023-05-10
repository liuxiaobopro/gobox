package http

import (
	"io/ioutil"
	"net/http"

	"github.com/liuxiaobopro/gobox/mapx"
)

type Client struct {
	Debug     bool                   // 是否开启debug模式
	Url       string                 // 请求地址
	Params    map[string]string      // 请求参数
	Header    map[string]string      // 请求头
	Form      map[string]interface{} // 表单参数
	UserAgent string                 // User-Agent
}

// Get get请求返回[]byte
func (Client *Client) Get() ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	if Client.Params != nil {
		Client.Url = Client.Url + "?" + mapx.MapToQuery(Client.Params)
	}
	req, err := http.NewRequest(http.MethodGet, Client.Url, nil)
	if err != nil {
		return nil, err
	}
	if Client.Header != nil {
		for k, v := range Client.Header {
			req.Header.Set(k, v)
		}
	}
	if Client.UserAgent != "" {
		req.Header.Set("User-Agent", Client.UserAgent)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

// Post post请求返回[]byte
func (Client *Client) Post() ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)
	req, err := http.NewRequest(http.MethodPost, Client.Url, nil)
	if err != nil {
		return nil, err
	}
	if Client.Header != nil {
		for k, v := range Client.Header {
			req.Header.Set(k, v)
		}
	}
	if Client.Form != nil {
		req.PostForm = mapx.MapToForm(Client.Form)
	}
	if Client.UserAgent != "" {
		req.Header.Set("User-Agent", Client.UserAgent)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
