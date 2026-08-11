package subscriberparam

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
var _ resource.Resource = &SubscriberparamResource{}
var _ resource.ResourceWithConfigure = (*SubscriberparamResource)(nil)
var _ resource.ResourceWithImportState = (*SubscriberparamResource)(nil)

func NewSubscriberparamResource() resource.Resource {
	return &SubscriberparamResource{}
}

// SubscriberparamResource defines the resource implementation.
type SubscriberparamResource struct {
	client *service.NitroClient
}

func (r *SubscriberparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SubscriberparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberparam"
}

func (r *SubscriberparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscriberparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscriberparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating subscriberparam resource")

	// Create API request body from the model
	subscriberparam := subscriberparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Subscriberparam.Type(), &subscriberparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create subscriberparam, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created subscriberparam resource")

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("subscriberparam-config")

	// Read the updated state back
	if !r.readSubscriberparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "subscriberparam not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscriberparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading subscriberparam resource")

	found := r.readSubscriberparamFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *SubscriberparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SubscriberparamResourceModel

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

	tflog.Debug(ctx, "Updating subscriberparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Idleaction.Equal(state.Idleaction) {
		tflog.Debug(ctx, "idleaction has changed for subscriberparam")
		if config.Idleaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "idleaction")
		} else {
			hasChange = true
		}
	}
	if !data.Idlettl.Equal(state.Idlettl) {
		tflog.Debug(ctx, "idlettl has changed for subscriberparam")
		if config.Idlettl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "idlettl")
		} else {
			hasChange = true
		}
	}
	if !data.Interfacetype.Equal(state.Interfacetype) {
		tflog.Debug(ctx, "interfacetype has changed for subscriberparam")
		if config.Interfacetype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "interfacetype")
		} else {
			hasChange = true
		}
	}
	if !data.Ipv6prefixlookuplist.Equal(state.Ipv6prefixlookuplist) {
		tflog.Debug(ctx, "ipv6prefixlookuplist has changed for subscriberparam")
		hasChange = true
	}
	if !data.Keytype.Equal(state.Keytype) {
		tflog.Debug(ctx, "keytype has changed for subscriberparam")
		if config.Keytype.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "keytype")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		subscriberparam := subscriberparamGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Subscriberparam.Type(), &subscriberparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update subscriberparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated subscriberparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for subscriberparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Subscriberparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset subscriberparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readSubscriberparamFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "subscriberparam not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SubscriberparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting subscriberparam resource")

	// subscriberparam does not support a DELETE operation (singleton global configuration).
	// Mirror the SDK v2 behavior: simply remove the resource from Terraform state.
	tflog.Trace(ctx, "Deleted subscriberparam resource from state")
}

// Helper function to read subscriberparam data from API
func (r *SubscriberparamResource) readSubscriberparamFromApi(ctx context.Context, data *SubscriberparamResourceModel, diags *diag.Diagnostics) bool {
	// Case 1: Simple find without ID (singleton)
	getResponseData, err := r.client.FindResource(service.Subscriberparam.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read subscriberparam, got error: %s", err))
		return false
	}

	subscriberparamSetAttrFromGet(ctx, data, getResponseData)

	return true
}
