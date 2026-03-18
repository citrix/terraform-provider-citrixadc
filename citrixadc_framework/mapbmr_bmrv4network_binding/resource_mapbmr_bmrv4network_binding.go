package mapbmr_bmrv4network_binding

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
var _ resource.Resource = &MapbmrBmrv4networkBindingResource{}
var _ resource.ResourceWithConfigure = (*MapbmrBmrv4networkBindingResource)(nil)
var _ resource.ResourceWithImportState = (*MapbmrBmrv4networkBindingResource)(nil)

func NewMapbmrBmrv4networkBindingResource() resource.Resource {
	return &MapbmrBmrv4networkBindingResource{}
}

// MapbmrBmrv4networkBindingResource defines the resource implementation.
type MapbmrBmrv4networkBindingResource struct {
	client *service.NitroClient
}

func (r *MapbmrBmrv4networkBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapbmrBmrv4networkBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapbmr_bmrv4network_binding"
}

func (r *MapbmrBmrv4networkBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapbmrBmrv4networkBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapbmr_bmrv4network_binding resource")

	// mapbmr_bmrv4network_binding := mapbmr_bmrv4network_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapbmr_bmrv4network_binding.Type(), &mapbmr_bmrv4network_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapbmr_bmrv4network_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("mapbmr_bmrv4network_binding-config")

	tflog.Trace(ctx, "Created mapbmr_bmrv4network_binding resource")

	// Read the updated state back
	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapbmr_bmrv4network_binding resource")

	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating mapbmr_bmrv4network_binding resource")

	// Create API request body from the model
	// mapbmr_bmrv4network_binding := mapbmr_bmrv4network_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapbmr_bmrv4network_binding.Type(), &mapbmr_bmrv4network_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update mapbmr_bmrv4network_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated mapbmr_bmrv4network_binding resource")

	// Read the updated state back
	r.readMapbmrBmrv4networkBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapbmrBmrv4networkBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapbmrBmrv4networkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapbmr_bmrv4network_binding resource")

	// For mapbmr_bmrv4network_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted mapbmr_bmrv4network_binding resource from state")
}

// Helper function to read mapbmr_bmrv4network_binding data from API
func (r *MapbmrBmrv4networkBindingResource) readMapbmrBmrv4networkBindingFromApi(ctx context.Context, data *MapbmrBmrv4networkBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Mapbmr_bmrv4network_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapbmr_bmrv4network_binding, got error: %s", err))
		return
	}

	mapbmr_bmrv4network_bindingSetAttrFromGet(ctx, data, getResponseData)

}
