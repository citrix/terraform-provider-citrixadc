package vlan_linkset_binding

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VlanLinksetBindingResource{}
var _ resource.ResourceWithConfigure = (*VlanLinksetBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VlanLinksetBindingResource)(nil)

func NewVlanLinksetBindingResource() resource.Resource {
	return &VlanLinksetBindingResource{}
}

// VlanLinksetBindingResource defines the resource implementation.
type VlanLinksetBindingResource struct {
	client *service.NitroClient
}

func (r *VlanLinksetBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VlanLinksetBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_linkset_binding"
}

func (r *VlanLinksetBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VlanLinksetBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VlanLinksetBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vlan_linkset_binding resource")
	vlan_linkset_binding := vlan_linkset_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vlan_linkset_binding.Type(), &vlan_linkset_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vlan_linkset_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vlan_linkset_binding resource")

	// Set ID for the resource before reading state.
	// Composite key: vlanid,ifnum (ownergroup/tagged are not part of the unique key).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vlanid:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlanid.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	found := r.readVlanLinksetBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "vlan_linkset_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanLinksetBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VlanLinksetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vlan_linkset_binding resource")

	found := r.readVlanLinksetBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "vlan_linkset_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VlanLinksetBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VlanLinksetBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vlan_linkset_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are
	// RequiresReplace, so Terraform recreates the resource on any change rather
	// than calling Update (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for vlan_linkset_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVlanLinksetBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// vlan_linkset_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/vlan_binding/<id>) and flattens the nested binding
// rows into a single slice.
//
// IMPORTANT (verified live on NS VPX): binding an interface through the
// vlan_linkset_binding endpoint causes the appliance to store the member as a
// "vlan_interface_binding" row in the aggregate vlan_binding/<id> response. The
// aggregate response never contains a "vlan_linkset_binding" array, and the
// direct vlan_linkset_binding/<id> endpoint returns a keyless empty body. The
// read therefore flattens the "vlan_interface_binding" nested key (the actual
// storage location of the bound member). Reading "vlan_linkset_binding" here
// always returned nothing, breaking Create's post-create read
// ("vlan_linkset_binding not found after create").
func vlan_linkset_bindingAggregateRead(client *service.NitroClient, vlanid string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vlan_binding",
		ResourceName:             vlanid,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["vlan_interface_binding"]
		if !ok || nested == nil {
			continue
		}
		nestedArr, ok := nested.([]interface{})
		if !ok {
			continue
		}
		for _, item := range nestedArr {
			if m, ok := item.(map[string]interface{}); ok {
				rows = append(rows, m)
			}
		}
	}
	return rows, nil
}

func (r *VlanLinksetBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VlanLinksetBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vlan_linkset_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent
	// (vlanid) as the resource (URL) name and ifnum/tagged/ownergroup passed as args.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	vlanid, ok := idMap["vlanid"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'vlanid' not found in ID")
		return
	}

	args := make([]string, 0)
	// ifnum value contains '/' (slot/port notation) and must be URL-encoded.
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() && data.Ifnum.ValueString() != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(data.Ifnum.ValueString())))
	}
	if !data.Tagged.IsNull() && !data.Tagged.IsUnknown() {
		args = append(args, fmt.Sprintf("tagged:%s", utils.UrlEncode(strconv.FormatBool(data.Tagged.ValueBool()))))
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", utils.UrlEncode(data.Ownergroup.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Vlan_linkset_binding.Type(), vlanid, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vlan_linkset_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vlan_linkset_binding binding")
}

// readVlanLinksetBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (vlan_binding/<vlanid>) and matches the row by ifnum.
// It returns true when the binding is found and the model was populated, false
// when the binding is genuinely absent (drift). Hard errors (parse / transport)
// are reported via diags.
func (r *VlanLinksetBindingResource) readVlanLinksetBindingFromApi(ctx context.Context, data *VlanLinksetBindingResourceModel, diags *diag.Diagnostics) bool {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"vlanid", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	vlanid_Name, ok := idMap["vlanid"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'vlanid' not found in ID string")
		return false
	}

	// The direct vlan_linkset_binding endpoint can return a keyless empty body;
	// read the bound interfaces from the aggregate parent endpoint instead.
	dataArr, err := vlan_linkset_bindingAggregateRead(r.client, vlanid_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vlan_linkset_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right ifnum (the binding member key).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ifnum (part of the composite key)
		if idVal, ok := idMap["ifnum"]; ok {
			if val, ok := v["ifnum"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["ifnum"].(string); ok {
			match = false
			continue
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	vlan_linkset_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
