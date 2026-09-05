package appfwxmlcontenttype

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
var _ resource.Resource = &AppfwxmlcontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwxmlcontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwxmlcontenttypeResource)(nil)

func NewAppfwxmlcontenttypeResource() resource.Resource {
	return &AppfwxmlcontenttypeResource{}
}

// AppfwxmlcontenttypeResource defines the resource implementation.
type AppfwxmlcontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwxmlcontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwxmlcontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwxmlcontenttype"
}

func (r *AppfwxmlcontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwxmlcontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwxmlcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwxmlcontenttype resource")
	// Get payload from plan
	appfwxmlcontenttype := appfwxmlcontenttypeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	xmlcontenttypevalue_value := data.Xmlcontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwxmlcontenttype.Type(), xmlcontenttypevalue_value, &appfwxmlcontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwxmlcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwxmlcontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Xmlcontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwxmlcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlcontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlcontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwxmlcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwxmlcontenttype resource")

	found := r.readAppfwxmlcontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwxmlcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwxmlcontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwxmlcontenttype resource")

	// All configurable attributes (isregex, xmlcontenttypevalue) are ForceNew
	// (RequiresReplace); there are no in-place updatable NITRO fields, so no
	// update API call is made. Refresh state from the API.
	if !r.readAppfwxmlcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwxmlcontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlcontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwxmlcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwxmlcontenttype resource")
	// Named resource - delete using DeleteResource
	xmlcontenttypevalue_value := data.Xmlcontenttypevalue.ValueString()
	err := r.client.DeleteResource(service.Appfwxmlcontenttype.Type(), xmlcontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwxmlcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwxmlcontenttype resource")
}

// Helper function to read appfwxmlcontenttype data from API
func (r *AppfwxmlcontenttypeResource) readAppfwxmlcontenttypeFromApi(ctx context.Context, data *AppfwxmlcontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	xmlcontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwxmlcontenttype.Type(), xmlcontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwxmlcontenttype, got error: %s", err))
		return false
	}

	appfwxmlcontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
