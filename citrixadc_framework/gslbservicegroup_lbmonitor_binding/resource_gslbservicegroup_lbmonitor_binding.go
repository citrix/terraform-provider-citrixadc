package gslbservicegroup_lbmonitor_binding

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
var _ resource.Resource = &GslbservicegroupLbmonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupLbmonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupLbmonitorBindingResource)(nil)

func NewGslbservicegroupLbmonitorBindingResource() resource.Resource {
	return &GslbservicegroupLbmonitorBindingResource{}
}

// GslbservicegroupLbmonitorBindingResource defines the resource implementation.
type GslbservicegroupLbmonitorBindingResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupLbmonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupLbmonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_lbmonitor_binding"
}

func (r *GslbservicegroupLbmonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupLbmonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup_lbmonitor_binding resource")

	// gslbservicegroup_lbmonitor_binding := gslbservicegroup_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_lbmonitor_binding.Type(), &gslbservicegroup_lbmonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbservicegroup_lbmonitor_binding-config")

	tflog.Trace(ctx, "Created gslbservicegroup_lbmonitor_binding resource")

	// Read the updated state back
	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupLbmonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup_lbmonitor_binding resource")

	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupLbmonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservicegroup_lbmonitor_binding resource")

	// Create API request body from the model
	// gslbservicegroup_lbmonitor_binding := gslbservicegroup_lbmonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_lbmonitor_binding.Type(), &gslbservicegroup_lbmonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup_lbmonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbservicegroup_lbmonitor_binding resource")

	// Read the updated state back
	r.readGslbservicegroupLbmonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupLbmonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupLbmonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup_lbmonitor_binding resource")

	// For gslbservicegroup_lbmonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbservicegroup_lbmonitor_binding resource from state")
}

// Helper function to read gslbservicegroup_lbmonitor_binding data from API
func (r *GslbservicegroupLbmonitorBindingResource) readGslbservicegroupLbmonitorBindingFromApi(ctx context.Context, data *GslbservicegroupLbmonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbservicegroup_lbmonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_lbmonitor_binding, got error: %s", err))
		return
	}

	gslbservicegroup_lbmonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
