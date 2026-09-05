package responderparam

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
var _ resource.Resource = &ResponderparamResource{}
var _ resource.ResourceWithConfigure = (*ResponderparamResource)(nil)
var _ resource.ResourceWithImportState = (*ResponderparamResource)(nil)

func NewResponderparamResource() resource.Resource {
	return &ResponderparamResource{}
}

// ResponderparamResource defines the resource implementation.
type ResponderparamResource struct {
	client *service.NitroClient
}

func (r *ResponderparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ResponderparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_responderparam"
}

func (r *ResponderparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ResponderparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ResponderparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating responderparam resource")

	// Build the NITRO payload from the plan (singleton).
	responderparam := responderparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call — singleton uses UpdateUnnamedResource (matches SDK v2).
	err := r.client.UpdateUnnamedResource(service.Responderparam.Type(), &responderparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create responderparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("responderparam-config")

	tflog.Trace(ctx, "Created responderparam resource")

	// Read the updated state back
	r.readResponderparamFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ResponderparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading responderparam resource")

	r.readResponderparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ResponderparamResourceModel

	// Read Terraform prior state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config (attributes removed from config are null here, not in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating responderparam resource")

	// Determine changed attributes and which were removed from config (-> unset).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, "timeout has changed for responderparam")
		if config.Timeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timeout")
		} else {
			hasChange = true
		}
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		tflog.Debug(ctx, "undefaction has changed for responderparam")
		if config.Undefaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (singleton).
		responderparam := responderparamGetThePayloadFromtheConfig(ctx, &data)

		// Make API call — singleton uses UpdateUnnamedResource (matches SDK v2).
		err := r.client.UpdateUnnamedResource(service.Responderparam.Type(), &responderparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update responderparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated responderparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for responderparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Responderparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset responderparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readResponderparamFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ResponderparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ResponderparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting responderparam resource")

	// For responderparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted responderparam resource from state")
}

// Helper function to read responderparam data from API
func (r *ResponderparamResource) readResponderparamFromApi(ctx context.Context, data *ResponderparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Responderparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read responderparam, got error: %s", err))
		return
	}

	responderparamSetAttrFromGet(ctx, data, getResponseData)

}
