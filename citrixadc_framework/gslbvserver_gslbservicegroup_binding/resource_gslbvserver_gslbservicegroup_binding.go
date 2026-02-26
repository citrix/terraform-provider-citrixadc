package gslbvserver_gslbservicegroup_binding

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
var _ resource.Resource = &GslbvserverGslbservicegroupBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverGslbservicegroupBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverGslbservicegroupBindingResource)(nil)

func NewGslbvserverGslbservicegroupBindingResource() resource.Resource {
	return &GslbvserverGslbservicegroupBindingResource{}
}

// GslbvserverGslbservicegroupBindingResource defines the resource implementation.
type GslbvserverGslbservicegroupBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverGslbservicegroupBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverGslbservicegroupBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_gslbservicegroup_binding"
}

func (r *GslbvserverGslbservicegroupBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverGslbservicegroupBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverGslbservicegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_gslbservicegroup_binding resource")

	// gslbvserver_gslbservicegroup_binding := gslbvserver_gslbservicegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_gslbservicegroup_binding.Type(), &gslbvserver_gslbservicegroup_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_gslbservicegroup_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbvserver_gslbservicegroup_binding-config")

	tflog.Trace(ctx, "Created gslbvserver_gslbservicegroup_binding resource")

	// Read the updated state back
	r.readGslbvserverGslbservicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbservicegroupBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverGslbservicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_gslbservicegroup_binding resource")

	r.readGslbvserverGslbservicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbservicegroupBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbvserverGslbservicegroupBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbvserver_gslbservicegroup_binding resource")

	// Create API request body from the model
	// gslbvserver_gslbservicegroup_binding := gslbvserver_gslbservicegroup_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_gslbservicegroup_binding.Type(), &gslbvserver_gslbservicegroup_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_gslbservicegroup_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbvserver_gslbservicegroup_binding resource")

	// Read the updated state back
	r.readGslbvserverGslbservicegroupBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverGslbservicegroupBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverGslbservicegroupBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_gslbservicegroup_binding resource")

	// For gslbvserver_gslbservicegroup_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbvserver_gslbservicegroup_binding resource from state")
}

// Helper function to read gslbvserver_gslbservicegroup_binding data from API
func (r *GslbvserverGslbservicegroupBindingResource) readGslbvserverGslbservicegroupBindingFromApi(ctx context.Context, data *GslbvserverGslbservicegroupBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbvserver_gslbservicegroup_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_gslbservicegroup_binding, got error: %s", err))
		return
	}

	gslbvserver_gslbservicegroup_bindingSetAttrFromGet(ctx, data, getResponseData)

}
