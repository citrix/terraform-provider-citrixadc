package sslechconfig

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SslechconfigResource{}
var _ resource.ResourceWithConfigure = (*SslechconfigResource)(nil)
var _ resource.ResourceWithImportState = (*SslechconfigResource)(nil)

func NewSslechconfigResource() resource.Resource {
	return &SslechconfigResource{}
}

// SslechconfigResource defines the resource implementation.
type SslechconfigResource struct {
	client *service.NitroClient
}

func (r *SslechconfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslechconfigResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslechconfig"
}

func (r *SslechconfigResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslechconfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslechconfigResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslechconfig resource")
	sslechconfig := sslechconfigGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	echconfigname_value := data.Echconfigname.ValueString()
	_, err := r.client.AddResource(service.Sslechconfig.Type(), echconfigname_value, &sslechconfig)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslechconfig, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslechconfig resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Echconfigname.ValueString()))

	// Read the updated state back
	r.readSslechconfigFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslechconfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslechconfigResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslechconfig resource")

	r.readSslechconfigFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslechconfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for sslechconfig (only add, delete, get,
	// get (all)). Every schema attribute is marked RequiresReplace, so Terraform
	// will never actually invoke Update with field changes. This body is a
	// documented no-op that preserves the prior ID and re-reads state.
	var data, state SslechconfigResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for sslechconfig; NITRO has no update endpoint and all attributes are RequiresReplace")

	r.readSslechconfigFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslechconfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslechconfigResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslechconfig resource")
	// Named resource - delete using DeleteResource
	echconfigname_value := data.Echconfigname.ValueString()
	err := r.client.DeleteResource(service.Sslechconfig.Type(), echconfigname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslechconfig, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslechconfig resource")
}

// Helper function to read sslechconfig data from API
func (r *SslechconfigResource) readSslechconfigFromApi(ctx context.Context, data *SslechconfigResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	echconfigname_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Sslechconfig.Type(), echconfigname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslechconfig, got error: %s", err))
		return
	}

	sslechconfigSetAttrFromGet(ctx, data, getResponseData)

}
