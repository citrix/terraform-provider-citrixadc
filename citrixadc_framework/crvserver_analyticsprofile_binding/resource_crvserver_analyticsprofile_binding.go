package crvserver_analyticsprofile_binding

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
var _ resource.Resource = &CrvserverAnalyticsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*CrvserverAnalyticsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverAnalyticsprofileBindingResource)(nil)

func NewCrvserverAnalyticsprofileBindingResource() resource.Resource {
	return &CrvserverAnalyticsprofileBindingResource{}
}

// CrvserverAnalyticsprofileBindingResource defines the resource implementation.
type CrvserverAnalyticsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *CrvserverAnalyticsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverAnalyticsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver_analyticsprofile_binding"
}

func (r *CrvserverAnalyticsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverAnalyticsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver_analyticsprofile_binding resource")

	// crvserver_analyticsprofile_binding := crvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_analyticsprofile_binding.Type(), &crvserver_analyticsprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("crvserver_analyticsprofile_binding-config")

	tflog.Trace(ctx, "Created crvserver_analyticsprofile_binding resource")

	// Read the updated state back
	r.readCrvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAnalyticsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver_analyticsprofile_binding resource")

	r.readCrvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAnalyticsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CrvserverAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating crvserver_analyticsprofile_binding resource")

	// Create API request body from the model
	// crvserver_analyticsprofile_binding := crvserver_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Crvserver_analyticsprofile_binding.Type(), &crvserver_analyticsprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated crvserver_analyticsprofile_binding resource")

	// Read the updated state back
	r.readCrvserverAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverAnalyticsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver_analyticsprofile_binding resource")

	// For crvserver_analyticsprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted crvserver_analyticsprofile_binding resource from state")
}

// Helper function to read crvserver_analyticsprofile_binding data from API
func (r *CrvserverAnalyticsprofileBindingResource) readCrvserverAnalyticsprofileBindingFromApi(ctx context.Context, data *CrvserverAnalyticsprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Crvserver_analyticsprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver_analyticsprofile_binding, got error: %s", err))
		return
	}

	crvserver_analyticsprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
