package sslpolicy

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
var _ resource.Resource = &SslpolicyResource{}
var _ resource.ResourceWithConfigure = (*SslpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*SslpolicyResource)(nil)

func NewSslpolicyResource() resource.Resource {
	return &SslpolicyResource{}
}

// SslpolicyResource defines the resource implementation.
type SslpolicyResource struct {
	client *service.NitroClient
}

func (r *SslpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslpolicy"
}

func (r *SslpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslpolicy resource")

	sslpolicy := sslpolicyGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Sslpolicy.Type(), name_value, &sslpolicy)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslpolicy resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(name_value)

	// Read the updated state back
	if !r.readSslpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslpolicy not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslpolicy resource")

	found := r.readSslpolicyFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslpolicyResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset candidates)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating sslpolicy resource")

	// Check if there are any changes in updateable attributes.
	// name is RequiresReplace and reqaction is not updateable via NITRO, so
	// neither is checked here.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Action.Equal(state.Action) {
		tflog.Debug(ctx, "action has changed for sslpolicy")
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for sslpolicy")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Rule.Equal(state.Rule) {
		tflog.Debug(ctx, "rule has changed for sslpolicy")
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for sslpolicy")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		sslpolicy := sslpolicyGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Sslpolicy.Type(), name_value, &sslpolicy)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslpolicy, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated sslpolicy resource")
	} else {
		tflog.Debug(ctx, "No changes detected for sslpolicy resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Sslpolicy.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslpolicy attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSslpolicyFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslpolicy not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslpolicy resource")

	// Named resource - delete using DeleteResource
	name_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Sslpolicy.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete sslpolicy, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted sslpolicy resource")
}

// Helper function to read sslpolicy data from API. Returns false when the
// resource no longer exists on the ADC.
func (r *SslpolicyResource) readSslpolicyFromApi(ctx context.Context, data *SslpolicyResourceModel, diags *diag.Diagnostics) bool {
	sslpolicyName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Sslpolicy.Type(), sslpolicyName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslpolicy, got error: %s", err))
		return false
	}

	sslpolicySetAttrFromGet(ctx, data, getResponseData)

	return true
}
