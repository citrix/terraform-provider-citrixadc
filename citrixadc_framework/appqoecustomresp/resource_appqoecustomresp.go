package appqoecustomresp

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
var _ resource.Resource = &AppqoecustomrespResource{}
var _ resource.ResourceWithConfigure = (*AppqoecustomrespResource)(nil)
var _ resource.ResourceWithImportState = (*AppqoecustomrespResource)(nil)

func NewAppqoecustomrespResource() resource.Resource {
	return &AppqoecustomrespResource{}
}

// AppqoecustomrespResource defines the resource implementation.
type AppqoecustomrespResource struct {
	client *service.NitroClient
}

func (r *AppqoecustomrespResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppqoecustomrespResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appqoecustomresp"
}

func (r *AppqoecustomrespResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppqoecustomrespResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appqoecustomresp resource")

	// appqoecustomresp := appqoecustomrespGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appqoecustomresp.Type(), &appqoecustomresp)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appqoecustomresp, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appqoecustomresp-config")

	tflog.Trace(ctx, "Created appqoecustomresp resource")

	// Read the updated state back
	r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appqoecustomresp resource")

	r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appqoecustomresp resource")

	// Create API request body from the model
	// appqoecustomresp := appqoecustomrespGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appqoecustomresp.Type(), &appqoecustomresp)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appqoecustomresp, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appqoecustomresp resource")

	// Read the updated state back
	r.readAppqoecustomrespFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppqoecustomrespResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppqoecustomrespResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appqoecustomresp resource")

	// For appqoecustomresp, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appqoecustomresp resource from state")
}

// Helper function to read appqoecustomresp data from API
func (r *AppqoecustomrespResource) readAppqoecustomrespFromApi(ctx context.Context, data *AppqoecustomrespResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appqoecustomresp.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appqoecustomresp, got error: %s", err))
		return
	}

	appqoecustomrespSetAttrFromGet(ctx, data, getResponseData)

}
