package nslimitselector

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
var _ resource.Resource = &NslimitselectorResource{}
var _ resource.ResourceWithConfigure = (*NslimitselectorResource)(nil)
var _ resource.ResourceWithImportState = (*NslimitselectorResource)(nil)

func NewNslimitselectorResource() resource.Resource {
	return &NslimitselectorResource{}
}

// NslimitselectorResource defines the resource implementation.
type NslimitselectorResource struct {
	client *service.NitroClient
}

func (r *NslimitselectorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslimitselectorResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslimitselector"
}

func (r *NslimitselectorResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslimitselectorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslimitselectorResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslimitselector resource")
	nslimitselector := nslimitselectorGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	selectorname_value := data.Selectorname.ValueString()
	_, err := r.client.AddResource(service.Nslimitselector.Type(), selectorname_value, &nslimitselector)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslimitselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nslimitselector resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Selectorname.ValueString()))

	// Read the updated state back
	if !r.readNslimitselectorFromApi(ctx, &data, &resp.Diagnostics) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nslimitselector %s immediately after create", data.Id.ValueString()))
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitselectorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslimitselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslimitselector resource")

	if !r.readNslimitselectorFromApi(ctx, &data, &resp.Diagnostics) {
		// Resource no longer exists on the appliance - treat as drift.
		tflog.Debug(ctx, fmt.Sprintf("Removing nslimitselector %s from state (not found)", data.Id.ValueString()))
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitselectorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NslimitselectorResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nslimitselector resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, fmt.Sprintf("rule has changed for nslimitselector"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		nslimitselector := nslimitselectorGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		selectorname_value := data.Selectorname.ValueString()
		_, err := r.client.UpdateResource(service.Nslimitselector.Type(), selectorname_value, &nslimitselector)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslimitselector, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nslimitselector resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nslimitselector resource, skipping update")
	}

	// Read the updated state back
	if !r.readNslimitselectorFromApi(ctx, &data, &resp.Diagnostics) {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nslimitselector %s after update", data.Id.ValueString()))
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslimitselectorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslimitselectorResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslimitselector resource")
	// Named resource - delete using DeleteResource
	selectorname_value := data.Selectorname.ValueString()
	err := r.client.DeleteResource(service.Nslimitselector.Type(), selectorname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nslimitselector, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nslimitselector resource")
}

// Helper function to read nslimitselector data from API.
// Returns true if the resource was found, false if it does not exist (drift).
// NOTE: The NITRO appliance returns this object under the "streamselector" key
// (nslimitselector is an alias of streamselector). A typed GET against the
// "nslimitselector" type returns 200 with a body the client cannot map back
// (key/field mismatch: response uses key "streamselector" and field "name"),
// so we read against the canonical "streamselector" type.
func (r *NslimitselectorResource) readNslimitselectorFromApi(ctx context.Context, data *NslimitselectorResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	selectorname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Streamselector.Type(), selectorname_Name)
	if err != nil {
		// FindResource returns an error both for a genuine 404 and an empty body.
		// Treat as not-found so the caller can decide (drift vs. hard error).
		tflog.Debug(ctx, fmt.Sprintf("nslimitselector %s not found: %s", selectorname_Name, err))
		return false
	}

	nslimitselectorSetAttrFromGet(ctx, data, getResponseData)
	return true

}
