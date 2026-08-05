package nsacls

import (
	"context"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/service"
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
//   - Read is a state-preserving no-op: nsacls has no aggregate GET; SDK v2 faked
//     it by scanning ALL device nsacls (leaking foreign rules), which is unsafe in
//     the framework. This matches the sibling rnat_clear migration.
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

	// nsacls has no aggregate GET. SDK v2's Read scanned ALL device nsacls and
	// blindly wrote them into state (leaking foreign rules and risking framework
	// "inconsistent result" errors). We preserve prior state instead.
	tflog.Debug(ctx, "Read is a state-preserving no-op for nsacls")

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
