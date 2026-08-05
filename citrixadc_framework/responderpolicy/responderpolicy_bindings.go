package responderpolicy

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/responder"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// small conversion helpers
// ---------------------------------------------------------------------------

func rpAsString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

func rpAsInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	if iv, err := strconv.ParseInt(rpAsString(v), 10, 64); err == nil {
		return iv
	}
	return 0
}

// ---------------------------------------------------------------------------
// globalbinding (responderglobal_responderpolicy_binding)
// ---------------------------------------------------------------------------

func (r *ResponderpolicyResource) extractGlobalbindings(ctx context.Context, s types.Set) []ResponderpolicyGlobalbindingModel {
	out := []ResponderpolicyGlobalbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *ResponderpolicyResource) addSingleGlobalbinding(policyname string, b ResponderpolicyGlobalbindingModel) error {
	bindingStruct := responder.Responderglobalpolicybinding{}
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
	return r.client.UpdateUnnamedResource("responderglobal_responderpolicy_binding", bindingStruct)
}

func (r *ResponderpolicyResource) deleteSingleGlobalbinding(policyname string, b ResponderpolicyGlobalbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Type.IsNull() && b.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%v", b.Type.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	// responderglobal is a GLOBAL (unnamed) bind point. The resource name must be
	// empty here (matching the SDK v2 resource, whose binding map never populated
	// policyname). Passing the policy name makes the existence-check GET
	// (responderglobal_responderpolicy_binding/<name>) return errorcode 1090, so
	// DeleteResourceWithArgs would treat the binding as "already deleted" and skip
	// the DELETE. With an empty name the check falls back to a ?filter= GET that
	// succeeds, letting the actual DELETE fire.
	return r.client.DeleteResourceWithArgs("responderglobal_responderpolicy_binding", "", args)
}

func globalKey(b ResponderpolicyGlobalbindingModel) string {
	return fmt.Sprintf("%s|%d", b.Type.ValueString(), b.Priority.ValueInt64())
}

func (r *ResponderpolicyResource) syncGlobalbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
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

func (r *ResponderpolicyResource) deleteAllGlobalbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractGlobalbindings(ctx, s) {
		if err := r.deleteSingleGlobalbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *ResponderpolicyResource) readGlobalbindings(ctx context.Context, policyname string, data *ResponderpolicyResourceModel) {
	if data.Globalbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("responderpolicy_responderglobal_binding", policyname)
	elems := make([]ResponderpolicyGlobalbindingModel, 0, len(bindings))
	for _, val := range bindings {
		e := ResponderpolicyGlobalbindingModel{
			Gotopriorityexpression: types.StringValue(rpAsString(val["gotopriorityexpression"])),
			Labelname:              types.StringValue(rpAsString(val["labelname"])),
			Labeltype:              types.StringValue(rpAsString(val["labeltype"])),
			Policyname:             types.StringValue(policyname),
			Priority:               types.Int64Value(rpAsInt64(val["priority"])),
			Type:                   types.StringNull(),
		}
		// The bind type is encoded in the "boundto" field ("<flow> <type>").
		if boundto, ok := val["boundto"]; ok && boundto != nil {
			boundtoSlice := strings.Split(rpAsString(boundto), " ")
			if len(boundtoSlice) > 1 {
				e.Type = types.StringValue(boundtoSlice[1])
			}
		}
		// Normalize labeltype the way the SDK v2 resource did: standalone
		// (reqvserver/resvserver) and cluster ("") map back to "vserver".
		lt := e.Labeltype.ValueString()
		if lt == "reqvserver" || lt == "resvserver" || lt == "" {
			e.Labeltype = types.StringValue("vserver")
		}
		// Deduce invoke from the presence of a label (NITRO does not echo it).
		invoke := false
		if e.Labelname.ValueString() != "" || e.Labeltype.ValueString() != "" {
			invoke = true
		}
		e.Invoke = types.BoolValue(invoke)
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: responderpolicyGlobalbindingAttrTypes}, elems)
	data.Globalbinding = setVal
}

// ---------------------------------------------------------------------------
// lbvserverbinding (lbvserver_responderpolicy_binding)
// ---------------------------------------------------------------------------

func (r *ResponderpolicyResource) extractLbvserverbindings(ctx context.Context, s types.Set) []ResponderpolicyLbvserverbindingModel {
	out := []ResponderpolicyLbvserverbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *ResponderpolicyResource) addSingleLbvserverbinding(policyname string, b ResponderpolicyLbvserverbindingModel) error {
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
	_, err := r.client.UpdateResource("lbvserver_responderpolicy_binding", b.Name.ValueString(), bindingStruct)
	return err
}

func (r *ResponderpolicyResource) deleteSingleLbvserverbinding(policyname string, b ResponderpolicyLbvserverbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Bindpoint.IsNull() && b.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%v", b.Bindpoint.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	return r.client.DeleteResourceWithArgs("lbvserver_responderpolicy_binding", b.Name.ValueString(), args)
}

func lbvserverKey(b ResponderpolicyLbvserverbindingModel) string {
	return fmt.Sprintf("%s|%s|%d", b.Name.ValueString(), b.Bindpoint.ValueString(), b.Priority.ValueInt64())
}

func (r *ResponderpolicyResource) syncLbvserverbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
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

func (r *ResponderpolicyResource) deleteAllLbvserverbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractLbvserverbindings(ctx, s) {
		if err := r.deleteSingleLbvserverbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *ResponderpolicyResource) readLbvserverbindings(ctx context.Context, policyname string, data *ResponderpolicyResourceModel) {
	if data.Lbvserverbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("responderpolicy_lbvserver_binding", policyname)
	elems := make([]ResponderpolicyLbvserverbindingModel, 0, len(bindings))
	for _, val := range bindings {
		boundtoSlice := strings.Split(rpAsString(val["boundto"]), " ")
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
		e := ResponderpolicyLbvserverbindingModel{
			Bindpoint:              types.StringValue(bindpoint),
			Gotopriorityexpression: types.StringNull(),
			Invoke:                 types.BoolValue(false),
			Labelname:              types.StringNull(),
			Labeltype:              types.StringNull(),
			Name:                   types.StringValue(vserverName),
			Priority:               types.Int64Value(0),
		}
		// Complete the record from the lb vserver side of the binding.
		vserverBindings, _ := r.client.FindResourceArray("lbvserver_responderpolicy_binding", vserverName)
		for _, vb := range vserverBindings {
			if rpAsString(vb["policyname"]) == policyname {
				e.Gotopriorityexpression = types.StringValue(rpAsString(vb["gotopriorityexpression"]))
				e.Labelname = types.StringValue(rpAsString(vb["labelname"]))
				e.Labeltype = types.StringValue(rpAsString(vb["labeltype"]))
				e.Priority = types.Int64Value(rpAsInt64(vb["priority"]))
				if inv, ok := vb["invoke"]; ok {
					e.Invoke = types.BoolValue(inv == true || inv == "true")
				}
				break
			}
		}
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: responderpolicyLbvserverbindingAttrTypes}, elems)
	data.Lbvserverbinding = setVal
}

// ---------------------------------------------------------------------------
// csvserverbinding (csvserver_responderpolicy_binding)
// ---------------------------------------------------------------------------

func (r *ResponderpolicyResource) extractCsvserverbindings(ctx context.Context, s types.Set) []ResponderpolicyCsvserverbindingModel {
	out := []ResponderpolicyCsvserverbindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *ResponderpolicyResource) addSingleCsvserverbinding(policyname string, b ResponderpolicyCsvserverbindingModel) error {
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
	_, err := r.client.UpdateResource("csvserver_responderpolicy_binding", b.Name.ValueString(), bindingStruct)
	return err
}

func (r *ResponderpolicyResource) deleteSingleCsvserverbinding(policyname string, b ResponderpolicyCsvserverbindingModel) error {
	args := make([]string, 0, 3)
	args = append(args, fmt.Sprintf("policyname:%v", policyname))
	if !b.Bindpoint.IsNull() && b.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%v", b.Bindpoint.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	return r.client.DeleteResourceWithArgs("csvserver_responderpolicy_binding", b.Name.ValueString(), args)
}

func csvserverKey(b ResponderpolicyCsvserverbindingModel) string {
	return fmt.Sprintf("%s|%s|%d", b.Name.ValueString(), b.Bindpoint.ValueString(), b.Priority.ValueInt64())
}

func (r *ResponderpolicyResource) syncCsvserverbindings(ctx context.Context, policyname string, oldSet, newSet types.Set) error {
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

func (r *ResponderpolicyResource) deleteAllCsvserverbindings(ctx context.Context, policyname string, s types.Set) error {
	for _, b := range r.extractCsvserverbindings(ctx, s) {
		if err := r.deleteSingleCsvserverbinding(policyname, b); err != nil {
			return err
		}
	}
	return nil
}

func (r *ResponderpolicyResource) readCsvserverbindings(ctx context.Context, policyname string, data *ResponderpolicyResourceModel) {
	if data.Csvserverbinding.IsNull() {
		return
	}
	bindings, _ := r.client.FindResourceArray("responderpolicy_csvserver_binding", policyname)
	elems := make([]ResponderpolicyCsvserverbindingModel, 0, len(bindings))
	for _, val := range bindings {
		boundtoSlice := strings.Split(rpAsString(val["boundto"]), " ")
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
		e := ResponderpolicyCsvserverbindingModel{
			Bindpoint:              types.StringValue(bindpoint),
			Gotopriorityexpression: types.StringNull(),
			Invoke:                 types.BoolValue(false),
			Labelname:              types.StringNull(),
			Labeltype:              types.StringNull(),
			Name:                   types.StringValue(vserverName),
			Policyname:             types.StringValue(policyname),
			Priority:               types.Int64Value(0),
			Targetlbvserver:        types.StringNull(),
		}
		vserverBindings, _ := r.client.FindResourceArray("csvserver_responderpolicy_binding", vserverName)
		for _, vb := range vserverBindings {
			if rpAsString(vb["policyname"]) == policyname {
				e.Gotopriorityexpression = types.StringValue(rpAsString(vb["gotopriorityexpression"]))
				e.Labelname = types.StringValue(rpAsString(vb["labelname"]))
				e.Labeltype = types.StringValue(rpAsString(vb["labeltype"]))
				e.Priority = types.Int64Value(rpAsInt64(vb["priority"]))
				if inv, ok := vb["invoke"]; ok {
					e.Invoke = types.BoolValue(inv == true || inv == "true")
				}
				if tv, ok := vb["targetlbvserver"]; ok && tv != nil {
					e.Targetlbvserver = types.StringValue(rpAsString(tv))
				}
				break
			}
		}
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: responderpolicyCsvserverbindingAttrTypes}, elems)
	data.Csvserverbinding = setVal
}

// ---------------------------------------------------------------------------
// combined apply/read orchestration
// ---------------------------------------------------------------------------

// applyBindingsOnCreate binds all convenience blocks after the base responderpolicy is added.
func (r *ResponderpolicyResource) applyBindingsOnCreate(ctx context.Context, policyname string, data *ResponderpolicyResourceModel, diags *diag.Diagnostics) {
	if !data.Globalbinding.IsNull() {
		if err := r.syncGlobalbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: responderpolicyGlobalbindingAttrTypes}), data.Globalbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind globalbinding for responderpolicy: %s", err))
			return
		}
	}
	if !data.Lbvserverbinding.IsNull() {
		if err := r.syncLbvserverbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: responderpolicyLbvserverbindingAttrTypes}), data.Lbvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind lbvserverbinding for responderpolicy: %s", err))
			return
		}
	}
	if !data.Csvserverbinding.IsNull() {
		if err := r.syncCsvserverbindings(ctx, policyname, types.SetNull(types.ObjectType{AttrTypes: responderpolicyCsvserverbindingAttrTypes}), data.Csvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind csvserverbinding for responderpolicy: %s", err))
			return
		}
	}
}

// applyBindingsOnUpdate reconciles all convenience blocks against prior state.
func (r *ResponderpolicyResource) applyBindingsOnUpdate(ctx context.Context, policyname string, data, state *ResponderpolicyResourceModel, diags *diag.Diagnostics) {
	if !data.Globalbinding.Equal(state.Globalbinding) {
		if err := r.syncGlobalbindings(ctx, policyname, state.Globalbinding, data.Globalbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync globalbinding for responderpolicy: %s", err))
			return
		}
	}
	if !data.Lbvserverbinding.Equal(state.Lbvserverbinding) {
		if err := r.syncLbvserverbindings(ctx, policyname, state.Lbvserverbinding, data.Lbvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync lbvserverbinding for responderpolicy: %s", err))
			return
		}
	}
	if !data.Csvserverbinding.Equal(state.Csvserverbinding) {
		if err := r.syncCsvserverbindings(ctx, policyname, state.Csvserverbinding, data.Csvserverbinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync csvserverbinding for responderpolicy: %s", err))
			return
		}
	}
}

// deleteAllBindings removes every managed convenience-block binding prior to
// deleting the base responderpolicy (bindings must be removed first).
func (r *ResponderpolicyResource) deleteAllBindings(ctx context.Context, policyname string, data *ResponderpolicyResourceModel, diags *diag.Diagnostics) {
	if err := r.deleteAllGlobalbindings(ctx, policyname, data.Globalbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete globalbinding for responderpolicy: %s", err))
		return
	}
	if err := r.deleteAllLbvserverbindings(ctx, policyname, data.Lbvserverbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserverbinding for responderpolicy: %s", err))
		return
	}
	if err := r.deleteAllCsvserverbindings(ctx, policyname, data.Csvserverbinding); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to delete csvserverbinding for responderpolicy: %s", err))
		return
	}
}

// readBindings refreshes all managed convenience blocks from the appliance.
func (r *ResponderpolicyResource) readBindings(ctx context.Context, policyname string, data *ResponderpolicyResourceModel) {
	r.readGlobalbindings(ctx, policyname, data)
	r.readLbvserverbindings(ctx, policyname, data)
	r.readCsvserverbindings(ctx, policyname, data)
}
