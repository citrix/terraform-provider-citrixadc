package vridparam

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
var _ resource.Resource = &VridparamResource{}
var _ resource.ResourceWithConfigure = (*VridparamResource)(nil)
var _ resource.ResourceWithImportState = (*VridparamResource)(nil)

func NewVridparamResource() resource.Resource {
	return &VridparamResource{}
}

// VridparamResource defines the resource implementation.
type VridparamResource struct {
	client *service.NitroClient
}

func (r *VridparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VridparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vridparam"
}

func (r *VridparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VridparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VridparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vridparam resource")
	vridparam := vridparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource (matches SDK v2 semantics).
	err := r.client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vridparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vridparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("vridparam-config")

	// Read the updated state back
	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VridparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vridparam resource")

	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VridparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vridparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Deadinterval.Equal(state.Deadinterval) {
		tflog.Debug(ctx, fmt.Sprintf("deadinterval has changed for vridparam"))
		if config.Deadinterval.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "deadinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Hellointerval.Equal(state.Hellointerval) {
		tflog.Debug(ctx, fmt.Sprintf("hellointerval has changed for vridparam"))
		if config.Hellointerval.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "hellointerval")
		} else {
			hasChange = true
		}
	}
	if !data.Sendtomaster.Equal(state.Sendtomaster) {
		tflog.Debug(ctx, fmt.Sprintf("sendtomaster has changed for vridparam"))
		if config.Sendtomaster.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sendtomaster")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		vridparam := vridparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource (matches SDK v2 semantics).
		err := r.client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vridparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vridparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vridparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Vridparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vridparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VridparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vridparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed vridparam from Terraform state")
}

// Helper function to read vridparam data from API
func (r *VridparamResource) readVridparamFromApi(ctx context.Context, data *VridparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Vridparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vridparam, got error: %s", err))
		return
	}

	vridparamSetAttrFromGet(ctx, data, getResponseData)

}
