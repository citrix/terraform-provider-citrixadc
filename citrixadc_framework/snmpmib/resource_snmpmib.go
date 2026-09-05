package snmpmib

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
var _ resource.Resource = &SnmpmibResource{}
var _ resource.ResourceWithConfigure = (*SnmpmibResource)(nil)
var _ resource.ResourceWithImportState = (*SnmpmibResource)(nil)

func NewSnmpmibResource() resource.Resource {
	return &SnmpmibResource{}
}

// SnmpmibResource defines the resource implementation.
type SnmpmibResource struct {
	client *service.NitroClient
}

func (r *SnmpmibResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmpmibResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmpmib"
}

func (r *SnmpmibResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmpmibResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmpmibResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmpmib resource")

	// Create API request body from the plan
	snmpmib := snmpmibGetThePayloadFromtheConfig(ctx, &data)

	// snmpmib is a singleton (unnamed) resource - push config with UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmpmib, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created snmpmib resource")

	// Read the updated state back (also sets data.Id)
	r.readSnmpmibFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmibResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmpmibResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmpmib resource")

	r.readSnmpmibFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmibResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SnmpmibResourceModel

	// Read Terraform prior state, plan and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmpmib resource")

	// Detect changes and collect attributes removed from config so they can be
	// reverted to their ADC defaults via a single ?action=unset call.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Contact.Equal(state.Contact) {
		tflog.Debug(ctx, "contact has changed for snmpmib")
		if config.Contact.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "contact")
		} else {
			hasChange = true
		}
	}
	if !data.Customid.Equal(state.Customid) {
		tflog.Debug(ctx, "customid has changed for snmpmib")
		if config.Customid.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "customid")
		} else {
			hasChange = true
		}
	}
	if !data.Location.Equal(state.Location) {
		tflog.Debug(ctx, "location has changed for snmpmib")
		if config.Location.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "location")
		} else {
			hasChange = true
		}
	}
	if !data.Name.Equal(state.Name) {
		tflog.Debug(ctx, "name has changed for snmpmib")
		if config.Name.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "name")
		} else {
			hasChange = true
		}
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		tflog.Debug(ctx, "ownernode has changed for snmpmib")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		snmpmib := snmpmibGetThePayloadFromtheConfig(ctx, &data)

		// snmpmib is a singleton (unnamed) resource - push config with UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Snmpmib.Type(), &snmpmib)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmpmib, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated snmpmib resource")
	} else {
		tflog.Debug(ctx, "No changes detected for snmpmib resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. snmpmib is a singleton, so no identifying key is needed.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Snmpmib.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset snmpmib attributes, got error: %s", err))
		return
	}

	// Read the updated state back (also sets data.Id)
	r.readSnmpmibFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmpmibResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmpmibResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmpmib resource")

	// For snmpmib, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmpmib resource from state")
}

// Helper function to read snmpmib data from API
func (r *SnmpmibResource) readSnmpmibFromApi(ctx context.Context, data *SnmpmibResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmpmib.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmpmib, got error: %s", err))
		return
	}

	snmpmibSetAttrFromGet(ctx, data, getResponseData)

}
