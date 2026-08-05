package tmtrafficaction

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
var _ resource.Resource = &TmtrafficactionResource{}
var _ resource.ResourceWithConfigure = (*TmtrafficactionResource)(nil)
var _ resource.ResourceWithImportState = (*TmtrafficactionResource)(nil)

func NewTmtrafficactionResource() resource.Resource {
	return &TmtrafficactionResource{}
}

// TmtrafficactionResource defines the resource implementation.
type TmtrafficactionResource struct {
	client *service.NitroClient
}

func (r *TmtrafficactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmtrafficactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmtrafficaction"
}

func (r *TmtrafficactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmtrafficactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmtrafficaction resource")

	tmtrafficaction := tmtrafficactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	tmtrafficactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Tmtrafficaction.Type(), tmtrafficactionName, &tmtrafficaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmtrafficaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tmtrafficaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(tmtrafficactionName)

	// Read the updated state back
	if !r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmtrafficaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmtrafficaction resource")

	found := r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TmtrafficactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TmtrafficactionResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tmtrafficaction resource")

	tmtrafficactionName := data.Name.ValueString()

	// Detect changes in updateable attributes (name is ForceNew, not updateable)
	hasChange := false
	if !data.Apptimeout.Equal(state.Apptimeout) {
		hasChange = true
	}
	if !data.Forcedtimeout.Equal(state.Forcedtimeout) {
		hasChange = true
	}
	if !data.Forcedtimeoutval.Equal(state.Forcedtimeoutval) {
		hasChange = true
	}
	if !data.Formssoaction.Equal(state.Formssoaction) {
		hasChange = true
	}
	if !data.Initiatelogout.Equal(state.Initiatelogout) {
		hasChange = true
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		hasChange = true
	}
	if !data.Passwdexpression.Equal(state.Passwdexpression) {
		hasChange = true
	}
	if !data.Persistentcookie.Equal(state.Persistentcookie) {
		hasChange = true
	}
	if !data.Samlssoprofile.Equal(state.Samlssoprofile) {
		hasChange = true
	}
	if !data.Sso.Equal(state.Sso) {
		hasChange = true
	}
	if !data.Userexpression.Equal(state.Userexpression) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model (Name is included in the payload)
		tmtrafficaction := tmtrafficactionGetThePayloadFromthePlan(ctx, &data)

		// Match SDK v2 semantics: unnamed PUT with name in the body
		err := r.client.UpdateUnnamedResource(service.Tmtrafficaction.Type(), &tmtrafficaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmtrafficaction %s, got error: %s", tmtrafficactionName, err))
			return
		}

		tflog.Trace(ctx, "Updated tmtrafficaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tmtrafficaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readTmtrafficactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "tmtrafficaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmtrafficactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmtrafficactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmtrafficaction resource")

	// Named resource - delete using DeleteResource
	tmtrafficactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Tmtrafficaction.Type(), tmtrafficactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tmtrafficaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tmtrafficaction resource")
}

// Helper function to read tmtrafficaction data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *TmtrafficactionResource) readTmtrafficactionFromApi(ctx context.Context, data *TmtrafficactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	tmtrafficactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Tmtrafficaction.Type(), tmtrafficactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmtrafficaction, got error: %s", err))
		return false
	}

	tmtrafficactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
