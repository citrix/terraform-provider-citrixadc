package sslpolicylabel

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
var _ resource.Resource = &SslpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*SslpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*SslpolicylabelResource)(nil)

func NewSslpolicylabelResource() resource.Resource {
	return &SslpolicylabelResource{}
}

// SslpolicylabelResource defines the resource implementation.
type SslpolicylabelResource struct {
	client *service.NitroClient
}

func (r *SslpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpolicylabel"
}

func (r *SslpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpolicylabel resource")

	// Create API request body from the model
	sslpolicylabel := sslpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (NITRO exposes add via POST)
	labelname_value := data.Labelname.ValueString()
	_, err := r.client.AddResource(service.Sslpolicylabel.Type(), labelname_value, &sslpolicylabel)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslpolicylabel resource")

	// Set ID for the resource before reading state (single unique attribute - plain value)
	data.Id = types.StringValue(data.Labelname.ValueString())

	// Read the updated state back
	if !r.readSslpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslpolicylabel not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslpolicylabel resource")

	found := r.readSslpolicylabelFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// sslpolicylabel has no NITRO update endpoint (only add/get/delete) and every
	// schema attribute (labelname, type) is RequiresReplace, so Terraform never
	// actually routes a change through Update. This body is a documented no-op that
	// preserves the prior ID and re-reads live state. (Pattern 5)
	var data, state SslpolicylabelResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for sslpolicylabel; all attributes are RequiresReplace")

	// Read the current state back
	if !r.readSslpolicylabelFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslpolicylabel not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslpolicylabel resource")

	// Named resource - delete using DeleteResource
	labelname_value := data.Labelname.ValueString()
	err := r.client.DeleteResource(service.Sslpolicylabel.Type(), labelname_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslpolicylabel, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslpolicylabel resource")
}

// Helper function to read sslpolicylabel data from API.
// Returns false (without error) when the resource no longer exists on the ADC.
func (r *SslpolicylabelResource) readSslpolicylabelFromApi(ctx context.Context, data *SslpolicylabelResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (labelname)
	labelname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslpolicylabel.Type(), labelname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslpolicylabel, got error: %s", err))
		return false
	}

	sslpolicylabelSetAttrFromGet(ctx, data, getResponseData)

	return true
}
