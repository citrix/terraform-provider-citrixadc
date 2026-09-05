package transformprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/transform"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TransformprofileResource{}
var _ resource.ResourceWithConfigure = (*TransformprofileResource)(nil)
var _ resource.ResourceWithImportState = (*TransformprofileResource)(nil)

func NewTransformprofileResource() resource.Resource {
	return &TransformprofileResource{}
}

// TransformprofileResource defines the resource implementation.
type TransformprofileResource struct {
	client *service.NitroClient
}

func (r *TransformprofileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TransformprofileResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_transformprofile"
}

func (r *TransformprofileResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TransformprofileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TransformprofileResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating transformprofile resource")
	transformprofileName := data.Name.ValueString()

	// Named resource - use AddResource.
	// Only name (and type) are valid for the create operation; the remaining
	// parameters must be applied through a follow-up update (matches SDK v2).
	addPayload := transform.Transformprofile{
		Name: data.Name.ValueString(),
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		addPayload.Type = data.Type.ValueString()
	}
	_, err := r.client.AddResource(service.Transformprofile.Type(), transformprofileName, &addPayload)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create transformprofile, got error: %s", err))
		return
	}

	// Follow-up update to include the parameters that are invalid for create.
	updatePayload := transformprofileGetThePayloadFromtheConfig(ctx, &data)
	doUpdate := false
	if updatePayload.Comment != "" {
		doUpdate = true
	}
	if updatePayload.Onlytransformabsurlinbody != "" {
		doUpdate = true
	}
	if updatePayload.Type != "" {
		doUpdate = true
	}
	if doUpdate {
		_, err := r.client.UpdateResource(service.Transformprofile.Type(), transformprofileName, &updatePayload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformprofile after create, got error: %s", err))
			return
		}
	}

	tflog.Trace(ctx, "Created transformprofile resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(transformprofileName)

	// Read the updated state back
	if !r.readTransformprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformprofile not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformprofileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TransformprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading transformprofile resource")

	found := r.readTransformprofileFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *TransformprofileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state TransformprofileResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (-> unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating transformprofile resource")
	transformprofileName := data.Name.ValueString()

	// Named resource - use UpdateResource for changed attributes only.
	payload := transform.Transformprofile{
		Name: data.Name.ValueString(),
	}
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for transformprofile")
		if config.Comment.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "comment")
		} else {
			payload.Comment = data.Comment.ValueString()
			hasChange = true
		}
	}
	if !data.Onlytransformabsurlinbody.Equal(state.Onlytransformabsurlinbody) {
		tflog.Debug(ctx, "onlytransformabsurlinbody has changed for transformprofile")
		if config.Onlytransformabsurlinbody.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "onlytransformabsurlinbody")
		} else {
			payload.Onlytransformabsurlinbody = data.Onlytransformabsurlinbody.ValueString()
			hasChange = true
		}
	}
	if !data.Type.Equal(state.Type) {
		tflog.Debug(ctx, "type has changed for transformprofile")
		payload.Type = data.Type.ValueString()
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Transformprofile.Type(), transformprofileName, &payload)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update transformprofile, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated transformprofile resource")
	} else {
		tflog.Debug(ctx, "No changes detected for transformprofile resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Transformprofile.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset transformprofile attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readTransformprofileFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "transformprofile not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TransformprofileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TransformprofileResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting transformprofile resource")
	// Named resource - delete using DeleteResource
	transformprofileName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Transformprofile.Type(), transformprofileName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete transformprofile, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted transformprofile resource")
}

// Helper function to read transformprofile data from API
func (r *TransformprofileResource) readTransformprofileFromApi(ctx context.Context, data *TransformprofileResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	transformprofileName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Transformprofile.Type(), transformprofileName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read transformprofile, got error: %s", err))
		return false
	}

	transformprofileSetAttrFromGet(ctx, data, getResponseData)

	return true
}
