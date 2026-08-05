package lbpolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbpolicyResource{}
var _ resource.ResourceWithConfigure = (*LbpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*LbpolicyResource)(nil)

func NewLbpolicyResource() resource.Resource {
	return &LbpolicyResource{}
}

// LbpolicyResource defines the resource implementation.
type LbpolicyResource struct {
	client *service.NitroClient
}

func (r *LbpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbpolicy"
}

func (r *LbpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbpolicy resource")

	lbpolicy := lbpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lbpolicy.Type(), name_value, &lbpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbpolicy resource")

	// Set ID for the resource before reading state (single unique attr -> plain name)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readLbpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbpolicy resource")

	r.readLbpolicyFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not found on the ADC - remove from state
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbpolicyResourceModel

	// Read Terraform prior state to preserve the live ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the live ID from prior state (tracks the current object name).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbpolicy resource")

	// Rename support: on a newname change, POST {name, newname} to ?action=rename,
	// then point the resource ID at the new name so subsequent reads/updates address
	// the live object. The rename SOURCE is the CURRENT LIVE name (state.Id), NOT
	// state.Name (which stays pinned to the originally configured value and would be
	// stale on a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming lbpolicy from %q to %q", oldName, newName))

		renamePayload := lb.Lbpolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Lbpolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename lbpolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Track it via the ID.
		data.Id = types.StringValue(newName)
	}

	// Detect changes in NITRO-updatable attributes (name is RequiresReplace and never
	// reaches Update; newname is handled above).
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for lbpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for lbpolicy")
		hasChange = true
	}
	if !data.Logaction.Equal(state.Logaction) {
		tflog.Debug(ctx, "logaction has changed for lbpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for lbpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for lbpolicy")
		hasChange = true
	}

	if hasChange {
		lbpolicy := lbpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Key/update by the CURRENT LIVE name (data.Id), which reflects any rename
		// that just happened above; also pin the payload name to the live name.
		liveName := data.Id.ValueString()
		lbpolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Lbpolicy.Type(), liveName, &lbpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lbpolicy resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for lbpolicy resource")
	}

	// Read the updated state back. Preserve the user-facing name and newname across
	// the read-back so a rename does not clobber the configured values.
	planName := data.Name
	planNewname := data.Newname
	r.readLbpolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbpolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id,
	// NOT data.Name (which is stale after a rename).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lbpolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbpolicy resource")
}

// readLbpolicyFromApi reads the lbpolicy from the ADC by its current live name
// (data.Id). On a not-found it nulls the ID as a sentinel so Read can remove the
// resource from state.
func (r *LbpolicyResource) readLbpolicyFromApi(ctx context.Context, data *LbpolicyResourceModel, diags *diag.Diagnostics) {
	// Case 2: Find with single ID attribute - ID is the plain (live) name.
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lbpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbpolicy, got error: %s", err))
		return
	}

	lbpolicySetAttrFromGet(ctx, data, getResponseData)
}
