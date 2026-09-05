package iptunnelparam

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
var _ resource.Resource = &IptunnelparamResource{}
var _ resource.ResourceWithConfigure = (*IptunnelparamResource)(nil)
var _ resource.ResourceWithImportState = (*IptunnelparamResource)(nil)

func NewIptunnelparamResource() resource.Resource {
	return &IptunnelparamResource{}
}

// IptunnelparamResource defines the resource implementation.
type IptunnelparamResource struct {
	client *service.NitroClient
}

func (r *IptunnelparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IptunnelparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iptunnelparam"
}

func (r *IptunnelparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IptunnelparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IptunnelparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating iptunnelparam resource")

	iptunnelparam := iptunnelparamGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - push configuration with UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create iptunnelparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("iptunnelparam-config")

	tflog.Trace(ctx, "Created iptunnelparam resource")

	// Read the updated state back
	r.readIptunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IptunnelparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IptunnelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading iptunnelparam resource")

	r.readIptunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IptunnelparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IptunnelparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating iptunnelparam resource")

	// Detect attributes that were removed from config so they can be unset
	// (reverted to their NITRO defaults) after the update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Dropfrag.Equal(state.Dropfrag) {
		if config.Dropfrag.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropfrag")
		} else {
			hasChange = true
		}
	}
	if !data.Dropfragcputhreshold.Equal(state.Dropfragcputhreshold) {
		if config.Dropfragcputhreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "dropfragcputhreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Enablestrictrx.Equal(state.Enablestrictrx) {
		if config.Enablestrictrx.IsNull() {
			attributesToUnset = append(attributesToUnset, "enablestrictrx")
		} else {
			hasChange = true
		}
	}
	if !data.Enablestricttx.Equal(state.Enablestricttx) {
		if config.Enablestricttx.IsNull() {
			attributesToUnset = append(attributesToUnset, "enablestricttx")
		} else {
			hasChange = true
		}
	}
	if !data.Srciproundrobin.Equal(state.Srciproundrobin) {
		if config.Srciproundrobin.IsNull() {
			attributesToUnset = append(attributesToUnset, "srciproundrobin")
		} else {
			hasChange = true
		}
	}
	if !data.Useclientsourceip.Equal(state.Useclientsourceip) {
		if config.Useclientsourceip.IsNull() {
			attributesToUnset = append(attributesToUnset, "useclientsourceip")
		} else {
			hasChange = true
		}
	}
	// Attributes without a documented NITRO default are not unset-eligible;
	// a change to them still requires an update.
	if !data.Mac.Equal(state.Mac) {
		hasChange = true
	}
	if !data.Srcip.Equal(state.Srcip) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		iptunnelparam := iptunnelparamGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - push configuration with UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Iptunnelparam.Type(), &iptunnelparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update iptunnelparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated iptunnelparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for iptunnelparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Iptunnelparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset iptunnelparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readIptunnelparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IptunnelparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IptunnelparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting iptunnelparam resource")

	// For iptunnelparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted iptunnelparam resource from state")
}

// Helper function to read iptunnelparam data from API
func (r *IptunnelparamResource) readIptunnelparamFromApi(ctx context.Context, data *IptunnelparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Iptunnelparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read iptunnelparam, got error: %s", err))
		return
	}

	iptunnelparamSetAttrFromGet(ctx, data, getResponseData)

}
