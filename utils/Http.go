package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	HttpJson      = "application/json"
	HttpText      = "text/plain"
	HttpXForm     = "application/x-www-form-urlencoded"
	HttpMultipart = "multipart/form-data"

	MultiTypeFile int = 1 << iota
	MultiTypeNetFile
)

var (
	Http = newHttpTools()
)

type MultiFile struct {
	Typ     int
	Content string
}

func NewMultiFile(typ int, content string) *MultiFile {
	return &MultiFile{Typ: typ, Content: content}
}

type httpTools struct{}

func newHttpTools() *httpTools {
	return &httpTools{}
}

func (h httpTools) Get(url string) (Bytes, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (h httpTools) FormPost(url string, data map[string]interface{}, headers map[string]string, timeout time.Duration) (Bytes, error) {
	body, _ := h.getBody(HttpXForm, data)
	return h.Post(HttpXForm, url, body, headers, timeout)
}

func (h httpTools) JsonPost(url string, data map[string]interface{}, headers map[string]string, timeout time.Duration) (Bytes, error) {
	body, _ := h.getBody(HttpJson, data)
	return h.Post(HttpJson, url, body, headers, timeout)
}

func (h httpTools) MultiForm(url string, data map[string]interface{}, files map[string][]byte, headers map[string]string, timeout time.Duration) (Bytes, error) {
	body, typ := h.multipartData(data, files)
	return h.Post(typ, url, body, headers, timeout)
}

func (h httpTools) multipartData(params map[string]interface{}, files map[string][]byte) (io.Reader, string) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	var v string
	for key, val := range params {
		switch val.(type) {
		case int:
			v = strconv.Itoa(val.(int))
			break
		case int64:
			v = strconv.FormatInt(val.(int64), 10)

		case string:
			v = val.(string)
			break
		}

		_ = writer.WriteField(key, v)
	}

	for key, file := range files {
		formFile, _ := writer.CreateFormFile(key, "")
		io.Copy(formFile, bytes.NewReader(file))
	}

	return body, writer.FormDataContentType()
}

func (h httpTools) getBody(contentTyp string, data map[string]interface{}) (io.Reader, string) {
	var body io.Reader
	contentTyp = strings.ToLower(contentTyp)
	switch contentTyp {
	case HttpJson:
		byteBody, _ := json.Marshal(data)
		body = bytes.NewBuffer(byteBody)
		break
	case HttpXForm:
		body = strings.NewReader(h.BuildQuery(data, ""))
		break

	//case HttpMultipart:
	//	body, contentTyp = h.multipartData(data)
	//	break
	default:
		body = strings.NewReader(h.BuildQuery(data, ""))
	}

	return body, contentTyp
}

func (h httpTools) Post(contentTyp string, url string, body io.Reader, headers map[string]string, timeout time.Duration) (Bytes, error) {
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", contentTyp)
	for k, v := range headers {
		request.Header.Add(k, v)
	}

	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func (h httpTools) BuildQuery(m map[string]interface{}, keyFmt string) string {
	urlValues := url.Values{}
	var nk, ret string
	for k, v := range m {
		if len(keyFmt) != 0 {
			nk = fmt.Sprintf(keyFmt, k)
		} else {
			nk = k
		}

		switch v.(type) {
		case int:
			urlValues.Add(nk, strconv.Itoa(v.(int)))
			break

		case string:
			urlValues.Add(nk, v.(string))
			break

		case map[string]interface{}:
			ret += h.BuildQuery(v.(map[string]interface{}), nk+"[%s]")
			ret += "&"
			break
		}
	}

	ret += urlValues.Encode()
	return ret
}
