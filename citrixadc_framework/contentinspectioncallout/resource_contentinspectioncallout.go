package contentinspectioncallout

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
var _ resource.Resource = &ContentinspectioncalloutResource{}
var _ resource.ResourceWithConfigure = (*ContentinspectioncalloutResource)(nil)
var _ resource.ResourceWithImportState = (*ContentinspectioncalloutResource)(nil)

func NewContentinspectioncalloutResource() resource.Resource {
	return &ContentinspectioncalloutResource{}
}

// ContentinspectioncalloutResource defines the resource implementation.
type ContentinspectioncalloutResource struct {
	client *service.NitroClient
}

func (r *ContentinspectioncalloutResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ContentinspectioncalloutResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectioncallout"
}

func (r *ContentinspectioncalloutResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ContentinspectioncalloutResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ContentinspectioncalloutResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating contentinspectioncallout resource")

	contentinspectioncallout := contentinspectioncalloutGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	contentinspectioncalloutName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Contentinspectioncallout.Type(), contentinspectioncalloutName, &contentinspectioncallout)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create contentinspectioncallout, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created contentinspectioncallout resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", contentinspectioncalloutName))

	// Read the updated state back
	if !r.readContentinspectioncalloutFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectioncallout not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectioncalloutResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ContentinspectioncalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading contentinspectioncallout resource")

	found := r.readContentinspectioncalloutFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ContentinspectioncalloutResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ContentinspectioncalloutResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating contentinspectioncallout resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	attributesToUnset := []string{}
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Profilename.Equal(state.Profilename) {
		tflog.Debug(ctx, "profilename has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Resultexpr.Equal(state.Resultexpr) {
		tflog.Debug(ctx, "resultexpr has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Returntype.Equal(state.Returntype) {
		tflog.Debug(ctx, "returntype has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Serverip.Equal(state.Serverip) {
		tflog.Debug(ctx, "serverip has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Servername.Equal(state.Servername) {
		tflog.Debug(ctx, "servername has changed for contentinspectioncallout")
		hasChange = true
	}
	if !data.Serverport.Equal(state.Serverport) {
		tflog.Debug(ctx, "serverport has changed for contentinspectioncallout")
		if config.Serverport.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "serverport")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// Create API request body from the model (updatable fields only)
		contentinspectioncallout := contentinspectioncalloutGetTheUpdatablePayloadFromThePlan(ctx, &data)

		// NITRO update verb is an unnamed PUT (name is carried in the payload).
		err := r.client.UpdateUnnamedResource(service.Contentinspectioncallout.Type(), &contentinspectioncallout)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update contentinspectioncallout, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated contentinspectioncallout resource")
	} else {
		tflog.Debug(ctx, "No changes detected for contentinspectioncallout resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their NITRO defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Contentinspectioncallout.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset contentinspectioncallout attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readContentinspectioncalloutFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "contentinspectioncallout not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ContentinspectioncalloutResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ContentinspectioncalloutResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting contentinspectioncallout resource")

	// Named resource - delete using DeleteResource keyed on the ID (name)
	contentinspectioncalloutName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Contentinspectioncallout.Type(), contentinspectioncalloutName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete contentinspectioncallout, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted contentinspectioncallout resource")
}

// Helper function to read contentinspectioncallout data from API
func (r *ContentinspectioncalloutResource) readContentinspectioncalloutFromApi(ctx context.Context, data *ContentinspectioncalloutResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value (name)
	contentinspectioncalloutName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Contentinspectioncallout.Type(), contentinspectioncalloutName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectioncallout, got error: %s", err))
		return false
	}

	contentinspectioncalloutSetAttrFromGet(ctx, data, getResponseData)

	return true
}
