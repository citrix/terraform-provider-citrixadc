package nsip

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
var _ resource.Resource = &NsipResource{}
var _ resource.ResourceWithConfigure = (*NsipResource)(nil)
var _ resource.ResourceWithImportState = (*NsipResource)(nil)

func NewNsipResource() resource.Resource {
	return &NsipResource{}
}

// NsipResource defines the resource implementation.
type NsipResource struct {
	client *service.NitroClient
}

func (r *NsipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NsipResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsip"
}

func (r *NsipResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NsipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NsipResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nsip resource")

	nsip := nsipGetThePayloadFromtheConfig(ctx, &data)

	// Named resource - the ipaddress is the resource name.
	ipaddress := data.Ipaddress.ValueString()
	_, err := r.client.AddResource(service.Nsip.Type(), ipaddress, &nsip)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nsip, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nsip resource")

	// The ID is the plain ipaddress value (matches the SDK v2 ID scheme and
	// resource_id_mapping.json single-key "ipaddress").
	data.Id = types.StringValue(ipaddress)

	// Read the updated state back
	if !r.readNsipFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsip not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NsipResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nsip resource")

	found := r.readNsipFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NsipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NsipResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nsip resource")

	ipaddress := data.Ipaddress.ValueString()

	// Build the update payload from the changed, updateable attributes only
	// (mirrors the SDK v2 HasChange semantics). The key/ForceNew attributes
	// (ipaddress, netmask, type, ownernode) never reach Update. The "state"
	// attribute is applied separately via the enable/disable action.
	nsip := ns.Nsip{
		Ipaddress: ipaddress,
	}
	hasChange := false

	if !data.Advertiseondefaultpartition.Equal(state.Advertiseondefaultpartition) && !data.Advertiseondefaultpartition.IsUnknown() {
		nsip.Advertiseondefaultpartition = data.Advertiseondefaultpartition.ValueString()
		hasChange = true
	}
	if !data.Arp.Equal(state.Arp) && !data.Arp.IsUnknown() {
		nsip.Arp = data.Arp.ValueString()
		hasChange = true
	}
	if !data.Arpowner.Equal(state.Arpowner) && !data.Arpowner.IsUnknown() {
		nsip.Arpowner = utils.IntPtr(int(data.Arpowner.ValueInt64()))
		hasChange = true
	}
	if !data.Arpresponse.Equal(state.Arpresponse) && !data.Arpresponse.IsUnknown() {
		nsip.Arpresponse = data.Arpresponse.ValueString()
		hasChange = true
	}
	if !data.Bgp.Equal(state.Bgp) && !data.Bgp.IsUnknown() {
		nsip.Bgp = data.Bgp.ValueString()
		hasChange = true
	}
	if !data.Decrementttl.Equal(state.Decrementttl) && !data.Decrementttl.IsUnknown() {
		nsip.Decrementttl = data.Decrementttl.ValueString()
		hasChange = true
	}
	if !data.Dynamicrouting.Equal(state.Dynamicrouting) && !data.Dynamicrouting.IsUnknown() {
		nsip.Dynamicrouting = data.Dynamicrouting.ValueString()
		hasChange = true
	}
	if !data.Ftp.Equal(state.Ftp) && !data.Ftp.IsUnknown() {
		nsip.Ftp = data.Ftp.ValueString()
		hasChange = true
	}
	if !data.Gui.Equal(state.Gui) && !data.Gui.IsUnknown() {
		nsip.Gui = data.Gui.ValueString()
		hasChange = true
	}
	if !data.Hostroute.Equal(state.Hostroute) && !data.Hostroute.IsUnknown() {
		nsip.Hostroute = data.Hostroute.ValueString()
		hasChange = true
	}
	if !data.Hostrtgw.Equal(state.Hostrtgw) && !data.Hostrtgw.IsUnknown() {
		nsip.Hostrtgw = data.Hostrtgw.ValueString()
		hasChange = true
	}
	if !data.Icmp.Equal(state.Icmp) && !data.Icmp.IsUnknown() {
		nsip.Icmp = data.Icmp.ValueString()
		hasChange = true
	}
	if !data.Icmpresponse.Equal(state.Icmpresponse) && !data.Icmpresponse.IsUnknown() {
		nsip.Icmpresponse = data.Icmpresponse.ValueString()
		hasChange = true
	}
	if !data.Metric.Equal(state.Metric) && !data.Metric.IsUnknown() {
		nsip.Metric = utils.IntPtr(int(data.Metric.ValueInt64()))
		hasChange = true
	}
	if !data.Mgmtaccess.Equal(state.Mgmtaccess) && !data.Mgmtaccess.IsUnknown() {
		nsip.Mgmtaccess = data.Mgmtaccess.ValueString()
		hasChange = true
	}
	if !data.Mptcpadvertise.Equal(state.Mptcpadvertise) && !data.Mptcpadvertise.IsUnknown() {
		nsip.Mptcpadvertise = data.Mptcpadvertise.ValueString()
		hasChange = true
	}
	if !data.Networkroute.Equal(state.Networkroute) && !data.Networkroute.IsUnknown() {
		nsip.Networkroute = data.Networkroute.ValueString()
		hasChange = true
	}
	if !data.Ospf.Equal(state.Ospf) && !data.Ospf.IsUnknown() {
		nsip.Ospf = data.Ospf.ValueString()
		hasChange = true
	}
	if !data.Ospfarea.Equal(state.Ospfarea) && !data.Ospfarea.IsUnknown() {
		nsip.Ospfarea = utils.IntPtr(int(data.Ospfarea.ValueInt64()))
		hasChange = true
	}
	if !data.Ospflsatype.Equal(state.Ospflsatype) && !data.Ospflsatype.IsUnknown() {
		nsip.Ospflsatype = data.Ospflsatype.ValueString()
		hasChange = true
	}
	if !data.Ownerdownresponse.Equal(state.Ownerdownresponse) && !data.Ownerdownresponse.IsUnknown() {
		nsip.Ownerdownresponse = data.Ownerdownresponse.ValueString()
		hasChange = true
	}
	if !data.Restrictaccess.Equal(state.Restrictaccess) && !data.Restrictaccess.IsUnknown() {
		nsip.Restrictaccess = data.Restrictaccess.ValueString()
		hasChange = true
	}
	if !data.Rip.Equal(state.Rip) && !data.Rip.IsUnknown() {
		nsip.Rip = data.Rip.ValueString()
		hasChange = true
	}
	if !data.Snmp.Equal(state.Snmp) && !data.Snmp.IsUnknown() {
		nsip.Snmp = data.Snmp.ValueString()
		hasChange = true
	}
	if !data.Ssh.Equal(state.Ssh) && !data.Ssh.IsUnknown() {
		nsip.Ssh = data.Ssh.ValueString()
		hasChange = true
	}
	if !data.Tag.Equal(state.Tag) && !data.Tag.IsUnknown() {
		nsip.Tag = utils.IntPtr(int(data.Tag.ValueInt64()))
		hasChange = true
	}
	if !data.Td.Equal(state.Td) && !data.Td.IsUnknown() {
		nsip.Td = utils.IntPtr(int(data.Td.ValueInt64()))
		hasChange = true
	}
	if !data.Telnet.Equal(state.Telnet) && !data.Telnet.IsUnknown() {
		nsip.Telnet = data.Telnet.ValueString()
		hasChange = true
	}
	if !data.Vrid.Equal(state.Vrid) && !data.Vrid.IsUnknown() {
		nsip.Vrid = utils.IntPtr(int(data.Vrid.ValueInt64()))
		hasChange = true
	}
	if !data.Vserver.Equal(state.Vserver) && !data.Vserver.IsUnknown() {
		nsip.Vserver = data.Vserver.ValueString()
		hasChange = true
	}
	if !data.Vserverrhilevel.Equal(state.Vserverrhilevel) && !data.Vserverrhilevel.IsUnknown() {
		nsip.Vserverrhilevel = data.Vserverrhilevel.ValueString()
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Nsip.Type(), ipaddress, &nsip)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nsip %s, got error: %s", ipaddress, err))
			return
		}
		tflog.Trace(ctx, "Updated nsip resource")
	} else {
		tflog.Debug(ctx, "No non-state changes detected for nsip resource, skipping update")
	}

	// Handle enable/disable via the state action (state is not a PUT field).
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doNsipStateChange(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling nsip %s, got error: %s", ipaddress, err))
			return
		}
	}

	// Read the updated state back
	if !r.readNsipFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nsip not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NsipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NsipResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nsip resource")

	ipaddress := data.Id.ValueString()
	argsMap := map[string]string{
		"td": fmt.Sprintf("%d", r.trafficDomain(&data)),
	}
	err := r.client.DeleteResourceWithArgsMap(service.Nsip.Type(), ipaddress, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nsip %s, got error: %s", ipaddress, err))
		return
	}

	tflog.Trace(ctx, "Deleted nsip resource")
}

// trafficDomain returns the td value from the model, defaulting to 0 when unset.
func (r *NsipResource) trafficDomain(data *NsipResourceModel) int64 {
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		return data.Td.ValueInt64()
	}
	return 0
}

// doNsipStateChange enables or disables the IP address via the NITRO action.
// A fresh, minimal payload is used because ActOnResource rejects superfluous
// attributes (mirrors the SDK v2 doNsipStateChange helper).
func (r *NsipResource) doNsipStateChange(ctx context.Context, data *NsipResourceModel) error {
	tflog.Debug(ctx, "In doNsipStateChange")

	nsip := ns.Nsip{
		Ipaddress: data.Ipaddress.ValueString(),
		Td:        utils.IntPtr(int(r.trafficDomain(data))),
	}

	newstate := data.State.ValueString()
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Nsip.Type(), nsip, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Nsip.Type(), nsip, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// readNsipFromApi reads the nsip from the appliance and maps it onto data.
// Returns false when the resource is missing (so callers can drop it from state).
func (r *NsipResource) readNsipFromApi(ctx context.Context, data *NsipResourceModel, diags *diag.Diagnostics) bool {
	ipaddress := data.Id.ValueString()
	argsMap := map[string]string{
		"td": fmt.Sprintf("%d", r.trafficDomain(data)),
	}
	findParams := service.FindParams{
		ResourceType:             service.Nsip.Type(),
		ResourceName:             ipaddress,
		ResourceMissingErrorCode: 258,
		ArgsMap:                  argsMap,
	}

	dataArr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nsip %s, got error: %s", ipaddress, err))
		return false
	}

	// Resource is missing
	if len(dataArr) == 0 {
		return false
	}

	// Iterate through results to find the one with the matching ipaddress
	foundIndex := -1
	for i, v := range dataArr {
		if addr, ok := v["ipaddress"].(string); ok && addr == ipaddress {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	nsipSetAttrFromGet(ctx, data, dataArr[foundIndex])

	return true
}
