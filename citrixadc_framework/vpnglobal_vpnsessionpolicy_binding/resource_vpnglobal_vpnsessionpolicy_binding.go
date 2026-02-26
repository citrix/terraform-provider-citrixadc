package vpnglobal_vpnsessionpolicy_binding

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
var _ resource.Resource = &VpnglobalVpnsessionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalVpnsessionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalVpnsessionpolicyBindingResource)(nil)

func NewVpnglobalVpnsessionpolicyBindingResource() resource.Resource {
	return &VpnglobalVpnsessionpolicyBindingResource{}
}

// VpnglobalVpnsessionpolicyBindingResource defines the resource implementation.
type VpnglobalVpnsessionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalVpnsessionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_vpnsessionpolicy_binding"
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_vpnsessionpolicy_binding resource")

	// vpnglobal_vpnsessionpolicy_binding := vpnglobal_vpnsessionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), &vpnglobal_vpnsessionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_vpnsessionpolicy_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_vpnsessionpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_vpnsessionpolicy_binding resource")

	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_vpnsessionpolicy_binding resource")

	// Create API request body from the model
	// vpnglobal_vpnsessionpolicy_binding := vpnglobal_vpnsessionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), &vpnglobal_vpnsessionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_vpnsessionpolicy_binding resource")

	// Read the updated state back
	r.readVpnglobalVpnsessionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalVpnsessionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalVpnsessionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_vpnsessionpolicy_binding resource")

	// For vpnglobal_vpnsessionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_vpnsessionpolicy_binding resource from state")
}

// Helper function to read vpnglobal_vpnsessionpolicy_binding data from API
func (r *VpnglobalVpnsessionpolicyBindingResource) readVpnglobalVpnsessionpolicyBindingFromApi(ctx context.Context, data *VpnglobalVpnsessionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_vpnsessionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_vpnsessionpolicy_binding, got error: %s", err))
		return
	}

	vpnglobal_vpnsessionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
