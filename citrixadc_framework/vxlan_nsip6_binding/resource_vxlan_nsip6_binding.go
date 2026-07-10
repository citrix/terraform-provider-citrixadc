package vxlan_nsip6_binding

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
var _ resource.Resource = &VxlanNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*VxlanNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*VxlanNsip6BindingResource)(nil)

func NewVxlanNsip6BindingResource() resource.Resource {
	return &VxlanNsip6BindingResource{}
}

// VxlanNsip6BindingResource defines the resource implementation.
type VxlanNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *VxlanNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VxlanNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vxlan_nsip6_binding"
}

func (r *VxlanNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VxlanNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vxlan_nsip6_binding resource")
	vxlan_nsip6_binding := vxlan_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vxlan_nsip6_binding.Type(), &vxlan_nsip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vxlan_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vxlan_nsip6_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(vxlan_nsip6_bindingComposeId(&data))

	// Read the updated state back
	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vxlan_nsip6_binding resource")

	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VxlanNsip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// All attributes are RequiresReplace; no NITRO update endpoint exists for
	// this binding. Update is a documented no-op.
	tflog.Debug(ctx, "Update is a no-op for vxlan_nsip6_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readVxlanNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VxlanNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VxlanNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vxlan_nsip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ParseIdString handles both the new key:value form and the legacy
	// positional "vxlanid,ipaddress" form.
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

	args := make([]string, 0)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(val)))
	}
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", utils.UrlEncode(data.Netmask.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Vxlan_nsip6_binding.Type(), vxlanid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vxlan_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vxlan_nsip6_binding binding")
}

// Helper function to read vxlan_nsip6_binding data from API
func (r *VxlanNsip6BindingResource) readVxlanNsip6BindingFromApi(ctx context.Context, data *VxlanNsip6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID.
	// ParseIdString handles both the new key:value form and the legacy
	// positional "vxlanid,ipaddress" form.
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

	ipaddress, ok := idMap["ipaddress"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'ipaddress' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Vxlan_nsip6_binding.Type(),
		ResourceName:             vxlanid,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vxlan_nsip6_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vxlan_nsip6_binding returned empty array.")
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "vxlan_nsip6_binding not found with the provided ID attributes")
		return
	}

	vxlan_nsip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
