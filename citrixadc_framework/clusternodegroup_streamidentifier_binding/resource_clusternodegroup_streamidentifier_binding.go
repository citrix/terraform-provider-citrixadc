package clusternodegroup_streamidentifier_binding

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
var _ resource.Resource = &ClusternodegroupStreamidentifierBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupStreamidentifierBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupStreamidentifierBindingResource)(nil)

func NewClusternodegroupStreamidentifierBindingResource() resource.Resource {
	return &ClusternodegroupStreamidentifierBindingResource{}
}

// ClusternodegroupStreamidentifierBindingResource defines the resource implementation.
type ClusternodegroupStreamidentifierBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupStreamidentifierBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_streamidentifier_binding"
}

func (r *ClusternodegroupStreamidentifierBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_streamidentifier_binding resource")

	// clusternodegroup_streamidentifier_binding := clusternodegroup_streamidentifier_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_streamidentifier_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_streamidentifier_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_streamidentifier_binding resource")

	// Read the updated state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_streamidentifier_binding resource")

	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_streamidentifier_binding resource")

	// Create API request body from the model
	// clusternodegroup_streamidentifier_binding := clusternodegroup_streamidentifier_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_streamidentifier_binding.Type(), &clusternodegroup_streamidentifier_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_streamidentifier_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_streamidentifier_binding resource")

	// Read the updated state back
	r.readClusternodegroupStreamidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupStreamidentifierBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupStreamidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_streamidentifier_binding resource")

	// For clusternodegroup_streamidentifier_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_streamidentifier_binding resource from state")
}

// Helper function to read clusternodegroup_streamidentifier_binding data from API
func (r *ClusternodegroupStreamidentifierBindingResource) readClusternodegroupStreamidentifierBindingFromApi(ctx context.Context, data *ClusternodegroupStreamidentifierBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_streamidentifier_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_streamidentifier_binding, got error: %s", err))
		return
	}

	clusternodegroup_streamidentifier_bindingSetAttrFromGet(ctx, data, getResponseData)

}
