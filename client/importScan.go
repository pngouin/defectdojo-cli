package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/pngouin/defectdojo-cli/config"
)

var (
	ErrScanTypeNotFound = errors.New("scanType not found")
)

type ImportScan struct {
	EngagementId int    `json:"engagement"`
	ScanType     string `json:"scan_type"`
	File         string
}

type ImportScanResponse struct {
	ScanDate       string `json:"scan_date"`
	ScanType       string `json:"scan_type"`
	ProductName    string `json:"product_name"`
	EngagementName string `json:"engagement_name"`
	EngagementId   int    `json:"engagement_id"`
	ProductId      int    `json:"product_id"`
}

var scanType []string = []string{
	"Anchore Engine Scan",
	"Anchore Grype",
	"Aqua Scan",
	"AuditJS Scan",
	"AWS Prowler Scan",
	"Burp REST API",
	"Bandit Scan",
	"CargoAudit Scan",
	"Checkmarx Scan detailed",
	"Checkmarx Scan",
	"Checkmarx OSA",
	"Coverity API",
	"Cobalt.io API",
	"Dependency Track Finding Packaging Format (FPF) Export",
	"Mobsfscan Scan",
	"SonarQube Scan detailed",
	"SonarQube Scan",
	"SonarQube API Import",
	"Dependency Check Scan",
	"Dockle Scan",
	"Nessus Scan",
	"Nexpose Scan",
	"NPM Audit Scan",
	"Yarn Audit Scan",
	"Whitesource Scan",
	"ZAP Scan",
	"Qualys Scan",
	"PHP Symfony Security Check",
	"Acunetix Scan",
	"Clair Scan",
	"Clair Klar Scan",
	"Veracode Scan",
	"Symfony Security Check",
	"DSOP Scan",
	"Terrascan Scan",
	"Trivy Scan",
	"TFSec Scan",
	"HackerOne Cases",
	"Snyk Scan",
	"GitLab Dependency Scanning Report",
	"GitLab SAST Report",
	"Checkov Scan",
	"SpotBugs Scan",
	"JFrog Xray Unified Scan",
	"Scout Suite Scan",
	"AWS Security Hub Scan",
	"Meterian Scan",
	"Github Vulnerability Scan",
	"Cloudsploit Scan",
	"KICS Scan",
	"SARIF",
	"Azure Security Center Recommendations Scan",
	"Hadolint Dockerfile check",
	"Semgrep JSON Report",
	"Generic Findings Import",
	"Trufflehog3 Scan",
	"Detect-secrets Scan",
	"Solar Appscreener Scan",
	"Gitleaks Scan",
	"pip-audit Scan",
	"Edgescan Scan",
	"Rubocop Scan",
	"JFrog Xray Scan",
	"CycloneDX Scan",
	"SSLyze Scan (JSON)",
	"Harbor Vulnerability Scan",
	"Rusty Hog Scan",
	"Hydra Scan",
}

func NewImportScanClient(config config.Config) ImportScanClient {
	client := NewHttpClient(config)
	return ImportScanClient{
		baseEndpoint: "/api/v2/import-scan/",
		config:       config,
		http:         &client,
	}
}

type ImportScanClient struct {
	baseEndpoint string
	config       config.Config
	http         *HttpClient
}

func (ic ImportScanClient) ListScan(formatJson bool) {
	if formatJson {
		data, _ := json.Marshal(scanType)
		fmt.Print(data)
		os.Exit(0)
	}
	for _, v := range scanType {
		fmt.Println(v)
	}
}

func (ic ImportScanClient) Send(scan ImportScan) (ImportScanResponse, error) {
	if !contains(scanType, scan.ScanType) {
		return ImportScanResponse{}, ErrScanTypeNotFound
	}
	body := make(map[string]string)
	body["engagement"] = strconv.Itoa(scan.EngagementId)
	body["scan_type"] = scan.ScanType

	resp, err := ic.http.Multipart(ic.baseEndpoint, body, "file", scan.File)
	if err != nil {
		return ImportScanResponse{}, err
	}
	var importScanResp ImportScanResponse
	err = json.NewDecoder(resp.Body).Decode(&importScanResp)
	return importScanResp, err
}
