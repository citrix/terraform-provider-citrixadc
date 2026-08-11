package nslicenseparameters

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
var _ resource.Resource = &NslicenseparametersResource{}
var _ resource.ResourceWithConfigure = (*NslicenseparametersResource)(nil)
var _ resource.ResourceWithImportState = (*NslicenseparametersResource)(nil)

func NewNslicenseparametersResource() resource.Resource {
	return &NslicenseparametersResource{}
}

// NslicenseparametersResource defines the resource implementation.
type NslicenseparametersResource struct {
	client *service.NitroClient
}

func (r *NslicenseparametersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NslicenseparametersResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nslicenseparameters"
}

func (r *NslicenseparametersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NslicenseparametersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nslicenseparameters resource")

	// Build the payload from the plan and push it to the ADC.
	// nslicenseparameters is a singleton config object (no primary key), so it is
	// written with UpdateUnnamedResource, mirroring the SDK v2 implementation.
	nslicenseparameters := nslicenseparametersGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Nslicenseparameters.Type(), &nslicenseparameters)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nslicenseparameters, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nslicenseparameters-config")

	tflog.Trace(ctx, "Created nslicenseparameters resource")

	// Read the updated state back
	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nslicenseparameters resource")

	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NslicenseparametersResourceModel

	// Read Terraform prior state, plan, and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nslicenseparameters resource")

	// Determine which attributes changed and which were removed from config
	// (and thus should be unset back to their NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Alert1gracetimeout.Equal(state.Alert1gracetimeout) {
		if config.Alert1gracetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "alert1gracetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Alert2gracetimeout.Equal(state.Alert2gracetimeout) {
		if config.Alert2gracetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "alert2gracetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Heartbeatinterval.Equal(state.Heartbeatinterval) {
		if config.Heartbeatinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "heartbeatinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Inventoryrefreshinterval.Equal(state.Inventoryrefreshinterval) {
		if config.Inventoryrefreshinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "inventoryrefreshinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Licenseexpiryalerttime.Equal(state.Licenseexpiryalerttime) {
		if config.Licenseexpiryalerttime.IsNull() {
			attributesToUnset = append(attributesToUnset, "licenseexpiryalerttime")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		nslicenseparameters := nslicenseparametersGetThePayloadFromtheConfig(ctx, &data)

		// Make API call (singleton config object - UpdateUnnamedResource)
		err := r.client.UpdateUnnamedResource(service.Nslicenseparameters.Type(), &nslicenseparameters)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nslicenseparameters, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nslicenseparameters resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nslicenseparameters resource, skipping update")
	}

	// Unset attributes that were removed from configuration (revert to ADC defaults)
	// Singleton resource - no identity fields required in the unset payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nslicenseparameters.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nslicenseparameters attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNslicenseparametersFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NslicenseparametersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NslicenseparametersResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nslicenseparameters resource")

	// For nslicenseparameters, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nslicenseparameters resource from state")
}

// Helper function to read nslicenseparameters data from API
func (r *NslicenseparametersResource) readNslicenseparametersFromApi(ctx context.Context, data *NslicenseparametersResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nslicenseparameters.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nslicenseparameters, got error: %s", err))
		return
	}

	nslicenseparametersSetAttrFromGet(ctx, data, getResponseData)

}
