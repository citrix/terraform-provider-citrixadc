package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/lb"
	"github.com/citrix/adc-nitro-go/resource/config/ssl"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ServiceResource{}
var _ resource.ResourceWithConfigure = (*ServiceResource)(nil)
var _ resource.ResourceWithImportState = (*ServiceResource)(nil)

func NewServiceResource() resource.Resource {
	return &ServiceResource{}
}

// ServiceResource defines the resource implementation.
type ServiceResource struct {
	client *service.NitroClient
}

func (r *ServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *ServiceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (r *ServiceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *ServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data ServiceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating service resource")

	// Determine the service name (generate one if the user did not specify it,
	// mirroring the SDK v2 resource behaviour).
	serviceName := data.Name.ValueString()
	if data.Name.IsNull() || data.Name.IsUnknown() || serviceName == "" {
		serviceName = fmt.Sprintf("tf-service-%d", time.Now().UnixNano())
	}
	data.Name = types.StringValue(serviceName)

	// Validate the referenced lbmonitor / lbvserver up front (SDK v2 parity).
	lbvserverName := data.Lbvserver.ValueString()
	hasLbvserver := !data.Lbvserver.IsNull() && !data.Lbvserver.IsUnknown() && lbvserverName != ""
	if hasLbvserver {
		if !r.client.ResourceExists(service.Lbvserver.Type(), lbvserverName) {
			resp.Diagnostics.AddError("Client Error", "Specified lb vserver does not exist on the NetScaler!")
			return
		}
	}
	lbmonitorName := data.Lbmonitor.ValueString()
	hasLbmonitor := !data.Lbmonitor.IsNull() && !data.Lbmonitor.IsUnknown() && lbmonitorName != ""
	if hasLbmonitor {
		if !r.client.ResourceExists(service.Lbmonitor.Type(), lbmonitorName) {
			resp.Diagnostics.AddError("Client Error", "Specified lb monitor does not exist on the NetScaler!")
			return
		}
	}

	svc := serviceGetThePayloadFromthePlan(ctx, &data)

	// Make API call - named resource, use AddResource
	if _, err := r.client.AddResource(service.Service.Type(), serviceName, &svc); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create service, got error: %s", err))
		return
	}

	// Bind the service to the lb vserver, if requested.
	if hasLbvserver {
		binding := lb.Lbvserverservicebinding{
			Name:        lbvserverName,
			Servicename: serviceName,
		}
		if err := r.client.BindResource(service.Lbvserver.Type(), lbvserverName, service.Service.Type(), serviceName, &binding); err != nil {
			// Roll back the service on bind failure (SDK v2 parity).
			_ = r.client.DeleteResource(service.Service.Type(), serviceName)
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind service %s to lbvserver %s: %s", serviceName, lbvserverName, err))
			return
		}
	}

	// Bind the service to the lb monitor, if requested.
	if hasLbmonitor {
		binding := lb.Lbmonitorservicebinding{
			Monitorname: lbmonitorName,
			Servicename: serviceName,
		}
		if err := r.client.BindResource(service.Lbmonitor.Type(), lbmonitorName, service.Service.Type(), serviceName, &binding); err != nil {
			_ = r.client.DeleteResource(service.Service.Type(), serviceName)
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind service %s to lbmonitor %s: %s", serviceName, lbmonitorName, err))
			return
		}
	}

	// Sync SSL service properties, if requested.
	if hasSslserviceProperties(&data) {
		if err := r.syncSslservice(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync sslservice for %s, got error: %s", serviceName, err))
			return
		}
	}

	// Set ID for the resource before reading state
	data.Id = types.StringValue(serviceName)

	tflog.Trace(ctx, "Created service resource")

	// Read the updated state back
	if !r.readServiceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "service not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data ServiceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading service resource")

	found := r.readServiceFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *ServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state ServiceResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (unset candidates).
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID and name from prior state (name is ForceNew, so it never changes here).
	data.Id = state.Id
	serviceName := state.Id.ValueString()

	tflog.Debug(ctx, "Updating service resource")

	// Detect changes on non-ForceNew, non-action attributes.
	hasChange := serviceHasUpdatableChange(&data, &state)

	// Collect attributes removed from config so they can be unset (reverted to
	// their NITRO defaults) after the base update. Each attribute has a schema
	// Default matching its documented NITRO default, so removal from config
	// produces a real plan diff that drives this Update.
	attributesToUnset := []string{}
	if !data.Accessdown.Equal(state.Accessdown) && config.Accessdown.IsNull() {
		attributesToUnset = append(attributesToUnset, "accessdown")
	}
	if !data.Appflowlog.Equal(state.Appflowlog) && config.Appflowlog.IsNull() {
		attributesToUnset = append(attributesToUnset, "appflowlog")
	}
	if !data.Cacheable.Equal(state.Cacheable) && config.Cacheable.IsNull() {
		attributesToUnset = append(attributesToUnset, "cacheable")
	}
	if !data.Downstateflush.Equal(state.Downstateflush) && config.Downstateflush.IsNull() {
		attributesToUnset = append(attributesToUnset, "downstateflush")
	}
	if !data.Healthmonitor.Equal(state.Healthmonitor) && config.Healthmonitor.IsNull() {
		attributesToUnset = append(attributesToUnset, "healthmonitor")
	}
	if !data.Processlocal.Equal(state.Processlocal) && config.Processlocal.IsNull() {
		attributesToUnset = append(attributesToUnset, "processlocal")
	}
	// Only treat state as changed when the plan carries a real, known value
	// (SDK v2 parity: state change was driven by an explicit config change).
	stateChange := !data.State.IsNull() && !data.State.IsUnknown() &&
		data.State.ValueString() != "" && !data.State.Equal(state.State)
	// Only treat the lb monitor / lb vserver convenience blocks as changed when the
	// plan carries a real, known value. On a refresh of SDK-v2-written state these
	// Optional(+Computed) convenience attrs plan as Unknown; treating Unknown as a
	// change would unbind a separately-managed service_lbmonitor_binding /
	// lbvserver_service_binding (SDK v2 parity: change was driven by an explicit
	// config change only).
	lbmonitorChanged := !data.Lbmonitor.IsUnknown() && !data.Lbmonitor.Equal(state.Lbmonitor)
	lbvserverChanged := !data.Lbvserver.IsUnknown() && !data.Lbvserver.Equal(state.Lbvserver)

	// Unbind existing lb monitor binding, if it changed.
	if lbmonitorChanged {
		oldLbmonitorName := state.Lbmonitor.ValueString()
		oldMonitorIsDefault := oldLbmonitorName == "ping-default" || oldLbmonitorName == "tcp-default"
		if oldLbmonitorName != "" && !oldMonitorIsDefault {
			if err := r.client.UnbindResource(service.Lbmonitor.Type(), oldLbmonitorName, service.Service.Type(), serviceName, "servicename"); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error unbinding lbmonitor %s from service %s: %s", oldLbmonitorName, serviceName, err))
				return
			}
		}
	}

	// Unbind existing lb vserver binding, if it changed.
	if lbvserverChanged {
		oldLbvserverName := state.Lbvserver.ValueString()
		if oldLbvserverName != "" {
			if err := r.client.UnbindResource(service.Lbvserver.Type(), oldLbvserverName, service.Service.Type(), serviceName, "servicename"); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error unbinding lbvserver %s from service %s: %s", oldLbvserverName, serviceName, err))
				return
			}
		}
	}

	if hasChange {
		svc := serviceGetTheUpdatablePayloadFromThePlan(ctx, &data)
		svc.Name = serviceName
		// Only issue the base SET when the payload carries a field beyond the name.
		// On a refresh of SDK-v2-written state the Optional+Computed attributes can
		// plan as Unknown and be dropped from the payload, which would otherwise
		// produce a name-only SET that NITRO rejects with errorcode 1094
		// ("Too few arguments"). This mirrors the SDK v2 per-attribute hasChange
		// gating (and lbvserverPayloadHasMutableFields).
		if servicePayloadHasMutableFields(&svc) {
			if _, err := r.client.UpdateResource(service.Service.Type(), serviceName, &svc); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update service %s, got error: %s", serviceName, err))
				return
			}
		} else {
			tflog.Debug(ctx, "Skipping base service SET: payload carries no field beyond name")
		}
	} else {
		tflog.Debug(ctx, "No updatable attribute changes detected for service resource")
	}

	// Rebind the new lb monitor, if it changed. Default monitors are auto-bound
	// upon the unbind of the last non-default monitor, so they are not bound explicitly.
	if lbmonitorChanged {
		newLbmonitorName := data.Lbmonitor.ValueString()
		newMonitorIsDefault := newLbmonitorName == "ping-default" || newLbmonitorName == "tcp-default"
		if newLbmonitorName != "" && !newMonitorIsDefault {
			binding := lb.Lbmonitorservicebinding{
				Monitorname: newLbmonitorName,
				Servicename: serviceName,
			}
			if err := r.client.BindResource(service.Lbmonitor.Type(), newLbmonitorName, service.Service.Type(), serviceName, &binding); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind lbmonitor %s to service %s: %s", newLbmonitorName, serviceName, err))
				return
			}
		}
	}

	// Rebind the new lb vserver, if it changed.
	if lbvserverChanged {
		newLbvserverName := data.Lbvserver.ValueString()
		if newLbvserverName != "" {
			binding := lb.Lbvserverservicebinding{
				Name:        newLbvserverName,
				Servicename: serviceName,
			}
			if err := r.client.BindResource(service.Lbvserver.Type(), newLbvserverName, service.Service.Type(), serviceName, &binding); err != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to bind lbvserver %s to service %s: %s", newLbvserverName, serviceName, err))
				return
			}
		}
	}

	// Sync SSL service properties, if requested.
	if hasSslserviceProperties(&data) {
		if err := r.syncSslservice(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to sync sslservice for %s, got error: %s", serviceName, err))
			return
		}
	}

	// Apply enable/disable state change.
	if stateChange {
		if err := r.doServiceStateChange(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling service %s: %s", serviceName, err))
			return
		}
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// defaults. Done after the base update so any default value the update payload
	// carried for a removed attribute is superseded by the unset.
	unsetIdPayload := map[string]interface{}{
		"name": serviceName,
	}
	if err := utils.ExecuteUnset(r.client, service.Service.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset service attributes, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Updated service resource")

	// Read the updated state back
	if !r.readServiceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "service not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ServiceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data ServiceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting service resource")

	serviceName := data.Id.ValueString()
	if err := r.client.DeleteResource(service.Service.Type(), serviceName); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete service %s, got error: %s", serviceName, err))
		return
	}

	tflog.Trace(ctx, "Deleted service resource")
}

// readServiceFromApi reads the service and its bindings/SSL properties from the
// ADC and applies them to the model. Returns false when the service no longer exists.
func (r *ServiceResource) readServiceFromApi(ctx context.Context, data *ServiceResourceModel, diags *diag.Diagnostics) bool {
	serviceName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Service.Type(), serviceName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read service %s, got error: %s", serviceName, err))
		return false
	}

	serviceSetAttrFromGet(ctx, data, getResponseData)

	// Read the bound lb vserver, only if it is being managed (SDK v2 parity).
	if !data.Lbvserver.IsNull() && !data.Lbvserver.IsUnknown() && data.Lbvserver.ValueString() != "" {
		vserverBindings, err := r.client.FindResourceArray(service.Svcbindings.Type(), serviceName)
		if err != nil {
			return false
		}
		boundVserver := ""
		for _, vserver := range vserverBindings {
			if vs, ok := vserver["vservername"]; ok && vs != nil {
				if s, ok2 := vs.(string); ok2 {
					boundVserver = s
					break
				}
			}
		}
		data.Lbvserver = types.StringValue(boundVserver)
	}

	// Read the bound lb monitor, only when the model attr is non-null (SDK v2 parity
	// with the guarded lbvserver convenience block above). A null lbmonitor means the
	// binding is not being managed through the service resource (it is managed by a
	// separate service_lbmonitor_binding), so adopting the externally-bound monitor
	// here would make Update believe it is a change and unbind it. Note: on Create/
	// Update an unconfigured (Computed) lbmonitor is Unknown, not null, so the default
	// monitor is still adopted back (SDK v2 parity, e.g. "tcp-default").
	if !data.Lbmonitor.IsNull() {
		boundMonitors, err := r.client.FindAllBoundResources(service.Service.Type(), serviceName, service.Lbmonitor.Type())
		if err != nil {
			return false
		}
		boundMonitor := ""
		for _, monitor := range boundMonitors {
			if mon, ok := monitor["monitor_name"]; ok && mon != nil {
				if s, ok2 := mon.(string); ok2 {
					boundMonitor = s
					break
				}
			}
		}
		data.Lbmonitor = types.StringValue(boundMonitor)
	}

	// Read the SSL service properties, if being managed.
	if hasSslserviceProperties(data) {
		if err := r.readSslservice(ctx, data); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to read sslservice for %s, got error: %s", serviceName, err))
			return false
		}
	}

	// Action-only attributes (delay, graceful, wait_until_disabled) and the SSL
	// convenience block (snienable, commonname) are not part of the service GET
	// response and are only read back when configured. When they are left
	// unconfigured their planned value stays Unknown; resolve it to null so the
	// post-apply state is fully known (avoids "invalid result object after apply").
	if data.Delay.IsUnknown() {
		data.Delay = types.Int64Null()
	}
	if data.Graceful.IsUnknown() {
		data.Graceful = types.StringNull()
	}
	if data.Snienable.IsUnknown() {
		data.Snienable = types.StringNull()
	}
	if data.Commonname.IsUnknown() {
		data.Commonname = types.StringNull()
	}
	if data.WaitUntilDisabled.IsUnknown() {
		data.WaitUntilDisabled = types.BoolNull()
	}

	return true
}

// servicePayloadHasMutableFields reports whether the service SET payload carries
// any attribute beyond the resource name. NITRO rejects a name-only SET with
// errorcode 1094 ("Too few arguments"), so the base update is skipped in that case.
// This mirrors citrixadc_framework/lbvserver's lbvserverPayloadHasMutableFields and
// the SDK v2 per-attribute hasChange gating: on a refresh of SDK-v2-written state the
// Optional+Computed attrs can plan as Unknown and get dropped from the payload,
// leaving a name-only SET.
func servicePayloadHasMutableFields(payload *basic.Service) bool {
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

// serviceHasUpdatableChange reports whether any non-ForceNew, non-action
// attribute differs between the plan and prior state.
func serviceHasUpdatableChange(data, state *ServiceResourceModel) bool {
	return !data.Internal.Equal(state.Internal) ||
		!data.Accessdown.Equal(state.Accessdown) ||
		!data.All.Equal(state.All) ||
		!data.Appflowlog.Equal(state.Appflowlog) ||
		!data.Cacheable.Equal(state.Cacheable) ||
		!data.Cip.Equal(state.Cip) ||
		!data.Cipheader.Equal(state.Cipheader) ||
		!data.Cka.Equal(state.Cka) ||
		!data.Clttimeout.Equal(state.Clttimeout) ||
		!data.Cmp.Equal(state.Cmp) ||
		!data.Comment.Equal(state.Comment) ||
		!data.Contentinspectionprofilename.Equal(state.Contentinspectionprofilename) ||
		!data.Customserverid.Equal(state.Customserverid) ||
		!data.Dnsprofilename.Equal(state.Dnsprofilename) ||
		!data.Downstateflush.Equal(state.Downstateflush) ||
		!data.Hashid.Equal(state.Hashid) ||
		!data.Healthmonitor.Equal(state.Healthmonitor) ||
		!data.Httpprofilename.Equal(state.Httpprofilename) ||
		!data.Ipaddress.Equal(state.Ipaddress) ||
		!data.Maxbandwidth.Equal(state.Maxbandwidth) ||
		!data.Maxclient.Equal(state.Maxclient) ||
		!data.Maxreq.Equal(state.Maxreq) ||
		!data.Monconnectionclose.Equal(state.Monconnectionclose) ||
		!data.Monitornamesvc.Equal(state.Monitornamesvc) ||
		!data.Monthreshold.Equal(state.Monthreshold) ||
		!data.Netprofile.Equal(state.Netprofile) ||
		!data.Pathmonitor.Equal(state.Pathmonitor) ||
		!data.Pathmonitorindv.Equal(state.Pathmonitorindv) ||
		!data.Processlocal.Equal(state.Processlocal) ||
		!data.Quicprofilename.Equal(state.Quicprofilename) ||
		!data.Rtspsessionidremap.Equal(state.Rtspsessionidremap) ||
		!data.Serverid.Equal(state.Serverid) ||
		!data.Sp.Equal(state.Sp) ||
		!data.Svrtimeout.Equal(state.Svrtimeout) ||
		!data.Tcpb.Equal(state.Tcpb) ||
		!data.Tcpprofilename.Equal(state.Tcpprofilename) ||
		!data.Useproxyport.Equal(state.Useproxyport) ||
		!data.Usip.Equal(state.Usip) ||
		!data.Weight.Equal(state.Weight)
}

// hasSslserviceProperties reports whether the SSL service convenience block is in use.
func hasSslserviceProperties(data *ServiceResourceModel) bool {
	if !data.Snienable.IsNull() && !data.Snienable.IsUnknown() && data.Snienable.ValueString() != "" {
		return true
	}
	if !data.Commonname.IsNull() && !data.Commonname.IsUnknown() && data.Commonname.ValueString() != "" {
		return true
	}
	return false
}

// syncSslservice pushes the SSL service properties (snienable, commonname).
func (r *ServiceResource) syncSslservice(ctx context.Context, data *ServiceResourceModel) error {
	tflog.Debug(ctx, "In syncSslservice Function")
	if !hasSslserviceProperties(data) {
		return nil
	}
	sslserviceStruct := ssl.Sslservice{
		Servicename: data.Name.ValueString(),
		Snienable:   data.Snienable.ValueString(),
		Commonname:  data.Commonname.ValueString(),
	}
	return r.client.UpdateUnnamedResource("sslservice", &sslserviceStruct)
}

// readSslservice reads the SSL service properties back into the model.
func (r *ServiceResource) readSslservice(ctx context.Context, data *ServiceResourceModel) error {
	tflog.Debug(ctx, "In readSslservice Function")
	if !hasSslserviceProperties(data) {
		return nil
	}
	findParams := service.FindParams{
		ResourceType: "sslservice",
		ResourceName: data.Name.ValueString(),
	}
	arr, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		return err
	}
	if len(arr) > 1 {
		return fmt.Errorf("too many sslservice results %v", arr)
	}
	if len(arr) == 1 {
		if val, ok := arr[0]["snienable"]; ok && val != nil {
			if s, ok2 := val.(string); ok2 {
				data.Snienable = types.StringValue(s)
			}
		}
		if val, ok := arr[0]["commonname"]; ok && val != nil {
			if s, ok2 := val.(string); ok2 {
				data.Commonname = types.StringValue(s)
			}
		}
	}
	return nil
}

// doServiceStateChange enables or disables the service using the NITRO action API.
func (r *ServiceResource) doServiceStateChange(ctx context.Context, data *ServiceResourceModel) error {
	tflog.Debug(ctx, "In doServiceStateChange Function")

	newstate := data.State.ValueString()

	// ActOnResource fails if superfluous attributes are supplied, so we build a
	// lean payload carrying only the fields the action requires (SDK v2 parity).
	switch newstate {
	case "ENABLED":
		svc := basic.Service{Name: data.Name.ValueString()}
		return r.client.ActOnResource(service.Service.Type(), svc, "enable")
	case "DISABLED":
		svc := basic.Service{Name: data.Name.ValueString()}
		if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
			svc.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
		}
		if !data.Graceful.IsNull() && !data.Graceful.IsUnknown() {
			svc.Graceful = data.Graceful.ValueString()
		}
		if err := r.client.ActOnResource(service.Service.Type(), svc, "disable"); err != nil {
			return err
		}
		if !data.WaitUntilDisabled.IsNull() && data.WaitUntilDisabled.ValueBool() {
			return r.serviceWaitDisableState(ctx, data)
		}
		return nil
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// serviceWaitDisableState polls the service until it reaches the OUT OF SERVICE state.
func (r *ServiceResource) serviceWaitDisableState(ctx context.Context, data *ServiceResourceModel) error {
	tflog.Debug(ctx, "In serviceWaitDisableState Function")

	timeout, err := time.ParseDuration(nonEmptyOrDefault(data.DisabledTimeout, "2m"))
	if err != nil {
		return err
	}
	pollInterval, err := time.ParseDuration(nonEmptyOrDefault(data.DisabledPollInterval, "5s"))
	if err != nil {
		return err
	}
	pollDelay, err := time.ParseDuration(nonEmptyOrDefault(data.DisabledPollDelay, "2s"))
	if err != nil {
		return err
	}

	serviceName := data.Name.ValueString()
	deadline := time.Now().Add(timeout)
	time.Sleep(pollDelay)
	for {
		getResponseData, err := r.client.FindResource(service.Service.Type(), serviceName)
		if err != nil {
			return err
		}
		if getResponseData["svrstate"] == "OUT OF SERVICE" {
			return nil
		}
		if time.Now().After(deadline) {
			return fmt.Errorf("timed out waiting for service %s to be disabled", serviceName)
		}
		time.Sleep(pollInterval)
	}
}

func nonEmptyOrDefault(v types.String, def string) string {
	if !v.IsNull() && !v.IsUnknown() && v.ValueString() != "" {
		return v.ValueString()
	}
	return def
}
