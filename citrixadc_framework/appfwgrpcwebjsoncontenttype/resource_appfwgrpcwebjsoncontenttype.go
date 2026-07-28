package appfwgrpcwebjsoncontenttype

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
var _ resource.Resource = &AppfwgrpcwebjsoncontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwgrpcwebjsoncontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwgrpcwebjsoncontenttypeResource)(nil)

func NewAppfwgrpcwebjsoncontenttypeResource() resource.Resource {
	return &AppfwgrpcwebjsoncontenttypeResource{}
}

// AppfwgrpcwebjsoncontenttypeResource defines the resource implementation.
type AppfwgrpcwebjsoncontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwgrpcwebjsoncontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwgrpcwebjsoncontenttype"
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwgrpcwebjsoncontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwgrpcwebjsoncontenttype resource")
	appfwgrpcwebjsoncontenttype := appfwgrpcwebjsoncontenttypeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	grpcwebjsoncontenttypevalue_value := data.Grpcwebjsoncontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwgrpcwebjsoncontenttype.Type(), grpcwebjsoncontenttypevalue_value, &appfwgrpcwebjsoncontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwgrpcwebjsoncontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwgrpcwebjsoncontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Grpcwebjsoncontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwgrpcwebjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpcwebjsoncontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwgrpcwebjsoncontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwgrpcwebjsoncontenttype resource")

	found := r.readAppfwgrpcwebjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwgrpcwebjsoncontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwgrpcwebjsoncontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for appfwgrpcwebjsoncontenttype; NITRO exposes no update endpoint and all attributes are RequiresReplace")

	// Read the updated state back
	if !r.readAppfwgrpcwebjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpcwebjsoncontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpcwebjsoncontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwgrpcwebjsoncontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwgrpcwebjsoncontenttype resource")
	// Named resource - delete using DeleteResource
	grpcwebjsoncontenttypevalue_value := data.Grpcwebjsoncontenttypevalue.ValueString()
	err := r.client.DeleteResource(service.Appfwgrpcwebjsoncontenttype.Type(), grpcwebjsoncontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwgrpcwebjsoncontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwgrpcwebjsoncontenttype resource")
}

// Helper function to read appfwgrpcwebjsoncontenttype data from API
func (r *AppfwgrpcwebjsoncontenttypeResource) readAppfwgrpcwebjsoncontenttypeFromApi(ctx context.Context, data *AppfwgrpcwebjsoncontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	grpcwebjsoncontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwgrpcwebjsoncontenttype.Type(), grpcwebjsoncontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwgrpcwebjsoncontenttype, got error: %s", err))
		return false
	}

	appfwgrpcwebjsoncontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
