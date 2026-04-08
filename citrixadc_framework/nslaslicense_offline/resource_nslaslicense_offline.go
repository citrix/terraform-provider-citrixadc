package nslaslicense_offline

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/citrix/adc-nitro-go/service"

	lasutils "github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NSLASLicenseOfflineResource{}

func NewNSLASLicenseOfflineResource() resource.Resource {
	return &NSLASLicenseOfflineResource{}
}

// NSLASLicenseOfflineResource defines the resource implementation.
type NSLASLicenseOfflineResource struct {
	client *service.NitroClient
}

// LASSecretsModel describes the LAS secrets JSON structure
type LASSecretsModel struct {
	Ccid        string `json:"ccid"`
	Client      string `json:"client"`
	Password    string `json:"password"`
	LasEndpoint string `json:"las_endpoint"`
	CcEndpoint  string `json:"cc_endpoint"`
}

func (r *NSLASLicenseOfflineResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslaslicense_offline"
}

func (r *NSLASLicenseOfflineResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NSLASLicenseOfflineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NSLASLicenseOfflineResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Validate provider configuration
	if r.client == nil {
		resp.Diagnostics.AddError(
			"Provider Not Configured",
			"The provider must be configured with endpoint, username, and password",
		)
		return
	}

	// Get provider configuration from NitroClient
	endpoint := r.client.GetURL()
	username := r.client.GetUsername()
	password := r.client.GetPassword()

	if endpoint == "" || username == "" || password == "" {
		resp.Diagnostics.AddError(
			"Provider Configuration Incomplete",
			"All provider fields (endpoint, username, password) must be configured",
		)
		return
	}

	// Extract device IP from provider endpoint
	deviceIP := extractIPFromEndpoint(endpoint)

	// Validate username
	if username != "nsroot" {
		resp.Diagnostics.AddError(
			"Invalid Username",
			"Username must be 'nsroot'",
		)
		return
	}

	// Early validation: Check if entitlement_name starts with a known prefix
	validEntitlementPrefixes := []string{
		"FIPS MPX 14",
		"FIPS MPX 15",
		"FIPS MPX 16",
		"FIPS MPX 89",
		"FIPS MPX 91",
		"FIPS MPX 92",
		"MPS 14",
		"MPX 15",
		"MPX 16",
		"MPX 17",
		"MPS 25",
		"MPX 26",
		"MPX 59",
		"MPX 89",
		"MPX 91",
		"MPX 92",
		"VPX",
	}
	entitlementNameVal := data.EntitlementName.ValueString()
	validEntitlement := false
	var matchedPrefix string
	for _, prefix := range validEntitlementPrefixes {
		if strings.HasPrefix(entitlementNameVal, prefix) {
			validEntitlement = true
			matchedPrefix = prefix
			break
		}
	}
	if !validEntitlement {
		resp.Diagnostics.AddError(
			"Invalid Entitlement Name",
			fmt.Sprintf("entitlement_name '%s' must start with a valid VPX/MPX model prefix. Valid prefixes: %s",
				entitlementNameVal, strings.Join(validEntitlementPrefixes, ", ")),
		)
		return
	}

	// // Early validation: Check if request_pem is valid before starting expensive operations
	// requestPEM := data.RequestPEM.ValueString()
	// if requestPEM != "" {
	// 	if _, ok := lasutils.PEMEntitlementMapping[requestPEM]; !ok {
	// 		resp.Diagnostics.AddError(
	// 			"Invalid Request PEM",
	// 			fmt.Sprintf("Request PEM '%s' is not a valid Platform Entitlement Model code. "+
	// 				"Please refer to the provider documentation for valid PEM codes.", requestPEM),
	// 		)
	// 		return
	// 	}
	// 	tflog.Debug(ctx, "Request PEM validation passed", map[string]interface{}{"pem": requestPEM})
	// }

	// Read LAS secrets from file
	lasSecretsPath := data.LASSecretsJson.ValueString()
	lasSecretsData, err := os.ReadFile(lasSecretsPath)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Read LAS Secrets File",
			fmt.Sprintf("Cannot read file %s: %s", lasSecretsPath, err.Error()),
		)
		return
	}

	var lasSecrets LASSecretsModel
	if err := json.Unmarshal(lasSecretsData, &lasSecrets); err != nil {
		resp.Diagnostics.AddError(
			"Failed to Parse LAS Secrets JSON",
			fmt.Sprintf("Invalid JSON in file %s: %s", lasSecretsPath, err.Error()),
		)
		return
	}

	// Validate LAS secrets
	if lasSecrets.Ccid == "" || lasSecrets.Client == "" || lasSecrets.Password == "" ||
		lasSecrets.LasEndpoint == "" || lasSecrets.CcEndpoint == "" {
		resp.Diagnostics.AddError(
			"Incomplete LAS Secrets",
			"LAS secrets JSON must contain: ccid, client, password, las_endpoint, cc_endpoint",
		)
		return
	}

	// // Validate FIPS settings
	// if data.IsFIPS.ValueBool() {
	// 	requestPEM := data.RequestPEM.ValueString()
	// 	if requestPEM != "" && len(requestPEM) >= 7 && requestPEM[:7] == "CNS_14" {
	// 		resp.Diagnostics.AddError(
	// 			"Invalid FIPS Configuration",
	// 			"MPX14k doesn't need fips argument",
	// 		)
	// 		return
	// 	}
	// }

	tflog.Info(ctx, "Starting offline LAS license generation for NetScaler", map[string]interface{}{
		"device_ip": deviceIP,
	})

	// Step 1: Check version compatibility
	var version, build string
	var compat *lasutils.VersionCompatibility

	compat, err = lasutils.CheckNSVersion(ctx, deviceIP, username, password, data.IsFIPS.ValueBool())
	if err != nil {
		resp.Diagnostics.AddError(
			"Version Check Failed",
			fmt.Sprintf("Failed to check NetScaler version: %s", err.Error()),
		)
		return
	}

	if !compat.LasOk {
		resp.Diagnostics.AddError(
			"Incompatible Version",
			fmt.Sprintf("Device version %s build %s is not compatible with LAS: %s", compat.Version, compat.Build, compat.Reason),
		)
		return
	}

	version = compat.Version
	build = compat.Build
	data.Version = types.StringValue(version)
	data.Build = types.StringValue(build)

	tflog.Info(ctx, "Version check passed", map[string]interface{}{
		"version": version,
		"build":   build,
	})

	// Step 2: Use default hostname
	hostname := "ns"

	// Step 3: Determine if new API is needed
	useHostname := lasutils.DetermineNewAPINeeded(version, build)

	tflog.Info(ctx, "API selection", map[string]interface{}{
		"useHostname": useHostname,
		"version":     version,
		"build":       build,
	})

	// Step 4: Generate offline request package for NS
	filename, packageData, err := lasutils.GetOfflineRequestPackageNS(ctx, deviceIP, hostname, username, password, useHostname)
	if err != nil {
		resp.Diagnostics.AddError(
			"Request Package Generation Failed",
			fmt.Sprintf("Failed to generate offline request package: %s", err.Error()),
		)
		return
	}

	tflog.Info(ctx, "Generated request package", map[string]interface{}{"filename": filename})

	// Step 5: Extract LSGUID from package
	lsguid, err := lasutils.ExtractLSGUIDFromPackageNS(ctx, packageData)
	if err != nil {
		resp.Diagnostics.AddError(
			"LSGUID Extraction Failed",
			fmt.Sprintf("Failed to extract LSGUID from request package: %s", err.Error()),
		)
		return
	}
	data.LSGUID = types.StringValue(lsguid)

	tflog.Info(ctx, "Extracted LSGUID", map[string]interface{}{"lsguid": lsguid})

	// Step 6: Use entitlement_name directly
	// entitlementName, err := lasutils.GetEntitlementNameForFixedBW(data.RequestPEM.ValueString(), data.RequestED.ValueString(), data.IsFIPS.ValueBool())
	// if err != nil {
	// 	resp.Diagnostics.AddError(
	// 		"Invalid Entitlement Configuration",
	// 		fmt.Sprintf("Failed to determine entitlement name: %s", err.Error()),
	// 	)
	// 	return
	// }
	lasEndpoint := "netscalerfixedbw"
	tflog.Info(ctx, "Using entitlement name", map[string]interface{}{"entitlementName": entitlementNameVal})

	// Step 7: Initialize LAS Token Generator
	ltg := lasutils.NewLASTokenGenerator(
		lasEndpoint,
		lsguid,
		lasSecrets.Ccid,
		lasSecrets.Client,
		lasSecrets.Password,
		lasSecrets.LasEndpoint,
		lasSecrets.CcEndpoint,
	)

	// Step 8: Generate bearer token
	_, err = ltg.GenerateBearerToken(ctx)
	if err != nil {
		resp.Diagnostics.AddError(
			"Bearer Token Generation Failed",
			fmt.Sprintf("Failed to generate bearer token: %s", err.Error()),
		)
		return
	}

	// Step 8a: Fetch customer entitlements from LAS and validate entitlement_name
	customerEntitlements, err := ltg.GetCustomerEntitlements(ctx, matchedPrefix)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to Fetch Customer Entitlements",
			fmt.Sprintf("Failed to retrieve entitlements for platform '%s': %s", matchedPrefix, err.Error()),
		)
		return
	}
	validEntitlementName := false
	validNames := []string{}
	if entitlements, ok := customerEntitlements["entitlements"].([]interface{}); ok {
		for _, e := range entitlements {
			if obj, ok := e.(map[string]interface{}); ok {
				if name, ok := obj["type"].(string); ok {
					validNames = append(validNames, name)
					if name == entitlementNameVal {
						validEntitlementName = true
					}
				}
			}
		}
	}
	if !validEntitlementName {
		resp.Diagnostics.AddError(
			"Invalid Entitlement Name",
			fmt.Sprintf("entitlement_name '%s' is not available for your account. Available entitlements: %s",
				entitlementNameVal, strings.Join(validNames, ", ")),
		)
		return
	}
	tflog.Info(ctx, "Entitlement name validated against LAS", map[string]interface{}{"entitlementName": entitlementNameVal})

	// Step 9: Get fingerprint
	fingerprint, err := ltg.GetFingerprintForLSGUID(ctx)
	if err != nil {
		tflog.Warn(ctx, "Failed to get fingerprint", map[string]interface{}{"error": err.Error()})
		fingerprint = ""
	}

	tflog.Info(ctx, "Fingerprint lookup complete", map[string]interface{}{"fingerprint": fingerprint})

	// Step 10: Import offline activation request
	var importToken string
	if data.RestrictedMode.ValueBool() {
		// Restricted mode: extract lsid and pubkey from package, use JSON-based API
		lsid, pubkey, err := lasutils.ExtractLSIDAndPubKeyFromPackageNS(ctx, packageData)
		if err != nil {
			resp.Diagnostics.AddError(
				"LSID/PubKey Extraction Failed",
				fmt.Sprintf("Failed to extract lsid and pubkey from request package: %s", err.Error()),
			)
			return
		}
		tflog.Info(ctx, "Restricted mode: using JSON-based import API", map[string]interface{}{"lsid": lsid})
		importToken, err = ltg.ImportRestrictedOfflineActivationRequest(ctx, lsid, pubkey)
		if err != nil {
			resp.Diagnostics.AddError(
				"Import Restricted Request Failed",
				fmt.Sprintf("Failed to import restricted offline activation request: %s", err.Error()),
			)
			return
		}
	} else {
		importToken, err = ltg.ImportOfflineActivationRequest(ctx, packageData, fingerprint)
		if err != nil {
			resp.Diagnostics.AddError(
				"Import Request Failed",
				fmt.Sprintf("Failed to import offline activation request: %s", err.Error()),
			)
			return
		}
	}

	tflog.Info(ctx, "Import successful", map[string]interface{}{"importToken": importToken})

	// Step 11: Generate offline activation
	activationResp, err := ltg.GenerateOfflineActivation(ctx, importToken, entitlementNameVal)
	if err != nil {
		resp.Diagnostics.AddError(
			"Activation Generation Failed",
			fmt.Sprintf("Failed to generate offline activation: %s", err.Error()),
		)
		return
	}

	activationID, ok := activationResp["newactivationid"].(string)
	if !ok {
		resp.Diagnostics.AddError(
			"Invalid Activation Response",
			"Failed to extract newactivationid from activation response",
		)
		return
	}

	activationFingerprint, ok := activationResp["lsfingerprint"].(string)
	if !ok {
		activationFingerprint = fingerprint
	}

	tflog.Info(ctx, "Activation generated", map[string]interface{}{"activationID": activationID})

	// Step 12: Export offline activation response (license blob)
	licenseBlob, err := ltg.ExportOfflineActivationResponse(ctx, activationID, activationFingerprint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Export Failed",
			fmt.Sprintf("Failed to export license blob: %s", err.Error()),
		)
		return
	}

	// Step 13: Save license blob to local file
	blobPath := fmt.Sprintf("/tmp/offline_token_%s_%s_activation.blob.tgz", deviceIP, hostname)
	if err := os.WriteFile(blobPath, licenseBlob, 0644); err != nil {
		resp.Diagnostics.AddError(
			"File Save Failed",
			fmt.Sprintf("Failed to save license blob: %s", err.Error()),
		)
		return
	}
	data.LicenseBlob = types.StringValue(blobPath)

	tflog.Info(ctx, "License blob saved", map[string]interface{}{"path": blobPath})

	// Step 14: Apply license blob to NetScaler device
	err = lasutils.ApplyLicenseBlobNS(ctx, deviceIP, username, password, licenseBlob)
	if err != nil {
		resp.Diagnostics.AddError(
			"License Application Failed",
			fmt.Sprintf("Failed to apply license blob to device: %s", err.Error()),
		)
		return
	}

	// Set computed values
	data.Id = types.StringValue(deviceIP)
	data.Status = types.StringValue("applied")
	data.LastUpdated = types.StringValue(time.Now().Format(time.RFC3339))

	tflog.Info(ctx, "License applied successfully", map[string]interface{}{"device_ip": deviceIP})

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// extractIPFromEndpoint extracts IP address from endpoint URL
func extractIPFromEndpoint(endpoint string) string {
	// Remove protocol prefix if present
	endpoint = strings.TrimPrefix(endpoint, "http://")
	endpoint = strings.TrimPrefix(endpoint, "https://")

	// Remove port if present
	if idx := strings.Index(endpoint, ":"); idx != -1 {
		endpoint = endpoint[:idx]
	}

	// Remove trailing slash
	endpoint = strings.TrimSuffix(endpoint, "/")

	return endpoint
}

func (r *NSLASLicenseOfflineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NSLASLicenseOfflineResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// For offline license, verify if the license blob file still exists
	if !data.LicenseBlob.IsNull() {
		blobPath := data.LicenseBlob.ValueString()
		if _, err := os.Stat(blobPath); os.IsNotExist(err) {
			tflog.Warn(ctx, "License blob file no longer exists", map[string]interface{}{"path": blobPath})
		}
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NSLASLicenseOfflineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NSLASLicenseOfflineResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NO-OP: Offline licenses cannot be updated in place
	// Any changes to key attributes require resource replacement
	tflog.Info(ctx, "Update operation is a no-op for offline license resource")

	// Save data into Terraform state as-is
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NSLASLicenseOfflineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NSLASLicenseOfflineResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NO-OP: Offline licenses are not removed from the device on destroy
	// The license remains active on the device; only Terraform state is removed
	tflog.Info(ctx, "Delete operation is a no-op for offline license resource")
}
