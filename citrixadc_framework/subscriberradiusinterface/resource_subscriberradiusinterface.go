package subscriberradiusinterface

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
var _ resource.Resource = &SubscriberradiusinterfaceResource{}
var _ resource.ResourceWithConfigure = (*SubscriberradiusinterfaceResource)(nil)
var _ resource.ResourceWithImportState = (*SubscriberradiusinterfaceResource)(nil)

func NewSubscriberradiusinterfaceResource() resource.Resource {
	return &SubscriberradiusinterfaceResource{}
}

// SubscriberradiusinterfaceResource defines the resource implementation.
type SubscriberradiusinterfaceResource struct {
	client *service.NitroClient
}

func (r *SubscriberradiusinterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SubscriberradiusinterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberradiusinterface"
}

func (r *SubscriberradiusinterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SubscriberradiusinterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SubscriberradiusinterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating subscriberradiusinterface resource")

	// Singleton resource - build payload and push with UpdateUnnamedResource
	subscriberradiusinterface := subscriberradiusinterfaceGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Subscriberradiusinterface.Type(), &subscriberradiusinterface)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create subscriberradiusinterface, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("subscriberradiusinterface-config")

	tflog.Trace(ctx, "Created subscriberradiusinterface resource")

	// Read the updated state back
	r.readSubscriberradiusinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberradiusinterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SubscriberradiusinterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading subscriberradiusinterface resource")

	r.readSubscriberradiusinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberradiusinterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state SubscriberradiusinterfaceResourceModel

	// Read Terraform prior state to preserve ID and detect changes
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

	tflog.Debug(ctx, "Updating subscriberradiusinterface resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Listeningservice.Equal(state.Listeningservice) {
		tflog.Debug(ctx, "listeningservice has changed for subscriberradiusinterface")
		hasChange = true
	}
	if !data.Radiusinterimasstart.Equal(state.Radiusinterimasstart) {
		tflog.Debug(ctx, "radiusinterimasstart has changed for subscriberradiusinterface")
		if config.Radiusinterimasstart.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "radiusinterimasstart")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Singleton resource - build payload and push with UpdateUnnamedResource
		subscriberradiusinterface := subscriberradiusinterfaceGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		err := r.client.UpdateUnnamedResource(service.Subscriberradiusinterface.Type(), &subscriberradiusinterface)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update subscriberradiusinterface, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No changes detected for subscriberradiusinterface resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Subscriberradiusinterface.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset subscriberradiusinterface attributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated subscriberradiusinterface resource")

	// Read the updated state back
	r.readSubscriberradiusinterfaceFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SubscriberradiusinterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SubscriberradiusinterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting subscriberradiusinterface resource")

	// For subscriberradiusinterface, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted subscriberradiusinterface resource from state")
}

// Helper function to read subscriberradiusinterface data from API
func (r *SubscriberradiusinterfaceResource) readSubscriberradiusinterfaceFromApi(ctx context.Context, data *SubscriberradiusinterfaceResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Subscriberradiusinterface.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read subscriberradiusinterface, got error: %s", err))
		return
	}

	subscriberradiusinterfaceSetAttrFromGet(ctx, data, getResponseData)

}
