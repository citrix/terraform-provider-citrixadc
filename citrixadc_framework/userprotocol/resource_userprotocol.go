package userprotocol

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
var _ resource.Resource = &UserprotocolResource{}
var _ resource.ResourceWithConfigure = (*UserprotocolResource)(nil)
var _ resource.ResourceWithImportState = (*UserprotocolResource)(nil)

func NewUserprotocolResource() resource.Resource {
	return &UserprotocolResource{}
}

// UserprotocolResource defines the resource implementation.
type UserprotocolResource struct {
	client *service.NitroClient
}

func (r *UserprotocolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *UserprotocolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_userprotocol"
}

func (r *UserprotocolResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *UserprotocolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data UserprotocolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating userprotocol resource")

	// userprotocol := userprotocolGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Userprotocol.Type(), &userprotocol)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create userprotocol, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("userprotocol-config")

	tflog.Trace(ctx, "Created userprotocol resource")

	// Read the updated state back
	r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserprotocolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data UserprotocolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading userprotocol resource")

	r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserprotocolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data UserprotocolResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating userprotocol resource")

	// Create API request body from the model
	// userprotocol := userprotocolGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Userprotocol.Type(), &userprotocol)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update userprotocol, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated userprotocol resource")

	// Read the updated state back
	r.readUserprotocolFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *UserprotocolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data UserprotocolResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting userprotocol resource")

	// For userprotocol, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted userprotocol resource from state")
}

// Helper function to read userprotocol data from API
func (r *UserprotocolResource) readUserprotocolFromApi(ctx context.Context, data *UserprotocolResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Userprotocol.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read userprotocol, got error: %s", err))
		return
	}

	userprotocolSetAttrFromGet(ctx, data, getResponseData)

}
