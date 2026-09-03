package lbvserver

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ---------------------------------------------------------------------------
// small conversion helpers
// ---------------------------------------------------------------------------

func lbListToStrings(ctx context.Context, l types.List) []string {
	out := []string{}
	if l.IsNull() || l.IsUnknown() {
		return out
	}
	l.ElementsAs(ctx, &out, false)
	return out
}

func lbSetToStrings(ctx context.Context, s types.Set) []string {
	out := []string{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func lbStringsToList(ctx context.Context, vals []string) types.List {
	v, _ := types.ListValueFrom(ctx, types.StringType, vals)
	return v
}

func lbStringsToSet(ctx context.Context, vals []string) types.Set {
	v, _ := types.SetValueFrom(ctx, types.StringType, vals)
	return v
}

func asString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}

// sslpolicyBindpoint returns the bind point (`type`) for an
// sslvserver_sslpolicy_binding read from NITRO. NITRO OMITS `type` on GET for the
// default REQUEST bind point, so an empty value is normalized to "REQUEST";
// explicit non-default bind points (INTERCEPT_REQ, CLIENTHELLO_REQ) are echoed
// as-is. Without this, a config that sets `type = "REQUEST"` fails the Framework's
// post-apply Set-element consistency check ("Provider produced inconsistent result
// after apply") because the read-back would be "" while the plan holds "REQUEST".
func sslpolicyBindpoint(v interface{}) string {
	if t := asString(v); t != "" {
		return t
	}
	return "REQUEST"
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (r *LbvserverResource) isCluster() bool {
	datalist, err := r.client.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		return false
	}
	return len(datalist) > 0
}

// ---------------------------------------------------------------------------
// ciphersuites (sslvserver_sslciphersuite_binding) — always synced
// ---------------------------------------------------------------------------

func (r *LbvserverResource) syncCiphersuites(ctx context.Context, name string, desired types.List) error {
	want := lbListToStrings(ctx, desired)
	actual, _ := r.client.FindResourceArray(service.Sslvserver_sslciphersuite_binding.Type(), name)
	have := make([]string, 0, len(actual))
	for _, b := range actual {
		have = append(have, asString(b["ciphername"]))
	}
	if desired.IsNull() && len(have) == 0 {
		return nil
	}
	if stringSlicesEqual(want, have) {
		return nil
	}
	for _, cn := range have {
		if err := r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslciphersuite_binding.Type(), name, map[string]string{"ciphername": cn}); err != nil {
			return err
		}
	}
	for _, cn := range want {
		binding := ssl.Sslvserverciphersuitebinding{Vservername: name, Ciphername: cn}
		if _, err := r.client.AddResource(service.Sslvserver_sslciphersuite_binding.Type(), name, binding); err != nil {
			return err
		}
	}
	return nil
}

func (r *LbvserverResource) readCiphersuites(ctx context.Context, name string, data *LbvserverResourceModel) {
	if data.Ciphersuites.IsNull() {
		return
	}
	actual, _ := r.client.FindResourceArray(service.Sslvserver_sslciphersuite_binding.Type(), name)
	vals := make([]string, 0, len(actual))
	for _, b := range actual {
		vals = append(vals, asString(b["ciphername"]))
	}
	data.Ciphersuites = lbStringsToList(ctx, vals)
}

// ---------------------------------------------------------------------------
// ciphers (sslvserver_sslcipher_binding) — cluster deployments only
// ---------------------------------------------------------------------------

func (r *LbvserverResource) syncCiphers(ctx context.Context, name string, desired types.List) error {
	want := lbListToStrings(ctx, desired)
	findParams := service.FindParams{ResourceType: "sslvserver_sslcipher_binding", ResourceName: name, ResourceMissingErrorCode: 258}
	actual, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	have := make([]string, 0, len(actual))
	for _, b := range actual {
		have = append(have, asString(b["cipheraliasname"]))
	}
	if desired.IsNull() && len(have) == 0 {
		return nil
	}
	if stringSlicesEqual(want, have) {
		return nil
	}
	for _, cn := range have {
		if err := r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslcipher_binding.Type(), name, map[string]string{"ciphername": cn}); err != nil {
			return err
		}
	}
	for _, cn := range want {
		binding := ssl.Sslvserverciphersuitebinding{Vservername: name, Ciphername: cn}
		if _, err := r.client.AddResource(service.Sslvserver_sslcipher_binding.Type(), name, binding); err != nil {
			return err
		}
	}
	return nil
}

func (r *LbvserverResource) readCiphers(ctx context.Context, name string, data *LbvserverResourceModel) {
	if data.Ciphers.IsNull() {
		return
	}
	actual, _ := r.client.FindResourceArray(service.Sslvserver_sslcipher_binding.Type(), name)
	vals := make([]string, 0, len(actual))
	for _, b := range actual {
		vals = append(vals, asString(b["cipheraliasname"]))
	}
	data.Ciphers = lbStringsToList(ctx, vals)
}

// ---------------------------------------------------------------------------
// sslcertkey (non-SNI) + snisslcertkeys (SNI) — sslvserver_sslcertkey_binding
// ---------------------------------------------------------------------------

func (r *LbvserverResource) bindSslcertkey(name, cert string) error {
	binding := ssl.Sslvservercertkeybinding{Vservername: name, Certkeyname: cert}
	return r.client.BindResource(service.Sslvserver.Type(), name, service.Sslcertkey.Type(), cert, &binding)
}

func (r *LbvserverResource) unbindSslcertkey(name, cert string) error {
	return r.client.UnbindResource(service.Sslvserver.Type(), name, service.Sslcertkey.Type(), cert, "certkeyname")
}

func (r *LbvserverResource) syncSnisslcert(ctx context.Context, name string, oldSet, newSet types.Set) error {
	oldList := lbSetToStrings(ctx, oldSet)
	newList := lbSetToStrings(ctx, newSet)
	inNew := map[string]bool{}
	for _, c := range newList {
		inNew[c] = true
	}
	inOld := map[string]bool{}
	for _, c := range oldList {
		inOld[c] = true
	}
	// unbind removed
	for _, c := range oldList {
		if !inNew[c] {
			args := map[string]string{"certkeyname": c, "snicert": "true"}
			if err := r.client.DeleteResourceWithArgsMap(service.Sslvserver_sslcertkey_binding.Type(), name, args); err != nil {
				return fmt.Errorf("error unbinding sni sslcertkey %s: %v", c, err)
			}
		}
	}
	// bind added
	for _, c := range newList {
		if !inOld[c] {
			binding := ssl.Sslvservercertkeybinding{Vservername: name, Certkeyname: c, Snicert: true}
			if err := r.client.BindResource(service.Sslvserver.Type(), name, service.Sslcertkey.Type(), c, &binding); err != nil {
				return fmt.Errorf("error binding sni sslcertkey %s: %v", c, err)
			}
		}
	}
	return nil
}

func (r *LbvserverResource) readSslcerts(ctx context.Context, name string, data *LbvserverResourceModel) {
	if data.Sslcertkey.IsNull() && data.Snisslcertkeys.IsNull() {
		return
	}
	bindings, err := r.client.FindAllBoundResources(service.Sslvserver.Type(), name, service.Sslcertkey.Type())
	if err != nil {
		return
	}
	boundCert := ""
	snicerts := make([]string, 0, len(bindings))
	for _, b := range bindings {
		cert, ok := b["certkeyname"]
		snicert, ok2 := b["snicert"]
		if ok && ok2 {
			if snicert == true || snicert == "true" {
				snicerts = append(snicerts, asString(cert))
			} else {
				boundCert = asString(cert)
			}
		}
	}
	if !data.Sslcertkey.IsNull() {
		data.Sslcertkey = types.StringValue(boundCert)
	}
	if !data.Snisslcertkeys.IsNull() {
		data.Snisslcertkeys = lbStringsToSet(ctx, snicerts)
	}
}

// ---------------------------------------------------------------------------
// sslprofile (bound via the sslvserver set endpoint)
// ---------------------------------------------------------------------------

func (r *LbvserverResource) setSslprofile(name, profile string) error {
	sslvserver := ssl.Sslvserver{Vservername: name, Sslprofile: profile}
	_, err := r.client.UpdateResource(service.Sslvserver.Type(), name, &sslvserver)
	return err
}

func (r *LbvserverResource) unsetSslprofile(name string) error {
	sslvserver := ssl.Sslvserver{Vservername: name, Sslprofile: "true"}
	return r.client.ActOnResource(service.Sslvserver.Type(), &sslvserver, "unset")
}

func (r *LbvserverResource) readSslprofile(name string, data *LbvserverResourceModel) {
	dataSsl, err := r.client.FindResource(service.Sslvserver.Type(), name)
	if err != nil || dataSsl == nil {
		// Non-SSL vserver (or no sslvserver view): there is no sslprofile. sslprofile
		// is Optional+Computed, so it must never be left Unknown after apply. Mirror the
		// SDK v2 behaviour, which set "" when the sslvserver GET returned nothing.
		if data.Sslprofile.IsNull() || data.Sslprofile.IsUnknown() {
			data.Sslprofile = types.StringValue("")
		}
		return
	}
	data.Sslprofile = types.StringValue(asString(dataSsl["sslprofile"]))
}

// ---------------------------------------------------------------------------
// sslpolicybinding (sslvserver_sslpolicy_binding)
// ---------------------------------------------------------------------------

func spbKey(policyname string, priority int64) string {
	return fmt.Sprintf("%s|%d", policyname, priority)
}

func (r *LbvserverResource) extractSslpolicyBindings(ctx context.Context, s types.Set) []SslpolicybindingModel {
	out := []SslpolicybindingModel{}
	if s.IsNull() || s.IsUnknown() {
		return out
	}
	s.ElementsAs(ctx, &out, false)
	return out
}

func (r *LbvserverResource) addSingleSslpolicyBinding(name string, b SslpolicybindingModel) error {
	bindingStruct := ssl.Sslvserverpolicybinding{Vservername: name}
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
	if !b.Policyname.IsNull() && !b.Policyname.IsUnknown() {
		bindingStruct.Policyname = b.Policyname.ValueString()
	}
	if !b.Priority.IsNull() && !b.Priority.IsUnknown() {
		bindingStruct.Priority = uint32(b.Priority.ValueInt64())
	}
	if !b.Type.IsNull() && !b.Type.IsUnknown() {
		bindingStruct.Type = b.Type.ValueString()
	}
	_, err := r.client.UpdateResource("sslvserver_sslpolicy_binding", name, bindingStruct)
	return err
}

func (r *LbvserverResource) deleteSingleSslpolicyBinding(name string, b SslpolicybindingModel) error {
	args := make([]string, 0, 3)
	if !b.Policyname.IsNull() && b.Policyname.ValueString() != "" {
		args = append(args, fmt.Sprintf("policyname:%s", b.Policyname.ValueString()))
	}
	if !b.Priority.IsNull() {
		args = append(args, fmt.Sprintf("priority:%d", b.Priority.ValueInt64()))
	}
	if !b.Type.IsNull() && b.Type.ValueString() != "" {
		args = append(args, fmt.Sprintf("type:%s", b.Type.ValueString()))
	}
	return r.client.DeleteResourceWithArgs("sslvserver_sslpolicy_binding", name, args)
}

func (r *LbvserverResource) syncSslpolicyBindings(ctx context.Context, name string, oldSet, newSet types.Set) error {
	oldB := r.extractSslpolicyBindings(ctx, oldSet)
	newB := r.extractSslpolicyBindings(ctx, newSet)
	oldKeys := map[string]bool{}
	for _, b := range oldB {
		oldKeys[spbKey(b.Policyname.ValueString(), b.Priority.ValueInt64())] = true
	}
	newKeys := map[string]bool{}
	for _, b := range newB {
		newKeys[spbKey(b.Policyname.ValueString(), b.Priority.ValueInt64())] = true
	}
	for _, b := range oldB {
		if !newKeys[spbKey(b.Policyname.ValueString(), b.Priority.ValueInt64())] {
			if err := r.deleteSingleSslpolicyBinding(name, b); err != nil {
				return err
			}
		}
	}
	for _, b := range newB {
		if !oldKeys[spbKey(b.Policyname.ValueString(), b.Priority.ValueInt64())] {
			if err := r.addSingleSslpolicyBinding(name, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *LbvserverResource) readSslpolicyBindings(ctx context.Context, name string, data *LbvserverResourceModel) {
	if data.Sslpolicybinding.IsNull() {
		return
	}
	findParams := service.FindParams{ResourceType: "sslvserver_sslpolicy_binding", ResourceName: name, ResourceMissingErrorCode: 258}
	bindings, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		// 1544 is returned on a non-SSL vserver; treat as no bindings.
		if strings.Contains(err.Error(), "\"errorcode\": 1544") {
			empty, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: sslpolicybindingAttrTypes}, []SslpolicybindingModel{})
			data.Sslpolicybinding = empty
			return
		}
		return
	}
	elems := make([]SslpolicybindingModel, 0, len(bindings))
	for _, b := range bindings {
		e := SslpolicybindingModel{
			Gotopriorityexpression: types.StringValue(asString(b["gotopriorityexpression"])),
			Invoke:                 types.BoolValue(b["invoke"] == true || b["invoke"] == "true"),
			Labelname:              types.StringValue(asString(b["labelname"])),
			Labeltype:              types.StringValue(asString(b["labeltype"])),
			Policyname:             types.StringValue(asString(b["policyname"])),
			Priority:               types.Int64Value(0),
			Type:                   types.StringValue(sslpolicyBindpoint(b["type"])),
		}
		if p, ok := b["priority"]; ok && p != nil {
			if iv, err := strconv.ParseInt(asString(p), 10, 64); err == nil {
				e.Priority = types.Int64Value(iv)
			}
		}
		elems = append(elems, e)
	}
	setVal, _ := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: sslpolicybindingAttrTypes}, elems)
	data.Sslpolicybinding = setVal
}

// ---------------------------------------------------------------------------
// combined apply/read orchestration
// ---------------------------------------------------------------------------

// applyBindingsOnCreate binds all convenience blocks after the base lbvserver is added.
// precheckSslcertkeysExist verifies referenced sslcertkey / snisslcertkeys exist
// on the appliance BEFORE the vserver is created, so a missing cert fails cleanly
// up-front instead of orphaning a just-created vserver (SDK v2 parity). Returns
// false (and appends an error) when a referenced cert is missing.
func (r *LbvserverResource) precheckSslcertkeysExist(ctx context.Context, data *LbvserverResourceModel, diags *diag.Diagnostics) bool {
	if !data.Sslcertkey.IsNull() && !data.Sslcertkey.IsUnknown() && data.Sslcertkey.ValueString() != "" {
		if !r.client.ResourceExists(service.Sslcertkey.Type(), data.Sslcertkey.ValueString()) {
			diags.AddError("Configuration Error", fmt.Sprintf("Specified sslcertkey %q does not exist on the NetScaler.", data.Sslcertkey.ValueString()))
			return false
		}
	}
	if !data.Snisslcertkeys.IsNull() && !data.Snisslcertkeys.IsUnknown() {
		var sniCerts []string
		diags.Append(data.Snisslcertkeys.ElementsAs(ctx, &sniCerts, false)...)
		if diags.HasError() {
			return false
		}
		var missing []string
		for _, c := range sniCerts {
			if c != "" && !r.client.ResourceExists(service.Sslcertkey.Type(), c) {
				missing = append(missing, c)
			}
		}
		if len(missing) > 0 {
			diags.AddError("Configuration Error", fmt.Sprintf("The following SNI sslcertkey(s) do not exist on the NetScaler: %v", missing))
			return false
		}
	}
	return true
}

func (r *LbvserverResource) applyBindingsOnCreate(ctx context.Context, name string, data *LbvserverResourceModel, diags *diag.Diagnostics) {
	if !data.Sslcertkey.IsNull() && data.Sslcertkey.ValueString() != "" {
		if err := r.bindSslcertkey(name, data.Sslcertkey.ValueString()); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind sslcertkey to lbvserver: %s", err))
			return
		}
	}
	if !data.Snisslcertkeys.IsNull() {
		if err := r.syncSnisslcert(ctx, name, types.SetNull(types.StringType), data.Snisslcertkeys); err != nil {
			diags.AddError("Client Error", err.Error())
			return
		}
	}
	if r.isCluster() {
		if err := r.syncCiphers(ctx, name, data.Ciphers); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync ciphers: %s", err))
			return
		}
	}
	if err := r.syncCiphersuites(ctx, name, data.Ciphersuites); err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to sync ciphersuites: %s", err))
		return
	}
	if !data.Sslprofile.IsNull() && data.Sslprofile.ValueString() != "" {
		if err := r.setSslprofile(name, data.Sslprofile.ValueString()); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind sslprofile to lbvserver: %s", err))
			return
		}
	}
	if !data.Sslpolicybinding.IsNull() {
		if err := r.syncSslpolicyBindings(ctx, name, types.SetNull(types.ObjectType{AttrTypes: sslpolicybindingAttrTypes}), data.Sslpolicybinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to bind sslpolicy to lbvserver: %s", err))
			return
		}
	}
}

// applyBindingsOnUpdate reconciles all convenience blocks against prior state.
func (r *LbvserverResource) applyBindingsOnUpdate(ctx context.Context, name string, data, state *LbvserverResourceModel, diags *diag.Diagnostics) {
	if !data.Sslcertkey.Equal(state.Sslcertkey) {
		if !state.Sslcertkey.IsNull() && state.Sslcertkey.ValueString() != "" {
			if err := r.unbindSslcertkey(name, state.Sslcertkey.ValueString()); err != nil {
				diags.AddError("Client Error", fmt.Sprintf("Unable to unbind old sslcertkey: %s", err))
				return
			}
		}
		if !data.Sslcertkey.IsNull() && data.Sslcertkey.ValueString() != "" {
			if err := r.bindSslcertkey(name, data.Sslcertkey.ValueString()); err != nil {
				diags.AddError("Client Error", fmt.Sprintf("Unable to bind new sslcertkey: %s", err))
				return
			}
		}
	}
	// sslprofile is Optional+Computed, so the framework marks it unknown whenever any
	// sibling attribute changes. Unknown means "keep the computed value" - it must NOT
	// be treated as an empty value requesting an unset (which fails with errorcode
	// 3679 on a default-SSL-profile ADC). Only reconcile a concrete, user-driven change.
	if !data.Sslprofile.IsUnknown() && !data.Sslprofile.Equal(state.Sslprofile) {
		if data.Sslprofile.IsNull() || data.Sslprofile.ValueString() == "" {
			if !state.Sslprofile.IsNull() && state.Sslprofile.ValueString() != "" {
				if err := r.unsetSslprofile(name); err != nil {
					diags.AddError("Client Error", fmt.Sprintf("Unable to unset sslprofile: %s", err))
					return
				}
			}
		} else {
			if err := r.setSslprofile(name, data.Sslprofile.ValueString()); err != nil {
				diags.AddError("Client Error", fmt.Sprintf("Unable to set sslprofile: %s", err))
				return
			}
		}
	}
	if !data.Snisslcertkeys.Equal(state.Snisslcertkeys) {
		if err := r.syncSnisslcert(ctx, name, state.Snisslcertkeys, data.Snisslcertkeys); err != nil {
			diags.AddError("Client Error", err.Error())
			return
		}
	}
	if r.isCluster() && !data.Ciphers.Equal(state.Ciphers) {
		if err := r.syncCiphers(ctx, name, data.Ciphers); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync ciphers: %s", err))
			return
		}
	}
	if !data.Ciphersuites.Equal(state.Ciphersuites) {
		if err := r.syncCiphersuites(ctx, name, data.Ciphersuites); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync ciphersuites: %s", err))
			return
		}
	}
	if !data.Sslpolicybinding.Equal(state.Sslpolicybinding) {
		if err := r.syncSslpolicyBindings(ctx, name, state.Sslpolicybinding, data.Sslpolicybinding); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to sync sslpolicy bindings: %s", err))
			return
		}
	}
}

// readBindings refreshes all managed convenience blocks from the appliance.
func (r *LbvserverResource) readBindings(ctx context.Context, name string, data *LbvserverResourceModel) {
	r.readSslcerts(ctx, name, data)
	r.readSslpolicyBindings(ctx, name, data)
	r.readSslprofile(name, data)
	r.readCiphersuites(ctx, name, data)
	if r.isCluster() {
		r.readCiphers(ctx, name, data)
	}
}
