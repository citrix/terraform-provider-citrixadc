package vridparam

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
var _ resource.Resource = &VridparamResource{}
var _ resource.ResourceWithConfigure = (*VridparamResource)(nil)
var _ resource.ResourceWithImportState = (*VridparamResource)(nil)

func NewVridparamResource() resource.Resource {
	return &VridparamResource{}
}

// VridparamResource defines the resource implementation.
type VridparamResource struct {
	client *service.NitroClient
}

func (r *VridparamResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VridparamResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vridparam"
}

func (r *VridparamResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VridparamResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VridparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vridparam resource")

	// vridparam := vridparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vridparam, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vridparam-config")

	tflog.Trace(ctx, "Created vridparam resource")

	// Read the updated state back
	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VridparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vridparam resource")

	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VridparamResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vridparam resource")

	// Create API request body from the model
	// vridparam := vridparamGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vridparam.Type(), &vridparam)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vridparam, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vridparam resource")

	// Read the updated state back
	r.readVridparamFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VridparamResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VridparamResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vridparam resource")

	// For vridparam, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vridparam resource from state")
}

// Helper function to read vridparam data from API
func (r *VridparamResource) readVridparamFromApi(ctx context.Context, data *VridparamResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vridparam.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vridparam, got error: %s", err))
		return
	}

	vridparamSetAttrFromGet(ctx, data, getResponseData)

}
