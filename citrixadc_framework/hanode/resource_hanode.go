package hanode

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &HanodeResource{}
var _ resource.ResourceWithConfigure = (*HanodeResource)(nil)
var _ resource.ResourceWithImportState = (*HanodeResource)(nil)

func NewHanodeResource() resource.Resource {
	return &HanodeResource{}
}

// HanodeResource defines the resource implementation.
type HanodeResource struct {
	client *service.NitroClient
}

func (r *HanodeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *HanodeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hanode"
}

func (r *HanodeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *HanodeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data HanodeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating hanode resource")

	hanode := hanodeGetThePayloadFromthePlan(ctx, &data)

	// The hanode id is the resource identifier. Self node (id == 0) has no name in the
	// NITRO URL and is configured via UpdateUnnamedResource; peer nodes (id 1-64) are
	// named and added via AddResource. This mirrors the SDK v2 createHanodeFunc.
	hanodeName := strconv.Itoa(int(data.Hanodeid.ValueInt64()))

	var err error
	if data.Hanodeid.ValueInt64() != 0 {
		_, err = r.client.AddResource(service.Hanode.Type(), hanodeName, &hanode)
	} else {
		err = r.client.UpdateUnnamedResource(service.Hanode.Type(), &hanode)
	}
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create hanode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created hanode resource")

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(hanodeName)

	// Read the updated state back
	if !r.readHanodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "hanode not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data HanodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading hanode resource")

	found := r.readHanodeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *HanodeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state HanodeResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating hanode resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Deadinterval.Equal(state.Deadinterval) {
		tflog.Debug(ctx, "deadinterval has changed for hanode")
		hasChange = true
	}
	if !data.Failsafe.Equal(state.Failsafe) {
		tflog.Debug(ctx, "failsafe has changed for hanode")
		hasChange = true
	}
	if !data.Haprop.Equal(state.Haprop) {
		tflog.Debug(ctx, "haprop has changed for hanode")
		hasChange = true
	}
	if !data.Hastatus.Equal(state.Hastatus) {
		tflog.Debug(ctx, "hastatus has changed for hanode")
		hasChange = true
	}
	if !data.Hasync.Equal(state.Hasync) {
		tflog.Debug(ctx, "hasync has changed for hanode")
		hasChange = true
	}
	if !data.Hellointerval.Equal(state.Hellointerval) {
		tflog.Debug(ctx, "hellointerval has changed for hanode")
		hasChange = true
	}
	if !data.Inc.Equal(state.Inc) {
		tflog.Debug(ctx, "inc has changed for hanode")
		hasChange = true
	}
	if !data.Maxflips.Equal(state.Maxflips) {
		tflog.Debug(ctx, "maxflips has changed for hanode")
		hasChange = true
	}
	if !data.Maxfliptime.Equal(state.Maxfliptime) {
		tflog.Debug(ctx, "maxfliptime has changed for hanode")
		hasChange = true
	}
	if !data.Syncstatusstrictmode.Equal(state.Syncstatusstrictmode) {
		tflog.Debug(ctx, "syncstatusstrictmode has changed for hanode")
		hasChange = true
	}
	if !data.Syncvlan.Equal(state.Syncvlan) {
		tflog.Debug(ctx, "syncvlan has changed for hanode")
		hasChange = true
	}

	if hasChange {
		// Update uses PUT /config/hanode with the id carried in the body (unnamed),
		// matching the SDK v2 updateHanodeFunc.
		hanode := hanodeGetTheUpdatablePayloadFromThePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Hanode.Type(), &hanode)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update hanode, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated hanode resource")
	} else {
		tflog.Debug(ctx, "No changes detected for hanode resource, skipping update")
	}

	// Read the updated state back
	if !r.readHanodeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "hanode not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *HanodeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data HanodeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting hanode resource")

	// Named resource - delete using DeleteResource keyed on the hanode id.
	err := r.client.DeleteResource(service.Hanode.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete hanode, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted hanode resource")
}

// Helper function to read hanode data from API. Returns false when the resource no
// longer exists on the ADC so the caller can drop it from state.
func (r *HanodeResource) readHanodeFromApi(ctx context.Context, data *HanodeResourceModel, diags *diag.Diagnostics) bool {
	hanodeName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Hanode.Type(), hanodeName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read hanode, got error: %s", err))
		return false
	}

	hanodeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
