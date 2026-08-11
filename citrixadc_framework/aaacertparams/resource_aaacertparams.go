package aaacertparams

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
var _ resource.Resource = &AaacertparamsResource{}
var _ resource.ResourceWithConfigure = (*AaacertparamsResource)(nil)
var _ resource.ResourceWithImportState = (*AaacertparamsResource)(nil)

func NewAaacertparamsResource() resource.Resource {
	return &AaacertparamsResource{}
}

// AaacertparamsResource defines the resource implementation.
type AaacertparamsResource struct {
	client *service.NitroClient
}

func (r *AaacertparamsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaacertparamsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaacertparams"
}

func (r *AaacertparamsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaacertparamsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaacertparamsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaacertparams resource")

	aaacertparams := aaacertparamsGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// aaacertparams is an unnamed (singleton) resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Aaacertparams.Type(), &aaacertparams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaacertparams, got error: %s", err))
		return
	}

	// Generate stable ID for this singleton configuration resource
	data.Id = types.StringValue("aaacertparams-config")

	tflog.Trace(ctx, "Created aaacertparams resource")

	// Read the updated state back
	r.readAaacertparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaacertparamsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaacertparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaacertparams resource")

	r.readAaacertparamsFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaacertparamsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AaacertparamsResourceModel

	// Read Terraform prior state, plan and config into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaacertparams resource")

	// Determine which attributes changed, and which were removed from config so
	// they must be unset (reverted to their appliance defaults).
	hasChange := false
	attributesToUnset := []string{}
	if !data.Defaultauthenticationgroup.Equal(state.Defaultauthenticationgroup) {
		if config.Defaultauthenticationgroup.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "defaultauthenticationgroup")
		} else {
			hasChange = true
		}
	}
	if !data.Groupnamefield.Equal(state.Groupnamefield) {
		if config.Groupnamefield.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "groupnamefield")
		} else {
			hasChange = true
		}
	}
	if !data.Usernamefield.Equal(state.Usernamefield) {
		if config.Usernamefield.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "usernamefield")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model
		aaacertparams := aaacertparamsGetThePayloadFromtheConfig(ctx, &data)

		// Make API call
		// aaacertparams is an unnamed (singleton) resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Aaacertparams.Type(), &aaacertparams)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaacertparams, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated aaacertparams resource")
	} else {
		tflog.Debug(ctx, "No changes detected for aaacertparams resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. aaacertparams is a singleton, so the unset payload
	// carries no identifying key.
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Aaacertparams.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset aaacertparams attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	r.readAaacertparamsFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaacertparamsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaacertparamsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaacertparams resource")

	// For aaacertparams, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaacertparams resource from state")
}

// Helper function to read aaacertparams data from API
func (r *AaacertparamsResource) readAaacertparamsFromApi(ctx context.Context, data *AaacertparamsResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaacertparams.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaacertparams, got error: %s", err))
		return
	}

	aaacertparamsSetAttrFromGet(ctx, data, getResponseData)

}
