package vxlan_nsip_binding

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
var _ resource.Resource = &VxlanNsipBindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanNsipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanNsipBindingResource)(nil)

func NewVxlanNsipBindingResource() resource.Resource {
	return &VxlanNsipBindingResource{}
}

// VxlanNsipBindingResource defines the resource implementation.
type VxlanNsipBindingResource struct {
	client *service.NitroClient
}

func (r *VxlanNsipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanNsipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan_nsip_binding"
}

func (r *VxlanNsipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanNsipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanNsipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan_nsip_binding resource")
	vxlan_nsip_binding := vxlan_nsip_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vxlan_nsip_binding.Type(), &vxlan_nsip_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vxlan_nsip_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vxlan_nsip_bindingComposeId(&data))

	// Read the updated state back
	r.readVxlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vxlan_nsip_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan_nsip_binding resource")

	r.readVxlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VxlanNsipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VxlanNsipBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; every attribute is
	// RequiresReplace, so Update is a documented no-op (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for vxlan_nsip_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVxlanNsipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vxlan_nsip_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanNsipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan_nsip_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// Legacy ID order: vxlanid,ipaddress.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vxlanid", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vxlanid, ok := idMap["vxlanid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vxlanid' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		argsMap["ipaddress"] = val
	}
	// netmask is not part of the ID; include it from state when present.
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		argsMap["netmask"] = data.Netmask.ValueString()
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vxlan_nsip_binding.Type(), vxlanid, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vxlan_nsip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vxlan_nsip_binding binding")
}

// Helper function to read vxlan_nsip_binding data from API
func (r *VxlanNsipBindingResource) readVxlanNsipBindingFromApi(ctx context.Context, data *VxlanNsipBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID. Legacy order: vxlanid,ipaddress.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vxlanid", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	vxlanid, ok := idMap["vxlanid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vxlanid' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vxlan_nsip_binding.Type(),
		ResourceName:             vxlanid,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan_nsip_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right ipaddress
	foundIndex := -1
	ipaddress := idMap["ipaddress"]
	for i, v := range dataArr {
		if val, ok := v["ipaddress"].(string); ok && val == ipaddress {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	vxlan_nsip_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
