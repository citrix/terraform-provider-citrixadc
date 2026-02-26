package netprofile_srcportset_binding

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
var _ resource.Resource = &NetprofileSrcportsetBindingResource{}
var _ resource.ResourceWithConfigure = (*NetprofileSrcportsetBindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetprofileSrcportsetBindingResource)(nil)

func NewNetprofileSrcportsetBindingResource() resource.Resource {
	return &NetprofileSrcportsetBindingResource{}
}

// NetprofileSrcportsetBindingResource defines the resource implementation.
type NetprofileSrcportsetBindingResource struct {
	client *service.NitroClient
}

func (r *NetprofileSrcportsetBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetprofileSrcportsetBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netprofile_srcportset_binding"
}

func (r *NetprofileSrcportsetBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetprofileSrcportsetBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetprofileSrcportsetBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netprofile_srcportset_binding resource")

	// netprofile_srcportset_binding := netprofile_srcportset_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netprofile_srcportset_binding.Type(), &netprofile_srcportset_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netprofile_srcportset_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("netprofile_srcportset_binding-config")

	tflog.Trace(ctx, "Created netprofile_srcportset_binding resource")

	// Read the updated state back
	r.readNetprofileSrcportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileSrcportsetBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetprofileSrcportsetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netprofile_srcportset_binding resource")

	r.readNetprofileSrcportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileSrcportsetBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetprofileSrcportsetBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating netprofile_srcportset_binding resource")

	// Create API request body from the model
	// netprofile_srcportset_binding := netprofile_srcportset_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Netprofile_srcportset_binding.Type(), &netprofile_srcportset_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netprofile_srcportset_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated netprofile_srcportset_binding resource")

	// Read the updated state back
	r.readNetprofileSrcportsetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetprofileSrcportsetBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetprofileSrcportsetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netprofile_srcportset_binding resource")

	// For netprofile_srcportset_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted netprofile_srcportset_binding resource from state")
}

// Helper function to read netprofile_srcportset_binding data from API
func (r *NetprofileSrcportsetBindingResource) readNetprofileSrcportsetBindingFromApi(ctx context.Context, data *NetprofileSrcportsetBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Netprofile_srcportset_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netprofile_srcportset_binding, got error: %s", err))
		return
	}

	netprofile_srcportset_bindingSetAttrFromGet(ctx, data, getResponseData)

}
