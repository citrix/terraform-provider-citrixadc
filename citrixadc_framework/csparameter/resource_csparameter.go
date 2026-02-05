package csparameter

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
var _ resource.Resource = &CsparameterResource{}
var _ resource.ResourceWithConfigure = (*CsparameterResource)(nil)
var _ resource.ResourceWithImportState = (*CsparameterResource)(nil)

func NewCsparameterResource() resource.Resource {
	return &CsparameterResource{}
}

// CsparameterResource defines the resource implementation.
type CsparameterResource struct {
	client *service.NitroClient
}

func (r *CsparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csparameter"
}

func (r *CsparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csparameter resource")

	// csparameter := csparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csparameter.Type(), &csparameter)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csparameter, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csparameter-config")

	tflog.Trace(ctx, "Created csparameter resource")

	// Read the updated state back
	r.readCsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csparameter resource")

	r.readCsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csparameter resource")

	// Create API request body from the model
	// csparameter := csparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csparameter.Type(), &csparameter)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csparameter, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csparameter resource")

	// Read the updated state back
	r.readCsparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csparameter resource")

	// For csparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csparameter resource from state")
}

// Helper function to read csparameter data from API
func (r *CsparameterResource) readCsparameterFromApi(ctx context.Context, data *CsparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csparameter, got error: %s", err))
		return
	}

	csparameterSetAttrFromGet(ctx, data, getResponseData)

}
