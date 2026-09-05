package smppparam

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
var _ resource.Resource = &SmppparamResource{}
var _ resource.ResourceWithConfigure = (*SmppparamResource)(nil)
var _ resource.ResourceWithImportState = (*SmppparamResource)(nil)

func NewSmppparamResource() resource.Resource {
	return &SmppparamResource{}
}

// SmppparamResource defines the resource implementation.
type SmppparamResource struct {
	client *service.NitroClient
}

func (r *SmppparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SmppparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_smppparam"
}

func (r *SmppparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SmppparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SmppparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating smppparam resource")

	// Create API request body from the plan
	smppparam := smppparamGetThePayloadFromtheConfig(ctx, &data)

	// smppparam is a singleton (unnamed) configuration resource
	err := r.client.UpdateUnnamedResource(service.Smppparam.Type(), &smppparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create smppparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("smppparam-config")

	tflog.Trace(ctx, "Created smppparam resource")

	// Read the updated state back
	r.readSmppparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SmppparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading smppparam resource")

	r.readSmppparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SmppparamResourceModel

	// Read Terraform prior state, plan, and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating smppparam resource")

	// Preserve ID across the update (singleton static ID)
	data.Id = types.StringValue("smppparam-config")

	// Determine which attributes changed and which were removed from config
	// (removed -> unset so the appliance reverts them to their defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Addrnpi.Equal(state.Addrnpi) {
		if config.Addrnpi.IsNull() {
			attributesToUnset = append(attributesToUnset, "addrnpi")
		} else {
			hasChange = true
		}
	}
	if !data.Addrrange.Equal(state.Addrrange) {
		if config.Addrrange.IsNull() {
			attributesToUnset = append(attributesToUnset, "addrrange")
		} else {
			hasChange = true
		}
	}
	if !data.Addrton.Equal(state.Addrton) {
		if config.Addrton.IsNull() {
			attributesToUnset = append(attributesToUnset, "addrton")
		} else {
			hasChange = true
		}
	}
	if !data.Clientmode.Equal(state.Clientmode) {
		if config.Clientmode.IsNull() {
			attributesToUnset = append(attributesToUnset, "clientmode")
		} else {
			hasChange = true
		}
	}
	if !data.Msgqueue.Equal(state.Msgqueue) {
		if config.Msgqueue.IsNull() {
			attributesToUnset = append(attributesToUnset, "msgqueue")
		} else {
			hasChange = true
		}
	}
	if !data.Msgqueuesize.Equal(state.Msgqueuesize) {
		if config.Msgqueuesize.IsNull() {
			attributesToUnset = append(attributesToUnset, "msgqueuesize")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the plan
		smppparam := smppparamGetThePayloadFromtheConfig(ctx, &data)

		// smppparam is a singleton (unnamed) configuration resource
		err := r.client.UpdateUnnamedResource(service.Smppparam.Type(), &smppparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update smppparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated smppparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for smppparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. smppparam is a singleton, so no id fields.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Smppparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset smppparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readSmppparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SmppparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SmppparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting smppparam resource")

	// For smppparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted smppparam resource from state")
}

// Helper function to read smppparam data from API
func (r *SmppparamResource) readSmppparamFromApi(ctx context.Context, data *SmppparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Smppparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read smppparam, got error: %s", err))
		return
	}

	smppparamSetAttrFromGet(ctx, data, getResponseData)

}
