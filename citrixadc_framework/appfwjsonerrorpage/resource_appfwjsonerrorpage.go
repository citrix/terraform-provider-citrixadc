package appfwjsonerrorpage

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
var _ resource.Resource = &AppfwjsonerrorpageResource{}
var _ resource.ResourceWithConfigure = (*AppfwjsonerrorpageResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwjsonerrorpageResource)(nil)

func NewAppfwjsonerrorpageResource() resource.Resource {
	return &AppfwjsonerrorpageResource{}
}

// AppfwjsonerrorpageResource defines the resource implementation.
type AppfwjsonerrorpageResource struct {
	client *service.NitroClient
}

func (r *AppfwjsonerrorpageResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwjsonerrorpageResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwjsonerrorpage"
}

func (r *AppfwjsonerrorpageResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwjsonerrorpageResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwjsonerrorpage resource")

	// appfwjsonerrorpage := appfwjsonerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwjsonerrorpage.Type(), &appfwjsonerrorpage)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwjsonerrorpage, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwjsonerrorpage-config")

	tflog.Trace(ctx, "Created appfwjsonerrorpage resource")

	// Read the updated state back
	r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsonerrorpageResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwjsonerrorpage resource")

	r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsonerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwjsonerrorpage resource")

	// Create API request body from the model
	// appfwjsonerrorpage := appfwjsonerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwjsonerrorpage.Type(), &appfwjsonerrorpage)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwjsonerrorpage, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwjsonerrorpage resource")

	// Read the updated state back
	r.readAppfwjsonerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsonerrorpageResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwjsonerrorpageResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwjsonerrorpage resource")

	// For appfwjsonerrorpage, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwjsonerrorpage resource from state")
}

// Helper function to read appfwjsonerrorpage data from API
func (r *AppfwjsonerrorpageResource) readAppfwjsonerrorpageFromApi(ctx context.Context, data *AppfwjsonerrorpageResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwjsonerrorpage.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwjsonerrorpage, got error: %s", err))
		return
	}

	appfwjsonerrorpageSetAttrFromGet(ctx, data, getResponseData)

}
