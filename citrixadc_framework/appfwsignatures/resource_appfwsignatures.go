package appfwsignatures

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
var _ resource.Resource = &AppfwsignaturesResource{}
var _ resource.ResourceWithConfigure = (*AppfwsignaturesResource)(nil)
var _ resource.ResourceWithImportState = (*AppfwsignaturesResource)(nil)

func NewAppfwsignaturesResource() resource.Resource {
	return &AppfwsignaturesResource{}
}

// AppfwsignaturesResource defines the resource implementation.
type AppfwsignaturesResource struct {
	client *service.NitroClient
}

func (r *AppfwsignaturesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppfwsignaturesResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwsignatures"
}

func (r *AppfwsignaturesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppfwsignaturesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appfwsignatures resource")

	// appfwsignatures := appfwsignaturesGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwsignatures.Type(), &appfwsignatures)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appfwsignatures, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appfwsignatures-config")

	tflog.Trace(ctx, "Created appfwsignatures resource")

	// Read the updated state back
	r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsignaturesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appfwsignatures resource")

	r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsignaturesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appfwsignatures resource")

	// Create API request body from the model
	// appfwsignatures := appfwsignaturesGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appfwsignatures.Type(), &appfwsignatures)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appfwsignatures, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appfwsignatures resource")

	// Read the updated state back
	r.readAppfwsignaturesFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppfwsignaturesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppfwsignaturesResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appfwsignatures resource")

	// For appfwsignatures, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appfwsignatures resource from state")
}

// Helper function to read appfwsignatures data from API
func (r *AppfwsignaturesResource) readAppfwsignaturesFromApi(ctx context.Context, data *AppfwsignaturesResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appfwsignatures.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appfwsignatures, got error: %s", err))
		return
	}

	appfwsignaturesSetAttrFromGet(ctx, data, getResponseData)

}
