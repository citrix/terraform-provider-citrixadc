package lbvserver

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkid "github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &LbvserverResource{}
var _ resource.ResourceWithConfigure = (*LbvserverResource)(nil)
var _ resource.ResourceWithImportState = (*LbvserverResource)(nil)

func NewLbvserverResource() resource.Resource {
	return &LbvserverResource{}
}

// LbvserverResource defines the resource implementation.
type LbvserverResource struct {
	client *service.NitroClient
}

// lbvserverUnsettableAttrs is the set of attributes for which the SDK v2 resource
// issued an explicit NITRO ?action=unset when they were removed from config or set
// to an empty/zero value. Preserving this behaviour keeps the migration backward
// compatible (Optional+Computed alone would leave the old value on the appliance).
var lbvserverUnsettableAttrs = []string{
	"quicprofilename", "apiprofile", "adfsproxyprofile", "appflowlog", "authentication",
	"authenticationhost", "authn401", "authnprofile", "authnvsname", "backuplbmethod",
	"backupvserver", "cacheable", "comment", "connfailover", "cookiename", "dbprofilename",
	"dbslb", "disableprimaryondown", "dnsprofilename", "downstateflush", "httpprofilename",
	"httpsredirecturl", "icmpvsrresponse", "insertvserveripport", "ipset", "l2conn",
	"lbmethod", "lbprofilename", "listenpolicy", "m", "macmoderetainvlan", "mssqlserverversion",
	"mysqlserverversion", "netmask", "netprofile", "newservicerequestunit", "oracleserverversion",
	"persistencebackup", "persistencetype", "persistmask", "probeprotocol", "pushlabel",
	"pushvserver", "redirectfromport", "redirurl", "tcpprofilename", "vipheader",
	"quicbridgeprofilename",
}

func (r *LbvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LbvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lbvserver"
}

func (r *LbvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LbvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LbvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lbvserver resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the user
	// does not supply a name, generate a unique one.
	lbvserverName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || lbvserverName == "" {
		lbvserverName = sdkid.PrefixedUniqueId("tf-lbvserver-")
		data.Name = types.StringValue(lbvserverName)
	}

	// Create API request body from the model
	lbvserver := lbvserverGetThePayloadFromtheConfig(ctx, &data)
	lbvserver.Name = lbvserverName

	// Pre-check referenced SSL cert key(s) exist before creating the vserver, so a
	// missing cert fails cleanly up-front instead of orphaning a created vserver
	// (SDK v2 parity).
	if !r.precheckSslcertkeysExist(ctx, &data, &resp.Diagnostics) {
		return
	}

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Lbvserver.Type(), lbvserverName, &lbvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lbvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lbvserver resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(lbvserverName)

	// Bind the convenience blocks (sslcertkey, sni certs, ciphers, ciphersuites,
	// sslprofile, sslpolicybinding).
	r.applyBindingsOnCreate(ctx, lbvserverName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		// Roll back the just-created vserver so a failed bind does not leave an
		// orphan on the appliance (SDK v2 parity).
		if delErr := r.client.DeleteResource(service.Lbvserver.Type(), lbvserverName); delErr != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to roll back lbvserver %q after a binding failure: %s", lbvserverName, delErr))
		}
		return
	}

	// Read the updated state back
	if !r.readLbvserverFromApi(ctx, &data, &resp.Diagnostics, true) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LbvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lbvserver resource")

	found := r.readLbvserverFromApi(ctx, &data, &resp.Diagnostics, false)
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

// lbvserverPayloadHasMutableFields reports whether the lbvserver SET payload carries
// any attribute beyond the resource name. NITRO rejects a name-only SET, so the base
// update is skipped in that case (mirrors the SDK v2 hasChange gating).
func lbvserverPayloadHasMutableFields(payload *lb.Lbvserver) bool {
	b, err := json.Marshal(payload)
	if err != nil {
		return true // fail open: let NITRO validate the payload
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return true
	}
	delete(m, "name")
	return len(m) > 0
}

func (r *LbvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state, config LbvserverResourceModel

	// Read prior state, plan and config.
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lbvserver resource")

	// The current live name on the appliance (tracks the ID).
	lbvserverName := state.Id.ValueString()

	// Handle an in-place rename via the NITRO ?action=rename endpoint. newname is
	// rename-only and is excluded from the add/set payloads.
	if !data.Newname.IsNull() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming lbvserver %s to %s", lbvserverName, newName))
		renamePayload := lb.Lbvserver{
			Name:    lbvserverName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Lbvserver.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename lbvserver, got error: %s", err))
			return
		}
		lbvserverName = newName
		data.Id = types.StringValue(newName)
		data.Name = types.StringValue(newName)
	}

	// Update the mutable attributes via the NITRO set (PUT) endpoint. The set
	// endpoint does NOT accept the immutable attributes (port, range, servicetype,
	// td, redirurlflags), the action-only "state" attribute, or "newname".
	lbvserver := lbvserverGetThePayloadFromtheConfig(ctx, &data)
	lbvserver.Name = lbvserverName
	lbvserver.Port = nil
	lbvserver.Range = nil
	lbvserver.Td = nil
	lbvserver.Servicetype = ""
	lbvserver.Redirurlflags = false
	lbvserver.State = ""

	if lbvserverPayloadHasMutableFields(&lbvserver) {
		_, err := r.client.UpdateResource(service.Lbvserver.Type(), lbvserverName, &lbvserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lbvserver, got error: %s", err))
			return
		}
	}

	// Unset attributes that were removed from config or set to empty/zero, mirroring
	// the SDK v2 ?action=unset behaviour.
	unsetAttrs := collectLbvserverUnsetAttrs(&config, &state)
	if err := r.executeLbvserverUnset(ctx, lbvserverName, unsetAttrs); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset lbvserver attributes, got error: %s", err))
		return
	}

	// The admin state is changed via the enable/disable actions.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() && data.State.ValueString() != "" {
		if err := r.doLbvserverStateChange(ctx, lbvserverName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change state of lbvserver, got error: %s", err))
			return
		}
	}

	// Reconcile the convenience-block bindings against prior state.
	r.applyBindingsOnUpdate(ctx, lbvserverName, &data, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Updated lbvserver resource")

	// Read the updated state back
	if !r.readLbvserverFromApi(ctx, &data, &resp.Diagnostics, true) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "lbvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LbvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LbvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lbvserver resource")

	// Named resource - delete using DeleteResource keyed on the live name (ID).
	err := r.client.DeleteResource(service.Lbvserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lbvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lbvserver resource")
}

// doLbvserverStateChange enables/disables the load balancing virtual server via the
// NITRO enable/disable actions. The set endpoint does not accept the "state" attribute.
func (r *LbvserverResource) doLbvserverStateChange(ctx context.Context, name string, newstate string) error {
	tflog.Debug(ctx, fmt.Sprintf("Changing state of lbvserver %s to %s", name, newstate))

	// A minimal payload is required; enable/disable will fail with superfluous attributes.
	lbvserver := lb.Lbvserver{
		Name: name,
	}

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Lbvserver.Type(), &lbvserver, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Lbvserver.Type(), &lbvserver, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// executeLbvserverUnset performs a single NITRO ?action=unset call for the collected
// attribute names.
func (r *LbvserverResource) executeLbvserverUnset(ctx context.Context, name string, attrs []string) error {
	if len(attrs) == 0 {
		return nil
	}
	tflog.Debug(ctx, fmt.Sprintf("Unsetting lbvserver %s attributes: %v", name, attrs))
	unsetData := map[string]interface{}{"name": name}
	for _, a := range attrs {
		unsetData[a] = "true"
	}
	return r.client.ActOnResource(service.Lbvserver.Type(), unsetData, "unset")
}

// collectLbvserverUnsetAttrs returns the attributes that should be unset: those with a
// non-empty/non-zero prior state value that are now absent from config or set to an
// empty/zero value.
func collectLbvserverUnsetAttrs(config, state *LbvserverResourceModel) []string {
	out := []string{}
	for _, name := range lbvserverUnsettableAttrs {
		cfgV := lbvserverFieldByTag(config, name)
		stV := lbvserverFieldByTag(state, name)
		if cfgV == nil || stV == nil {
			continue
		}
		if lbvserverStateHasValue(stV) && lbvserverConfigIsEmpty(cfgV) {
			out = append(out, name)
		}
	}
	return out
}

func lbvserverFieldByTag(m *LbvserverResourceModel, tag string) attr.Value {
	rv := reflect.ValueOf(m).Elem()
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		if rt.Field(i).Tag.Get("tfsdk") == tag {
			if v, ok := rv.Field(i).Interface().(attr.Value); ok {
				return v
			}
		}
	}
	return nil
}

func lbvserverStateHasValue(v attr.Value) bool {
	switch t := v.(type) {
	case types.String:
		return !t.IsNull() && !t.IsUnknown() && t.ValueString() != ""
	case types.Int64:
		return !t.IsNull() && !t.IsUnknown() && t.ValueInt64() != 0
	}
	return false
}

func lbvserverConfigIsEmpty(v attr.Value) bool {
	switch t := v.(type) {
	case types.String:
		return t.IsNull() || (!t.IsUnknown() && t.ValueString() == "")
	case types.Int64:
		return t.IsNull() || (!t.IsUnknown() && t.ValueInt64() == 0)
	}
	return true
}

// readLbvserverFromApi reads lbvserver data from the appliance. Returns false when the
// resource no longer exists on the appliance. When preservePlan is true (Create/Update)
// known planned scalar values are retained across the read-back so the applied state
// stays consistent with the plan; plain Read passes false to adopt appliance values
// (drift detection).
func (r *LbvserverResource) readLbvserverFromApi(ctx context.Context, data *LbvserverResourceModel, diags *diag.Diagnostics, preservePlan bool) bool {
	// Single unique attribute - the ID is the plain live name.
	lbvserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Lbvserver.Type(), lbvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lbvserver, got error: %s", err))
		return false
	}

	var planSnapshot LbvserverResourceModel
	if preservePlan {
		planSnapshot = *data
	}

	lbvserverSetAttrFromGet(ctx, data, getResponseData)

	// On Create/Update the read-back must not overwrite a known planned value with a
	// server-side value that was never in the plan. NITRO occasionally starts
	// reporting a default (e.g. backuppersistencetimeout=2) only after a SET, which
	// would otherwise trip Terraform's "inconsistent result after apply" check for
	// Optional+Computed attributes that were null in the plan. Preserve every known
	// planned scalar and let GET resolve only the unknowns. Plain Read passes
	// preservePlan=false so drift is still detected. (Pattern 7)
	if preservePlan {
		lbvserverRestoreKnownScalars(data, &planSnapshot)
	}

	// Refresh the managed convenience-block bindings.
	r.readBindings(ctx, lbvserverName, data)

	return true
}

// lbvserverRestoreKnownScalars copies every known (non-unknown) scalar attr.Value from
// plan into post, so the Create/Update read-back never replaces a value the plan already
// fixed with a differing appliance value (which would fail Terraform's post-apply
// consistency check). Unknown planned values are left as resolved by the GET read-back.
// Only scalar attributes (String/Int64/Bool) are handled; List/Set fields (persistavpno
// and the convenience-block bindings) keep their existing read-back/readBindings handling.
func lbvserverRestoreKnownScalars(post, plan *LbvserverResourceModel) {
	pv := reflect.ValueOf(plan).Elem()
	ov := reflect.ValueOf(post).Elem()
	for i := 0; i < pv.NumField(); i++ {
		planField := pv.Field(i)
		av, ok := planField.Interface().(attr.Value)
		if !ok {
			continue
		}
		switch av.(type) {
		case types.String, types.Int64, types.Bool:
			if av.IsUnknown() {
				continue // let the GET read-back resolve unknowns
			}
			ov.Field(i).Set(planField)
		}
	}
}
