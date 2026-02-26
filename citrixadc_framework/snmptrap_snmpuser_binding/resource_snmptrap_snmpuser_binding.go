package snmptrap_snmpuser_binding

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
var _ resource.Resource = &SnmptrapSnmpuserBindingResource{}
var _ resource.ResourceWithConfigure = (*SnmptrapSnmpuserBindingResource)(nil)
var _ resource.ResourceWithImportState = (*SnmptrapSnmpuserBindingResource)(nil)

func NewSnmptrapSnmpuserBindingResource() resource.Resource {
	return &SnmptrapSnmpuserBindingResource{}
}

// SnmptrapSnmpuserBindingResource defines the resource implementation.
type SnmptrapSnmpuserBindingResource struct {
	client *service.NitroClient
}

func (r *SnmptrapSnmpuserBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SnmptrapSnmpuserBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_snmptrap_snmpuser_binding"
}

func (r *SnmptrapSnmpuserBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SnmptrapSnmpuserBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating snmptrap_snmpuser_binding resource")

	// snmptrap_snmpuser_binding := snmptrap_snmpuser_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmptrap_snmpuser_binding.Type(), &snmptrap_snmpuser_binding)
	// if err != nil {
	//	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create snmptrap_snmpuser_binding, got error: %s", err))
	//	 return
	// }

	// Generate unique ID for this configuration resource
	data.Id = types.StringValue("snmptrap_snmpuser_binding-config")

	tflog.Trace(ctx, "Created snmptrap_snmpuser_binding resource")

	// Read the updated state back
	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading snmptrap_snmpuser_binding resource")

	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating snmptrap_snmpuser_binding resource")

	// Create API request body from the model
	// snmptrap_snmpuser_binding := snmptrap_snmpuser_bindingGetThePayloadFromtheConfig(ctx, &data)

	// Make API call
	// err := r.client.UpdateUnnamedResource(service.Snmptrap_snmpuser_binding.Type(), &snmptrap_snmpuser_binding)
	// if err != nil {
	// 	 resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update snmptrap_snmpuser_binding, got error: %s", err))
	//	 return
	// }

	tflog.Trace(ctx, "Updated snmptrap_snmpuser_binding resource")

	// Read the updated state back
	r.readSnmptrapSnmpuserBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SnmptrapSnmpuserBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SnmptrapSnmpuserBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting snmptrap_snmpuser_binding resource")

	// For snmptrap_snmpuser_binding, we don't actually delete the resource as it's a global configuration
	// We just remove it from state
	tflog.Trace(ctx, "Deleted snmptrap_snmpuser_binding resource from state")
}

// Helper function to read snmptrap_snmpuser_binding data from API
func (r *SnmptrapSnmpuserBindingResource) readSnmptrapSnmpuserBindingFromApi(ctx context.Context, data *SnmptrapSnmpuserBindingResourceModel, diags *diag.Diagnostics) {
	getResponseData, err := r.client.FindResource(service.Snmptrap_snmpuser_binding.Type(), "")
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read snmptrap_snmpuser_binding, got error: %s", err))
		return
	}

	snmptrap_snmpuser_bindingSetAttrFromGet(ctx, data, getResponseData)

}
