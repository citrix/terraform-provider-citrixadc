package ipsecparameter

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
var _ resource.Resource = &IpsecparameterResource{}
var _ resource.ResourceWithConfigure = (*IpsecparameterResource)(nil)
var _ resource.ResourceWithImportState = (*IpsecparameterResource)(nil)

func NewIpsecparameterResource() resource.Resource {
	return &IpsecparameterResource{}
}

// IpsecparameterResource defines the resource implementation.
type IpsecparameterResource struct {
	client *service.NitroClient
}

func (r *IpsecparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IpsecparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecparameter"
}

func (r *IpsecparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IpsecparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IpsecparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ipsecparameter resource")

	// Create API request body from the model
	ipsecparameter := ipsecparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed (singleton) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ipsecparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("ipsecparameter-config")

	tflog.Trace(ctx, "Created ipsecparameter resource")

	// Read the updated state back
	r.readIpsecparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IpsecparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ipsecparameter resource")

	r.readIpsecparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IpsecparameterResourceModel

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

	tflog.Debug(ctx, "Updating ipsecparameter resource")

	// Determine attributes removed from config so they can be unset (reverted
	// to their NITRO defaults) after the update.
	hasChange := false
	attributesToUnset := []string{}
	// encalgo and hashalgo are list attributes without a schema Default, so they
	// stay sticky when removed from config (no plan diff) and cannot be reliably
	// unset here; they are intentionally excluded from the unset set.
	if !data.Encalgo.Equal(state.Encalgo) && !config.Encalgo.IsNull() {
		hasChange = true
	}
	if !data.Hashalgo.Equal(state.Hashalgo) && !config.Hashalgo.IsNull() {
		hasChange = true
	}
	if !data.Ikeretryinterval.Equal(state.Ikeretryinterval) {
		if config.Ikeretryinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "ikeretryinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Ikeversion.Equal(state.Ikeversion) {
		if config.Ikeversion.IsNull() {
			attributesToUnset = append(attributesToUnset, "ikeversion")
		} else {
			hasChange = true
		}
	}
	if !data.Lifetime.Equal(state.Lifetime) {
		if config.Lifetime.IsNull() {
			attributesToUnset = append(attributesToUnset, "lifetime")
		} else {
			hasChange = true
		}
	}
	if !data.Livenesscheckinterval.Equal(state.Livenesscheckinterval) {
		if config.Livenesscheckinterval.IsNull() {
			attributesToUnset = append(attributesToUnset, "livenesscheckinterval")
		} else {
			hasChange = true
		}
	}
	if !data.Perfectforwardsecrecy.Equal(state.Perfectforwardsecrecy) {
		if config.Perfectforwardsecrecy.IsNull() {
			attributesToUnset = append(attributesToUnset, "perfectforwardsecrecy")
		} else {
			hasChange = true
		}
	}
	if !data.Replaywindowsize.Equal(state.Replaywindowsize) {
		if config.Replaywindowsize.IsNull() {
			attributesToUnset = append(attributesToUnset, "replaywindowsize")
		} else {
			hasChange = true
		}
	}
	if !data.Retransmissiontime.Equal(state.Retransmissiontime) {
		if config.Retransmissiontime.IsNull() {
			attributesToUnset = append(attributesToUnset, "retransmissiontime")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		ipsecparameter := ipsecparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed (singleton) resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Ipsecparameter.Type(), &ipsecparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ipsecparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated ipsecparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for ipsecparameter resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Ipsecparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset ipsecparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readIpsecparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IpsecparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IpsecparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ipsecparameter resource")

	// For ipsecparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted ipsecparameter resource from state")
}

// Helper function to read ipsecparameter data from API
func (r *IpsecparameterResource) readIpsecparameterFromApi(ctx context.Context, data *IpsecparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Ipsecparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ipsecparameter, got error: %s", err))
		return
	}

	ipsecparameterSetAttrFromGet(ctx, data, getResponseData)

}
