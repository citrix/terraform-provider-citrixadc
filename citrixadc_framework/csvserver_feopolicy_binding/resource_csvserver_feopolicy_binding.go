package csvserver_feopolicy_binding

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
var _ resource.Resource = &CsvserverFeopolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverFeopolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverFeopolicyBindingResource)(nil)

func NewCsvserverFeopolicyBindingResource() resource.Resource {
	return &CsvserverFeopolicyBindingResource{}
}

// CsvserverFeopolicyBindingResource defines the resource implementation.
type CsvserverFeopolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverFeopolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverFeopolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_feopolicy_binding"
}

func (r *CsvserverFeopolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverFeopolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_feopolicy_binding resource")

	// csvserver_feopolicy_binding := csvserver_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_feopolicy_binding.Type(), &csvserver_feopolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_feopolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_feopolicy_binding resource")

	// Read the updated state back
	r.readCsvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverFeopolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_feopolicy_binding resource")

	r.readCsvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverFeopolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverFeopolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_feopolicy_binding resource")

	// Create API request body from the model
	// csvserver_feopolicy_binding := csvserver_feopolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_feopolicy_binding.Type(), &csvserver_feopolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_feopolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_feopolicy_binding resource")

	// Read the updated state back
	r.readCsvserverFeopolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverFeopolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverFeopolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_feopolicy_binding resource")

	// For csvserver_feopolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_feopolicy_binding resource from state")
}

// Helper function to read csvserver_feopolicy_binding data from API
func (r *CsvserverFeopolicyBindingResource) readCsvserverFeopolicyBindingFromApi(ctx context.Context, data *CsvserverFeopolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_feopolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_feopolicy_binding, got error: %s", err))
		return
	}

	csvserver_feopolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
