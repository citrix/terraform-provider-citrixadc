package csvserver_cspolicy_binding

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
var _ resource.Resource = &CsvserverCspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverCspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverCspolicyBindingResource)(nil)

func NewCsvserverCspolicyBindingResource() resource.Resource {
	return &CsvserverCspolicyBindingResource{}
}

// CsvserverCspolicyBindingResource defines the resource implementation.
type CsvserverCspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverCspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverCspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_cspolicy_binding"
}

func (r *CsvserverCspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverCspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverCspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_cspolicy_binding resource")

	// csvserver_cspolicy_binding := csvserver_cspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_cspolicy_binding.Type(), &csvserver_cspolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_cspolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_cspolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_cspolicy_binding resource")

	// Read the updated state back
	r.readCsvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_cspolicy_binding resource")

	r.readCsvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverCspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_cspolicy_binding resource")

	// Create API request body from the model
	// csvserver_cspolicy_binding := csvserver_cspolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_cspolicy_binding.Type(), &csvserver_cspolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_cspolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_cspolicy_binding resource")

	// Read the updated state back
	r.readCsvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverCspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_cspolicy_binding resource")

	// For csvserver_cspolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_cspolicy_binding resource from state")
}

// Helper function to read csvserver_cspolicy_binding data from API
func (r *CsvserverCspolicyBindingResource) readCsvserverCspolicyBindingFromApi(ctx context.Context, data *CsvserverCspolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_cspolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_cspolicy_binding, got error: %s", err))
		return
	}

	csvserver_cspolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
