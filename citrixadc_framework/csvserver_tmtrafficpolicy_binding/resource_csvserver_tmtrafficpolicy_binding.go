package csvserver_tmtrafficpolicy_binding

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
var _ resource.Resource = &CsvserverTmtrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverTmtrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverTmtrafficpolicyBindingResource)(nil)

func NewCsvserverTmtrafficpolicyBindingResource() resource.Resource {
	return &CsvserverTmtrafficpolicyBindingResource{}
}

// CsvserverTmtrafficpolicyBindingResource defines the resource implementation.
type CsvserverTmtrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverTmtrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_tmtrafficpolicy_binding"
}

func (r *CsvserverTmtrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_tmtrafficpolicy_binding resource")

	// csvserver_tmtrafficpolicy_binding := csvserver_tmtrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_tmtrafficpolicy_binding.Type(), &csvserver_tmtrafficpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_tmtrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_tmtrafficpolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_tmtrafficpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_tmtrafficpolicy_binding resource")

	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_tmtrafficpolicy_binding resource")

	// Create API request body from the model
	// csvserver_tmtrafficpolicy_binding := csvserver_tmtrafficpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_tmtrafficpolicy_binding.Type(), &csvserver_tmtrafficpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_tmtrafficpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_tmtrafficpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverTmtrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverTmtrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverTmtrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_tmtrafficpolicy_binding resource")

	// For csvserver_tmtrafficpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_tmtrafficpolicy_binding resource from state")
}

// Helper function to read csvserver_tmtrafficpolicy_binding data from API
func (r *CsvserverTmtrafficpolicyBindingResource) readCsvserverTmtrafficpolicyBindingFromApi(ctx context.Context, data *CsvserverTmtrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_tmtrafficpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_tmtrafficpolicy_binding, got error: %s", err))
		return
	}

	csvserver_tmtrafficpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
