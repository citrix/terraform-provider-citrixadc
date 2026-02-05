package aaapreauthenticationaction

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
var _ resource.Resource = &AaapreauthenticationactionResource{}
var _ resource.ResourceWithConfigure = (*AaapreauthenticationactionResource)(nil)
var _ resource.ResourceWithImportState = (*AaapreauthenticationactionResource)(nil)

func NewAaapreauthenticationactionResource() resource.Resource {
	return &AaapreauthenticationactionResource{}
}

// AaapreauthenticationactionResource defines the resource implementation.
type AaapreauthenticationactionResource struct {
	client *service.NitroClient
}

func (r *AaapreauthenticationactionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AaapreauthenticationactionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_aaapreauthenticationaction"
}

func (r *AaapreauthenticationactionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AaapreauthenticationactionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating aaapreauthenticationaction resource")

	// aaapreauthenticationaction := aaapreauthenticationactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationaction.Type(), &aaapreauthenticationaction)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create aaapreauthenticationaction, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("aaapreauthenticationaction-config")

	tflog.Trace(ctx, "Created aaapreauthenticationaction resource")

	// Read the updated state back
	r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationactionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading aaapreauthenticationaction resource")

	r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationactionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating aaapreauthenticationaction resource")

	// Create API request body from the model
	// aaapreauthenticationaction := aaapreauthenticationactionGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Aaapreauthenticationaction.Type(), &aaapreauthenticationaction)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update aaapreauthenticationaction, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated aaapreauthenticationaction resource")

	// Read the updated state back
	r.readAaapreauthenticationactionFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AaapreauthenticationactionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AaapreauthenticationactionResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting aaapreauthenticationaction resource")

	// For aaapreauthenticationaction, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted aaapreauthenticationaction resource from state")
}

// Helper function to read aaapreauthenticationaction data from API
func (r *AaapreauthenticationactionResource) readAaapreauthenticationactionFromApi(ctx context.Context, data *AaapreauthenticationactionResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Aaapreauthenticationaction.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read aaapreauthenticationaction, got error: %s", err))
		return
	}

	aaapreauthenticationactionSetAttrFromGet(ctx, data, getResponseData)

}
