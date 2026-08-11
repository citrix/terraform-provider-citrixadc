package inatparam

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
var _ resource.Resource = &InatparamResource{}
var _ resource.ResourceWithConfigure = (*InatparamResource)(nil)
var _ resource.ResourceWithImportState = (*InatparamResource)(nil)

func NewInatparamResource() resource.Resource {
	return &InatparamResource{}
}

// InatparamResource defines the resource implementation.
type InatparamResource struct {
	client *service.NitroClient
}

func (r *InatparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *InatparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inatparam"
}

func (r *InatparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InatparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InatparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating inatparam resource")

	inatparam := inatparamGetThePayloadFromtheConfig(ctx, &data)

	// inatparam has no NITRO ADD operation; it is configured with an
	// unnamed (PUT) update, matching the legacy SDK v2 behavior.
	err := r.client.UpdateUnnamedResource(service.Inatparam.Type(), &inatparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create inatparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created inatparam resource")

	// Read the updated state back
	if !r.readInatparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "inatparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InatparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InatparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading inatparam resource")

	found := r.readInatparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *InatparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state InatparamResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating inatparam resource")

	// Determine attributes that were removed from config so they can be unset
	// (reverted to their NITRO defaults) after the update. Only nat46v6prefix
	// supports the NITRO unset operation on inatparam; the nat46* toggle/mtu
	// attributes are rejected with "Invalid argument" by the appliance.
	attributesToUnset := []string{}
	if !data.Nat46v6prefix.Equal(state.Nat46v6prefix) {
		if config.Nat46v6prefix.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "nat46v6prefix")
		}
	}

	// Create API request body from the model
	inatparam := inatparamGetThePayloadFromtheConfig(ctx, &data)

	// inatparam is configured with an unnamed (PUT) update.
	err := r.client.UpdateUnnamedResource(service.Inatparam.Type(), &inatparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update inatparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated inatparam resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. inatparam is keyed on the traffic domain (td); the unset
	// applies to the default traffic domain (0) unless a non-zero td is set.
	unsetIdPayload := map[string]interface{}{}
	if !data.Td.IsNull() && !data.Td.IsUnknown() && data.Td.ValueInt64() != 0 {
		unsetIdPayload["td"] = int(data.Td.ValueInt64())
	}
	if err := utils.ExecuteUnset(r.client, service.Inatparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset inatparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readInatparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "inatparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InatparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// inatparam does not support a NITRO DELETE operation. Matching the legacy
	// SDK v2 behavior, deletion is a no-op: the resource is only removed from
	// Terraform state (done automatically by the framework once Delete returns).
	tflog.Debug(ctx, "Deleting inatparam resource (no-op; DELETE not supported by NITRO)")
}

// Helper function to read inatparam data from API
func (r *InatparamResource) readInatparamFromApi(ctx context.Context, data *InatparamResourceModel, diags *diag.Diagnostics) bool {
	// inatparam is keyed on the traffic domain (td). Default traffic domain is 0.
	tdName := fmt.Sprintf("%d", data.Td.ValueInt64())
	// On import, td is not yet populated; the ID carries the td value.
	if (data.Td.IsNull() || data.Td.IsUnknown()) && data.Id.ValueString() != "" {
		tdName = data.Id.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Inatparam.Type(), tdName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read inatparam, got error: %s", err))
		return false
	}

	inatparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
