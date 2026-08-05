package route

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &RouteResource{}
var _ resource.ResourceWithConfigure = (*RouteResource)(nil)
var _ resource.ResourceWithImportState = (*RouteResource)(nil)

func NewRouteResource() resource.Resource {
	return &RouteResource{}
}

// RouteResource defines the resource implementation.
type RouteResource struct {
	client *service.NitroClient
}

func (r *RouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *RouteResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route"
}

func (r *RouteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *RouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data RouteResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating route resource")

	route := routeGetThePayloadFromthePlan(ctx, &data)

	// Named/array resource - use AddResource (NITRO add is HTTP POST).
	_, err := r.client.AddResource(service.Route.Type(), route.Network, &route)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create route, got error: %s", err))
		return
	}

	// Convenience block (preserved from SDK v2): optionally delete the default
	// static route (0.0.0.0/0.0.0.0) after adding this route, remembering the
	// original gateway so it can be restored on destroy.
	if data.DeleteDefaultRoute.ValueBool() {
		tflog.Debug(ctx, "delete_default_route is true, finding and deleting default route (0.0.0.0/0.0.0.0)")
		findParams := service.FindParams{
			ResourceType: service.Route.Type(),
		}
		dataArray, findErr := r.client.FindResourceArrayWithParams(findParams)
		if findErr != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("delete_default_route is true but failed to list routes: %s", findErr.Error()))
			return
		}
		defaultRouteFound := false
		for _, rt := range dataArray {
			routeNetwork, _ := rt["network"].(string)
			routeNetmask, _ := rt["netmask"].(string)
			routeGateway, _ := rt["gateway"].(string)
			if routeNetwork == "0.0.0.0" && routeNetmask == "0.0.0.0" {
				defaultRouteFound = true
				delArgs := map[string]string{
					"network": url.QueryEscape("0.0.0.0"),
					"netmask": url.QueryEscape("0.0.0.0"),
					"gateway": url.QueryEscape(routeGateway),
				}
				delErr := r.client.DeleteResourceWithArgsMap(service.Route.Type(), "", delArgs)
				if delErr != nil {
					resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error deleting default route (gateway %s): %s", routeGateway, delErr.Error()))
					return
				}
				// Save the original gateway so we can restore it on destroy.
				data.OriginalDefaultGateway = types.StringValue(routeGateway)
			}
		}
		if !defaultRouteFound {
			tflog.Warn(ctx, "delete_default_route is true but no default route (0.0.0.0/0.0.0.0) was found")
		}
	}

	// original_default_gateway is Computed; ensure it is always known after apply.
	if data.OriginalDefaultGateway.IsUnknown() || data.OriginalDefaultGateway.IsNull() {
		data.OriginalDefaultGateway = types.StringValue("")
	}

	// Preserve the SDK v2 ID scheme: network__netmask__gateway
	data.Id = types.StringValue(data.Network.ValueString() + "__" + data.Netmask.ValueString() + "__" + data.Gateway.ValueString())

	tflog.Trace(ctx, "Created route resource")

	// Read the updated state back
	if !r.readRouteFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "route not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data RouteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading route resource")

	found := r.readRouteFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *RouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state RouteResourceModel

	// Read Terraform prior state (to preserve ID and detect changes)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating route resource")

	route := network.Route{}
	hasChange := false
	// Only send fields that are genuinely changed in config. When an
	// Optional+Computed attribute is absent from config, the framework marks it
	// Unknown ("known after apply") on an update plan; treating that as a change
	// would push stale zero/empty values (e.g. vlan=0) that NITRO rejects
	// (errorcode 278, "Invalid argument [vlan]"). Guarding with !IsUnknown()
	// mirrors the SDK v2 d.HasChange() semantics, which only fired on real
	// config changes.
	if !data.Protocol.IsUnknown() && !data.Protocol.Equal(state.Protocol) {
		var protocolList []string
		data.Protocol.ElementsAs(ctx, &protocolList, false)
		route.Protocol = protocolList
		hasChange = true
	}
	if !data.Advertise.IsUnknown() && !data.Advertise.Equal(state.Advertise) {
		route.Advertise = data.Advertise.ValueString()
		hasChange = true
	}
	if !data.Cost.IsUnknown() && !data.Cost.Equal(state.Cost) {
		route.Cost = utils.IntPtr(int(data.Cost.ValueInt64()))
		hasChange = true
	}
	if !data.Cost1.IsUnknown() && !data.Cost1.Equal(state.Cost1) {
		route.Cost1 = utils.IntPtr(int(data.Cost1.ValueInt64()))
		hasChange = true
	}
	if !data.Detail.IsUnknown() && !data.Detail.Equal(state.Detail) {
		route.Detail = data.Detail.ValueBool()
		hasChange = true
	}
	if !data.Distance.IsUnknown() && !data.Distance.Equal(state.Distance) {
		route.Distance = utils.IntPtr(int(data.Distance.ValueInt64()))
		hasChange = true
	}
	if !data.Monitor.IsUnknown() && !data.Monitor.Equal(state.Monitor) {
		route.Monitor = data.Monitor.ValueString()
		hasChange = true
	}
	if !data.Msr.IsUnknown() && !data.Msr.Equal(state.Msr) {
		route.Msr = data.Msr.ValueString()
		hasChange = true
	}
	if !data.Ownergroup.IsUnknown() && !data.Ownergroup.Equal(state.Ownergroup) {
		route.Ownergroup = data.Ownergroup.ValueString()
		hasChange = true
	}
	if !data.Routetype.IsUnknown() && !data.Routetype.Equal(state.Routetype) {
		route.Routetype = data.Routetype.ValueString()
		hasChange = true
	}
	if !data.Td.IsUnknown() && !data.Td.Equal(state.Td) {
		route.Td = utils.IntPtr(int(data.Td.ValueInt64()))
		hasChange = true
	}
	if !data.Vlan.IsUnknown() && !data.Vlan.Equal(state.Vlan) {
		route.Vlan = utils.IntPtr(int(data.Vlan.ValueInt64()))
		hasChange = true
	}
	if !data.Weight.IsUnknown() && !data.Weight.Equal(state.Weight) {
		route.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
		hasChange = true
	}

	if hasChange {
		// network, netmask, gateway are mandatory in the UPDATE payload.
		route.Network = data.Network.ValueString()
		route.Netmask = data.Netmask.ValueString()
		route.Gateway = data.Gateway.ValueString()
		err := r.client.UpdateUnnamedResource(service.Route.Type(), &route)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update route, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated route resource")
	} else {
		tflog.Debug(ctx, "No changes detected for route resource, skipping update")
	}

	// Read the updated state back
	if !r.readRouteFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "route not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *RouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data RouteResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting route resource")

	argsMap := make(map[string]string)
	argsMap["network"] = url.QueryEscape(data.Network.ValueString())
	argsMap["netmask"] = url.QueryEscape(data.Netmask.ValueString())
	argsMap["gateway"] = url.QueryEscape(data.Gateway.ValueString())
	if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() && data.Ownergroup.ValueString() != "" {
		argsMap["ownergroup"] = url.QueryEscape(data.Ownergroup.ValueString())
	}

	// Convenience block (preserved from SDK v2): restore the original default
	// route before deleting the managed route.
	if data.DeleteDefaultRoute.ValueBool() {
		originalGw := data.OriginalDefaultGateway.ValueString()
		if originalGw != "" {
			tflog.Debug(ctx, fmt.Sprintf("Restoring default route (0.0.0.0/0.0.0.0) with gateway %s", originalGw))
			defaultRoute := network.Route{
				Network: "0.0.0.0",
				Netmask: "0.0.0.0",
				Gateway: originalGw,
			}
			_, restoreErr := r.client.AddResource(service.Route.Type(), "0.0.0.0", &defaultRoute)
			if restoreErr != nil {
				resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error restoring default route (0.0.0.0/0.0.0.0) with gateway %s: %s", originalGw, restoreErr.Error()))
				return
			}
		}
	}

	err := r.client.DeleteResourceWithArgsMap(service.Route.Type(), "", argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete route, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted route resource")
}

// readRouteFromApi lists all routes and locates the one identified by this
// resource's ID (network__netmask__gateway), matching ownergroup when set.
// Returns false when the route no longer exists.
func (r *RouteResource) readRouteFromApi(ctx context.Context, data *RouteResourceModel, diags *diag.Diagnostics) bool {
	findParams := service.FindParams{
		ResourceType: service.Route.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read route, got error: %s", err))
		return false
	}
	if len(dataArray) == 0 {
		return false
	}

	idSlice := strings.SplitN(data.Id.ValueString(), "__", 3)
	if len(idSlice) != 3 {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse route ID %q; expected network__netmask__gateway", data.Id.ValueString()))
		return false
	}
	networkId := idSlice[0]
	netmaskId := idSlice[1]
	gatewayId := idSlice[2]

	foundIndex := -1
	for i, rt := range dataArray {
		match := true
		if rt["network"] != networkId {
			match = false
		}
		if rt["netmask"] != netmaskId {
			match = false
		}
		if rt["gateway"] != gatewayId {
			match = false
		}
		if !data.Ownergroup.IsNull() && !data.Ownergroup.IsUnknown() && data.Ownergroup.ValueString() != "" {
			if rt["ownergroup"] != data.Ownergroup.ValueString() {
				match = false
			}
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	routeSetAttrFromGet(ctx, data, dataArray[foundIndex])
	return true
}
