package contentinspectionaction

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
var _ resource.Resource = &ContentinspectionactionResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectionactionResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectionactionResource)(nil)

func NewContentinspectionactionResource() resource.Resource {
	return &ContentinspectionactionResource{}
}

// ContentinspectionactionResource defines the resource implementation.
type ContentinspectionactionResource struct {
	client *service.NitroClient
}

func (r *ContentinspectionactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectionactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionaction"
}

func (r *ContentinspectionactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectionactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectionaction resource")

	contentinspectionaction := contentinspectionactionGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Contentinspectionaction.Type(), name_value, &contentinspectionaction)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectionaction resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionaction not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectionaction resource")

	found := r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ContentinspectionactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ContentinspectionactionResourceModel

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

	tflog.Debug(ctx, "Updating contentinspectionaction resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Icapprofilename.Equal(state.Icapprofilename) {
		tflog.Debug(ctx, "icapprofilename has changed for contentinspectionaction")
		hasChange = true
	}
	if !data.Ifserverdown.Equal(state.Ifserverdown) {
		tflog.Debug(ctx, "ifserverdown has changed for contentinspectionaction")
		if config.Ifserverdown.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "ifserverdown")
		} else {
			hasChange = true
		}
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, "serverip has changed for contentinspectionaction")
		hasChange = true
	}
	if !data.Servername.Equal(state.Servername) {
		tflog.Debug(ctx, "servername has changed for contentinspectionaction")
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, "serverport has changed for contentinspectionaction")
		if config.Serverport.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverport")
		} else {
			hasChange = true
		}
	}
	if !data.Wasmprofilename.Equal(state.Wasmprofilename) {
		tflog.Debug(ctx, "wasmprofilename has changed for contentinspectionaction")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model (updatable fields only)
		contentinspectionaction := contentinspectionactionGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Make API call
		// Update is an unnamed PUT (name is carried in the payload)
		err := r.client.UpdateUnnamedResource(service.Contentinspectionaction.Type(), &contentinspectionaction)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectionaction, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectionaction resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectionaction resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Contentinspectionaction.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset contentinspectionaction attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readContentinspectionactionFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectionaction not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectionactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectionactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectionaction resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Contentinspectionaction.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectionaction, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectionaction resource")
}

// Helper function to read contentinspectionaction data from API
func (r *ContentinspectionactionResource) readContentinspectionactionFromApi(ctx context.Context, data *ContentinspectionactionResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	name_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Contentinspectionaction.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionaction, got error: %s", err))
		return false
	}

	contentinspectionactionSetAttrFromGet(ctx, data, getResponseData)

	return true
}
