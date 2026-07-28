package vrid6_interface_binding

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
var _ resource.Resource = &Vrid6InterfaceBindingResource{}
var _ resource.ResourceWithConfigure = (*Vrid6InterfaceBindingResource)(nil)
var _ resource.ResourceWithImportState = (*Vrid6InterfaceBindingResource)(nil)

func NewVrid6InterfaceBindingResource() resource.Resource {
	return &Vrid6InterfaceBindingResource{}
}

// Vrid6InterfaceBindingResource defines the resource implementation.
type Vrid6InterfaceBindingResource struct {
	client *service.NitroClient
}

func (r *Vrid6InterfaceBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Vrid6InterfaceBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vrid6_interface_binding"
}

func (r *Vrid6InterfaceBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Vrid6InterfaceBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Vrid6InterfaceBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vrid6_interface_binding resource")
	vrid6_interface_binding := vrid6_interface_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vrid6_interface_binding.Type(), &vrid6_interface_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vrid6_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vrid6_interface_binding resource")

	// Set ID for the resource before reading state
	// Composite key: id,ifnum
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	found := r.readVrid6InterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "vrid6_interface_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Vrid6InterfaceBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Vrid6InterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vrid6_interface_binding resource")

	found := r.readVrid6InterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "vrid6_interface_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Vrid6InterfaceBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state Vrid6InterfaceBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for vrid6_interface_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for vrid6_interface_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readVrid6InterfaceBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Vrid6InterfaceBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Vrid6InterfaceBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vrid6_interface_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (id) as the
	// resource (URL) name and ifnum passed as a UrlEncoded arg (interface ids contain '/').
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	id_value, ok := idMap["id"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'id' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["ifnum"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ifnum:%s", utils.UrlEncode(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Vrid6_interface_binding.Type(), id_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vrid6_interface_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vrid6_interface_binding binding")
}

// vrid6_interface_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/vrid6_binding/<id>) and flattens the nested
// "vrid6_interface_binding" arrays into a single slice of binding rows.
//
// The direct by-name endpoint can return a keyless empty body, so the bound
// members are read from the parent aggregate for consistency and robustness.
func vrid6_interface_bindingAggregateRead(client *service.NitroClient, idValue string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "vrid6_binding",
		ResourceName:             idValue,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["vrid6_interface_binding"]
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

// readVrid6InterfaceBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (vrid6_binding/<id>) and matches the row by ifnum.
// It returns true when the binding is found and the model was populated, false
// when the binding is genuinely absent (drift). Hard errors (parse / transport)
// are reported via diags.
func (r *Vrid6InterfaceBindingResource) readVrid6InterfaceBindingFromApi(ctx context.Context, data *Vrid6InterfaceBindingResourceModel, diags *diag.Diagnostics) bool {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"id", "ifnum"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	id_Name, ok := idMap["id"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'id' not found in ID string")
		return false
	}

	dataArr, err := vrid6_interface_bindingAggregateRead(r.client, id_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vrid6_interface_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right member (ifnum).
	//
	// IMPORTANT (verified live on NS VPX): vrid6_interface_binding rows are of the
	// form {"id","vlan","flags"} and do NOT echo "ifnum". When a row carries an
	// ifnum, match on it; otherwise fall back to row presence (the parent vrid6 id
	// already scopes the result).
	foundIndex := -1
	wantIfnum := idMap["ifnum"]
	for i, v := range dataArr {
		if raw, ok := v["ifnum"]; ok && raw != nil {
			if val, ok := raw.(string); ok && val == wantIfnum {
				foundIndex = i
				break
			}
			continue
		}
		foundIndex = i
		break
	}

	// Binding row not present in the aggregate response: drift.
	if foundIndex == -1 {
		return false
	}

	vrid6_interface_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])

	// Backfill identity attributes from the parsed composite ID so that
	// `terraform import` (which starts with no prior state) fully round-trips.
	// Both vrid_id and ifnum are RequiresReplace identity values and are always
	// encoded in the ID. In particular, ifnum is NOT echoed by the aggregate GET
	// (rows are {"id","vlan","flags"}), so it can only be recovered from the ID.
	// This is done after the found/len self-heal checks so drift still nulls the
	// state, and it does not change data.Id (SetAttrFromGet does not rebuild it).
	if v, err := strconv.ParseInt(id_Name, 10, 64); err == nil {
		data.VridId = types.Int64Value(v)
	}
	if ifnum, ok := idMap["ifnum"]; ok {
		data.Ifnum = types.StringValue(ifnum)
	}
	return true
}
