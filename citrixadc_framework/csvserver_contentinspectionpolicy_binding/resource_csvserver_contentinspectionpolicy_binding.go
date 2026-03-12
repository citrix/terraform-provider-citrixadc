package csvserver_contentinspectionpolicy_binding

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
var _ resource.Resource = &CsvserverContentinspectionpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverContentinspectionpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverContentinspectionpolicyBindingResource)(nil)

func NewCsvserverContentinspectionpolicyBindingResource() resource.Resource {
	return &CsvserverContentinspectionpolicyBindingResource{}
}

// CsvserverContentinspectionpolicyBindingResource defines the resource implementation.
type CsvserverContentinspectionpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverContentinspectionpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverContentinspectionpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_contentinspectionpolicy_binding"
}

func (r *CsvserverContentinspectionpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverContentinspectionpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_contentinspectionpolicy_binding resource")

	// csvserver_contentinspectionpolicy_binding := csvserver_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_contentinspectionpolicy_binding.Type(), &csvserver_contentinspectionpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_contentinspectionpolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverContentinspectionpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_contentinspectionpolicy_binding resource")

	r.readCsvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverContentinspectionpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_contentinspectionpolicy_binding resource")

	// Create API request body from the model
	// csvserver_contentinspectionpolicy_binding := csvserver_contentinspectionpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_contentinspectionpolicy_binding.Type(), &csvserver_contentinspectionpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_contentinspectionpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_contentinspectionpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverContentinspectionpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverContentinspectionpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverContentinspectionpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_contentinspectionpolicy_binding resource")

	// For csvserver_contentinspectionpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_contentinspectionpolicy_binding resource from state")
}

// Helper function to read csvserver_contentinspectionpolicy_binding data from API
func (r *CsvserverContentinspectionpolicyBindingResource) readCsvserverContentinspectionpolicyBindingFromApi(ctx context.Context, data *CsvserverContentinspectionpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_contentinspectionpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_contentinspectionpolicy_binding, got error: %s", err))
		return
	}

	csvserver_contentinspectionpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
