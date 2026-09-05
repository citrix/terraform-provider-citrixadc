package sslzerotouchparam

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
var _ resource.Resource = &SslzerotouchparamResource{}
var _ resource.ResourceWithConfigure = (*SslzerotouchparamResource)(nil)
var _ resource.ResourceWithImportState = (*SslzerotouchparamResource)(nil)

func NewSslzerotouchparamResource() resource.Resource {
	return &SslzerotouchparamResource{}
}

// SslzerotouchparamResource defines the resource implementation.
type SslzerotouchparamResource struct {
	client *service.NitroClient
}

func (r *SslzerotouchparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SslzerotouchparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslzerotouchparam"
}

func (r *SslzerotouchparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SslzerotouchparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SslzerotouchparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating sslzerotouchparam resource")
	sslzerotouchparam := sslzerotouchparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Sslzerotouchparam.Type(), &sslzerotouchparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create sslzerotouchparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created sslzerotouchparam resource")

	// Read the updated state back (also sets the ID)
	if !r.readSslzerotouchparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslzerotouchparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslzerotouchparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SslzerotouchparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading sslzerotouchparam resource")

	found := r.readSslzerotouchparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SslzerotouchparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SslzerotouchparamResourceModel

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

	tflog.Debug(ctx, "Updating sslzerotouchparam resource")

	// Determine which attributes changed. Attributes removed from config are
	// unset (reverted to their NITRO defaults); attributes still present but
	// changed are pushed via update. An empty singleton PUT is rejected by
	// NITRO, so the update is guarded by hasChange.
	hasChange := false
	attributesToUnset := []string{}

	if !data.Ocspcachetimeout.Equal(state.Ocspcachetimeout) {
		if config.Ocspcachetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspcachetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspbatchingdepth.Equal(state.Ocspbatchingdepth) {
		if config.Ocspbatchingdepth.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspbatchingdepth")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspbatchingdelay.Equal(state.Ocspbatchingdelay) {
		if config.Ocspbatchingdelay.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspbatchingdelay")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspresptimeout.Equal(state.Ocspresptimeout) {
		if config.Ocspresptimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspresptimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspurlresolvetimeout.Equal(state.Ocspurlresolvetimeout) {
		if config.Ocspurlresolvetimeout.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspurlresolvetimeout")
		} else {
			hasChange = true
		}
	}
	if !data.Ocsptrustresponder.Equal(state.Ocsptrustresponder) {
		if config.Ocsptrustresponder.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocsptrustresponder")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspproducedattimeskew.Equal(state.Ocspproducedattimeskew) {
		if config.Ocspproducedattimeskew.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspproducedattimeskew")
		} else {
			hasChange = true
		}
	}
	if !data.Ocspusenonce.Equal(state.Ocspusenonce) {
		if config.Ocspusenonce.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocspusenonce")
		} else {
			hasChange = true
		}
	}
	if !data.Ocsphttpmethod.Equal(state.Ocsphttpmethod) {
		if config.Ocsphttpmethod.IsNull() {
			attributesToUnset = append(attributesToUnset, "ocsphttpmethod")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		sslzerotouchparam := sslzerotouchparamGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Sslzerotouchparam.Type(), &sslzerotouchparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update sslzerotouchparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated sslzerotouchparam resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for sslzerotouchparam resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Sslzerotouchparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset sslzerotouchparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSslzerotouchparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "sslzerotouchparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SslzerotouchparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SslzerotouchparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting sslzerotouchparam resource")

	// sslzerotouchparam is a global configuration singleton and does not support
	// a DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed sslzerotouchparam from Terraform state")
}

// Helper function to read sslzerotouchparam data from API
func (r *SslzerotouchparamResource) readSslzerotouchparamFromApi(ctx context.Context, data *SslzerotouchparamResourceModel, diags *diag.Diagnostics) bool {
	// Case 1: Simple find without ID (singleton)
	getResponseData, err := r.client.FindResource(service.Sslzerotouchparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read sslzerotouchparam, got error: %s", err))
		return false
	}

	sslzerotouchparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
