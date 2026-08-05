package vpnformssoaction

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
var _ resource.Resource = &VpnformssoactionResource{}
var _ resource.ResourceWithConfigure = (*VpnformssoactionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnformssoactionResource)(nil)

func NewVpnformssoactionResource() resource.Resource {
	return &VpnformssoactionResource{}
}

// VpnformssoactionResource defines the resource implementation.
type VpnformssoactionResource struct {
	client *service.NitroClient
}

func (r *VpnformssoactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnformssoactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnformssoaction"
}

func (r *VpnformssoactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnformssoactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnformssoactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnformssoaction resource")

	vpnformssoaction := vpnformssoactionGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	vpnformssoactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnformssoaction.Type(), vpnformssoactionName, &vpnformssoaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnformssoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnformssoaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vpnformssoactionName)

	// Read the updated state back
	if !r.readVpnformssoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnformssoaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnformssoactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnformssoaction resource")

	found := r.readVpnformssoactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnformssoactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnformssoactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnformssoaction resource")

	// Only the non-ForceNew attributes can reach Update; the ForceNew attributes
	// (actionurl, name, passwdfield, ssosuccessrule, userfield) trigger a replace.
	hasChange := false
	if !data.Namevaluepair.Equal(state.Namevaluepair) {
		tflog.Debug(ctx, "namevaluepair has changed for vpnformssoaction")
		hasChange = true
	}
	if !data.Nvtype.Equal(state.Nvtype) {
		tflog.Debug(ctx, "nvtype has changed for vpnformssoaction")
		hasChange = true
	}
	if !data.Responsesize.Equal(state.Responsesize) {
		tflog.Debug(ctx, "responsesize has changed for vpnformssoaction")
		hasChange = true
	}
	if !data.Submitmethod.Equal(state.Submitmethod) {
		tflog.Debug(ctx, "submitmethod has changed for vpnformssoaction")
		hasChange = true
	}

	if hasChange {
		vpnformssoaction := vpnformssoactionGetThePayloadFromthePlan(ctx, &data)
		// Named resource - use UpdateResource
		vpnformssoactionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Vpnformssoaction.Type(), vpnformssoactionName, &vpnformssoaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnformssoaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnformssoaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnformssoaction resource, skipping update")
	}

	// Read the updated state back
	if !r.readVpnformssoactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnformssoaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnformssoactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnformssoactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnformssoaction resource")

	// Named resource - delete using DeleteResource
	vpnformssoactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnformssoaction.Type(), vpnformssoactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnformssoaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnformssoaction resource")
}

// Helper function to read vpnformssoaction data from API
func (r *VpnformssoactionResource) readVpnformssoactionFromApi(ctx context.Context, data *VpnformssoactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	vpnformssoactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnformssoaction.Type(), vpnformssoactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnformssoaction, got error: %s", err))
		return false
	}

	vpnformssoactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
