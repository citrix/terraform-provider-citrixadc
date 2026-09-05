package lsnclient

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
var _ resource.Resource = &LsnclientResource{}
var _ resource.ResourceWithConfigure = (*LsnclientResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientResource)(nil)

func NewLsnclientResource() resource.Resource {
	return &LsnclientResource{}
}

// LsnclientResource defines the resource implementation.
type LsnclientResource struct {
	client *service.NitroClient
}

func (r *LsnclientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient"
}

func (r *LsnclientResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient resource")

	// Create API request body from the model
	lsnclient := lsnclientGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	clientname := data.Clientname.ValueString()
	_, err := r.client.AddResource(service.Lsnclient.Type(), clientname, &lsnclient)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnclient resource")

	// Set ID for the resource before reading state
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", clientname))

	// Read the updated state back
	if !r.readLsnclientFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnclient not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient resource")

	found := r.readLsnclientFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *LsnclientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnclientResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnclient resource")

	// lsnclient has only the ForceNew "clientname" attribute; there are no
	// NITRO-updatable attributes, so no update API call is required. Any change to
	// clientname triggers RequiresReplace (destroy/recreate) instead of Update.

	// Read the updated state back
	if !r.readLsnclientFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lsnclient not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient resource")

	// Named resource - delete using DeleteResource
	clientname := data.Id.ValueString()
	err := r.client.DeleteResource(service.Lsnclient.Type(), clientname)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnclient, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnclient resource")
}

// Helper function to read lsnclient data from API
func (r *LsnclientResource) readLsnclientFromApi(ctx context.Context, data *LsnclientResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	clientnameName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lsnclient.Type(), clientnameName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient, got error: %s", err))
		return false
	}

	lsnclientSetAttrFromGet(ctx, data, getResponseData)

	return true
}
