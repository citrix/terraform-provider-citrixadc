package nsdiameter

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
var _ resource.Resource = &NsdiameterResource{}
var _ resource.ResourceWithConfigure = (*NsdiameterResource)(nil)
var _ resource.ResourceWithImportState = (*NsdiameterResource)(nil)

func NewNsdiameterResource() resource.Resource {
	return &NsdiameterResource{}
}

// NsdiameterResource defines the resource implementation.
type NsdiameterResource struct {
	client *service.NitroClient
}

func (r *NsdiameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsdiameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsdiameter"
}

func (r *NsdiameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsdiameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsdiameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsdiameter resource")

	nsdiameter := nsdiameterGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource (mirrors SDK v2 create).
	err := r.client.UpdateUnnamedResource(service.Nsdiameter.Type(), &nsdiameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsdiameter, got error: %s", err))
		return
	}

	// Static ID for this singleton configuration resource
	data.Id = types.StringValue("nsdiameter-config")

	tflog.Trace(ctx, "Created nsdiameter resource")

	// Read the updated state back
	r.readNsdiameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsdiameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsdiameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsdiameter resource")

	r.readNsdiameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsdiameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsdiameterResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsdiameter resource")

	// Preserve the singleton ID across the update.
	data.Id = types.StringValue("nsdiameter-config")

	// Detect changes and, for unset-eligible attributes removed from config,
	// collect them to unset (revert to NITRO defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Identity.Equal(state.Identity) {
		hasChange = true
	}
	if !data.Ownernode.Equal(state.Ownernode) {
		hasChange = true
	}
	if !data.Realm.Equal(state.Realm) {
		hasChange = true
	}
	if !data.Serverclosepropagation.Equal(state.Serverclosepropagation) {
		if config.Serverclosepropagation.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverclosepropagation")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		nsdiameter := nsdiameterGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource (mirrors SDK v2 update).
		err := r.client.UpdateUnnamedResource(service.Nsdiameter.Type(), &nsdiameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsdiameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated nsdiameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsdiameter resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Nsdiameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsdiameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readNsdiameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsdiameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsdiameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsdiameter resource")

	// For nsdiameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nsdiameter resource from state")
}

// Helper function to read nsdiameter data from API
func (r *NsdiameterResource) readNsdiameterFromApi(ctx context.Context, data *NsdiameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nsdiameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsdiameter, got error: %s", err))
		return
	}

	nsdiameterSetAttrFromGet(ctx, data, getResponseData)

}
