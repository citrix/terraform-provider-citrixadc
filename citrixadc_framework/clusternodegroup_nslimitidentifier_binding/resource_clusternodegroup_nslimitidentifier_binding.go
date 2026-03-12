package clusternodegroup_nslimitidentifier_binding

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
var _ resource.Resource = &ClusternodegroupNslimitidentifierBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupNslimitidentifierBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupNslimitidentifierBindingResource)(nil)

func NewClusternodegroupNslimitidentifierBindingResource() resource.Resource {
	return &ClusternodegroupNslimitidentifierBindingResource{}
}

// ClusternodegroupNslimitidentifierBindingResource defines the resource implementation.
type ClusternodegroupNslimitidentifierBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupNslimitidentifierBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_nslimitidentifier_binding"
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupNslimitidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_nslimitidentifier_binding resource")

	// clusternodegroup_nslimitidentifier_binding := clusternodegroup_nslimitidentifier_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_nslimitidentifier_binding.Type(), &clusternodegroup_nslimitidentifier_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_nslimitidentifier_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_nslimitidentifier_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_nslimitidentifier_binding resource")

	// Read the updated state back
	r.readClusternodegroupNslimitidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupNslimitidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_nslimitidentifier_binding resource")

	r.readClusternodegroupNslimitidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupNslimitidentifierBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_nslimitidentifier_binding resource")

	// Create API request body from the model
	// clusternodegroup_nslimitidentifier_binding := clusternodegroup_nslimitidentifier_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_nslimitidentifier_binding.Type(), &clusternodegroup_nslimitidentifier_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_nslimitidentifier_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_nslimitidentifier_binding resource")

	// Read the updated state back
	r.readClusternodegroupNslimitidentifierBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupNslimitidentifierBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupNslimitidentifierBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_nslimitidentifier_binding resource")

	// For clusternodegroup_nslimitidentifier_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_nslimitidentifier_binding resource from state")
}

// Helper function to read clusternodegroup_nslimitidentifier_binding data from API
func (r *ClusternodegroupNslimitidentifierBindingResource) readClusternodegroupNslimitidentifierBindingFromApi(ctx context.Context, data *ClusternodegroupNslimitidentifierBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_nslimitidentifier_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_nslimitidentifier_binding, got error: %s", err))
		return
	}

	clusternodegroup_nslimitidentifier_bindingSetAttrFromGet(ctx, data, getResponseData)

}
