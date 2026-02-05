package l2param

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
var _ resource.Resource = &L2paramResource{}
var _ resource.ResourceWithConfigure = (*L2paramResource)(nil)
var _ resource.ResourceWithImportState = (*L2paramResource)(nil)

func NewL2paramResource() resource.Resource {
	return &L2paramResource{}
}

// L2paramResource defines the resource implementation.
type L2paramResource struct {
	client *service.NitroClient
}

func (r *L2paramResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *L2paramResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2param"
}

func (r *L2paramResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *L2paramResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data L2paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating l2param resource")

	// l2param := l2paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create l2param, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("l2param-config")

	tflog.Trace(ctx, "Created l2param resource")

	// Read the updated state back
	r.readL2paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L2paramResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data L2paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading l2param resource")

	r.readL2paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L2paramResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data L2paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating l2param resource")

	// Create API request body from the model
	// l2param := l2paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.L2param.Type(), &l2param)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update l2param, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated l2param resource")

	// Read the updated state back
	r.readL2paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L2paramResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data L2paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting l2param resource")

	// For l2param, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted l2param resource from state")
}

// Helper function to read l2param data from API
func (r *L2paramResource) readL2paramFromApi(ctx context.Context, data *L2paramResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.L2param.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read l2param, got error: %s", err))
		return
	}

	l2paramSetAttrFromGet(ctx, data, getResponseData)

}
