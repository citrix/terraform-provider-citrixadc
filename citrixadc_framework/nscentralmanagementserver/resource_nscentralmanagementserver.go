package nscentralmanagementserver

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
var _ resource.Resource = &NscentralmanagementserverResource{}
var _ resource.ResourceWithConfigure = (*NscentralmanagementserverResource)(nil)
var _ resource.ResourceWithImportState = (*NscentralmanagementserverResource)(nil)
var _ resource.ResourceWithValidateConfig = (*NscentralmanagementserverResource)(nil)

func NewNscentralmanagementserverResource() resource.Resource {
	return &NscentralmanagementserverResource{}
}

// NscentralmanagementserverResource defines the resource implementation.
type NscentralmanagementserverResource struct {
	client *service.NitroClient
}

func (r *NscentralmanagementserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NscentralmanagementserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nscentralmanagementserver"
}

func (r *NscentralmanagementserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NscentralmanagementserverResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data NscentralmanagementserverResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// NITRO CLI requires exactly one of ipaddress / servername as the address of the
	// central management server. tfdata marks both is_required:false, so enforce here.
	ipSet := !data.Ipaddress.IsNull() && !data.Ipaddress.IsUnknown() && data.Ipaddress.ValueString() != ""
	nameSet := !data.Servername.IsNull() && !data.Servername.IsUnknown() && data.Servername.ValueString() != ""

	if !ipSet && !nameSet {
		resp.Diagnostics.AddError(
			"Missing Required Attribute",
			"Exactly one of \"ipaddress\" or \"servername\" must be set for nscentralmanagementserver.",
		)
	} else if ipSet && nameSet {
		resp.Diagnostics.AddError(
			"Conflicting Attributes",
			"Only one of \"ipaddress\" or \"servername\" may be set for nscentralmanagementserver, not both.",
		)
	}
}

func (r *NscentralmanagementserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NscentralmanagementserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nscentralmanagementserver resource")
	// Get payload from plan (regular attributes)
	nscentralmanagementserver := nscentralmanagementserverGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	nscentralmanagementserverGetThePayloadFromtheConfig(ctx, &config, &nscentralmanagementserver)

	// Make API call
	// Named resource - use AddResource
	type_value := data.Type.ValueString()
	_, err := r.client.AddResource(service.Nscentralmanagementserver.Type(), type_value, &nscentralmanagementserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nscentralmanagementserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nscentralmanagementserver resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Type.ValueString()))

	// Read the updated state back
	r.readNscentralmanagementserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscentralmanagementserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NscentralmanagementserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nscentralmanagementserver resource")

	r.readNscentralmanagementserverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Resource was deleted out-of-band; remove it from state so it can be recreated
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscentralmanagementserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO exposes no update endpoint for nscentralmanagementserver (only add/delete/get).
	// Every schema attribute uses RequiresReplace, so Terraform never reaches Update with a
	// real change; this body is a documented no-op that just re-reads live state.
	var data, state NscentralmanagementserverResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for nscentralmanagementserver; all attributes are RequiresReplace")

	// Read the updated state back
	r.readNscentralmanagementserverFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscentralmanagementserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NscentralmanagementserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nscentralmanagementserver resource")
	// Named resource - delete using DeleteResource
	type_value := data.Type.ValueString()
	err := r.client.DeleteResource(service.Nscentralmanagementserver.Type(), type_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nscentralmanagementserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nscentralmanagementserver resource")
}

// Helper function to read nscentralmanagementserver data from API
func (r *NscentralmanagementserverResource) readNscentralmanagementserverFromApi(ctx context.Context, data *NscentralmanagementserverResourceModel, diags *diag.Diagnostics) {

	// Case 2: Find with single ID attribute - ID is the plain value
	type_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Nscentralmanagementserver.Type(), type_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nscentralmanagementserver, got error: %s", err))
		return
	}

	nscentralmanagementserverSetAttrFromGet(ctx, data, getResponseData)

}
