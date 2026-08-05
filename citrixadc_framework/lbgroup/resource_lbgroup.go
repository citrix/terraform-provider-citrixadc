package lbgroup

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
var _ resource.Resource = &LbgroupResource{}
var _ resource.ResourceWithConfigure = (*LbgroupResource)(nil)
var _ resource.ResourceWithImportState = (*LbgroupResource)(nil)

func NewLbgroupResource() resource.Resource {
	return &LbgroupResource{}
}

// LbgroupResource defines the resource implementation.
type LbgroupResource struct {
	client *service.NitroClient
}

func (r *LbgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbgroup"
}

func (r *LbgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbgroup resource")

	lbgroup := lbgroupGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	lbgroupName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Lbgroup.Type(), lbgroupName, &lbgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbgroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(lbgroupName)

	// Read the updated state back
	if !r.readLbgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbgroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbgroup resource")

	found := r.readLbgroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LbgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LbgroupResourceModel

	// Read Terraform prior state to preserve the live ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the live ID from prior state (tracks the current object name).
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbgroup resource")

	// Handle in-place rename via the NITRO rename action. The rename source must
	// be the CURRENT LIVE name, which is held in state.Id (not the configured
	// name attribute, which is pinned to the originally-configured value).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		renamePayload := lb.Lbgroup{
			Name:    state.Id.ValueString(),
			Newname: data.Newname.ValueString(),
		}
		err := r.client.ActOnResource(service.Lbgroup.Type(), &renamePayload, "rename")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename lbgroup, got error: %s", err))
			return
		}
		// The live object is now named newname; track it in the ID.
		data.Id = types.StringValue(data.Newname.ValueString())
		tflog.Trace(ctx, "Renamed lbgroup resource")
	}

	// Check whether any updatable attribute changed.
	hasChange := false
	if !data.Backuppersistencetimeout.Equal(state.Backuppersistencetimeout) {
		hasChange = true
	}
	if !data.Cookiedomain.Equal(state.Cookiedomain) {
		hasChange = true
	}
	if !data.Cookiename.Equal(state.Cookiename) {
		hasChange = true
	}
	if !data.Mastervserver.Equal(state.Mastervserver) {
		hasChange = true
	}
	if !data.Persistencebackup.Equal(state.Persistencebackup) {
		hasChange = true
	}
	if !data.Persistencetype.Equal(state.Persistencetype) {
		hasChange = true
	}
	if !data.Persistmask.Equal(state.Persistmask) {
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		hasChange = true
	}
	if !data.Timeout.Equal(state.Timeout) {
		hasChange = true
	}
	if !data.Usevserverpersistency.Equal(state.Usevserverpersistency) {
		hasChange = true
	}
	if !data.V6persistmasklen.Equal(state.V6persistmasklen) {
		hasChange = true
	}

	if hasChange {
		lbgroup := lbgroupGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Key off the current live name (after a possible rename).
		lbgroup.Name = data.Id.ValueString()
		_, err := r.client.UpdateResource(service.Lbgroup.Type(), data.Id.ValueString(), &lbgroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbgroup, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated lbgroup resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lbgroup resource, skipping update")
	}

	// Read the updated state back
	if !r.readLbgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbgroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbgroup resource")

	// Named resource - delete by live ID (correct even after a rename).
	err := r.client.DeleteResource(service.Lbgroup.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbgroup resource")
}

// Helper function to read lbgroup data from API
func (r *LbgroupResource) readLbgroupFromApi(ctx context.Context, data *LbgroupResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the live name.
	lbgroupName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lbgroup.Type(), lbgroupName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbgroup, got error: %s", err))
		return false
	}

	lbgroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
