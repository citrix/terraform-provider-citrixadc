package nsvariable

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
var _ resource.Resource = &NsvariableResource{}
var _ resource.ResourceWithConfigure = (*NsvariableResource)(nil)
var _ resource.ResourceWithImportState = (*NsvariableResource)(nil)

func NewNsvariableResource() resource.Resource {
	return &NsvariableResource{}
}

// NsvariableResource defines the resource implementation.
type NsvariableResource struct {
	client *service.NitroClient
}

func (r *NsvariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsvariableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsvariable"
}

func (r *NsvariableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsvariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsvariableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsvariable resource")

	nsvariable := nsvariableGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	nsvariableName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Nsvariable.Type(), nsvariableName, &nsvariable)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsvariable, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsvariable resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", nsvariableName))

	// Read the updated state back
	if !r.readNsvariableFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsvariable not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsvariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsvariableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsvariable resource")

	found := r.readNsvariableFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsvariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsvariableResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsvariable resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for nsvariable")
		hasChange = true
	}
	if !data.Expires.Equal(state.Expires) {
		tflog.Debug(ctx, "expires has changed for nsvariable")
		hasChange = true
	}
	if !data.Iffull.Equal(state.Iffull) {
		tflog.Debug(ctx, "iffull has changed for nsvariable")
		hasChange = true
	}
	if !data.Ifnovalue.Equal(state.Ifnovalue) {
		tflog.Debug(ctx, "ifnovalue has changed for nsvariable")
		hasChange = true
	}
	if !data.Ifvaluetoobig.Equal(state.Ifvaluetoobig) {
		tflog.Debug(ctx, "ifvaluetoobig has changed for nsvariable")
		hasChange = true
	}
	if !data.Init.Equal(state.Init) {
		tflog.Debug(ctx, "init has changed for nsvariable")
		hasChange = true
	}

	if hasChange {
		nsvariable := nsvariableGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		nsvariableName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Nsvariable.Type(), nsvariableName, &nsvariable)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsvariable, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nsvariable resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsvariable resource, skipping update")
	}

	// Read the updated state back
	if !r.readNsvariableFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsvariable not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsvariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsvariableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsvariable resource")

	// Named resource - delete using DeleteResource
	nsvariableName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nsvariable.Type(), nsvariableName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsvariable, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsvariable resource")
}

// Helper function to read nsvariable data from API
func (r *NsvariableResource) readNsvariableFromApi(ctx context.Context, data *NsvariableResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	nsvariableName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsvariable.Type(), nsvariableName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsvariable, got error: %s", err))
		return false
	}

	nsvariableSetAttrFromGet(ctx, data, getResponseData)

	return true
}
