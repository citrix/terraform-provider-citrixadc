package auditnslogaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AuditnslogactionResource{}
var _ resource.ResourceWithConfigure = (*AuditnslogactionResource)(nil)
var _ resource.ResourceWithImportState = (*AuditnslogactionResource)(nil)

func NewAuditnslogactionResource() resource.Resource {
	return &AuditnslogactionResource{}
}

// AuditnslogactionResource defines the resource implementation.
type AuditnslogactionResource struct {
	client *service.NitroClient
}

func (r *AuditnslogactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditnslogactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditnslogaction"
}

func (r *AuditnslogactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditnslogactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditnslogactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditnslogaction resource")
	// Get payload from plan
	auditnslogaction := auditnslogactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	auditnslogactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Auditnslogaction.Type(), auditnslogactionName, &auditnslogaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditnslogaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created auditnslogaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(auditnslogactionName)

	// Read the updated state back
	if !r.readAuditnslogactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditnslogaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditnslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditnslogaction resource")

	found := r.readAuditnslogactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuditnslogactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating auditnslogaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acl.Equal(state.Acl) {
		if config.Acl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "acl")
		} else {
			hasChange = true
		}
	}
	if !data.Alg.Equal(state.Alg) {
		if config.Alg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "alg")
		} else {
			hasChange = true
		}
	}
	if !data.Appflowexport.Equal(state.Appflowexport) {
		if config.Appflowexport.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "appflowexport")
		} else {
			hasChange = true
		}
	}
	if !data.Contentinspectionlog.Equal(state.Contentinspectionlog) {
		if config.Contentinspectionlog.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "contentinspectionlog")
		} else {
			hasChange = true
		}
	}
	if !data.Dateformat.Equal(state.Dateformat) {
		if config.Dateformat.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "dateformat")
		} else {
			hasChange = true
		}
	}
	if !data.Domainresolvenow.Equal(state.Domainresolvenow) {
		hasChange = true
	}
	if !data.Domainresolveretry.Equal(state.Domainresolveretry) {
		hasChange = true
	}
	if !data.Logfacility.Equal(state.Logfacility) {
		if config.Logfacility.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logfacility")
		} else {
			hasChange = true
		}
	}
	if !data.Loglevel.Equal(state.Loglevel) {
		hasChange = true
	}
	if !data.Lsn.Equal(state.Lsn) {
		if config.Lsn.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "lsn")
		} else {
			hasChange = true
		}
	}
	if !data.Protocolviolations.Equal(state.Protocolviolations) {
		if config.Protocolviolations.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "protocolviolations")
		} else {
			hasChange = true
		}
	}
	if !data.Serverdomainname.Equal(state.Serverdomainname) {
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		if config.Serverport.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverport")
		} else {
			hasChange = true
		}
	}
	if !data.Sslinterception.Equal(state.Sslinterception) {
		if config.Sslinterception.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sslinterception")
		} else {
			hasChange = true
		}
	}
	if !data.Subscriberlog.Equal(state.Subscriberlog) {
		if config.Subscriberlog.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "subscriberlog")
		} else {
			hasChange = true
		}
	}
	if !data.Tcp.Equal(state.Tcp) {
		if config.Tcp.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "tcp")
		} else {
			hasChange = true
		}
	}
	if !data.Timezone.Equal(state.Timezone) {
		if config.Timezone.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timezone")
		} else {
			hasChange = true
		}
	}
	if !data.Urlfiltering.Equal(state.Urlfiltering) {
		hasChange = true
	}
	if !data.Userdefinedauditlog.Equal(state.Userdefinedauditlog) {
		if config.Userdefinedauditlog.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "userdefinedauditlog")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		auditnslogaction := auditnslogactionGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// NITRO update for auditnslogaction is a PUT to /config/auditnslogaction (name is in the body)
		err := r.client.UpdateUnnamedResource(service.Auditnslogaction.Type(), &auditnslogaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditnslogaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated auditnslogaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for auditnslogaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Auditnslogaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset auditnslogaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAuditnslogactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "auditnslogaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditnslogactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditnslogactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditnslogaction resource")
	// Named resource - delete using DeleteResource
	auditnslogactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Auditnslogaction.Type(), auditnslogactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete auditnslogaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted auditnslogaction resource")
}

// Helper function to read auditnslogaction data from API
func (r *AuditnslogactionResource) readAuditnslogactionFromApi(ctx context.Context, data *AuditnslogactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	auditnslogactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Auditnslogaction.Type(), auditnslogactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditnslogaction, got error: %s", err))
		return false
	}

	auditnslogactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
