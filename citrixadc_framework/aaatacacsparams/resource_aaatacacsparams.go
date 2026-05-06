package aaatacacsparams

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AaatacacsparamsResource{}
var _ resource.ResourceWithConfigure = (*AaatacacsparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AaatacacsparamsResource)(nil)

func NewAaatacacsparamsResource() resource.Resource {
	return &AaatacacsparamsResource{}
}

// AaatacacsparamsResource defines the resource implementation.
type AaatacacsparamsResource struct {
	client *service.NitroClient
}

func (r *AaatacacsparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaatacacsparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaatacacsparams"
}

func (r *AaatacacsparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaatacacsparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config AaatacacsparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaatacacsparams resource")
	// Get payload from plan (regular attributes)
	aaatacacsparams := aaatacacsparamsGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	aaatacacsparamsGetThePayloadFromtheConfig(ctx, &config, &aaatacacsparams)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaatacacsparams, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaatacacsparams resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("aaatacacsparams-config")

	// Read the updated state back
	r.readAaatacacsparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaatacacsparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaatacacsparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaatacacsparams resource")

	r.readAaatacacsparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaatacacsparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaatacacsparamsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating aaatacacsparams resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Accounting.Equal(state.Accounting) {
		tflog.Debug(ctx, fmt.Sprintf("accounting has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Auditfailedcmds.Equal(state.Auditfailedcmds) {
		tflog.Debug(ctx, fmt.Sprintf("auditfailedcmds has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Authorization.Equal(state.Authorization) {
		tflog.Debug(ctx, fmt.Sprintf("authorization has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Authtimeout.Equal(state.Authtimeout) {
		tflog.Debug(ctx, fmt.Sprintf("authtimeout has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		tflog.Debug(ctx, fmt.Sprintf("defaultauthenticationgroup has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Groupattrname.Equal(state.Groupattrname) {
		tflog.Debug(ctx, fmt.Sprintf("groupattrname has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, fmt.Sprintf("serverip has changed for aaatacacsparams"))
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, fmt.Sprintf("serverport has changed for aaatacacsparams"))
		hasChange = true
	}
	// Check secret attribute tacacssecret or its version tracker
	if !data.Tacacssecret.Equal(state.Tacacssecret) {
		tflog.Debug(ctx, fmt.Sprintf("tacacssecret has changed for aaatacacsparams"))
		hasChange = true
	} else if !data.TacacssecretWoVersion.Equal(state.TacacssecretWoVersion) {
		tflog.Debug(ctx, fmt.Sprintf("tacacssecret_wo_version has changed for aaatacacsparams"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Get payload from plan (regular attributes)
		aaatacacsparams := aaatacacsparamsGetThePayloadFromthePlan(ctx, &data)
		// Add write-only attributes from config to the payload
		aaatacacsparamsGetThePayloadFromtheConfig(ctx, &config, &aaatacacsparams)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaatacacsparams.Type(), &aaatacacsparams)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaatacacsparams, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaatacacsparams resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaatacacsparams resource, skipping update")
	}

	// Read the updated state back
	r.readAaatacacsparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaatacacsparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaatacacsparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaatacacsparams resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed aaatacacsparams from Terraform state")
}

// Helper function to read aaatacacsparams data from API
func (r *AaatacacsparamsResource) readAaatacacsparamsFromApi(ctx context.Context, data *AaatacacsparamsResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Aaatacacsparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaatacacsparams, got error: %s", err))
		return
	}

	aaatacacsparamsSetAttrFromGet(ctx, data, getResponseData)

}
