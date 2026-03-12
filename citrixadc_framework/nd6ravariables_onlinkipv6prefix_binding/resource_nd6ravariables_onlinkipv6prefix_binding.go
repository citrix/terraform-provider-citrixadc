package nd6ravariables_onlinkipv6prefix_binding

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
var _ resource.Resource = &Nd6ravariablesOnlinkipv6prefixBindingResource{}
var _ resource.ResourceWithConfigure = (*Nd6ravariablesOnlinkipv6prefixBindingResource)(nil)
var _ resource.ResourceWithImportState = (*Nd6ravariablesOnlinkipv6prefixBindingResource)(nil)

func NewNd6ravariablesOnlinkipv6prefixBindingResource() resource.Resource {
	return &Nd6ravariablesOnlinkipv6prefixBindingResource{}
}

// Nd6ravariablesOnlinkipv6prefixBindingResource defines the resource implementation.
type Nd6ravariablesOnlinkipv6prefixBindingResource struct {
	client *service.NitroClient
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nd6ravariables_onlinkipv6prefix_binding"
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nd6ravariables_onlinkipv6prefix_binding resource")

	// nd6ravariables_onlinkipv6prefix_binding := nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), &nd6ravariables_onlinkipv6prefix_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("nd6ravariables_onlinkipv6prefix_binding-config")

	tflog.Trace(ctx, "Created nd6ravariables_onlinkipv6prefix_binding resource")

	// Read the updated state back
	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nd6ravariables_onlinkipv6prefix_binding resource")

	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nd6ravariables_onlinkipv6prefix_binding resource")

	// Create API request body from the model
	// nd6ravariables_onlinkipv6prefix_binding := nd6ravariables_onlinkipv6prefix_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), &nd6ravariables_onlinkipv6prefix_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated nd6ravariables_onlinkipv6prefix_binding resource")

	// Read the updated state back
	r.readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nd6ravariablesOnlinkipv6prefixBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nd6ravariables_onlinkipv6prefix_binding resource")

	// For nd6ravariables_onlinkipv6prefix_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted nd6ravariables_onlinkipv6prefix_binding resource from state")
}

// Helper function to read nd6ravariables_onlinkipv6prefix_binding data from API
func (r *Nd6ravariablesOnlinkipv6prefixBindingResource) readNd6ravariablesOnlinkipv6prefixBindingFromApi(ctx context.Context, data *Nd6ravariablesOnlinkipv6prefixBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Nd6ravariables_onlinkipv6prefix_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nd6ravariables_onlinkipv6prefix_binding, got error: %s", err))
		return
	}

	nd6ravariables_onlinkipv6prefix_bindingSetAttrFromGet(ctx, data, getResponseData)

}
