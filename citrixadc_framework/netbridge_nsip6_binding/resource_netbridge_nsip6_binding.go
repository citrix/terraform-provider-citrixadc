package netbridge_nsip6_binding

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
var _ resource.Resource = &NetbridgeNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*NetbridgeNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetbridgeNsip6BindingResource)(nil)

func NewNetbridgeNsip6BindingResource() resource.Resource {
	return &NetbridgeNsip6BindingResource{}
}

// NetbridgeNsip6BindingResource defines the resource implementation.
type NetbridgeNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *NetbridgeNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetbridgeNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netbridge_nsip6_binding"
}

func (r *NetbridgeNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetbridgeNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netbridge_nsip6_binding resource")

	// netbridge_nsip6_binding := netbridge_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_nsip6_binding.Type(), &netbridge_nsip6_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netbridge_nsip6_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("netbridge_nsip6_binding-config")

	tflog.Trace(ctx, "Created netbridge_nsip6_binding resource")

	// Read the updated state back
	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netbridge_nsip6_binding resource")

	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating netbridge_nsip6_binding resource")

	// Create API request body from the model
	// netbridge_nsip6_binding := netbridge_nsip6_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_nsip6_binding.Type(), &netbridge_nsip6_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netbridge_nsip6_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated netbridge_nsip6_binding resource")

	// Read the updated state back
	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netbridge_nsip6_binding resource")

	// For netbridge_nsip6_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted netbridge_nsip6_binding resource from state")
}

// Helper function to read netbridge_nsip6_binding data from API
func (r *NetbridgeNsip6BindingResource) readNetbridgeNsip6BindingFromApi(ctx context.Context, data *NetbridgeNsip6BindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Netbridge_nsip6_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netbridge_nsip6_binding, got error: %s", err))
		return
	}

	netbridge_nsip6_bindingSetAttrFromGet(ctx, data, getResponseData)

}
