package dnspolicylabel_dnspolicy_binding

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
var _ resource.Resource = &DnspolicylabelDnspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*DnspolicylabelDnspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*DnspolicylabelDnspolicyBindingResource)(nil)

func NewDnspolicylabelDnspolicyBindingResource() resource.Resource {
	return &DnspolicylabelDnspolicyBindingResource{}
}

// DnspolicylabelDnspolicyBindingResource defines the resource implementation.
type DnspolicylabelDnspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *DnspolicylabelDnspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnspolicylabelDnspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnspolicylabel_dnspolicy_binding"
}

func (r *DnspolicylabelDnspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnspolicylabelDnspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnspolicylabelDnspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnspolicylabel_dnspolicy_binding resource")

	// dnspolicylabel_dnspolicy_binding := dnspolicylabel_dnspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnspolicylabel_dnspolicy_binding.Type(), &dnspolicylabel_dnspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnspolicylabel_dnspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnspolicylabel_dnspolicy_binding-config")

	tflog.Trace(ctx, "Created dnspolicylabel_dnspolicy_binding resource")

	// Read the updated state back
	r.readDnspolicylabelDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnspolicylabelDnspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnspolicylabelDnspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnspolicylabel_dnspolicy_binding resource")

	r.readDnspolicylabelDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnspolicylabelDnspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnspolicylabelDnspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnspolicylabel_dnspolicy_binding resource")

	// Create API request body from the model
	// dnspolicylabel_dnspolicy_binding := dnspolicylabel_dnspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnspolicylabel_dnspolicy_binding.Type(), &dnspolicylabel_dnspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnspolicylabel_dnspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnspolicylabel_dnspolicy_binding resource")

	// Read the updated state back
	r.readDnspolicylabelDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnspolicylabelDnspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnspolicylabelDnspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnspolicylabel_dnspolicy_binding resource")

	// For dnspolicylabel_dnspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnspolicylabel_dnspolicy_binding resource from state")
}

// Helper function to read dnspolicylabel_dnspolicy_binding data from API
func (r *DnspolicylabelDnspolicyBindingResource) readDnspolicylabelDnspolicyBindingFromApi(ctx context.Context, data *DnspolicylabelDnspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnspolicylabel_dnspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnspolicylabel_dnspolicy_binding, got error: %s", err))
		return
	}

	dnspolicylabel_dnspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
