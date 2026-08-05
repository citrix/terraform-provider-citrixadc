package rewritepolicy

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/rewrite"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// small conversion helpers
// ---------------------------------------------------------------------------

func rwAsString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

func rwAsInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	if iv, err := strconv.ParseInt(rwAsString(v), 10, 64); err == nil {
		return iv
	}
	return 0
}

// ---------------------------------------------------------------------------
// globalbinding (rewriteglobal_rewritepolicy_binding)
// ---------------------------------------------------------------------------

func (r *RewritepolicyResource) extractGlobalbindings(ctx context.Context, s types.Set) []RewritepolicyGlobalbindingModel {
	out := []RewritepolicyGlobalbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *RewritepolicyResource) addSingleGlobalbinding(policyname string, b RewritepolicyGlobalbindingModel) error {
	bindingStruct := rewrite.Rewriteglobalpolicybinding{}
	bindingStruct.Policyname = policyname
	if !b.Gotopriorityexpression.IsNull() && !b.Gotopriorityexpression.IsUnknown() {
		bindingStruct.Gotopriorityexpression = b.Gotopriorityexpression.ValueString()
	}
	if !b.Invoke.IsNull() && !b.Invoke.IsUnknown() {
		bindingStruct.Invoke = b.Invoke.ValueBool()
	}
	if !b.Labelname.IsNull() && !b.Labelname.IsUnknown() {
		bindingStruct.Labelname = b.Labelname.ValueString()
	}
	if !b.Labeltype.IsNull() && !b.Labeltype.IsUnknown() {
		bindingStruct.Labeltype = b.Labeltype.ValueString()
	}
	if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
		bindingStruct.Priority = uint32(b.Priority.ValueInt64())
	}
	if !b.Type.IsNull() && !b.Type.IsUnknown() {
		bindingStruct.Type = b.Type.ValueString()
	}
	return r.client.UpdateUnnamedResource("rewriteglobal_rewritepolicy_binding", bindingStruct)
}

func (r *RewritepolicyResource) deleteSingleGlobalbinding(policyname string, b RewritepolicyGlobalbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Type.IsNull() && b.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%v", b.Type.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	// rewriteglobal_rewritepolicy_binding is a global binding with NO name segment.
	// The resourceName MUST be empty so DeleteResourceWithArgs issues the request
	// against /rewriteglobal_rewritepolicy_binding?args=... . Passing policyname as
	// the name segment (/rewriteglobal_rewritepolicy_binding/<policyname>?args=...)
	// makes the internal existence-check GET fail (errorcode 1090 "No such argument
	// [arguid]"), which DeleteResourceWithArgs interprets as "already deleted" and
	// silently skips the delete, leaving the policy bound (errorcode 810 on the
	// subsequent policy delete).
	return r.client.DeleteResourceWithArgs("rewriteglobal_rewritepolicy_binding", "", args)
}

func globalKey(b RewritepolicyGlobalbindingModel) string {
	return fmt.Sprintf("%s|%d", b.Type.ValueString(), b.Priority.ValueInt64())
}

func (r *RewritepolicyResource) syncGlobalbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
	oldB := r.extractGlobalbindings(ctx, oldSet)
	newB := r.extractGlobalbindings(ctx, newSet)
	newKeys := map[string]bool{}
	for _, b := range newB {
		newKeys[globalKey(b)] = true
	}
	oldKeys := map[string]bool{}
	for _, b := range oldB {
		oldKeys[globalKey(b)] = true
	}
	for _, b := range oldB {
		if !newKeys[globalKey(b)] {
			if err := r.deleteSingleGlobalbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	for _, b := range newB {
		if !oldKeys[globalKey(b)] {
			if err := r.addSingleGlobalbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RewritepolicyResource) deleteAllGlobalbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractGlobalbindings(ctx, s) {
		if err := r.deleteSingleGlobalbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *RewritepolicyResource) readGlobalbindings(ctx context.Context, policyname string, data *RewritepolicyResourceModel) {
	if data.Globalbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("rewritepolicy_rewriteglobal_binding", policyname)
	elems := make([]RewritepolicyGlobalbindingModel, 0, len(bindings))
	for _, val := range bindings {
		e := RewritepolicyGlobalbindingModel{
			Globalbindtype:         types.StringNull(),
			Gotopriorityexpression: types.StringValue(rwAsString(val["gotopriorityexpression"])),
			Labelname:              types.StringValue(rwAsString(val["labelname"])),
			Labeltype:              types.StringValue(rwAsString(val["labeltype"])),
			Policyname:             types.StringValue(policyname),
			Priority:               types.Int64Value(rwAsInt64(val["priority"])),
			Type:                   types.StringNull(),
		}
		// The bind type is encoded in the "boundto" field ("<flow> <type>").
		if boundto, ok := val["boundto"]; ok && boundto != nil {
			boundtoSlice := strings.Split(rwAsString(boundto), " ")
			if len(boundtoSlice) > 1 {
				e.Type = types.StringValue(boundtoSlice[1])
			}
		}
		// Deduce invoke from the presence of a label (NITRO does not echo it).
		invoke := false
		if e.Labelname.ValueString() != "" || e.Labeltype.ValueString() != "" {
			invoke = true
		}
		e.Invoke = types.BoolValue(invoke)
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: rewritepolicyGlobalbindingAttrTypes}, elems)
	data.Globalbinding = setVal
}

// ---------------------------------------------------------------------------
// lbvserverbinding (lbvserver_rewritepolicy_binding)
// ---------------------------------------------------------------------------

func (r *RewritepolicyResource) extractLbvserverbindings(ctx context.Context, s types.Set) []RewritepolicyLbvserverbindingModel {
	out := []RewritepolicyLbvserverbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *RewritepolicyResource) addSingleLbvserverbinding(policyname string, b RewritepolicyLbvserverbindingModel) error {
	bindingStruct := lb.Lbvserverpolicybinding{}
	bindingStruct.Policyname = policyname
	if !b.Bindpoint.IsNull() && !b.Bindpoint.IsUnknown() {
		bindingStruct.Bindpoint = b.Bindpoint.ValueString()
	}
	if !b.Gotopriorityexpression.IsNull() && !b.Gotopriorityexpression.IsUnknown() {
		bindingStruct.Gotopriorityexpression = b.Gotopriorityexpression.ValueString()
	}
	if !b.Invoke.IsNull() && !b.Invoke.IsUnknown() {
		bindingStruct.Invoke = b.Invoke.ValueBool()
	}
	if !b.Labelname.IsNull() && !b.Labelname.IsUnknown() {
		bindingStruct.Labelname = b.Labelname.ValueString()
	}
	if !b.Labeltype.IsNull() && !b.Labeltype.IsUnknown() {
		bindingStruct.Labeltype = b.Labeltype.ValueString()
	}
	if !b.Name.IsNull() && !b.Name.IsUnknown() {
		bindingStruct.Name = b.Name.ValueString()
	}
	if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
		bindingStruct.Priority = uint32(b.Priority.ValueInt64())
	}
	// The binding is keyed on the lb vserver name; PUT (UpdateResource) semantics.
	_, err := r.client.UpdateResource("lbvserver_rewritepolicy_binding", b.Name.ValueString(), bindingStruct)
	return err
}

func (r *RewritepolicyResource) deleteSingleLbvserverbinding(policyname string, b RewritepolicyLbvserverbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Bindpoint.IsNull() && b.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%v", b.Bindpoint.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	return r.client.DeleteResourceWithArgs("lbvserver_rewritepolicy_binding", b.Name.ValueString(), args)
}

func lbvserverKey(b RewritepolicyLbvserverbindingModel) string {
	return fmt.Sprintf("%s|%s|%d", b.Name.ValueString(), b.Bindpoint.ValueString(), b.Priority.ValueInt64())
}

func (r *RewritepolicyResource) syncLbvserverbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
	oldB := r.extractLbvserverbindings(ctx, oldSet)
	newB := r.extractLbvserverbindings(ctx, newSet)
	newKeys := map[string]bool{}
	for _, b := range newB {
		newKeys[lbvserverKey(b)] = true
	}
	oldKeys := map[string]bool{}
	for _, b := range oldB {
		oldKeys[lbvserverKey(b)] = true
	}
	for _, b := range oldB {
		if !newKeys[lbvserverKey(b)] {
			if err := r.deleteSingleLbvserverbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	for _, b := range newB {
		if !oldKeys[lbvserverKey(b)] {
			if err := r.addSingleLbvserverbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RewritepolicyResource) deleteAllLbvserverbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractLbvserverbindings(ctx, s) {
		if err := r.deleteSingleLbvserverbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *RewritepolicyResource) readLbvserverbindings(ctx context.Context, policyname string, data *RewritepolicyResourceModel) {
	if data.Lbvserverbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("rewritepolicy_lbvserver_binding", policyname)
	elems := make([]RewritepolicyLbvserverbindingModel, 0, len(bindings))
	for _, val := range bindings {
		boundtoSlice := strings.Split(rwAsString(val["boundto"]), " ")
		bindpoint := ""
		vserverName := ""
		if len(boundtoSlice) > 2 {
			switch boundtoSlice[0] {
			case "REQ", "REQUEST":
				bindpoint = "REQUEST"
			case "RES", "RESPONSE":
				bindpoint = "RESPONSE"
			}
			vserverName = boundtoSlice[2]
		}
		e := RewritepolicyLbvserverbindingModel{
			Bindpoint:              types.StringValue(bindpoint),
			Gotopriorityexpression: types.StringNull(),
			Invoke:                 types.BoolValue(false),
			Labelname:              types.StringNull(),
			Labeltype:              types.StringNull(),
			Name:                   types.StringValue(vserverName),
			Priority:               types.Int64Value(0),
		}
		// Complete the record from the lb vserver side of the binding.
		vserverBindings, _ := r.client.FindResourceArray("lbvserver_rewritepolicy_binding", vserverName)
		for _, vb := range vserverBindings {
			if rwAsString(vb["policyname"]) == policyname {
				e.Gotopriorityexpression = types.StringValue(rwAsString(vb["gotopriorityexpression"]))
				e.Labelname = types.StringValue(rwAsString(vb["labelname"]))
				e.Labeltype = types.StringValue(rwAsString(vb["labeltype"]))
				e.Priority = types.Int64Value(rwAsInt64(vb["priority"]))
				if inv, ok := vb["invoke"]; ok {
					e.Invoke = types.BoolValue(inv == true || inv == "true")
				}
				break
			}
		}
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: rewritepolicyLbvserverbindingAttrTypes}, elems)
	data.Lbvserverbinding = setVal
}

// ---------------------------------------------------------------------------
// csvserverbinding (csvserver_rewritepolicy_binding)
// ---------------------------------------------------------------------------

func (r *RewritepolicyResource) extractCsvserverbindings(ctx context.Context, s types.Set) []RewritepolicyCsvserverbindingModel {
	out := []RewritepolicyCsvserverbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *RewritepolicyResource) addSingleCsvserverbinding(policyname string, b RewritepolicyCsvserverbindingModel) error {
	bindingStruct := cs.Csvserverpolicybinding{}
	bindingStruct.Policyname = policyname
	if !b.Bindpoint.IsNull() && !b.Bindpoint.IsUnknown() {
		bindingStruct.Bindpoint = b.Bindpoint.ValueString()
	}
	if !b.Gotopriorityexpression.IsNull() && !b.Gotopriorityexpression.IsUnknown() {
		bindingStruct.Gotopriorityexpression = b.Gotopriorityexpression.ValueString()
	}
	if !b.Invoke.IsNull() && !b.Invoke.IsUnknown() {
		bindingStruct.Invoke = b.Invoke.ValueBool()
	}
	if !b.Labelname.IsNull() && !b.Labelname.IsUnknown() {
		bindingStruct.Labelname = b.Labelname.ValueString()
	}
	if !b.Labeltype.IsNull() && !b.Labeltype.IsUnknown() {
		bindingStruct.Labeltype = b.Labeltype.ValueString()
	}
	if !b.Name.IsNull() && !b.Name.IsUnknown() {
		bindingStruct.Name = b.Name.ValueString()
	}
	if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
		bindingStruct.Priority = uint32(b.Priority.ValueInt64())
	}
	if !b.Targetlbvserver.IsNull() && !b.Targetlbvserver.IsUnknown() {
		bindingStruct.Targetlbvserver = b.Targetlbvserver.ValueString()
	}
	// The binding is keyed on the cs vserver name; PUT (UpdateResource) semantics.
	_, err := r.client.UpdateResource("csvserver_rewritepolicy_binding", b.Name.ValueString(), bindingStruct)
	return err
}

func (r *RewritepolicyResource) deleteSingleCsvserverbinding(policyname string, b RewritepolicyCsvserverbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Bindpoint.IsNull() && b.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%v", b.Bindpoint.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	return r.client.DeleteResourceWithArgs("csvserver_rewritepolicy_binding", b.Name.ValueString(), args)
}

func csvserverKey(b RewritepolicyCsvserverbindingModel) string {
	return fmt.Sprintf("%s|%s|%d", b.Name.ValueString(), b.Bindpoint.ValueString(), b.Priority.ValueInt64())
}

func (r *RewritepolicyResource) syncCsvserverbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
	oldB := r.extractCsvserverbindings(ctx, oldSet)
	newB := r.extractCsvserverbindings(ctx, newSet)
	newKeys := map[string]bool{}
	for _, b := range newB {
		newKeys[csvserverKey(b)] = true
	}
	oldKeys := map[string]bool{}
	for _, b := range oldB {
		oldKeys[csvserverKey(b)] = true
	}
	for _, b := range oldB {
		if !newKeys[csvserverKey(b)] {
			if err := r.deleteSingleCsvserverbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	for _, b := range newB {
		if !oldKeys[csvserverKey(b)] {
			if err := r.addSingleCsvserverbinding(policyname, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RewritepolicyResource) deleteAllCsvserverbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractCsvserverbindings(ctx, s) {
		if err := r.deleteSingleCsvserverbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *RewritepolicyResource) readCsvserverbindings(ctx context.Context, policyname string, data *RewritepolicyResourceModel) {
	if data.Csvserverbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("rewritepolicy_csvserver_binding", policyname)
	elems := make([]RewritepolicyCsvserverbindingModel, 0, len(bindings))
	for _, val := range bindings {
		boundtoSlice := strings.Split(rwAsString(val["boundto"]), " ")
		bindpoint := ""
		vserverName := ""
		if len(boundtoSlice) > 2 {
			switch boundtoSlice[0] {
			case "REQ", "REQUEST":
				bindpoint = "REQUEST"
			case "RES", "RESPONSE":
				bindpoint = "RESPONSE"
			}
			vserverName = boundtoSlice[2]
		}
		e := RewritepolicyCsvserverbindingModel{
			Bindpoint:              types.StringValue(bindpoint),
			Gotopriorityexpression: types.StringNull(),
			Invoke:                 types.BoolValue(false),
			Labelname:              types.StringNull(),
			Labeltype:              types.StringNull(),
			Name:                   types.StringValue(vserverName),
			Priority:               types.Int64Value(0),
			Targetlbvserver:        types.StringNull(),
		}
		vserverBindings, _ := r.client.FindResourceArray("csvserver_rewritepolicy_binding", vserverName)
		for _, vb := range vserverBindings {
			if rwAsString(vb["policyname"]) == policyname {
				e.Gotopriorityexpression = types.StringValue(rwAsString(vb["gotopriorityexpression"]))
				e.Labelname = types.StringValue(rwAsString(vb["labelname"]))
				e.Labeltype = types.StringValue(rwAsString(vb["labeltype"]))
				e.Priority = types.Int64Value(rwAsInt64(vb["priority"]))
				if inv, ok := vb["invoke"]; ok {
					e.Invoke = types.BoolValue(inv == true || inv == "true")
				}
				if tv, ok := vb["targetlbvserver"]; ok && tv != nil {
					e.Targetlbvserver = types.StringValue(rwAsString(tv))
				}
				break
			}
		}
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: rewritepolicyCsvserverbindingAttrTypes}, elems)
	data.Csvserverbinding = setVal
}

// ---------------------------------------------------------------------------
// combined apply/read orchestration
// ---------------------------------------------------------------------------

// applyBindingsOnCreate binds all convenience blocks after the base rewritepolicy is added.
func (r *RewritepolicyResource) applyBindingsOnCreate(ctx context.Context, policyname string, data *RewritepolicyResourceModel, diags *diag.Diagnostics) {
	if !data.Globalbinding.IsNull() {
		if err := r.syncGlobalbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: rewritepolicyGlobalbindingAttrTypes}), data.Globalbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind globalbinding for rewritepolicy: %s", err))
			return
		}
	}
	if !data.Lbvserverbinding.IsNull() {
		if err := r.syncLbvserverbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: rewritepolicyLbvserverbindingAttrTypes}), data.Lbvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind lbvserverbinding for rewritepolicy: %s", err))
			return
		}
	}
	if !data.Csvserverbinding.IsNull() {
		if err := r.syncCsvserverbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: rewritepolicyCsvserverbindingAttrTypes}), data.Csvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind csvserverbinding for rewritepolicy: %s", err))
			return
		}
	}
}

// applyBindingsOnUpdate reconciles all convenience blocks against prior state.
func (r *RewritepolicyResource) applyBindingsOnUpdate(ctx context.Context, policyname string, data, state *RewritepolicyResourceModel, diags *diag.Diagnostics) {
	if !data.Globalbinding.Equal(state.Globalbinding) {
		if err := r.syncGlobalbindings(ctx, policyname, state.Globalbinding, data.Globalbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync globalbinding for rewritepolicy: %s", err))
			return
		}
	}
	if !data.Lbvserverbinding.Equal(state.Lbvserverbinding) {
		if err := r.syncLbvserverbindings(ctx, policyname, state.Lbvserverbinding, data.Lbvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync lbvserverbinding for rewritepolicy: %s", err))
			return
		}
	}
	if !data.Csvserverbinding.Equal(state.Csvserverbinding) {
		if err := r.syncCsvserverbindings(ctx, policyname, state.Csvserverbinding, data.Csvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync csvserverbinding for rewritepolicy: %s", err))
			return
		}
	}
}

// deleteAllBindings removes every managed convenience-block binding prior to
// deleting the base rewritepolicy (bindings must be removed first).
func (r *RewritepolicyResource) deleteAllBindings(ctx context.Context, policyname string, data *RewritepolicyResourceModel, diags *diag.Diagnostics) {
	if err := r.deleteAllGlobalbindings(ctx, policyname, data.Globalbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete globalbinding for rewritepolicy: %s", err))
		return
	}
	if err := r.deleteAllLbvserverbindings(ctx, policyname, data.Lbvserverbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserverbinding for rewritepolicy: %s", err))
		return
	}
	if err := r.deleteAllCsvserverbindings(ctx, policyname, data.Csvserverbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete csvserverbinding for rewritepolicy: %s", err))
		return
	}
}

// readBindings refreshes all managed convenience blocks from the appliance.
func (r *RewritepolicyResource) readBindings(ctx context.Context, policyname string, data *RewritepolicyResourceModel) {
	r.readGlobalbindings(ctx, policyname, data)
	r.readLbvserverbindings(ctx, policyname, data)
	r.readCsvserverbindings(ctx, policyname, data)
}
