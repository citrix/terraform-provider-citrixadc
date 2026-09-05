package systemparameter

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
var _ resource.Resource = &SystemparameterResource{}
var _ resource.ResourceWithConfigure = (*SystemparameterResource)(nil)
var _ resource.ResourceWithImportState = (*SystemparameterResource)(nil)

func NewSystemparameterResource() resource.Resource {
	return &SystemparameterResource{}
}

// SystemparameterResource defines the resource implementation.
type SystemparameterResource struct {
	client *service.NitroClient
}

func (r *SystemparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemparameter"
}

func (r *SystemparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemparameter resource")

	// Create API request body from the model
	systemparameter := systemparameterGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("systemparameter-config")

	tflog.Trace(ctx, "Created systemparameter resource")

	// Read the updated state back
	r.readSystemparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemparameter resource")

	r.readSystemparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SystemparameterResourceModel

	// Read Terraform prior state, plan and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemparameter resource")

	// Determine whether there is a real update and which attributes were
	// removed from config (so they should be unset back to NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Cliloglevel.Equal(state.Cliloglevel) {
		if config.Cliloglevel.IsNull() {
			attributesToUnset = append(attributesToUnset, "cliloglevel")
		} else {
			hasChange = true
		}
	}
	if !data.Doppler.Equal(state.Doppler) {
		if config.Doppler.IsNull() {
			attributesToUnset = append(attributesToUnset, "doppler")
		} else {
			hasChange = true
		}
	}
	if !data.Googleanalytics.Equal(state.Googleanalytics) {
		if config.Googleanalytics.IsNull() {
			attributesToUnset = append(attributesToUnset, "googleanalytics")
		} else {
			hasChange = true
		}
	}
	if !data.Natpcbforceflushlimit.Equal(state.Natpcbforceflushlimit) {
		if config.Natpcbforceflushlimit.IsNull() {
			attributesToUnset = append(attributesToUnset, "natpcbforceflushlimit")
		} else {
			hasChange = true
		}
	}
	if !data.Natpcbrstontimeout.Equal(state.Natpcbrstontimeout) {
		if config.Natpcbrstontimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "natpcbrstontimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Rbaonresponse.Equal(state.Rbaonresponse) {
		if config.Rbaonresponse.IsNull() {
			attributesToUnset = append(attributesToUnset, "rbaonresponse")
		} else {
			hasChange = true
		}
	}
	if !data.Reauthonauthparamchange.Equal(state.Reauthonauthparamchange) {
		if config.Reauthonauthparamchange.IsNull() {
			attributesToUnset = append(attributesToUnset, "reauthonauthparamchange")
		} else {
			hasChange = true
		}
	}
	if !data.Removesensitivefiles.Equal(state.Removesensitivefiles) {
		if config.Removesensitivefiles.IsNull() {
			attributesToUnset = append(attributesToUnset, "removesensitivefiles")
		} else {
			hasChange = true
		}
	}
	if !data.Restrictedtimeout.Equal(state.Restrictedtimeout) {
		if config.Restrictedtimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "restrictedtimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Timeout.Equal(state.Timeout) {
		if config.Timeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "timeout")
		} else {
			hasChange = true
		}
	}
	if !data.Totalauthtimeout.Equal(state.Totalauthtimeout) {
		if config.Totalauthtimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "totalauthtimeout")
		} else {
			hasChange = true
		}
	}
	// Any other attribute change (attributes without unset wiring) still needs
	// a normal update.
	if !data.Basicauth.Equal(state.Basicauth) ||
		!data.Daystoexpire.Equal(state.Daystoexpire) ||
		!data.Denylist.Equal(state.Denylist) ||
		!data.Denylistlogging.Equal(state.Denylistlogging) ||
		!data.Fipsusermode.Equal(state.Fipsusermode) ||
		!data.Forcepasswordchange.Equal(state.Forcepasswordchange) ||
		!data.Localauth.Equal(state.Localauth) ||
		!data.Maxclient.Equal(state.Maxclient) ||
		!data.Maxsessionperuser.Equal(state.Maxsessionperuser) ||
		!data.Minpasswordlen.Equal(state.Minpasswordlen) ||
		!data.Passwordhistorycontrol.Equal(state.Passwordhistorycontrol) ||
		!data.Promptstring.Equal(state.Promptstring) ||
		!data.Pwdhistorycount.Equal(state.Pwdhistorycount) ||
		!data.Strongpassword.Equal(state.Strongpassword) ||
		!data.Wafprotection.Equal(state.Wafprotection) ||
		!data.Warnpriorndays.Equal(state.Warnpriorndays) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		systemparameter := systemparameterGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Systemparameter.Type(), &systemparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated systemparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for systemparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Systemparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset systemparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readSystemparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemparameter resource")

	// For systemparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted systemparameter resource from state")
}

// Helper function to read systemparameter data from API
func (r *SystemparameterResource) readSystemparameterFromApi(ctx context.Context, data *SystemparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Systemparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemparameter, got error: %s", err))
		return
	}

	systemparameterSetAttrFromGet(ctx, data, getResponseData)

}
