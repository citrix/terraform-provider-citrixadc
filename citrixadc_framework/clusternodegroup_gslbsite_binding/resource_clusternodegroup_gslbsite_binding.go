package clusternodegroup_gslbsite_binding

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
var _ resource.Resource = &ClusternodegroupGslbsiteBindingResource{}
var _ resource.ResourceWithConfigure = (*ClusternodegroupGslbsiteBindingResource)(nil)
var _ resource.ResourceWithImportState = (*ClusternodegroupGslbsiteBindingResource)(nil)

func NewClusternodegroupGslbsiteBindingResource() resource.Resource {
	return &ClusternodegroupGslbsiteBindingResource{}
}

// ClusternodegroupGslbsiteBindingResource defines the resource implementation.
type ClusternodegroupGslbsiteBindingResource struct {
	client *service.NitroClient
}

func (r *ClusternodegroupGslbsiteBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusternodegroupGslbsiteBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusternodegroup_gslbsite_binding"
}

func (r *ClusternodegroupGslbsiteBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ClusternodegroupGslbsiteBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating clusternodegroup_gslbsite_binding resource")

	// clusternodegroup_gslbsite_binding := clusternodegroup_gslbsite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_gslbsite_binding.Type(), &clusternodegroup_gslbsite_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create clusternodegroup_gslbsite_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("clusternodegroup_gslbsite_binding-config")

	tflog.Trace(ctx, "Created clusternodegroup_gslbsite_binding resource")

	// Read the updated state back
	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading clusternodegroup_gslbsite_binding resource")

	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating clusternodegroup_gslbsite_binding resource")

	// Create API request body from the model
	// clusternodegroup_gslbsite_binding := clusternodegroup_gslbsite_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Clusternodegroup_gslbsite_binding.Type(), &clusternodegroup_gslbsite_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update clusternodegroup_gslbsite_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated clusternodegroup_gslbsite_binding resource")

	// Read the updated state back
	r.readClusternodegroupGslbsiteBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusternodegroupGslbsiteBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusternodegroupGslbsiteBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting clusternodegroup_gslbsite_binding resource")

	// For clusternodegroup_gslbsite_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted clusternodegroup_gslbsite_binding resource from state")
}

// Helper function to read clusternodegroup_gslbsite_binding data from API
func (r *ClusternodegroupGslbsiteBindingResource) readClusternodegroupGslbsiteBindingFromApi(ctx context.Context, data *ClusternodegroupGslbsiteBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Clusternodegroup_gslbsite_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read clusternodegroup_gslbsite_binding, got error: %s", err))
		return
	}

	clusternodegroup_gslbsite_bindingSetAttrFromGet(ctx, data, getResponseData)

}
