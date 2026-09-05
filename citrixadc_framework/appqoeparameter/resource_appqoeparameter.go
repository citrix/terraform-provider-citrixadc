package appqoeparameter

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
var _ resource.Resource = &AppqoeparameterResource{}
var _ resource.ResourceWithConfigure = (*AppqoeparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoeparameterResource)(nil)

func NewAppqoeparameterResource() resource.Resource {
	return &AppqoeparameterResource{}
}

// AppqoeparameterResource defines the resource implementation.
type AppqoeparameterResource struct {
	client *service.NitroClient
}

func (r *AppqoeparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoeparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoeparameter"
}

func (r *AppqoeparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoeparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoeparameter resource")

	appqoeparameter := appqoeparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Unnamed/singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoeparameter, got error: %s", err))
		return
	}

	// Set a fixed synthetic ID for this singleton configuration resource
	data.Id = types.StringValue("appqoeparameter-config")

	tflog.Trace(ctx, "Created appqoeparameter resource")

	// Read the updated state back
	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoeparameter resource")

	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppqoeparameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appqoeparameter resource")

	// Determine which attributes changed and which were removed from config (unset)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Avgwaitingclient.Equal(state.Avgwaitingclient) {
		if config.Avgwaitingclient.IsNull() {
			attributesToUnset = append(attributesToUnset, "avgwaitingclient")
		} else {
			hasChange = true
		}
	}
	if !data.Dosattackthresh.Equal(state.Dosattackthresh) {
		if config.Dosattackthresh.IsNull() {
			attributesToUnset = append(attributesToUnset, "dosattackthresh")
		} else {
			hasChange = true
		}
	}
	if !data.Maxaltrespbandwidth.Equal(state.Maxaltrespbandwidth) {
		if config.Maxaltrespbandwidth.IsNull() {
			attributesToUnset = append(attributesToUnset, "maxaltrespbandwidth")
		} else {
			hasChange = true
		}
	}
	if !data.Sessionlife.Equal(state.Sessionlife) {
		if config.Sessionlife.IsNull() {
			attributesToUnset = append(attributesToUnset, "sessionlife")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		appqoeparameter := appqoeparameterGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// Unnamed/singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appqoeparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appqoeparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appqoeparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts to defaults
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Appqoeparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset appqoeparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoeparameter resource")

	// For appqoeparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appqoeparameter resource from state")
}

// Helper function to read appqoeparameter data from API
func (r *AppqoeparameterResource) readAppqoeparameterFromApi(ctx context.Context, data *AppqoeparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appqoeparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoeparameter, got error: %s", err))
		return
	}

	appqoeparameterSetAttrFromGet(ctx, data, getResponseData)

}
