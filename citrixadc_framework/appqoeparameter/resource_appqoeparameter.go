package appqoeparameter

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
var _ resource.Resource = &AppqoeparameterResource{}
var _ resource.ResourceWithConfigure = (*AppqoeparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoeparameterResource)(nil)

func NewAppqoeparameterResource() resource.Resource {
	return &AppqoeparameterResource{}
}

// AppqoeparameterResource defines the resource implementation.
type AppqoeparameterResource struct {
	client *service.NitroClient
}

func (r *AppqoeparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoeparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoeparameter"
}

func (r *AppqoeparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoeparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoeparameter resource")

	// appqoeparameter := appqoeparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoeparameter, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appqoeparameter-config")

	tflog.Trace(ctx, "Created appqoeparameter resource")

	// Read the updated state back
	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoeparameter resource")

	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appqoeparameter resource")

	// Create API request body from the model
	// appqoeparameter := appqoeparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appqoeparameter.Type(), &appqoeparameter)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appqoeparameter, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appqoeparameter resource")

	// Read the updated state back
	r.readAppqoeparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoeparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoeparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoeparameter resource")

	// For appqoeparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appqoeparameter resource from state")
}

// Helper function to read appqoeparameter data from API
func (r *AppqoeparameterResource) readAppqoeparameterFromApi(ctx context.Context, data *AppqoeparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appqoeparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoeparameter, got error: %s", err))
		return
	}

	appqoeparameterSetAttrFromGet(ctx, data, getResponseData)

}
