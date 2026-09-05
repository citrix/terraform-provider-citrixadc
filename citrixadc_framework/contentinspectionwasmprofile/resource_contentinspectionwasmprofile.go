package contentinspectionwasmprofile

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
var _ resource.Resource = &ContentinspectionwasmprofileResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionwasmprofileResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionwasmprofileResource)(nil)

func NewContentinspectionwasmprofileResource() resource.Resource {
	return &ContentinspectionwasmprofileResource{}
}

// ContentinspectionwasmprofileResource defines the resource implementation.
type ContentinspectionwasmprofileResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionwasmprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionwasmprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionwasmprofile"
}

func (r *ContentinspectionwasmprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionwasmprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionwasmprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionwasmprofile resource")

	contentinspectionwasmprofile := contentinspectionwasmprofileGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Contentinspectionwasmprofile.Type(), name_value, &contentinspectionwasmprofile)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionwasmprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionwasmprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readContentinspectionwasmprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionwasmprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionwasmprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionwasmprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionwasmprofile resource")

	found := r.readContentinspectionwasmprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ContentinspectionwasmprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ContentinspectionwasmprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from configuration (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating contentinspectionwasmprofile resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Anomalousdatasize.Equal(state.Anomalousdatasize) {
		tflog.Debug(ctx, "anomalousdatasize has changed for contentinspectionwasmprofile")
		if config.Anomalousdatasize.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "anomalousdatasize")
		} else {
			hasChange = true
		}
	}
	if !data.Anomalousttfbtime.Equal(state.Anomalousttfbtime) {
		tflog.Debug(ctx, "anomalousttfbtime has changed for contentinspectionwasmprofile")
		if config.Anomalousttfbtime.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "anomalousttfbtime")
		} else {
			hasChange = true
		}
	}
	if !data.Maxbodylen.Equal(state.Maxbodylen) {
		tflog.Debug(ctx, "maxbodylen has changed for contentinspectionwasmprofile")
		if config.Maxbodylen.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "maxbodylen")
		} else {
			hasChange = true
		}
	}
	if !data.Timeout.Equal(state.Timeout) {
		tflog.Debug(ctx, "timeout has changed for contentinspectionwasmprofile")
		if config.Timeout.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timeout")
		} else {
			hasChange = true
		}
	}
	if !data.Timeoutaction.Equal(state.Timeoutaction) {
		tflog.Debug(ctx, "timeoutaction has changed for contentinspectionwasmprofile")
		if config.Timeoutaction.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "timeoutaction")
		} else {
			hasChange = true
		}
	}
	// wasmmodule does not support the NITRO unset operation; it is only pushed via update.
	if !data.Wasmmodule.Equal(state.Wasmmodule) {
		tflog.Debug(ctx, "wasmmodule has changed for contentinspectionwasmprofile")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		contentinspectionwasmprofile := contentinspectionwasmprofileGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Update is an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Contentinspectionwasmprofile.Type(), &contentinspectionwasmprofile)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionwasmprofile, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectionwasmprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionwasmprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Contentinspectionwasmprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset contentinspectionwasmprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readContentinspectionwasmprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionwasmprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionwasmprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionwasmprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionwasmprofile resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Contentinspectionwasmprofile.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionwasmprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionwasmprofile resource")
}

// Helper function to read contentinspectionwasmprofile data from API
func (r *ContentinspectionwasmprofileResource) readContentinspectionwasmprofileFromApi(ctx context.Context, data *ContentinspectionwasmprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Contentinspectionwasmprofile.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionwasmprofile, got error: %s", err))
		return false
	}

	contentinspectionwasmprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
