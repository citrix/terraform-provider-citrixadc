package appfwmultipartformcontenttype

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
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

	// appfwmultipartformcontenttype := appfwmultipartformcontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwmultipartformcontenttype.Type(), &appfwmultipartformcontenttype)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwmultipartformcontenttype, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwmultipartformcontenttype-config")

	tflog.Trace(ctx, "Created appfwmultipartformcontenttype resource")

	// Read the updated state back
	r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwmultipartformcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwmultipartformcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwmultipartformcontenttype resource")

	// Create API request body from the model
	// appfwmultipartformcontenttype := appfwmultipartformcontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwmultipartformcontenttype.Type(), &appfwmultipartformcontenttype)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwmultipartformcontenttype, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwmultipartformcontenttype resource")

	// Read the updated state back
	r.readAppfwmultipartformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

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

	// For appfwmultipartformcontenttype, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwmultipartformcontenttype resource from state")
}

// Helper function to read appfwmultipartformcontenttype data from API
func (r *AppfwmultipartformcontenttypeResource) readAppfwmultipartformcontenttypeFromApi(ctx context.Context, data *AppfwmultipartformcontenttypeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwmultipartformcontenttype.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwmultipartformcontenttype, got error: %s", err))
		return
	}

	appfwmultipartformcontenttypeSetAttrFromGet(ctx, data, getResponseData)

}
