package nsratecontrol

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
var _ resource.Resource = &NsratecontrolResource{}
var _ resource.ResourceWithConfigure = (*NsratecontrolResource)(nil)
var _ resource.ResourceWithImportState = (*NsratecontrolResource)(nil)

func NewNsratecontrolResource() resource.Resource {
	return &NsratecontrolResource{}
}

// NsratecontrolResource defines the resource implementation.
type NsratecontrolResource struct {
	client *service.NitroClient
}

func (r *NsratecontrolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsratecontrolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsratecontrol"
}

func (r *NsratecontrolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsratecontrolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsratecontrolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsratecontrol resource")

	nsratecontrol := nsratecontrolGetThePayloadFromtheConfig(ctx, &data)

	// Make API call. nsratecontrol is a singleton (unnamed) configuration resource,
	// so it is configured via UpdateUnnamedResource (matches SDK v2 behavior).
	err := r.client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsratecontrol, got error: %s", err))
		return
	}

	// Static ID for this singleton configuration resource
	data.Id = types.StringValue("nsratecontrol-config")

	tflog.Trace(ctx, "Created nsratecontrol resource")

	// Read the updated state back
	r.readNsratecontrolFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsratecontrolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsratecontrolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsratecontrol resource")

	r.readNsratecontrolFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsratecontrolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsratecontrolResourceModel

	// Read Terraform prior state, plan and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsratecontrol resource")

	// Preserve ID from prior state
	data.Id = types.StringValue("nsratecontrol-config")

	// Determine which attributes changed and which were removed from config
	// (so they should be unset back to the appliance defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Icmpthreshold.Equal(state.Icmpthreshold) {
		if config.Icmpthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "icmpthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Tcprstthreshold.Equal(state.Tcprstthreshold) {
		if config.Tcprstthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcprstthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Tcpthreshold.Equal(state.Tcpthreshold) {
		if config.Tcpthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "tcpthreshold")
		} else {
			hasChange = true
		}
	}
	if !data.Udpthreshold.Equal(state.Udpthreshold) {
		if config.Udpthreshold.IsNull() {
			attributesToUnset = append(attributesToUnset, "udpthreshold")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		nsratecontrol := nsratecontrolGetThePayloadFromtheConfig(ctx, &data)

		// Make API call (singleton resource -> UpdateUnnamedResource)
		err := r.client.UpdateUnnamedResource(service.Nsratecontrol.Type(), &nsratecontrol)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsratecontrol, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsratecontrol resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsratecontrol resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nsratecontrol.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsratecontrol attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNsratecontrolFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsratecontrolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsratecontrolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsratecontrol resource")

	// For nsratecontrol, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsratecontrol resource from state")
}

// Helper function to read nsratecontrol data from API
func (r *NsratecontrolResource) readNsratecontrolFromApi(ctx context.Context, data *NsratecontrolResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsratecontrol.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsratecontrol, got error: %s", err))
		return
	}

	nsratecontrolSetAttrFromGet(ctx, data, getResponseData)

}
