package appfwxmlerrorpage

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
var _ resource.Resource = &AppfwxmlerrorpageResource{}
var _ resource.ResourceWithConfigure = (*AppfwxmlerrorpageResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwxmlerrorpageResource)(nil)

func NewAppfwxmlerrorpageResource() resource.Resource {
	return &AppfwxmlerrorpageResource{}
}

// AppfwxmlerrorpageResource defines the resource implementation.
type AppfwxmlerrorpageResource struct {
	client *service.NitroClient
}

func (r *AppfwxmlerrorpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwxmlerrorpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwxmlerrorpage"
}

func (r *AppfwxmlerrorpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwxmlerrorpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwxmlerrorpage resource")

	// appfwxmlerrorpage := appfwxmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwxmlerrorpage.Type(), &appfwxmlerrorpage)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwxmlerrorpage, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwxmlerrorpage-config")

	tflog.Trace(ctx, "Created appfwxmlerrorpage resource")

	// Read the updated state back
	r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlerrorpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwxmlerrorpage resource")

	r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwxmlerrorpage resource")

	// Create API request body from the model
	// appfwxmlerrorpage := appfwxmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwxmlerrorpage.Type(), &appfwxmlerrorpage)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwxmlerrorpage, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwxmlerrorpage resource")

	// Read the updated state back
	r.readAppfwxmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwxmlerrorpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwxmlerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwxmlerrorpage resource")

	// For appfwxmlerrorpage, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwxmlerrorpage resource from state")
}

// Helper function to read appfwxmlerrorpage data from API
func (r *AppfwxmlerrorpageResource) readAppfwxmlerrorpageFromApi(ctx context.Context, data *AppfwxmlerrorpageResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwxmlerrorpage.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwxmlerrorpage, got error: %s", err))
		return
	}

	appfwxmlerrorpageSetAttrFromGet(ctx, data, getResponseData)

}
