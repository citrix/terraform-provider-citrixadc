package arpparam

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
var _ resource.Resource = &ArpparamResource{}
var _ resource.ResourceWithConfigure = (*ArpparamResource)(nil)
var _ resource.ResourceWithImportState = (*ArpparamResource)(nil)

func NewArpparamResource() resource.Resource {
	return &ArpparamResource{}
}

// ArpparamResource defines the resource implementation.
type ArpparamResource struct {
	client *service.NitroClient
}

func (r *ArpparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ArpparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_arpparam"
}

func (r *ArpparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ArpparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ArpparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating arpparam resource")

	arpparam := arpparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed/singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create arpparam, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("arpparam-config")

	tflog.Trace(ctx, "Created arpparam resource")

	// Read the updated state back
	r.readArpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArpparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ArpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading arpparam resource")

	r.readArpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArpparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ArpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating arpparam resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Spoofvalidation.Equal(state.Spoofvalidation) {
		tflog.Debug(ctx, "spoofvalidation has changed for arpparam")
		if config.Spoofvalidation.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "spoofvalidation")
		} else {
			hasChange = true
		}
	}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, "timeout has changed for arpparam")
		if config.Timeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timeout")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		arpparam := arpparamGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed/singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Arpparam.Type(), &arpparam)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update arpparam, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated arpparam resource")
	} else {
		tflog.Debug(ctx, "No changes detected for arpparam resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Arpparam.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset arpparam attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readArpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ArpparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ArpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting arpparam resource")

	// For arpparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted arpparam resource from state")
}

// Helper function to read arpparam data from API
func (r *ArpparamResource) readArpparamFromApi(ctx context.Context, data *ArpparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Arpparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read arpparam, got error: %s", err))
		return
	}

	arpparamSetAttrFromGet(ctx, data, getResponseData)

}
