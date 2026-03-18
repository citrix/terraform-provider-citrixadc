package cmppolicylabel_cmppolicy_binding

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
var _ resource.Resource = &CmppolicylabelCmppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CmppolicylabelCmppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CmppolicylabelCmppolicyBindingResource)(nil)

func NewCmppolicylabelCmppolicyBindingResource() resource.Resource {
	return &CmppolicylabelCmppolicyBindingResource{}
}

// CmppolicylabelCmppolicyBindingResource defines the resource implementation.
type CmppolicylabelCmppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CmppolicylabelCmppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CmppolicylabelCmppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cmppolicylabel_cmppolicy_binding"
}

func (r *CmppolicylabelCmppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CmppolicylabelCmppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CmppolicylabelCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating cmppolicylabel_cmppolicy_binding resource")

	// cmppolicylabel_cmppolicy_binding := cmppolicylabel_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmppolicylabel_cmppolicy_binding.Type(), &cmppolicylabel_cmppolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create cmppolicylabel_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("cmppolicylabel_cmppolicy_binding-config")

	tflog.Trace(ctx, "Created cmppolicylabel_cmppolicy_binding resource")

	// Read the updated state back
	r.readCmppolicylabelCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmppolicylabelCmppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CmppolicylabelCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cmppolicylabel_cmppolicy_binding resource")

	r.readCmppolicylabelCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmppolicylabelCmppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CmppolicylabelCmppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating cmppolicylabel_cmppolicy_binding resource")

	// Create API request body from the model
	// cmppolicylabel_cmppolicy_binding := cmppolicylabel_cmppolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Cmppolicylabel_cmppolicy_binding.Type(), &cmppolicylabel_cmppolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update cmppolicylabel_cmppolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated cmppolicylabel_cmppolicy_binding resource")

	// Read the updated state back
	r.readCmppolicylabelCmppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CmppolicylabelCmppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CmppolicylabelCmppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cmppolicylabel_cmppolicy_binding resource")

	// For cmppolicylabel_cmppolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted cmppolicylabel_cmppolicy_binding resource from state")
}

// Helper function to read cmppolicylabel_cmppolicy_binding data from API
func (r *CmppolicylabelCmppolicyBindingResource) readCmppolicylabelCmppolicyBindingFromApi(ctx context.Context, data *CmppolicylabelCmppolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Cmppolicylabel_cmppolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read cmppolicylabel_cmppolicy_binding, got error: %s", err))
		return
	}

	cmppolicylabel_cmppolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
