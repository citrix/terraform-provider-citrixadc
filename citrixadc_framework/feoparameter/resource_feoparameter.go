package feoparameter

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
var _ resource.Resource = &FeoparameterResource{}
var _ resource.ResourceWithConfigure = (*FeoparameterResource)(nil)
var _ resource.ResourceWithImportState = (*FeoparameterResource)(nil)

func NewFeoparameterResource() resource.Resource {
	return &FeoparameterResource{}
}

// FeoparameterResource defines the resource implementation.
type FeoparameterResource struct {
	client *service.NitroClient
}

func (r *FeoparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *FeoparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_feoparameter"
}

func (r *FeoparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *FeoparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data FeoparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating feoparameter resource")

	feoparameter := feoparameterGetThePayloadFromthePlan(ctx, &data)

	// Singleton (unnamed) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Feoparameter.Type(), &feoparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create feoparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("feoparameter-config")

	tflog.Trace(ctx, "Created feoparameter resource")

	// Read the updated state back
	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data FeoparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading feoparameter resource")

	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state FeoparameterResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (now null)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating feoparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Cssinlinethressize.Equal(state.Cssinlinethressize) {
		tflog.Debug(ctx, "cssinlinethressize has changed for feoparameter")
		if config.Cssinlinethressize.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cssinlinethressize")
		} else {
			hasChange = true
		}
	}
	if !data.Imginlinethressize.Equal(state.Imginlinethressize) {
		tflog.Debug(ctx, "imginlinethressize has changed for feoparameter")
		if config.Imginlinethressize.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "imginlinethressize")
		} else {
			hasChange = true
		}
	}
	if !data.Jpegqualitypercent.Equal(state.Jpegqualitypercent) {
		tflog.Debug(ctx, "jpegqualitypercent has changed for feoparameter")
		if config.Jpegqualitypercent.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "jpegqualitypercent")
		} else {
			hasChange = true
		}
	}
	if !data.Jsinlinethressize.Equal(state.Jsinlinethressize) {
		tflog.Debug(ctx, "jsinlinethressize has changed for feoparameter")
		if config.Jsinlinethressize.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "jsinlinethressize")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		feoparameter := feoparameterGetThePayloadFromthePlan(ctx, &data)

		// Singleton (unnamed) resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Feoparameter.Type(), &feoparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update feoparameter, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated feoparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for feoparameter resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// NITRO defaults. feoparameter is a singleton, so the unset id payload is empty.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Feoparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset feoparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readFeoparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FeoparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FeoparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting feoparameter resource")

	// feoparameter is a global configuration singleton and does not support a
	// NITRO DELETE operation (matches SDK v2). Just remove it from state.
	tflog.Trace(ctx, "Deleted feoparameter resource from state")
}

// Helper function to read feoparameter data from API
func (r *FeoparameterResource) readFeoparameterFromApi(ctx context.Context, data *FeoparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Feoparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read feoparameter, got error: %s", err))
		return
	}

	feoparameterSetAttrFromGet(ctx, data, getResponseData)

}
