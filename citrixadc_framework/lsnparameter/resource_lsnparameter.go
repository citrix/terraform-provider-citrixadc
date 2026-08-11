package lsnparameter

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
var _ resource.Resource = &LsnparameterResource{}
var _ resource.ResourceWithConfigure = (*LsnparameterResource)(nil)
var _ resource.ResourceWithImportState = (*LsnparameterResource)(nil)

func NewLsnparameterResource() resource.Resource {
	return &LsnparameterResource{}
}

// LsnparameterResource defines the resource implementation.
type LsnparameterResource struct {
	client *service.NitroClient
}

func (r *LsnparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnparameter"
}

func (r *LsnparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnparameter resource")

	// Singleton resource - use UpdateUnnamedResource (NITRO exposes only get/update)
	lsnparameter := lsnparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Lsnparameter.Type(), &lsnparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsnparameter-config")

	tflog.Trace(ctx, "Created lsnparameter resource")

	// Read the updated state back
	r.readLsnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnparameter resource")

	r.readLsnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LsnparameterResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Memlimit.Equal(state.Memlimit) {
		tflog.Debug(ctx, "memlimit has changed for lsnparameter")
		hasChange = true
	}
	if !data.Sessionsync.Equal(state.Sessionsync) {
		tflog.Debug(ctx, "sessionsync has changed for lsnparameter")
		if config.Sessionsync.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sessionsync")
		} else {
			hasChange = true
		}
	}
	if !data.Subscrsessionremoval.Equal(state.Subscrsessionremoval) {
		tflog.Debug(ctx, "subscrsessionremoval has changed for lsnparameter")
		if config.Subscrsessionremoval.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "subscrsessionremoval")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Singleton resource - use UpdateUnnamedResource
		lsnparameter := lsnparameterGetThePayloadFromtheConfig(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Lsnparameter.Type(), &lsnparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnparameter, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsnparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Lsnparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lsnparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readLsnparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnparameter resource")

	// For lsnparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsnparameter resource from state")
}

// Helper function to read lsnparameter data from API
func (r *LsnparameterResource) readLsnparameterFromApi(ctx context.Context, data *LsnparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsnparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnparameter, got error: %s", err))
		return
	}

	lsnparameterSetAttrFromGet(ctx, data, getResponseData)

}
