package lbvserver_analyticsprofile_binding

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
var _ resource.Resource = &LbvserverAnalyticsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*LbvserverAnalyticsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverAnalyticsprofileBindingResource)(nil)

func NewLbvserverAnalyticsprofileBindingResource() resource.Resource {
	return &LbvserverAnalyticsprofileBindingResource{}
}

// LbvserverAnalyticsprofileBindingResource defines the resource implementation.
type LbvserverAnalyticsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *LbvserverAnalyticsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverAnalyticsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver_analyticsprofile_binding"
}

func (r *LbvserverAnalyticsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverAnalyticsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver_analyticsprofile_binding resource")

	// lbvserver_analyticsprofile_binding := lbvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_analyticsprofile_binding.Type(), &lbvserver_analyticsprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("lbvserver_analyticsprofile_binding-config")

	tflog.Trace(ctx, "Created lbvserver_analyticsprofile_binding resource")

	// Read the updated state back
	r.readLbvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAnalyticsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver_analyticsprofile_binding resource")

	r.readLbvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAnalyticsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data LbvserverAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating lbvserver_analyticsprofile_binding resource")

	// Create API request body from the model
	// lbvserver_analyticsprofile_binding := lbvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Lbvserver_analyticsprofile_binding.Type(), &lbvserver_analyticsprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated lbvserver_analyticsprofile_binding resource")

	// Read the updated state back
	r.readLbvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverAnalyticsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver_analyticsprofile_binding resource")

	// For lbvserver_analyticsprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted lbvserver_analyticsprofile_binding resource from state")
}

// Helper function to read lbvserver_analyticsprofile_binding data from API
func (r *LbvserverAnalyticsprofileBindingResource) readLbvserverAnalyticsprofileBindingFromApi(ctx context.Context, data *LbvserverAnalyticsprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Lbvserver_analyticsprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver_analyticsprofile_binding, got error: %s", err))
		return
	}

	lbvserver_analyticsprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
