package aaapreauthenticationparameter

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
var _ resource.Resource = &AaapreauthenticationparameterResource{}
var _ resource.ResourceWithConfigure = (*AaapreauthenticationparameterResource)(nil)
var _ resource.ResourceWithImportState = (*AaapreauthenticationparameterResource)(nil)

func NewAaapreauthenticationparameterResource() resource.Resource {
	return &AaapreauthenticationparameterResource{}
}

// AaapreauthenticationparameterResource defines the resource implementation.
type AaapreauthenticationparameterResource struct {
	client *service.NitroClient
}

func (r *AaapreauthenticationparameterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaapreauthenticationparameterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapreauthenticationparameter"
}

func (r *AaapreauthenticationparameterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaapreauthenticationparameterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaapreauthenticationparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaapreauthenticationparameter resource")

	// aaapreauthenticationparameter := aaapreauthenticationparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaapreauthenticationparameter, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaapreauthenticationparameter-config")

	tflog.Trace(ctx, "Created aaapreauthenticationparameter resource")

	// Read the updated state back
	r.readAaapreauthenticationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationparameterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaapreauthenticationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaapreauthenticationparameter resource")

	r.readAaapreauthenticationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationparameterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaapreauthenticationparameterResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaapreauthenticationparameter resource")

	// Create API request body from the model
	// aaapreauthenticationparameter := aaapreauthenticationparameterGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationparameter.Type(), &aaapreauthenticationparameter)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaapreauthenticationparameter, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaapreauthenticationparameter resource")

	// Read the updated state back
	r.readAaapreauthenticationparameterFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationparameterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaapreauthenticationparameterResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaapreauthenticationparameter resource")

	// For aaapreauthenticationparameter, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaapreauthenticationparameter resource from state")
}

// Helper function to read aaapreauthenticationparameter data from API
func (r *AaapreauthenticationparameterResource) readAaapreauthenticationparameterFromApi(ctx context.Context, data *AaapreauthenticationparameterResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaapreauthenticationparameter.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaapreauthenticationparameter, got error: %s", err))
		return
	}

	aaapreauthenticationparameterSetAttrFromGet(ctx, data, getResponseData)

}
