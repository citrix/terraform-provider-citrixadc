package clusternodegroup_lbvserver_binding

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
var _ resource.Resource = &ClusternodegroupLbvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupLbvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupLbvserverBindingResource)(nil)

func NewClusternodegroupLbvserverBindingResource() resource.Resource {
	return &ClusternodegroupLbvserverBindingResource{}
}

// ClusternodegroupLbvserverBindingResource defines the resource implementation.
type ClusternodegroupLbvserverBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupLbvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupLbvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_lbvserver_binding"
}

func (r *ClusternodegroupLbvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupLbvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_lbvserver_binding resource")

	// clusternodegroup_lbvserver_binding := clusternodegroup_lbvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_lbvserver_binding.Type(), &clusternodegroup_lbvserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_lbvserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_lbvserver_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_lbvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupLbvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_lbvserver_binding resource")

	r.readClusternodegroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupLbvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupLbvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_lbvserver_binding resource")

	// Create API request body from the model
	// clusternodegroup_lbvserver_binding := clusternodegroup_lbvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_lbvserver_binding.Type(), &clusternodegroup_lbvserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_lbvserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_lbvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupLbvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupLbvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupLbvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_lbvserver_binding resource")

	// For clusternodegroup_lbvserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_lbvserver_binding resource from state")
}

// Helper function to read clusternodegroup_lbvserver_binding data from API
func (r *ClusternodegroupLbvserverBindingResource) readClusternodegroupLbvserverBindingFromApi(ctx context.Context, data *ClusternodegroupLbvserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_lbvserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_lbvserver_binding, got error: %s", err))
		return
	}

	clusternodegroup_lbvserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
