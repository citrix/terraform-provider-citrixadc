package gslbservice

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"
	"github.com/citrix/adc-nitro-go/resource/config/gslb"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &GslbserviceResource{}
var _ resource.ResourceWithConfigure = (*GslbserviceResource)(nil)
var _ resource.ResourceWithImportState = (*GslbserviceResource)(nil)

func NewGslbserviceResource() resource.Resource {
	return &GslbserviceResource{}
}

// GslbserviceResource defines the resource implementation.
type GslbserviceResource struct {
	client *service.NitroClient
}

func (r *GslbserviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *GslbserviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gslbservice"
}

func (r *GslbserviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *GslbserviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data GslbserviceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating gslbservice resource")

	servicename := data.Servicename.ValueString()
	gslbservice := gslbserviceGetThePayloadFromthePlan(ctx, &data)

	_, err := r.client.AddResource(service.Gslbservice.Type(), servicename, &gslbservice)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create gslbservice, got error: %s", err))
		return
	}
	data.Id = types.StringValue(servicename)

	// Add lbmonitor bindings from the plan (old set is empty on create).
	resp.Diagnostics.Append(r.syncLbmonitorBindings(ctx, servicename, types.SetNull(types.ObjectType{AttrTypes: lbmonitorbindingAttrTypes}), data.Lbmonitorbinding)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Created gslbservice resource")

	r.readGslbserviceFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data GslbserviceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading gslbservice resource")

	found := r.readGslbserviceFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		resp.State.RemoveResource(ctx)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data GslbserviceResourceModel
	var state GslbserviceResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Updating gslbservice resource")

	servicename := data.Servicename.ValueString()

	gslbservice := gslb.Gslbservice{Servicename: servicename}
	hasChange := false

	if !data.Appflowlog.Equal(state.Appflowlog) {
		gslbservice.Appflowlog = data.Appflowlog.ValueString()
		hasChange = true
	}
	if !data.Cip.Equal(state.Cip) {
		gslbservice.Cip = data.Cip.ValueString()
		hasChange = true
	}
	if !data.Cipheader.Equal(state.Cipheader) {
		gslbservice.Cipheader = data.Cipheader.ValueString()
		hasChange = true
	}
	if !data.Comment.Equal(state.Comment) {
		gslbservice.Comment = data.Comment.ValueString()
		hasChange = true
	}
	if !data.Downstateflush.Equal(state.Downstateflush) {
		gslbservice.Downstateflush = data.Downstateflush.ValueString()
		hasChange = true
	}
	if !data.Hashid.Equal(state.Hashid) {
		gslbservice.Hashid = utils.IntPtr(int(data.Hashid.ValueInt64()))
		hasChange = true
	}
	if !data.Healthmonitor.Equal(state.Healthmonitor) {
		gslbservice.Healthmonitor = data.Healthmonitor.ValueString()
		hasChange = true
	}
	if !data.Ipaddress.Equal(state.Ipaddress) {
		gslbservice.Ipaddress = data.Ipaddress.ValueString()
		hasChange = true
	}
	if !data.Maxaaausers.Equal(state.Maxaaausers) {
		gslbservice.Maxaaausers = utils.IntPtr(int(data.Maxaaausers.ValueInt64()))
		hasChange = true
	}
	if !data.Maxbandwidth.Equal(state.Maxbandwidth) {
		gslbservice.Maxbandwidth = utils.IntPtr(int(data.Maxbandwidth.ValueInt64()))
		hasChange = true
	}
	if !data.Maxclient.Equal(state.Maxclient) {
		gslbservice.Maxclient = utils.IntPtr(int(data.Maxclient.ValueInt64()))
		hasChange = true
	}
	if !data.Monitornamesvc.Equal(state.Monitornamesvc) {
		gslbservice.Monitornamesvc = data.Monitornamesvc.ValueString()
		hasChange = true
	}
	if !data.Monthreshold.Equal(state.Monthreshold) {
		gslbservice.Monthreshold = utils.IntPtr(int(data.Monthreshold.ValueInt64()))
		hasChange = true
	}
	if !data.Naptrdomainttl.Equal(state.Naptrdomainttl) {
		gslbservice.Naptrdomainttl = utils.IntPtr(int(data.Naptrdomainttl.ValueInt64()))
		hasChange = true
	}
	if !data.Naptrorder.Equal(state.Naptrorder) {
		gslbservice.Naptrorder = utils.IntPtr(int(data.Naptrorder.ValueInt64()))
		hasChange = true
	}
	if !data.Naptrpreference.Equal(state.Naptrpreference) {
		gslbservice.Naptrpreference = utils.IntPtr(int(data.Naptrpreference.ValueInt64()))
		hasChange = true
	}
	if !data.Naptrreplacement.Equal(state.Naptrreplacement) {
		gslbservice.Naptrreplacement = data.Naptrreplacement.ValueString()
		hasChange = true
	}
	if !data.Naptrservices.Equal(state.Naptrservices) {
		gslbservice.Naptrservices = data.Naptrservices.ValueString()
		hasChange = true
	}
	if !data.Publicip.Equal(state.Publicip) {
		gslbservice.Publicip = data.Publicip.ValueString()
		hasChange = true
	}
	if !data.Publicport.Equal(state.Publicport) {
		gslbservice.Publicport = utils.IntPtr(int(data.Publicport.ValueInt64()))
		hasChange = true
	}
	if !data.Sitepersistence.Equal(state.Sitepersistence) {
		gslbservice.Sitepersistence = data.Sitepersistence.ValueString()
		hasChange = true
	}
	if !data.Siteprefix.Equal(state.Siteprefix) {
		gslbservice.Siteprefix = data.Siteprefix.ValueString()
		hasChange = true
	}
	if !data.Viewip.Equal(state.Viewip) {
		gslbservice.Viewip = data.Viewip.ValueString()
		hasChange = true
	}
	if !data.Viewname.Equal(state.Viewname) {
		gslbservice.Viewname = data.Viewname.ValueString()
		hasChange = true
	}
	if !data.Weight.Equal(state.Weight) {
		gslbservice.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
		hasChange = true
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Gslbservice.Type(), servicename, &gslbservice)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update gslbservice %s, got error: %s", servicename, err))
			return
		}
	}

	// State change is applied via an enable/disable action, not the update payload.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doGslbServiceStateChange(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to change state of gslbservice %s, got error: %s", servicename, err))
			return
		}
	}

	// Reconcile lbmonitor bindings.
	resp.Diagnostics.Append(r.syncLbmonitorBindings(ctx, servicename, state.Lbmonitorbinding, data.Lbmonitorbinding)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, "Updated gslbservice resource")

	r.readGslbserviceFromApi(ctx, &data, &resp.Diagnostics)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *GslbserviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data GslbserviceResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting gslbservice resource")

	err := r.client.DeleteResource(service.Gslbservice.Type(), data.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete gslbservice %s, got error: %s", data.Id.ValueString(), err))
		return
	}

	tflog.Trace(ctx, "Deleted gslbservice resource")
}

// readGslbserviceFromApi reads the gslbservice (and its lbmonitor bindings) into the model.
// Returns false if the resource no longer exists on the ADC.
func (r *GslbserviceResource) readGslbserviceFromApi(ctx context.Context, data *GslbserviceResourceModel, diags *diag.Diagnostics) bool {
	servicename := data.Id.ValueString()
	if servicename == "" {
		servicename = data.Servicename.ValueString()
	}

	getResponseData, err := r.client.FindResource(service.Gslbservice.Type(), servicename)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read gslbservice %s, got error: %s", servicename, err))
		return false
	}
	if getResponseData == nil {
		return false
	}

	gslbserviceSetAttrFromGet(ctx, data, getResponseData)

	// Refresh lbmonitor bindings only when the resource manages them (state/plan non-null).
	if !data.Lbmonitorbinding.IsNull() {
		diags.Append(r.readLbmonitorBindings(ctx, servicename, data)...)
	}

	return true
}

// doGslbServiceStateChange enables/disables the underlying service, mirroring SDK v2.
func (r *GslbserviceResource) doGslbServiceStateChange(ctx context.Context, data *GslbserviceResourceModel) error {
	svc := basic.Service{Name: data.Servicename.ValueString()}
	newstate := data.State.ValueString()

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Service.Type(), svc, "enable")
	case "DISABLED":
		if !data.Delay.IsNull() && !data.Delay.IsUnknown() {
			svc.Delay = utils.IntPtr(int(data.Delay.ValueInt64()))
		}
		return r.client.ActOnResource(service.Service.Type(), svc, "disable")
	default:
		return fmt.Errorf("%q is not a valid state; use ENABLED or DISABLED", newstate)
	}
}

// syncLbmonitorBindings adds/removes gslbservice_lbmonitor_binding entries to match newSet.
func (r *GslbserviceResource) syncLbmonitorBindings(ctx context.Context, servicename string, oldSet, newSet types.Set) diag.Diagnostics {
	var diags diag.Diagnostics

	var oldBindings, newBindings []LbmonitorbindingModel
	if !oldSet.IsNull() && !oldSet.IsUnknown() {
		diags.Append(oldSet.ElementsAs(ctx, &oldBindings, false)...)
	}
	if !newSet.IsNull() && !newSet.IsUnknown() {
		diags.Append(newSet.ElementsAs(ctx, &newBindings, false)...)
	}
	if diags.HasError() {
		return diags
	}

	newNames := make(map[string]bool)
	for _, b := range newBindings {
		newNames[b.MonitorName.ValueString()] = true
	}
	oldNames := make(map[string]bool)
	for _, b := range oldBindings {
		oldNames[b.MonitorName.ValueString()] = true
	}

	// Remove bindings that are no longer present.
	for _, b := range oldBindings {
		mn := b.MonitorName.ValueString()
		if mn == "" || newNames[mn] {
			continue
		}
		args := []string{fmt.Sprintf("monitor_name:%s", mn)}
		if err := r.client.DeleteResourceWithArgs("gslbservice_lbmonitor_binding", servicename, args); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to delete lbmonitor binding %s from gslbservice %s: %s", mn, servicename, err))
			return diags
		}
	}

	// Add bindings that are new.
	for _, b := range newBindings {
		mn := b.MonitorName.ValueString()
		if mn == "" || oldNames[mn] {
			continue
		}
		bind := gslb.Gslbservicemonitorbinding{Servicename: servicename}
		if !b.Weight.IsNull() && !b.Weight.IsUnknown() {
			bind.Weight = uint32(b.Weight.ValueInt64())
		}
		bind.Monitorname = mn
		if !b.Monstate.IsNull() && !b.Monstate.IsUnknown() {
			bind.Monstate = b.Monstate.ValueString()
		}
		if _, err := r.client.UpdateResource("gslbservice_lbmonitor_binding", servicename, bind); err != nil {
			diags.AddError("Client Error", fmt.Sprintf("Unable to add lbmonitor binding %s to gslbservice %s: %s", mn, servicename, err))
			return diags
		}
	}

	return diags
}

// readLbmonitorBindings populates data.Lbmonitorbinding from the ADC.
func (r *GslbserviceResource) readLbmonitorBindings(ctx context.Context, servicename string, data *GslbserviceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics

	bindings, _ := r.client.FindResourceArray("gslbservice_lbmonitor_binding", servicename)
	elems := make([]LbmonitorbindingModel, 0, len(bindings))
	for _, m := range bindings {
		e := LbmonitorbindingModel{
			Weight:      types.Int64Null(),
			MonitorName: types.StringNull(),
			Monstate:    types.StringNull(),
		}
		if v, ok := m["weight"]; ok && v != nil {
			if iv, err := utils.ConvertToInt64(v); err == nil {
				e.Weight = types.Int64Value(iv)
			}
		}
		if v, ok := m["monitor_name"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.MonitorName = types.StringValue(s)
			}
		}
		if v, ok := m["monstate"]; ok && v != nil {
			if s, isStr := v.(string); isStr {
				e.Monstate = types.StringValue(s)
			}
		}
		elems = append(elems, e)
	}

	setVal, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: lbmonitorbindingAttrTypes}, elems)
	diags.Append(d...)
	if diags.HasError() {
		return diags
	}
	data.Lbmonitorbinding = setVal
	return diags
}
