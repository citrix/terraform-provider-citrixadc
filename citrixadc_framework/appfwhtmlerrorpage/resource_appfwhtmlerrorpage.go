package appfwhtmlerrorpage

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

	// appfwhtmlerrorpage := appfwhtmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwhtmlerrorpage.Type(), &appfwhtmlerrorpage)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwhtmlerrorpage, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwhtmlerrorpage-config")

	tflog.Trace(ctx, "Created appfwhtmlerrorpage resource")

	// Read the updated state back
	r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwhtmlerrorpageResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwhtmlerrorpageResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwhtmlerrorpage resource")

	// Create API request body from the model
	// appfwhtmlerrorpage := appfwhtmlerrorpageGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwhtmlerrorpage.Type(), &appfwhtmlerrorpage)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwhtmlerrorpage, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwhtmlerrorpage resource")

	// Read the updated state back
	r.readAppfwhtmlerrorpageFromApi(ctx, &data, &resp.Diagnostics)

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

	// For appfwhtmlerrorpage, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwhtmlerrorpage resource from state")
}

// Helper function to read appfwhtmlerrorpage data from API
func (r *AppfwhtmlerrorpageResource) readAppfwhtmlerrorpageFromApi(ctx context.Context, data *AppfwhtmlerrorpageResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwhtmlerrorpage.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwhtmlerrorpage, got error: %s", err))
		return
	}

	appfwhtmlerrorpageSetAttrFromGet(ctx, data, getResponseData)

}
