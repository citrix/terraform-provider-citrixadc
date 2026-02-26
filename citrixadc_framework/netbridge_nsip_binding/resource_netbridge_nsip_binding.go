package netbridge_nsip_binding

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
var _ resource.Resource = &NetbridgeNsipBindingResource{}
var _ resource.ResourceWithConfigure = (*NetbridgeNsipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetbridgeNsipBindingResource)(nil)

func NewNetbridgeNsipBindingResource() resource.Resource {
	return &NetbridgeNsipBindingResource{}
}

// NetbridgeNsipBindingResource defines the resource implementation.
type NetbridgeNsipBindingResource struct {
	client *service.NitroClient
}

func (r *NetbridgeNsipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetbridgeNsipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netbridge_nsip_binding"
}

func (r *NetbridgeNsipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetbridgeNsipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetbridgeNsipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netbridge_nsip_binding resource")

	// netbridge_nsip_binding := netbridge_nsip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_nsip_binding.Type(), &netbridge_nsip_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netbridge_nsip_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("netbridge_nsip_binding-config")

	tflog.Trace(ctx, "Created netbridge_nsip_binding resource")

	// Read the updated state back
	r.readNetbridgeNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetbridgeNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netbridge_nsip_binding resource")

	r.readNetbridgeNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetbridgeNsipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating netbridge_nsip_binding resource")

	// Create API request body from the model
	// netbridge_nsip_binding := netbridge_nsip_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netbridge_nsip_binding.Type(), &netbridge_nsip_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netbridge_nsip_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated netbridge_nsip_binding resource")

	// Read the updated state back
	r.readNetbridgeNsipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetbridgeNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netbridge_nsip_binding resource")

	// For netbridge_nsip_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted netbridge_nsip_binding resource from state")
}

// Helper function to read netbridge_nsip_binding data from API
func (r *NetbridgeNsipBindingResource) readNetbridgeNsipBindingFromApi(ctx context.Context, data *NetbridgeNsipBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Netbridge_nsip_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netbridge_nsip_binding, got error: %s", err))
		return
	}

	netbridge_nsip_bindingSetAttrFromGet(ctx, data, getResponseData)

}
