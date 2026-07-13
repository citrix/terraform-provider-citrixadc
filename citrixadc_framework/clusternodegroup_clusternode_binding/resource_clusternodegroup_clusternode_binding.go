package clusternodegroup_clusternode_binding

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
var _ resource.Resource = &ClusternodegroupClusternodeBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupClusternodeBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupClusternodeBindingResource)(nil)

func NewClusternodegroupClusternodeBindingResource() resource.Resource {
	return &ClusternodegroupClusternodeBindingResource{}
}

// ClusternodegroupClusternodeBindingResource defines the resource implementation.
type ClusternodegroupClusternodeBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupClusternodeBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupClusternodeBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_clusternode_binding"
}

func (r *ClusternodegroupClusternodeBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupClusternodeBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_clusternode_binding resource")

	// clusternodegroup_clusternode_binding := clusternodegroup_clusternode_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_clusternode_binding.Type(), &clusternodegroup_clusternode_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_clusternode_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_clusternode_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_clusternode_binding resource")

	// Read the updated state back
	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_clusternode_binding resource")

	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_clusternode_binding resource")

	// Create API request body from the model
	// clusternodegroup_clusternode_binding := clusternodegroup_clusternode_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_clusternode_binding.Type(), &clusternodegroup_clusternode_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_clusternode_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_clusternode_binding resource")

	// Read the updated state back
	r.readClusternodegroupClusternodeBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupClusternodeBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupClusternodeBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_clusternode_binding resource")

	// For clusternodegroup_clusternode_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_clusternode_binding resource from state")
}

// Helper function to read clusternodegroup_clusternode_binding data from API
func (r *ClusternodegroupClusternodeBindingResource) readClusternodegroupClusternodeBindingFromApi(ctx context.Context, data *ClusternodegroupClusternodeBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_clusternode_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_clusternode_binding, got error: %s", err))
		return
	}

	clusternodegroup_clusternode_bindingSetAttrFromGet(ctx, data, getResponseData)

}
