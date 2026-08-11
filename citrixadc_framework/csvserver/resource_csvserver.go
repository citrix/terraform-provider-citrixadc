package csvserver

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cs"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	sdkv2resource "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CsvserverResource{}
var _ resource.ResourceWithConfigure = (*CsvserverResource)(nil)
var _ resource.ResourceWithImportState = (*CsvserverResource)(nil)

func NewCsvserverResource() resource.Resource {
	return &CsvserverResource{}
}

// CsvserverResource defines the resource implementation.
type CsvserverResource struct {
	client *service.NitroClient
}

func (r *CsvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CsvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_csvserver"
}

func (r *CsvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CsvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CsvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating csvserver resource")

	// Backward-compatible with the SDK v2 resource: name is optional. When the
	// user does not supply a name, generate a unique one.
	csvserverName := data.Name.ValueString()
	if data.Name.IsNull() || csvserverName == "" {
		csvserverName = sdkv2resource.PrefixedUniqueId("tf-csvserver-")
		data.Name = types.StringValue(csvserverName)
	}

	// Create API request body from the model
	csvserver := csvserverGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - use AddResource (NITRO add is HTTP POST)
	_, err := r.client.AddResource(service.Csvserver.Type(), csvserverName, &csvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create csvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created csvserver resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(csvserverName)

	// Bind the convenience blocks (sslcertkey, sni certs, ciphers, ciphersuites,
	// sslprofile, lbvserverbinding, sslpolicybinding).
	r.applyBindingsOnCreate(ctx, csvserverName, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read the updated state back
	if !r.readCsvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "csvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CsvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading csvserver resource")

	found := r.readCsvserverFromApi(ctx, &data, &resp.Diagnostics)
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

// csvserverPayloadHasMutableFields reports whether the csvserver SET payload
// carries any attribute beyond the resource name. NITRO returns errorcode 1094
// ("Too few arguments") for a name-only csvserver SET, so the base update is
// skipped in that case. All cs.Csvserver fields use json omitempty, so a
// name-only payload marshals to {"name":...}.
func csvserverPayloadHasMutableFields(payload *cs.Csvserver) bool {
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

func (r *CsvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state CsvserverResourceModel

	// Read Terraform prior state to preserve ID / detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates to unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating csvserver resource")

	// Collect attributes that were removed from config so they can be reverted to
	// their NITRO defaults via ?action=unset after the base SET.
	attributesToUnset := []string{}
	if !data.Appflowlog.Equal(state.Appflowlog) && config.Appflowlog.IsNull() {
		attributesToUnset = append(attributesToUnset, "appflowlog")
	}
	if !data.Clttimeout.Equal(state.Clttimeout) && config.Clttimeout.IsNull() {
		attributesToUnset = append(attributesToUnset, "clttimeout")
	}
	if !data.Icmpvsrresponse.Equal(state.Icmpvsrresponse) && config.Icmpvsrresponse.IsNull() {
		attributesToUnset = append(attributesToUnset, "icmpvsrresponse")
	}
	if !data.L2conn.Equal(state.L2conn) && config.L2conn.IsNull() {
		attributesToUnset = append(attributesToUnset, "l2conn")
	}

	// The current live name on the appliance (tracks the ID, which follows any
	// prior rename).
	liveName := state.Id.ValueString()

	// Handle an in-place rename via the NITRO ?action=rename endpoint. newname is
	// rename-only and is excluded from the add/set payloads.
	if !data.Newname.IsNull() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming csvserver %s to %s", liveName, newName))
		renamePayload := cs.Csvserver{
			Name:    liveName,
			Newname: newName,
		}
		if err := r.client.ActOnResource(service.Csvserver.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename csvserver, got error: %s", err))
			return
		}
		liveName = newName
		data.Id = types.StringValue(newName)
	}

	// Update the mutable attributes via the NITRO set (PUT) endpoint. The set
	// endpoint does NOT accept the immutable attributes (port, range, servicetype,
	// targettype, td), the action-only "state" attribute, or "newname", so these
	// are cleared from the payload.
	csvserver := csvserverGetThePayloadFromtheConfig(ctx, &data)
	csvserver.Name = liveName
	csvserver.Port = nil
	csvserver.Range = nil
	csvserver.Td = nil
	csvserver.Servicetype = ""
	csvserver.Targettype = ""
	csvserver.State = ""

	// NITRO rejects a SET that carries only the name with errorcode 1094 ("Too
	// few arguments"). This happens on a state upgrade of a csvserver whose only
	// base attributes are immutable (e.g. a GSLB-target vserver): after the
	// immutable fields are cleared, and with the computed attributes still
	// unknown at plan time, the payload is name-only. Mirror the SDK v2 update's
	// hasChange gating and skip the base SET when there is nothing mutable to
	// send.
	if csvserverPayloadHasMutableFields(&csvserver) {
		_, err := r.client.UpdateResource(service.Csvserver.Type(), liveName, &csvserver)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update csvserver, got error: %s", err))
			return
		}
	}

	// The admin state is changed via the enable/disable actions (the set endpoint
	// does not accept "state").
	if !data.State.Equal(state.State) && !data.State.IsNull() && data.State.ValueString() != "" {
		if err := r.doCsvserverStateChange(ctx, liveName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change state of csvserver, got error: %s", err))
			return
		}
	}

	// Reconcile the convenience-block bindings against prior state.
	r.applyBindingsOnUpdate(ctx, liveName, &data, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Unset attributes that were removed from config so the appliance reverts them
	// to their defaults. Done after the base SET so any default value the SET
	// payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": liveName,
	}
	if err := utils.ExecuteUnset(r.client, service.Csvserver.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset csvserver attributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated csvserver resource")

	// Read the updated state back
	if !r.readCsvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "csvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CsvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CsvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting csvserver resource")

	// Named resource - delete using DeleteResource keyed on the live name (ID).
	err := r.client.DeleteResource(service.Csvserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete csvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted csvserver resource")
}

// doCsvserverStateChange enables/disables the content switching virtual server
// via the NITRO enable/disable actions. The set endpoint does not accept the
// "state" attribute directly.
func (r *CsvserverResource) doCsvserverStateChange(ctx context.Context, name string, newstate string) error {
	tflog.Debug(ctx, fmt.Sprintf("Changing state of csvserver %s to %s", name, newstate))

	// A minimal payload is required; enable/disable will fail with superfluous
	// attributes.
	csvserver := cs.Csvserver{
		Name: name,
	}

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Csvserver.Type(), &csvserver, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Csvserver.Type(), &csvserver, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read csvserver data from API. Returns false when the
// resource no longer exists on the appliance.
func (r *CsvserverResource) readCsvserverFromApi(ctx context.Context, data *CsvserverResourceModel, diags *diag.Diagnostics) bool {
	// Single unique attribute - the ID is the plain live name.
	csvserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Csvserver.Type(), csvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read csvserver, got error: %s", err))
		return false
	}

	csvserverSetAttrFromGet(ctx, data, getResponseData)

	// Refresh the managed convenience-block bindings.
	r.readBindings(ctx, csvserverName, data)

	return true
}
