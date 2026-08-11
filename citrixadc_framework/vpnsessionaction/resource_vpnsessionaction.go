package vpnsessionaction

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
var _ resource.Resource = &VpnsessionactionResource{}
var _ resource.ResourceWithConfigure = (*VpnsessionactionResource)(nil)
var _ resource.ResourceWithImportState = (*VpnsessionactionResource)(nil)

func NewVpnsessionactionResource() resource.Resource {
	return &VpnsessionactionResource{}
}

// VpnsessionactionResource defines the resource implementation.
type VpnsessionactionResource struct {
	client *service.NitroClient
}

func (r *VpnsessionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnsessionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnsessionaction"
}

func (r *VpnsessionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnsessionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnsessionaction resource")

	// Create API request body from the model
	vpnsessionaction := vpnsessionactionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource
	vpnsessionactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Vpnsessionaction.Type(), vpnsessionactionName, &vpnsessionaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnsessionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnsessionaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vpnsessionactionName)

	// Read the updated state back
	if !r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsessionaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsessionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnsessionaction resource")

	found := r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnsessionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state VpnsessionactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset them)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnsessionaction resource")

	// Determine attributes removed from config so they can be unset (reverted
	// to their NITRO defaults) after the update.
	attributesToUnset := []string{}
	if !data.Advancedclientlessvpnmode.Equal(state.Advancedclientlessvpnmode) {
		if config.Advancedclientlessvpnmode.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "advancedclientlessvpnmode")
		}
	}

	// Create API request body from the model
	vpnsessionaction := vpnsessionactionGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use UpdateResource
	vpnsessionactionName := data.Name.ValueString()
	_, err := r.client.UpdateResource(service.Vpnsessionaction.Type(), vpnsessionactionName, &vpnsessionaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnsessionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated vpnsessionaction resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Vpnsessionaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset vpnsessionaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readVpnsessionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "vpnsessionaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnsessionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnsessionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnsessionaction resource")

	// Named resource - delete using DeleteResource keyed by ID (the live name)
	vpnsessionactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Vpnsessionaction.Type(), vpnsessionactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnsessionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnsessionaction resource")
}

// Helper function to read vpnsessionaction data from API
func (r *VpnsessionactionResource) readVpnsessionactionFromApi(ctx context.Context, data *VpnsessionactionResourceModel, diags *diag.Diagnostics) bool {

	// Named resource - ID is the plain name value
	vpnsessionactionName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Vpnsessionaction.Type(), vpnsessionactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnsessionaction, got error: %s", err))
		return false
	}

	vpnsessionactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
