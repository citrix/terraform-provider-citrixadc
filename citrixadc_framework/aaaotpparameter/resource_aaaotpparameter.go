package aaaotpparameter

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &AaaotpparameterResource{}
var _ resource.ResourceWithConfigure = (*AaaotpparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AaaotpparameterResource)(nil)

func NewAaaotpparameterResource() resource.Resource {
	return &AaaotpparameterResource{}
}

// AaaotpparameterResource defines the resource implementation.
type AaaotpparameterResource struct {
	client *service.NitroClient
}

func (r *AaaotpparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaotpparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaotpparameter"
}

func (r *AaaotpparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaotpparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaaotpparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaotpparameter resource")

	// Create API request body from the model
	aaaotpparameter := aaaotpparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaotpparameter.Type(), &aaaotpparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaotpparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaotpparameter resource")

	// Read the updated state back (also sets the ID)
	if !r.readAaaotpparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaaotpparameter not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaotpparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaotpparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaotpparameter resource")

	found := r.readAaaotpparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AaaotpparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaaotpparameterResourceModel

	// Read Terraform prior state to preserve ID
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

	tflog.Debug(ctx, "Updating aaaotpparameter resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Encryption.Equal(state.Encryption) {
		tflog.Debug(ctx, "encryption has changed for aaaotpparameter")
		if config.Encryption.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "encryption")
		} else {
			hasChange = true
		}
	}
	if !data.Maxotpdevices.Equal(state.Maxotpdevices) {
		tflog.Debug(ctx, "maxotpdevices has changed for aaaotpparameter")
		if config.Maxotpdevices.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "maxotpdevices")
		} else {
			hasChange = true
		}
	}
	if !data.Otptype.Equal(state.Otptype) {
		tflog.Debug(ctx, "otptype has changed for aaaotpparameter")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		aaaotpparameter := aaaotpparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaaotpparameter.Type(), &aaaotpparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaotpparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaaotpparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaaotpparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Aaaotpparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset aaaotpparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAaaotpparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "aaaotpparameter not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaotpparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaotpparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaotpparameter resource")

	// aaaotpparameter is a global configuration singleton and does not support a
	// DELETE operation (matches the SDK v2 behavior which only cleared the ID).
	// We simply remove it from state.
	tflog.Trace(ctx, "Deleted aaaotpparameter resource from state")
}

// Helper function to read aaaotpparameter data from API
func (r *AaaotpparameterResource) readAaaotpparameterFromApi(ctx context.Context, data *AaaotpparameterResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Aaaotpparameter.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaotpparameter, got error: %s", err))
		return false
	}

	aaaotpparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
