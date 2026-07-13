package appfwgrpcwebtextcontenttype

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
var _ resource.Resource = &AppfwgrpcwebtextcontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwgrpcwebtextcontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwgrpcwebtextcontenttypeResource)(nil)

func NewAppfwgrpcwebtextcontenttypeResource() resource.Resource {
	return &AppfwgrpcwebtextcontenttypeResource{}
}

// AppfwgrpcwebtextcontenttypeResource defines the resource implementation.
type AppfwgrpcwebtextcontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwgrpcwebtextcontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwgrpcwebtextcontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwgrpcwebtextcontenttype"
}

func (r *AppfwgrpcwebtextcontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwgrpcwebtextcontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwgrpcwebtextcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwgrpcwebtextcontenttype resource")
	appfwgrpcwebtextcontenttype := appfwgrpcwebtextcontenttypeGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	grpcwebtextcontenttypevalue_value := data.Grpcwebtextcontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwgrpcwebtextcontenttype.Type(), grpcwebtextcontenttypevalue_value, &appfwgrpcwebtextcontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwgrpcwebtextcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwgrpcwebtextcontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Grpcwebtextcontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwgrpcwebtextcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpcwebtextcontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpcwebtextcontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwgrpcwebtextcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwgrpcwebtextcontenttype resource")

	found := r.readAppfwgrpcwebtextcontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwgrpcwebtextcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwgrpcwebtextcontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwgrpcwebtextcontenttype resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		appfwgrpcwebtextcontenttype := appfwgrpcwebtextcontenttypeGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Named resource - use UpdateResource
		grpcwebtextcontenttypevalue_value := data.Grpcwebtextcontenttypevalue.ValueString()
		_, err := r.client.UpdateResource(service.Appfwgrpcwebtextcontenttype.Type(), grpcwebtextcontenttypevalue_value, &appfwgrpcwebtextcontenttype)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwgrpcwebtextcontenttype, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated appfwgrpcwebtextcontenttype resource")
	} else {
		tflog.Debug(ctx, "No changes detected for appfwgrpcwebtextcontenttype resource, skipping update")
	}

	// Read the updated state back
	if !r.readAppfwgrpcwebtextcontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwgrpcwebtextcontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwgrpcwebtextcontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwgrpcwebtextcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwgrpcwebtextcontenttype resource")
	// Named resource - delete using DeleteResource
	grpcwebtextcontenttypevalue_value := data.Grpcwebtextcontenttypevalue.ValueString()
	err := r.client.DeleteResource(service.Appfwgrpcwebtextcontenttype.Type(), grpcwebtextcontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwgrpcwebtextcontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwgrpcwebtextcontenttype resource")
}

// Helper function to read appfwgrpcwebtextcontenttype data from API
func (r *AppfwgrpcwebtextcontenttypeResource) readAppfwgrpcwebtextcontenttypeFromApi(ctx context.Context, data *AppfwgrpcwebtextcontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	grpcwebtextcontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwgrpcwebtextcontenttype.Type(), grpcwebtextcontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwgrpcwebtextcontenttype, got error: %s", err))
		return false
	}

	appfwgrpcwebtextcontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
