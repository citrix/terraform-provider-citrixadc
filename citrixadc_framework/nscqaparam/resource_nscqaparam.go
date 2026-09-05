package nscqaparam

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
var _ resource.Resource = &NscqaparamResource{}
var _ resource.ResourceWithConfigure = (*NscqaparamResource)(nil)
var _ resource.ResourceWithImportState = (*NscqaparamResource)(nil)

func NewNscqaparamResource() resource.Resource {
	return &NscqaparamResource{}
}

// NscqaparamResource defines the resource implementation.
type NscqaparamResource struct {
	client *service.NitroClient
}

func (r *NscqaparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NscqaparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nscqaparam"
}

func (r *NscqaparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NscqaparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NscqaparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nscqaparam resource")

	nscqaparam := nscqaparamGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nscqaparam.Type(), &nscqaparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nscqaparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nscqaparam-config")

	tflog.Trace(ctx, "Created nscqaparam resource")

	// Read the updated state back
	r.readNscqaparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscqaparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NscqaparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nscqaparam resource")

	r.readNscqaparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscqaparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NscqaparamResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nscqaparam resource")

	// Determine which attributes were removed from config so they can be unset
	// (reverted to NITRO defaults) after the update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Harqretxdelay.Equal(state.Harqretxdelay) {
		if config.Harqretxdelay.IsNull() {
			attributesToUnset = append(attributesToUnset, "harqretxdelay")
		} else {
			hasChange = true
		}
	}
	if !data.Lr1probthresh.Equal(state.Lr1probthresh) {
		if config.Lr1probthresh.IsNull() {
			attributesToUnset = append(attributesToUnset, "lr1probthresh")
		} else {
			hasChange = true
		}
	}
	if !data.Minrttnet1.Equal(state.Minrttnet1) {
		if config.Minrttnet1.IsNull() {
			attributesToUnset = append(attributesToUnset, "minrttnet1")
		} else {
			hasChange = true
		}
	}
	if !data.Minrttnet2.Equal(state.Minrttnet2) {
		if config.Minrttnet2.IsNull() {
			attributesToUnset = append(attributesToUnset, "minrttnet2")
		} else {
			hasChange = true
		}
	}
	if !data.Minrttnet3.Equal(state.Minrttnet3) {
		if config.Minrttnet3.IsNull() {
			attributesToUnset = append(attributesToUnset, "minrttnet3")
		} else {
			hasChange = true
		}
	}

	// Create API request body from the model
	nscqaparam := nscqaparamGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nscqaparam.Type(), &nscqaparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nscqaparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated nscqaparam resource")

	_ = hasChange

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nscqaparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nscqaparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNscqaparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NscqaparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NscqaparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nscqaparam resource")

	// For nscqaparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nscqaparam resource from state")
}

// Helper function to read nscqaparam data from API
func (r *NscqaparamResource) readNscqaparamFromApi(ctx context.Context, data *NscqaparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nscqaparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nscqaparam, got error: %s", err))
		return
	}

	nscqaparamSetAttrFromGet(ctx, data, getResponseData)

}
