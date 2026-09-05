package cacheparameter

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
var _ resource.Resource = &CacheparameterResource{}
var _ resource.ResourceWithConfigure = (*CacheparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CacheparameterResource)(nil)

func NewCacheparameterResource() resource.Resource {
	return &CacheparameterResource{}
}

// CacheparameterResource defines the resource implementation.
type CacheparameterResource struct {
	client *service.NitroClient
}

func (r *CacheparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CacheparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheparameter"
}

func (r *CacheparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CacheparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CacheparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cacheparameter resource")

	// Create API request body from the plan
	cacheparameter := cacheparameterGetThePayloadFromtheConfig(ctx, &data)

	// Singleton resource - use UpdateUnnamedResource to configure the ADC
	err := r.client.UpdateUnnamedResource(service.Cacheparameter.Type(), &cacheparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cacheparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cacheparameter-config")

	tflog.Trace(ctx, "Created cacheparameter resource")

	// Read the updated state back
	r.readCacheparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CacheparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cacheparameter resource")

	r.readCacheparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CacheparameterResourceModel

	// Read Terraform prior state, plan and config into the models
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cacheparameter resource")

	// Determine changed attributes and which were removed from config (-> unset).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Cacheevictionpolicy.Equal(state.Cacheevictionpolicy) {
		if config.Cacheevictionpolicy.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cacheevictionpolicy")
		} else {
			hasChange = true
		}
	}
	if !data.Enablehaobjpersist.Equal(state.Enablehaobjpersist) {
		hasChange = true
	}
	if !data.Maxpostlen.Equal(state.Maxpostlen) {
		if config.Maxpostlen.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "maxpostlen")
		} else {
			hasChange = true
		}
	}
	// Remaining attributes have no documented unset default; any change is a plain update.
	if !data.Enablebypass.Equal(state.Enablebypass) {
		hasChange = true
	}
	if !data.Memlimit.Equal(state.Memlimit) {
		hasChange = true
	}
	if !data.Prefetchmaxpending.Equal(state.Prefetchmaxpending) {
		hasChange = true
	}
	if !data.Undefaction.Equal(state.Undefaction) {
		hasChange = true
	}
	if !data.Verifyusing.Equal(state.Verifyusing) {
		hasChange = true
	}
	if !data.Via.Equal(state.Via) {
		hasChange = true
	}

	if hasChange {
		// Create API request body from the plan
		cacheparameter := cacheparameterGetThePayloadFromtheConfig(ctx, &data)

		// Singleton resource - use UpdateUnnamedResource to configure the ADC
		err := r.client.UpdateUnnamedResource(service.Cacheparameter.Type(), &cacheparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cacheparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated cacheparameter resource")
	} else {
		tflog.Debug(ctx, "No changes detected for cacheparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Cacheparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset cacheparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readCacheparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CacheparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CacheparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cacheparameter resource")

	// For cacheparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cacheparameter resource from state")
}

// Helper function to read cacheparameter data from API
func (r *CacheparameterResource) readCacheparameterFromApi(ctx context.Context, data *CacheparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cacheparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cacheparameter, got error: %s", err))
		return
	}

	cacheparameterSetAttrFromGet(ctx, data, getResponseData)

}
