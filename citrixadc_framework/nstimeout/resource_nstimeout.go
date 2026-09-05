package nstimeout

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
var _ resource.Resource = &NstimeoutResource{}
var _ resource.ResourceWithConfigure = (*NstimeoutResource)(nil)
var _ resource.ResourceWithImportState = (*NstimeoutResource)(nil)

func NewNstimeoutResource() resource.Resource {
	return &NstimeoutResource{}
}

// NstimeoutResource defines the resource implementation.
type NstimeoutResource struct {
	client *service.NitroClient
}

func (r *NstimeoutResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NstimeoutResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nstimeout"
}

func (r *NstimeoutResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NstimeoutResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config NstimeoutResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to build the payload only from user-configured values (matches SDK v2 GetRawConfig)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nstimeout resource")

	// Build the payload from the config (unconfigured attributes are null and skipped)
	nstimeout := nstimeoutGetThePayloadFromtheConfig(ctx, &config)

	// Make API call - nstimeout is a singleton, use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nstimeout, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nstimeout-config")

	tflog.Trace(ctx, "Created nstimeout resource")

	// Read the updated state back
	r.readNstimeoutFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimeoutResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NstimeoutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nstimeout resource")

	r.readNstimeoutFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimeoutResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NstimeoutResourceModel

	// Read Terraform prior state
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to build the payload only from user-configured values (matches SDK v2 GetRawConfig)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nstimeout resource")

	// Determine attributes removed from config so they can be unset (reverted to
	// NITRO defaults) after the update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Anyclient.Equal(state.Anyclient) {
		if config.Anyclient.IsNull() {
			attributesToUnset = append(attributesToUnset, "anyclient")
		} else {
			hasChange = true
		}
	}
	if !data.Httpclient.Equal(state.Httpclient) {
		if config.Httpclient.IsNull() {
			attributesToUnset = append(attributesToUnset, "httpclient")
		} else {
			hasChange = true
		}
	}
	if !data.Reducedrsttimeout.Equal(state.Reducedrsttimeout) {
		if config.Reducedrsttimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "reducedrsttimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Server.Equal(state.Server) {
		if config.Server.IsNull() {
			attributesToUnset = append(attributesToUnset, "server")
		} else {
			hasChange = true
		}
	}
	if !data.Zombie.Equal(state.Zombie) {
		if config.Zombie.IsNull() {
			attributesToUnset = append(attributesToUnset, "zombie")
		} else {
			hasChange = true
		}
	}
	// Non-unset-wired attributes: a configured change requires an update call.
	// (Only when present in config -- otherwise a Computed attr planned as
	// "known after apply" would spuriously force an empty-payload update.)
	if (!config.Anyserver.IsNull() && !data.Anyserver.Equal(state.Anyserver)) ||
		(!config.Anytcpclient.IsNull() && !data.Anytcpclient.Equal(state.Anytcpclient)) ||
		(!config.Anytcpserver.IsNull() && !data.Anytcpserver.Equal(state.Anytcpserver)) ||
		(!config.Client.IsNull() && !data.Client.Equal(state.Client)) ||
		(!config.Halfclose.IsNull() && !data.Halfclose.Equal(state.Halfclose)) ||
		(!config.Httpserver.IsNull() && !data.Httpserver.Equal(state.Httpserver)) ||
		(!config.Newconnidletimeout.IsNull() && !data.Newconnidletimeout.Equal(state.Newconnidletimeout)) ||
		(!config.Nontcpzombie.IsNull() && !data.Nontcpzombie.Equal(state.Nontcpzombie)) ||
		(!config.Reducedfintimeout.IsNull() && !data.Reducedfintimeout.Equal(state.Reducedfintimeout)) ||
		(!config.Tcpclient.IsNull() && !data.Tcpclient.Equal(state.Tcpclient)) ||
		(!config.Tcpserver.IsNull() && !data.Tcpserver.Equal(state.Tcpserver)) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the config (unconfigured attributes are null and skipped)
		nstimeout := nstimeoutGetThePayloadFromtheConfig(ctx, &config)

		// Make API call - nstimeout is a singleton, use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Nstimeout.Type(), &nstimeout)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nstimeout, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nstimeout resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nstimeout resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. nstimeout is a singleton, so the id payload is empty.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nstimeout.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nstimeout attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNstimeoutFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NstimeoutResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NstimeoutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nstimeout resource")

	// For nstimeout, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nstimeout resource from state")
}

// Helper function to read nstimeout data from API
func (r *NstimeoutResource) readNstimeoutFromApi(ctx context.Context, data *NstimeoutResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nstimeout.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nstimeout, got error: %s", err))
		return
	}

	nstimeoutSetAttrFromGet(ctx, data, getResponseData)

}
