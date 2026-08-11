package route6

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &Route6Resource{}
var _ resource.ResourceWithConfigure = (*Route6Resource)(nil)
var _ resource.ResourceWithImportState = (*Route6Resource)(nil)

func NewRoute6Resource() resource.Resource {
	return &Route6Resource{}
}

// Route6Resource defines the resource implementation.
type Route6Resource struct {
	client *service.NitroClient
}

func (r *Route6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// SDK v2 ID scheme is the plain network value; passthrough it into the id
	// attribute so Read can resolve the resource by network.
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *Route6Resource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route6"
}

func (r *Route6Resource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *Route6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data Route6ResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating route6 resource")

	route6 := route6GetThePayloadFromtheConfig(ctx, &data)

	// route6 add is POST /nitro/v1/config/route6 (named-style resource, no name
	// in the URL). SDK v2 used AddResource here as well.
	_, err := r.client.AddResource(service.Route6.Type(), "", &route6)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create route6, got error: %s", err))
		return
	}

	// SDK v2 ID scheme: d.SetId(network) — plain network value.
	data.Id = types.StringValue(data.Network.ValueString())

	tflog.Trace(ctx, "Created route6 resource")

	// Read the updated state back
	found := r.readRoute6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// The route was created but could not be read back; resolve any still
		// unknown Computed attributes (empty GET response nulls them) so the
		// apply does not fail with an inconsistent-result / unknown-value error.
		route6SetAttrFromGet(ctx, &data, map[string]interface{}{})
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Route6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data Route6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading route6 resource")

	found := r.readRoute6FromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !found {
		// SDK v2 cleared state (d.SetId("")) when the route could not be found.
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *Route6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state Route6ResourceModel

	// Read Terraform prior state to preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read raw config to detect attributes removed from configuration (unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve the ID (plain network) from prior state.
	data.Id = state.Id

	tflog.Debug(ctx, "Updating route6 resource")

	// Detect whether any mutable attribute changed (network/mgmt are
	// RequiresReplace so they never reach here). Attributes removed from config
	// that support NITRO unset are collected instead of being pushed.
	hasChange := false
	attributesToUnset := []string{}
	if !data.Advertise.Equal(state.Advertise) {
		hasChange = true
	}
	if !data.Cost.Equal(state.Cost) {
		if config.Cost.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "cost")
		} else {
			hasChange = true
		}
	}
	if !data.Distance.Equal(state.Distance) {
		if config.Distance.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "distance")
		} else {
			hasChange = true
		}
	}
	if !data.Gateway.Equal(state.Gateway) {
		hasChange = true
	}
	if !data.Monitor.Equal(state.Monitor) {
		hasChange = true
	}
	if !data.Msr.Equal(state.Msr) {
		if config.Msr.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "msr")
		} else {
			hasChange = true
		}
	}
	if !data.Td.Equal(state.Td) {
		hasChange = true
	}
	if !data.Vlan.Equal(state.Vlan) {
		hasChange = true
	}
	if !data.Vxlan.Equal(state.Vxlan) {
		hasChange = true
	}
	if !data.Weight.Equal(state.Weight) {
		if config.Weight.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "weight")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		// route6 update is PUT /nitro/v1/config/route6 (unnamed, full identifying
		// payload). Any non-ForceNew attribute change is pushed via
		// UpdateUnnamedResource.
		route6 := route6GetThePayloadForUpdate(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Route6.Type(), &route6)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update route6, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated route6 resource")
	} else {
		tflog.Debug(ctx, "No changes detected for route6 resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to their
	// defaults. route6 is an unnamed resource identified by network plus the
	// route-identity keys; include them so NITRO locates the exact route.
	unsetIdPayload := map[string]interface{}{
		"network": data.Network.ValueString(),
	}
	if !data.Gateway.IsNull() && !data.Gateway.IsUnknown() && data.Gateway.ValueString() != "" {
		unsetIdPayload["gateway"] = data.Gateway.ValueString()
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		unsetIdPayload["vlan"] = int(data.Vlan.ValueInt64())
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() && data.Vxlan.ValueInt64() != 0 {
		unsetIdPayload["vxlan"] = int(data.Vxlan.ValueInt64())
	}
	if !data.Td.IsNull() && !data.Td.IsUnknown() {
		unsetIdPayload["td"] = int(data.Td.ValueInt64())
	}
	if err := utils.ExecuteUnset(r.client, service.Route6.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset route6 attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	found := r.readRoute6FromApi(ctx, &data, &resp.Diagnostics)
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

func (r *Route6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data Route6ResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting route6 resource")

	// SDK v2 deleted via DeleteResourceWithArgs with the (url-escaped) network,
	// gateway, vlan, vxlan and ownergroup as disambiguating args. The internal
	// client does not escape the args, so escape strings here to match SDK v2.
	args := make([]string, 0)
	if !data.Network.IsNull() && !data.Network.IsUnknown() && data.Network.ValueString() != "" {
		args = append(args, fmt.Sprintf("network:%s", url.QueryEscape(data.Network.ValueString())))
	}
	if !data.Gateway.IsNull() && !data.Gateway.IsUnknown() && data.Gateway.ValueString() != "" {
		args = append(args, fmt.Sprintf("gateway:%s", url.QueryEscape(data.Gateway.ValueString())))
	}
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() && data.Vlan.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("vlan:%d", data.Vlan.ValueInt64()))
	}
	if !data.Vxlan.IsNull() && !data.Vxlan.IsUnknown() && data.Vxlan.ValueInt64() != 0 {
		args = append(args, fmt.Sprintf("vxlan:%d", data.Vxlan.ValueInt64()))
	}
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() && data.Ownergroup.ValueString() != "" {
		args = append(args, fmt.Sprintf("ownergroup:%s", url.QueryEscape(data.Ownergroup.ValueString())))
	}

	err := r.client.DeleteResourceWithArgs(service.Route6.Type(), "", args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete route6, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted route6 resource")
}

// readRoute6FromApi reads the route6 data from the ADC and maps it onto the
// model. It returns true when a matching route is found and false when it does
// not exist (so the caller can clear state). Only genuine response errors are
// surfaced via diags. Matching mirrors SDK v2: network (the ID) AND vlan.
func (r *Route6Resource) readRoute6FromApi(ctx context.Context, data *Route6ResourceModel, diags *diag.Diagnostics) bool {
	route6Network := data.Id.ValueString()

	wantedVlan := int64(0)
	if !data.Vlan.IsNull() && !data.Vlan.IsUnknown() {
		wantedVlan = data.Vlan.ValueInt64()
	}

	findParams := service.FindParams{
		ResourceType:             service.Route6.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		// SDK v2 cleared state on any find error; treat as not-found.
		tflog.Warn(ctx, fmt.Sprintf("Unable to list route6, clearing state for %s: %s", route6Network, err.Error()))
		return false
	}

	if len(dataArray) == 0 {
		tflog.Warn(ctx, "route6 does not exist; clearing state")
		return false
	}

	foundIndex := -1
	for i, v := range dataArray {
		n, ok := v["network"].(string)
		if !ok || n != route6Network {
			continue
		}
		if route6MapInt64(v, "vlan") != wantedVlan {
			continue
		}
		foundIndex = i
		break
	}
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("route6 with network %s and vlan %d not found; clearing state", route6Network, wantedVlan))
		return false
	}

	route6SetAttrFromGet(ctx, data, dataArray[foundIndex])

	return true
}

// route6MapInt64 extracts an integer field from a NITRO GET record, returning 0
// when the field is absent or unparseable (matches NITRO omitting a default-0).
func route6MapInt64(v map[string]interface{}, key string) int64 {
	if raw, ok := v[key]; ok && raw != nil {
		if i, err := utils.ConvertToInt64(raw); err == nil {
			return i
		}
	}
	return 0
}
