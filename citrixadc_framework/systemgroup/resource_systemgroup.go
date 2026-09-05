package systemgroup

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/system"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NITRO binding type tokens (mirrors the SDK v2 implementation).
const (
	systemgroupCmdpolicyBindingType  = "systemgroup_systemcmdpolicy_binding"
	systemgroupSystemuserBindingType = "systemgroup_systemuser_binding"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SystemgroupResource{}
var _ resource.ResourceWithConfigure = (*SystemgroupResource)(nil)
var _ resource.ResourceWithImportState = (*SystemgroupResource)(nil)

func NewSystemgroupResource() resource.Resource {
	return &SystemgroupResource{}
}

// SystemgroupResource defines the resource implementation.
type SystemgroupResource struct {
	client *service.NitroClient
}

func (r *SystemgroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *SystemgroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_systemgroup"
}

func (r *SystemgroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *SystemgroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SystemgroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating systemgroup resource")

	groupname := data.Groupname.ValueString()
	systemgroup := systemgroupGetThePayloadFromthePlan(ctx, &data)

	// Named resource - create with AddResource (HTTP POST).
	_, err := r.client.AddResource(service.Systemgroup.Type(), groupname, &systemgroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create systemgroup, got error: %s", err))
		return
	}

	// Set ID before reading state back.
	data.Id = types.StringValue(groupname)

	// Manage cmdpolicy bindings only when explicitly configured.
	if !data.Cmdpolicybinding.IsNull() && !data.Cmdpolicybinding.IsUnknown() {
		resp.Diagnostics.Append(r.syncCmdpolicyBindings(ctx, groupname, types.SetNull(types.ObjectType{AttrTypes: cmdpolicybindingAttrTypes}), data.Cmdpolicybinding)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Manage systemuser bindings only when explicitly configured.
	if !data.Systemusers.IsNull() && !data.Systemusers.IsUnknown() {
		resp.Diagnostics.Append(r.syncSystemuserBindings(ctx, groupname, types.SetNull(types.StringType), data.Systemusers)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Trace(ctx, "Created systemgroup resource")

	// Read the updated state back
	if !r.readSystemgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemgroup not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SystemgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading systemgroup resource")

	found := r.readSystemgroupFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state SystemgroupResourceModel

	// Read Terraform plan and prior state data into the models
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating systemgroup resource")

	groupname := data.Groupname.ValueString()
	systemgroup := system.Systemgroup{Groupname: groupname}
	hasChange := false

	if !data.Warnpriorndays.Equal(state.Warnpriorndays) {
		systemgroup.Warnpriorndays = utils.IntPtr(int(data.Warnpriorndays.ValueInt64()))
		hasChange = true
	}
	if !data.Daystoexpire.Equal(state.Daystoexpire) {
		systemgroup.Daystoexpire = utils.IntPtr(int(data.Daystoexpire.ValueInt64()))
		hasChange = true
	}
	if !data.Promptstring.Equal(state.Promptstring) {
		systemgroup.Promptstring = data.Promptstring.ValueString()
		hasChange = true
	}
	if !data.Timeout.Equal(state.Timeout) {
		systemgroup.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
		hasChange = true
	}
	if !data.Allowedmanagementinterface.Equal(state.Allowedmanagementinterface) {
		var iface []string
		if !data.Allowedmanagementinterface.IsNull() && !data.Allowedmanagementinterface.IsUnknown() {
			data.Allowedmanagementinterface.ElementsAs(ctx, &iface, false)
		}
		systemgroup.Allowedmanagementinterface = iface
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Systemgroup.Type(), groupname, &systemgroup)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update systemgroup %s, got error: %s", groupname, err))
			return
		}
	}

	// Reconcile cmdpolicy bindings.
	if !data.Cmdpolicybinding.Equal(state.Cmdpolicybinding) {
		resp.Diagnostics.Append(r.syncCmdpolicyBindings(ctx, groupname, state.Cmdpolicybinding, data.Cmdpolicybinding)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Reconcile systemuser bindings.
	if !data.Systemusers.Equal(state.Systemusers) {
		resp.Diagnostics.Append(r.syncSystemuserBindings(ctx, groupname, state.Systemusers, data.Systemusers)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	tflog.Trace(ctx, "Updated systemgroup resource")

	// Read the updated state back
	if !r.readSystemgroupFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "systemgroup not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SystemgroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SystemgroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting systemgroup resource")

	// Named resource - delete using DeleteResource. Deleting the group cascades
	// to its bindings on the ADC.
	err := r.client.DeleteResource(service.Systemgroup.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete systemgroup %s, got error: %s", data.Id.ValueString(), err))
		return
	}

	tflog.Trace(ctx, "Deleted systemgroup resource")
}

// readSystemgroupFromApi reads the systemgroup (and any managed bindings) into
// the model. Returns false if the resource no longer exists on the ADC.
func (r *SystemgroupResource) readSystemgroupFromApi(ctx context.Context, data *SystemgroupResourceModel, diags *diag.Diagnostics) bool {
	groupname := data.Id.ValueString()
	if groupname == "" {
		groupname = data.Groupname.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Systemgroup.Type(), groupname)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read systemgroup %s, got error: %s", groupname, err))
		return false
	}
	if getResponseData == nil {
		return false
	}

	systemgroupSetAttrFromGet(ctx, data, getResponseData)

	// Refresh cmdpolicy bindings only when the resource manages them.
	if !data.Cmdpolicybinding.IsNull() && !data.Cmdpolicybinding.IsUnknown() {
		diags.Append(r.readCmdpolicyBindings(ctx, groupname, data)...)
	}

	// Refresh systemuser bindings only when the resource manages them.
	if !data.Systemusers.IsNull() && !data.Systemusers.IsUnknown() {
		diags.Append(r.readSystemuserBindings(ctx, groupname, data)...)
	}

	return true
}

// syncCmdpolicyBindings adds/removes systemgroup_systemcmdpolicy_binding entries
// so the live ADC state matches newSet. A (policyname, priority) pair is treated
// as an element: a priority change is a remove-then-add, mirroring SDK v2.
func (r *SystemgroupResource) syncCmdpolicyBindings(ctx context.Context, groupname string, oldSet, newSet types.Set) diag.Diagnostics {
	var diags diag.Diagnostics

	var oldBindings, newBindings []CmdpolicybindingModel
	if !oldSet.IsNull() && !oldSet.IsUnknown() {
		diags.Append(oldSet.ElementsAs(ctx, &oldBindings, false)...)
	}
	if !newSet.IsNull() && !newSet.IsUnknown() {
		diags.Append(newSet.ElementsAs(ctx, &newBindings, false)...)
	}
	if diags.HasError() {
		return diags
	}

	keyOf := func(b CmdpolicybindingModel) string {
		p := "null"
		if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
			p = fmt.Sprintf("%d", b.Priority.ValueInt64())
		}
		return b.Policyname.ValueString() + "|" + p
	}

	newKeys := make(map[string]bool)
	for _, b := range newBindings {
		newKeys[keyOf(b)] = true
	}
	oldKeys := make(map[string]bool)
	for _, b := range oldBindings {
		oldKeys[keyOf(b)] = true
	}

	// Remove bindings no longer present (delete first, so priority changes work).
	for _, b := range oldBindings {
		if newKeys[keyOf(b)] {
			continue
		}
		pn := b.Policyname.ValueString()
		if pn == "" {
			continue
		}
		args := []string{fmt.Sprintf("policyname:%s", pn)}
		if err := r.client.DeleteResourceWithArgs(systemgroupCmdpolicyBindingType, groupname, args); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to delete cmdpolicy binding %s from systemgroup %s: %s", pn, groupname, err))
			return diags
		}
	}

	// Add new bindings (HTTP PUT via UpdateResource, mirroring SDK v2).
	for _, b := range newBindings {
		if oldKeys[keyOf(b)] {
			continue
		}
		pn := b.Policyname.ValueString()
		if pn == "" {
			continue
		}
		bind := system.Systemgroupsystemcmdpolicybinding{
			Groupname:  groupname,
			Policyname: pn,
		}
		if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
			bind.Priority = utils.IntPtr(int(b.Priority.ValueInt64()))
		}
		if _, err := r.client.UpdateResource(systemgroupCmdpolicyBindingType, groupname, bind); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to add cmdpolicy binding %s to systemgroup %s: %s", pn, groupname, err))
			return diags
		}
	}

	return diags
}

// readCmdpolicyBindings populates data.Cmdpolicybinding from the ADC.
func (r *SystemgroupResource) readCmdpolicyBindings(ctx context.Context, groupname string, data *SystemgroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	bindings, _ := r.client.FindResourceArray(systemgroupCmdpolicyBindingType, groupname)
	elems := make([]CmdpolicybindingModel, 0, len(bindings))
	for _, m := range bindings {
		e := CmdpolicybindingModel{
			Policyname: types.StringNull(),
			Priority:   types.Int64Null(),
		}
		if v, ok := m["policyname"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Policyname = types.StringValue(s)
			}
		}
		if v, ok := m["priority"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Priority = types.Int64Value(iv)
			}
		}
		elems = append(elems, e)
	}

	setVal, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: cmdpolicybindingAttrTypes}, elems)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Cmdpolicybinding = setVal
	return diags
}

// syncSystemuserBindings adds/removes systemgroup_systemuser_binding entries so
// the live ADC state matches newSet.
func (r *SystemgroupResource) syncSystemuserBindings(ctx context.Context, groupname string, oldSet, newSet types.Set) diag.Diagnostics {
	var diags diag.Diagnostics

	var oldUsers, newUsers []string
	if !oldSet.IsNull() && !oldSet.IsUnknown() {
		diags.Append(oldSet.ElementsAs(ctx, &oldUsers, false)...)
	}
	if !newSet.IsNull() && !newSet.IsUnknown() {
		diags.Append(newSet.ElementsAs(ctx, &newUsers, false)...)
	}
	if diags.HasError() {
		return diags
	}

	newNames := make(map[string]bool)
	for _, u := range newUsers {
		newNames[u] = true
	}
	oldNames := make(map[string]bool)
	for _, u := range oldUsers {
		oldNames[u] = true
	}

	// Remove users no longer present.
	for _, u := range oldUsers {
		if u == "" || newNames[u] {
			continue
		}
		args := []string{fmt.Sprintf("username:%s", u)}
		if err := r.client.DeleteResourceWithArgs(systemgroupSystemuserBindingType, groupname, args); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to delete systemuser binding %s from systemgroup %s: %s", u, groupname, err))
			return diags
		}
	}

	// Add new users (HTTP PUT via UpdateResource, mirroring SDK v2).
	for _, u := range newUsers {
		if u == "" || oldNames[u] {
			continue
		}
		bind := system.Systemgroupsystemuserbinding{
			Groupname: groupname,
			Username:  u,
		}
		if _, err := r.client.UpdateResource(systemgroupSystemuserBindingType, groupname, bind); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to add systemuser binding %s to systemgroup %s: %s", u, groupname, err))
			return diags
		}
	}

	return diags
}

// readSystemuserBindings populates data.Systemusers from the ADC.
func (r *SystemgroupResource) readSystemuserBindings(ctx context.Context, groupname string, data *SystemgroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	bindings, _ := r.client.FindResourceArray(systemgroupSystemuserBindingType, groupname)
	elems := make([]string, 0, len(bindings))
	for _, m := range bindings {
		if v, ok := m["username"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				elems = append(elems, s)
			}
		}
	}

	setVal, d := types.SetValueFrom(ctx, types.StringType, elems)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Systemusers = setVal
	return diags
}
