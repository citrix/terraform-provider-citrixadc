package nsip6

import (
	"context"
	"fmt"
	"net/url"

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
var _ resource.Resource = &Nsip6Resource{}
var _ resource.ResourceWithConfigure = (*Nsip6Resource)(nil)
var _ resource.ResourceWithImportState = (*Nsip6Resource)(nil)

func NewNsip6Resource() resource.Resource {
	return &Nsip6Resource{}
}

// Nsip6Resource defines the resource implementation.
type Nsip6Resource struct {
	client *service.NitroClient
}

func (r *Nsip6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Nsip6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsip6"
}

func (r *Nsip6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Nsip6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Nsip6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsip6 resource")

	nsip6 := nsip6GetThePayloadFromtheConfig(ctx, &data)

	// Named resource - the ipv6address is the resource key (matches SDK v2
	// AddResource(type, ipv6address, ...)).
	ipv6address := data.Ipv6address.ValueString()
	_, err := r.client.AddResource(service.Nsip6.Type(), ipv6address, &nsip6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsip6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsip6 resource")

	// The ID is the plain ipv6address value (matches the SDK v2 ID scheme and
	// resource_id_mapping.json single-key "ipv6address").
	data.Id = types.StringValue(ipv6address)

	// Read the updated state back
	if !r.readNsip6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsip6 not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsip6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Nsip6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsip6 resource")

	found := r.readNsip6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Nsip6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Nsip6ResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (candidates for unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state (ipv6address is ForceNew, so it never changes)
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsip6 resource")

	// Build the update payload from the changed, updateable attributes only
	// (mirrors the SDK v2 HasChange semantics). The key/ForceNew attributes
	// (ipv6address, ownernode, scope, type, vlan) never reach Update. The SDK v2
	// resource issues the update with an empty resource name and the ipv6address
	// carried in the body, which is preserved here because the address contains
	// characters ("/", ":") that cannot go in the URL path.
	nsip6 := ns.Nsip6{
		Ipv6address: data.Ipv6address.ValueString(),
	}
	hasChange := false
	attributesToUnset := []string{}

	if !data.Advertiseondefaultpartition.IsUnknown() && !data.Advertiseondefaultpartition.Equal(state.Advertiseondefaultpartition) {
		nsip6.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
		hasChange = true
	}
	if !data.Decrementhoplimit.IsUnknown() && !data.Decrementhoplimit.Equal(state.Decrementhoplimit) {
		nsip6.Decrementhoplimit = data.Decrementhoplimit.ValueString()
		hasChange = true
	}
	if !data.Dynamicrouting.IsUnknown() && !data.Dynamicrouting.Equal(state.Dynamicrouting) {
		nsip6.Dynamicrouting = data.Dynamicrouting.ValueString()
		hasChange = true
	}
	if !data.Ftp.IsUnknown() && !data.Ftp.Equal(state.Ftp) {
		nsip6.Ftp = data.Ftp.ValueString()
		hasChange = true
	}
	if !data.Gui.IsUnknown() && !data.Gui.Equal(state.Gui) {
		nsip6.Gui = data.Gui.ValueString()
		hasChange = true
	}
	if !data.Hostroute.IsUnknown() && !data.Hostroute.Equal(state.Hostroute) {
		nsip6.Hostroute = data.Hostroute.ValueString()
		hasChange = true
	}
	if !data.Icmp.IsUnknown() && !data.Icmp.Equal(state.Icmp) {
		nsip6.Icmp = data.Icmp.ValueString()
		hasChange = true
	}
	if !data.Icmpresponse.IsUnknown() && !data.Icmpresponse.Equal(state.Icmpresponse) {
		nsip6.Icmpresponse = data.Icmpresponse.ValueString()
		hasChange = true
	}
	if !data.Ip6hostrtgw.IsUnknown() && !data.Ip6hostrtgw.Equal(state.Ip6hostrtgw) {
		nsip6.Ip6hostrtgw = data.Ip6hostrtgw.ValueString()
		hasChange = true
	}
	if !data.Map.IsUnknown() && !data.Map.Equal(state.Map) {
		nsip6.Map = data.Map.ValueString()
		hasChange = true
	}
	if !data.Metric.IsUnknown() && !data.Metric.Equal(state.Metric) {
		nsip6.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
		hasChange = true
	}
	if !data.Mgmtaccess.IsUnknown() && !data.Mgmtaccess.Equal(state.Mgmtaccess) {
		if config.Mgmtaccess.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "mgmtaccess")
		} else {
			nsip6.Mgmtaccess = data.Mgmtaccess.ValueString()
			hasChange = true
		}
	}
	if !data.Mptcpadvertise.IsUnknown() && !data.Mptcpadvertise.Equal(state.Mptcpadvertise) {
		nsip6.Mptcpadvertise = data.Mptcpadvertise.ValueString()
		hasChange = true
	}
	if !data.Nd.IsUnknown() && !data.Nd.Equal(state.Nd) {
		nsip6.Nd = data.Nd.ValueString()
		hasChange = true
	}
	if !data.Ndowner.IsUnknown() && !data.Ndowner.Equal(state.Ndowner) {
		nsip6.Ndowner = utils.IntPtr(int(data.Ndowner.ValueInt64()))
		hasChange = true
	}
	if !data.Networkroute.IsUnknown() && !data.Networkroute.Equal(state.Networkroute) {
		nsip6.Networkroute = data.Networkroute.ValueString()
		hasChange = true
	}
	if !data.Ospf6lsatype.IsUnknown() && !data.Ospf6lsatype.Equal(state.Ospf6lsatype) {
		nsip6.Ospf6lsatype = data.Ospf6lsatype.ValueString()
		hasChange = true
	}
	if !data.Ospfarea.IsUnknown() && !data.Ospfarea.Equal(state.Ospfarea) {
		nsip6.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
		hasChange = true
	}
	if !data.Ownerdownresponse.IsUnknown() && !data.Ownerdownresponse.Equal(state.Ownerdownresponse) {
		nsip6.Ownerdownresponse = data.Ownerdownresponse.ValueString()
		hasChange = true
	}
	if !data.Restrictaccess.IsUnknown() && !data.Restrictaccess.Equal(state.Restrictaccess) {
		if config.Restrictaccess.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "restrictaccess")
		} else {
			nsip6.Restrictaccess = data.Restrictaccess.ValueString()
			hasChange = true
		}
	}
	if !data.Snmp.IsUnknown() && !data.Snmp.Equal(state.Snmp) {
		nsip6.Snmp = data.Snmp.ValueString()
		hasChange = true
	}
	if !data.Ssh.IsUnknown() && !data.Ssh.Equal(state.Ssh) {
		nsip6.Ssh = data.Ssh.ValueString()
		hasChange = true
	}
	if !data.State.IsUnknown() && !data.State.Equal(state.State) {
		nsip6.State = data.State.ValueString()
		hasChange = true
	}
	if !data.Tag.IsUnknown() && !data.Tag.Equal(state.Tag) {
		nsip6.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
		hasChange = true
	}
	if !data.Td.IsUnknown() && !data.Td.Equal(state.Td) {
		nsip6.Td = utils.IntPtr(int(data.Td.ValueInt64()))
		hasChange = true
	}
	if !data.Telnet.IsUnknown() && !data.Telnet.Equal(state.Telnet) {
		nsip6.Telnet = data.Telnet.ValueString()
		hasChange = true
	}
	if !data.Vrid6.IsUnknown() && !data.Vrid6.Equal(state.Vrid6) {
		nsip6.Vrid6 = utils.IntPtr(int(data.Vrid6.ValueInt64()))
		hasChange = true
	}
	if !data.Vserver.IsUnknown() && !data.Vserver.Equal(state.Vserver) {
		nsip6.Vserver = data.Vserver.ValueString()
		hasChange = true
	}
	if !data.Vserverrhilevel.IsUnknown() && !data.Vserverrhilevel.Equal(state.Vserverrhilevel) {
		nsip6.Vserverrhilevel = data.Vserverrhilevel.ValueString()
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Nsip6.Type(), "", &nsip6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsip6 %s, got error: %s", data.Id.ValueString(), err))
			return
		}
		tflog.Trace(ctx, "Updated nsip6 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for nsip6 resource, skipping update")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults. The resource is keyed by ipv6address (which cannot
	// go in the URL path); it is carried in the unset body along with td.
	unsetIdPayload := map[string]interface{}{
		"ipv6address": data.Ipv6address.ValueString(),
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		unsetIdPayload["td"] = int(data.Td.ValueInt64())
	}
	if err := utils.ExecuteUnset(r.client, service.Nsip6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nsip6 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNsip6FromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsip6 not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Nsip6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Nsip6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsip6 resource")

	// The ipv6address contains "/" and ":", so it cannot be placed in the URL
	// path. Mirror the SDK v2 delete: empty resource name with the (URL-escaped)
	// ipv6address and td carried in the query args.
	ipv6address := data.Id.ValueString()
	argsMap := map[string]string{
		"ipv6address": url.QueryEscape(ipv6address),
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		argsMap["td"] = fmt.Sprintf("%d", data.Td.ValueInt64())
	}
	err := r.client.DeleteResourceWithArgsMap(service.Nsip6.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsip6 %s, got error: %s", ipv6address, err))
		return
	}

	tflog.Trace(ctx, "Deleted nsip6 resource")
}

// readNsip6FromApi reads the nsip6 from the appliance and maps it onto data.
// Returns false when the resource is missing (so callers can drop it from
// state). Mirrors the SDK v2 read, which lists all nsip6 objects and matches on
// ipv6address (the address cannot be used directly in the URL path).
func (r *Nsip6Resource) readNsip6FromApi(ctx context.Context, data *Nsip6ResourceModel, diags *diag.Diagnostics) bool {
	ipv6address := data.Id.ValueString()

	findParams := service.FindParams{
		ResourceType:             service.Nsip6.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsip6 %s, got error: %s", ipv6address, err))
		return false
	}

	// Resource is missing
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the matching ipv6address
	foundIndex := -1
	for i, v := range dataArr {
		if addr, ok := v["ipv6address"].(string); ok && addr == ipv6address {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	nsip6SetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
