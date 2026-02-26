package lbmetrictable_metric_binding

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
var _ resource.Resource = &LbmetrictableMetricBindingResource{}
var _ resource.ResourceWithConfigure = (*LbmetrictableMetricBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbmetrictableMetricBindingResource)(nil)

func NewLbmetrictableMetricBindingResource() resource.Resource {
	return &LbmetrictableMetricBindingResource{}
}

// LbmetrictableMetricBindingResource defines the resource implementation.
type LbmetrictableMetricBindingResource struct {
	client *service.NitroClient
}

func (r *LbmetrictableMetricBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbmetrictableMetricBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbmetrictable_metric_binding"
}

func (r *LbmetrictableMetricBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbmetrictableMetricBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbmetrictable_metric_binding resource")

	// lbmetrictable_metric_binding := lbmetrictable_metric_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmetrictable_metric_binding.Type(), &lbmetrictable_metric_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbmetrictable_metric_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbmetrictable_metric_binding-config")

	tflog.Trace(ctx, "Created lbmetrictable_metric_binding resource")

	// Read the updated state back
	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmetrictableMetricBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbmetrictable_metric_binding resource")

	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmetrictableMetricBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbmetrictable_metric_binding resource")

	// Create API request body from the model
	// lbmetrictable_metric_binding := lbmetrictable_metric_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbmetrictable_metric_binding.Type(), &lbmetrictable_metric_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbmetrictable_metric_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbmetrictable_metric_binding resource")

	// Read the updated state back
	r.readLbmetrictableMetricBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbmetrictableMetricBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbmetrictableMetricBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbmetrictable_metric_binding resource")

	// For lbmetrictable_metric_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbmetrictable_metric_binding resource from state")
}

// Helper function to read lbmetrictable_metric_binding data from API
func (r *LbmetrictableMetricBindingResource) readLbmetrictableMetricBindingFromApi(ctx context.Context, data *LbmetrictableMetricBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbmetrictable_metric_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbmetrictable_metric_binding, got error: %s", err))
		return
	}

	lbmetrictable_metric_bindingSetAttrFromGet(ctx, data, getResponseData)

}
