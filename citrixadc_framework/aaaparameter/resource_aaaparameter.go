package aaaparameter

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
var _ resource.Resource = &AaaparameterResource{}
var _ resource.ResourceWithConfigure = (*AaaparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AaaparameterResource)(nil)

func NewAaaparameterResource() resource.Resource {
	return &AaaparameterResource{}
}

// AaaparameterResource defines the resource implementation.
type AaaparameterResource struct {
	client *service.NitroClient
}

func (r *AaaparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaaparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaaparameter"
}

func (r *AaaparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaaparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaaparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaaparameter resource")
	aaaparameter := aaaparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaparameter.Type(), &aaaparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaaparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created aaaparameter resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue("aaaparameter-config")

	// Read the updated state back
	r.readAaaparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaaparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaaparameter resource")

	r.readAaaparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaaparameterResourceModel

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

	tflog.Debug(ctx, "Updating aaaparameter resource")

	// Determine which attributes were removed from config so they can be unset
	// (reverted to their NITRO defaults) on the appliance.
	attributesToUnset := []string{}
	if !data.Aaadloglevel.Equal(state.Aaadloglevel) && config.Aaadloglevel.IsNull() {
		attributesToUnset = append(attributesToUnset, "aaadloglevel")
	}
	if !data.Apitokencache.Equal(state.Apitokencache) && config.Apitokencache.IsNull() {
		attributesToUnset = append(attributesToUnset, "apitokencache")
	}
	if !data.Classicendpoints.Equal(state.Classicendpoints) && config.Classicendpoints.IsNull() {
		attributesToUnset = append(attributesToUnset, "classicendpoints")
	}
	if !data.Defaultcspheader.Equal(state.Defaultcspheader) && config.Defaultcspheader.IsNull() {
		attributesToUnset = append(attributesToUnset, "defaultcspheader")
	}
	if !data.Enablesessionstickiness.Equal(state.Enablesessionstickiness) && config.Enablesessionstickiness.IsNull() {
		attributesToUnset = append(attributesToUnset, "enablesessionstickiness")
	}
	if !data.Enhancedepa.Equal(state.Enhancedepa) && config.Enhancedepa.IsNull() {
		attributesToUnset = append(attributesToUnset, "enhancedepa")
	}
	if !data.Httponlycookie.Equal(state.Httponlycookie) && config.Httponlycookie.IsNull() {
		attributesToUnset = append(attributesToUnset, "httponlycookie")
	}
	if !data.Loginencryption.Equal(state.Loginencryption) && config.Loginencryption.IsNull() {
		attributesToUnset = append(attributesToUnset, "loginencryption")
	}
	if !data.Maxkbquestions.Equal(state.Maxkbquestions) && config.Maxkbquestions.IsNull() {
		attributesToUnset = append(attributesToUnset, "maxkbquestions")
	}
	if !data.Persistentloginattempts.Equal(state.Persistentloginattempts) && config.Persistentloginattempts.IsNull() {
		attributesToUnset = append(attributesToUnset, "persistentloginattempts")
	}
	if !data.Securityinsights.Equal(state.Securityinsights) && config.Securityinsights.IsNull() {
		attributesToUnset = append(attributesToUnset, "securityinsights")
	}
	if !data.Webviewendpoints.Equal(state.Webviewendpoints) && config.Webviewendpoints.IsNull() {
		attributesToUnset = append(attributesToUnset, "webviewendpoints")
	}

	// Create API request body from the model
	aaaparameter := aaaparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Singleton resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaaparameter.Type(), &aaaparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaaparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated aaaparameter resource")

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. Singleton resource -> empty id payload.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Aaaparameter.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset aaaparameter attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readAaaparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaaparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaaparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaaparameter resource")

	// Singleton resource - no delete operation on ADC, just remove from state
	tflog.Trace(ctx, "Removed aaaparameter from Terraform state")
}

// Helper function to read aaaparameter data from API
func (r *AaaparameterResource) readAaaparameterFromApi(ctx context.Context, data *AaaparameterResourceModel, diags *diag.Diagnostics) {

	// Case 1: Simple find without ID (singleton)
	getResponseData, err := r.client.FindResource(service.Aaaparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaaparameter, got error: %s", err))
		return
	}

	aaaparameterSetAttrFromGet(ctx, data, getResponseData)

}
