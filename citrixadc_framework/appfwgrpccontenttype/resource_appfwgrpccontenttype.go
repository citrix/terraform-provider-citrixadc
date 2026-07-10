package appfwgrpccontenttype

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
var _ resource.Resource = &AppfwgrpccontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwgrpccontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwgrpccontenttypeResource)(nil)

func NewAppfwgrpccontenttypeResource() resource.Resource {
	return &AppfwgrpccontenttypeResource{}
}

// AppfwgrpccontenttypeResource defines the resource implementation.
type AppfwgrpccontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwgrpccontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwgrpccontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwgrpccontenttype"
}

func (r *AppfwgrpccontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwgrpccontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwgrpccontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwgrpccontenttype resource")
	appfwgrpccontenttype := appfwgrpccontenttypeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	grpccontenttypevalue_value := data.Grpccontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwgrpccontenttype.Type(), grpccontenttypevalue_value, &appfwgrpccontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwgrpccontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwgrpccontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Grpccontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwgrpccontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpccontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpccontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwgrpccontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwgrpccontenttype resource")

	found := r.readAppfwgrpccontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwgrpccontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwgrpccontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Update is a no-op for appfwgrpccontenttype; NITRO exposes no update endpoint and all attributes are RequiresReplace")

	// Read the updated state back
	if !r.readAppfwgrpccontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpccontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpccontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwgrpccontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwgrpccontenttype resource")
	// Named resource - delete using DeleteResource
	grpccontenttypevalue_value := data.Grpccontenttypevalue.ValueString()
	err := r.client.DeleteResource(service.Appfwgrpccontenttype.Type(), grpccontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwgrpccontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwgrpccontenttype resource")
}

// Helper function to read appfwgrpccontenttype data from API
func (r *AppfwgrpccontenttypeResource) readAppfwgrpccontenttypeFromApi(ctx context.Context, data *AppfwgrpccontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	grpccontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwgrpccontenttype.Type(), grpccontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwgrpccontenttype, got error: %s", err))
		return false
	}

	appfwgrpccontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
