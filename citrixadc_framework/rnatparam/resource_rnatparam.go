package rnatparam

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
var _ resource.Resource = &RnatparamResource{}
var _ resource.ResourceWithConfigure = (*RnatparamResource)(nil)
var _ resource.ResourceWithImportState = (*RnatparamResource)(nil)

func NewRnatparamResource() resource.Resource {
	return &RnatparamResource{}
}

// RnatparamResource defines the resource implementation.
type RnatparamResource struct {
	client *service.NitroClient
}

func (r *RnatparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RnatparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rnatparam"
}

func (r *RnatparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RnatparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RnatparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating rnatparam resource")
	rnatparam := rnatparamGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Rnatparam.Type(), &rnatparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create rnatparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created rnatparam resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("rnatparam-config")

	// Read the updated state back
	r.readRnatparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RnatparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading rnatparam resource")

	r.readRnatparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RnatparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating rnatparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Srcippersistency.Equal(state.Srcippersistency) {
		tflog.Debug(ctx, fmt.Sprintf("srcippersistency has changed for rnatparam"))
		hasChange = true
	}
	if !data.Tcpproxy.Equal(state.Tcpproxy) {
		tflog.Debug(ctx, fmt.Sprintf("tcpproxy has changed for rnatparam"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		rnatparam := rnatparamGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Rnatparam.Type(), &rnatparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update rnatparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated rnatparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for rnatparam resource, skipping update")
	}

	// Read the updated state back
	r.readRnatparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RnatparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RnatparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting rnatparam resource")
	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed rnatparam from Terraform state")
}

// Helper function to read rnatparam data from API
func (r *RnatparamResource) readRnatparamFromApi(ctx context.Context, data *RnatparamResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID
	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Rnatparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read rnatparam, got error: %s", err))
		return
	}

	rnatparamSetAttrFromGet(ctx, data, getResponseData)

}
