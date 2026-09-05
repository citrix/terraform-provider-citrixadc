package appfwhtmlerrorpage

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
var _ resource.Resource = &AppfwhtmlerrorpageResource{}
var _ resource.ResourceWithConfigure = (*AppfwhtmlerrorpageResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwhtmlerrorpageResource)(nil)

func NewAppfwhtmlerrorpageResource() resource.Resource {
	return &AppfwhtmlerrorpageResource{}
}

// AppfwhtmlerrorpageResource defines the resource implementation.
type AppfwhtmlerrorpageResource struct {
	client *service.NitroClient
}

func (r *AppfwhtmlerrorpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwhtmlerrorpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwhtmlerrorpage"
}

func (r *AppfwhtmlerrorpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwhtmlerrorpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwhtmlerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwhtmlerrorpage resource")

	// Build the payload from the plan
	appfwhtmlerrorpage := appfwhtmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// Named resource created via the "Import" action (mirrors SDK v2 ActOnResource)
	err := r.client.ActOnResource(service.Appfwhtmlerrorpage.Type(), &appfwhtmlerrorpage, "Import")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwhtmlerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created appfwhtmlerrorpage resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Read the updated state back
	if !r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwhtmlerrorpage not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwhtmlerrorpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwhtmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwhtmlerrorpage resource")

	found := r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *AppfwhtmlerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AppfwhtmlerrorpageResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating appfwhtmlerrorpage resource")

	// All attributes are ForceNew/RequiresReplace; there are no in-place updatable
	// fields for appfwhtmlerrorpage, so no NITRO update call is issued here.
	tflog.Trace(ctx, "Updated appfwhtmlerrorpage resource")

	// Read the updated state back
	if !r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "appfwhtmlerrorpage not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwhtmlerrorpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwhtmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwhtmlerrorpage resource")
	// Named resource - delete using DeleteResource
	name_value := data.Name.ValueString()
	err := r.client.DeleteResource(service.Appfwhtmlerrorpage.Type(), name_value)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete appfwhtmlerrorpage, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted appfwhtmlerrorpage resource")
}

// Helper function to read appfwhtmlerrorpage data from API
func (r *AppfwhtmlerrorpageResource) readAppfwhtmlerrorpageFromApi(ctx context.Context, data *AppfwhtmlerrorpageResourceModel, diags *diag.Diagnostics) bool {

	// Case 2: Find with single ID attribute - ID is the plain value
	appfwhtmlerrorpage_Name := data.Id.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = r.client.FindResource(service.Appfwhtmlerrorpage.Type(), appfwhtmlerrorpage_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwhtmlerrorpage, got error: %s", err))
		return false
	}

	appfwhtmlerrorpageSetAttrFromGet(ctx, data, getResponseData)

	return true
}
