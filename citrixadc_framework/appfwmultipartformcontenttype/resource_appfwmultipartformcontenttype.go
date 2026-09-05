package appfwmultipartformcontenttype

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
var _ resource.Resource = &AppfwmultipartformcontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwmultipartformcontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwmultipartformcontenttypeResource)(nil)

func NewAppfwmultipartformcontenttypeResource() resource.Resource {
	return &AppfwmultipartformcontenttypeResource{}
}

// AppfwmultipartformcontenttypeResource defines the resource implementation.
type AppfwmultipartformcontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwmultipartformcontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwmultipartformcontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwmultipartformcontenttype"
}

func (r *AppfwmultipartformcontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwmultipartformcontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwmultipartformcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwmultipartformcontenttype resource")
	// Get payload from plan
	appfwmultipartformcontenttype := appfwmultipartformcontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	multipartformcontenttypevalue_value := data.Multipartformcontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwmultipartformcontenttype.Type(), multipartformcontenttypevalue_value, &appfwmultipartformcontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwmultipartformcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwmultipartformcontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Multipartformcontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwmultipartformcontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwmultipartformcontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwmultipartformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwmultipartformcontenttype resource")

	found := r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwmultipartformcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwmultipartformcontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwmultipartformcontenttype resource")
	// appfwmultipartformcontenttype has no NITRO update operation and every attribute forces
	// replacement (ForceNew), so this path performs no API mutation and only refreshes state.

	// Read the updated state back
	if !r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwmultipartformcontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwmultipartformcontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwmultipartformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwmultipartformcontenttype resource")
	// Named resource - delete using DeleteResource
	multipartformcontenttypevalue_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwmultipartformcontenttype.Type(), multipartformcontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwmultipartformcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwmultipartformcontenttype resource")
}

// Helper function to read appfwmultipartformcontenttype data from API
func (r *AppfwmultipartformcontenttypeResource) readAppfwmultipartformcontenttypeFromApi(ctx context.Context, data *AppfwmultipartformcontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	multipartformcontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwmultipartformcontenttype.Type(), multipartformcontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwmultipartformcontenttype, got error: %s", err))
		return false
	}

	appfwmultipartformcontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
