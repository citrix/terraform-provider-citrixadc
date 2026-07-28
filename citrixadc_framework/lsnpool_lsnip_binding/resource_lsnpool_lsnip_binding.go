package lsnpool_lsnip_binding

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
var _ resource.Resource = &LsnpoolLsnipBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnpoolLsnipBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnpoolLsnipBindingResource)(nil)

func NewLsnpoolLsnipBindingResource() resource.Resource {
	return &LsnpoolLsnipBindingResource{}
}

// LsnpoolLsnipBindingResource defines the resource implementation.
type LsnpoolLsnipBindingResource struct {
	client *service.NitroClient
}

func (r *LsnpoolLsnipBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnpoolLsnipBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnpool_lsnip_binding"
}

func (r *LsnpoolLsnipBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnpoolLsnipBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnpoolLsnipBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnpool_lsnip_binding resource")
	lsnpool_lsnip_binding := lsnpool_lsnip_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add is HTTP PUT (bind), use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnpool_lsnip_binding.Type(), &lsnpool_lsnip_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnpool_lsnip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnpool_lsnip_binding resource")

	// Set ID for the resource before reading state
	// Composite key: poolname,lsnip (ownernode is cluster-only, not part of the unique key)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("poolname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Poolname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("lsnip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Lsnip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	found := r.readLsnpoolLsnipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.Diagnostics.AddError("Client Error", "lsnpool_lsnip_binding not found after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnpoolLsnipBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnpoolLsnipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnpool_lsnip_binding resource")

	found := r.readLsnpoolLsnipBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding genuinely absent on the appliance: treat as drift and clear state.
	if !found {
		tflog.Debug(ctx, "lsnpool_lsnip_binding not found on appliance; removing from state")
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnpoolLsnipBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnpoolLsnipBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for lsnpool_lsnip_binding: NITRO exposes only add (PUT)
	// and delete (no update/change endpoint), and all schema attributes are RequiresReplace, so Terraform
	// recreates the resource on any change rather than calling Update.
	tflog.Debug(ctx, "Update is a no-op for lsnpool_lsnip_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readLsnpoolLsnipBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// lsnpool_lsnip_bindingAggregateRead queries the AGGREGATE parent endpoint
// (GET /nitro/v1/config/lsnpool_binding/<poolname>) and flattens the nested
// "lsnpool_lsnip_binding" arrays into a single slice of binding rows.
//
// On firmware NS14.1 the direct endpoint
// (GET /nitro/v1/config/lsnpool_lsnip_binding/<poolname>) returns a keyless empty
// body, so the bound IP ranges are only retrievable via the parent aggregate.
func lsnpool_lsnip_bindingAggregateRead(client *service.NitroClient, poolname string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             "lsnpool_binding",
		ResourceName:             poolname,
		ResourceMissingErrorCode: 258,
	}
	parentArr, err := client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]interface{}, 0)
	for _, parent := range parentArr {
		nested, ok := parent["lsnpool_lsnip_binding"]
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

func (r *LsnpoolLsnipBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnpoolLsnipBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnpool_lsnip_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs with the parent (poolname) as the
	// resource (URL) name and lsnip (+ ownernode when set) passed as args.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"poolname", "lsnip"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	poolname, ok := idMap["poolname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'poolname' not found in ID")
		return
	}

	args := make([]string, 0)
	if val, ok := idMap["lsnip"]; ok && val != "" {
		args = append(args, fmt.Sprintf("lsnip:%s", utils.UrlEncode(val)))
	}
	// ownernode is not part of the composite ID; include it from state only when configured (cluster mode).
	if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
		args = append(args, fmt.Sprintf("ownernode:%s", utils.UrlEncode(strconv.FormatInt(data.Ownernode.ValueInt64(), 10))))
	}

	err = r.client.DeleteResourceWithArgs(service.Lsnpool_lsnip_binding.Type(), poolname, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnpool_lsnip_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnpool_lsnip_binding binding")
}

// readLsnpoolLsnipBindingFromApi reads the binding from the appliance via the
// AGGREGATE parent endpoint (lsnpool_binding/<poolname>) and matches the row by
// lsnip (and ownernode when configured). It returns true when the binding is
// found and the model was populated, false when the binding is genuinely absent
// (drift). Hard errors (parse / transport) are reported via diags.
func (r *LsnpoolLsnipBindingResource) readLsnpoolLsnipBindingFromApi(ctx context.Context, data *LsnpoolLsnipBindingResourceModel, diags *diag.Diagnostics) bool {

	// Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"poolname", "lsnip"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return false
	}

	poolname_Name, ok := idMap["poolname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'poolname' not found in ID string")
		return false
	}

	// The direct lsnpool_lsnip_binding endpoint returns a keyless empty body on
	// NS14.1; read the bound IPs from the aggregate parent endpoint instead.
	dataArr, err := lsnpool_lsnip_bindingAggregateRead(r.client, poolname_Name)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnpool_lsnip_binding, got error: %s", err))
		return false
	}

	// Binding genuinely absent (parent missing or no nested rows): report drift.
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check lsnip (part of the composite key)
		if idVal, ok := idMap["lsnip"]; ok {
			if val, ok := v["lsnip"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["lsnip"].(string); ok {
			match = false
			continue
		}

		// Check ownernode only when configured (cluster mode); it is not part of the composite ID.
		if !data.Ownernode.IsNull() && !data.Ownernode.IsUnknown() {
			if val, ok := v["ownernode"]; ok {
				intVal, _ := utils.ConvertToInt64(val)
				if intVal != data.Ownernode.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
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

	lsnpool_lsnip_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
	return true
}
