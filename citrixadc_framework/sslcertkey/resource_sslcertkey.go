package sslcertkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslCertKeyResource{}
var _ resource.ResourceWithConfigure = (*SslCertKeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslCertKeyResource)(nil)

func NewSslCertKeyResource() resource.Resource {
	return &SslCertKeyResource{}
}

// SslCertKeyResource defines the resource implementation.
type SslCertKeyResource struct {
	client *service.NitroClient
}

func (r *SslCertKeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslCertKeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcertkey"
}

func (r *SslCertKeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslCertKeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslCertKeyResourceModel
	var config SslCertKeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to access write-only attributes (like passplain)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcertkey resource")

	// Get certkey name - if not provided, generate one
	var sslcertkeyName string
	if data.Certkey.IsNull() || data.Certkey.ValueString() == "" {
		sslcertkeyName = "tf-sslcertkey-" + fmt.Sprintf("%d", len(data.Cert.ValueString()))
		data.Certkey = types.StringValue(sslcertkeyName)
	} else {
		sslcertkeyName = data.Certkey.ValueString()
	}

	// Use config (not plan) to get write-only attributes like passplain
	sslcertkey := sslcertkeyGetThePayloadFromtheConfig(ctx, &config)

	// Make API call
	_, err := r.client.AddResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkey)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcertkey, got error: %s", err))
		return
	}

	data.Id = types.StringValue(sslcertkeyName)

	tflog.Trace(ctx, "Created sslcertkey resource")

	// Handle linked certificate if configured
	if !data.LinkCertKeyName.IsNull() && data.LinkCertKeyName.ValueString() != "" {
		if err := r.handleLinkedCertificate(ctx, &data, &resp.Diagnostics); err != nil {
			tflog.Error(ctx, "Error linking certificate during creation")
			// If linking fails, delete the created certificate
			delErr := r.client.DeleteResource(service.Sslcertkey.Type(), sslcertkeyName)
			if delErr != nil {
				resp.Diagnostics.AddError("Cleanup Error",
					fmt.Sprintf("Failed to delete certificate after link error. Link error: %s, Delete error: %s", err, delErr))
			} else {
				resp.Diagnostics.AddError("Client Error",
					fmt.Sprintf("Failed to link certificate: %s. Certificate has been deleted.", err))
			}
			return
		}
	}

	sslcertkeyUpdate := ssl.Sslcertkey{
		Certkey: sslcertkeyName,
		Cert:    data.Cert.ValueString(),
	}

	// Nodomaincheck is set to true post creation
	if !data.NoDomainCheck.IsNull() {
		sslcertkeyUpdate.Nodomaincheck = data.NoDomainCheck.ValueBool()
	}
	if !data.Key.IsNull() {
		sslcertkeyUpdate.Key = data.Key.ValueString()
	}
	if !data.Password.IsNull() {
		sslcertkeyUpdate.Password = data.Password.ValueBool()
	}
	if !data.Passplain.IsNull() {
		sslcertkeyUpdate.Passplain = data.Passplain.ValueString()
	} else if !data.PassplainWo.IsNull() {
		sslcertkeyUpdate.Passplain = data.PassplainWo.ValueString()
	}

	// Update resource using "update" action after create.
	err = r.client.ActOnResource(service.Sslcertkey.Type(), &sslcertkeyUpdate, "update")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkey, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readSslCertKeyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslCertKeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslCertKeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcertkey resource")
	r.readSslCertKeyFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslCertKeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state, config SslCertKeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating sslcertkey resource")

	sslcertkeyName := plan.Certkey.ValueString()

	// Determine which type of update is needed
	needsUpdate := false
	needsChange := false
	needsClear := false

	sslcertkeyUpdate := ssl.Sslcertkey{
		Certkey: sslcertkeyName,
	}
	sslcertkeyChange := ssl.Sslcertkey{
		Certkey: sslcertkeyName,
	}
	sslcertkeyClear := ssl.Sslcertkey{
		Certkey: sslcertkeyName,
	}

	// Check for changes that require Update API
	if !plan.Expirymonitor.Equal(state.Expirymonitor) {
		tflog.Debug(ctx, "Expirymonitor has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyUpdate.Expirymonitor = plan.Expirymonitor.ValueString()
		needsUpdate = true
	}
	if !plan.NotificationPeriod.Equal(state.NotificationPeriod) {
		tflog.Debug(ctx, "Notificationperiod has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		if !plan.NotificationPeriod.IsNull() {
			sslcertkeyUpdate.Notificationperiod = utils.IntPtr(int(plan.NotificationPeriod.ValueInt64()))
		}
		needsUpdate = true
	}
	if !plan.DeleteCertKeyFilesOnRemoval.Equal(state.DeleteCertKeyFilesOnRemoval) {
		tflog.Debug(ctx, "DeleteCertKeyFilesOnRemoval has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyUpdate.Deletecertkeyfilesonremoval = plan.DeleteCertKeyFilesOnRemoval.ValueString()
		needsUpdate = true
	}

	// Check for changes that require Change API
	if !plan.NoDomainCheck.Equal(state.NoDomainCheck) {
		tflog.Debug(ctx, "NoDomainCheck has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Nodomaincheck = plan.NoDomainCheck.ValueBool()
		needsChange = true
	}
	if !plan.Cert.Equal(state.Cert) {
		tflog.Debug(ctx, "Cert has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Cert = plan.Cert.ValueString()
		needsChange = true
	}
	if !plan.Key.Equal(state.Key) {
		tflog.Debug(ctx, "Key has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Key = plan.Key.ValueString()
		needsChange = true
	}
	// Check if certificate or key file content has changed by comparing hashes
	if !plan.Cert.IsNull() && !plan.CertHash.Equal(state.CertHash) {
		tflog.Debug(ctx, "Cert file content has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Cert = plan.Cert.ValueString()
		needsChange = true
	}
	if !plan.Key.IsNull() && !plan.KeyHash.Equal(state.KeyHash) {
		tflog.Debug(ctx, "Key file content has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Key = plan.Key.ValueString()
		needsChange = true
	}
	if !plan.Password.Equal(state.Password) {
		tflog.Debug(ctx, "Password has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Password = plan.Password.ValueBool()
		needsChange = true
	}
	if !plan.Fipskey.Equal(state.Fipskey) {
		tflog.Debug(ctx, "Fipskey has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Fipskey = plan.Fipskey.ValueString()
		needsChange = true
	}
	if !plan.Inform.Equal(state.Inform) {
		tflog.Debug(ctx, "Inform has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyChange.Inform = plan.Inform.ValueString()
		needsChange = true
	}

	// Check for changes that require Clear API
	if !config.OcspStaplingCache.IsNull() && !plan.OcspStaplingCache.Equal(state.OcspStaplingCache) {
		tflog.Debug(ctx, "OcspStaplingCache has changed for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
		sslcertkeyClear.Ocspstaplingcache = plan.OcspStaplingCache.ValueBool()
		needsClear = true
	}

	// Execute Update API if needed
	if needsUpdate {
		// Expirymonitor is always expected by NITRO API
		if !plan.Expirymonitor.IsNull() {
			sslcertkeyUpdate.Expirymonitor = plan.Expirymonitor.ValueString()
		}
		_, err := r.client.UpdateResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyUpdate)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslcertkey %s, got error: %s", sslcertkeyName, err))
			return
		}
	}

	// Execute Change API if needed
	if needsChange {
		// Nodomaincheck is a flag for the change operation
		if !plan.NoDomainCheck.IsNull() {
			sslcertkeyChange.Nodomaincheck = plan.NoDomainCheck.ValueBool()
		}
		// Include password if configured (required for PKCS12 cert/key changes)
		// Use config value since passplain is WriteOnly
		if !config.Passplain.IsNull() && config.Passplain.ValueString() != "" {
			tflog.Debug(ctx, "Including passplain in change operation for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
			sslcertkeyChange.Passplain = config.Passplain.ValueString()
		} else if !config.PassplainWo.IsNull() && config.PassplainWo.ValueString() != "" {
			tflog.Debug(ctx, "Including passplain_wo in change operation for sslcertkey", map[string]interface{}{"certkey": sslcertkeyName})
			sslcertkeyChange.Passplain = config.PassplainWo.ValueString()
		}
		_, err := r.client.ChangeResource(service.Sslcertkey.Type(), sslcertkeyName, &sslcertkeyChange)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change sslcertkey %s, got error: %s", sslcertkeyName, err))
			return
		}
	}

	// Execute Clear API if needed
	if needsClear {
		err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslcertkeyClear, "clear")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to clear sslcertkey %s, got error: %s", sslcertkeyName, err))
			return
		}
	}

	// Handle linked certificate changes
	if err := r.handleLinkedCertificate(ctx, &plan, &resp.Diagnostics); err != nil {
		tflog.Error(ctx, "Error linking certificate during update")
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to handle linked certificate for %s, got error: %s", sslcertkeyName, err))
		return
	}

	tflog.Trace(ctx, "Updated sslcertkey resource")

	// Ensure plan has the correct ID before reading from API
	plan.Id = state.Id

	// Read the updated state back - update plan with fresh API data
	r.readSslCertKeyFromApi(ctx, &plan, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SslCertKeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslCertKeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcertkey resource")

	sslcertkeyName := data.Id.ValueString()

	// Unlink certificate before deletion
	if err := r.unlinkCertificate(ctx, &data, &resp.Diagnostics); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unlink certificate %s, got error: %s", sslcertkeyName, err))
		return
	}

	// Build delete arguments
	args := make([]string, 0)
	if !data.DeleteFromDevice.IsNull() {
		args = append(args, fmt.Sprintf("deletefromdevice:%t", data.DeleteFromDevice.ValueBool()))
	}
	if !data.DeleteCertKeyFilesOnRemoval.IsNull() {
		args = append(args, fmt.Sprintf("deletecertkeyfilesonremoval:%s", data.DeleteCertKeyFilesOnRemoval.ValueString()))
	}

	// Make API call with arguments
	err := r.client.DeleteResourceWithArgs(service.Sslcertkey.Type(), sslcertkeyName, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcertkey %s, got error: %s", sslcertkeyName, err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcertkey resource")
}

// Helper function to read sslcertkey data from API
func (r *SslCertKeyResource) readSslCertKeyFromApi(ctx context.Context, data *SslCertKeyResourceModel, diags *diag.Diagnostics) {
	sslcertkeyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Clearing sslcertkey state %s", sslcertkeyName))
		data.Id = types.StringNull()
		return
	}

	sslcertkeySetAttrFromGet(ctx, data, getResponseData)
}

// Helper function to handle linked certificate
func (r *SslCertKeyResource) handleLinkedCertificate(ctx context.Context, data *SslCertKeyResourceModel, diags *diag.Diagnostics) error {
	tflog.Debug(ctx, "In handleLinkedCertificate")

	sslcertkeyName := data.Certkey.ValueString()

	// Get current state from API
	getResponseData, err := r.client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error finding sslcertkey %s", sslcertkeyName))
		data.Id = types.StringNull()
		return err
	}

	// Get actual and configured linked certificate names
	var actualLinkedCertKeyname interface{} = nil
	if val, ok := getResponseData["linkcertkeyname"]; ok {
		actualLinkedCertKeyname = val
	}

	configuredLinkedCertKeyname := ""
	if !data.LinkCertKeyName.IsNull() {
		configuredLinkedCertKeyname = data.LinkCertKeyName.ValueString()
	}

	// Check for noop conditions
	if actualLinkedCertKeyname != nil && actualLinkedCertKeyname.(string) == configuredLinkedCertKeyname {
		tflog.Debug(ctx, fmt.Sprintf("actual and configured linked certificates identical: %s", actualLinkedCertKeyname))
		return nil
	}

	if actualLinkedCertKeyname == nil && configuredLinkedCertKeyname == "" {
		tflog.Debug(ctx, "actual and configured linked certificates both empty")
		return nil
	}

	// Unlink existing certificate if present
	if err := r.unlinkCertificate(ctx, data, diags); err != nil {
		return err
	}

	// Link new certificate if configured
	if configuredLinkedCertKeyname != "" {
		tflog.Debug(ctx, fmt.Sprintf("Linking certkey: %s", configuredLinkedCertKeyname))
		sslCertkey := ssl.Sslcertkey{
			Certkey:         sslcertkeyName,
			Linkcertkeyname: configuredLinkedCertKeyname,
		}
		if err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "link"); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error linking certificate: %v", err))
			return err
		}
	} else {
		tflog.Debug(ctx, "configured linked certkey is empty, nothing to do")
	}

	return nil
}

// Helper function to unlink certificate
func (r *SslCertKeyResource) unlinkCertificate(ctx context.Context, data *SslCertKeyResourceModel, diags *diag.Diagnostics) error {
	sslcertkeyName := data.Certkey.ValueString()
	if sslcertkeyName == "" {
		sslcertkeyName = data.Id.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Sslcertkey.Type(), sslcertkeyName)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("Error finding sslcertkey %s", sslcertkeyName))
		data.Id = types.StringNull()
		return err
	}

	actualLinkedCertKeyname := getResponseData["linkcertkeyname"]

	if actualLinkedCertKeyname != nil {
		tflog.Debug(ctx, fmt.Sprintf("Unlinking certkey: %s", actualLinkedCertKeyname))

		sslCertkey := ssl.Sslcertkey{
			Certkey: sslcertkeyName,
		}
		if err := r.client.ActOnResource(service.Sslcertkey.Type(), &sslCertkey, "unlink"); err != nil {
			tflog.Error(ctx, fmt.Sprintf("Error unlinking certificate: %v", err))
			return err
		}
	} else {
		tflog.Debug(ctx, "actual linked certkey is nil, nothing to do")
	}

	return nil
}
