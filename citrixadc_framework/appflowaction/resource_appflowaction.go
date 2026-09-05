package appflowaction

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
var _ resource.Resource = &AppflowactionResource{}
var _ resource.ResourceWithConfigure = (*AppflowactionResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowactionResource)(nil)

func NewAppflowactionResource() resource.Resource {
	return &AppflowactionResource{}
}

// AppflowactionResource defines the resource implementation.
type AppflowactionResource struct {
	client *service.NitroClient
}

func (r *AppflowactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowaction"
}

func (r *AppflowactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowaction resource")
	// Get payload from plan
	appflowaction := appflowactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	appflowactionName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Appflowaction.Type(), appflowactionName, &appflowaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appflowaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", appflowactionName))

	// Read the updated state back
	if !r.readAppflowactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appflowaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowaction resource")

	found := r.readAppflowactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppflowactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state AppflowactionResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appflowaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Botinsight.Equal(state.Botinsight) {
		tflog.Debug(ctx, "botinsight has changed for appflowaction")
		if config.Botinsight.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "botinsight")
		} else {
			hasChange = true
		}
	}
	if !data.Ciinsight.Equal(state.Ciinsight) {
		tflog.Debug(ctx, "ciinsight has changed for appflowaction")
		if config.Ciinsight.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ciinsight")
		} else {
			hasChange = true
		}
	}
	if !data.Clientsidemeasurements.Equal(state.Clientsidemeasurements) {
		tflog.Debug(ctx, "clientsidemeasurements has changed for appflowaction")
		if config.Clientsidemeasurements.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "clientsidemeasurements")
		} else {
			hasChange = true
		}
	}
	if !data.Collectors.Equal(state.Collectors) {
		tflog.Debug(ctx, "collectors has changed for appflowaction")
		if config.Collectors.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "collectors")
		} else {
			hasChange = true
		}
	}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for appflowaction")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			hasChange = true
		}
	}
	if !data.Distributionalgorithm.Equal(state.Distributionalgorithm) {
		tflog.Debug(ctx, "distributionalgorithm has changed for appflowaction")
		if config.Distributionalgorithm.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "distributionalgorithm")
		} else {
			hasChange = true
		}
	}
	if !data.Pagetracking.Equal(state.Pagetracking) {
		tflog.Debug(ctx, "pagetracking has changed for appflowaction")
		if config.Pagetracking.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "pagetracking")
		} else {
			hasChange = true
		}
	}
	if !data.Securityinsight.Equal(state.Securityinsight) {
		tflog.Debug(ctx, "securityinsight has changed for appflowaction")
		if config.Securityinsight.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "securityinsight")
		} else {
			hasChange = true
		}
	}
	if !data.Videoanalytics.Equal(state.Videoanalytics) {
		tflog.Debug(ctx, "videoanalytics has changed for appflowaction")
		if config.Videoanalytics.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "videoanalytics")
		} else {
			hasChange = true
		}
	}
	if !data.Webinsight.Equal(state.Webinsight) {
		tflog.Debug(ctx, "webinsight has changed for appflowaction")
		if config.Webinsight.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "webinsight")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model, restricted to NITRO-updatable fields
		appflowaction := appflowactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		appflowactionName := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Appflowaction.Type(), appflowactionName, &appflowaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appflowaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appflowaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Appflowaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset appflowaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readAppflowactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appflowaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowaction resource")
	// Named resource - delete using DeleteResource
	appflowactionName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appflowaction.Type(), appflowactionName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appflowaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appflowaction resource")
}

// Helper function to read appflowaction data from API
func (r *AppflowactionResource) readAppflowactionFromApi(ctx context.Context, data *AppflowactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	appflowactionName := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appflowaction.Type(), appflowactionName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowaction, got error: %s", err))
		return false
	}

	appflowactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
