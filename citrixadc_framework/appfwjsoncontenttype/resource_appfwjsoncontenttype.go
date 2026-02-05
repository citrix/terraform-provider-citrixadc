package appfwjsoncontenttype

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

	// appfwjsoncontenttype := appfwjsoncontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwjsoncontenttype.Type(), &appfwjsoncontenttype)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwjsoncontenttype, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwjsoncontenttype-config")

	tflog.Trace(ctx, "Created appfwjsoncontenttype resource")

	// Read the updated state back
	r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics)

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

	r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwjsoncontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwjsoncontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwjsoncontenttype resource")

	// Create API request body from the model
	// appfwjsoncontenttype := appfwjsoncontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwjsoncontenttype.Type(), &appfwjsoncontenttype)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwjsoncontenttype, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwjsoncontenttype resource")

	// Read the updated state back
	r.readAppfwjsoncontenttypeFromApi(ctx, &data, &resp.Diagnostics)

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

	// For appfwjsoncontenttype, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwjsoncontenttype resource from state")
}

// Helper function to read appfwjsoncontenttype data from API
func (r *AppfwjsoncontenttypeResource) readAppfwjsoncontenttypeFromApi(ctx context.Context, data *AppfwjsoncontenttypeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwjsoncontenttype.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwjsoncontenttype, got error: %s", err))
		return
	}

	appfwjsoncontenttypeSetAttrFromGet(ctx, data, getResponseData)

}
