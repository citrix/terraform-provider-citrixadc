package tunnelglobal_tunneltrafficpolicy_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &TunnelglobalTunneltrafficpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*TunnelglobalTunneltrafficpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*TunnelglobalTunneltrafficpolicyBindingResource)(nil)

func NewTunnelglobalTunneltrafficpolicyBindingResource() resource.Resource {
	return &TunnelglobalTunneltrafficpolicyBindingResource{}
}

// TunnelglobalTunneltrafficpolicyBindingResource defines the resource implementation.
type TunnelglobalTunneltrafficpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tunnelglobal_tunneltrafficpolicy_binding"
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating tunnelglobal_tunneltrafficpolicy_binding resource")
	tunnelglobal_tunneltrafficpolicy_binding := tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), &tunnelglobal_tunneltrafficpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created tunnelglobal_tunneltrafficpolicy_binding resource")

	// Set ID for the resource before reading state.
	// Backward-compatible with SDK v2 which set the ID to the plain policyname value
	// (resource_id_mapping.json: "policyname"). type is a non-echoed bindpoint filter,
	// not part of the identity, so the ID stays a single plain value.
	data.Id = types.StringValue(data.Policyname.ValueString())

	// Read the updated state back
	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "tunnelglobal_tunneltrafficpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading tunnelglobal_tunneltrafficpolicy_binding resource")

	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating tunnelglobal_tunneltrafficpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		tunnelglobal_tunneltrafficpolicy_binding := tunnelglobal_tunneltrafficpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), &tunnelglobal_tunneltrafficpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated tunnelglobal_tunneltrafficpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for tunnelglobal_tunneltrafficpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "tunnelglobal_tunneltrafficpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *TunnelglobalTunneltrafficpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data TunnelglobalTunneltrafficpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting tunnelglobal_tunneltrafficpolicy_binding resource")
	// Global binding - delete using DeleteResourceWithArgs with empty resource name.
	// Single-key plain-value ID: the ID is the policyname. type/priority are extra
	// delete disambiguators (mirrors SDK v2 delete args). URL-encode slashy/special
	// values per DeleteResourceWithArgs convention.
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("policyname:%s", url.QueryEscape(data.Id.ValueString())))
	if !data.Type.IsNull() && data.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%s", url.QueryEscape(data.Type.ValueString())))
	}
	if !data.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", data.Priority.ValueInt64()))
	}

	err := r.client.DeleteResourceWithArgs(service.Tunnelglobal_tunneltrafficpolicy_binding.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted tunnelglobal_tunneltrafficpolicy_binding binding")
}

// Helper function to read tunnelglobal_tunneltrafficpolicy_binding data from API
func (r *TunnelglobalTunneltrafficpolicyBindingResource) readTunnelglobalTunneltrafficpolicyBindingFromApi(ctx context.Context, data *TunnelglobalTunneltrafficpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Single-key plain-value ID: the ID is the plain policyname (Pattern 10).
	policyname := data.Id.ValueString()

	var dataArr []map[string]interface{}
	var argsMap map[string]string = make(map[string]string)
	if !data.Type.IsNull() && data.Type.ValueString() != "" {
		argsMap["type"] = data.Type.ValueString()
	}

	findParams := service.FindParams{
		ResourceType:             service.Tunnelglobal_tunneltrafficpolicy_binding.Type(),
		ArgsMap:                  argsMap,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read tunnelglobal_tunneltrafficpolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching policyname
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policyname"].(string); ok && val == policyname {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	tunnelglobal_tunneltrafficpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
