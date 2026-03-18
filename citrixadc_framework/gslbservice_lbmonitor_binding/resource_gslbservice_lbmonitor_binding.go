package gslbservice_lbmonitor_binding

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
var _ resource.Resource = &GslbserviceLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbserviceLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbserviceLbmonitorBindingResource)(nil)

func NewGslbserviceLbmonitorBindingResource() resource.Resource {
	return &GslbserviceLbmonitorBindingResource{}
}

// GslbserviceLbmonitorBindingResource defines the resource implementation.
type GslbserviceLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *GslbserviceLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbserviceLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice_lbmonitor_binding"
}

func (r *GslbserviceLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbserviceLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbserviceLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservice_lbmonitor_binding resource")

	// gslbservice_lbmonitor_binding := gslbservice_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservice_lbmonitor_binding.Type(), &gslbservice_lbmonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservice_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbservice_lbmonitor_binding-config")

	tflog.Trace(ctx, "Created gslbservice_lbmonitor_binding resource")

	// Read the updated state back
	r.readGslbserviceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbserviceLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservice_lbmonitor_binding resource")

	r.readGslbserviceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbserviceLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservice_lbmonitor_binding resource")

	// Create API request body from the model
	// gslbservice_lbmonitor_binding := gslbservice_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservice_lbmonitor_binding.Type(), &gslbservice_lbmonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservice_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbservice_lbmonitor_binding resource")

	// Read the updated state back
	r.readGslbserviceLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbserviceLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservice_lbmonitor_binding resource")

	// For gslbservice_lbmonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbservice_lbmonitor_binding resource from state")
}

// Helper function to read gslbservice_lbmonitor_binding data from API
func (r *GslbserviceLbmonitorBindingResource) readGslbserviceLbmonitorBindingFromApi(ctx context.Context, data *GslbserviceLbmonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbservice_lbmonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice_lbmonitor_binding, got error: %s", err))
		return
	}

	gslbservice_lbmonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
