package gslbvserver_spilloverpolicy_binding

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
var _ resource.Resource = &GslbvserverSpilloverpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverSpilloverpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverSpilloverpolicyBindingResource)(nil)

func NewGslbvserverSpilloverpolicyBindingResource() resource.Resource {
	return &GslbvserverSpilloverpolicyBindingResource{}
}

// GslbvserverSpilloverpolicyBindingResource defines the resource implementation.
type GslbvserverSpilloverpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *GslbvserverSpilloverpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverSpilloverpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver_spilloverpolicy_binding"
}

func (r *GslbvserverSpilloverpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverSpilloverpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverSpilloverpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver_spilloverpolicy_binding resource")

	// gslbvserver_spilloverpolicy_binding := gslbvserver_spilloverpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_spilloverpolicy_binding.Type(), &gslbvserver_spilloverpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver_spilloverpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbvserver_spilloverpolicy_binding-config")

	tflog.Trace(ctx, "Created gslbvserver_spilloverpolicy_binding resource")

	// Read the updated state back
	r.readGslbvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverSpilloverpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverSpilloverpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver_spilloverpolicy_binding resource")

	r.readGslbvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverSpilloverpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbvserverSpilloverpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbvserver_spilloverpolicy_binding resource")

	// Create API request body from the model
	// gslbvserver_spilloverpolicy_binding := gslbvserver_spilloverpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbvserver_spilloverpolicy_binding.Type(), &gslbvserver_spilloverpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver_spilloverpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbvserver_spilloverpolicy_binding resource")

	// Read the updated state back
	r.readGslbvserverSpilloverpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverSpilloverpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverSpilloverpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver_spilloverpolicy_binding resource")

	// For gslbvserver_spilloverpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbvserver_spilloverpolicy_binding resource from state")
}

// Helper function to read gslbvserver_spilloverpolicy_binding data from API
func (r *GslbvserverSpilloverpolicyBindingResource) readGslbvserverSpilloverpolicyBindingFromApi(ctx context.Context, data *GslbvserverSpilloverpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbvserver_spilloverpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver_spilloverpolicy_binding, got error: %s", err))
		return
	}

	gslbvserver_spilloverpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
