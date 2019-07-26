package curl

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	METHOD_POST = "POST"
	METHOD_GET  = "GET"
)

type M map[string]interface{}

type FileInfo struct {
	Name   string
	Stream io.Reader
}

type FormData struct {
	File   map[string]FileInfo
	Params map[string]string
}

type HttpReq struct {
	Url        string
	Method     string
	Header     map[string]string
	Params     M
	FormData   FormData
	Body       []byte
	BodyReader io.Reader
	CertFile   string
	KeyFile    string
	Timeout    time.Duration
}

func (h *HttpReq) buildBody() {
	if h.Body != nil || h.BodyReader != nil {
		return
	}

	var data string
	for k, v := range h.Params {
		if data != "" {
			data += "&"
		}
		data += k + "=" + v.(string)
	}

	switch h.Method {
	case METHOD_POST:
		h.Body = []byte(data)
		break
	case METHOD_GET:
		urlArr := strings.Split(h.Url, "?")
		if len(urlArr) == 2 {
			if data != "" {
				urlArr[1] = urlArr[1] + "&" + data
			}
			//将GET请求的参数进行转义
			h.Url = urlArr[0] + "?" + url.PathEscape(urlArr[1])
		}
		break
	}
}

func (h *HttpReq) Do() ([]byte, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if h.CertFile != "" {
		cert, err := tls.LoadX509KeyPair(h.CertFile, h.KeyFile)
		if err != nil {
			return nil, err
		}
		tr.DisableCompression = true
		tr.TLSClientConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
		}
	}
	var client = &http.Client{
		Transport: tr,
		Timeout:   h.Timeout,
	}
	h.buildBody()

	var bReader io.Reader
	if h.BodyReader != nil {
		bReader = h.BodyReader
	} else {
		bReader = bytes.NewReader(h.Body)
	}

	req, err := http.NewRequest(h.Method, h.Url, bReader)
	if err != nil {
		return nil, err
	}

	if h.Header != nil {
		for k, v := range h.Header {
			req.Header.Set(k, v)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (h *HttpReq) Get() ([]byte, error) {

	h.Method = METHOD_GET
	return h.Do()
}

func (h *HttpReq) Post() ([]byte, error) {

	h.Method = METHOD_POST
	h.Header = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	return h.Do()
}

func (h *HttpReq) PostForm() ([]byte, error) {

	h.Method = METHOD_POST

	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)

	if h.FormData.File != nil {
		for k, file := range h.FormData.File {
			part, err := w.CreateFormFile(k, file.Name)
			if err != nil {
				return nil, err
			}
			if _, err = io.Copy(part, file.Stream); err != nil {
				return nil, err
			}
		}
	}

	if h.FormData.Params != nil {
		for k, v := range h.FormData.Params {
			if err := w.WriteField(k, v); err != nil {
				return nil, err
			}
		}
	}

	if err := w.Close(); err != nil {
		return nil, err
	}

	if h.Header == nil {
		h.Header = make(map[string]string)
	}
	h.Header["Content-Type"] = w.FormDataContentType()
	h.BodyReader = &buf

	return h.Do()
}
