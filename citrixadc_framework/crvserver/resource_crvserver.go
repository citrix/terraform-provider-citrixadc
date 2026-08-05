package crvserver

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cr"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CrvserverResource{}
var _ resource.ResourceWithConfigure = (*CrvserverResource)(nil)
var _ resource.ResourceWithImportState = (*CrvserverResource)(nil)

func NewCrvserverResource() resource.Resource {
	return &CrvserverResource{}
}

// CrvserverResource defines the resource implementation.
type CrvserverResource struct {
	client *service.NitroClient
}

func (r *CrvserverResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *CrvserverResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_crvserver"
}

func (r *CrvserverResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *CrvserverResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data CrvserverResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating crvserver resource")

	crvserver := crvserverGetThePayloadFromthePlan(ctx, &data)

	// Named resource - use AddResource (NITRO doc: add -> HTTP POST)
	crvserverName := data.Name.ValueString()
	_, err := r.client.AddResource(service.Crvserver.Type(), crvserverName, &crvserver)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create crvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created crvserver resource")

	// ID is the resource name (single unique attribute), matching SDK v2.
	data.Id = types.StringValue(crvserverName)

	// Read the updated state back
	if !r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "crvserver not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data CrvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading crvserver resource")

	found := r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *CrvserverResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state CrvserverResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the live name (ID) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating crvserver resource")

	// Handle in-place rename via NITRO ?action=rename when newname changes.
	// The rename source must be the CURRENT LIVE name, tracked by state.Id.
	if !data.Newname.IsNull() && !data.Newname.IsUnknown() && data.Newname.ValueString() != "" && !data.Newname.Equal(state.Newname) {
		renamePayload := cr.Crvserver{
			Name:    state.Id.ValueString(),
			Newname: data.Newname.ValueString(),
		}
		if err := r.client.ActOnResource(service.Crvserver.Type(), &renamePayload, "rename"); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to rename crvserver, got error: %s", err))
			return
		}
		// The live object is now named newname; track it via the ID.
		data.Id = types.StringValue(data.Newname.ValueString())
	}

	crvserverName := data.Id.ValueString()

	// Build a payload from only the changed, updatable attributes (mirrors SDK v2).
	crvserver := cr.Crvserver{
		Name: crvserverName,
	}
	hasChange := false

	if !data.Appflowlog.IsUnknown() && !data.Appflowlog.Equal(state.Appflowlog) {
		crvserver.Appflowlog = data.Appflowlog.ValueString()
		hasChange = true
	}
	if !data.Arp.IsUnknown() && !data.Arp.Equal(state.Arp) {
		crvserver.Arp = data.Arp.ValueString()
		hasChange = true
	}
	if !data.Backendssl.IsUnknown() && !data.Backendssl.Equal(state.Backendssl) {
		crvserver.Backendssl = data.Backendssl.ValueString()
		hasChange = true
	}
	if !data.Backupvserver.IsUnknown() && !data.Backupvserver.Equal(state.Backupvserver) {
		crvserver.Backupvserver = data.Backupvserver.ValueString()
		hasChange = true
	}
	if !data.Cachetype.IsUnknown() && !data.Cachetype.Equal(state.Cachetype) {
		crvserver.Cachetype = data.Cachetype.ValueString()
		hasChange = true
	}
	if !data.Cachevserver.IsUnknown() && !data.Cachevserver.Equal(state.Cachevserver) {
		crvserver.Cachevserver = data.Cachevserver.ValueString()
		hasChange = true
	}
	if !data.Clttimeout.IsUnknown() && !data.Clttimeout.Equal(state.Clttimeout) {
		crvserver.Clttimeout = utils.IntPtr(int(data.Clttimeout.ValueInt64()))
		hasChange = true
	}
	if !data.Comment.IsUnknown() && !data.Comment.Equal(state.Comment) {
		crvserver.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if !data.Destinationvserver.IsUnknown() && !data.Destinationvserver.Equal(state.Destinationvserver) {
		crvserver.Destinationvserver = data.Destinationvserver.ValueString()
		hasChange = true
	}
	if !data.Disableprimaryondown.IsUnknown() && !data.Disableprimaryondown.Equal(state.Disableprimaryondown) {
		crvserver.Disableprimaryondown = data.Disableprimaryondown.ValueString()
		hasChange = true
	}
	if !data.Disallowserviceaccess.IsUnknown() && !data.Disallowserviceaccess.Equal(state.Disallowserviceaccess) {
		crvserver.Disallowserviceaccess = data.Disallowserviceaccess.ValueString()
		hasChange = true
	}
	if !data.Dnsvservername.IsUnknown() && !data.Dnsvservername.Equal(state.Dnsvservername) {
		crvserver.Dnsvservername = data.Dnsvservername.ValueString()
		hasChange = true
	}
	if !data.Domain.IsUnknown() && !data.Domain.Equal(state.Domain) {
		crvserver.Domain = data.Domain.ValueString()
		hasChange = true
	}
	if !data.Downstateflush.IsUnknown() && !data.Downstateflush.Equal(state.Downstateflush) {
		crvserver.Downstateflush = data.Downstateflush.ValueString()
		hasChange = true
	}
	if !data.Format.IsUnknown() && !data.Format.Equal(state.Format) {
		crvserver.Format = data.Format.ValueString()
		hasChange = true
	}
	if !data.Ghost.IsUnknown() && !data.Ghost.Equal(state.Ghost) {
		crvserver.Ghost = data.Ghost.ValueString()
		hasChange = true
	}
	if !data.Httpprofilename.IsUnknown() && !data.Httpprofilename.Equal(state.Httpprofilename) {
		crvserver.Httpprofilename = data.Httpprofilename.ValueString()
		hasChange = true
	}
	if !data.Icmpvsrresponse.IsUnknown() && !data.Icmpvsrresponse.Equal(state.Icmpvsrresponse) {
		crvserver.Icmpvsrresponse = data.Icmpvsrresponse.ValueString()
		hasChange = true
	}
	if !data.Ipset.IsUnknown() && !data.Ipset.Equal(state.Ipset) {
		crvserver.Ipset = data.Ipset.ValueString()
		hasChange = true
	}
	if !data.Ipv46.IsUnknown() && !data.Ipv46.Equal(state.Ipv46) {
		crvserver.Ipv46 = data.Ipv46.ValueString()
		hasChange = true
	}
	if !data.L2conn.IsUnknown() && !data.L2conn.Equal(state.L2conn) {
		crvserver.L2conn = data.L2conn.ValueString()
		hasChange = true
	}
	if !data.Listenpolicy.IsUnknown() && !data.Listenpolicy.Equal(state.Listenpolicy) {
		crvserver.Listenpolicy = data.Listenpolicy.ValueString()
		hasChange = true
	}
	if !data.Listenpriority.IsUnknown() && !data.Listenpriority.Equal(state.Listenpriority) {
		crvserver.Listenpriority = utils.IntPtr(int(data.Listenpriority.ValueInt64()))
		hasChange = true
	}
	if !data.Map.IsUnknown() && !data.Map.Equal(state.Map) {
		crvserver.Map = data.Map.ValueString()
		hasChange = true
	}
	if !data.Netprofile.IsUnknown() && !data.Netprofile.Equal(state.Netprofile) {
		crvserver.Netprofile = data.Netprofile.ValueString()
		hasChange = true
	}
	if !data.Onpolicymatch.IsUnknown() && !data.Onpolicymatch.Equal(state.Onpolicymatch) {
		crvserver.Onpolicymatch = data.Onpolicymatch.ValueString()
		hasChange = true
	}
	if !data.Originusip.IsUnknown() && !data.Originusip.Equal(state.Originusip) {
		crvserver.Originusip = data.Originusip.ValueString()
		hasChange = true
	}
	if !data.Port.IsUnknown() && !data.Port.Equal(state.Port) {
		crvserver.Port = utils.IntPtr(int(data.Port.ValueInt64()))
		hasChange = true
	}
	if !data.Precedence.IsUnknown() && !data.Precedence.Equal(state.Precedence) {
		crvserver.Precedence = data.Precedence.ValueString()
		hasChange = true
	}
	if !data.Probeport.IsUnknown() && !data.Probeport.Equal(state.Probeport) {
		crvserver.Probeport = utils.IntPtr(int(data.Probeport.ValueInt64()))
		hasChange = true
	}
	if !data.Probeprotocol.IsUnknown() && !data.Probeprotocol.Equal(state.Probeprotocol) {
		crvserver.Probeprotocol = data.Probeprotocol.ValueString()
		hasChange = true
	}
	if !data.Probesuccessresponsecode.IsUnknown() && !data.Probesuccessresponsecode.Equal(state.Probesuccessresponsecode) {
		crvserver.Probesuccessresponsecode = data.Probesuccessresponsecode.ValueString()
		hasChange = true
	}
	if !data.Range.IsUnknown() && !data.Range.Equal(state.Range) {
		crvserver.Range = utils.IntPtr(int(data.Range.ValueInt64()))
		hasChange = true
	}
	if !data.Redirect.IsUnknown() && !data.Redirect.Equal(state.Redirect) {
		crvserver.Redirect = data.Redirect.ValueString()
		hasChange = true
	}
	if !data.Redirecturl.IsUnknown() && !data.Redirecturl.Equal(state.Redirecturl) {
		crvserver.Redirecturl = data.Redirecturl.ValueString()
		hasChange = true
	}
	if !data.Reuse.IsUnknown() && !data.Reuse.Equal(state.Reuse) {
		crvserver.Reuse = data.Reuse.ValueString()
		hasChange = true
	}
	if !data.Rhistate.IsUnknown() && !data.Rhistate.Equal(state.Rhistate) {
		crvserver.Rhistate = data.Rhistate.ValueString()
		hasChange = true
	}
	if !data.Servicetype.IsUnknown() && !data.Servicetype.Equal(state.Servicetype) {
		crvserver.Servicetype = data.Servicetype.ValueString()
		hasChange = true
	}
	if !data.Sopersistencetimeout.IsUnknown() && !data.Sopersistencetimeout.Equal(state.Sopersistencetimeout) {
		crvserver.Sopersistencetimeout = utils.IntPtr(int(data.Sopersistencetimeout.ValueInt64()))
		hasChange = true
	}
	if !data.Sothreshold.IsUnknown() && !data.Sothreshold.Equal(state.Sothreshold) {
		crvserver.Sothreshold = utils.IntPtr(int(data.Sothreshold.ValueInt64()))
		hasChange = true
	}
	if !data.Srcipexpr.IsUnknown() && !data.Srcipexpr.Equal(state.Srcipexpr) {
		crvserver.Srcipexpr = data.Srcipexpr.ValueString()
		hasChange = true
	}
	if !data.Tcpprobeport.IsUnknown() && !data.Tcpprobeport.Equal(state.Tcpprobeport) {
		crvserver.Tcpprobeport = utils.IntPtr(int(data.Tcpprobeport.ValueInt64()))
		hasChange = true
	}
	if !data.Tcpprofilename.IsUnknown() && !data.Tcpprofilename.Equal(state.Tcpprofilename) {
		crvserver.Tcpprofilename = data.Tcpprofilename.ValueString()
		hasChange = true
	}
	if !data.Td.IsUnknown() && !data.Td.Equal(state.Td) {
		crvserver.Td = utils.IntPtr(int(data.Td.ValueInt64()))
		hasChange = true
	}
	if !data.Useoriginipportforcache.IsUnknown() && !data.Useoriginipportforcache.Equal(state.Useoriginipportforcache) {
		crvserver.Useoriginipportforcache = data.Useoriginipportforcache.ValueString()
		hasChange = true
	}
	if !data.Useportrange.IsUnknown() && !data.Useportrange.Equal(state.Useportrange) {
		crvserver.Useportrange = data.Useportrange.ValueString()
		hasChange = true
	}
	if !data.Via.IsUnknown() && !data.Via.Equal(state.Via) {
		crvserver.Via = data.Via.ValueString()
		hasChange = true
	}

	// The state attribute is applied via the enable/disable actions, matching SDK v2.
	stateChange := !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown()

	if stateChange {
		if err := doCrvserverStateChange(r.client, crvserverName, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling crvserver %s: %s", crvserverName, err))
			return
		}
	}

	if hasChange {
		if _, err := r.client.UpdateResource(service.Crvserver.Type(), crvserverName, &crvserver); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update crvserver, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated crvserver resource")
	} else {
		tflog.Debug(ctx, "No updatable-attribute changes detected for crvserver resource")
	}

	// Read the updated state back
	if !r.readCrvserverFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "crvserver not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CrvserverResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data CrvserverResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting crvserver resource")

	// Named resource - delete by the live name (ID).
	err := r.client.DeleteResource(service.Crvserver.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete crvserver, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted crvserver resource")
}

// Helper function to read crvserver data from API. Returns false if the resource
// no longer exists on the ADC.
func (r *CrvserverResource) readCrvserverFromApi(ctx context.Context, data *CrvserverResourceModel, diags *diag.Diagnostics) bool {
	crvserverName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Crvserver.Type(), crvserverName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read crvserver, got error: %s", err))
		return false
	}

	crvserverSetAttrFromGet(ctx, data, getResponseData)

	return true
}

// doCrvserverStateChange enables or disables the cache redirection virtual server,
// mirroring the SDK v2 behavior for the state attribute.
func doCrvserverStateChange(client *service.NitroClient, name, newstate string) error {
	crvserver := cr.Crvserver{
		Name: name,
	}

	switch newstate {
	case "ENABLED":
		return client.ActOnResource(service.Crvserver.Type(), &crvserver, "enable")
	case "DISABLED":
		return client.ActOnResource(service.Crvserver.Type(), &crvserver, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}
