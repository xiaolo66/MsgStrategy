package MsgStrategy

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Request struct {
	Url    string
	Method string
	Header http.Header
	Body   string
}

type Option func(*Request)

func SetUrl(url string) Option {
	return func(r *Request) {
		r.Url = url
	}
}

func SetMethod(method string) Option {
	return func(r *Request) {
		r.Method = method
	}
}

func SetBody(body string) Option {
	return func(r *Request) {
		r.Body = body
	}
}

func SetHeader(f, s string) Option {
	header := http.Header{}
	header.Set(f, s)
	return func(r *Request) {
		r.Header = header
	}
}

var Res = Request{}

func (r *Request) SetConfig(option ...Option) {
	for _, v := range option {
		v(r)
	}
	return
}

func (r *Request) GetMsg() (string, error) {
	client := &http.Client{}
	proxy, err := url.Parse("http://127.0.0.1:4780")
	if err != nil {
		Log.Errorf("parse proxy failed")
		return "", err
	}
	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
	}
	client.Transport = tr
	req, err := http.NewRequest(r.Method, r.Url, strings.NewReader(r.Body))
	if err != nil {
		return "", err
	}
	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	Log.Infoln("Fetch Binance Announce Success")
	return string(b), err
}
