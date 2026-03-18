package icaglobal_icapolicy_binding

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
var _ resource.Resource = &IcaglobalIcapolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*IcaglobalIcapolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*IcaglobalIcapolicyBindingResource)(nil)

func NewIcaglobalIcapolicyBindingResource() resource.Resource {
	return &IcaglobalIcapolicyBindingResource{}
}

// IcaglobalIcapolicyBindingResource defines the resource implementation.
type IcaglobalIcapolicyBindingResource struct {
	client *service.NitroClient
}

func (r *IcaglobalIcapolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *IcaglobalIcapolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_icaglobal_icapolicy_binding"
}

func (r *IcaglobalIcapolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *IcaglobalIcapolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data IcaglobalIcapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating icaglobal_icapolicy_binding resource")

	// icaglobal_icapolicy_binding := icaglobal_icapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Icaglobal_icapolicy_binding.Type(), &icaglobal_icapolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create icaglobal_icapolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("icaglobal_icapolicy_binding-config")

	tflog.Trace(ctx, "Created icaglobal_icapolicy_binding resource")

	// Read the updated state back
	r.readIcaglobalIcapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaglobalIcapolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data IcaglobalIcapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading icaglobal_icapolicy_binding resource")

	r.readIcaglobalIcapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaglobalIcapolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data IcaglobalIcapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating icaglobal_icapolicy_binding resource")

	// Create API request body from the model
	// icaglobal_icapolicy_binding := icaglobal_icapolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Icaglobal_icapolicy_binding.Type(), &icaglobal_icapolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update icaglobal_icapolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated icaglobal_icapolicy_binding resource")

	// Read the updated state back
	r.readIcaglobalIcapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *IcaglobalIcapolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data IcaglobalIcapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting icaglobal_icapolicy_binding resource")

	// For icaglobal_icapolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted icaglobal_icapolicy_binding resource from state")
}

// Helper function to read icaglobal_icapolicy_binding data from API
func (r *IcaglobalIcapolicyBindingResource) readIcaglobalIcapolicyBindingFromApi(ctx context.Context, data *IcaglobalIcapolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Icaglobal_icapolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read icaglobal_icapolicy_binding, got error: %s", err))
		return
	}

	icaglobal_icapolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
