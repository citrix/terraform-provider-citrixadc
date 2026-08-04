package appfwfieldtype

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
var _ resource.Resource = &AppfwfieldtypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwfieldtypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwfieldtypeResource)(nil)

func NewAppfwfieldtypeResource() resource.Resource {
	return &AppfwfieldtypeResource{}
}

// AppfwfieldtypeResource defines the resource implementation.
type AppfwfieldtypeResource struct {
	client *service.NitroClient
}

func (r *AppfwfieldtypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwfieldtypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwfieldtype"
}

func (r *AppfwfieldtypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwfieldtypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwfieldtypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwfieldtype resource")
	// Get payload from plan
	appfwfieldtype := appfwfieldtypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Appfwfieldtype.Type(), name_value, &appfwfieldtype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwfieldtype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwfieldtype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppfwfieldtypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwfieldtype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwfieldtypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwfieldtypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwfieldtype resource")

	found := r.readAppfwfieldtypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwfieldtypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwfieldtypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwfieldtype resource")

	// Check if there are any changes in updateable attributes
	hasChange := false
	if !data.Comment.Equal(state.Comment) {
		tflog.Debug(ctx, "comment has changed for appfwfieldtype")
		hasChange = true
	}
	if !data.Nocharmaps.Equal(state.Nocharmaps) {
		tflog.Debug(ctx, "nocharmaps has changed for appfwfieldtype")
		hasChange = true
	}
	if !data.Priority.Equal(state.Priority) {
		tflog.Debug(ctx, "priority has changed for appfwfieldtype")
		hasChange = true
	}
	if !data.Regex.Equal(state.Regex) {
		tflog.Debug(ctx, "regex has changed for appfwfieldtype")
		hasChange = true
	}

	if hasChange {
		// Create API request body from the model
		appfwfieldtype := appfwfieldtypeGetThePayloadFromtheConfig(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		name_value := data.Name.ValueString()
		_, err := r.client.UpdateResource(service.Appfwfieldtype.Type(), name_value, &appfwfieldtype)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwfieldtype, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwfieldtype resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwfieldtype resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppfwfieldtypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwfieldtype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwfieldtypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwfieldtypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwfieldtype resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwfieldtype.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwfieldtype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwfieldtype resource")
}

// Helper function to read appfwfieldtype data from API
func (r *AppfwfieldtypeResource) readAppfwfieldtypeFromApi(ctx context.Context, data *AppfwfieldtypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	appfwfieldtype_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwfieldtype.Type(), appfwfieldtype_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwfieldtype, got error: %s", err))
		return false
	}

	appfwfieldtypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
