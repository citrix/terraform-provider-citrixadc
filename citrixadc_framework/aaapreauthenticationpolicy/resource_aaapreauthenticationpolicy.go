package aaapreauthenticationpolicy

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
var _ resource.Resource = &AaapreauthenticationpolicyResource{}
var _ resource.ResourceWithConfigure = (*AaapreauthenticationpolicyResource)(nil)
var _ resource.ResourceWithImportState = (*AaapreauthenticationpolicyResource)(nil)

func NewAaapreauthenticationpolicyResource() resource.Resource {
	return &AaapreauthenticationpolicyResource{}
}

// AaapreauthenticationpolicyResource defines the resource implementation.
type AaapreauthenticationpolicyResource struct {
	client *service.NitroClient
}

func (r *AaapreauthenticationpolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaapreauthenticationpolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapreauthenticationpolicy"
}

func (r *AaapreauthenticationpolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaapreauthenticationpolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaapreauthenticationpolicy resource")

	// aaapreauthenticationpolicy := aaapreauthenticationpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationpolicy.Type(), &aaapreauthenticationpolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaapreauthenticationpolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaapreauthenticationpolicy-config")

	tflog.Trace(ctx, "Created aaapreauthenticationpolicy resource")

	// Read the updated state back
	r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationpolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaapreauthenticationpolicy resource")

	r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationpolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaapreauthenticationpolicy resource")

	// Create API request body from the model
	// aaapreauthenticationpolicy := aaapreauthenticationpolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationpolicy.Type(), &aaapreauthenticationpolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaapreauthenticationpolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaapreauthenticationpolicy resource")

	// Read the updated state back
	r.readAaapreauthenticationpolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationpolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaapreauthenticationpolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaapreauthenticationpolicy resource")

	// For aaapreauthenticationpolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaapreauthenticationpolicy resource from state")
}

// Helper function to read aaapreauthenticationpolicy data from API
func (r *AaapreauthenticationpolicyResource) readAaapreauthenticationpolicyFromApi(ctx context.Context, data *AaapreauthenticationpolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaapreauthenticationpolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaapreauthenticationpolicy, got error: %s", err))
		return
	}

	aaapreauthenticationpolicySetAttrFromGet(ctx, data, getResponseData)

}
