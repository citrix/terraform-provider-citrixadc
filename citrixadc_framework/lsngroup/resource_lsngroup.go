package lsngroup

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
var _ resource.Resource = &LsngroupResource{}
var _ resource.ResourceWithConfigure = (*LsngroupResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupResource)(nil)

func NewLsngroupResource() resource.Resource {
	return &LsngroupResource{}
}

// LsngroupResource defines the resource implementation.
type LsngroupResource struct {
	client *service.NitroClient
}

func (r *LsngroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup"
}

func (r *LsngroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup resource")

	// Build the create payload from the plan
	lsngroup := lsngroupGetThePayloadFromthePlan(ctx, &data)

	// Named resource - create using AddResource (POST)
	groupname := data.Groupname.ValueString()
	_, err := r.client.AddResource(service.Lsngroup.Type(), groupname, &lsngroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup, got error: %s", err))
		return
	}

	// ID is the resource name (single unique attribute: groupname)
	data.Id = types.StringValue(groupname)

	tflog.Trace(ctx, "Created lsngroup resource")

	// Read the updated state back
	if !r.readLsngroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsngroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup resource")

	found := r.readLsngroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsngroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state LsngroupResourceModel

	// Read prior state (to preserve ID and detect changes)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsngroup resource")

	// Only updateable (non-ForceNew) attributes are checked here. ForceNew
	// attributes (clientname, allocpolicy, ip6profile, nattype) trigger a
	// resource replacement and never reach Update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Ftp.Equal(state.Ftp) {
		if config.Ftp.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ftp")
		} else {
			hasChange = true
		}
	}
	if !data.Ftpcm.Equal(state.Ftpcm) {
		if config.Ftpcm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ftpcm")
		} else {
			hasChange = true
		}
	}
	if !data.Logging.Equal(state.Logging) {
		hasChange = true
	}
	if !data.Portblocksize.Equal(state.Portblocksize) {
		hasChange = true
	}
	if !data.Pptp.Equal(state.Pptp) {
		if config.Pptp.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "pptp")
		} else {
			hasChange = true
		}
	}
	if !data.Rtspalg.Equal(state.Rtspalg) {
		if config.Rtspalg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "rtspalg")
		} else {
			hasChange = true
		}
	}
	if !data.Sessionlogging.Equal(state.Sessionlogging) {
		hasChange = true
	}
	if !data.Sessionsync.Equal(state.Sessionsync) {
		if config.Sessionsync.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sessionsync")
		} else {
			hasChange = true
		}
	}
	if !data.Sipalg.Equal(state.Sipalg) {
		if config.Sipalg.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sipalg")
		} else {
			hasChange = true
		}
	}
	if !data.Snmptraplimit.Equal(state.Snmptraplimit) {
		if config.Snmptraplimit.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "snmptraplimit")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		lsngroup := lsngroupGetTheUpdatablePayloadFromthePlan(ctx, &data)
		// lsngroup update is a PUT on the unnamed endpoint (groupname is carried in the body)
		err := r.client.UpdateUnnamedResource(service.Lsngroup.Type(), &lsngroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lsngroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsngroup resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their NITRO defaults. The unset key for lsngroup is groupname.
	unsetIdPayload := map[string]interface{}{
		"groupname": data.Groupname.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Lsngroup.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lsngroup attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readLsngroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsngroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup resource")

	// Named resource - delete using DeleteResource
	err := r.client.DeleteResource(service.Lsngroup.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsngroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsngroup resource")
}

// Helper function to read lsngroup data from API. Returns false when the
// resource no longer exists on the appliance.
func (r *LsngroupResource) readLsngroupFromApi(ctx context.Context, data *LsngroupResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute: the ID is the plain groupname value.
	groupname := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsngroup.Type(), groupname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup, got error: %s", err))
		return false
	}

	lsngroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
