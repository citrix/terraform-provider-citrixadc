package vpnnexthopserver

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
var _ resource.Resource = &VpnnexthopserverResource{}
var _ resource.ResourceWithConfigure = (*VpnnexthopserverResource)(nil)
var _ resource.ResourceWithImportState = (*VpnnexthopserverResource)(nil)

func NewVpnnexthopserverResource() resource.Resource {
	return &VpnnexthopserverResource{}
}

// VpnnexthopserverResource defines the resource implementation.
type VpnnexthopserverResource struct {
	client *service.NitroClient
}

func (r *VpnnexthopserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnnexthopserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnnexthopserver"
}

func (r *VpnnexthopserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnnexthopserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnnexthopserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnnexthopserver resource")

	vpnnexthopserver := vpnnexthopserverGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpnnexthopserverName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName, &vpnnexthopserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnnexthopserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnnexthopserver resource")

	// Set ID for the resource before reading state (single unique attr - plain value)
	data.Id = types.StringValue(vpnnexthopserverName)

	// Read the updated state back
	if !r.readVpnnexthopserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnnexthopserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnnexthopserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnnexthopserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnnexthopserver resource")

	found := r.readVpnnexthopserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnnexthopserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnnexthopserverResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnnexthopserver resource")

	// All attributes are non-updateable (ForceNew/RequiresReplace), so Terraform
	// recreates the resource on any change and Update carries no NITRO write.
	// Re-sync the state from the appliance.
	if !r.readVpnnexthopserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnnexthopserver not found immediately after update")
		}
		return
	}

	tflog.Trace(ctx, "Updated vpnnexthopserver resource")

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnnexthopserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnnexthopserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnnexthopserver resource")

	// Named resource - delete using DeleteResource keyed off the live ID
	vpnnexthopserverName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnnexthopserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnnexthopserver resource")
}

// Helper function to read vpnnexthopserver data from API.
// Returns false (without an error diagnostic) when the resource is not found so
// callers can drop it from state.
func (r *VpnnexthopserverResource) readVpnnexthopserverFromApi(ctx context.Context, data *VpnnexthopserverResourceModel, diags *diag.Diagnostics) bool {
	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vpnnexthopserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnnexthopserver.Type(), vpnnexthopserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnnexthopserver, got error: %s", err))
		return false
	}

	vpnnexthopserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}
