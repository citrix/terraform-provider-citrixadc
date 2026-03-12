package dnsglobal_dnspolicy_binding

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
var _ resource.Resource = &DnsglobalDnspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*DnsglobalDnspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*DnsglobalDnspolicyBindingResource)(nil)

func NewDnsglobalDnspolicyBindingResource() resource.Resource {
	return &DnsglobalDnspolicyBindingResource{}
}

// DnsglobalDnspolicyBindingResource defines the resource implementation.
type DnsglobalDnspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *DnsglobalDnspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *DnsglobalDnspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dnsglobal_dnspolicy_binding"
}

func (r *DnsglobalDnspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *DnsglobalDnspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data DnsglobalDnspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating dnsglobal_dnspolicy_binding resource")

	// dnsglobal_dnspolicy_binding := dnsglobal_dnspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsglobal_dnspolicy_binding.Type(), &dnsglobal_dnspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create dnsglobal_dnspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("dnsglobal_dnspolicy_binding-config")

	tflog.Trace(ctx, "Created dnsglobal_dnspolicy_binding resource")

	// Read the updated state back
	r.readDnsglobalDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsglobalDnspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DnsglobalDnspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading dnsglobal_dnspolicy_binding resource")

	r.readDnsglobalDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsglobalDnspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DnsglobalDnspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating dnsglobal_dnspolicy_binding resource")

	// Create API request body from the model
	// dnsglobal_dnspolicy_binding := dnsglobal_dnspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Dnsglobal_dnspolicy_binding.Type(), &dnsglobal_dnspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update dnsglobal_dnspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated dnsglobal_dnspolicy_binding resource")

	// Read the updated state back
	r.readDnsglobalDnspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *DnsglobalDnspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data DnsglobalDnspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting dnsglobal_dnspolicy_binding resource")

	// For dnsglobal_dnspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted dnsglobal_dnspolicy_binding resource from state")
}

// Helper function to read dnsglobal_dnspolicy_binding data from API
func (r *DnsglobalDnspolicyBindingResource) readDnsglobalDnspolicyBindingFromApi(ctx context.Context, data *DnsglobalDnspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Dnsglobal_dnspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read dnsglobal_dnspolicy_binding, got error: %s", err))
		return
	}

	dnsglobal_dnspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
