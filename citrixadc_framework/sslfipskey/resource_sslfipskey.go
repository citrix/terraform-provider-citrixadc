package sslfipskey

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
var _ resource.Resource = &SslfipskeyResource{}
var _ resource.ResourceWithConfigure = (*SslfipskeyResource)(nil)
var _ resource.ResourceWithImportState = (*SslfipskeyResource)(nil)

func NewSslfipskeyResource() resource.Resource {
	return &SslfipskeyResource{}
}

// SslfipskeyResource defines the resource implementation.
type SslfipskeyResource struct {
	client *service.NitroClient
}

func (r *SslfipskeyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslfipskeyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslfipskey"
}

func (r *SslfipskeyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslfipskeyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslfipskey resource")

	sslfipskey := sslfipskeyGetThePayloadFromtheConfig(ctx, &data)

	// sslfipskey is created via the NITRO "create" action (matches SDK v2
	// client.ActOnResource(..., "create")), not a plain AddResource.
	err := r.client.ActOnResource(service.Sslfipskey.Type(), &sslfipskey, "create")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslfipskey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslfipskey resource")

	// Single unique attribute: ID is the fipskeyname (matches SDK v2 d.SetId).
	data.Id = types.StringValue(data.Fipskeyname.ValueString())

	// Read the updated state back
	if !r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslfipskey not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipskeyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslfipskey resource")

	found := r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslfipskeyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SslfipskeyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// sslfipskey has no updateable attributes in SDK v2 (every attribute is
	// ForceNew), so there is no NITRO write here. Any attribute change forces
	// replacement via the RequiresReplace plan modifiers. Simply re-read the
	// live object so computed values stay consistent.
	tflog.Debug(ctx, "Updating sslfipskey resource (read-only refresh; all attributes force replacement)")

	if !r.readSslfipskeyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslfipskey not found during update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslfipskeyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslfipskeyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslfipskey resource")

	// Named resource - delete by fipskeyname (the ID), matching SDK v2.
	err := r.client.DeleteResource(service.Sslfipskey.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslfipskey, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslfipskey resource")
}

// Helper function to read sslfipskey data from API. Returns false (without an
// error diagnostic) when the resource no longer exists on the ADC.
func (r *SslfipskeyResource) readSslfipskeyFromApi(ctx context.Context, data *SslfipskeyResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain fipskeyname value.
	fipskeyname_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslfipskey.Type(), fipskeyname_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslfipskey, got error: %s", err))
		return false
	}

	sslfipskeySetAttrFromGet(ctx, data, getResponseData)

	return true
}
