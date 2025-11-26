package lbparameter

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
var _ resource.Resource = &LbParameterResource{}
var _ resource.ResourceWithConfigure = (*LbParameterResource)(nil)
var _ resource.ResourceWithImportState = (*LbParameterResource)(nil)

func NewLbParameterResource() resource.Resource {
	return &LbParameterResource{}
}

// LbParameterResource defines the resource implementation.
type LbParameterResource struct {
	client *service.NitroClient
}

func (r *LbParameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbParameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbparameter"
}

func (r *LbParameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbParameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbParameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbparameter resource")

	lbparameter := lbparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbparameter, got error: %s", err))
		return
	}

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbparameter-config")

	tflog.Trace(ctx, "Created lbparameter resource")

	// Read the updated state back
	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbParameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbparameter resource")

	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbParameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbparameter resource")

	// Create API request body from the model
	lbparameter := lbparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	err := r.client.UpdateUnnamedResource(service.Lbparameter.Type(), &lbparameter)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbparameter, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated lbparameter resource")

	// Read the updated state back
	r.readLbParameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbParameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbParameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbparameter resource")

	// For lbparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbparameter resource from state")
}

// Helper function to read lbparameter data from API
func (r *LbParameterResource) readLbParameterFromApi(ctx context.Context, data *LbParameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbparameter, got error: %s", err))
		return
	}

	lbparameterSetAttrFromGet(ctx, data, getResponseData)

}
