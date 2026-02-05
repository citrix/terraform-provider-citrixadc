package vpnclientlessaccesspolicy

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
var _ resource.Resource = &VpnclientlessaccesspolicyResource{}
var _ resource.ResourceWithConfigure = (*VpnclientlessaccesspolicyResource)(nil)
var _ resource.ResourceWithImportState = (*VpnclientlessaccesspolicyResource)(nil)

func NewVpnclientlessaccesspolicyResource() resource.Resource {
	return &VpnclientlessaccesspolicyResource{}
}

// VpnclientlessaccesspolicyResource defines the resource implementation.
type VpnclientlessaccesspolicyResource struct {
	client *service.NitroClient
}

func (r *VpnclientlessaccesspolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnclientlessaccesspolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnclientlessaccesspolicy"
}

func (r *VpnclientlessaccesspolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnclientlessaccesspolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnclientlessaccesspolicy resource")

	// vpnclientlessaccesspolicy := vpnclientlessaccesspolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnclientlessaccesspolicy.Type(), &vpnclientlessaccesspolicy)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnclientlessaccesspolicy, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnclientlessaccesspolicy-config")

	tflog.Trace(ctx, "Created vpnclientlessaccesspolicy resource")

	// Read the updated state back
	r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccesspolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnclientlessaccesspolicy resource")

	r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccesspolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnclientlessaccesspolicy resource")

	// Create API request body from the model
	// vpnclientlessaccesspolicy := vpnclientlessaccesspolicyGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnclientlessaccesspolicy.Type(), &vpnclientlessaccesspolicy)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnclientlessaccesspolicy, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnclientlessaccesspolicy resource")

	// Read the updated state back
	r.readVpnclientlessaccesspolicyFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnclientlessaccesspolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnclientlessaccesspolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnclientlessaccesspolicy resource")

	// For vpnclientlessaccesspolicy, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnclientlessaccesspolicy resource from state")
}

// Helper function to read vpnclientlessaccesspolicy data from API
func (r *VpnclientlessaccesspolicyResource) readVpnclientlessaccesspolicyFromApi(ctx context.Context, data *VpnclientlessaccesspolicyResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnclientlessaccesspolicy.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnclientlessaccesspolicy, got error: %s", err))
		return
	}

	vpnclientlessaccesspolicySetAttrFromGet(ctx, data, getResponseData)

}
