package gslbservice_dnsview_binding

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
var _ resource.Resource = &GslbserviceDnsviewBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbserviceDnsviewBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbserviceDnsviewBindingResource)(nil)

func NewGslbserviceDnsviewBindingResource() resource.Resource {
	return &GslbserviceDnsviewBindingResource{}
}

// GslbserviceDnsviewBindingResource defines the resource implementation.
type GslbserviceDnsviewBindingResource struct {
	client *service.NitroClient
}

func (r *GslbserviceDnsviewBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbserviceDnsviewBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice_dnsview_binding"
}

func (r *GslbserviceDnsviewBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbserviceDnsviewBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservice_dnsview_binding resource")

	// gslbservice_dnsview_binding := gslbservice_dnsview_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservice_dnsview_binding.Type(), &gslbservice_dnsview_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservice_dnsview_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbservice_dnsview_binding-config")

	tflog.Trace(ctx, "Created gslbservice_dnsview_binding resource")

	// Read the updated state back
	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservice_dnsview_binding resource")

	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservice_dnsview_binding resource")

	// Create API request body from the model
	// gslbservice_dnsview_binding := gslbservice_dnsview_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservice_dnsview_binding.Type(), &gslbservice_dnsview_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservice_dnsview_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbservice_dnsview_binding resource")

	// Read the updated state back
	r.readGslbserviceDnsviewBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceDnsviewBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbserviceDnsviewBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservice_dnsview_binding resource")

	// For gslbservice_dnsview_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbservice_dnsview_binding resource from state")
}

// Helper function to read gslbservice_dnsview_binding data from API
func (r *GslbserviceDnsviewBindingResource) readGslbserviceDnsviewBindingFromApi(ctx context.Context, data *GslbserviceDnsviewBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbservice_dnsview_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice_dnsview_binding, got error: %s", err))
		return
	}

	gslbservice_dnsview_bindingSetAttrFromGet(ctx, data, getResponseData)

}
