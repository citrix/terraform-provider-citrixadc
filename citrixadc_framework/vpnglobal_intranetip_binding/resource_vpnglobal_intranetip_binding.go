package vpnglobal_intranetip_binding

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
var _ resource.Resource = &VpnglobalIntranetipBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnglobalIntranetipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnglobalIntranetipBindingResource)(nil)

func NewVpnglobalIntranetipBindingResource() resource.Resource {
	return &VpnglobalIntranetipBindingResource{}
}

// VpnglobalIntranetipBindingResource defines the resource implementation.
type VpnglobalIntranetipBindingResource struct {
	client *service.NitroClient
}

func (r *VpnglobalIntranetipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnglobalIntranetipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnglobal_intranetip_binding"
}

func (r *VpnglobalIntranetipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnglobalIntranetipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnglobalIntranetipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnglobal_intranetip_binding resource")

	// vpnglobal_intranetip_binding := vpnglobal_intranetip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_intranetip_binding.Type(), &vpnglobal_intranetip_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnglobal_intranetip_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("vpnglobal_intranetip_binding-config")

	tflog.Trace(ctx, "Created vpnglobal_intranetip_binding resource")

	// Read the updated state back
	r.readVpnglobalIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnglobalIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnglobal_intranetip_binding resource")

	r.readVpnglobalIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data VpnglobalIntranetipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating vpnglobal_intranetip_binding resource")

	// Create API request body from the model
	// vpnglobal_intranetip_binding := vpnglobal_intranetip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Vpnglobal_intranetip_binding.Type(), &vpnglobal_intranetip_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnglobal_intranetip_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated vpnglobal_intranetip_binding resource")

	// Read the updated state back
	r.readVpnglobalIntranetipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnglobalIntranetipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnglobalIntranetipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnglobal_intranetip_binding resource")

	// For vpnglobal_intranetip_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted vpnglobal_intranetip_binding resource from state")
}

// Helper function to read vpnglobal_intranetip_binding data from API
func (r *VpnglobalIntranetipBindingResource) readVpnglobalIntranetipBindingFromApi(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Vpnglobal_intranetip_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnglobal_intranetip_binding, got error: %s", err))
		return
	}

	vpnglobal_intranetip_bindingSetAttrFromGet(ctx, data, getResponseData)

}
