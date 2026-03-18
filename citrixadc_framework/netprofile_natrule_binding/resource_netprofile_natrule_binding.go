package netprofile_natrule_binding

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
var _ resource.Resource = &NetprofileNatruleBindingResource{}
var _ resource.ResourceWithConfigure = (*NetprofileNatruleBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetprofileNatruleBindingResource)(nil)

func NewNetprofileNatruleBindingResource() resource.Resource {
	return &NetprofileNatruleBindingResource{}
}

// NetprofileNatruleBindingResource defines the resource implementation.
type NetprofileNatruleBindingResource struct {
	client *service.NitroClient
}

func (r *NetprofileNatruleBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetprofileNatruleBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netprofile_natrule_binding"
}

func (r *NetprofileNatruleBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetprofileNatruleBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetprofileNatruleBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netprofile_natrule_binding resource")

	// netprofile_natrule_binding := netprofile_natrule_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netprofile_natrule_binding.Type(), &netprofile_natrule_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netprofile_natrule_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("netprofile_natrule_binding-config")

	tflog.Trace(ctx, "Created netprofile_natrule_binding resource")

	// Read the updated state back
	r.readNetprofileNatruleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileNatruleBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetprofileNatruleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netprofile_natrule_binding resource")

	r.readNetprofileNatruleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileNatruleBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetprofileNatruleBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating netprofile_natrule_binding resource")

	// Create API request body from the model
	// netprofile_natrule_binding := netprofile_natrule_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netprofile_natrule_binding.Type(), &netprofile_natrule_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netprofile_natrule_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated netprofile_natrule_binding resource")

	// Read the updated state back
	r.readNetprofileNatruleBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileNatruleBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetprofileNatruleBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netprofile_natrule_binding resource")

	// For netprofile_natrule_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted netprofile_natrule_binding resource from state")
}

// Helper function to read netprofile_natrule_binding data from API
func (r *NetprofileNatruleBindingResource) readNetprofileNatruleBindingFromApi(ctx context.Context, data *NetprofileNatruleBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Netprofile_natrule_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netprofile_natrule_binding, got error: %s", err))
		return
	}

	netprofile_natrule_bindingSetAttrFromGet(ctx, data, getResponseData)

}
