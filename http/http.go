package http

import (
	"bytes"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/liuxiaobopro/gobox/mapx"
	numberx "github.com/liuxiaobopro/gobox/number"
)

type Client struct {
	Debug     bool                   // 是否开启debug模式
	Url       string                 // 请求地址
	Params    map[string]string      // 请求参数
	Header    map[string]string      // 请求头
	Form      map[string]interface{} // 表单参数
	UserAgent string                 // User-Agent
	Json      []byte                 // json参数
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
	return io.ReadAll(resp.Body)
}

// Post post请求返回[]byte
func (Client *Client) Post() ([]byte, error) {
	var (
		resp *http.Response
		err  error

		jsonData io.Reader
	)

	if Client.Json != nil {
		jsonData = bytes.NewBuffer(Client.Json)
	}

	req, err := http.NewRequest(http.MethodPost, Client.Url, jsonData)
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
	return io.ReadAll(resp.Body)
}

// GetRandomUserAgent 获取随机User-Agent
func GetRandomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36 Edge/16.16299",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:99.0) Gecko/20100101 Firefox/99.0",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.85 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64; rv:99.0) Gecko/20100101 Firefox/99.0",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPad; CPU OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Android 10; Mobile; rv:99.0) Gecko/99.0 Firefox/99.0",
	}

	// 使用当前时间作为随机数生成器的种子
	rand.Seed(time.Now().UnixNano())

	// 生成随机索引
	randomIndex := numberx.RandomInt(0, len(userAgents))

	// 返回随机选择的User-Agent字符串
	return userAgents[randomIndex]
}

// IsValidUrl 是否是有效的url
func IsValidUrl(str string) bool {
	if parsedURL, err := url.Parse(str); err != nil {
		return false
	} else {
		if parsedURL.Scheme == "" || parsedURL.Host == "" {
			return false
		}

		url := parsedURL.Scheme + "://" + parsedURL.Host

		headers := make(http.Header)
		headers.Set("User-Agent", GetRandomUserAgent())

		// 创建一个自定义的http.Client
		client := &http.Client{
			Timeout: 3 * time.Second, // 设置超时时间
		}

		// 创建HTTP请求
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return false
		}

		req.Header = headers

		// 发送HTTP请求
		resp, err := client.Do(req)
		if err != nil {
			return false
		}

		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			return false
		} else {
			return true
		}
	}
}
