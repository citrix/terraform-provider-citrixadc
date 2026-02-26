package gslbvserver_domain_binding

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
var _ resource.Resource = &GslbvserverDomainBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverDomainBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverDomainBindingResource)(nil)

func NewGslbvserverDomainBindingResource() resource.Resource {
	return &GslbvserverDomainBindingResource{}
}

// GslbvserverDomainBindingResource defines the resource implementation.
type GslbvserverDomainBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverDomainBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverDomainBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_domain_binding"
}

func (r *GslbvserverDomainBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverDomainBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_domain_binding resource")

	// gslbvserver_domain_binding := gslbvserver_domain_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_domain_binding.Type(), &gslbvserver_domain_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_domain_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbvserver_domain_binding-config")

	tflog.Trace(ctx, "Created gslbvserver_domain_binding resource")

	// Read the updated state back
	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_domain_binding resource")

	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbvserver_domain_binding resource")

	// Create API request body from the model
	// gslbvserver_domain_binding := gslbvserver_domain_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_domain_binding.Type(), &gslbvserver_domain_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_domain_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbvserver_domain_binding resource")

	// Read the updated state back
	r.readGslbvserverDomainBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverDomainBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverDomainBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_domain_binding resource")

	// For gslbvserver_domain_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbvserver_domain_binding resource from state")
}

// Helper function to read gslbvserver_domain_binding data from API
func (r *GslbvserverDomainBindingResource) readGslbvserverDomainBindingFromApi(ctx context.Context, data *GslbvserverDomainBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbvserver_domain_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_domain_binding, got error: %s", err))
		return
	}

	gslbvserver_domain_bindingSetAttrFromGet(ctx, data, getResponseData)

}
