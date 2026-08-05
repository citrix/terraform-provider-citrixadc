package lbaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbactionResource{}
var _ resource.ResourceWithConfigure = (*LbactionResource)(nil)
var _ resource.ResourceWithImportState = (*LbactionResource)(nil)

func NewLbactionResource() resource.Resource {
	return &LbactionResource{}
}

// LbactionResource defines the resource implementation.
type LbactionResource struct {
	client *service.NitroClient
}

func (r *LbactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbaction"
}

func (r *LbactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbaction resource")

	lbaction := lbactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO add is POST)
	lbactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lbaction.Type(), lbactionName, &lbaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbaction resource")

	// Set ID for the resource before reading state (single unique attribute - plain value)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readLbactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbaction resource")

	found := r.readLbactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbaction resource")

	// Rename support: NITRO exposes a `rename` action for lbaction. name and type are
	// ForceNew (RequiresReplace), so the only key change that reaches Update is
	// `newname`. On a newname change, POST {name, newname} to ?action=rename, then
	// point the resource ID at the new name so subsequent reads/updates/deletes address
	// the live object.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming lbaction from %q to %q", oldName, newName))

		renamePayload := lb.Lbaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Lbaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename lbaction, got error: %s", err))
			return
		}
		// The live object is now named newName.
		data.Id = types.StringValue(newName)
	}

	// Regular update (comment/value) via NITRO PUT.
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for lbaction")
		hasChange = true
	}
	if !data.Value.Equal(state.Value) {
		tflog.Debug(ctx, "value has changed for lbaction")
		hasChange = true
	}

	if hasChange {
		lbaction := lbactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Address the current live name (== newname after a rename, else name).
		liveName := data.Id.ValueString()
		lbaction.Name = liveName
		_, err := r.client.UpdateResource(service.Lbaction.Type(), liveName, &lbaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lbaction resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for lbaction resource")
	}

	// Read the current state back, preserving the user-facing key/rename inputs so a
	// rename does not clobber the configured name and trigger an inconsistent result.
	planName := data.Name
	planNewname := data.Newname
	r.readLbactionFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id.
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lbaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbaction resource")
}

// Helper function to read lbaction data from API
func (r *LbactionResource) readLbactionFromApi(ctx context.Context, data *LbactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (the live name)
	lbactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lbaction.Type(), lbactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbaction, got error: %s", err))
		return false
	}

	lbactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
