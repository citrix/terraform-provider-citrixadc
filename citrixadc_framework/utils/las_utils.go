package utils

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

// PEMEntitlementMapping maps PEM codes to entitlement names for NetScaler VPX and MPX
var PEMEntitlementMapping = map[string]string{
	"CNS_8905_SERVER":   "MPX 8905",
	"CNS_8910_SERVER":   "MPX 8910",
	"CNS_8920_SERVER":   "MPX 8920",
	"CNS_8930_SERVER":   "MPX 8930",
	"CNS_9110_SERVER":   "MPX 9110",
	"CNS_9120_SERVER":   "MPX 9120",
	"CNS_9130_SERVER":   "MPX 9130",
	"CNS_5901_SERVER":   "MPX 5901",
	"CNS_5905_SERVER":   "MPX 5905",
	"CNS_5910_SERVER":   "MPX 5910",
	"CNS_14020_SERVER":  "FIPS MPX 14020",
	"CNS_14030_SERVER":  "FIPS MPX 14030",
	"CNS_14060_SERVER":  "FIPS MPX 14060",
	"CNS_14080_SERVER":  "FIPS MPX 14080",
	"CNS_14500_SERVER":  "FIPS MPX 14500",
	"CNS_16030_SERVER":  "MPX 16030",
	"CNS_16040_SERVER":  "MPX 16040",
	"CNS_16060_SERVER":  "MPX 16060",
	"CNS_16120_SERVER":  "MPX 16120",
	"CNS_16200_SERVER":  "MPX 16200",
	"CNS_15120_SERVER":  "MPX 15120 / 15120-50G",
	"CNS_26200_SERVER":  "MPX 26200 / 26200-50S / 26200-100G Premium",
	"CNS_9205_SERVER":   "MPX 9205",
	"CNS_9210_SERVER":   "MPX 9210",
	"CNS_9220_SERVER":   "MPX 9220",
	"CNS_9240_SERVER":   "MPX 9240",
	"CNS_9260_SERVER":   "MPX 9260",
	"CNS_9280_SERVER":   "MPX 9280",
	"CNS_9295_SERVER":   "MPX 9295",
	"CNS_9299_SERVER":   "MPX 9299",
	"CNS_17020_SERVER":  "MPX 17020",
	"CNS_17050_SERVER":  "MPX 17050",
	"CNS_17100_SERVER":  "MPX 17100",
	"CNS_17150_SERVER":  "MPX 17150",
	"CNS_17200_SERVER":  "MPX 17200",
	"CNS_17250_SERVER":  "MPX 17250",
	"CNS_17300_SERVER":  "MPX 17300",
	"CNS_17400_SERVER":  "MPX 17400",
	"CNS_17500_SERVER":  "MPX 17500",
	"CNS_V25000_SERVER": "VPX 25000",
	"CNS_V10000_SERVER": "VPX 10000",
	"CNS_V5000_SERVER":  "VPX 5000",
	"CNS_V3000_SERVER":  "VPX 3000",
	"CNS_V1000_SERVER":  "VPX 1000",
	"CNS_V200_SERVER":   "VPX 200",
	"CNS_V25_SERVER":    "VPX 25",
	"CNS_V10_SERVER":    "VPX 10",
}

// VersionCompatibility holds version compatibility check results
type VersionCompatibility struct {
	Version string
	Build   string
	LasOk   bool
	Reason  string
}

// LASTokenGenerator handles LAS token generation
type LASTokenGenerator struct {
	Endpoint      string
	LSGUID        string
	CCID          string
	SecretClient  string
	SecretPwd     string
	BaseURL       string
	CCTokenURL    string
	BearerCache   string
	BearerToken   string
	HTTPClient    *http.Client
	InsecureHTTPS bool
}

// NewLASTokenGenerator creates a new LAS token generator
func NewLASTokenGenerator(endpoint, lsguid, ccid, client, password, baseURL, ccTokenURL string) *LASTokenGenerator {
	return &LASTokenGenerator{
		Endpoint:      endpoint,
		LSGUID:        lsguid,
		CCID:          ccid,
		SecretClient:  client,
		SecretPwd:     password,
		BaseURL:       baseURL,
		CCTokenURL:    ccTokenURL,
		BearerCache:   "/tmp/las_bearer_cache",
		InsecureHTTPS: true,
		HTTPClient: &http.Client{
			Timeout: 60 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}
}

// RunCurlHTTPSFallback tries HTTPS first, falls back to HTTP
func RunCurlHTTPSFallback(ctx context.Context, url, method string, auth *BasicAuth, body []byte, headers map[string]string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// Try HTTPS first
	httpsURL := strings.Replace(url, "http://", "https://", 1)
	resp, err := makeHTTPRequest(ctx, client, httpsURL, method, auth, body, headers)
	if err == nil {
		return resp, nil
	}

	tflog.Debug(ctx, "HTTPS request failed, falling back to HTTP", map[string]interface{}{"error": err.Error()})

	// Fallback to HTTP
	httpURL := strings.Replace(httpsURL, "https://", "http://", 1)
	return makeHTTPRequest(ctx, client, httpURL, method, auth, body, headers)
}

// BasicAuth holds basic authentication credentials
type BasicAuth struct {
	Username string
	Password string
}

func makeHTTPRequest(ctx context.Context, client *http.Client, url, method string, auth *BasicAuth, body []byte, headers map[string]string) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = bytes.NewReader(body)
	}

	// Log request details
	tflog.Debug(ctx, "HTTP Request", map[string]interface{}{
		"method": method,
		"url":    url,
	})
	if body != nil && len(body) > 0 {
		tflog.Debug(ctx, "Request Body", map[string]interface{}{
			"body": string(body),
		})
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if auth != nil {
		req.SetBasicAuth(auth.Username, auth.Password)
		tflog.Debug(ctx, "Using Basic Auth", map[string]interface{}{"username": auth.Username})
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Log response details
	tflog.Debug(ctx, "HTTP Response", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if len(respBody) > 0 {
		tflog.Debug(ctx, "Response Body", map[string]interface{}{
			"body": string(respBody),
		})
	}

	if resp.StatusCode >= 400 {
		return respBody, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// CheckNSVersion checks NetScaler version and LAS compatibility
func CheckNSVersion(ctx context.Context, ip, username, password string, isFIPS bool) (*VersionCompatibility, error) {
	url := fmt.Sprintf("http://%s/nitro/v1/config/nsversion", ip)
	auth := &BasicAuth{Username: username, Password: password}

	respBody, err := RunCurlHTTPSFallback(ctx, url, "GET", auth, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get NS version: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	nsversion, ok := result["nsversion"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("missing 'nsversion' in response")
	}

	versionStr, ok := nsversion["version"].(string)
	if !ok || versionStr == "" {
		return nil, fmt.Errorf("missing 'version' field")
	}

	tflog.Info(ctx, "NetScaler version", map[string]interface{}{"ip": ip, "version": versionStr})

	// Extract version (e.g., "NetScaler NS14.1: Build 4.3401.a.nc")
	versionRegex := regexp.MustCompile(`NS(\d+\.\d+)`)
	versionMatch := versionRegex.FindStringSubmatch(versionStr)
	if len(versionMatch) < 2 {
		return nil, fmt.Errorf("unable to parse version from: %s", versionStr)
	}
	version := versionMatch[1]

	// Extract build (major.minor)
	buildRegex := regexp.MustCompile(`Build\s+(\d+)\.(\d+)`)
	buildMatch := buildRegex.FindStringSubmatch(versionStr)
	if len(buildMatch) < 3 {
		return nil, fmt.Errorf("unable to parse build from: %s", versionStr)
	}

	majorBuild, _ := strconv.Atoi(buildMatch[1])
	minorBuild, _ := strconv.Atoi(buildMatch[2])
	build := fmt.Sprintf("%d.%d", majorBuild, minorBuild)

	// Check LAS compatibility
	compat := &VersionCompatibility{
		Version: version,
		Build:   build,
	}

	switch version {
	case "14.1":
		if isBuildGE(majorBuild, minorBuild, 51, 80) {
			compat.LasOk = true
			compat.Reason = "Meets minimum required build 14.1-51.80"
		} else {
			compat.Reason = "Minimum required build is 14.1-51.80"
		}
	case "13.1":
		if isFIPS {
			if isBuildGE(majorBuild, minorBuild, 37, 247) {
				compat.LasOk = true
				compat.Reason = "Meets minimum required build 13.1-37.247 (FIPS)"
			} else {
				compat.Reason = "Minimum required build is 13.1-37.247 (FIPS)"
			}
		} else {
			if isBuildGE(majorBuild, minorBuild, 60, 29) {
				compat.LasOk = true
				compat.Reason = "Meets minimum required build 13.1-60.29"
			} else {
				compat.Reason = "Minimum required build is 13.1-60.29"
			}
		}
	default:
		compat.Reason = fmt.Sprintf("Unsupported version %s for LAS compatibility", version)
	}

	return compat, nil
}

func isBuildGE(aMajor, aMinor, bMajor, bMinor int) bool {
	return (aMajor > bMajor) || (aMajor == bMajor && aMinor >= bMinor)
}

// GetOfflineRequestPackageNS generates offline activation request package for NetScaler
func GetOfflineRequestPackageNS(ctx context.Context, ip, hostname, username, password string, useHostname bool) (string, []byte, error) {
	auth := &BasicAuth{Username: username, Password: password}
	var url string

	if useHostname {
		url = fmt.Sprintf("http://%s/nitro/v1/config/nslicenseactivationdata?args=usehostname:true", ip)
	} else {
		url = fmt.Sprintf("http://%s/nitro/v1/config/nslicenseactivationdata", ip)
	}

	// Make API call to generate request package
	respBody, err := RunCurlHTTPSFallback(ctx, url, "GET", auth, nil, nil)
	if err != nil {
		return "", nil, fmt.Errorf("failed to generate request package: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	// Extract filename
	var filename string
	if data, ok := result["nslicenseactivationdata"].(map[string]interface{}); ok {
		filename, _ = data["filename"].(string)
	}

	if filename == "" {
		return "", nil, fmt.Errorf("failed to extract filename from response")
	}

	// Download the file via SCP
	remotePath := "/nsconfig/license/" + filename
	fileContent, err := SCPDownload(ctx, ip, username, password, remotePath)
	if err != nil {
		return "", nil, fmt.Errorf("failed to download request package: %w", err)
	}

	tflog.Info(ctx, "Generated request package for NetScaler", map[string]interface{}{"ip": ip, "filename": filename})
	return filename, fileContent, nil
}

// SCPDownload downloads a file via SCP
func SCPDownload(ctx context.Context, ip, username, password, remotePath string) ([]byte, error) {
	tflog.Debug(ctx, "Starting SFTP download", map[string]interface{}{
		"ip":         ip,
		"remotePath": remotePath,
	})

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		tflog.Error(ctx, "Failed to dial SSH", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Create SFTP client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		tflog.Error(ctx, "Failed to create SFTP client", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to create SFTP client: %w", err)
	}
	defer sftpClient.Close()

	tflog.Debug(ctx, "Opening remote file via SFTP", map[string]interface{}{"path": remotePath})

	// Open the remote file
	remoteFile, err := sftpClient.Open(remotePath)
	if err != nil {
		tflog.Error(ctx, "Failed to open remote file", map[string]interface{}{
			"error": err.Error(),
			"path":  remotePath,
		})
		return nil, fmt.Errorf("failed to open remote file: %w", err)
	}
	defer remoteFile.Close()

	// Read the file content
	var buf bytes.Buffer
	bytesRead, err := io.Copy(&buf, remoteFile)
	if err != nil {
		tflog.Error(ctx, "Failed to read remote file", map[string]interface{}{"error": err.Error()})
		return nil, fmt.Errorf("failed to read remote file: %w", err)
	}

	tflog.Debug(ctx, "SFTP download successful", map[string]interface{}{
		"bytesDownloaded": bytesRead,
	})
	return buf.Bytes(), nil
}

// SCPUpload uploads a file via SFTP
func SCPUpload(ctx context.Context, ip, username, password, remotePath string, content []byte) error {
	tflog.Debug(ctx, "Starting SFTP upload", map[string]interface{}{
		"ip":         ip,
		"remotePath": remotePath,
		"size":       len(content),
	})

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         30 * time.Second,
	}

	client, err := ssh.Dial("tcp", ip+":22", config)
	if err != nil {
		tflog.Error(ctx, "Failed to dial SSH", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to dial: %w", err)
	}
	defer client.Close()

	// Create SFTP client
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		tflog.Error(ctx, "Failed to create SFTP client", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to create SFTP client: %w", err)
	}
	defer sftpClient.Close()

	tflog.Debug(ctx, "Creating remote file via SFTP", map[string]interface{}{"path": remotePath})

	// Create the remote file
	remoteFile, err := sftpClient.Create(remotePath)
	if err != nil {
		tflog.Error(ctx, "Failed to create remote file", map[string]interface{}{
			"error": err.Error(),
			"path":  remotePath,
		})
		return fmt.Errorf("failed to create remote file: %w", err)
	}
	defer remoteFile.Close()

	// Write content to the file
	bytesWritten, err := remoteFile.Write(content)
	if err != nil {
		tflog.Error(ctx, "Failed to write to remote file", map[string]interface{}{"error": err.Error()})
		return fmt.Errorf("failed to write to remote file: %w", err)
	}

	tflog.Info(ctx, "Uploaded file via SFTP", map[string]interface{}{
		"remotePath":   remotePath,
		"bytesWritten": bytesWritten,
	})
	return nil
}

// ExtractLSGUIDFromPackageNS extracts LSGUID from NetScaler request package
func ExtractLSGUIDFromPackageNS(ctx context.Context, packageData []byte) (string, error) {
	jsonFile := "ns_offline_activation_request.json"

	// Decompress and extract
	gzReader, err := gzip.NewReader(bytes.NewReader(packageData))
	if err != nil {
		return "", fmt.Errorf("failed to create gzip reader: %w", err)
	}
	defer gzReader.Close()

	tarReader := tar.NewReader(gzReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read tar: %w", err)
		}

		if header.Name == jsonFile || strings.HasSuffix(header.Name, "/"+jsonFile) {
			jsonData, err := io.ReadAll(tarReader)
			if err != nil {
				return "", fmt.Errorf("failed to read JSON file: %w", err)
			}

			var result map[string]interface{}
			if err := json.Unmarshal(jsonData, &result); err != nil {
				return "", fmt.Errorf("failed to parse JSON: %w", err)
			}

			lsguid, ok := result["lsguid"].(string)
			if !ok || lsguid == "" {
				return "", fmt.Errorf("lsguid not found in JSON")
			}

			tflog.Info(ctx, "Extracted LSGUID from NetScaler package", map[string]interface{}{"lsguid": lsguid})
			return lsguid, nil
		}
	}

	return "", fmt.Errorf("JSON file %s not found in package", jsonFile)
}

// GenerateBearerToken generates a bearer token for LAS API
func (ltg *LASTokenGenerator) GenerateBearerToken(ctx context.Context) (string, error) {
	payload := map[string]string{
		"clientId":     ltg.SecretClient,
		"clientSecret": ltg.SecretPwd,
	}
	body, _ := json.Marshal(payload)

	// Log request (sanitize sensitive data)
	tflog.Debug(ctx, "API Call: GenerateBearerToken", map[string]interface{}{
		"url":    ltg.CCTokenURL,
		"method": "POST",
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"clientId":     ltg.SecretClient,
		"clientSecret": "***REDACTED***",
	})

	req, err := http.NewRequestWithContext(ctx, "POST", ltg.CCTokenURL, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to generate bearer token: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Bearer Token Generation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return "", fmt.Errorf("token not found in response")
	}

	ltg.BearerToken = token
	tflog.Info(ctx, "Generated bearer token")
	tflog.Debug(ctx, "Response Payload", map[string]interface{}{
		"token": "***REDACTED***",
	})
	return token, nil
}

// GetCustomerEntitlements fetches valid entitlements from the LAS endpoint for a given platform
func (ltg *LASTokenGenerator) GetCustomerEntitlements(ctx context.Context, platform string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s/customerentitlements", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	payload := map[string]string{
		"ver":      "1.0",
		"platform": platform,
	}
	body, _ := json.Marshal(payload)

	tflog.Debug(ctx, "API Call: GetCustomerEntitlements", map[string]interface{}{
		"url":      url,
		"method":   "POST",
		"platform": platform,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get customer entitlements: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Get Customer Entitlements Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
			"platform":    platform,
		})
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	tflog.Info(ctx, "Fetched customer entitlements", map[string]interface{}{"platform": platform})
	return result, nil
}

// GetFingerprintForLSGUID gets fingerprint for LSGUID and deregisters if exists
func (ltg *LASTokenGenerator) GetFingerprintForLSGUID(ctx context.Context) (string, error) {
	url := fmt.Sprintf("%s/support/%s/%s/listls", ltg.BaseURL, ltg.CCID, ltg.Endpoint)
	payload := map[string]string{"ver": "1.0"}
	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: GetFingerprintForLSGUID", map[string]interface{}{
		"url":    url,
		"method": "POST",
		"lsguid": ltg.LSGUID,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to list license servers: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	// Find and deregister if exists
	if lstlasactivatedls, ok := result["lstlasactivatedls"].([]interface{}); ok {
		for _, ls := range lstlasactivatedls {
			if lsObj, ok := ls.(map[string]interface{}); ok {
				if lsguid, _ := lsObj["lsguid"].(string); lsguid == ltg.LSGUID {
					fingerprint, _ := lsObj["lsfingerprint"].(string)
					return fingerprint, nil
				}
			}
		}
	}

	return "", nil
}

// ImportOfflineActivationRequest imports the offline activation request
func (ltg *LASTokenGenerator) ImportOfflineActivationRequest(ctx context.Context, requestPackage []byte, fingerprint string) (string, error) {
	url := fmt.Sprintf("%s/support/%s/%s/importofflineactivationrequest", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	// Create multipart form
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Write file part
	part, err := writer.CreateFormFile("file", "request.tgz")
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}
	if _, err := part.Write(requestPackage); err != nil {
		return "", fmt.Errorf("failed to write file part: %w", err)
	}

	// Write data part
	dataJSON, _ := json.Marshal(map[string]string{
		"ver":           "1.0",
		"lsfingerprint": fingerprint,
	})
	if err := writer.WriteField("data", string(dataJSON)); err != nil {
		return "", fmt.Errorf("failed to write data field: %w", err)
	}

	writer.Close()

	// Log request
	tflog.Debug(ctx, "API Call: ImportOfflineActivationRequest", map[string]interface{}{
		"url":          url,
		"method":       "POST",
		"fingerprint":  fingerprint,
		"package_size": len(requestPackage),
	})
	tflog.Debug(ctx, "Request Data Field", map[string]interface{}{
		"data": string(dataJSON),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, &buf)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to import request: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Import Request Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return "", fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", fmt.Errorf("invalid JSON response: %w", err)
	}

	importToken, ok := result["importrequesttoken"].(string)
	if !ok || importToken == "" {
		return "", fmt.Errorf("importrequesttoken not found in response")
	}

	tflog.Info(ctx, "Imported offline activation request", map[string]interface{}{"importToken": importToken})
	return importToken, nil
}

// GenerateOfflineActivation generates the offline activation
func (ltg *LASTokenGenerator) GenerateOfflineActivation(ctx context.Context, importToken, entitlementName string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s/generateofflineactivation", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	payload := map[string]interface{}{
		"ver":                "1.0",
		"importrequesttoken": importToken,
	}

	if ltg.Endpoint == "netscalerfixedbw" && entitlementName != "" {
		payload["entitlementname"] = entitlementName
	}

	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: GenerateOfflineActivation", map[string]interface{}{
		"url":    url,
		"method": "POST",
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate activation: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Generate Activation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("invalid JSON response: %w", err)
	}

	tflog.Info(ctx, "Generated offline activation", map[string]interface{}{"result": result})
	return result, nil
}

// ExportOfflineActivationResponse exports the license blob
func (ltg *LASTokenGenerator) ExportOfflineActivationResponse(ctx context.Context, activationID, fingerprint string) ([]byte, error) {
	url := fmt.Sprintf("%s/support/%s/%s/exportofflineactivationresponse", ltg.BaseURL, ltg.CCID, ltg.Endpoint)

	payload := map[string]string{
		"ver":             "1.0",
		"lsfingerprint":   fingerprint,
		"newactivationid": activationID,
	}
	body, _ := json.Marshal(payload)

	// Log request
	tflog.Debug(ctx, "API Call: ExportOfflineActivationResponse", map[string]interface{}{
		"url":          url,
		"method":       "POST",
		"activationID": activationID,
		"fingerprint":  fingerprint,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": string(body),
	})

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "CWSAuth bearer="+ltg.BearerToken)

	resp, err := ltg.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to export blob: %w", err)
	}
	defer resp.Body.Close()

	tflog.Debug(ctx, "Response Status", map[string]interface{}{
		"status_code": resp.StatusCode,
	})

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	tflog.Debug(ctx, "Response Body Size", map[string]interface{}{
		"size_bytes": len(respBody),
	})

	if resp.StatusCode >= 400 {
		tflog.Error(ctx, "Export Activation Failed", map[string]interface{}{
			"status_code": resp.StatusCode,
			"response":    string(respBody),
		})
		return nil, fmt.Errorf("HTTP error %d: %s", resp.StatusCode, string(respBody))
	}

	tflog.Info(ctx, "Exported offline activation response")
	return respBody, nil
}

// ApplyLicenseBlobNS applies license blob to NetScaler
func ApplyLicenseBlobNS(ctx context.Context, ip, username, password string, blobContent []byte) error {
	// Create temp filename
	filename := fmt.Sprintf("offline_token_%s_activation.blob.tgz", ip)

	// Upload blob to device
	remotePath := "/nsconfig/license/" + filename
	if err := SCPUpload(ctx, ip, username, password, remotePath, blobContent); err != nil {
		return fmt.Errorf("failed to upload license blob: %w", err)
	}

	// Apply license - NITRO API expects form-encoded data with "object" key containing JSON
	auth := &BasicAuth{Username: username, Password: password}
	url := fmt.Sprintf("http://%s/nitro/v1/config/nslaslicense", ip)

	payload := map[string]interface{}{
		"params": map[string]string{
			"action":  "apply",
			"warning": "YES",
		},
		"nslaslicense": map[string]interface{}{
			"filename":       filename,
			"filelocation":   "/nsconfig/license",
			"fixedbandwidth": true,
		},
	}
	payloadJSON, _ := json.Marshal(payload)

	// Form-encode with "object" key
	formData := fmt.Sprintf("object=%s", string(payloadJSON))
	body := []byte(formData)
	headers := map[string]string{"Content-Type": "application/x-www-form-urlencoded"}

	// Log request
	tflog.Debug(ctx, "API Call: ApplyLicenseBlobNS", map[string]interface{}{
		"url":      url,
		"method":   "POST",
		"ip":       ip,
		"filename": filename,
	})
	tflog.Debug(ctx, "Request Payload", map[string]interface{}{
		"body": formData,
	})

	respBody, err := RunCurlHTTPSFallback(ctx, url, "POST", auth, body, headers)
	if err != nil {
		return fmt.Errorf("failed to apply license: %w", err)
	}

	tflog.Debug(ctx, "Response Body", map[string]interface{}{
		"body": string(respBody),
	})

	// Check for error code 1043 (invalid license blob)
	if strings.Contains(string(respBody), "\"errorcode\": 1043") {
		tflog.Error(ctx, "Invalid license blob", map[string]interface{}{"ip": ip, "response": string(respBody)})
		return fmt.Errorf("invalid license blob (error code 1043): the license file is invalid or incompatible with this device")
	}

	// Check if reboot is needed (error code 1125)
	if strings.Contains(string(respBody), "\"errorcode\": 1125") {
		tflog.Info(ctx, "Rebooting device after license application", map[string]interface{}{"ip": ip})
		// Trigger reboot
		rebootURL := fmt.Sprintf("http://%s/nitro/v1/config/reboot", ip)
		rebootPayload := map[string]interface{}{
			"params": map[string]string{
				"warning": "YES",
			},
			"reboot": map[string]bool{
				"warm": true,
			},
		}
		rebootJSON, _ := json.Marshal(rebootPayload)
		rebootFormData := fmt.Sprintf("object=%s", string(rebootJSON))
		rebootBody := []byte(rebootFormData)
		RunCurlHTTPSFallback(ctx, rebootURL, "POST", auth, rebootBody, headers)
	} else if !strings.Contains(string(respBody), "\"errorcode\": 0") {
		// If we don't get error code 0 (success) or 1125 (reboot required), something went wrong
		tflog.Warn(ctx, "Unexpected response from license application", map[string]interface{}{
			"ip":       ip,
			"response": string(respBody),
		})
	}

	tflog.Info(ctx, "Applied license blob", map[string]interface{}{"ip": ip})
	return nil
}

// GetEntitlementNameForFixedBW gets the entitlement name for NetScaler fixed bandwidth
func GetEntitlementNameForFixedBW(requestPEM, requestED string, isFIPS bool) (string, error) {
	baseEntString, ok := PEMEntitlementMapping[requestPEM]
	if !ok {
		return "", fmt.Errorf("PEM not found: %s", requestPEM)
	}

	// FIPS validation
	if isFIPS {
		fipsPEMs := []string{"CNS_8910_SERVER", "CNS_8920_SERVER", "CNS_9130_SERVER", "CNS_15120_SERVER",
			"CNS_V5000_SERVER", "CNS_V3000_SERVER", "CNS_V1000_SERVER", "CNS_V200_SERVER", "CNS_V25_SERVER"}
		found := false
		for _, pem := range fipsPEMs {
			if requestPEM == pem {
				found = true
				break
			}
		}
		if !found {
			return "", fmt.Errorf("FIPS not supported for PEM: %s", requestPEM)
		}

		mpxFipsPEMs := []string{"CNS_8910_SERVER", "CNS_8920_SERVER", "CNS_9130_SERVER", "CNS_15120_SERVER"}
		for _, pem := range mpxFipsPEMs {
			if requestPEM == pem && requestED != "Premium" {
				return "", fmt.Errorf("MPX FIPS only supported for Premium edition")
			}
		}

		if requestPEM == "CNS_15120_SERVER" {
			baseEntString = "FIPS MPX 15120-50G"
		} else {
			baseEntString = "FIPS " + baseEntString
		}
	}

	// NetScaler edition formatting
	if requestED == "Advanced" || requestED == "Standard" || requestED == "Premium" {
		return baseEntString + " " + requestED, nil
	}

	return "", fmt.Errorf("invalid edition: %s", requestED)
}

// DetermineNewAPINeeded checks if new API is needed based on NetScaler version
func DetermineNewAPINeeded(version, build string) bool {
	parts := strings.Split(build, ".")
	if len(parts) < 2 {
		return false
	}
	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])

	// Version 14.1 checks
	if version == "14.1" {
		if major > 68 {
			return true
		}
		if major == 68 && minor >= 3 {
			return true
		}
		if major == 60 && minor >= 55 {
			return true
		}
		if major == 66 && minor >= 32 {
			return true
		}
	}

	// Version 13.1 checks for NetScaler
	if version == "13.1" {
		if major > 62 {
			return true
		}
		if major == 62 && minor >= 6 {
			return true
		}
		if major == 61 && minor >= 26 {
			return true
		}
		// FIPS
		if major == 37 && minor >= 256 {
			return true
		}
	}

	return false
}
