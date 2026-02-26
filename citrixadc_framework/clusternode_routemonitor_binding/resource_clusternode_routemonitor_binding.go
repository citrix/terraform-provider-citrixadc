package clusternode_routemonitor_binding

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
var _ resource.Resource = &ClusternodeRoutemonitorBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodeRoutemonitorBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodeRoutemonitorBindingResource)(nil)

func NewClusternodeRoutemonitorBindingResource() resource.Resource {
	return &ClusternodeRoutemonitorBindingResource{}
}

// ClusternodeRoutemonitorBindingResource defines the resource implementation.
type ClusternodeRoutemonitorBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodeRoutemonitorBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodeRoutemonitorBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternode_routemonitor_binding"
}

func (r *ClusternodeRoutemonitorBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodeRoutemonitorBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternode_routemonitor_binding resource")

	// clusternode_routemonitor_binding := clusternode_routemonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternode_routemonitor_binding.Type(), &clusternode_routemonitor_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternode_routemonitor_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternode_routemonitor_binding-config")

	tflog.Trace(ctx, "Created clusternode_routemonitor_binding resource")

	// Read the updated state back
	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternode_routemonitor_binding resource")

	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternode_routemonitor_binding resource")

	// Create API request body from the model
	// clusternode_routemonitor_binding := clusternode_routemonitor_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternode_routemonitor_binding.Type(), &clusternode_routemonitor_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternode_routemonitor_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternode_routemonitor_binding resource")

	// Read the updated state back
	r.readClusternodeRoutemonitorBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodeRoutemonitorBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodeRoutemonitorBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternode_routemonitor_binding resource")

	// For clusternode_routemonitor_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternode_routemonitor_binding resource from state")
}

// Helper function to read clusternode_routemonitor_binding data from API
func (r *ClusternodeRoutemonitorBindingResource) readClusternodeRoutemonitorBindingFromApi(ctx context.Context, data *ClusternodeRoutemonitorBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternode_routemonitor_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternode_routemonitor_binding, got error: %s", err))
		return
	}

	clusternode_routemonitor_bindingSetAttrFromGet(ctx, data, getResponseData)

}
