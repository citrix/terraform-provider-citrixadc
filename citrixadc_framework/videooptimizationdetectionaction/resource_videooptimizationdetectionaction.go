package videooptimizationdetectionaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/videooptimization"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VideooptimizationdetectionactionResource{}
var _ resource.ResourceWithConfigure = (*VideooptimizationdetectionactionResource)(nil)
var _ resource.ResourceWithImportState = (*VideooptimizationdetectionactionResource)(nil)

func NewVideooptimizationdetectionactionResource() resource.Resource {
	return &VideooptimizationdetectionactionResource{}
}

// VideooptimizationdetectionactionResource defines the resource implementation.
type VideooptimizationdetectionactionResource struct {
	client *service.NitroClient
}

func (r *VideooptimizationdetectionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VideooptimizationdetectionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_videooptimizationdetectionaction"
}

func (r *VideooptimizationdetectionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VideooptimizationdetectionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating videooptimizationdetectionaction resource")

	videooptimizationdetectionaction := videooptimizationdetectionactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Videooptimizationdetectionaction.Type(), name_value, &videooptimizationdetectionaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create videooptimizationdetectionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created videooptimizationdetectionaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading videooptimizationdetectionaction resource")

	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource no longer exists on the ADC - remove it from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VideooptimizationdetectionactionResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (holds the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating videooptimizationdetectionaction resource")

	// Rename support: NITRO exposes an ?action=rename verb and a `newname` attribute.
	// A newname change drives an in-place rename, NOT a destroy/recreate. Every other
	// mutable change (type, comment) is handled by UpdateResource below; `name` is
	// RequiresReplace and never reaches Update.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID (== name at
		// create, == the prior newname after a rename), NOT state.Name which stays
		// pinned to the originally configured value.
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming videooptimizationdetectionaction from %q to %q", oldName, newName))

		renamePayload := videooptimization.Videooptimizationdetectionaction{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Videooptimizationdetectionaction.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename videooptimizationdetectionaction, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read below
		// (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Detect changes in the NITRO-updatable attributes (type, comment).
	hasChange := false
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, "type has changed for videooptimizationdetectionaction")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for videooptimizationdetectionaction")
		hasChange = true
	}

	if hasChange {
		videooptimizationdetectionaction := videooptimizationdetectionactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// The update must address the CURRENT LIVE name (== data.Id after any rename).
		liveName := data.Id.ValueString()
		videooptimizationdetectionaction.Name = liveName
		_, err := r.client.UpdateResource(service.Videooptimizationdetectionaction.Type(), liveName, &videooptimizationdetectionaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update videooptimizationdetectionaction, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated videooptimizationdetectionaction resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for videooptimizationdetectionaction resource, skipping update")
	}

	// Read the current state back. The live object may now be named newName, so we
	// must NOT let GET clobber the user-facing name/newname attributes. Capture the
	// plan values and restore them after the read to avoid an inconsistent-result /
	// perpetual diff.
	planName := data.Name
	planNewname := data.Newname
	r.readVideooptimizationdetectionactionFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VideooptimizationdetectionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VideooptimizationdetectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting videooptimizationdetectionaction resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Videooptimizationdetectionaction.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete videooptimizationdetectionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted videooptimizationdetectionaction resource")
}

// Helper function to read videooptimizationdetectionaction data from API
func (r *VideooptimizationdetectionactionResource) readVideooptimizationdetectionactionFromApi(ctx context.Context, data *VideooptimizationdetectionactionResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Videooptimizationdetectionaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read videooptimizationdetectionaction, got error: %s", err))
		return
	}

	videooptimizationdetectionactionSetAttrFromGet(ctx, data, getResponseData)
}
