package Interface

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/network"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &InterfaceResource{}
var _ resource.ResourceWithConfigure = (*InterfaceResource)(nil)
var _ resource.ResourceWithImportState = (*InterfaceResource)(nil)

func NewInterfaceResource() resource.Resource {
	return &InterfaceResource{}
}

// InterfaceResource defines the resource implementation.
type InterfaceResource struct {
	client *service.NitroClient
}

func (r *InterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *InterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface"
}

func (r *InterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *InterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data InterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating interface resource")

	interfaceId := data.Interfaceid.ValueString()

	// Build the payload from the plan (skips unknown/unconfigured Computed attributes).
	Interface := interfaceGetThePayloadFromthePlan(ctx, &data)

	// The Interface resource exposes no NITRO "add" verb; it is configured with an
	// unnamed "update" (PUT) call, exactly like the SDK v2 resource.
	_, err := r.client.UpdateResource(service.Interface.Type(), "", &Interface)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create interface %s, got error: %s", interfaceId, err))
		return
	}

	// Apply the enable/disable action if the state attribute is configured.
	if !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doInterfaceStateChange(interfaceId, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set state for interface %s, got error: %s", interfaceId, err))
			return
		}
	}

	tflog.Trace(ctx, "Created interface resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(interfaceId)

	// Read the updated state back
	if !r.readInterfaceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "interface not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data InterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading interface resource")

	found := r.readInterfaceFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *InterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state InterfaceResourceModel

	// Read Terraform prior state to detect changes and preserve the ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating interface resource")

	interfaceId := data.Interfaceid.ValueString()

	// Build a payload containing only the attributes that actually changed, mirroring
	// the SDK v2 update behaviour (avoids re-applying negotiated/computed values).
	Interface := network.Interface{
		Id: interfaceId,
	}
	hasChange := false

	if !data.Autoneg.Equal(state.Autoneg) {
		hasChange = true
		if !data.Autoneg.IsNull() && !data.Autoneg.IsUnknown() {
			Interface.Autoneg = data.Autoneg.ValueString()
		}
	}
	if !data.Bandwidthhigh.Equal(state.Bandwidthhigh) {
		hasChange = true
		if !data.Bandwidthhigh.IsNull() && !data.Bandwidthhigh.IsUnknown() {
			Interface.Bandwidthhigh = intPtrFromInt64(data.Bandwidthhigh.ValueInt64())
		}
	}
	if !data.Bandwidthnormal.Equal(state.Bandwidthnormal) {
		hasChange = true
		if !data.Bandwidthnormal.IsNull() && !data.Bandwidthnormal.IsUnknown() {
			Interface.Bandwidthnormal = intPtrFromInt64(data.Bandwidthnormal.ValueInt64())
		}
	}
	if !data.Duplex.Equal(state.Duplex) {
		hasChange = true
		if !data.Duplex.IsNull() && !data.Duplex.IsUnknown() {
			Interface.Duplex = data.Duplex.ValueString()
		}
	}
	if !data.Flowctl.Equal(state.Flowctl) {
		hasChange = true
		if !data.Flowctl.IsNull() && !data.Flowctl.IsUnknown() {
			Interface.Flowctl = data.Flowctl.ValueString()
		}
	}
	if !data.Haheartbeat.Equal(state.Haheartbeat) {
		hasChange = true
		if !data.Haheartbeat.IsNull() && !data.Haheartbeat.IsUnknown() {
			Interface.Haheartbeat = data.Haheartbeat.ValueString()
		}
	}
	if !data.Hamonitor.Equal(state.Hamonitor) {
		hasChange = true
		if !data.Hamonitor.IsNull() && !data.Hamonitor.IsUnknown() {
			Interface.Hamonitor = data.Hamonitor.ValueString()
		}
	}
	if !data.Ifalias.Equal(state.Ifalias) {
		hasChange = true
		if !data.Ifalias.IsNull() && !data.Ifalias.IsUnknown() {
			Interface.Ifalias = data.Ifalias.ValueString()
		}
	}
	if !data.Lacpkey.Equal(state.Lacpkey) {
		hasChange = true
		if !data.Lacpkey.IsNull() && !data.Lacpkey.IsUnknown() {
			Interface.Lacpkey = intPtrFromInt64(data.Lacpkey.ValueInt64())
		}
	}
	if !data.Lacpmode.Equal(state.Lacpmode) {
		hasChange = true
		if !data.Lacpmode.IsNull() && !data.Lacpmode.IsUnknown() {
			Interface.Lacpmode = data.Lacpmode.ValueString()
		}
	}
	if !data.Lacppriority.Equal(state.Lacppriority) {
		hasChange = true
		if !data.Lacppriority.IsNull() && !data.Lacppriority.IsUnknown() {
			Interface.Lacppriority = intPtrFromInt64(data.Lacppriority.ValueInt64())
		}
	}
	if !data.Lacptimeout.Equal(state.Lacptimeout) {
		hasChange = true
		if !data.Lacptimeout.IsNull() && !data.Lacptimeout.IsUnknown() {
			Interface.Lacptimeout = data.Lacptimeout.ValueString()
		}
	}
	if !data.Lagtype.Equal(state.Lagtype) {
		hasChange = true
		if !data.Lagtype.IsNull() && !data.Lagtype.IsUnknown() {
			Interface.Lagtype = data.Lagtype.ValueString()
		}
	}
	if !data.Linkredundancy.Equal(state.Linkredundancy) {
		hasChange = true
		if !data.Linkredundancy.IsNull() && !data.Linkredundancy.IsUnknown() {
			Interface.Linkredundancy = data.Linkredundancy.ValueString()
		}
	}
	if !data.Lldpmode.Equal(state.Lldpmode) {
		hasChange = true
		if !data.Lldpmode.IsNull() && !data.Lldpmode.IsUnknown() {
			Interface.Lldpmode = data.Lldpmode.ValueString()
		}
	}
	if !data.Lrsetpriority.Equal(state.Lrsetpriority) {
		hasChange = true
		if !data.Lrsetpriority.IsNull() && !data.Lrsetpriority.IsUnknown() {
			Interface.Lrsetpriority = intPtrFromInt64(data.Lrsetpriority.ValueInt64())
		}
	}
	if !data.Mtu.Equal(state.Mtu) {
		hasChange = true
		if !data.Mtu.IsNull() && !data.Mtu.IsUnknown() {
			Interface.Mtu = intPtrFromInt64(data.Mtu.ValueInt64())
		}
	}
	if !data.Ringsize.Equal(state.Ringsize) {
		hasChange = true
		if !data.Ringsize.IsNull() && !data.Ringsize.IsUnknown() {
			Interface.Ringsize = intPtrFromInt64(data.Ringsize.ValueInt64())
		}
	}
	if !data.Ringtype.Equal(state.Ringtype) {
		hasChange = true
		if !data.Ringtype.IsNull() && !data.Ringtype.IsUnknown() {
			Interface.Ringtype = data.Ringtype.ValueString()
		}
	}
	if !data.Speed.Equal(state.Speed) {
		hasChange = true
		if !data.Speed.IsNull() && !data.Speed.IsUnknown() {
			Interface.Speed = data.Speed.ValueString()
		}
	}
	if !data.Tagall.Equal(state.Tagall) {
		hasChange = true
		if !data.Tagall.IsNull() && !data.Tagall.IsUnknown() {
			Interface.Tagall = data.Tagall.ValueString()
		}
	}
	if !data.Throughput.Equal(state.Throughput) {
		hasChange = true
		if !data.Throughput.IsNull() && !data.Throughput.IsUnknown() {
			Interface.Throughput = intPtrFromInt64(data.Throughput.ValueInt64())
		}
	}
	if !data.Trunk.Equal(state.Trunk) {
		hasChange = true
		if !data.Trunk.IsNull() && !data.Trunk.IsUnknown() {
			Interface.Trunk = data.Trunk.ValueString()
		}
	}
	if !data.Trunkallowedvlan.Equal(state.Trunkallowedvlan) {
		hasChange = true
		if !data.Trunkallowedvlan.IsNull() && !data.Trunkallowedvlan.IsUnknown() {
			var trunkallowedvlanList []string
			data.Trunkallowedvlan.ElementsAs(ctx, &trunkallowedvlanList, false)
			Interface.Trunkallowedvlan = trunkallowedvlanList
		}
	}
	if !data.Trunkmode.Equal(state.Trunkmode) {
		hasChange = true
		if !data.Trunkmode.IsNull() && !data.Trunkmode.IsUnknown() {
			Interface.Trunkmode = data.Trunkmode.ValueString()
		}
	}

	if hasChange {
		_, err := r.client.UpdateResource(service.Interface.Type(), "", &Interface)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update interface %s, got error: %s", interfaceId, err))
			return
		}
		tflog.Trace(ctx, "Updated interface resource")
	} else {
		tflog.Debug(ctx, "No changes detected for interface resource, skipping update")
	}

	// Apply the enable/disable action if the state attribute changed.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doInterfaceStateChange(interfaceId, data.State.ValueString()); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set state for interface %s, got error: %s", interfaceId, err))
			return
		}
	}

	// Read the updated state back
	if !r.readInterfaceFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "interface not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *InterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data InterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// An interface is a hardware entity and cannot be deleted through NITRO (the
	// resource exposes no delete verb). Mirroring the SDK v2 resource, we only
	// remove it from Terraform state; the framework does this automatically.
	tflog.Trace(ctx, "Deleted interface resource from state (interface hardware is not removed)")
}

// Helper function to read interface data from API. Returns false if the interface
// is not present on the appliance.
func (r *InterfaceResource) readInterfaceFromApi(ctx context.Context, data *InterfaceResourceModel, diags *diag.Diagnostics) bool {
	interfaceId := data.Id.ValueString()

	array, err := r.client.FindAllResources(service.Interface.Type())
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read interface %s, got error: %s", interfaceId, err))
		return false
	}

	foundIndex := -1
	for i, item := range array {
		if id, ok := item["id"].(string); ok && id == interfaceId {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		return false
	}

	interfaceSetAttrFromGet(ctx, data, array[foundIndex])

	return true
}

// doInterfaceStateChange enables or disables the interface, mirroring the SDK v2
// resource's use of the NITRO enable/disable actions.
func (r *InterfaceResource) doInterfaceStateChange(interfaceId string, newstate string) error {
	Interface := network.Interface{
		Id: interfaceId,
	}

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Interface.Type(), Interface, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Interface.Type(), Interface, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use \"ENABLED\" or \"DISABLED\"", newstate)
	}
}

// intPtrFromInt64 returns a pointer to an int converted from an int64, for use with
// the optional pointer fields of the NITRO Interface struct.
func intPtrFromInt64(v int64) *int {
	i := int(v)
	return &i
}
