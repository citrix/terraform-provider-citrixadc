package lbmonitor_metric_binding

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
var _ resource.Resource = &LbmonitorMetricBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmonitorMetricBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmonitorMetricBindingResource)(nil)

func NewLbmonitorMetricBindingResource() resource.Resource {
	return &LbmonitorMetricBindingResource{}
}

// LbmonitorMetricBindingResource defines the resource implementation.
type LbmonitorMetricBindingResource struct {
	client *service.NitroClient
}

func (r *LbmonitorMetricBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmonitorMetricBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmonitor_metric_binding"
}

func (r *LbmonitorMetricBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmonitorMetricBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmonitor_metric_binding resource")

	// lbmonitor_metric_binding := lbmonitor_metric_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmonitor_metric_binding.Type(), &lbmonitor_metric_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmonitor_metric_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbmonitor_metric_binding-config")

	tflog.Trace(ctx, "Created lbmonitor_metric_binding resource")

	// Read the updated state back
	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmonitor_metric_binding resource")

	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbmonitor_metric_binding resource")

	// Create API request body from the model
	// lbmonitor_metric_binding := lbmonitor_metric_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmonitor_metric_binding.Type(), &lbmonitor_metric_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmonitor_metric_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbmonitor_metric_binding resource")

	// Read the updated state back
	r.readLbmonitorMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmonitorMetricBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmonitorMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmonitor_metric_binding resource")

	// For lbmonitor_metric_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbmonitor_metric_binding resource from state")
}

// Helper function to read lbmonitor_metric_binding data from API
func (r *LbmonitorMetricBindingResource) readLbmonitorMetricBindingFromApi(ctx context.Context, data *LbmonitorMetricBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbmonitor_metric_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmonitor_metric_binding, got error: %s", err))
		return
	}

	lbmonitor_metric_bindingSetAttrFromGet(ctx, data, getResponseData)

}
