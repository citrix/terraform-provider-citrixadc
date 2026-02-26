package gslbservicegroup_gslbservicegroupmember_binding

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
var _ resource.Resource = &GslbservicegroupGslbservicegroupmemberBindingResource{}
var _ resource.ResourceWithConfigure = (*GslbservicegroupGslbservicegroupmemberBindingResource)(nil)
var _ resource.ResourceWithImportState = (*GslbservicegroupGslbservicegroupmemberBindingResource)(nil)

func NewGslbservicegroupGslbservicegroupmemberBindingResource() resource.Resource {
	return &GslbservicegroupGslbservicegroupmemberBindingResource{}
}

// GslbservicegroupGslbservicegroupmemberBindingResource defines the resource implementation.
type GslbservicegroupGslbservicegroupmemberBindingResource struct {
	client *service.NitroClient
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservicegroup_gslbservicegroupmember_binding"
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservicegroup_gslbservicegroupmember_binding resource")

	// gslbservicegroup_gslbservicegroupmember_binding := gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), &gslbservicegroup_gslbservicegroupmember_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("gslbservicegroup_gslbservicegroupmember_binding-config")

	tflog.Trace(ctx, "Created gslbservicegroup_gslbservicegroupmember_binding resource")

	// Read the updated state back
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservicegroup_gslbservicegroupmember_binding resource")

	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservicegroup_gslbservicegroupmember_binding resource")

	// Create API request body from the model
	// gslbservicegroup_gslbservicegroupmember_binding := gslbservicegroup_gslbservicegroupmember_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), &gslbservicegroup_gslbservicegroupmember_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated gslbservicegroup_gslbservicegroupmember_binding resource")

	// Read the updated state back
	r.readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbservicegroupGslbservicegroupmemberBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbservicegroupGslbservicegroupmemberBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservicegroup_gslbservicegroupmember_binding resource")

	// For gslbservicegroup_gslbservicegroupmember_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted gslbservicegroup_gslbservicegroupmember_binding resource from state")
}

// Helper function to read gslbservicegroup_gslbservicegroupmember_binding data from API
func (r *GslbservicegroupGslbservicegroupmemberBindingResource) readGslbservicegroupGslbservicegroupmemberBindingFromApi(ctx context.Context, data *GslbservicegroupGslbservicegroupmemberBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Gslbservicegroup_gslbservicegroupmember_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservicegroup_gslbservicegroupmember_binding, got error: %s", err))
		return
	}

	gslbservicegroup_gslbservicegroupmember_bindingSetAttrFromGet(ctx, data, getResponseData)

}
