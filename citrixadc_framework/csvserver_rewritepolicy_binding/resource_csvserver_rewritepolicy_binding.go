package csvserver_rewritepolicy_binding

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
var _ resource.Resource = &CsvserverRewritepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverRewritepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverRewritepolicyBindingResource)(nil)

func NewCsvserverRewritepolicyBindingResource() resource.Resource {
	return &CsvserverRewritepolicyBindingResource{}
}

// CsvserverRewritepolicyBindingResource defines the resource implementation.
type CsvserverRewritepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverRewritepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverRewritepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_rewritepolicy_binding"
}

func (r *CsvserverRewritepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverRewritepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_rewritepolicy_binding resource")

	// csvserver_rewritepolicy_binding := csvserver_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_rewritepolicy_binding.Type(), &csvserver_rewritepolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_rewritepolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_rewritepolicy_binding resource")

	// Read the updated state back
	r.readCsvserverRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverRewritepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_rewritepolicy_binding resource")

	r.readCsvserverRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverRewritepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverRewritepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_rewritepolicy_binding resource")

	// Create API request body from the model
	// csvserver_rewritepolicy_binding := csvserver_rewritepolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_rewritepolicy_binding.Type(), &csvserver_rewritepolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_rewritepolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_rewritepolicy_binding resource")

	// Read the updated state back
	r.readCsvserverRewritepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverRewritepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverRewritepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_rewritepolicy_binding resource")

	// For csvserver_rewritepolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_rewritepolicy_binding resource from state")
}

// Helper function to read csvserver_rewritepolicy_binding data from API
func (r *CsvserverRewritepolicyBindingResource) readCsvserverRewritepolicyBindingFromApi(ctx context.Context, data *CsvserverRewritepolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_rewritepolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_rewritepolicy_binding, got error: %s", err))
		return
	}

	csvserver_rewritepolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
