package appflowpolicylabel

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
var _ resource.Resource = &AppflowpolicylabelResource{}
var _ resource.ResourceWithConfigure = (*AppflowpolicylabelResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowpolicylabelResource)(nil)

func NewAppflowpolicylabelResource() resource.Resource {
	return &AppflowpolicylabelResource{}
}

// AppflowpolicylabelResource defines the resource implementation.
type AppflowpolicylabelResource struct {
	client *service.NitroClient
}

func (r *AppflowpolicylabelResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowpolicylabelResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowpolicylabel"
}

func (r *AppflowpolicylabelResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowpolicylabelResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowpolicylabel resource")

	// appflowpolicylabel := appflowpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel.Type(), &appflowpolicylabel)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowpolicylabel, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appflowpolicylabel-config")

	tflog.Trace(ctx, "Created appflowpolicylabel resource")

	// Read the updated state back
	r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowpolicylabel resource")

	r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appflowpolicylabel resource")

	// Create API request body from the model
	// appflowpolicylabel := appflowpolicylabelGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowpolicylabel.Type(), &appflowpolicylabel)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowpolicylabel, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appflowpolicylabel resource")

	// Read the updated state back
	r.readAppflowpolicylabelFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowpolicylabelResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowpolicylabelResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowpolicylabel resource")

	// For appflowpolicylabel, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appflowpolicylabel resource from state")
}

// Helper function to read appflowpolicylabel data from API
func (r *AppflowpolicylabelResource) readAppflowpolicylabelFromApi(ctx context.Context, data *AppflowpolicylabelResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appflowpolicylabel.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowpolicylabel, got error: %s", err))
		return
	}

	appflowpolicylabelSetAttrFromGet(ctx, data, getResponseData)

}
