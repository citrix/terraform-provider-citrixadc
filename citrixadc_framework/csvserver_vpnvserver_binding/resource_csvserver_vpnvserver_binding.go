package csvserver_vpnvserver_binding

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
var _ resource.Resource = &CsvserverVpnvserverBindingResource{}
var _ resource.ResourceWithConfigure = (*CsvserverVpnvserverBindingResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverVpnvserverBindingResource)(nil)

func NewCsvserverVpnvserverBindingResource() resource.Resource {
	return &CsvserverVpnvserverBindingResource{}
}

// CsvserverVpnvserverBindingResource defines the resource implementation.
type CsvserverVpnvserverBindingResource struct {
	client *service.NitroClient
}

func (r *CsvserverVpnvserverBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverVpnvserverBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver_vpnvserver_binding"
}

func (r *CsvserverVpnvserverBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverVpnvserverBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverVpnvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver_vpnvserver_binding resource")

	// csvserver_vpnvserver_binding := csvserver_vpnvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_vpnvserver_binding.Type(), &csvserver_vpnvserver_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver_vpnvserver_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("csvserver_vpnvserver_binding-config")

	tflog.Trace(ctx, "Created csvserver_vpnvserver_binding resource")

	// Read the updated state back
	r.readCsvserverVpnvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverVpnvserverBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverVpnvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver_vpnvserver_binding resource")

	r.readCsvserverVpnvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverVpnvserverBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data CsvserverVpnvserverBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating csvserver_vpnvserver_binding resource")

	// Create API request body from the model
	// csvserver_vpnvserver_binding := csvserver_vpnvserver_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Csvserver_vpnvserver_binding.Type(), &csvserver_vpnvserver_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver_vpnvserver_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated csvserver_vpnvserver_binding resource")

	// Read the updated state back
	r.readCsvserverVpnvserverBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverVpnvserverBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverVpnvserverBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver_vpnvserver_binding resource")

	// For csvserver_vpnvserver_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted csvserver_vpnvserver_binding resource from state")
}

// Helper function to read csvserver_vpnvserver_binding data from API
func (r *CsvserverVpnvserverBindingResource) readCsvserverVpnvserverBindingFromApi(ctx context.Context, data *CsvserverVpnvserverBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Csvserver_vpnvserver_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver_vpnvserver_binding, got error: %s", err))
		return
	}

	csvserver_vpnvserver_bindingSetAttrFromGet(ctx, data, getResponseData)

}
