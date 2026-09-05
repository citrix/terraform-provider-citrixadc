package cluster

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	adccluster "github.com/citrix/adc-nitro-go/resource/config/cluster"
	"github.com/citrix/adc-nitro-go/resource/config/ns"
	"github.com/citrix/adc-nitro-go/resource/config/router"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// cluster is a high-level bootstrap orchestrator, migrated 1:1 from the SDK v2
// "citrixadc_cluster" resource. It is NOT a thin wrapper over the NITRO cluster
// "join" action: it creates the cluster instance on the first (CCO) node, adds
// the CLIP, enables/reboots the node, then joins the remaining nodes to the
// cluster. The resource identity is the cluster instance id (clid), matching the
// SDK v2 id format and resource_id_mapping.json ("cluster": "clid").
//
// WARNING: applying this resource bootstraps/modifies a Citrix ADC cluster,
// which is disruptive (reboots nodes). Intended for deliberate operator use.
var _ resource.Resource = &ClusterResource{}
var _ resource.ResourceWithConfigure = (*ClusterResource)(nil)
var _ resource.ResourceWithImportState = (*ClusterResource)(nil)

func NewClusterResource() resource.Resource {
	return &ClusterResource{}
}

// ClusterResource defines the resource implementation.
type ClusterResource struct {
	client *service.NitroClient
}

func (r *ClusterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ClusterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster"
}

func (r *ClusterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// -------------------------------------------------------------------------
// CRUD
// -------------------------------------------------------------------------

func (r *ClusterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ClusterResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nodes := r.getClusternodes(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(nodes) == 0 {
		resp.Diagnostics.AddError("Configuration Error", "at least one clusternode block is required")
		return
	}

	tflog.Debug(ctx, "Bootstrapping cluster")
	if err := r.bootstrapCluster(ctx, &data, nodes); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to bootstrap cluster, got error: %s", err))
		return
	}

	// ID is the cluster instance id (clid), matching SDK v2 d.SetId(clid).
	data.Id = types.StringValue(strconv.Itoa(int(data.Clid.ValueInt64())))

	// Read back the cluster instance to resolve computed attributes.
	if _, err := r.readClusterInstanceIntoModel(ctx, &data); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("cluster created but read-back failed: %s", err))
		r.nullifyUnknownComputed(&data)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ClusterResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading cluster resource")

	found, err := r.readClusterInstanceIntoModel(ctx, &data)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cluster, got error: %s", err))
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state ClusterResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state.
	data.Id = state.Id

	oldNodes := r.getClusternodesFrom(ctx, state.Clusternode, &resp.Diagnostics)
	newNodes := r.getClusternodes(ctx, &data, &resp.Diagnostics)
	oldGroups := r.getClusternodegroupsFrom(ctx, state.Clusternodegroup, &resp.Diagnostics)
	newGroups := r.getClusternodegroups(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(newNodes) == 0 {
		resp.Diagnostics.AddError("Configuration Error", "at least one clusternode block is required")
		return
	}

	tflog.Debug(ctx, "Updating cluster resource")

	clid := int(data.Clid.ValueInt64())
	clidStr := strconv.Itoa(clid)

	// --- Update cluster instance attributes ---
	clusterinstance := adccluster.Clusterinstance{Clid: utils.IntPtr(clid)}
	hasChange := false
	// Change detection: only treat an attribute as changed when the planned value
	// is KNOWN, NON-NULL and differs from state. Two cases would otherwise make the
	// clusterinstance payload collapse to a clid-only set, which NITRO rejects with
	// errorcode 1094 "Too few arguments":
	//   1) An Optional+Computed attribute not in config is unknown in the plan
	//      (guarded by !IsUnknown).
	//   2) On an SDK v2 -> Framework upgrade, an Optional-only attribute (e.g.
	//      nodegroup, which cannot carry UseStateForUnknown) is null in the Framework
	//      config but was stored as "" by SDK v2, so null != "" spuriously reports a
	//      change and writes strOrEmpty(null)="" (guarded by !IsNull).
	// This matches SDK v2's d.HasChange semantics and the payload-builder convention
	// (guard !IsNull() && !IsUnknown()).
	if !data.Backplanebasedview.IsNull() && !data.Backplanebasedview.IsUnknown() && !data.Backplanebasedview.Equal(state.Backplanebasedview) {
		clusterinstance.Backplanebasedview = strOrEmpty(data.Backplanebasedview)
		hasChange = true
	}
	if !data.Deadinterval.IsNull() && !data.Deadinterval.IsUnknown() && !data.Deadinterval.Equal(state.Deadinterval) {
		clusterinstance.Deadinterval = utils.IntPtr(intOrZero(data.Deadinterval))
		hasChange = true
	}
	if !data.Hellointerval.IsNull() && !data.Hellointerval.IsUnknown() && !data.Hellointerval.Equal(state.Hellointerval) {
		clusterinstance.Hellointerval = utils.IntPtr(intOrZero(data.Hellointerval))
		hasChange = true
	}
	if !data.Inc.IsNull() && !data.Inc.IsUnknown() && !data.Inc.Equal(state.Inc) {
		clusterinstance.Inc = strOrEmpty(data.Inc)
		hasChange = true
	}
	if !data.Nodegroup.IsNull() && !data.Nodegroup.IsUnknown() && !data.Nodegroup.Equal(state.Nodegroup) {
		clusterinstance.Nodegroup = strOrEmpty(data.Nodegroup)
		hasChange = true
	}
	if !data.Preemption.IsNull() && !data.Preemption.IsUnknown() && !data.Preemption.Equal(state.Preemption) {
		clusterinstance.Preemption = strOrEmpty(data.Preemption)
		hasChange = true
	}
	if !data.Processlocal.IsNull() && !data.Processlocal.IsUnknown() && !data.Processlocal.Equal(state.Processlocal) {
		clusterinstance.Processlocal = strOrEmpty(data.Processlocal)
		hasChange = true
	}
	if !data.Quorumtype.IsNull() && !data.Quorumtype.IsUnknown() && !data.Quorumtype.Equal(state.Quorumtype) {
		clusterinstance.Quorumtype = strOrEmpty(data.Quorumtype)
		hasChange = true
	}
	if !data.Retainconnectionsoncluster.IsNull() && !data.Retainconnectionsoncluster.IsUnknown() && !data.Retainconnectionsoncluster.Equal(state.Retainconnectionsoncluster) {
		clusterinstance.Retainconnectionsoncluster = strOrEmpty(data.Retainconnectionsoncluster)
		hasChange = true
	}

	if hasChange {
		if _, err := r.client.UpdateResource(service.Clusterinstance.Type(), clidStr, &clusterinstance); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error updating clusterinstance %s: %s", clidStr, err))
			return
		}
	}

	clusterNodegroupChanged := !clusternodegroupsEqual(oldGroups, newGroups)
	l3 := isClusterModeL3(&data, newGroups)

	// Add and update node groups before nodes.
	if l3 && clusterNodegroupChanged {
		if err := r.addClusterNodegroups(ctx, newGroups); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error adding cluster nodegroups: %s", err))
			return
		}
		if err := r.updateClusterNodegroups(ctx, newGroups); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error updating cluster nodegroups: %s", err))
			return
		}
	}

	if err := r.updateClusterNodes(ctx, &data, oldNodes, newNodes); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error updating cluster nodes: %s", err))
		return
	}

	// Delete node groups after nodes.
	if l3 && clusterNodegroupChanged {
		if err := r.deleteClusterNodegroups(ctx, oldGroups, newGroups); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error deleting cluster nodegroups: %s", err))
			return
		}
	}

	if _, err := r.readClusterInstanceIntoModel(ctx, &data); err != nil {
		tflog.Warn(ctx, fmt.Sprintf("cluster updated but read-back failed: %s", err))
		r.nullifyUnknownComputed(&data)
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ClusterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ClusterResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting cluster resource")

	nodes := r.getClusternodes(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if len(nodes) == 0 {
		return
	}

	nodeids := getSortedClusternodeIds(nodes)
	// Delete member nodes first (excluding the CCO, which is index 0).
	for _, nodeid := range nodeids[1:] {
		if node, ok := getClusterNodeByid(nodes, nodeid); ok {
			if err := r.deleteSingleClusterNode(ctx, &data, node, true); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cluster node %d: %s", nodeid, err))
				return
			}
		}
	}

	// Delete the CCO last. Do not wait for CLIP migration on the last node.
	if node, ok := getClusterNodeByid(nodes, nodeids[0]); ok {
		if err := r.deleteSingleClusterNode(ctx, &data, node, false); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete cluster CCO node %d: %s", nodeids[0], err))
			return
		}
	}
}

// -------------------------------------------------------------------------
// Read-back helpers
// -------------------------------------------------------------------------

// readClusterInstanceIntoModel resolves computed cluster-instance attributes
// from the ADC. To stay compatible with SDK v2 while satisfying the stricter
// Framework consistency checks, a user-configured (known) value is never
// overwritten by the API; only unknown/null computed attributes are filled.
func (r *ClusterResource) readClusterInstanceIntoModel(ctx context.Context, data *ClusterResourceModel) (bool, error) {
	datalist, err := r.client.FindAllResources(service.Clusterinstance.Type())
	if err != nil {
		return false, err
	}
	if len(datalist) == 0 {
		return false, nil
	}
	d := datalist[0]

	fillInt := func(cur types.Int64, key string) types.Int64 {
		if !cur.IsNull() && !cur.IsUnknown() {
			return cur
		}
		if v, ok := d[key]; ok {
			if iv, e := utils.ConvertToInt64(v); e == nil {
				return types.Int64Value(iv)
			}
		}
		return types.Int64Null()
	}
	fillStr := func(cur types.String, key string) types.String {
		if !cur.IsNull() && !cur.IsUnknown() {
			return cur
		}
		if v, ok := d[key]; ok {
			if s, ok2 := v.(string); ok2 {
				return types.StringValue(s)
			}
		}
		return types.StringNull()
	}

	data.Clid = fillInt(data.Clid, "clid")
	data.Deadinterval = fillInt(data.Deadinterval, "deadinterval")
	data.Hellointerval = fillInt(data.Hellointerval, "hellointerval")
	data.Backplanebasedview = fillStr(data.Backplanebasedview, "backplanebasedview")
	data.Inc = fillStr(data.Inc, "inc")
	data.Preemption = fillStr(data.Preemption, "preemption")
	data.Processlocal = fillStr(data.Processlocal, "processlocal")
	data.Quorumtype = fillStr(data.Quorumtype, "quorumtype")
	data.Retainconnectionsoncluster = fillStr(data.Retainconnectionsoncluster, "retainconnectionsoncluster")
	// nodegroup (top-level) is Optional-only; preserve config to avoid perpetual diffs.

	return true, nil
}

// nullifyUnknownComputed ensures no computed attribute is left unknown when a
// read-back fails, preventing "inconsistent result after apply" errors.
func (r *ClusterResource) nullifyUnknownComputed(data *ClusterResourceModel) {
	if data.Backplanebasedview.IsUnknown() {
		data.Backplanebasedview = types.StringNull()
	}
	if data.Deadinterval.IsUnknown() {
		data.Deadinterval = types.Int64Null()
	}
	if data.Hellointerval.IsUnknown() {
		data.Hellointerval = types.Int64Null()
	}
	if data.Inc.IsUnknown() {
		data.Inc = types.StringNull()
	}
	if data.Preemption.IsUnknown() {
		data.Preemption = types.StringNull()
	}
	if data.Processlocal.IsUnknown() {
		data.Processlocal = types.StringNull()
	}
	if data.Quorumtype.IsUnknown() {
		data.Quorumtype = types.StringNull()
	}
	if data.Retainconnectionsoncluster.IsUnknown() {
		data.Retainconnectionsoncluster = types.StringNull()
	}
}

// -------------------------------------------------------------------------
// Bootstrap
// -------------------------------------------------------------------------

func (r *ClusterResource) bootstrapCluster(ctx context.Context, data *ClusterResourceModel, nodes []ClusternodeModel) error {
	groups := r.getClusternodegroupsFrom(ctx, data.Clusternodegroup, nil)

	if err := r.createFirstClusterNode(ctx, data, nodes, groups); err != nil {
		return err
	}

	if isClusterModeL3(data, groups) {
		if err := r.addClusterNodegroups(ctx, groups); err != nil {
			return err
		}
	}

	// Join the rest of the nodes to the cluster.
	nodeids := getSortedClusternodeIds(nodes)
	for _, nodeid := range nodeids[1:] {
		node, ok := getClusterNodeByid(nodes, nodeid)
		if !ok {
			continue
		}
		if err := r.addSingleClusterNode(ctx, data, node); err != nil {
			return err
		}
	}

	return nil
}

func (r *ClusterResource) createFirstClusterNode(ctx context.Context, data *ClusterResourceModel, nodes []ClusternodeModel, groups []ClusternodegroupModel) error {
	nodeids := getSortedClusternodeIds(nodes)
	firstNode, ok := getClusterNodeByid(nodes, nodeids[0])
	if !ok {
		return fmt.Errorf("could not resolve first cluster node")
	}

	nodeClient, err := r.instantiateNodeClient(firstNode)
	if err != nil {
		return err
	}

	clid := int(data.Clid.ValueInt64())
	clusterId := strconv.Itoa(clid)

	clusterinstance := adccluster.Clusterinstance{
		Backplanebasedview:         strOrEmpty(data.Backplanebasedview),
		Inc:                        strOrEmpty(data.Inc),
		Nodegroup:                  strOrEmpty(data.Nodegroup),
		Preemption:                 strOrEmpty(data.Preemption),
		Processlocal:               strOrEmpty(data.Processlocal),
		Quorumtype:                 strOrEmpty(data.Quorumtype),
		Retainconnectionsoncluster: strOrEmpty(data.Retainconnectionsoncluster),
		Clid:                       utils.IntPtr(clid),
	}
	if !data.Deadinterval.IsNull() && !data.Deadinterval.IsUnknown() {
		clusterinstance.Deadinterval = utils.IntPtr(intOrZero(data.Deadinterval))
	}
	if !data.Hellointerval.IsNull() && !data.Hellointerval.IsUnknown() {
		clusterinstance.Hellointerval = utils.IntPtr(intOrZero(data.Hellointerval))
	}

	if _, err = nodeClient.AddResource(service.Clusterinstance.Type(), clusterId, &clusterinstance); err != nil {
		return err
	}

	// In L3 mode create the node group prior to adding the node.
	if isClusterModeL3(data, groups) {
		nodegroupName := strOrEmpty(firstNode.Nodegroup)
		ng, found := getClusterNodegroupByName(groups, nodegroupName)
		if !found {
			return fmt.Errorf("cannot find node group %s in configuration", nodegroupName)
		}
		clusternodegroup := clusternodegroupFromModel(ng)
		if _, err := nodeClient.AddResource(service.Clusternodegroup.Type(), clusternodegroup.Name, &clusternodegroup); err != nil {
			return err
		}
	}

	// Add the first cluster node.
	clusternode := clusternodeFromModel(firstNode)
	if _, err = nodeClient.AddResource(service.Clusternode.Type(), strconv.Itoa(intOrZero(firstNode.Nodeid)), &clusternode); err != nil {
		return err
	}

	// Add the CLIP to the first node.
	clip := strOrEmpty(data.Clip)
	nsip := ns.Nsip{
		Ipaddress: clip,
		Netmask:   "255.255.255.255",
		Type:      "CLIP",
	}
	if _, err = nodeClient.AddResource(service.Nsip.Type(), clip, &nsip); err != nil {
		return err
	}

	// Enable the cluster instance on the first node.
	clusterinstanceEnabler := adccluster.Clusterinstance{Clid: utils.IntPtr(clid)}
	if err = nodeClient.ActOnResource(service.Clusterinstance.Type(), &clusterinstanceEnabler, "enable"); err != nil {
		return err
	}

	// Save config and reboot the first node.
	nodeClient.SaveConfig()
	tflog.Debug(ctx, "Rebooting first cluster node")
	if err = nodeClient.ActOnResource("reboot", &ns.Reboot{Warm: true}, ""); err != nil {
		return err
	}

	// Poll the CLIP for bootstrap completion.
	delay, interval, total, perPoll, err := parseTimeouts(
		data.BootstrapPollDelay, data.BootstrapPollInterval, data.BootstrapTotalTimeout, data.BootstrapPollTimeout)
	if err != nil {
		return err
	}
	if err = r.pollClipReachable(delay, interval, total, perPoll); err != nil {
		return err
	}

	// Verify that the first node is actually part of the cluster.
	nodeData, err := r.client.FindAllResources(service.Clusternode.Type())
	if err != nil {
		return err
	}
	if len(nodeData) == 0 {
		return fmt.Errorf("no cluster nodes found after bootstrap; the CLIP may not be fully ready yet")
	}
	fetchedIpaddress := nodeData[0]["ipaddress"]
	configIpaddress := strOrEmpty(firstNode.Ipaddress)
	if fetchedIpaddress != configIpaddress {
		return fmt.Errorf("fetched first node address differs from configuration. Fetched: %v Configuration: %s", fetchedIpaddress, configIpaddress)
	}

	return nil
}

// -------------------------------------------------------------------------
// Node group management
// -------------------------------------------------------------------------

func (r *ClusterResource) addClusterNodegroups(ctx context.Context, groups []ClusternodegroupModel) error {
	for _, ng := range groups {
		name := strOrEmpty(ng.Name)
		exists, err := clusternodegroupExistsInCluster(r.client, name)
		if err != nil {
			return err
		}
		if !exists {
			if err := r.addSingleClusterNodegroup(ctx, ng); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ClusterResource) addSingleClusterNodegroup(ctx context.Context, ng ClusternodegroupModel) error {
	clusternodegroup := clusternodegroupFromModel(ng)
	_, err := r.client.AddResource(service.Clusternodegroup.Type(), clusternodegroup.Name, &clusternodegroup)
	return err
}

func (r *ClusterResource) updateClusterNodegroups(ctx context.Context, groups []ClusternodegroupModel) error {
	for _, ng := range groups {
		name := strOrEmpty(ng.Name)
		exists, err := clusternodegroupExistsInCluster(r.client, name)
		if err != nil {
			return err
		}
		if !exists {
			continue
		}
		fetched, err := getClusternodegroupFromCluster(r.client, name)
		if err != nil {
			return err
		}
		needsUpdate := false
		if p := intOrZero(ng.Priority); p != 0 {
			if fv, ok := fetched["priority"]; !ok || fmt.Sprintf("%v", fv) != strconv.Itoa(p) {
				needsUpdate = true
			}
		}
		for key, cfg := range map[string]string{
			"state":  strOrEmpty(ng.State),
			"sticky": strOrEmpty(ng.Sticky),
			"strict": strOrEmpty(ng.Strict),
		} {
			if cfg != "" {
				if fv, ok := fetched[key]; !ok || fmt.Sprintf("%v", fv) != cfg {
					needsUpdate = true
				}
			}
		}
		if needsUpdate {
			clusternodegroup := clusternodegroupFromModel(ng)
			if err := r.client.UpdateUnnamedResource(service.Clusternodegroup.Type(), &clusternodegroup); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *ClusterResource) deleteClusterNodegroups(ctx context.Context, oldGroups, newGroups []ClusternodegroupModel) error {
	nameInSet := func(groups []ClusternodegroupModel, name string) bool {
		for _, g := range groups {
			if strOrEmpty(g.Name) == name {
				return true
			}
		}
		return false
	}
	for _, g := range oldGroups {
		name := strOrEmpty(g.Name)
		if name == "DEFAULT_NG" {
			continue
		}
		if !nameInSet(newGroups, name) {
			if err := r.client.DeleteResource(service.Clusternodegroup.Type(), name); err != nil {
				return err
			}
		}
	}
	return nil
}

// -------------------------------------------------------------------------
// Node management
// -------------------------------------------------------------------------

func (r *ClusterResource) updateClusterNodes(ctx context.Context, data *ClusterResourceModel, oldNodes, newNodes []ClusternodeModel) error {
	removed := nodeSetDifference(oldNodes, newNodes)
	added := nodeSetDifference(newNodes, oldNodes)

	var toRemove []ClusternodeModel
	var toCreate []ClusternodeModel

	// Inline updates and recreates for nodes with matching nodeid.
	for _, oldVal := range removed {
		for _, newVal := range added {
			if compareNeedsUpdate(oldVal, newVal) {
				tflog.Debug(ctx, fmt.Sprintf("Updating node %d", intOrZero(oldVal.Nodeid)))
				if err := r.updateSingleClusterNode(ctx, newVal); err != nil {
					return err
				}
				break
			}
			if compareNeedsRecreate(oldVal, newVal) {
				toRemove = append(toRemove, oldVal)
				toCreate = append(toCreate, newVal)
				break
			}
		}
	}

	// Recreates: remove all first (so node swaps work), then create.
	for _, v := range toRemove {
		if err := r.deleteSingleClusterNode(ctx, data, v, true); err != nil {
			return err
		}
	}
	for _, v := range toCreate {
		if err := r.addSingleClusterNode(ctx, data, v); err != nil {
			return err
		}
	}

	// Create brand-new nodes (nodeid not present in old set).
	for _, newVal := range added {
		needsCreate := true
		for _, oldVal := range removed {
			if intOrZero(oldVal.Nodeid) == intOrZero(newVal.Nodeid) {
				needsCreate = false
				break
			}
		}
		if needsCreate {
			tflog.Debug(ctx, fmt.Sprintf("Creating node %d", intOrZero(newVal.Nodeid)))
			if err := r.addSingleClusterNode(ctx, data, newVal); err != nil {
				return err
			}
		}
	}

	// Delete old nodes (nodeid not present in new set).
	for _, oldVal := range removed {
		needsDelete := true
		for _, newVal := range added {
			if intOrZero(oldVal.Nodeid) == intOrZero(newVal.Nodeid) {
				needsDelete = false
				break
			}
		}
		if needsDelete {
			if err := r.deleteSingleClusterNode(ctx, data, oldVal, true); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ClusterResource) addSingleClusterNode(ctx context.Context, data *ClusterResourceModel, node ClusternodeModel) error {
	// Add the cluster node at the CLIP.
	clusternode := clusternodeFromModel(node)
	if _, err := r.client.AddResource(service.Clusternode.Type(), strconv.Itoa(intOrZero(node.Nodeid)), &clusternode); err != nil {
		return err
	}

	// Register the node SNIP on the CLIP if requested.
	if boolOrFalse(node.Addsnip) {
		nodeNsip := ns.Nsip{
			Ipaddress:  strOrEmpty(node.SnipIpaddress),
			Mgmtaccess: "ENABLED",
			Netmask:    strOrEmpty(node.SnipNetmask),
			Type:       "SNIP",
		}
		if _, err := r.client.AddResource(service.Nsip.Type(), strOrEmpty(node.Ipaddress), &nodeNsip); err != nil {
			return err
		}
	}

	// Apply optional VTYSH commands.
	if boolOrFalse(node.VtyshEnable) {
		var cmds []string
		if !node.Vtysh.IsNull() && !node.Vtysh.IsUnknown() {
			node.Vtysh.ElementsAs(ctx, &cmds, false)
		}
		for _, cmd := range cmds {
			routerdynamicrouting := router.Routerdynamicrouting{Commandstring: cmd}
			if err := r.client.ActOnResource(service.Routerdynamicrouting.Type(), &routerdynamicrouting, "apply"); err != nil {
				return err
			}
		}
	}

	// Instantiate a node client and join the cluster from the node.
	nodeClient, err := r.instantiateNodeClient(node)
	if err != nil {
		return err
	}
	joinPayload := adccluster.Cluster{
		Clip:     strOrEmpty(data.Clip),
		Password: r.client.GetPassword(),
	}
	if err := nodeClient.ActOnResource(service.Cluster.Type(), &joinPayload, "join"); err != nil {
		return err
	}

	// Save config and reboot the node.
	nodeClient.SaveConfig()
	tflog.Debug(ctx, fmt.Sprintf("Rebooting node %d", intOrZero(node.Nodeid)))
	if err := nodeClient.ActOnResource("reboot", &ns.Reboot{Warm: true}, ""); err != nil {
		return err
	}

	// Poll the CLIP until the added node id becomes ACTIVE.
	delay, interval, total, err := parseThreeTimeouts(
		data.NodeAddPollDelay, data.NodeAddPollInterval, data.NodeAddTotalTimeout)
	if err != nil {
		return err
	}
	nodeid := intOrZero(node.Nodeid)
	return poll(delay, interval, total, func() (bool, error) {
		return r.pollClusterNodeWithid(nodeid)
	})
}

func (r *ClusterResource) deleteSingleClusterNode(ctx context.Context, data *ClusterResourceModel, node ClusternodeModel, waitClip bool) error {
	nodeId := strconv.Itoa(intOrZero(node.Nodeid))
	// Deleting the CCO node produces an expected connection reset.
	if err := r.client.DeleteResource(service.Clusternode.Type(), nodeId); err != nil {
		if !strings.Contains(err.Error(), "read: connection reset by peer") {
			return err
		}
		tflog.Debug(ctx, fmt.Sprintf("lost CLIP when deleting node %d", intOrZero(node.Nodeid)))
	}

	if waitClip {
		delay, interval, total, perPoll, err := parseTimeouts(
			data.ClipMigrationPollDelay, data.ClipMigrationPollInterval, data.ClipMigrationTotalTimeout, data.ClipMigrationPollTimeout)
		if err != nil {
			return err
		}
		if err := r.pollClipReachable(delay, interval, total, perPoll); err != nil {
			return err
		}
	}
	return nil
}

func (r *ClusterResource) updateSingleClusterNode(ctx context.Context, node ClusternodeModel) error {
	// Only include attributes valid in an HTTP PUT.
	clusternode := adccluster.Clusternode{
		Backplane:  strOrEmpty(node.Backplane),
		Delay:      utils.IntPtr(intOrZero(node.Delay)),
		Nodeid:     utils.IntPtr(intOrZero(node.Nodeid)),
		Priority:   utils.IntPtr(intOrZero(node.Priority)),
		State:      strOrEmpty(node.State),
		Tunnelmode: strOrEmpty(node.Tunnelmode),
	}
	return r.client.UpdateUnnamedResource(service.Clusternode.Type(), &clusternode)
}

// -------------------------------------------------------------------------
// Node comparison helpers (ported from SDK v2)
// -------------------------------------------------------------------------

func compareNeedsUpdate(oldNode, newNode ClusternodeModel) bool {
	if intOrZero(oldNode.Nodeid) != intOrZero(newNode.Nodeid) {
		return false
	}
	// Non-updateable attributes changing means recreate, not update.
	if strOrEmpty(oldNode.Ipaddress) != strOrEmpty(newNode.Ipaddress) {
		return false
	}
	if strOrEmpty(oldNode.Nodegroup) != strOrEmpty(newNode.Nodegroup) {
		return false
	}

	needsUpdate := false
	if strOrEmpty(oldNode.Backplane) != strOrEmpty(newNode.Backplane) {
		needsUpdate = true
	}
	if intOrZero(oldNode.Delay) != intOrZero(newNode.Delay) {
		needsUpdate = true
	}
	if intOrZero(oldNode.Priority) != intOrZero(newNode.Priority) {
		needsUpdate = true
	}
	if strOrEmpty(oldNode.State) != strOrEmpty(newNode.State) {
		needsUpdate = true
	}
	if strOrEmpty(oldNode.Tunnelmode) != strOrEmpty(newNode.Tunnelmode) {
		needsUpdate = true
	}
	return needsUpdate
}

func compareNeedsRecreate(oldNode, newNode ClusternodeModel) bool {
	if intOrZero(oldNode.Nodeid) != intOrZero(newNode.Nodeid) {
		return false
	}
	needsRecreate := false
	if strOrEmpty(oldNode.Ipaddress) != strOrEmpty(newNode.Ipaddress) {
		needsRecreate = true
	}
	if strOrEmpty(oldNode.Nodegroup) != strOrEmpty(newNode.Nodegroup) {
		needsRecreate = true
	}
	return needsRecreate
}

// nodeSetDifference returns the nodes in a that are not in b, using the same
// identity fields SDK v2 hashed for its clusternode TypeSet (excludes the
// config-only fields such as endpoint/username/password/vtysh).
func nodeSetDifference(a, b []ClusternodeModel) []ClusternodeModel {
	var out []ClusternodeModel
	for _, x := range a {
		found := false
		for _, y := range b {
			if nodeHashEqual(x, y) {
				found = true
				break
			}
		}
		if !found {
			out = append(out, x)
		}
	}
	return out
}

func nodeHashEqual(a, b ClusternodeModel) bool {
	return strOrEmpty(a.Backplane) == strOrEmpty(b.Backplane) &&
		strOrEmpty(a.Clearnodegroupconfig) == strOrEmpty(b.Clearnodegroupconfig) &&
		intOrZero(a.Delay) == intOrZero(b.Delay) &&
		strOrEmpty(a.Ipaddress) == strOrEmpty(b.Ipaddress) &&
		strOrEmpty(a.Nodegroup) == strOrEmpty(b.Nodegroup) &&
		intOrZero(a.Nodeid) == intOrZero(b.Nodeid) &&
		intOrZero(a.Priority) == intOrZero(b.Priority) &&
		strOrEmpty(a.State) == strOrEmpty(b.State) &&
		strOrEmpty(a.Tunnelmode) == strOrEmpty(b.Tunnelmode)
}

func clusternodegroupsEqual(a, b []ClusternodegroupModel) bool {
	if len(a) != len(b) {
		return false
	}
	match := func(x ClusternodegroupModel, set []ClusternodegroupModel) bool {
		for _, y := range set {
			if strOrEmpty(x.Name) == strOrEmpty(y.Name) &&
				intOrZero(x.Priority) == intOrZero(y.Priority) &&
				strOrEmpty(x.State) == strOrEmpty(y.State) &&
				strOrEmpty(x.Sticky) == strOrEmpty(y.Sticky) &&
				strOrEmpty(x.Strict) == strOrEmpty(y.Strict) {
				return true
			}
		}
		return false
	}
	for _, x := range a {
		if !match(x, b) {
			return false
		}
	}
	return true
}

// -------------------------------------------------------------------------
// NITRO node group lookups
// -------------------------------------------------------------------------

func clusternodegroupExistsInCluster(client *service.NitroClient, name string) (bool, error) {
	data, err := getClusternodegroupsByName(client, name)
	if err != nil {
		return false, err
	}
	switch len(data) {
	case 0:
		return false, nil
	case 1:
		return true, nil
	default:
		return false, fmt.Errorf("got multiple node groups existing for name %s (%v)", name, data)
	}
}

func getClusternodegroupFromCluster(client *service.NitroClient, name string) (map[string]interface{}, error) {
	data, err := getClusternodegroupsByName(client, name)
	if err != nil {
		return nil, err
	}
	switch len(data) {
	case 0:
		return nil, nil
	case 1:
		return data[0], nil
	default:
		return nil, fmt.Errorf("got multiple node groups existing for name %s (%v)", name, data)
	}
}

func getClusternodegroupsByName(client *service.NitroClient, name string) ([]map[string]interface{}, error) {
	findParams := service.FindParams{
		ResourceType:             service.Clusternodegroup.Type(),
		ResourceName:             name,
		ResourceMissingErrorCode: 258,
	}
	return client.FindResourceArrayWithParams(findParams)
}

// -------------------------------------------------------------------------
// Node client + polling
// -------------------------------------------------------------------------

func (r *ClusterResource) instantiateNodeClient(node ClusternodeModel) (*service.NitroClient, error) {
	nodeEndpoint := strOrEmpty(node.Endpoint)

	nodeUsername := strOrEmpty(node.Username)
	if nodeUsername == "" {
		nodeUsername = r.client.GetUsername()
	}
	nodePassword := strOrEmpty(node.Password)
	if nodePassword == "" {
		nodePassword = r.client.GetPassword()
	}
	sslVerify := !boolOrFalse(node.InsecureSkipVerify)

	params := service.NitroParams{
		Url:       nodeEndpoint,
		Username:  nodeUsername,
		Password:  nodePassword,
		SslVerify: sslVerify,
	}
	return service.NewNitroClientFromParams(params)
}

// pollClipReachable polls the CLIP (provider endpoint) until it answers.
func (r *ClusterResource) pollClipReachable(delay, interval, total, perPoll time.Duration) error {
	return poll(delay, interval, total, func() (bool, error) {
		err := r.pollNode(perPoll)
		if err != nil {
			if err.Error() == "Timeout" {
				return false, nil
			}
			return false, err
		}
		return true, nil
	})
}

func (r *ClusterResource) pollNode(timeout time.Duration) error {
	username := r.client.GetUsername()
	password := r.client.GetPassword()
	endpoint := r.client.GetURL()
	url := fmt.Sprintf("%s/nitro/v1/config/nslicense", endpoint)

	c := http.Client{Timeout: timeout}
	buff := &bytes.Buffer{}
	req, _ := http.NewRequest("GET", url, buff)
	req.Header.Set("X-NITRO-USER", username)
	req.Header.Set("X-NITRO-PASS", password)
	resp, err := c.Do(req)
	if err != nil {
		if !strings.Contains(err.Error(), "Client.Timeout exceeded") {
			return err
		}
		return fmt.Errorf("Timeout")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Timeout")
	}
	return nil
}

func (r *ClusterResource) pollClusterNodeWithid(nodeid int) (bool, error) {
	data, err := r.client.FindAllResources(service.Clusternode.Type())
	if err != nil {
		return false, err
	}
	nodeidFound := false
	for _, val := range data {
		valNodeid, err := utils.ConvertToInt64(val["nodeid"])
		if err != nil {
			return false, err
		}
		if int(valNodeid) == nodeid {
			nodeidFound = true
			if val["masterstate"] == "ACTIVE" {
				return true, nil
			}
			break
		}
	}
	if !nodeidFound {
		return false, fmt.Errorf("node id %d not in retrieved nodes list", nodeid)
	}
	return false, nil
}

// poll runs fn after an initial delay, then on each interval, until fn reports
// done or the total timeout elapses.
func poll(delay, interval, total time.Duration, fn func() (bool, error)) error {
	if delay > 0 {
		time.Sleep(delay)
	}
	deadline := time.Now().Add(total)
	for {
		done, err := fn()
		if err != nil {
			return err
		}
		if done {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for cluster operation to complete")
		}
		if interval > 0 {
			time.Sleep(interval)
		}
	}
}

func parseTimeouts(delay, interval, total, perPoll types.String) (time.Duration, time.Duration, time.Duration, time.Duration, error) {
	d, i, t, err := parseThreeTimeouts(delay, interval, total)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	p, err := time.ParseDuration(strOrEmpty(perPoll))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return d, i, t, p, nil
}

func parseThreeTimeouts(delay, interval, total types.String) (time.Duration, time.Duration, time.Duration, error) {
	d, err := time.ParseDuration(strOrEmpty(delay))
	if err != nil {
		return 0, 0, 0, err
	}
	i, err := time.ParseDuration(strOrEmpty(interval))
	if err != nil {
		return 0, 0, 0, err
	}
	t, err := time.ParseDuration(strOrEmpty(total))
	if err != nil {
		return 0, 0, 0, err
	}
	return d, i, t, nil
}

// -------------------------------------------------------------------------
// Model <-> payload / set helpers
// -------------------------------------------------------------------------

func (r *ClusterResource) getClusternodes(ctx context.Context, data *ClusterResourceModel, diags *diag.Diagnostics) []ClusternodeModel {
	return r.getClusternodesFrom(ctx, data.Clusternode, diags)
}

func (r *ClusterResource) getClusternodesFrom(ctx context.Context, set types.Set, diags *diag.Diagnostics) []ClusternodeModel {
	var nodes []ClusternodeModel
	if set.IsNull() || set.IsUnknown() {
		return nodes
	}
	d := set.ElementsAs(ctx, &nodes, false)
	if diags != nil {
		diags.Append(d...)
	}
	return nodes
}

func (r *ClusterResource) getClusternodegroups(ctx context.Context, data *ClusterResourceModel, diags *diag.Diagnostics) []ClusternodegroupModel {
	return r.getClusternodegroupsFrom(ctx, data.Clusternodegroup, diags)
}

func (r *ClusterResource) getClusternodegroupsFrom(ctx context.Context, set types.Set, diags *diag.Diagnostics) []ClusternodegroupModel {
	var groups []ClusternodegroupModel
	if set.IsNull() || set.IsUnknown() {
		return groups
	}
	d := set.ElementsAs(ctx, &groups, false)
	if diags != nil {
		diags.Append(d...)
	}
	return groups
}

func clusternodeFromModel(node ClusternodeModel) adccluster.Clusternode {
	return adccluster.Clusternode{
		Backplane:            strOrEmpty(node.Backplane),
		Clearnodegroupconfig: strOrEmpty(node.Clearnodegroupconfig),
		Delay:                utils.IntPtr(intOrZero(node.Delay)),
		Ipaddress:            strOrEmpty(node.Ipaddress),
		Nodegroup:            strOrEmpty(node.Nodegroup),
		Nodeid:               utils.IntPtr(intOrZero(node.Nodeid)),
		Priority:             utils.IntPtr(intOrZero(node.Priority)),
		State:                strOrEmpty(node.State),
		Tunnelmode:           strOrEmpty(node.Tunnelmode),
	}
}

func clusternodegroupFromModel(ng ClusternodegroupModel) adccluster.Clusternodegroup {
	return adccluster.Clusternodegroup{
		Name:     strOrEmpty(ng.Name),
		Priority: utils.IntPtr(intOrZero(ng.Priority)),
		State:    strOrEmpty(ng.State),
		Sticky:   strOrEmpty(ng.Sticky),
		Strict:   strOrEmpty(ng.Strict),
	}
}

// -------------------------------------------------------------------------
// Small utilities
// -------------------------------------------------------------------------

func isClusterModeL3(data *ClusterResourceModel, groups []ClusternodegroupModel) bool {
	if len(groups) == 0 {
		return false
	}
	return strOrEmpty(data.Inc) == "ENABLED"
}

type nodePriority struct {
	nodeid   int
	priority int
}

func getSortedClusternodeIds(nodes []ClusternodeModel) []int {
	ps := make([]nodePriority, 0, len(nodes))
	for _, n := range nodes {
		ps = append(ps, nodePriority{nodeid: intOrZero(n.Nodeid), priority: intOrZero(n.Priority)})
	}
	sort.Slice(ps, func(i, j int) bool {
		if ps[i].priority == ps[j].priority {
			return ps[i].nodeid < ps[j].nodeid
		}
		return ps[i].priority < ps[j].priority
	})
	ids := make([]int, 0, len(ps))
	for _, p := range ps {
		ids = append(ids, p.nodeid)
	}
	return ids
}

func getClusterNodeByid(nodes []ClusternodeModel, id int) (ClusternodeModel, bool) {
	for _, n := range nodes {
		if intOrZero(n.Nodeid) == id {
			return n, true
		}
	}
	return ClusternodeModel{}, false
}

func getClusterNodegroupByName(groups []ClusternodegroupModel, name string) (ClusternodegroupModel, bool) {
	for _, g := range groups {
		if strOrEmpty(g.Name) == name {
			return g, true
		}
	}
	return ClusternodegroupModel{}, false
}

func strOrEmpty(s types.String) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	return s.ValueString()
}

func intOrZero(i types.Int64) int {
	if i.IsNull() || i.IsUnknown() {
		return 0
	}
	return int(i.ValueInt64())
}

func boolOrFalse(b types.Bool) bool {
	if b.IsNull() || b.IsUnknown() {
		return false
	}
	return b.ValueBool()
}
