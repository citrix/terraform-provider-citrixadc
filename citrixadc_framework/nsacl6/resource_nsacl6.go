package nsacl6

import (
	"context"
	"fmt"

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
var _ resource.Resource = &Nsacl6Resource{}
var _ resource.ResourceWithConfigure = (*Nsacl6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nsacl6Resource)(nil)

func NewNsacl6Resource() resource.Resource {
	return &Nsacl6Resource{}
}

// Nsacl6Resource defines the resource implementation.
type Nsacl6Resource struct {
	client *service.NitroClient
}

func (r *Nsacl6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nsacl6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsacl6"
}

func (r *Nsacl6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nsacl6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nsacl6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsacl6 resource")

	// Named resource - use AddResource
	nsacl6 := nsacl6GetThePayloadFromthePlan(ctx, &data)
	acl6name := data.Acl6name.ValueString()
	_, err := r.client.AddResource(service.Nsacl6.Type(), acl6name, &nsacl6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsacl6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsacl6 resource")

	// Set ID for the resource before reading state (matches SDK v2: d.SetId(acl6name))
	data.Id = types.StringValue(acl6name)

	// Read the updated state back
	if !r.readNsacl6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsacl6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacl6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nsacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsacl6 resource")

	found := r.readNsacl6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nsacl6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Nsacl6ResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsacl6 resource")

	acl6name := data.Acl6name.ValueString()

	// Detect changes in updatable attributes (state is handled separately via enable/disable)
	hasChange := false
	attributesToUnset := []string{}
	if !data.Acl6action.Equal(state.Acl6action) {
		hasChange = true
	}
	if !data.Aclaction.Equal(state.Aclaction) {
		hasChange = true
	}
	if !data.Destipop.Equal(state.Destipop) {
		hasChange = true
	}
	if !data.Destipv6.Equal(state.Destipv6) {
		hasChange = true
	}
	if !data.Destipv6val.Equal(state.Destipv6val) {
		hasChange = true
	}
	if !data.Destport.Equal(state.Destport) {
		hasChange = true
	}
	if !data.Destportop.Equal(state.Destportop) {
		hasChange = true
	}
	if !data.Destportval.Equal(state.Destportval) {
		hasChange = true
	}
	if !data.Dfdhash.Equal(state.Dfdhash) {
		hasChange = true
	}
	if !data.Dfdprefix.Equal(state.Dfdprefix) {
		hasChange = true
	}
	if !data.Established.Equal(state.Established) {
		hasChange = true
	}
	if !data.Icmpcode.Equal(state.Icmpcode) {
		hasChange = true
	}
	if !data.Icmptype.Equal(state.Icmptype) {
		hasChange = true
	}
	if !data.Interface.Equal(state.Interface) {
		hasChange = true
	}
	if !data.Logstate.Equal(state.Logstate) {
		if config.Logstate.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "logstate")
		} else {
			hasChange = true
		}
	}
	if !data.Nodeid.Equal(state.Nodeid) {
		hasChange = true
	}
	if !data.Priority.Equal(state.Priority) {
		hasChange = true
	}
	if !data.Protocol.Equal(state.Protocol) {
		hasChange = true
	}
	if !data.Protocolnumber.Equal(state.Protocolnumber) {
		hasChange = true
	}
	if !data.Ratelimit.Equal(state.Ratelimit) {
		hasChange = true
	}
	if !data.Srcipop.Equal(state.Srcipop) {
		hasChange = true
	}
	if !data.Srcipv6.Equal(state.Srcipv6) {
		hasChange = true
	}
	if !data.Srcipv6val.Equal(state.Srcipv6val) {
		hasChange = true
	}
	if !data.Srcmac.Equal(state.Srcmac) {
		hasChange = true
	}
	if !data.Srcmacmask.Equal(state.Srcmacmask) {
		hasChange = true
	}
	if !data.Srcport.Equal(state.Srcport) {
		hasChange = true
	}
	if !data.Srcportop.Equal(state.Srcportop) {
		hasChange = true
	}
	if !data.Srcportval.Equal(state.Srcportval) {
		hasChange = true
	}
	if !data.Stateful.Equal(state.Stateful) {
		if config.Stateful.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "stateful")
		} else {
			hasChange = true
		}
	}
	if !data.Td.Equal(state.Td) {
		hasChange = true
	}
	if !data.Ttl.Equal(state.Ttl) {
		hasChange = true
	}
	if !data.Vlan.Equal(state.Vlan) {
		hasChange = true
	}
	if !data.Vxlan.Equal(state.Vxlan) {
		hasChange = true
	}

	if hasChange {
		nsacl6 := nsacl6GetTheUpdatablePayloadFromThePlan(ctx, &data)
		_, err := r.client.UpdateResource(service.Nsacl6.Type(), acl6name, &nsacl6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsacl6, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nsacl6 resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for nsacl6 resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. Done after the update so any default value the
	// update payload carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"acl6name": data.Acl6name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nsacl6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsacl6 attributes, got error: %s", err))
		return
	}

	// Handle state change (ENABLED/DISABLED) via enable/disable action, matching SDK v2
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doNsacl6StateChange(&data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling nsacl6 %s: %s", acl6name, err))
			return
		}
	}

	// Read the updated state back
	if !r.readNsacl6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsacl6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsacl6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nsacl6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsacl6 resource")

	// Named resource - delete using DeleteResource keyed by ID (acl6name)
	err := r.client.DeleteResource(service.Nsacl6.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsacl6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nsacl6 resource")
}

// doNsacl6StateChange enables or disables the ACL6 rule, mirroring the SDK v2 doNsacl6StateSchange.
func (r *Nsacl6Resource) doNsacl6StateChange(data *Nsacl6ResourceModel) error {
	// We need a struct with only the key since ActOnResource fails on superfluous attributes.
	nsacl6 := ns.Nsacl6{
		Acl6name: data.Acl6name.ValueString(),
	}

	newstate := data.State.ValueString()
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Nsacl6.Type(), nsacl6, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Nsacl6.Type(), nsacl6, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// Helper function to read nsacl6 data from API. Returns false when the resource no longer exists.
func (r *Nsacl6Resource) readNsacl6FromApi(ctx context.Context, data *Nsacl6ResourceModel, diags *diag.Diagnostics) bool {
	acl6name := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nsacl6.Type(), acl6name)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsacl6, got error: %s", err))
		return false
	}

	nsacl6SetAttrFromGet(ctx, data, getResponseData)

	return true
}
