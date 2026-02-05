package l3param

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
var _ resource.Resource = &L3paramResource{}
var _ resource.ResourceWithConfigure = (*L3paramResource)(nil)
var _ resource.ResourceWithImportState = (*L3paramResource)(nil)

func NewL3paramResource() resource.Resource {
	return &L3paramResource{}
}

// L3paramResource defines the resource implementation.
type L3paramResource struct {
	client *service.NitroClient
}

func (r *L3paramResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *L3paramResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l3param"
}

func (r *L3paramResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *L3paramResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data L3paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating l3param resource")

	// l3param := l3paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create l3param, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("l3param-config")

	tflog.Trace(ctx, "Created l3param resource")

	// Read the updated state back
	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data L3paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading l3param resource")

	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data L3paramResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating l3param resource")

	// Create API request body from the model
	// l3param := l3paramGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.L3param.Type(), &l3param)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update l3param, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated l3param resource")

	// Read the updated state back
	r.readL3paramFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *L3paramResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data L3paramResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting l3param resource")

	// For l3param, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted l3param resource from state")
}

// Helper function to read l3param data from API
func (r *L3paramResource) readL3paramFromApi(ctx context.Context, data *L3paramResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.L3param.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read l3param, got error: %s", err))
		return
	}

	l3paramSetAttrFromGet(ctx, data, getResponseData)

}
