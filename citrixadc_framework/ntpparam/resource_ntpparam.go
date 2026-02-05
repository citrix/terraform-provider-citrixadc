package ntpparam

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
var _ resource.Resource = &NtpparamResource{}
var _ resource.ResourceWithConfigure = (*NtpparamResource)(nil)
var _ resource.ResourceWithImportState = (*NtpparamResource)(nil)

func NewNtpparamResource() resource.Resource {
	return &NtpparamResource{}
}

// NtpparamResource defines the resource implementation.
type NtpparamResource struct {
	client *service.NitroClient
}

func (r *NtpparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NtpparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ntpparam"
}

func (r *NtpparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NtpparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NtpparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating ntpparam resource")

	// ntpparam := ntpparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create ntpparam, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("ntpparam-config")

	tflog.Trace(ctx, "Created ntpparam resource")

	// Read the updated state back
	r.readNtpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NtpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading ntpparam resource")

	r.readNtpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NtpparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating ntpparam resource")

	// Create API request body from the model
	// ntpparam := ntpparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Ntpparam.Type(), &ntpparam)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update ntpparam, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated ntpparam resource")

	// Read the updated state back
	r.readNtpparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NtpparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NtpparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting ntpparam resource")

	// For ntpparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted ntpparam resource from state")
}

// Helper function to read ntpparam data from API
func (r *NtpparamResource) readNtpparamFromApi(ctx context.Context, data *NtpparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Ntpparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read ntpparam, got error: %s", err))
		return
	}

	ntpparamSetAttrFromGet(ctx, data, getResponseData)

}
