package client

import (
	"encoding/json"
	"strconv"

	"github.com/pngouin/defectdojo-cli/config"
)

type ImportLanguages struct {
	Product int
	File    string
}

type ImportLanguagesResponse struct {
	Product int
}

func NewImportLanguagesClient(config config.Config) ImportLanguagesClient {
	client := NewHttpClient(config)
	return ImportLanguagesClient{
		baseEndpoint: "/api/v2/import-languages/",
		config:       config,
		http:         &client,
	}
}

type ImportLanguagesClient struct {
	baseEndpoint string
	config       config.Config
	http         *HttpClient
}

func (lc ImportLanguagesClient) Send(lang ImportLanguages) (ImportLanguagesResponse, error) {
	body := make(map[string]string)
	body["product"] = strconv.Itoa(lang.Product)

	resp, err := lc.http.Multipart(lc.baseEndpoint, body, "file", lang.File)
	if err != nil {
		return ImportLanguagesResponse{}, err
	}

	var languagesResp ImportLanguagesResponse
	err = json.NewDecoder(resp.Body).Decode(&languagesResp)
	return languagesResp, err
}
