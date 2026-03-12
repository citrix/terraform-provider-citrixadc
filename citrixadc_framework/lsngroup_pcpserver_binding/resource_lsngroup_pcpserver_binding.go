package lsngroup_pcpserver_binding

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
var _ resource.Resource = &LsngroupPcpserverBindingResource{}
var _ resource.ResourceWithConfigure = (*LsngroupPcpserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsngroupPcpserverBindingResource)(nil)

func NewLsngroupPcpserverBindingResource() resource.Resource {
	return &LsngroupPcpserverBindingResource{}
}

// LsngroupPcpserverBindingResource defines the resource implementation.
type LsngroupPcpserverBindingResource struct {
	client *service.NitroClient
}

func (r *LsngroupPcpserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsngroupPcpserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsngroup_pcpserver_binding"
}

func (r *LsngroupPcpserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsngroupPcpserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsngroupPcpserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsngroup_pcpserver_binding resource")

	// lsngroup_pcpserver_binding := lsngroup_pcpserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_pcpserver_binding.Type(), &lsngroup_pcpserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsngroup_pcpserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lsngroup_pcpserver_binding-config")

	tflog.Trace(ctx, "Created lsngroup_pcpserver_binding resource")

	// Read the updated state back
	r.readLsngroupPcpserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupPcpserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsngroupPcpserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsngroup_pcpserver_binding resource")

	r.readLsngroupPcpserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupPcpserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LsngroupPcpserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lsngroup_pcpserver_binding resource")

	// Create API request body from the model
	// lsngroup_pcpserver_binding := lsngroup_pcpserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lsngroup_pcpserver_binding.Type(), &lsngroup_pcpserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsngroup_pcpserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lsngroup_pcpserver_binding resource")

	// Read the updated state back
	r.readLsngroupPcpserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsngroupPcpserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsngroupPcpserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsngroup_pcpserver_binding resource")

	// For lsngroup_pcpserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lsngroup_pcpserver_binding resource from state")
}

// Helper function to read lsngroup_pcpserver_binding data from API
func (r *LsngroupPcpserverBindingResource) readLsngroupPcpserverBindingFromApi(ctx context.Context, data *LsngroupPcpserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lsngroup_pcpserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsngroup_pcpserver_binding, got error: %s", err))
		return
	}

	lsngroup_pcpserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
