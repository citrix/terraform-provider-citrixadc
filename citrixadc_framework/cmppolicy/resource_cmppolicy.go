package cmppolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CmppolicyResource{}
var _ resource.ResourceWithConfigure = (*CmppolicyResource)(nil)
var _ resource.ResourceWithImportState = (*CmppolicyResource)(nil)

func NewCmppolicyResource() resource.Resource {
	return &CmppolicyResource{}
}

// CmppolicyResource defines the resource implementation.
type CmppolicyResource struct {
	client *service.NitroClient
}

func (r *CmppolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmppolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmppolicy"
}

func (r *CmppolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmppolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmppolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmppolicy resource")

	cmppolicy := cmppolicyGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	cmppolicyName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Cmppolicy.Type(), cmppolicyName, &cmppolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created cmppolicy resource")

	// Set ID for the resource before reading state (single unique attr -> plain name)
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	r.readCmppolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmppolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmppolicy resource")

	r.readCmppolicyFromApi(ctx, &data, &resp.Diagnostics)

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

func (r *CmppolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CmppolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the CURRENT LIVE name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating cmppolicy resource")

	// Rename branch: cmppolicy exposes a `rename` action (NITRO ?action=rename). The
	// `name` primary key is RequiresReplace, so a name change recreates the resource
	// and never reaches Update; the ONLY key-changing path that lands here is a
	// `newname` change, which must drive an in-place rename (mirrors SDK v2 rename
	// convention in citrixadc/resource_citrixadc_appfwpolicy.go).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		// The rename SOURCE is the CURRENT LIVE name, tracked by the ID - NOT
		// state.Name (which stays pinned to the originally configured value and would
		// point at a no-longer-live name on a second rename).
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming cmppolicy from %q to %q", oldName, newName))

		renamePayload := cmp.Cmppolicy{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Cmppolicy.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename cmppolicy, got error: %s", err))
			return
		}

		// The live object is now named newName. Point the ID at it so the update and
		// read below (and all future reads) address the renamed resource.
		data.Id = types.StringValue(newName)
	}

	// In-place update branch: rule and resaction are updateable (NITRO PUT). Only send
	// the update when one of them actually changed (matches SDK v2 hasChange logic).
	hasChange := false
	if !data.Resaction.Equal(state.Resaction) {
		tflog.Debug(ctx, "resaction has changed for cmppolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for cmppolicy")
		hasChange = true
	}

	if hasChange {
		cmppolicy := cmppolicyGetThePayloadFromthePlan(ctx, &data)
		// Target the live name (post-rename) for the update PUT body.
		liveName := data.Id.ValueString()
		cmppolicy.Name = liveName
		_, err := r.client.UpdateResource(service.Cmppolicy.Type(), liveName, &cmppolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmppolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated cmppolicy resource")
	} else {
		tflog.Debug(ctx, "No updateable changes detected for cmppolicy resource, skipping update")
	}

	// Read the current state back. The resource may now be physically named newName,
	// so capture the plan's name + newname and restore them after the read to avoid an
	// inconsistent-result / perpetual diff (SetAttrFromGet preserves name, but be
	// explicit for the rename case).
	planName := data.Name
	planNewname := data.Newname
	r.readCmppolicyFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmppolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmppolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmppolicy resource")
	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE name
	// (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, dangling the object).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Cmppolicy.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cmppolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted cmppolicy resource")
}

// Helper function to read cmppolicy data from API
func (r *CmppolicyResource) readCmppolicyFromApi(ctx context.Context, data *CmppolicyResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain (live) name
	cmppolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Cmppolicy.Type(), cmppolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmppolicy, got error: %s", err))
		return
	}

	cmppolicySetAttrFromGet(ctx, data, getResponseData)
}
