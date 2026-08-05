package vpntrafficpolicy

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
var _ resource.Resource = &VpntrafficpolicyResource{}
var _ resource.ResourceWithConfigure = (*VpntrafficpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VpntrafficpolicyResource)(nil)

func NewVpntrafficpolicyResource() resource.Resource {
	return &VpntrafficpolicyResource{}
}

// VpntrafficpolicyResource defines the resource implementation.
type VpntrafficpolicyResource struct {
	client *service.NitroClient
}

func (r *VpntrafficpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpntrafficpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpntrafficpolicy"
}

func (r *VpntrafficpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpntrafficpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpntrafficpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpntrafficpolicy resource")

	vpntrafficpolicy := vpntrafficpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpntrafficpolicy.Type(), name_value, &vpntrafficpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpntrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpntrafficpolicy resource")

	// Set ID for the resource before reading state (plain single-key value)
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readVpntrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpntrafficpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpntrafficpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpntrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpntrafficpolicy resource")

	found := r.readVpntrafficpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpntrafficpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpntrafficpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpntrafficpolicy resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for vpntrafficpolicy")
		hasChange = true
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for vpntrafficpolicy")
		hasChange = true
	}

	if hasChange {
		vpntrafficpolicy := vpntrafficpolicyGetThePayloadFromtheConfig(ctx, &data)
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpntrafficpolicy.Type(), name_value, &vpntrafficpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpntrafficpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated vpntrafficpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpntrafficpolicy resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpntrafficpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpntrafficpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpntrafficpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpntrafficpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpntrafficpolicy resource")
	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpntrafficpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpntrafficpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpntrafficpolicy resource")
}

// Helper function to read vpntrafficpolicy data from API
func (r *VpntrafficpolicyResource) readVpntrafficpolicyFromApi(ctx context.Context, data *VpntrafficpolicyResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpntrafficpolicy.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpntrafficpolicy, got error: %s", err))
		return false
	}

	vpntrafficpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
