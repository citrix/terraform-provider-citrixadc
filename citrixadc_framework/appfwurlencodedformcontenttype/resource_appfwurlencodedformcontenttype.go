package appfwurlencodedformcontenttype

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
var _ resource.Resource = &AppfwurlencodedformcontenttypeResource{}
var _ resource.ResourceWithConfigure = (*AppfwurlencodedformcontenttypeResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwurlencodedformcontenttypeResource)(nil)

func NewAppfwurlencodedformcontenttypeResource() resource.Resource {
	return &AppfwurlencodedformcontenttypeResource{}
}

// AppfwurlencodedformcontenttypeResource defines the resource implementation.
type AppfwurlencodedformcontenttypeResource struct {
	client *service.NitroClient
}

func (r *AppfwurlencodedformcontenttypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwurlencodedformcontenttypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwurlencodedformcontenttype"
}

func (r *AppfwurlencodedformcontenttypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwurlencodedformcontenttypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwurlencodedformcontenttype resource")

	// appfwurlencodedformcontenttype := appfwurlencodedformcontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwurlencodedformcontenttype.Type(), &appfwurlencodedformcontenttype)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwurlencodedformcontenttype, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwurlencodedformcontenttype-config")

	tflog.Trace(ctx, "Created appfwurlencodedformcontenttype resource")

	// Read the updated state back
	r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwurlencodedformcontenttypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwurlencodedformcontenttype resource")

	r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwurlencodedformcontenttypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwurlencodedformcontenttype resource")

	// Create API request body from the model
	// appfwurlencodedformcontenttype := appfwurlencodedformcontenttypeGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwurlencodedformcontenttype.Type(), &appfwurlencodedformcontenttype)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwurlencodedformcontenttype, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwurlencodedformcontenttype resource")

	// Read the updated state back
	r.readAppfwurlencodedformcontenttypeFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwurlencodedformcontenttypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwurlencodedformcontenttypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwurlencodedformcontenttype resource")

	// For appfwurlencodedformcontenttype, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwurlencodedformcontenttype resource from state")
}

// Helper function to read appfwurlencodedformcontenttype data from API
func (r *AppfwurlencodedformcontenttypeResource) readAppfwurlencodedformcontenttypeFromApi(ctx context.Context, data *AppfwurlencodedformcontenttypeResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwurlencodedformcontenttype.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwurlencodedformcontenttype, got error: %s", err))
		return
	}

	appfwurlencodedformcontenttypeSetAttrFromGet(ctx, data, getResponseData)

}
