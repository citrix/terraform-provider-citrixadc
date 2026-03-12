package mapdomain_mapbmr_binding

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
var _ resource.Resource = &MapdomainMapbmrBindingResource{}
var _ resource.ResourceWithConfigure = (*MapdomainMapbmrBindingResource)(nil)
var _ resource.ResourceWithImportState = (*MapdomainMapbmrBindingResource)(nil)

func NewMapdomainMapbmrBindingResource() resource.Resource {
	return &MapdomainMapbmrBindingResource{}
}

// MapdomainMapbmrBindingResource defines the resource implementation.
type MapdomainMapbmrBindingResource struct {
	client *service.NitroClient
}

func (r *MapdomainMapbmrBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *MapdomainMapbmrBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mapdomain_mapbmr_binding"
}

func (r *MapdomainMapbmrBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *MapdomainMapbmrBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data MapdomainMapbmrBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating mapdomain_mapbmr_binding resource")

	// mapdomain_mapbmr_binding := mapdomain_mapbmr_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapdomain_mapbmr_binding.Type(), &mapdomain_mapbmr_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create mapdomain_mapbmr_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("mapdomain_mapbmr_binding-config")

	tflog.Trace(ctx, "Created mapdomain_mapbmr_binding resource")

	// Read the updated state back
	r.readMapdomainMapbmrBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdomainMapbmrBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data MapdomainMapbmrBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading mapdomain_mapbmr_binding resource")

	r.readMapdomainMapbmrBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdomainMapbmrBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data MapdomainMapbmrBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating mapdomain_mapbmr_binding resource")

	// Create API request body from the model
	// mapdomain_mapbmr_binding := mapdomain_mapbmr_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Mapdomain_mapbmr_binding.Type(), &mapdomain_mapbmr_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update mapdomain_mapbmr_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated mapdomain_mapbmr_binding resource")

	// Read the updated state back
	r.readMapdomainMapbmrBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *MapdomainMapbmrBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data MapdomainMapbmrBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting mapdomain_mapbmr_binding resource")

	// For mapdomain_mapbmr_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted mapdomain_mapbmr_binding resource from state")
}

// Helper function to read mapdomain_mapbmr_binding data from API
func (r *MapdomainMapbmrBindingResource) readMapdomainMapbmrBindingFromApi(ctx context.Context, data *MapdomainMapbmrBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Mapdomain_mapbmr_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read mapdomain_mapbmr_binding, got error: %s", err))
		return
	}

	mapdomain_mapbmr_bindingSetAttrFromGet(ctx, data, getResponseData)

}
