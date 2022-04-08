package client

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/pngouin/defectdojo-cli/config"
)

type HttpClient struct {
	config config.Config
	client *http.Client
}

func NewHttpClient(config config.Config) HttpClient {
	return HttpClient{
		config: config,
		client: &http.Client{
			Timeout: 90*time.Second,
		},
	}
}

func (h HttpClient) Get(url string) (*http.Response, error) {
	return h.do(http.MethodGet, url, nil)
}

func (h HttpClient) Post(url string, body []byte) (*http.Response, error) {
	bodyRead := bytes.NewReader(body)
	return h.do(http.MethodPost, url, bodyRead)
}

func (h HttpClient) do(method string, url string, body io.Reader) (*http.Response, error) {
	completePath := fmt.Sprint(h.config.Host, url)
	req, err := http.NewRequest(method, completePath, body)
	if err != nil {
		return &http.Response{}, err
	}
	
	req.Header.Add("Authorization", h.getToken())
	req.Header.Add("Content-Type", "application/json")
	return h.client.Do(req)
}

func (h HttpClient) Multipart(url string, params map[string]string, paramName string, filePath string) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return &http.Response{}, err
	}
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		return &http.Response{}, err
	}
	fileInfo, err := file.Stat()
	if err != nil {
		return &http.Response{}, err
	}
	file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fileInfo.Name())
	if err != nil {
		return &http.Response{}, err
	}
	_, err = part.Write(fileContent)
	if err != nil {
		return &http.Response{}, err
	}
	for key, val := range params {
		err = writer.WriteField(key, val)
		if err != nil {
			return &http.Response{}, err
		}
	}
	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", writer.Boundary())
	writer.Close()
	completePath := fmt.Sprint(h.config.Host, url)
	req, err := http.NewRequest(http.MethodPost, completePath, body)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("Authorization", h.getToken())
	req.Header.Add("Content-Type", contentType)

	return h.client.Do(req)
}

func (h HttpClient) getToken() string {
	return fmt.Sprint("Token ", h.config.ApiKey)
}