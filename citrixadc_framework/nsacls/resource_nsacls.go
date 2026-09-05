package nsacls

import (
	"context"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &NsaclsResource{}
var _ resource.ResourceWithConfigure = (*NsaclsResource)(nil)
var _ resource.ResourceWithImportState = (*NsaclsResource)(nil)

func NewNsaclsResource() resource.Resource {
	return &NsaclsResource{}
}

// NsaclsResource defines the resource implementation.
//
// This mirrors the SDK v2 `citrixadc_nsacls` convenience resource. It is NOT a
// plain singleton: it manages a *set* of individual `nsacl` rules identified by a
// synthetic handle (aclsname):
//   - Create adds every rule in the `acl` set with AddResource(nsacl) and then
//     applies the ACL set with ApplyResource(nsacls) (POST ?action=apply).
//   - Update diffs the old and new sets by aclname: removed rules are deleted,
//     added rules are added, and changed rules are re-added (exactly as SDK v2's
//     createSingleAcl-for-update behavior), then ApplyResource(nsacls) is called.
//   - Delete deletes every rule in state with DeleteResource(nsacl) and applies.
//   - Read refreshes ONLY the rules this resource manages (those already in
//     state), one per-rule GET keyed by aclname. nsacls has no aggregate GET, and
//     SDK v2 faked it with an unscoped scan of ALL device nsacls (leaking foreign
//     rules). The state-scoped per-rule refresh restores drift detection without
//     that leak; see the Read method for the echo-only / omit-on-default rules.
type NsaclsResource struct {
	client *service.NitroClient
}

func (r *NsaclsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsaclsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsacls"
}

func (r *NsaclsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsaclsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsaclsResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsacls resource")

	// Resolve the aclsname handle: use the configured value if present,
	// otherwise generate a unique "tf-nsacl-*" name (mirrors SDK v2
	// resource.PrefixedUniqueId("tf-nsacl-")).
	var aclsName string
	if !data.Aclsname.IsNull() && !data.Aclsname.IsUnknown() && data.Aclsname.ValueString() != "" {
		aclsName = data.Aclsname.ValueString()
	} else {
		aclsName = fmt.Sprintf("tf-nsacl-%d", time.Now().UnixNano())
	}
	data.Aclsname = types.StringValue(aclsName)

	// Add every rule in the set. SDK v2 ignored per-rule errors here.
	elems := nsaclsElementsFromSet(ctx, data.Acl, &resp.Diagnostics)
	for i := range elems {
		if err := r.createSingleAcl(ctx, &elems[i]); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("error adding nsacl rule: %s", err))
		}
	}

	// Apply the ACL set. SDK v2 passed the configured type here.
	if err := r.applyNsacls(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nsacls, got error: %s", err))
		return
	}

	// Resolve computed attributes to known values.
	r.resolveComputed(&data)

	// ID scheme matches SDK v2: the aclsname handle.
	data.Id = types.StringValue(aclsName)

	tflog.Trace(ctx, "Created nsacls resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaclsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsaclsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsacls resource")

	// Refresh only the rules this resource manages (those already in prior state),
	// keyed by aclname, via a per-rule GET. nsacls has no aggregate GET, and SDK
	// v2's unscoped FindAllResources(nsacl) scan leaked foreign rules into state
	// (perpetual plans / "inconsistent result"). Scoping the refresh to the rules
	// in state avoids that leak entirely: a managed rule deleted on the appliance
	// is dropped (recreated on the next apply -> self-healing), and a managed rule
	// whose configured fields changed out-of-band is refreshed so the drift shows
	// in the plan.
	//
	// The refresh is ECHO-ONLY / OMIT-ON-DEFAULT: only attributes the user actually
	// set (non-null in prior state) are overwritten from the device. Unset (null)
	// attributes are never populated with NITRO's per-rule defaults - doing so would
	// flip the set element's identity hash and reintroduce the very
	// "inconsistent result after apply" / perpetual-plan failure the previous no-op
	// read was avoiding (the acl nested attributes are Optional, not Computed).
	r.refreshManagedAcls(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// The effective ADC default for `type` is CLASSIC. The released SDK v2 provider
	// stored `type` as null in state when it was left unset. Carrying that null
	// forward (UseStateForUnknown copies the refreshed state into the plan) makes
	// the planned `type` null, while Create/Update resolve it to "CLASSIC" -> the
	// post-apply value differs from the plan ("inconsistent result after apply").
	// Normalize the refreshed state to the effective default so the refreshed
	// state (and therefore the plan) already holds "CLASSIC", matching apply.
	// `type` is not configured in this case, so RequiresReplaceIfConfigured does
	// not fire and no spurious replace is produced.
	if data.Type.IsNull() || data.Type.IsUnknown() || data.Type.ValueString() == "" {
		data.Type = types.StringValue("CLASSIC")
	}

	// Reproduce the SDK v2 toggle: acls_apply_trigger is reset to "No" after every
	// read, so a config value of "Yes" always produces a diff and re-applies.
	data.AclsApplyTrigger = types.StringValue("No")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaclsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsaclsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating nsacls resource")

	oldElems := nsaclsElementsFromSet(ctx, state.Acl, &resp.Diagnostics)
	newElems := nsaclsElementsFromSet(ctx, data.Acl, &resp.Diagnostics)

	oldByName := make(map[string]NsaclEntryModel, len(oldElems))
	for i := range oldElems {
		oldByName[oldElems[i].Aclname.ValueString()] = oldElems[i]
	}
	newByName := make(map[string]NsaclEntryModel, len(newElems))
	for i := range newElems {
		newByName[newElems[i].Aclname.ValueString()] = newElems[i]
	}

	// Rules present in the old set but not the new set are deleted.
	for name := range oldByName {
		if _, ok := newByName[name]; !ok {
			e := oldByName[name]
			if err := r.deleteSingleAcl(ctx, &e); err != nil {
				tflog.Debug(ctx, fmt.Sprintf("error deleting nsacl rule %s: %s", name, err))
			}
		}
	}

	// Rules present in the new set are added (new name) or re-added (changed).
	// This mirrors SDK v2, which calls createSingleAcl for both add and update.
	for i := range newElems {
		name := newElems[i].Aclname.ValueString()
		if oe, ok := oldByName[name]; ok {
			if oe == newElems[i] {
				continue // unchanged
			}
		}
		if err := r.createSingleAcl(ctx, &newElems[i]); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("error applying nsacl rule %s: %s", name, err))
		}
	}

	// Apply the ACL set (SDK v2 update applied with an empty nsacls payload).
	if err := r.applyNsacls(ctx, &data); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nsacls, got error: %s", err))
		return
	}

	// Preserve the ID/handle from prior state.
	data.Id = state.Id
	if data.Aclsname.IsNull() || data.Aclsname.IsUnknown() {
		data.Aclsname = state.Aclsname
	}
	r.resolveComputed(&data)

	tflog.Trace(ctx, "Updated nsacls resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaclsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsaclsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsacls resource")

	// Delete every managed rule. SDK v2 ignored per-rule errors here.
	elems := nsaclsElementsFromSet(ctx, data.Acl, &resp.Diagnostics)
	for i := range elems {
		if err := r.deleteSingleAcl(ctx, &elems[i]); err != nil {
			tflog.Debug(ctx, fmt.Sprintf("error deleting nsacl rule: %s", err))
		}
	}

	// Apply the (now empty) ACL set, matching SDK v2 delete.
	empty := ns.Nsacls{}
	if err := r.client.ApplyResource(service.Nsacls.Type(), &empty); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to apply nsacls on delete, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsacls resource")
}

// createSingleAcl builds and adds a single nsacl rule (AddResource). SDK v2 uses
// AddResource for both create and update of individual rules.
func (r *NsaclsResource) createSingleAcl(ctx context.Context, m *NsaclEntryModel) error {
	nsacl, err := nsaclsBuildNsaclPayload(ctx, m)
	if err != nil {
		return err
	}
	_, err = r.client.AddResource(service.Nsacl.Type(), nsacl.Aclname, &nsacl)
	return err
}

// deleteSingleAcl deletes a single nsacl rule by name (DeleteResource).
func (r *NsaclsResource) deleteSingleAcl(ctx context.Context, m *NsaclEntryModel) error {
	return r.client.DeleteResource(service.Nsacl.Type(), m.Aclname.ValueString())
}

// applyNsacls issues the POST ?action=apply on the nsacls object, passing the
// configured type (omitempty means an unset type applies as the CLASSIC default,
// matching SDK v2).
func (r *NsaclsResource) applyNsacls(ctx context.Context, data *NsaclsResourceModel) error {
	nsacls := ns.Nsacls{}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		nsacls.Type = data.Type.ValueString()
	}
	return r.client.ApplyResource(service.Nsacls.Type(), &nsacls)
}

// resolveComputed fills computed attributes with known values so state is
// consistent with the plan after Create/Update.
func (r *NsaclsResource) resolveComputed(data *NsaclsResourceModel) {
	// type: the effective ADC default is CLASSIC when unset.
	if data.Type.IsNull() || data.Type.IsUnknown() {
		data.Type = types.StringValue("CLASSIC")
	}
	// acls_apply_trigger: default toggle position is "No".
	if data.AclsApplyTrigger.IsNull() || data.AclsApplyTrigger.IsUnknown() {
		data.AclsApplyTrigger = types.StringValue("No")
	}
	// acl: resolve an omitted (unknown) set to a known null value.
	if data.Acl.IsUnknown() {
		data.Acl = types.SetNull(nsaclObjectType())
	}
}

// nsaclsElementsFromSet decodes the `acl` set into a slice, tolerating
// null/unknown sets (returns an empty slice).
func nsaclsElementsFromSet(ctx context.Context, set types.Set, diags *diag.Diagnostics) []NsaclEntryModel {
	var elems []NsaclEntryModel
	if set.IsNull() || set.IsUnknown() {
		return elems
	}
	diags.Append(set.ElementsAs(ctx, &elems, false)...)
	return elems
}

// refreshManagedAcls re-reads each rule in prior state (keyed by aclname) from the
// appliance and rebuilds data.Acl. Rules absent on the appliance are dropped
// (out-of-band deletion -> recreated on next apply). This is scoped strictly to
// the rules already in state, so unmanaged/foreign rules are never pulled in.
func (r *NsaclsResource) refreshManagedAcls(ctx context.Context, data *NsaclsResourceModel, diags *diag.Diagnostics) {
	if data.Acl.IsNull() || data.Acl.IsUnknown() {
		return
	}
	elems := nsaclsElementsFromSet(ctx, data.Acl, diags)
	if diags.HasError() {
		return
	}

	refreshed := make([]NsaclEntryModel, 0, len(elems))
	for i := range elems {
		name := elems[i].Aclname.ValueString()
		if name == "" {
			// aclname is Required; keep any degenerate element as-is.
			refreshed = append(refreshed, elems[i])
			continue
		}
		dev, err := r.client.FindResource(service.Nsacl.Type(), name)
		if err != nil {
			if utils.IsNotFoundError(err) {
				// Deleted out-of-band: drop it so the next apply recreates it.
				tflog.Debug(ctx, fmt.Sprintf("nsacl rule %q no longer exists on the appliance; dropping from state", name))
				continue
			}
			diags.AddError("Client Error", fmt.Sprintf("Unable to read nsacl rule %q, got error: %s", name, err))
			return
		}
		refreshed = append(refreshed, refreshNsaclEntry(elems[i], dev))
	}

	setVal, d := types.SetValueFrom(ctx, nsaclObjectType(), refreshed)
	diags.Append(d...)
	if diags.HasError() {
		return
	}
	data.Acl = setVal
}

// refreshNsaclEntry returns a copy of the prior-state rule with only the
// user-configured (non-null) attributes overwritten from the device GET map.
// aclname (the match key) is preserved as-is. Attributes that were null in state
// are left null (omit-on-default) so device defaults never flip the set-element
// identity.
func refreshNsaclEntry(state NsaclEntryModel, dev map[string]interface{}) NsaclEntryModel {
	out := state
	out.Aclaction = echoStr(state.Aclaction, dev, "aclaction")
	out.Destipop = echoStr(state.Destipop, dev, "destipop")
	out.Destipval = echoStr(state.Destipval, dev, "destipval")
	out.Destportop = echoStr(state.Destportop, dev, "destportop")
	out.Destportval = echoStr(state.Destportval, dev, "destportval")
	out.Interface = echoStr(state.Interface, dev, "interface")
	out.Logstate = echoStr(state.Logstate, dev, "logstate")
	out.Protocol = echoStr(state.Protocol, dev, "protocol")
	out.Srcipop = echoStr(state.Srcipop, dev, "srcipop")
	out.Srcipval = echoStr(state.Srcipval, dev, "srcipval")
	out.Srcmac = echoStr(state.Srcmac, dev, "srcmac")
	out.Srcportop = echoStr(state.Srcportop, dev, "srcportop")
	out.Srcportval = echoStr(state.Srcportval, dev, "srcportval")
	out.State = echoStr(state.State, dev, "state")
	out.Srcportdataset = echoStr(state.Srcportdataset, dev, "srcportdataset")
	out.Srcipdataset = echoStr(state.Srcipdataset, dev, "srcipdataset")
	out.Destportdataset = echoStr(state.Destportdataset, dev, "destportdataset")
	out.Destipdataset = echoStr(state.Destipdataset, dev, "destipdataset")

	out.Established = echoBool(state.Established, dev, "established")

	out.Icmpcode = echoInt(state.Icmpcode, dev, "icmpcode")
	out.Icmptype = echoInt(state.Icmptype, dev, "icmptype")
	out.Priority = echoInt(state.Priority, dev, "priority")
	out.Protocolnumber = echoInt(state.Protocolnumber, dev, "protocolnumber")
	out.Ratelimit = echoInt(state.Ratelimit, dev, "ratelimit")
	out.Td = echoInt(state.Td, dev, "td")
	out.Ttl = echoInt(state.Ttl, dev, "ttl")
	out.Vlan = echoInt(state.Vlan, dev, "vlan")

	return out
}

// echoStr overwrites a string attribute from the device only when it was set in
// prior state (omit-on-default) and the device echoes the key.
func echoStr(cur types.String, dev map[string]interface{}, key string) types.String {
	if cur.IsNull() || cur.IsUnknown() {
		return cur
	}
	if v, ok := dev[key]; ok && v != nil {
		return types.StringValue(nsaclDevString(v))
	}
	return cur
}

// echoInt overwrites an int attribute from the device only when it was set in
// prior state and the device echoes a parseable value.
func echoInt(cur types.Int64, dev map[string]interface{}, key string) types.Int64 {
	if cur.IsNull() || cur.IsUnknown() {
		return cur
	}
	if v, ok := dev[key]; ok && v != nil {
		if iv, err := utils.ConvertToInt64(v); err == nil {
			return types.Int64Value(iv)
		}
	}
	return cur
}

// echoBool overwrites a bool attribute from the device only when it was set in
// prior state and the device echoes the key.
func echoBool(cur types.Bool, dev map[string]interface{}, key string) types.Bool {
	if cur.IsNull() || cur.IsUnknown() {
		return cur
	}
	if v, ok := dev[key]; ok && v != nil {
		return types.BoolValue(v == true || v == "true")
	}
	return cur
}

// nsaclDevString coerces a NITRO GET value into a string.
func nsaclDevString(v interface{}) string {
	if s, ok := v.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", v)
}
