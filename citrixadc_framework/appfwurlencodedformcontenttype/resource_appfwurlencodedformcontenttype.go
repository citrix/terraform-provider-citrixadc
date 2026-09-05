package appfwurlencodedformcontenttype

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
var _ resource.Resource = &AppfwurlencodedformcontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwurlencodedformcontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwurlencodedformcontenttypeResource)(nil)

func NewAppfwurlencodedformcontenttypeResource() resource.Resource {
	return &AppfwurlencodedformcontenttypeResource{}
}

// AppfwurlencodedformcontenttypeResource defines the resource implementation.
type AppfwurlencodedformcontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwurlencodedformcontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwurlencodedformcontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwurlencodedformcontenttype"
}

func (r *AppfwurlencodedformcontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwurlencodedformcontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwurlencodedformcontenttype resource")

	// Get payload from plan
	appfwurlencodedformcontenttype := appfwurlencodedformcontenttypeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	urlencodedformcontenttypevalue_value := data.Urlencodedformcontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwurlencodedformcontenttype.Type(), urlencodedformcontenttypevalue_value, &appfwurlencodedformcontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwurlencodedformcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwurlencodedformcontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Urlencodedformcontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwurlencodedformcontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwurlencodedformcontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwurlencodedformcontenttype resource")

	found := r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwurlencodedformcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwurlencodedformcontenttype resource")

	// All configurable attributes (urlencodedformcontenttypevalue, isregex) are
	// ForceNew/RequiresReplace, so there is no in-place update path. Simply read
	// the current state back from the API.
	if !r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwurlencodedformcontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwurlencodedformcontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwurlencodedformcontenttype resource")
	// Named resource - delete using DeleteResource
	urlencodedformcontenttypevalue_value := data.Id.ValueString()
	err := r.client.DeleteResource(service.Appfwurlencodedformcontenttype.Type(), urlencodedformcontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwurlencodedformcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwurlencodedformcontenttype resource")
}

// Helper function to read appfwurlencodedformcontenttype data from API
func (r *AppfwurlencodedformcontenttypeResource) readAppfwurlencodedformcontenttypeFromApi(ctx context.Context, data *AppfwurlencodedformcontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	urlencodedformcontenttypevalue_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Appfwurlencodedformcontenttype.Type(), urlencodedformcontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwurlencodedformcontenttype, got error: %s", err))
		return false
	}

	appfwurlencodedformcontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
