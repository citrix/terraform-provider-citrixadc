package gslbvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbvserverResource{}
var _ resource.ResourceWithConfigure = (*GslbvserverResource)(nil)
var _ resource.ResourceWithImportState = (*GslbvserverResource)(nil)

func NewGslbvserverResource() resource.Resource {
	return &GslbvserverResource{}
}

// GslbvserverResource defines the resource implementation.
type GslbvserverResource struct {
	client *service.NitroClient
}

func (r *GslbvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbvserver"
}

func (r *GslbvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbvserver resource")

	gslbvserver := gslbvserverGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource keyed on the primary attribute (name).
	name_value := data.Name.ValueString()
	_, err := r.client.AddResource(service.Gslbvserver.Type(), name_value, &gslbvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbvserver, got error: %s", err))
		return
	}

	// backupvserver is not accepted by the NITRO `add` verb; apply it via a follow-up
	// update once the vserver exists, mirroring the SDK v2 behavior.
	if !data.Backupvserver.IsNull() && !data.Backupvserver.IsUnknown() && data.Backupvserver.ValueString() != "" {
		backupPayload := gslb.Gslbvserver{
			Name:          name_value,
			Backupvserver: data.Backupvserver.ValueString(),
		}
		if _, err := r.client.UpdateResource(service.Gslbvserver.Type(), name_value, &backupPayload); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set backupvserver on gslbvserver, got error: %s", err))
			return
		}
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	// Bind the "domain" and "service" convenience blocks (SDK v2 parity). The old
	// sets are empty on create.
	nullDomainSet := types.SetNull(types.ObjectType{AttrTypes: domainbindingAttrTypes})
	nullServiceSet := types.SetNull(types.ObjectType{AttrTypes: servicebindingAttrTypes})
	resp.Diagnostics.Append(r.syncDomainBindings(ctx, name_value, nullDomainSet, data.Domain)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(r.syncServiceBindings(ctx, name_value, nullServiceSet, data.Service)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Created gslbvserver resource")

	// Read the updated state back
	r.readGslbvserverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbvserver resource")

	r.readGslbvserverFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// If the resource no longer exists on the ADC, remove it from state.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state GslbvserverResourceModel

	// Read Terraform prior state to preserve ID and to compute changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (tracks the current live name)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating gslbvserver resource")

	// Rename support: gslbvserver exposes a NITRO `rename` action. `name` and
	// `servicetype` are RequiresReplace, so Terraform recreates the resource for
	// those; the ONLY key change that lands in Update is `newname`. On a newname
	// change, POST {name, newname} to ?action=rename and repoint the ID at the new
	// name so subsequent reads address the live object. The rename SOURCE is the
	// current live name (state.Id), NOT state.Name (which stays pinned to the
	// originally configured value and would be stale after a second rename).
	if !data.Newname.Equal(state.Newname) && !data.Newname.IsNull() && data.Newname.ValueString() != "" {
		oldName := state.Id.ValueString()
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming gslbvserver from %q to %q", oldName, newName))

		renamePayload := gslb.Gslbvserver{
			Name:    oldName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Gslbvserver.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename gslbvserver, got error: %s", err))
			return
		}
		data.Id = types.StringValue(newName)
	}

	// The current live name (after any rename above).
	liveName := data.Id.ValueString()

	// Apply changes to the updatable attributes (excludes name/servicetype/state/newname).
	gslbvserver, hasChange := gslbvserverGetTheUpdatablePayloadFromThePlan(ctx, &data, &state)
	if hasChange {
		gslbvserver.Name = liveName
		if _, err := r.client.UpdateResource(service.Gslbvserver.Type(), liveName, &gslbvserver); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbvserver, got error: %s", err))
			return
		}
	} else {
		tflog.Debug(ctx, "No updatable field changes detected for gslbvserver resource")
	}

	// Collect attributes that were removed from config so the appliance reverts them
	// to their NITRO defaults via the unset action. Each of these has a schema Default
	// so config removal produces a plan diff (data != state) that lands in Update.
	attributesToUnset := []string{}
	if !data.Appflowlog.Equal(state.Appflowlog) && config.Appflowlog.IsNull() {
		attributesToUnset = append(attributesToUnset, "appflowlog")
	}
	if !data.Considereffectivestate.Equal(state.Considereffectivestate) && config.Considereffectivestate.IsNull() {
		attributesToUnset = append(attributesToUnset, "considereffectivestate")
	}
	if !data.Disableprimaryondown.Equal(state.Disableprimaryondown) && config.Disableprimaryondown.IsNull() {
		attributesToUnset = append(attributesToUnset, "disableprimaryondown")
	}
	if !data.Dnsrecordtype.Equal(state.Dnsrecordtype) && config.Dnsrecordtype.IsNull() {
		attributesToUnset = append(attributesToUnset, "dnsrecordtype")
	}
	if !data.Ecs.Equal(state.Ecs) && config.Ecs.IsNull() {
		attributesToUnset = append(attributesToUnset, "ecs")
	}
	if !data.Ecsaddrvalidation.Equal(state.Ecsaddrvalidation) && config.Ecsaddrvalidation.IsNull() {
		attributesToUnset = append(attributesToUnset, "ecsaddrvalidation")
	}
	if !data.Edr.Equal(state.Edr) && config.Edr.IsNull() {
		attributesToUnset = append(attributesToUnset, "edr")
	}
	if !data.Lbmethod.Equal(state.Lbmethod) && config.Lbmethod.IsNull() {
		attributesToUnset = append(attributesToUnset, "lbmethod")
	}
	if !data.Toggleorder.Equal(state.Toggleorder) && config.Toggleorder.IsNull() {
		attributesToUnset = append(attributesToUnset, "toggleorder")
	}

	// Unset attributes removed from config so the appliance reverts them to defaults.
	// Done after the update so any default value the update payload carried for a
	// removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": liveName,
	}
	if err := utils.ExecuteUnset(r.client, service.Gslbvserver.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset gslbvserver attributes, got error: %s", err))
		return
	}

	// State (ENABLED/DISABLED) is not part of the update payload; it is driven by the
	// enable/disable NITRO actions, mirroring the SDK v2 behavior.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doGslbvserverStateChange(liveName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change state of gslbvserver, got error: %s", err))
			return
		}
	}

	// Reconcile the "domain" and "service" convenience blocks (SDK v2 parity).
	resp.Diagnostics.Append(r.syncDomainBindings(ctx, liveName, state.Domain, data.Domain)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(r.syncServiceBindings(ctx, liveName, state.Service, data.Service)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Updated gslbvserver resource")

	// Read the updated state back. The read must not clobber the user-facing name /
	// newname inputs, so capture and restore them around the read.
	planName := data.Name
	planNewname := data.Newname
	r.readGslbvserverFromApi(ctx, &data, &resp.Diagnostics)
	data.Name = planName
	data.Newname = planNewname
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbvserver resource")

	// Named resource - delete using DeleteResource. The ID holds the CURRENT LIVE
	// name (== name at create, == newname after a rename), so delete by data.Id, NOT
	// data.Name (which stays at the originally configured value and would target a
	// non-existent name after a rename, leaving the object dangling).
	liveName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Gslbvserver.Type(), liveName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbvserver resource")
}

// doGslbvserverStateChange enables or disables the GSLB virtual server via the NITRO
// enable/disable actions (state is not settable through the update payload).
func (r *GslbvserverResource) doGslbvserverStateChange(name string, newstate string) error {
	// A minimal struct is required - ActOnResource fails on superfluous attributes.
	gslbvserver := gslb.Gslbvserver{
		Name: name,
	}
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Gslbvserver.Type(), &gslbvserver, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Gslbvserver.Type(), &gslbvserver, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read gslbvserver data from API
func (r *GslbvserverResource) readGslbvserverFromApi(ctx context.Context, data *GslbvserverResourceModel, diags *diag.Diagnostics) {
	// Case 2: Find with single ID attribute - ID is the plain value (the live name).
	name_Name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Gslbvserver.Type(), name_Name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbvserver, got error: %s", err))
		return
	}

	gslbvserverSetAttrFromGet(ctx, data, getResponseData)

	// Refresh the convenience binding blocks only when the resource manages them
	// (state/plan non-null). This keeps resources that never use these blocks free of
	// spurious set values.
	if !data.Domain.IsNull() {
		diags.Append(r.readDomainBindings(ctx, name_Name, data)...)
	}
	if !data.Service.IsNull() {
		diags.Append(r.readServiceBindings(ctx, name_Name, data)...)
	}
}

// syncDomainBindings adds/removes gslbvserver_domain_binding entries to match newSet,
// mirroring the SDK v2 "domain" block reconciliation.
func (r *GslbvserverResource) syncDomainBindings(ctx context.Context, vserverName string, oldSet, newSet types.Set) diag.Diagnostics {
	var diags diag.Diagnostics

	// If the config does not declare the "domain" convenience block, do not
	// reconcile. A legacy SDK v2 state upgrade populates the parent's domain
	// block from the ADC, so reconciling against a null config would delete
	// domain bindings that are owned by a separate gslbvserver_domain_binding
	// resource. A non-null (even empty) set still reconciles as an explicit set.
	if newSet.IsNull() {
		return diags
	}

	var oldBindings, newBindings []DomainbindingModel
	if !oldSet.IsNull() && !oldSet.IsUnknown() {
		diags.Append(oldSet.ElementsAs(ctx, &oldBindings, false)...)
	}
	if !newSet.IsNull() && !newSet.IsUnknown() {
		diags.Append(newSet.ElementsAs(ctx, &newBindings, false)...)
	}
	if diags.HasError() {
		return diags
	}

	newNames := make(map[string]bool)
	for _, b := range newBindings {
		newNames[b.Domainname.ValueString()] = true
	}
	oldNames := make(map[string]bool)
	for _, b := range oldBindings {
		oldNames[b.Domainname.ValueString()] = true
	}

	// Remove domains no longer present.
	for _, b := range oldBindings {
		dn := b.Domainname.ValueString()
		if dn == "" || newNames[dn] {
			continue
		}
		args := []string{fmt.Sprintf("domainname:%s", dn)}
		if err := r.client.DeleteResourceWithArgs(service.Gslbvserver_domain_binding.Type(), vserverName, args); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to delete domain binding %s from gslbvserver %s: %s", dn, vserverName, err))
			return diags
		}
	}

	// Add new domains.
	for _, b := range newBindings {
		dn := b.Domainname.ValueString()
		if dn == "" || oldNames[dn] {
			continue
		}
		bind := gslb.Gslbvserverdomainbinding{Name: vserverName, Domainname: dn}
		if !b.Backupip.IsNull() && !b.Backupip.IsUnknown() {
			bind.Backupip = b.Backupip.ValueString()
		}
		if !b.Cookiedomain.IsNull() && !b.Cookiedomain.IsUnknown() {
			bind.Cookiedomain = b.Cookiedomain.ValueString()
		}
		if !b.Cookietimeout.IsNull() && !b.Cookietimeout.IsUnknown() {
			bind.Cookietimeout = utils.IntPtr(int(b.Cookietimeout.ValueInt64()))
		}
		if !b.Sitedomainttl.IsNull() && !b.Sitedomainttl.IsUnknown() {
			bind.Sitedomainttl = utils.IntPtr(int(b.Sitedomainttl.ValueInt64()))
		}
		if !b.Ttl.IsNull() && !b.Ttl.IsUnknown() {
			bind.Ttl = utils.IntPtr(int(b.Ttl.ValueInt64()))
		}
		if !b.Backupipflag.IsNull() && !b.Backupipflag.IsUnknown() {
			bind.Backupipflag = b.Backupipflag.ValueBool()
		}
		if !b.Cookiedomainflag.IsNull() && !b.Cookiedomainflag.IsUnknown() {
			bind.Cookiedomainflag = b.Cookiedomainflag.ValueBool()
		}
		if _, err := r.client.UpdateResource(service.Gslbvserver_domain_binding.Type(), vserverName, &bind); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to add domain binding %s to gslbvserver %s: %s", dn, vserverName, err))
			return diags
		}
	}

	return diags
}

// readDomainBindings populates data.Domain from the ADC.
func (r *GslbvserverResource) readDomainBindings(ctx context.Context, vserverName string, data *GslbvserverResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	bindings, _ := r.client.FindResourceArray(service.Gslbvserver_domain_binding.Type(), vserverName)
	elems := make([]DomainbindingModel, 0, len(bindings))
	for _, m := range bindings {
		e := DomainbindingModel{
			Backupip:         types.StringNull(),
			Backupipflag:     types.BoolNull(),
			Cookiedomain:     types.StringNull(),
			Cookiedomainflag: types.BoolNull(),
			Cookietimeout:    types.Int64Null(),
			Domainname:       types.StringNull(),
			Name:             types.StringNull(),
			Sitedomainttl:    types.Int64Null(),
			Ttl:              types.Int64Null(),
		}
		if v, ok := m["backupip"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Backupip = types.StringValue(s)
			}
		}
		if v, ok := m["backupipflag"]; ok && v != nil {
			if bv, isBool := v.(bool); isBool {
				e.Backupipflag = types.BoolValue(bv)
			}
		}
		if v, ok := m["cookie_domain"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Cookiedomain = types.StringValue(s)
			}
		}
		if v, ok := m["cookie_domainflag"]; ok && v != nil {
			if bv, isBool := v.(bool); isBool {
				e.Cookiedomainflag = types.BoolValue(bv)
			}
		}
		if v, ok := m["cookietimeout"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Cookietimeout = types.Int64Value(iv)
			}
		}
		if v, ok := m["domainname"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Domainname = types.StringValue(s)
			}
		}
		if v, ok := m["name"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Name = types.StringValue(s)
			}
		}
		if v, ok := m["sitedomainttl"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Sitedomainttl = types.Int64Value(iv)
			}
		}
		if v, ok := m["ttl"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Ttl = types.Int64Value(iv)
			}
		}
		elems = append(elems, e)
	}

	setVal, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: domainbindingAttrTypes}, elems)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Domain = setVal
	return diags
}

// syncServiceBindings adds/removes gslbvserver_gslbservice_binding entries to match
// newSet, mirroring the SDK v2 "service" block reconciliation.
func (r *GslbvserverResource) syncServiceBindings(ctx context.Context, vserverName string, oldSet, newSet types.Set) diag.Diagnostics {
	var diags diag.Diagnostics

	// If the config does not declare the "service" convenience block, do not
	// reconcile. A legacy SDK v2 state upgrade populates the parent's service
	// block from the ADC, so reconciling against a null config would delete
	// service bindings that are owned by a separate gslbvserver_gslbservice_binding
	// resource. A non-null (even empty) set still reconciles as an explicit set.
	if newSet.IsNull() {
		return diags
	}

	var oldBindings, newBindings []ServicebindingModel
	if !oldSet.IsNull() && !oldSet.IsUnknown() {
		diags.Append(oldSet.ElementsAs(ctx, &oldBindings, false)...)
	}
	if !newSet.IsNull() && !newSet.IsUnknown() {
		diags.Append(newSet.ElementsAs(ctx, &newBindings, false)...)
	}
	if diags.HasError() {
		return diags
	}

	newNames := make(map[string]bool)
	for _, b := range newBindings {
		newNames[b.Servicename.ValueString()] = true
	}
	oldNames := make(map[string]bool)
	for _, b := range oldBindings {
		oldNames[b.Servicename.ValueString()] = true
	}

	// Remove services no longer present.
	for _, b := range oldBindings {
		sn := b.Servicename.ValueString()
		if sn == "" || newNames[sn] {
			continue
		}
		args := []string{fmt.Sprintf("servicename:%s", sn)}
		if err := r.client.DeleteResourceWithArgs(service.Gslbvserver_gslbservice_binding.Type(), vserverName, args); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to delete service binding %s from gslbvserver %s: %s", sn, vserverName, err))
			return diags
		}
	}

	// Add new services.
	for _, b := range newBindings {
		sn := b.Servicename.ValueString()
		if sn == "" || oldNames[sn] {
			continue
		}
		bind := gslb.Gslbvserverservicebinding{Name: vserverName, Servicename: sn}
		if !b.Domainname.IsNull() && !b.Domainname.IsUnknown() {
			bind.Domainname = b.Domainname.ValueString()
		}
		if !b.Weight.IsNull() && !b.Weight.IsUnknown() {
			bind.Weight = uint32(b.Weight.ValueInt64())
		}
		if _, err := r.client.UpdateResource(service.Gslbvserver_gslbservice_binding.Type(), vserverName, &bind); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to add service binding %s to gslbvserver %s: %s", sn, vserverName, err))
			return diags
		}
	}

	return diags
}

// readServiceBindings populates data.Service from the ADC.
func (r *GslbvserverResource) readServiceBindings(ctx context.Context, vserverName string, data *GslbvserverResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	bindings, _ := r.client.FindResourceArray(service.Gslbvserver_gslbservice_binding.Type(), vserverName)
	elems := make([]ServicebindingModel, 0, len(bindings))
	for _, m := range bindings {
		e := ServicebindingModel{
			Domainname:  types.StringNull(),
			Servicename: types.StringNull(),
			Weight:      types.Int64Null(),
		}
		if v, ok := m["domainname"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Domainname = types.StringValue(s)
			}
		}
		if v, ok := m["servicename"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Servicename = types.StringValue(s)
			}
		}
		if v, ok := m["weight"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Weight = types.Int64Value(iv)
			}
		}
		elems = append(elems, e)
	}

	setVal, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: servicebindingAttrTypes}, elems)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Service = setVal
	return diags
}
