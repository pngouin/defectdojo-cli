package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/pngouin/defectdojo-cli/config"
)

type EngagementStatus string
type EngagementType string

type Engagement struct {
	ProductId   int              `json:"product"`
	TargetStart string           `json:"target_start"`
	TargetEnd   string           `json:"target_end"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Version     string           `json:"version"`
	Status      EngagementStatus `json:"status"`
	Type        EngagementType   `json:"engagement_type"`
	CommitHash  string           `json:"commit_hash"`
	BranchTag   string           `json:"branch_tag"`
}

type EngagementResponse struct {
	Id int `json:"id"`
	Engagement
}

var (
	ErrCannotCloseEngagement = errors.New("cannot close engagement")
)

const (
	NotStarted         EngagementStatus = "Not Started"
	Blocked            EngagementStatus = "Blocked"
	Cancelled          EngagementStatus = "Cancelled"
	Completed          EngagementStatus = "Completed"
	InProgress         EngagementStatus = "In Progress"
	OnHold             EngagementStatus = "On Hold"
	WaitingforResource EngagementStatus = "Waiting for Resource"

	Interactive EngagementType = "Interactive"
	CICD        EngagementType = "CI/CD"
)

func NewEngagementClient(config config.Config) EngagementClient {
	client := NewHttpClient(config)
	return EngagementClient{
		baseEndpoint: "/api/v2/engagements/",
		config:       config,
		http:         &client,
	}
}

type EngagementClient struct {
	baseEndpoint string
	config       config.Config
	http         *HttpClient
}

func (ec EngagementClient) Create(e Engagement) (EngagementResponse, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return EngagementResponse{}, err
	}
	resp, err := ec.http.Post(ec.baseEndpoint, body)
	if err != nil {
		return EngagementResponse{}, err
	}
	var engagementResponse EngagementResponse
	err = json.NewDecoder(resp.Body).Decode(&engagementResponse)
	return engagementResponse, err
}

func (ec EngagementClient) Get(id string) (EngagementResponse, error) {
	path := fmt.Sprint(ec.baseEndpoint, id, "/")
	resp, err := ec.http.Get(path)
	if err != nil {
		return EngagementResponse{}, err
	}
	var engagementResponse EngagementResponse
	err = json.NewDecoder(resp.Body).Decode(&engagementResponse)
	return engagementResponse, err
}

func (ec EngagementClient) Close(id string) (error) {
	path := fmt.Sprint(ec.baseEndpoint, id, "/close/")
	resp, err := ec.http.Post(path, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return ErrCannotCloseEngagement
	}
	return nil
}
