package vxlan

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
var _ resource.Resource = &VxlanResource{}
var _ resource.ResourceWithConfigure = (*VxlanResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanResource)(nil)

func NewVxlanResource() resource.Resource {
	return &VxlanResource{}
}

// VxlanResource defines the resource implementation.
type VxlanResource struct {
	client *service.NitroClient
}

func (r *VxlanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan"
}

func (r *VxlanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan resource")

	// vxlan := vxlanGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan.Type(), &vxlan)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vxlan-config")

	tflog.Trace(ctx, "Created vxlan resource")

	// Read the updated state back
	r.readVxlanFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan resource")

	r.readVxlanFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VxlanResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vxlan resource")

	// Create API request body from the model
	// vxlan := vxlanGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vxlan.Type(), &vxlan)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vxlan, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vxlan resource")

	// Read the updated state back
	r.readVxlanFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan resource")

	// For vxlan, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vxlan resource from state")
}

// Helper function to read vxlan data from API
func (r *VxlanResource) readVxlanFromApi(ctx context.Context, data *VxlanResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vxlan.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan, got error: %s", err))
		return
	}

	vxlanSetAttrFromGet(ctx, data, getResponseData)

}
