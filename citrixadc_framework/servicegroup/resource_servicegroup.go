package servicegroup

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServicegroupResource{}
var _ resource.ResourceWithConfigure = (*ServicegroupResource)(nil)
var _ resource.ResourceWithImportState = (*ServicegroupResource)(nil)

func NewServicegroupResource() resource.Resource {
	return &ServicegroupResource{}
}

// ServicegroupResource defines the resource implementation.
type ServicegroupResource struct {
	client *service.NitroClient
}

func (r *ServicegroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServicegroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_servicegroup"
}

func (r *ServicegroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServicegroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServicegroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating servicegroup resource")

	servicegroupName := data.Servicegroupname.ValueString()

	// Validate lbmonitor exists (backward-compat with SDK v2).
	if !data.Lbmonitor.IsNull() && !data.Lbmonitor.IsUnknown() && data.Lbmonitor.ValueString() != "" {
		if !r.client.ResourceExists(service.Lbmonitor.Type(), data.Lbmonitor.ValueString()) {
			resp.Diagnostics.AddError("Client Error", "Specified lb monitor does not exist on netscaler!")
			return
		}
	}

	// Validate lbvservers exist (backward-compat with SDK v2).
	lbvservers := setToStringSlice(ctx, data.Lbvservers, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	for _, lbvserver := range lbvservers {
		if !r.client.ResourceExists(service.Lbvserver.Type(), lbvserver) {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Specified lb vserver %s does not exist on netscaler!", lbvserver))
			return
		}
	}

	groupmembers := setToStringSlice(ctx, data.Servicegroupmembers, &resp.Diagnostics)
	groupmembersByServername := setToStringSlice(ctx, data.ServicegroupmembersByServername, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build and push the add payload.
	servicegroup := servicegroupGetThePayloadFromthePlan(ctx, &data)
	_, err := r.client.AddResource(service.Servicegroup.Type(), servicegroupName, &servicegroup)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create servicegroup, got error: %s", err))
		return
	}

	// lbvservers bindings
	if !data.Lbvservers.IsNull() && len(lbvservers) > 0 {
		if err := addLbvserverBindings(r.client, servicegroupName, lbvservers); err != nil {
			resp.Diagnostics.AddError("Client Error", err.Error())
			return
		}
	}

	// lbmonitor binding
	if !data.Lbmonitor.IsNull() && !data.Lbmonitor.IsUnknown() && data.Lbmonitor.ValueString() != "" {
		lbmonitorName := data.Lbmonitor.ValueString()
		binding := lb.Lbmonitorservicebinding{
			Monitorname:      lbmonitorName,
			Servicegroupname: servicegroupName,
		}
		if err := r.client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Servicegroup.Type(), servicegroupName, &binding); err != nil {
			// Roll back the servicegroup on a failed monitor bind (SDK v2 parity).
			_ = r.client.DeleteResource(service.Servicegroup.Type(), servicegroupName)
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind lb monitor %s to servicegroup %s", lbmonitorName, servicegroupName))
			return
		}
	}

	// servicegroupmembers bindings
	if !data.Servicegroupmembers.IsNull() && len(groupmembers) > 0 {
		createServicegroupMemberBindings(r.client, servicegroupName, groupmembers, false)
	}
	if !data.ServicegroupmembersByServername.IsNull() && len(groupmembersByServername) > 0 {
		createServicegroupMemberBindings(r.client, servicegroupName, groupmembersByServername, true)
	}

	tflog.Trace(ctx, "Created servicegroup resource")

	// Set ID for the resource before reading state (plain servicegroupname value).
	data.Id = types.StringValue(servicegroupName)

	// Read the updated state back
	r.readServicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading servicegroup resource")

	r.readServicegroupFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// servicegroupPayloadHasMutableFields reports whether the servicegroup SET payload
// carries any attribute beyond the resource name key. NITRO rejects a name-only SET
// with errorcode 1094 ("Too few arguments"), so the base update is skipped in that
// case (mirrors lbvserverPayloadHasMutableFields and the SDK v2 hasChange gating).
func servicegroupPayloadHasMutableFields(payload *basic.Servicegroup) bool {
	b, err := json.Marshal(payload)
	if err != nil {
		return true // fail open: let NITRO validate the payload
	}
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return true
	}
	delete(m, "servicegroupname")
	return len(m) > 0
}

func (r *ServicegroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ServicegroupResourceModel

	// Read Terraform prior state to preserve ID and diff convenience blocks
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id
	servicegroupName := data.Id.ValueString()

	tflog.Debug(ctx, "Updating servicegroup resource")

	// Collect mutable attributes that were removed from config so they are reverted
	// to their NITRO defaults via ?action=unset. The unsetOnRemove plan modifiers on
	// these attributes force the planned value to unknown on config removal, so a
	// diff against prior state is detected here. For each attribute being unset we
	// copy the prior-state value back into data so it does not, by itself, flip
	// servicegroupAttrsChanged (which would trigger a full base SET carrying other
	// read-back defaults like nameserver=0.0.0.0 that NITRO rejects). The subsequent
	// read-back repopulates data from the appliance.
	attributesToUnset := []string{}
	if !data.Appflowlog.Equal(state.Appflowlog) && config.Appflowlog.IsNull() {
		attributesToUnset = append(attributesToUnset, "appflowlog")
		data.Appflowlog = state.Appflowlog
	}
	if !data.Cacheable.Equal(state.Cacheable) && config.Cacheable.IsNull() {
		attributesToUnset = append(attributesToUnset, "cacheable")
		data.Cacheable = state.Cacheable
	}
	if !data.Downstateflush.Equal(state.Downstateflush) && config.Downstateflush.IsNull() {
		attributesToUnset = append(attributesToUnset, "downstateflush")
		data.Downstateflush = state.Downstateflush
	}
	if !data.Healthmonitor.Equal(state.Healthmonitor) && config.Healthmonitor.IsNull() {
		attributesToUnset = append(attributesToUnset, "healthmonitor")
		data.Healthmonitor = state.Healthmonitor
	}
	if !data.Monconnectionclose.Equal(state.Monconnectionclose) && config.Monconnectionclose.IsNull() {
		attributesToUnset = append(attributesToUnset, "monconnectionclose")
		data.Monconnectionclose = state.Monconnectionclose
	}
	if !data.Sp.Equal(state.Sp) && config.Sp.IsNull() {
		attributesToUnset = append(attributesToUnset, "sp")
		data.Sp = state.Sp
	}

	// Determine which convenience-block changes to apply.
	planLbvservers := setToStringSlice(ctx, data.Lbvservers, &resp.Diagnostics)
	stateLbvservers := setToStringSlice(ctx, state.Lbvservers, &resp.Diagnostics)
	planMembers := setToStringSlice(ctx, data.Servicegroupmembers, &resp.Diagnostics)
	stateMembers := setToStringSlice(ctx, state.Servicegroupmembers, &resp.Diagnostics)
	planMembersByName := setToStringSlice(ctx, data.ServicegroupmembersByServername, &resp.Diagnostics)
	stateMembersByName := setToStringSlice(ctx, state.ServicegroupmembersByServername, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Detect whether any updateable servicegroup attribute changed. Guard the SET
	// with servicegroupPayloadHasMutableFields so a payload that carries only the
	// name key (e.g. every mutable attr resolved to unknown/null on a refresh of
	// SDK-v2-written state) never reaches NITRO as a name-only SET (errorcode 1094
	// "Too few arguments"). Belt-and-suspenders with the UseStateForUnknown plan
	// modifiers that keep those attrs known in the plan.
	if servicegroupAttrsChanged(&data, &state) {
		updatePayload := servicegroupGetTheUpdatePayloadFromthePlan(ctx, &data)
		updatePayload.Servicegroupname = servicegroupName
		if servicegroupPayloadHasMutableFields(&updatePayload) {
			if _, err := r.client.UpdateResource(service.Servicegroup.Type(), servicegroupName, &updatePayload); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update servicegroup %s, got error: %s", servicegroupName, err))
				return
			}
		}
	}

	// lbvservers binding diff.
	if !data.Lbvservers.Equal(state.Lbvservers) {
		add, remove := stringSliceDiff(stateLbvservers, planLbvservers)
		if len(remove) > 0 {
			if err := removeLbvserverBindings(r.client, servicegroupName, remove); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
		if len(add) > 0 {
			if err := addLbvserverBindings(r.client, servicegroupName, add); err != nil {
				resp.Diagnostics.AddError("Client Error", err.Error())
				return
			}
		}
	}

	// lbmonitor binding change (unbind old, bind new).
	if !data.Lbmonitor.Equal(state.Lbmonitor) {
		oldLbmonitor := state.Lbmonitor.ValueString()
		if oldLbmonitor != "" {
			if err := r.client.UnbindResource(service.Lbmonitor.Type(), oldLbmonitor, service.Servicegroup.Type(), servicegroupName, "servicegroupname"); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error unbinding lbmonitor %s from servicegroup %s", oldLbmonitor, servicegroupName))
				return
			}
		}
		newLbmonitor := data.Lbmonitor.ValueString()
		if !data.Lbmonitor.IsNull() && newLbmonitor != "" {
			binding := lb.Lbmonitorservicebinding{
				Monitorname:      newLbmonitor,
				Servicegroupname: servicegroupName,
			}
			if err := r.client.BindResource(service.Lbmonitor.Type(), newLbmonitor, service.Servicegroup.Type(), servicegroupName, &binding); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind lb monitor %s to servicegroup %s", newLbmonitor, servicegroupName))
				return
			}
		}
	}

	// servicegroupmembers diff.
	if !data.Servicegroupmembers.Equal(state.Servicegroupmembers) {
		add, remove := stringSliceDiff(stateMembers, planMembers)
		if len(remove) > 0 {
			removeServicegroupMemberBindings(r.client, servicegroupName, remove, false)
		}
		if len(add) > 0 {
			createServicegroupMemberBindings(r.client, servicegroupName, add, false)
		}
	}

	// servicegroupmembers_by_servername diff.
	if !data.ServicegroupmembersByServername.Equal(state.ServicegroupmembersByServername) {
		add, remove := stringSliceDiff(stateMembersByName, planMembersByName)
		if len(remove) > 0 {
			removeServicegroupMemberBindings(r.client, servicegroupName, remove, true)
		}
		if len(add) > 0 {
			createServicegroupMemberBindings(r.client, servicegroupName, add, true)
		}
	}

	// state change - enable/disable action (SDK v2 parity).
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doServicegroupStateChange(&data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling servicegroup %s, got error: %s", servicegroupName, err))
			return
		}
	}

	// Unset attributes removed from config so the appliance reverts them to defaults.
	unsetIdPayload := map[string]interface{}{
		"servicegroupname": servicegroupName,
	}
	if err := utils.ExecuteUnset(r.client, service.Servicegroup.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset servicegroup attributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated servicegroup resource")

	// Read the updated state back.
	r.readServicegroupFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServicegroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServicegroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting servicegroup resource")

	servicegroupName := data.Id.ValueString()
	if err := r.client.DeleteResource(service.Servicegroup.Type(), servicegroupName); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete servicegroup %s, got error: %s", servicegroupName, err))
		return
	}

	tflog.Trace(ctx, "Deleted servicegroup resource")
}

// readServicegroupFromApi reads the servicegroup and its convenience bindings.
func (r *ServicegroupResource) readServicegroupFromApi(ctx context.Context, data *ServicegroupResourceModel, diags *diag.Diagnostics) {
	servicegroupName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Servicegroup.Type(), servicegroupName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			data.Id = types.StringNull()
			return
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read servicegroup, got error: %s", err))
		return
	}

	servicegroupSetAttrFromGet(ctx, data, getResponseData)

	// Convenience block: servicegroupmembers / servicegroupmembers_by_servername.
	if !data.Servicegroupmembers.IsNull() || !data.ServicegroupmembersByServername.IsNull() {
		boundMembers, err := r.client.FindAllBoundResources(service.Servicegroup.Type(), servicegroupName, "servicegroupmember")
		if err == nil {
			members := make([]string, 0, len(boundMembers))
			membersByName := make([]string, 0, len(boundMembers))
			for _, member := range boundMembers {
				ip, _ := member["ip"].(string)
				servername, _ := member["servername"].(string)
				port, _ := member["port"].(float64)
				weight, _ := member["weight"].(string)
				if servername == ip {
					members = append(members, fmt.Sprintf("%s:%.0f:%s", ip, port, weight))
				} else {
					membersByName = append(membersByName, fmt.Sprintf("%s:%.0f:%s", servername, port, weight))
				}
			}
			if !data.Servicegroupmembers.IsNull() {
				sv, d := types.SetValueFrom(ctx, types.StringType, members)
				diags.Append(d...)
				data.Servicegroupmembers = sv
			}
			if !data.ServicegroupmembersByServername.IsNull() {
				sv, d := types.SetValueFrom(ctx, types.StringType, membersByName)
				diags.Append(d...)
				data.ServicegroupmembersByServername = sv
			}
		}
	}

	// Convenience block: lbvservers.
	if !data.Lbvservers.IsNull() {
		vserverBindings, err := r.client.FindResourceArray(service.Servicegroupbindings.Type(), servicegroupName)
		if err == nil {
			lbvservers := make([]string, 0, len(vserverBindings))
			for _, vserver := range vserverBindings {
				if vs, ok := vserver["vservername"]; ok {
					lbvservers = append(lbvservers, vs.(string))
				}
			}
			sv, d := types.SetValueFrom(ctx, types.StringType, lbvservers)
			diags.Append(d...)
			data.Lbvservers = sv
		}
	}

	// Convenience block: lbmonitor.
	if !data.Lbmonitor.IsNull() {
		boundMonitors, err := r.client.FindAllBoundResources(service.Servicegroup.Type(), servicegroupName, service.Lbmonitor.Type())
		if err == nil {
			boundMonitor := ""
			for _, monitor := range boundMonitors {
				if mon, ok := monitor["monitor_name"]; ok {
					boundMonitor = mon.(string)
					break
				}
			}
			data.Lbmonitor = types.StringValue(boundMonitor)
		}
	}
}

// doServicegroupStateChange enables/disables the servicegroup via the NITRO action
// endpoints (SDK v2 doServicegroupStateChange parity).
func (r *ServicegroupResource) doServicegroupStateChange(data *ServicegroupResourceModel) error {
	serviceGroup := basic.Servicegroup{
		Servicegroupname: data.Servicegroupname.ValueString(),
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		serviceGroup.Servername = data.Servername.ValueString()
	}
	if !data.Port.IsNull() && !data.Port.IsUnknown() {
		serviceGroup.Port = utils.IntPtr(int(data.Port.ValueInt64()))
	}

	newstate := data.State.ValueString()
	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Servicegroup.Type(), serviceGroup, "enable")
	case "DISABLED":
		if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
			serviceGroup.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
		}
		if !data.Graceful.IsNull() && !data.Graceful.IsUnknown() {
			serviceGroup.Graceful = data.Graceful.ValueString()
		}
		return r.client.ActOnResource(service.Servicegroup.Type(), serviceGroup, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// servicegroupAttrsChanged reports whether any updateable (non-ForceNew, non-state,
// non-convenience) servicegroup attribute changed between prior state and plan.
func servicegroupAttrsChanged(data, state *ServicegroupResourceModel) bool {
	return !data.Appflowlog.Equal(state.Appflowlog) ||
		!data.Autodelayedtrofs.Equal(state.Autodelayedtrofs) ||
		!data.Autodisabledelay.Equal(state.Autodisabledelay) ||
		!data.Autodisablegraceful.Equal(state.Autodisablegraceful) ||
		!data.Cacheable.Equal(state.Cacheable) ||
		!data.Cip.Equal(state.Cip) ||
		!data.Cipheader.Equal(state.Cipheader) ||
		!data.Cka.Equal(state.Cka) ||
		!data.Clttimeout.Equal(state.Clttimeout) ||
		!data.Cmp.Equal(state.Cmp) ||
		!data.Comment.Equal(state.Comment) ||
		!data.Customserverid.Equal(state.Customserverid) ||
		!data.Dbsttl.Equal(state.Dbsttl) ||
		!data.Downstateflush.Equal(state.Downstateflush) ||
		!data.Dupweight.Equal(state.Dupweight) ||
		!data.Hashid.Equal(state.Hashid) ||
		!data.Healthmonitor.Equal(state.Healthmonitor) ||
		!data.Httpprofilename.Equal(state.Httpprofilename) ||
		!data.Maxbandwidth.Equal(state.Maxbandwidth) ||
		!data.Maxclient.Equal(state.Maxclient) ||
		!data.Maxreq.Equal(state.Maxreq) ||
		!data.Monconnectionclose.Equal(state.Monconnectionclose) ||
		!data.Monitornamesvc.Equal(state.Monitornamesvc) ||
		!data.Monthreshold.Equal(state.Monthreshold) ||
		!data.Nameserver.Equal(state.Nameserver) ||
		!data.Netprofile.Equal(state.Netprofile) ||
		!data.Pathmonitor.Equal(state.Pathmonitor) ||
		!data.Pathmonitorindv.Equal(state.Pathmonitorindv) ||
		!data.Port.Equal(state.Port) ||
		!data.Quicprofilename.Equal(state.Quicprofilename) ||
		!data.Rtspsessionidremap.Equal(state.Rtspsessionidremap) ||
		!data.Serverid.Equal(state.Serverid) ||
		!data.Servername.Equal(state.Servername) ||
		!data.Sp.Equal(state.Sp) ||
		!data.Svrtimeout.Equal(state.Svrtimeout) ||
		!data.Tcpb.Equal(state.Tcpb) ||
		!data.Tcpprofilename.Equal(state.Tcpprofilename) ||
		!data.Useproxyport.Equal(state.Useproxyport) ||
		!data.Usip.Equal(state.Usip) ||
		!data.Weight.Equal(state.Weight)
}

// setToStringSlice converts a types.Set (of strings) to a []string. Null/unknown
// sets yield an empty slice.
func setToStringSlice(ctx context.Context, set types.Set, diags *diag.Diagnostics) []string {
	if set.IsNull() || set.IsUnknown() {
		return []string{}
	}
	out := make([]string, 0, len(set.Elements()))
	diags.Append(set.ElementsAs(ctx, &out, false)...)
	return out
}

// stringSliceDiff returns the elements to add (in newList, not oldList) and remove
// (in oldList, not newList).
func stringSliceDiff(oldList, newList []string) (add, remove []string) {
	oldSet := make(map[string]bool, len(oldList))
	for _, s := range oldList {
		oldSet[s] = true
	}
	newSet := make(map[string]bool, len(newList))
	for _, s := range newList {
		newSet[s] = true
	}
	for _, s := range newList {
		if !oldSet[s] {
			add = append(add, s)
		}
	}
	for _, s := range oldList {
		if !newSet[s] {
			remove = append(remove, s)
		}
	}
	return add, remove
}

// createServicegroupMemberBindings binds members in the form ip:port:weight (or
// servername:port:weight when bindByServername is true). SDK v2 parity.
func createServicegroupMemberBindings(client *service.NitroClient, servicegroupName string, groupmembers []string, bindByServername bool) error {
	for _, member := range groupmembers {
		parts := strings.Split(member, ":")
		var ip, servername string
		if !bindByServername {
			ip = parts[0]
		} else {
			servername = parts[0]
		}
		if len(parts) < 2 {
			continue
		}
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		weightFound := false
		var weight int
		if len(parts) > 2 {
			weight, err = strconv.Atoi(parts[2])
			weightFound = err == nil
		}
		var binding basic.Servicegroupservicegroupmemberbinding
		if !bindByServername {
			binding = basic.Servicegroupservicegroupmemberbinding{
				Servicegroupname: servicegroupName,
				Ip:               ip,
				Port:             utils.IntPtr(port),
			}
		} else {
			binding = basic.Servicegroupservicegroupmemberbinding{
				Servicegroupname: servicegroupName,
				Servername:       servername,
				Port:             utils.IntPtr(port),
			}
		}
		if weightFound {
			binding.Weight = utils.IntPtr(weight)
		}
		if _, err := client.AddResource(service.Servicegroup_servicegroupmember_binding.Type(), servicegroupName, &binding); err != nil {
			continue
		}
	}
	return nil
}

// removeServicegroupMemberBindings unbinds members. SDK v2 parity.
func removeServicegroupMemberBindings(client *service.NitroClient, servicegroupName string, groupmembers []string, bindByServername bool) error {
	for _, member := range groupmembers {
		parts := strings.Split(member, ":")
		var ip, servername, port string
		if !bindByServername {
			ip = parts[0]
		} else {
			servername = parts[0]
		}
		if len(parts) < 2 {
			continue
		}
		port = parts[1]
		args := make([]string, 1)
		if !bindByServername {
			args[0] = fmt.Sprintf("ip:%s,port:%s", ip, port)
		} else {
			args[0] = fmt.Sprintf("servername:%s,port:%s", servername, port)
		}
		if err := client.DeleteResourceWithArgs(service.Servicegroup_servicegroupmember_binding.Type(), servicegroupName, args); err != nil {
			continue
		}
	}
	return nil
}

// addLbvserverBindings binds the servicegroup to the given lb vservers. SDK v2 parity.
func addLbvserverBindings(client *service.NitroClient, servicegroupName string, lbvservers []string) error {
	for _, lbvserverName := range lbvservers {
		binding := lb.Lbvserverservicegroupbinding{
			Name:             lbvserverName,
			Servicegroupname: servicegroupName,
		}
		if err := client.BindResource(service.Lbvserver.Type(), lbvserverName, service.Servicegroup.Type(), servicegroupName, &binding); err != nil {
			return fmt.Errorf("failed to bind servicegroup %s to lbvserver %s", servicegroupName, lbvserverName)
		}
	}
	return nil
}

// removeLbvserverBindings unbinds the servicegroup from the given lb vservers. SDK v2 parity.
func removeLbvserverBindings(client *service.NitroClient, servicegroupName string, lbvservers []string) error {
	for _, lbvserverName := range lbvservers {
		if err := client.UnbindResource(service.Lbvserver.Type(), lbvserverName, service.Servicegroup.Type(), servicegroupName, "servicegroupname"); err != nil {
			return fmt.Errorf("error unbinding lbvserver %s from servicegroup %s", lbvserverName, servicegroupName)
		}
	}
	return nil
}
