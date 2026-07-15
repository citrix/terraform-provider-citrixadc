package nsextension

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
var _ resource.Resource = &NsextensionResource{}
var _ resource.ResourceWithConfigure = (*NsextensionResource)(nil)
var _ resource.ResourceWithImportState = (*NsextensionResource)(nil)

func NewNsextensionResource() resource.Resource {
	return &NsextensionResource{}
}

// NsextensionResource defines the resource implementation.
type NsextensionResource struct {
	client *service.NitroClient
}

func (r *NsextensionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsextensionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsextension"
}

func (r *NsextensionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsextensionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsextensionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsextension resource")
	nsextension := nsextensionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// nsextension is imported via the ?action=Import endpoint (uploads from src)
	err := r.client.ActOnResource(service.Nsextension.Type(), &nsextension, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsextension, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsextension resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readNsextensionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsextensionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsextension resource")

	r.readNsextensionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Object deleted out-of-band: remove from state so a later apply re-creates it.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsextensionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsextension resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, fmt.Sprintf("comment has changed for nsextension"))
		hasChange = true
	}
	if !data.Trace.Equal(state.Trace) {
		tflog.Debug(ctx, fmt.Sprintf("trace has changed for nsextension"))
		hasChange = true
	}
	if !data.Tracefunctions.Equal(state.Tracefunctions) {
		tflog.Debug(ctx, fmt.Sprintf("tracefunctions has changed for nsextension"))
		hasChange = true
	}
	if !data.Tracevariables.Equal(state.Tracevariables) {
		tflog.Debug(ctx, fmt.Sprintf("tracevariables has changed for nsextension"))
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model (update accepts a different field
		// set than Import: only name/trace/tracefunctions/tracevariables/comment)
		nsextension := nsextensionGetTheUpdatePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - update via PUT
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsextension.Type(), name_value, &nsextension)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsextension, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsextension resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsextension resource, skipping update")
	}

	// Read the updated state back
	r.readNsextensionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsextensionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsextensionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsextension resource")
	// Named resource - delete using DeleteResource by name (DELETE /nsextension/<name>)
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsextension.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsextension, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsextension resource")
}

// Helper function to read nsextension data from API
func (r *NsextensionResource) readNsextensionFromApi(ctx context.Context, data *NsextensionResourceModel, diags *diag.Diagnostics) {

	// ID is the plain name value (single key)
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nsextension.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsextension, got error: %s", err))
		return
	}

	nsextensionSetAttrFromGet(ctx, data, getResponseData)
}
