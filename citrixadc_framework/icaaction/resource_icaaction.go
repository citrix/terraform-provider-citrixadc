package icaaction

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
var _ resource.Resource = &IcaactionResource{}
var _ resource.ResourceWithConfigure = (*IcaactionResource)(nil)
var _ resource.ResourceWithImportState = (*IcaactionResource)(nil)

func NewIcaactionResource() resource.Resource {
	return &IcaactionResource{}
}

// IcaactionResource defines the resource implementation.
type IcaactionResource struct {
	client *service.NitroClient
}

func (r *IcaactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcaactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icaaction"
}

func (r *IcaactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcaactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcaactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icaaction resource")
	// Get payload from plan
	icaaction := icaactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	icaactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Icaaction.Type(), icaactionName, &icaaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created icaaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(icaactionName)

	// Read the updated state back
	if !r.readIcaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icaaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icaaction resource")

	found := r.readIcaactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *IcaactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state IcaactionResourceModel

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

	tflog.Debug(ctx, "Updating icaaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Accessprofilename.Equal(state.Accessprofilename) {
		tflog.Debug(ctx, "accessprofilename has changed for icaaction")
		if config.Accessprofilename.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "accessprofilename")
		} else {
			hasChange = true
		}
	}
	if !data.Latencyprofilename.Equal(state.Latencyprofilename) {
		tflog.Debug(ctx, "latencyprofilename has changed for icaaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		icaaction := icaactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// NITRO update for icaaction is a PUT to /config/icaaction (name carried in the
		// payload, not the URL) - use UpdateUnnamedResource (matches SDK v2 behavior).
		err := r.client.UpdateUnnamedResource(service.Icaaction.Type(), &icaaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icaaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated icaaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for icaaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Icaaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset icaaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readIcaactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "icaaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcaactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icaaction resource")
	// Named resource - delete using DeleteResource
	icaactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Icaaction.Type(), icaactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete icaaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted icaaction resource")
}

// Helper function to read icaaction data from API
func (r *IcaactionResource) readIcaactionFromApi(ctx context.Context, data *IcaactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	icaactionName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Icaaction.Type(), icaactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icaaction, got error: %s", err))
		return false
	}

	icaactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
