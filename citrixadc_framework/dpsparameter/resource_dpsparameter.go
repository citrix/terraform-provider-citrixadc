package dpsparameter

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
var _ resource.Resource = &DpsparameterResource{}
var _ resource.ResourceWithConfigure = (*DpsparameterResource)(nil)
var _ resource.ResourceWithImportState = (*DpsparameterResource)(nil)

func NewDpsparameterResource() resource.Resource {
	return &DpsparameterResource{}
}

// DpsparameterResource defines the resource implementation.
type DpsparameterResource struct {
	client *service.NitroClient
}

func (r *DpsparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DpsparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dpsparameter"
}

func (r *DpsparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DpsparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DpsparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dpsparameter resource")

	// Create API request body from the model (untyped payload)
	dpsparameter := dpsparameterGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Unnamed singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Dpsparameter.Type(), &dpsparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dpsparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created dpsparameter resource")

	// Read the updated state back (also sets the ID)
	if !r.readDpsparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dpsparameter not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DpsparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DpsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dpsparameter resource")

	found := r.readDpsparameterFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *DpsparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state DpsparameterResourceModel

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

	tflog.Debug(ctx, "Updating dpsparameter resource")

	// Determine which attributes changed. All three writable attributes
	// (customerid, deployment, serviceurl) support the NITRO unset operation, so
	// a value removed from config is unset rather than pushed via update.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Customerid.Equal(state.Customerid) {
		tflog.Debug(ctx, "customerid has changed for dpsparameter")
		if config.Customerid.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "customerid")
		} else {
			hasChange = true
		}
	}
	if !data.Deployment.Equal(state.Deployment) {
		tflog.Debug(ctx, "deployment has changed for dpsparameter")
		if config.Deployment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "deployment")
		} else {
			hasChange = true
		}
	}
	if !data.Serviceurl.Equal(state.Serviceurl) {
		tflog.Debug(ctx, "serviceurl has changed for dpsparameter")
		if config.Serviceurl.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serviceurl")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (untyped payload)
		dpsparameter := dpsparameterGetThePayloadFromthePlan(ctx, &data)

		// Make API call
		// Unnamed singleton resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Dpsparameter.Type(), &dpsparameter)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dpsparameter, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated dpsparameter resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for dpsparameter resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Dpsparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset dpsparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readDpsparameterFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "dpsparameter not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DpsparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DpsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dpsparameter resource")

	// dpsparameter is a global configuration singleton and does not support a
	// DELETE operation. We simply remove it from Terraform state.
	tflog.Trace(ctx, "Removed dpsparameter from Terraform state")
}

// Helper function to read dpsparameter data from API
func (r *DpsparameterResource) readDpsparameterFromApi(ctx context.Context, data *DpsparameterResourceModel, diags *diag.Diagnostics) bool {
	getResponseData, err := r.client.FindResource(service.Dpsparameter.Type(), "")
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dpsparameter, got error: %s", err))
		return false
	}

	dpsparameterSetAttrFromGet(ctx, data, getResponseData)

	return true
}
