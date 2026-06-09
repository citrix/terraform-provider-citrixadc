package clusternodegroup

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
var _ resource.Resource = &ClusternodegroupResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupResource)(nil)

func NewClusternodegroupResource() resource.Resource {
	return &ClusternodegroupResource{}
}

// ClusternodegroupResource defines the resource implementation.
type ClusternodegroupResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup"
}

func (r *ClusternodegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup resource")
	clusternodegroup := clusternodegroupGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Clusternodegroup.Type(), name_value, &clusternodegroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created clusternodegroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readClusternodegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup resource")

	r.readClusternodegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusternodegroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating clusternodegroup resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, fmt.Sprintf("priority has changed for clusternodegroup"))
		hasChange = true
	}
	if !data.State.Equal(state.State) {
		tflog.Debug(ctx, fmt.Sprintf("state has changed for clusternodegroup"))
		hasChange = true
	}
	if !data.Strict.Equal(state.Strict) {
		tflog.Debug(ctx, fmt.Sprintf("strict has changed for clusternodegroup"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		// Use the update-specific payload builder which omits the create-only `sticky`
		// attribute (NITRO set/PUT does not accept it).
		clusternodegroup := clusternodegroupGetThePayloadForUpdate(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Clusternodegroup.Type(), name_value, &clusternodegroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated clusternodegroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for clusternodegroup resource, skipping update")
	}

	// Read the updated state back
	r.readClusternodegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Clusternodegroup.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete clusternodegroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted clusternodegroup resource")
}

// Helper function to read clusternodegroup data from API
func (r *ClusternodegroupResource) readClusternodegroupFromApi(ctx context.Context, data *ClusternodegroupResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Clusternodegroup.Type(), name_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup, got error: %s", err))
		return
	}

	clusternodegroupSetAttrFromGet(ctx, data, getResponseData)

}
