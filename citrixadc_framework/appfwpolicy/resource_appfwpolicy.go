package appfwpolicy

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
var _ resource.Resource = &AppfwpolicyResource{}
var _ resource.ResourceWithConfigure = (*AppfwpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwpolicyResource)(nil)

func NewAppfwpolicyResource() resource.Resource {
	return &AppfwpolicyResource{}
}

// AppfwpolicyResource defines the resource implementation.
type AppfwpolicyResource struct {
	client *service.NitroClient
}

func (r *AppfwpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwpolicy"
}

func (r *AppfwpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwpolicy resource")

	// appfwpolicy := appfwpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwpolicy.Type(), &appfwpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwpolicy-config")

	tflog.Trace(ctx, "Created appfwpolicy resource")

	// Read the updated state back
	r.readAppfwpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwpolicy resource")

	r.readAppfwpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwpolicy resource")

	// Create API request body from the model
	// appfwpolicy := appfwpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwpolicy.Type(), &appfwpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwpolicy resource")

	// Read the updated state back
	r.readAppfwpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwpolicy resource")

	// For appfwpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwpolicy resource from state")
}

// Helper function to read appfwpolicy data from API
func (r *AppfwpolicyResource) readAppfwpolicyFromApi(ctx context.Context, data *AppfwpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwpolicy, got error: %s", err))
		return
	}

	appfwpolicySetAttrFromGet(ctx, data, getResponseData)

}
