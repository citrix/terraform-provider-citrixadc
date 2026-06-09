package lbwlm

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
var _ resource.Resource = &LbwlmResource{}
var _ resource.ResourceWithConfigure = (*LbwlmResource)(nil)
var _ resource.ResourceWithImportState = (*LbwlmResource)(nil)

func NewLbwlmResource() resource.Resource {
	return &LbwlmResource{}
}

// LbwlmResource defines the resource implementation.
type LbwlmResource struct {
	client *service.NitroClient
}

func (r *LbwlmResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbwlmResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbwlm"
}

func (r *LbwlmResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbwlmResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbwlmResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbwlm resource")
	lbwlm := lbwlmGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	wlmname_value := data.Wlmname.ValueString()
	_, err := r.client.AddResource(service.Lbwlm.Type(), wlmname_value, &lbwlm)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbwlm, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbwlm resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Wlmname.ValueString()))

	// Read the updated state back
	r.readLbwlmFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbwlmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbwlm resource")

	r.readLbwlmFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbwlmResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbwlm resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Katimeout.Equal(state.Katimeout) {
		tflog.Debug(ctx, fmt.Sprintf("katimeout has changed for lbwlm"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model.
		// Update (PUT) accepts ONLY wlmname + katimeout; use the dedicated update payload builder (Pattern 9).
		lbwlm := lbwlmGetTheUpdatePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		wlmname_value := data.Wlmname.ValueString()
		_, err := r.client.UpdateResource(service.Lbwlm.Type(), wlmname_value, &lbwlm)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbwlm, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lbwlm resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbwlm resource, skipping update")
	}

	// Read the updated state back
	r.readLbwlmFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbwlmResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbwlmResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbwlm resource")
	// Named resource - delete using DeleteResource
	wlmname_value := data.Wlmname.ValueString()
	err := r.client.DeleteResource(service.Lbwlm.Type(), wlmname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbwlm, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbwlm resource")
}

// Helper function to read lbwlm data from API
func (r *LbwlmResource) readLbwlmFromApi(ctx context.Context, data *LbwlmResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	wlmname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Lbwlm.Type(), wlmname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbwlm, got error: %s", err))
		return
	}

	lbwlmSetAttrFromGet(ctx, data, getResponseData)

}
