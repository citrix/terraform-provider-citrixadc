package tmsessionparameter

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
var _ resource.Resource = &TmsessionparameterResource{}
var _ resource.ResourceWithConfigure = (*TmsessionparameterResource)(nil)
var _ resource.ResourceWithImportState = (*TmsessionparameterResource)(nil)

func NewTmsessionparameterResource() resource.Resource {
	return &TmsessionparameterResource{}
}

// TmsessionparameterResource defines the resource implementation.
type TmsessionparameterResource struct {
	client *service.NitroClient
}

func (r *TmsessionparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TmsessionparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tmsessionparameter"
}

func (r *TmsessionparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TmsessionparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TmsessionparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tmsessionparameter resource")

	// Build the payload from the plan
	tmsessionparameter := tmsessionparameterGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource (NITRO exposes only update)
	err := r.client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tmsessionparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("tmsessionparameter-config")

	tflog.Trace(ctx, "Created tmsessionparameter resource")

	// Read the updated state back
	r.readTmsessionparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TmsessionparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tmsessionparameter resource")

	r.readTmsessionparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TmsessionparameterResourceModel

	// Read prior state (to detect changes and preserve ID)
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

	tflog.Debug(ctx, "Updating tmsessionparameter resource")

	// Detect changes in the updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Defaultauthorizationaction.Equal(state.Defaultauthorizationaction) {
		if config.Defaultauthorizationaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "defaultauthorizationaction")
		} else {
			hasChange = true
		}
	}
	if !data.Homepage.Equal(state.Homepage) {
		if config.Homepage.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "homepage")
		} else {
			hasChange = true
		}
	}
	if !data.Httponlycookie.Equal(state.Httponlycookie) {
		if config.Httponlycookie.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "httponlycookie")
		} else {
			hasChange = true
		}
	}
	if !data.Kcdaccount.Equal(state.Kcdaccount) {
		hasChange = true
	}
	if !data.Persistentcookie.Equal(state.Persistentcookie) {
		hasChange = true
	}
	if !data.Persistentcookievalidity.Equal(state.Persistentcookievalidity) {
		hasChange = true
	}
	if !data.Sesstimeout.Equal(state.Sesstimeout) {
		if config.Sesstimeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sesstimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Sso.Equal(state.Sso) {
		if config.Sso.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "sso")
		} else {
			hasChange = true
		}
	}
	if !data.Ssocredential.Equal(state.Ssocredential) {
		if config.Ssocredential.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ssocredential")
		} else {
			hasChange = true
		}
	}
	if !data.Ssodomain.Equal(state.Ssodomain) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		tmsessionparameter := tmsessionparameterGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Tmsessionparameter.Type(), &tmsessionparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tmsessionparameter, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated tmsessionparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tmsessionparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Tmsessionparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset tmsessionparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readTmsessionparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TmsessionparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TmsessionparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tmsessionparameter resource")

	// For tmsessionparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted tmsessionparameter resource from state")
}

// Helper function to read tmsessionparameter data from API
func (r *TmsessionparameterResource) readTmsessionparameterFromApi(ctx context.Context, data *TmsessionparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Tmsessionparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tmsessionparameter, got error: %s", err))
		return
	}

	tmsessionparameterSetAttrFromGet(ctx, data, getResponseData)

}
