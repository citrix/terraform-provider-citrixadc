package csvserver_responderpolicy_binding

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
var _ resource.Resource = &CsvserverResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverResponderpolicyBindingResource)(nil)

func NewCsvserverResponderpolicyBindingResource() resource.Resource {
	return &CsvserverResponderpolicyBindingResource{}
}

// CsvserverResponderpolicyBindingResource defines the resource implementation.
type CsvserverResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_responderpolicy_binding"
}

func (r *CsvserverResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_responderpolicy_binding resource")

	// csvserver_responderpolicy_binding := csvserver_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_responderpolicy_binding.Type(), &csvserver_responderpolicy_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_responderpolicy_binding-config")

	tflog.Trace(ctx, "Created csvserver_responderpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_responderpolicy_binding resource")

	r.readCsvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_responderpolicy_binding resource")

	// Create API request body from the model
	// csvserver_responderpolicy_binding := csvserver_responderpolicy_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_responderpolicy_binding.Type(), &csvserver_responderpolicy_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_responderpolicy_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_responderpolicy_binding resource")

	// Read the updated state back
	r.readCsvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_responderpolicy_binding resource")

	// For csvserver_responderpolicy_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_responderpolicy_binding resource from state")
}

// Helper function to read csvserver_responderpolicy_binding data from API
func (r *CsvserverResponderpolicyBindingResource) readCsvserverResponderpolicyBindingFromApi(ctx context.Context, data *CsvserverResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_responderpolicy_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	csvserver_responderpolicy_bindingSetAttrFromGet(ctx, data, getResponseData)

}
