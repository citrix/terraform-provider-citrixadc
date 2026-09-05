package appfwjsoncontenttype

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
var _ resource.Resource = &AppfwjsoncontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwjsoncontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwjsoncontenttypeResource)(nil)

func NewAppfwjsoncontenttypeResource() resource.Resource {
	return &AppfwjsoncontenttypeResource{}
}

// AppfwjsoncontenttypeResource defines the resource implementation.
type AppfwjsoncontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwjsoncontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwjsoncontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwjsoncontenttype"
}

func (r *AppfwjsoncontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwjsoncontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwjsoncontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwjsoncontenttype resource")
	// Get payload from plan
	appfwjsoncontenttype := appfwjsoncontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource - use AddResource
	jsoncontenttypevalue_value := data.Jsoncontenttypevalue.ValueString()
	_, err := r.client.AddResource(service.Appfwjsoncontenttype.Type(), jsoncontenttypevalue_value, &appfwjsoncontenttype)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwjsoncontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwjsoncontenttype resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Jsoncontenttypevalue.ValueString()))

	// Read the updated state back
	if !r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwjsoncontenttype not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsoncontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwjsoncontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwjsoncontenttype resource")

	found := r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwjsoncontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwjsoncontenttypeResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwjsoncontenttype resource")

	// All attributes (isregex, jsoncontenttypevalue) are ForceNew/RequiresReplace and
	// NITRO exposes no update operation for appfwjsoncontenttype, so there is nothing
	// to update in place. Any attribute change triggers a destroy/create instead.
	tflog.Debug(ctx, "No updateable attributes for appfwjsoncontenttype resource, skipping update")

	// Read the updated state back
	if !r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwjsoncontenttype not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsoncontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwjsoncontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwjsoncontenttype resource")
	// Named resource - delete using DeleteResource
	jsoncontenttypevalue_value := data.Jsoncontenttypevalue.ValueString()
	err := r.client.DeleteResource(service.Appfwjsoncontenttype.Type(), jsoncontenttypevalue_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwjsoncontenttype, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwjsoncontenttype resource")
}

// Helper function to read appfwjsoncontenttype data from API
func (r *AppfwjsoncontenttypeResource) readAppfwjsoncontenttypeFromApi(ctx context.Context, data *AppfwjsoncontenttypeResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	jsoncontenttypevalue_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwjsoncontenttype.Type(), jsoncontenttypevalue_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwjsoncontenttype, got error: %s", err))
		return false
	}

	appfwjsoncontenttypeSetAttrFromGet(ctx, data, getResponseData)

	return true
}
