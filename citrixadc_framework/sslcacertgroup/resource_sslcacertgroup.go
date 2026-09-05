package sslcacertgroup

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
var _ resource.Resource = &SslcacertgroupResource{}
var _ resource.ResourceWithConfigure = (*SslcacertgroupResource)(nil)
var _ resource.ResourceWithImportState = (*SslcacertgroupResource)(nil)

func NewSslcacertgroupResource() resource.Resource {
	return &SslcacertgroupResource{}
}

// SslcacertgroupResource defines the resource implementation.
type SslcacertgroupResource struct {
	client *service.NitroClient
}

func (r *SslcacertgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslcacertgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcacertgroup"
}

func (r *SslcacertgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslcacertgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslcacertgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslcacertgroup resource")

	sslcacertgroup := sslcacertgroupGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	cacertgroupname_value := data.Cacertgroupname.ValueString()
	_, err := r.client.AddResource(service.Sslcacertgroup.Type(), cacertgroupname_value, &sslcacertgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslcacertgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslcacertgroup resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Cacertgroupname.ValueString()))

	// Read the updated state back
	if !r.readSslcacertgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcacertgroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslcacertgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslcacertgroup resource")

	found := r.readSslcacertgroupFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslcacertgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslcacertgroupResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslcacertgroup resource")

	// sslcacertgroup has no NITRO-updatable attributes: cacertgroupname is the
	// only user-settable attribute and it is ForceNew (RequiresReplace), so any
	// change to it triggers a destroy/recreate rather than an in-place update.
	// Nothing to push to NITRO here; just re-read the current state.

	// Read the updated state back
	if !r.readSslcacertgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslcacertgroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslcacertgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslcacertgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslcacertgroup resource")

	// Named resource - delete using DeleteResource
	cacertgroupname_value := data.Cacertgroupname.ValueString()
	err := r.client.DeleteResource(service.Sslcacertgroup.Type(), cacertgroupname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslcacertgroup, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslcacertgroup resource")
}

// Helper function to read sslcacertgroup data from API
func (r *SslcacertgroupResource) readSslcacertgroupFromApi(ctx context.Context, data *SslcacertgroupResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (cacertgroupname)
	cacertgroupname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslcacertgroup.Type(), cacertgroupname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslcacertgroup, got error: %s", err))
		return false
	}

	sslcacertgroupSetAttrFromGet(ctx, data, getResponseData)

	return true
}
