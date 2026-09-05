package responderpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ResponderpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*ResponderpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderpolicylabelResource)(nil)

func NewResponderpolicylabelResource() resource.Resource {
	return &ResponderpolicylabelResource{}
}

// ResponderpolicylabelResource defines the resource implementation.
type ResponderpolicylabelResource struct {
	client *service.NitroClient
}

func (r *ResponderpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderpolicylabel"
}

func (r *ResponderpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderpolicylabel resource")
	responderpolicylabel := responderpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (NITRO exposes `add`, HTTP POST).
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Responderpolicylabel.Type(), labelname_value, &responderpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created responderpolicylabel resource")

	// Set ID for the resource before reading state (single unique attr - plain value).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	// Read the updated state back
	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderpolicylabel resource")

	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ResponderpolicylabelResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Rename support: responderpolicylabel exposes NO set/update endpoint. The only
	// in-place mutation NITRO offers is the `rename` action. Every other schema
	// attribute (labelname, policylabeltype, comment) uses RequiresReplace, so
	// Terraform recreates the resource on any of those changes and never reaches here
	// for them. The ONLY change that lands in Update is `newname`.
	//
	// Mirrors the SDK v2 convention (see citrixadc/resource_citrixadc_appfwpolicy.go):
	// on a newname change, POST {labelname, newname} to ?action=rename, then point the
	// resource ID at the new name so subsequent reads address the live object.
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Labelname (which stays pinned to the originally configured value and
		// would point at the wrong name on a SECOND rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming responderpolicylabel from %q to %q", oldName, newName))

		renamePayload := responder.Responderpolicylabel{
			Labelname: oldName,
			Newname:   newName,
		}
		if err := r.client.ActOnResource(service.Responderpolicylabel.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename responderpolicylabel, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the read below
		// (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// Read the current state back. Capture the plan values and restore them after the
	// read so a rename's GET (which returns the new live name) does not clobber the
	// user-facing labelname/newname attributes and cause an inconsistent-result diff.
	planLabelname := data.Labelname
	planNewname := data.Newname
	r.readResponderpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
	data.Labelname = planLabelname
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderpolicylabel resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== labelname at create, == newname after a rename), so delete by data.Id, NOT
	// data.Labelname (stale/old after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Responderpolicylabel.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete responderpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted responderpolicylabel resource")
}

// Helper function to read responderpolicylabel data from API
func (r *ResponderpolicylabelResource) readResponderpolicylabelFromApi(ctx context.Context, data *ResponderpolicylabelResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	labelname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Responderpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderpolicylabel, got error: %s", err))
		return
	}

	responderpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
