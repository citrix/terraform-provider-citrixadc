package bridgegroup_nsip6_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BridgegroupNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*BridgegroupNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*BridgegroupNsip6BindingResource)(nil)

func NewBridgegroupNsip6BindingResource() resource.Resource {
	return &BridgegroupNsip6BindingResource{}
}

// BridgegroupNsip6BindingResource defines the resource implementation.
type BridgegroupNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *BridgegroupNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgegroupNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgegroup_nsip6_binding"
}

func (r *BridgegroupNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *BridgegroupNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgegroupNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgegroup_nsip6_binding resource")
	bridgegroup_nsip6_binding := bridgegroup_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Bridgegroup_nsip6_binding.Type(), &bridgegroup_nsip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgegroup_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created bridgegroup_nsip6_binding resource")

	// Set ID for the resource (legacy SDK v2 order: bridgegroup_id, ipaddress)
	data.Id = types.StringValue(bridgegroup_nsip6_bindingComposeId(&data))

	// Read the updated state back
	r.readBridgegroupNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "bridgegroup_nsip6_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgegroupNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgegroup_nsip6_binding resource")

	r.readBridgegroupNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BridgegroupNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state BridgegroupNsip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// All attributes are RequiresReplace; this binding has no NITRO update endpoint.
	tflog.Debug(ctx, "Update is a no-op for bridgegroup_nsip6_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readBridgegroupNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "bridgegroup_nsip6_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgegroupNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgegroupNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgegroup_nsip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"bridgegroup_id", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	bridgegroupId, ok := idMap["bridgegroup_id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'bridgegroup_id' not found in ID")
		return
	}

	// Delete args: ipaddress (required) plus any set optional unique attrs.
	args := make([]string, 0)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(val)))
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", utils.UrlEncode(data.Netmask.ValueString())))
	}
	if !data.Td.IsNull() {
		args = append(args, fmt.Sprintf("td:%d", data.Td.ValueInt64()))
	}
	if !data.Ownergroup.IsNull() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(data.Ownergroup.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Bridgegroup_nsip6_binding.Type(), bridgegroupId, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete bridgegroup_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted bridgegroup_nsip6_binding binding")
}

// Helper function to read bridgegroup_nsip6_binding data from API
func (r *BridgegroupNsip6BindingResource) readBridgegroupNsip6BindingFromApi(ctx context.Context, data *BridgegroupNsip6BindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID (legacy order: bridgegroup_id, ipaddress)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"bridgegroup_id", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	bridgegroupId, ok := idMap["bridgegroup_id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'bridgegroup_id' not found in ID string")
		return
	}

	ipaddress, ok := idMap["ipaddress"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ipaddress' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Bridgegroup_nsip6_binding.Type(),
		ResourceName:             bridgegroupId,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgegroup_nsip6_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the matching ipaddress
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ipaddress"].(string); ok && val == ipaddress {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	bridgegroup_nsip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
