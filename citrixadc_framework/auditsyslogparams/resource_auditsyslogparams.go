package auditsyslogparams

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
var _ resource.Resource = &AuditsyslogparamsResource{}
var _ resource.ResourceWithConfigure = (*AuditsyslogparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AuditsyslogparamsResource)(nil)

func NewAuditsyslogparamsResource() resource.Resource {
	return &AuditsyslogparamsResource{}
}

// AuditsyslogparamsResource defines the resource implementation.
type AuditsyslogparamsResource struct {
	client *service.NitroClient
}

func (r *AuditsyslogparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuditsyslogparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_auditsyslogparams"
}

func (r *AuditsyslogparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuditsyslogparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuditsyslogparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating auditsyslogparams resource")

	// Create API request body from the model
	auditsyslogparams := auditsyslogparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed (singleton) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Auditsyslogparams.Type(), &auditsyslogparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create auditsyslogparams, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("auditsyslogparams-config")

	tflog.Trace(ctx, "Created auditsyslogparams resource")

	// Read the updated state back
	r.readAuditsyslogparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuditsyslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading auditsyslogparams resource")

	r.readAuditsyslogparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AuditsyslogparamsResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating auditsyslogparams resource")

	// Detect attributes removed from config so they can be unset (reverted to
	// the appliance default) rather than pushed via update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dateformat.Equal(state.Dateformat) {
		if config.Dateformat.IsNull() {
			attributesToUnset = append(attributesToUnset, "dateformat")
		} else {
			hasChange = true
		}
	}
	if !data.Tcp.Equal(state.Tcp) {
		if config.Tcp.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcp")
		} else {
			hasChange = true
		}
	}
	if !data.Protocolviolations.Equal(state.Protocolviolations) {
		if config.Protocolviolations.IsNull() {
			attributesToUnset = append(attributesToUnset, "protocolviolations")
		} else {
			hasChange = true
		}
	}
	if !data.Streamanalytics.Equal(state.Streamanalytics) {
		if config.Streamanalytics.IsNull() {
			attributesToUnset = append(attributesToUnset, "streamanalytics")
		} else {
			hasChange = true
		}
	}
	if !data.Denylistviolations.Equal(state.Denylistviolations) {
		hasChange = true
	}
	_ = hasChange

	// Create API request body from the model
	auditsyslogparams := auditsyslogparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed (singleton) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Auditsyslogparams.Type(), &auditsyslogparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update auditsyslogparams, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated auditsyslogparams resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Done after the update so any default the update payload
	// carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Auditsyslogparams.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset auditsyslogparams attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readAuditsyslogparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuditsyslogparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuditsyslogparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting auditsyslogparams resource")

	// For auditsyslogparams, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted auditsyslogparams resource from state")
}

// Helper function to read auditsyslogparams data from API
func (r *AuditsyslogparamsResource) readAuditsyslogparamsFromApi(ctx context.Context, data *AuditsyslogparamsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Auditsyslogparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read auditsyslogparams, got error: %s", err))
		return
	}

	auditsyslogparamsSetAttrFromGet(ctx, data, getResponseData)

}
