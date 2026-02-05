package cmpparameter

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
var _ resource.Resource = &CmpparameterResource{}
var _ resource.ResourceWithConfigure = (*CmpparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CmpparameterResource)(nil)

func NewCmpparameterResource() resource.Resource {
	return &CmpparameterResource{}
}

// CmpparameterResource defines the resource implementation.
type CmpparameterResource struct {
	client *service.NitroClient
}

func (r *CmpparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmpparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmpparameter"
}

func (r *CmpparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmpparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmpparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmpparameter resource")

	// cmpparameter := cmpparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmpparameter.Type(), &cmpparameter)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmpparameter, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cmpparameter-config")

	tflog.Trace(ctx, "Created cmpparameter resource")

	// Read the updated state back
	r.readCmpparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmpparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmpparameter resource")

	r.readCmpparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CmpparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cmpparameter resource")

	// Create API request body from the model
	// cmpparameter := cmpparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmpparameter.Type(), &cmpparameter)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmpparameter, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cmpparameter resource")

	// Read the updated state back
	r.readCmpparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmpparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmpparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmpparameter resource")

	// For cmpparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cmpparameter resource from state")
}

// Helper function to read cmpparameter data from API
func (r *CmpparameterResource) readCmpparameterFromApi(ctx context.Context, data *CmpparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cmpparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmpparameter, got error: %s", err))
		return
	}

	cmpparameterSetAttrFromGet(ctx, data, getResponseData)

}
