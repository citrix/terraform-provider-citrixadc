package appflowaction_analyticsprofile_binding

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
var _ resource.Resource = &AppflowactionAnalyticsprofileBindingResource{}
var _ resource.ResourceWithConfigure = (*AppflowactionAnalyticsprofileBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AppflowactionAnalyticsprofileBindingResource)(nil)

func NewAppflowactionAnalyticsprofileBindingResource() resource.Resource {
	return &AppflowactionAnalyticsprofileBindingResource{}
}

// AppflowactionAnalyticsprofileBindingResource defines the resource implementation.
type AppflowactionAnalyticsprofileBindingResource struct {
	client *service.NitroClient
}

func (r *AppflowactionAnalyticsprofileBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AppflowactionAnalyticsprofileBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appflowaction_analyticsprofile_binding"
}

func (r *AppflowactionAnalyticsprofileBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AppflowactionAnalyticsprofileBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AppflowactionAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating appflowaction_analyticsprofile_binding resource")

	// appflowaction_analyticsprofile_binding := appflowaction_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowaction_analyticsprofile_binding.Type(), &appflowaction_analyticsprofile_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create appflowaction_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("appflowaction_analyticsprofile_binding-config")

	tflog.Trace(ctx, "Created appflowaction_analyticsprofile_binding resource")

	// Read the updated state back
	r.readAppflowactionAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowactionAnalyticsprofileBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AppflowactionAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading appflowaction_analyticsprofile_binding resource")

	r.readAppflowactionAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowactionAnalyticsprofileBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data AppflowactionAnalyticsprofileBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating appflowaction_analyticsprofile_binding resource")

	// Create API request body from the model
	// appflowaction_analyticsprofile_binding := appflowaction_analyticsprofile_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Appflowaction_analyticsprofile_binding.Type(), &appflowaction_analyticsprofile_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update appflowaction_analyticsprofile_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated appflowaction_analyticsprofile_binding resource")

	// Read the updated state back
	r.readAppflowactionAnalyticsprofileBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AppflowactionAnalyticsprofileBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AppflowactionAnalyticsprofileBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting appflowaction_analyticsprofile_binding resource")

	// For appflowaction_analyticsprofile_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted appflowaction_analyticsprofile_binding resource from state")
}

// Helper function to read appflowaction_analyticsprofile_binding data from API
func (r *AppflowactionAnalyticsprofileBindingResource) readAppflowactionAnalyticsprofileBindingFromApi(ctx context.Context, data *AppflowactionAnalyticsprofileBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Appflowaction_analyticsprofile_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read appflowaction_analyticsprofile_binding, got error: %s", err))
		return
	}

	appflowaction_analyticsprofile_bindingSetAttrFromGet(ctx, data, getResponseData)

}
