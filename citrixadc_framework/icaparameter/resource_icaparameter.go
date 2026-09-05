package icaparameter

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
var _ resource.Resource = &IcaparameterResource{}
var _ resource.ResourceWithConfigure = (*IcaparameterResource)(nil)
var _ resource.ResourceWithImportState = (*IcaparameterResource)(nil)

func NewIcaparameterResource() resource.Resource {
	return &IcaparameterResource{}
}

// IcaparameterResource defines the resource implementation.
type IcaparameterResource struct {
	client *service.NitroClient
}

func (r *IcaparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcaparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icaparameter"
}

func (r *IcaparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcaparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcaparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icaparameter resource")

	// Create API request body from the plan
	icaparameter := icaparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Icaparameter.Type(), &icaparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icaparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("icaparameter-config")

	tflog.Trace(ctx, "Created icaparameter resource")

	// Read the updated state back
	r.readIcaparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcaparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icaparameter resource")

	r.readIcaparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IcaparameterResourceModel

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

	tflog.Debug(ctx, "Updating icaparameter resource")

	// Detect changes and, for attributes removed from config, collect them to unset.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dfpersistence.Equal(state.Dfpersistence) {
		if config.Dfpersistence.IsNull() {
			attributesToUnset = append(attributesToUnset, "dfpersistence")
		} else {
			hasChange = true
		}
	}
	if !data.Edtlosstolerant.Equal(state.Edtlosstolerant) {
		hasChange = true
	}
	if !data.Edtpmtuddf.Equal(state.Edtpmtuddf) {
		if config.Edtpmtuddf.IsNull() {
			attributesToUnset = append(attributesToUnset, "edtpmtuddf")
		} else {
			hasChange = true
		}
	}
	if !data.Edtpmtuddftimeout.Equal(state.Edtpmtuddftimeout) {
		if config.Edtpmtuddftimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "edtpmtuddftimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Edtpmtudrediscovery.Equal(state.Edtpmtudrediscovery) {
		if config.Edtpmtudrediscovery.IsNull() {
			attributesToUnset = append(attributesToUnset, "edtpmtudrediscovery")
		} else {
			hasChange = true
		}
	}
	if !data.Enablesronhafailover.Equal(state.Enablesronhafailover) {
		if config.Enablesronhafailover.IsNull() {
			attributesToUnset = append(attributesToUnset, "enablesronhafailover")
		} else {
			hasChange = true
		}
	}
	if !data.Hdxinsightnonnsap.Equal(state.Hdxinsightnonnsap) {
		if config.Hdxinsightnonnsap.IsNull() {
			attributesToUnset = append(attributesToUnset, "hdxinsightnonnsap")
		} else {
			hasChange = true
		}
	}
	if !data.Insightonlytodirector.Equal(state.Insightonlytodirector) {
		if config.Insightonlytodirector.IsNull() {
			attributesToUnset = append(attributesToUnset, "insightonlytodirector")
		} else {
			hasChange = true
		}
	}
	if !data.L7latencyfrequency.Equal(state.L7latencyfrequency) {
		if config.L7latencyfrequency.IsNull() {
			attributesToUnset = append(attributesToUnset, "l7latencyfrequency")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the plan
		icaparameter := icaparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Icaparameter.Type(), &icaparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icaparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated icaparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for icaparameter resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Icaparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset icaparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readIcaparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcaparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icaparameter resource")

	// For icaparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted icaparameter resource from state")
}

// Helper function to read icaparameter data from API
func (r *IcaparameterResource) readIcaparameterFromApi(ctx context.Context, data *IcaparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Icaparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icaparameter, got error: %s", err))
		return
	}

	icaparameterSetAttrFromGet(ctx, data, getResponseData)

}
