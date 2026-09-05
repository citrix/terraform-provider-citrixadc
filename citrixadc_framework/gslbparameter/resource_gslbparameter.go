package gslbparameter

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
var _ resource.Resource = &GslbparameterResource{}
var _ resource.ResourceWithConfigure = (*GslbparameterResource)(nil)
var _ resource.ResourceWithImportState = (*GslbparameterResource)(nil)

func NewGslbparameterResource() resource.Resource {
	return &GslbparameterResource{}
}

// GslbparameterResource defines the resource implementation.
type GslbparameterResource struct {
	client *service.NitroClient
}

func (r *GslbparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbparameter"
}

func (r *GslbparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbparameter resource")

	// Create API request body from the plan
	gslbparameter := gslbparameterGetThePayloadFromthePlan(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Gslbparameter.Type(), &gslbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbparameter-config")

	tflog.Trace(ctx, "Created gslbparameter resource")

	// Read the updated state back
	r.readGslbparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbparameter resource")

	r.readGslbparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state GslbparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbparameter resource")

	// Determine which attributes changed and which were removed from config (unset)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Automaticconfigsync.Equal(state.Automaticconfigsync) {
		if config.Automaticconfigsync.IsNull() {
			attributesToUnset = append(attributesToUnset, "automaticconfigsync")
		} else {
			hasChange = true
		}
	}
	if !data.Dropldnsreq.Equal(state.Dropldnsreq) {
		if config.Dropldnsreq.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropldnsreq")
		} else {
			hasChange = true
		}
	}
	if !data.Gslbconfigsyncmonitor.Equal(state.Gslbconfigsyncmonitor) {
		if config.Gslbconfigsyncmonitor.IsNull() {
			attributesToUnset = append(attributesToUnset, "gslbconfigsyncmonitor")
		} else {
			hasChange = true
		}
	}
	if !data.Gslbsyncinterval.Equal(state.Gslbsyncinterval) {
		if config.Gslbsyncinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "gslbsyncinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Gslbsynclocfiles.Equal(state.Gslbsynclocfiles) {
		if config.Gslbsynclocfiles.IsNull() {
			attributesToUnset = append(attributesToUnset, "gslbsynclocfiles")
		} else {
			hasChange = true
		}
	}
	if !data.Gslbsyncmode.Equal(state.Gslbsyncmode) {
		if config.Gslbsyncmode.IsNull() {
			attributesToUnset = append(attributesToUnset, "gslbsyncmode")
		} else {
			hasChange = true
		}
	}
	if !data.Gslbsyncsaveconfigcommand.Equal(state.Gslbsyncsaveconfigcommand) {
		if config.Gslbsyncsaveconfigcommand.IsNull() {
			attributesToUnset = append(attributesToUnset, "gslbsyncsaveconfigcommand")
		} else {
			hasChange = true
		}
	}
	if !data.Ldnsentrytimeout.Equal(state.Ldnsentrytimeout) {
		if config.Ldnsentrytimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "ldnsentrytimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Mepkeepalivetimeout.Equal(state.Mepkeepalivetimeout) {
		if config.Mepkeepalivetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "mepkeepalivetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Rtttolerance.Equal(state.Rtttolerance) {
		if config.Rtttolerance.IsNull() {
			attributesToUnset = append(attributesToUnset, "rtttolerance")
		} else {
			hasChange = true
		}
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		if config.Undefaction.IsNull() {
			attributesToUnset = append(attributesToUnset, "undefaction")
		} else {
			hasChange = true
		}
	}
	if !data.V6ldnsmasklen.Equal(state.V6ldnsmasklen) {
		if config.V6ldnsmasklen.IsNull() {
			attributesToUnset = append(attributesToUnset, "v6ldnsmasklen")
		} else {
			hasChange = true
		}
	}
	// Attributes without a documented server default (not unset-eligible) still
	// participate in normal updates.
	if !data.Gslbsvcstatedelaytime.Equal(state.Gslbsvcstatedelaytime) {
		hasChange = true
	}
	if !data.Svcstatelearningtime.Equal(state.Svcstatelearningtime) {
		hasChange = true
	}
	if !data.Ldnsmask.Equal(state.Ldnsmask) {
		hasChange = true
	}
	if !data.Ldnsprobeorder.Equal(state.Ldnsprobeorder) {
		hasChange = true
	}
	if !data.Sourceipwhitelisting.Equal(state.Sourceipwhitelisting) {
		hasChange = true
	}
	if !data.Usekrpcchannelforsync.Equal(state.Usekrpcchannelforsync) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		gslbparameter := gslbparameterGetThePayloadFromthePlan(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Gslbparameter.Type(), &gslbparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated gslbparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for gslbparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts to defaults
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Gslbparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset gslbparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readGslbparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbparameter resource")

	// For gslbparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbparameter resource from state")
}

// Helper function to read gslbparameter data from API
func (r *GslbparameterResource) readGslbparameterFromApi(ctx context.Context, data *GslbparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbparameter, got error: %s", err))
		return
	}

	gslbparameterSetAttrFromGet(ctx, data, getResponseData)

}
