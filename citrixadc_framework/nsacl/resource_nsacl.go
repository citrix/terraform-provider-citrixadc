package nsacl

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
var _ resource.Resource = &NsaclResource{}
var _ resource.ResourceWithConfigure = (*NsaclResource)(nil)
var _ resource.ResourceWithImportState = (*NsaclResource)(nil)

func NewNsaclResource() resource.Resource {
	return &NsaclResource{}
}

// NsaclResource defines the resource implementation.
type NsaclResource struct {
	client *service.NitroClient
}

func (r *NsaclResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsaclResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsacl"
}

func (r *NsaclResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsaclResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsaclResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsacl resource")

	// Determine the acl name. Matching SDK v2, aclname is optional; if not
	// provided a unique name is generated.
	var aclName string
	if data.Aclname.IsNull() || data.Aclname.IsUnknown() || data.Aclname.ValueString() == "" {
		aclName = fmt.Sprintf("tf-nsacl-%d", time.Now().UnixNano())
		data.Aclname = types.StringValue(aclName)
	} else {
		aclName = data.Aclname.ValueString()
	}

	nsacl := nsaclGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource
	_, err := r.client.AddResource(service.Nsacl.Type(), aclName, &nsacl)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsacl, got error: %s", err))
		return
	}

	// Set ID for the resource before reading state back
	data.Id = types.StringValue(aclName)

	tflog.Trace(ctx, "Created nsacl resource")

	// Read the updated state back
	if !r.readNsaclFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsacl not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaclResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsacl resource")

	found := r.readNsaclFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsaclResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NsaclResourceModel

	// Read Terraform prior state to determine the live name and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Determine attributes removed from config so they can be unset (reverted to
	// their NITRO defaults) after any update. The regular update below (driven by
	// nsaclHasUpdatableChange) pushes the schema default for these on removal;
	// the unset guarantees the appliance reverts to its documented default.
	attributesToUnset := []string{}
	if !data.Logstate.Equal(state.Logstate) && config.Logstate.IsNull() {
		attributesToUnset = append(attributesToUnset, "logstate")
	}
	if !data.Stateful.Equal(state.Stateful) && config.Stateful.IsNull() {
		attributesToUnset = append(attributesToUnset, "stateful")
	}

	// Preserve the live ID from prior state
	data.Id = state.Id
	// The current live name is tracked by the ID, not the (possibly stale) key.
	liveName := state.Id.ValueString()

	tflog.Debug(ctx, "Updating nsacl resource")

	// Handle in-place rename via the NITRO rename action when newname changes.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		newName := data.Newname.ValueString()
		tflog.Debug(ctx, fmt.Sprintf("Renaming nsacl %s to %s", liveName, newName))
		renamePayload := ns.Nsacl{
			Aclname: liveName,
			Newname: newName,
		}
		err := r.client.ActOnResource(service.Nsacl.Type(), &renamePayload, "rename")
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename nsacl, got error: %s", err))
			return
		}
		liveName = newName
		data.Id = types.StringValue(newName)
	}

	// Detect changes on the updatable (non-ForceNew, non-action) attributes.
	if nsaclHasUpdatableChange(&data, &state) {
		nsacl := nsaclGetTheUpdatablePayloadFromThePlan(ctx, &data)
		// Target the live name (the object may have just been renamed).
		nsacl.Aclname = liveName
		_, err := r.client.UpdateResource(service.Nsacl.Type(), liveName, &nsacl)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsacl %s, got error: %s", liveName, err))
			return
		}
		tflog.Trace(ctx, "Updated nsacl resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for nsacl resource, skipping update")
	}

	// Handle enable/disable state change via the NITRO action (as SDK v2 did).
	if !data.State.IsNull() && !data.State.IsUnknown() && !data.State.Equal(state.State) {
		err := r.doNsaclStateChange(ctx, liveName, data.State.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable/disable nsacl %s, got error: %s", liveName, err))
			return
		}
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"aclname": liveName,
	}
	if err := utils.ExecuteUnset(r.client, service.Nsacl.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsacl attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsaclFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsacl not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsaclResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsaclResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsacl resource")

	// Delete by the live name held in the ID.
	err := r.client.DeleteResource(service.Nsacl.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsacl, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsacl resource")
}

// doNsaclStateChange enables or disables the ACL rule via the NITRO action,
// mirroring the SDK v2 implementation.
func (r *NsaclResource) doNsaclStateChange(ctx context.Context, aclName string, newState string) error {
	tflog.Debug(ctx, fmt.Sprintf("In doNsaclStateChange for %s -> %s", aclName, newState))

	nsacl := ns.Nsacl{
		Aclname: aclName,
	}

	switch newState {
	case "ENABLED":
		return r.client.ActOnResource(service.Nsacl.Type(), &nsacl, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Nsacl.Type(), &nsacl, "disable")
	default:
		return fmt.Errorf("\"%s\" is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newState)
	}
}

// Helper function to read nsacl data from API. Returns false if the resource
// no longer exists on the appliance.
func (r *NsaclResource) readNsaclFromApi(ctx context.Context, data *NsaclResourceModel, diags *diag.Diagnostics) bool {
	aclName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsacl.Type(), aclName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsacl, got error: %s", err))
		return false
	}

	nsaclSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// nsaclHasUpdatableChange reports whether any NITRO-updatable attribute differs
// between plan and prior state. ForceNew attributes (type, datasets), the
// rename-only newname, and the action-driven state are excluded.
func nsaclHasUpdatableChange(data, state *NsaclResourceModel) bool {
	return !data.Aclaction.Equal(state.Aclaction) ||
		!data.Interface.Equal(state.Interface) ||
		!data.Destip.Equal(state.Destip) ||
		!data.Destipop.Equal(state.Destipop) ||
		!data.Destipval.Equal(state.Destipval) ||
		!data.Destport.Equal(state.Destport) ||
		!data.Destportop.Equal(state.Destportop) ||
		!data.Destportval.Equal(state.Destportval) ||
		!data.Dfdhash.Equal(state.Dfdhash) ||
		!data.Established.Equal(state.Established) ||
		!data.Icmpcode.Equal(state.Icmpcode) ||
		!data.Icmptype.Equal(state.Icmptype) ||
		!data.Logstate.Equal(state.Logstate) ||
		!data.Nodeid.Equal(state.Nodeid) ||
		!data.Priority.Equal(state.Priority) ||
		!data.Protocol.Equal(state.Protocol) ||
		!data.Protocolnumber.Equal(state.Protocolnumber) ||
		!data.Ratelimit.Equal(state.Ratelimit) ||
		!data.Srcip.Equal(state.Srcip) ||
		!data.Srcipop.Equal(state.Srcipop) ||
		!data.Srcipval.Equal(state.Srcipval) ||
		!data.Srcmac.Equal(state.Srcmac) ||
		!data.Srcmacmask.Equal(state.Srcmacmask) ||
		!data.Srcport.Equal(state.Srcport) ||
		!data.Srcportop.Equal(state.Srcportop) ||
		!data.Srcportval.Equal(state.Srcportval) ||
		!data.Stateful.Equal(state.Stateful) ||
		!data.Vlan.Equal(state.Vlan) ||
		!data.Vxlan.Equal(state.Vxlan)
}
