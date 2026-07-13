package clusternodegroup_authenticationvserver_binding

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
var _ resource.Resource = &ClusternodegroupAuthenticationvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupAuthenticationvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupAuthenticationvserverBindingResource)(nil)

func NewClusternodegroupAuthenticationvserverBindingResource() resource.Resource {
	return &ClusternodegroupAuthenticationvserverBindingResource{}
}

// ClusternodegroupAuthenticationvserverBindingResource defines the resource implementation.
type ClusternodegroupAuthenticationvserverBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_authenticationvserver_binding"
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupAuthenticationvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_authenticationvserver_binding resource")

	// clusternodegroup_authenticationvserver_binding := clusternodegroup_authenticationvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_authenticationvserver_binding.Type(), &clusternodegroup_authenticationvserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_authenticationvserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_authenticationvserver_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_authenticationvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupAuthenticationvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupAuthenticationvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_authenticationvserver_binding resource")

	r.readClusternodegroupAuthenticationvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupAuthenticationvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_authenticationvserver_binding resource")

	// Create API request body from the model
	// clusternodegroup_authenticationvserver_binding := clusternodegroup_authenticationvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_authenticationvserver_binding.Type(), &clusternodegroup_authenticationvserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_authenticationvserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_authenticationvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupAuthenticationvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupAuthenticationvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupAuthenticationvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_authenticationvserver_binding resource")

	// For clusternodegroup_authenticationvserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_authenticationvserver_binding resource from state")
}

// Helper function to read clusternodegroup_authenticationvserver_binding data from API
func (r *ClusternodegroupAuthenticationvserverBindingResource) readClusternodegroupAuthenticationvserverBindingFromApi(ctx context.Context, data *ClusternodegroupAuthenticationvserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_authenticationvserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_authenticationvserver_binding, got error: %s", err))
		return
	}

	clusternodegroup_authenticationvserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
