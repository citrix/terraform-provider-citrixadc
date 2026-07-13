package clusternodegroup_crvserver_binding

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
var _ resource.Resource = &ClusternodegroupCrvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupCrvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupCrvserverBindingResource)(nil)

func NewClusternodegroupCrvserverBindingResource() resource.Resource {
	return &ClusternodegroupCrvserverBindingResource{}
}

// ClusternodegroupCrvserverBindingResource defines the resource implementation.
type ClusternodegroupCrvserverBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupCrvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupCrvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_crvserver_binding"
}

func (r *ClusternodegroupCrvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupCrvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupCrvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_crvserver_binding resource")

	// clusternodegroup_crvserver_binding := clusternodegroup_crvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_crvserver_binding.Type(), &clusternodegroup_crvserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_crvserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_crvserver_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_crvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCrvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupCrvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_crvserver_binding resource")

	r.readClusternodegroupCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCrvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupCrvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_crvserver_binding resource")

	// Create API request body from the model
	// clusternodegroup_crvserver_binding := clusternodegroup_crvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_crvserver_binding.Type(), &clusternodegroup_crvserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_crvserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_crvserver_binding resource")

	// Read the updated state back
	r.readClusternodegroupCrvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupCrvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupCrvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_crvserver_binding resource")

	// For clusternodegroup_crvserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_crvserver_binding resource from state")
}

// Helper function to read clusternodegroup_crvserver_binding data from API
func (r *ClusternodegroupCrvserverBindingResource) readClusternodegroupCrvserverBindingFromApi(ctx context.Context, data *ClusternodegroupCrvserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_crvserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_crvserver_binding, got error: %s", err))
		return
	}

	clusternodegroup_crvserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
