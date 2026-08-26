package nspbr

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
var _ resource.Resource = &NspbrResource{}
var _ resource.ResourceWithConfigure = (*NspbrResource)(nil)
var _ resource.ResourceWithImportState = (*NspbrResource)(nil)

func NewNspbrResource() resource.Resource {
	return &NspbrResource{}
}

// NspbrResource defines the resource implementation.
type NspbrResource struct {
	client *service.NitroClient
}

func (r *NspbrResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NspbrResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nspbr"
}

func (r *NspbrResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NspbrResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NspbrResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating nspbr resource")

	// nspbr is a named resource created via POST add. The `state` attribute
	// (ENABLED/DISABLED) is included in the add payload, mirroring SDK v2.
	nspbr := nspbrGetThePayloadFromthePlan(ctx, &data)
	nspbrName := data.Name.ValueString()

	_, err := r.client.AddResource(service.Nspbr.Type(), nspbrName, &nspbr)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create nspbr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created nspbr resource")

	// Set ID for the resource before reading state (plain name value, matches SDK v2 d.SetId(name))
	data.Id = types.StringValue(nspbrName)

	// Read the updated state back
	if !r.readNspbrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspbr not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspbrResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NspbrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading nspbr resource")

	found := r.readNspbrFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NspbrResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state NspbrResourceModel

	// Read Terraform prior state to preserve ID and detect changes
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read config to detect attributes removed from config (to be unset)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating nspbr resource")

	// Detect changes on updatable attributes (name and iptunnelname are ForceNew;
	// state is handled via the enable/disable action).
	hasChange := false
	attributesToUnset := []string{}
	updatableAttrs := []struct {
		name  string
		equal bool
	}{
		{"action", data.Action.Equal(state.Action)},
		{"destip", data.Destip.Equal(state.Destip)},
		{"destipdataset", data.Destipdataset.Equal(state.Destipdataset)},
		{"destipop", data.Destipop.Equal(state.Destipop)},
		{"destipval", data.Destipval.Equal(state.Destipval)},
		{"destport", data.Destport.Equal(state.Destport)},
		{"destportdataset", data.Destportdataset.Equal(state.Destportdataset)},
		{"destportop", data.Destportop.Equal(state.Destportop)},
		{"destportval", data.Destportval.Equal(state.Destportval)},
		{"detail", data.Detail.Equal(state.Detail)},
		{"interface", data.Interface.Equal(state.Interface)},
		{"monitor", data.Monitor.Equal(state.Monitor)},
		{"nexthop", data.Nexthop.Equal(state.Nexthop)},
		{"nexthopval", data.Nexthopval.Equal(state.Nexthopval)},
		{"ownergroup", data.Ownergroup.Equal(state.Ownergroup)},
		{"priority", data.Priority.Equal(state.Priority)},
		{"protocol", data.Protocol.Equal(state.Protocol)},
		{"protocolnumber", data.Protocolnumber.Equal(state.Protocolnumber)},
		{"srcip", data.Srcip.Equal(state.Srcip)},
		{"srcipdataset", data.Srcipdataset.Equal(state.Srcipdataset)},
		{"srcipop", data.Srcipop.Equal(state.Srcipop)},
		{"srcipval", data.Srcipval.Equal(state.Srcipval)},
		{"srcmac", data.Srcmac.Equal(state.Srcmac)},
		{"srcport", data.Srcport.Equal(state.Srcport)},
		{"srcportdataset", data.Srcportdataset.Equal(state.Srcportdataset)},
		{"srcportop", data.Srcportop.Equal(state.Srcportop)},
		{"srcportval", data.Srcportval.Equal(state.Srcportval)},
		{"targettd", data.Targettd.Equal(state.Targettd)},
		{"td", data.Td.Equal(state.Td)},
		{"vlan", data.Vlan.Equal(state.Vlan)},
		{"vxlan", data.Vxlan.Equal(state.Vxlan)},
		{"vxlanvlanmap", data.Vxlanvlanmap.Equal(state.Vxlanvlanmap)},
	}
	for _, a := range updatableAttrs {
		if !a.equal {
			tflog.Debug(ctx, fmt.Sprintf("%s has changed for nspbr, starting update", a.name))
			hasChange = true
		}
	}

	// Unsettable attributes: when removed from config, revert them to the NITRO
	// default via the unset action instead of sending an update.
	if !data.Srcmacmask.Equal(state.Srcmacmask) {
		tflog.Debug(ctx, "srcmacmask has changed for nspbr")
		if config.Srcmacmask.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "srcmacmask")
		} else {
			hasChange = true
		}
	}

	// State change is applied via the enable/disable action, not the update call.
	if !data.State.Equal(state.State) && !data.State.IsNull() && !data.State.IsUnknown() {
		if err := r.doNspbrStateChange(ctx, &data); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error enabling/disabling nspbr %s: %s", data.Name.ValueString(), err))
			return
		}
	}

	if hasChange {
		nspbr := nspbrGetTheUpdatablePayloadFromThePlan(ctx, &data)
		err := r.client.UpdateUnnamedResource(service.Nspbr.Type(), &nspbr)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update nspbr, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated nspbr resource")
	} else {
		tflog.Debug(ctx, "No updatable changes detected for nspbr resource, skipping update call")
	}

	// Unset attributes that were removed from config so the appliance reverts
	// them to their defaults.
	unsetIdPayload := map[string]interface{}{
		"name": data.Name.ValueString(),
	}
	if err := utils.ExecuteUnset(r.client, service.Nspbr.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset nspbr attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readNspbrFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "nspbr not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NspbrResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NspbrResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting nspbr resource")

	nspbrName := data.Id.ValueString()
	err := r.client.DeleteResource(service.Nspbr.Type(), nspbrName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete nspbr, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted nspbr resource")
}

// doNspbrStateChange enables or disables the PBR via the NITRO action endpoint
// (mirrors SDK v2 doNspbrStateChange).
func (r *NspbrResource) doNspbrStateChange(ctx context.Context, data *NspbrResourceModel) error {
	tflog.Debug(ctx, "In doNspbrStateChange")

	// A fresh struct with only the name is used; ActOnResource fails on superfluous attributes.
	nspbr := ns.Nspbr{
		Name: data.Name.ValueString(),
	}
	newstate := data.State.ValueString()

	switch newstate {
	case "ENABLED":
		return r.client.ActOnResource(service.Nspbr.Type(), nspbr, "enable")
	case "DISABLED":
		return r.client.ActOnResource(service.Nspbr.Type(), nspbr, "disable")
	default:
		return fmt.Errorf("%q is not a valid state. Use (\"ENABLED\", \"DISABLED\")", newstate)
	}
}

// readNspbrFromApi reads the nspbr from NITRO by name (the plain-value ID) and
// maps it onto the model. Returns false (without error) if the resource is gone.
func (r *NspbrResource) readNspbrFromApi(ctx context.Context, data *NspbrResourceModel, diags *diag.Diagnostics) bool {
	nspbrName := data.Id.ValueString()

	getResponseData, err := r.client.FindResource(service.Nspbr.Type(), nspbrName)
	if err != nil {
		if utils.IsNotFoundError(err) {
			return false
		}
		diags.AddError("Client Error", fmt.Sprintf("Unable to read nspbr, got error: %s", err))
		return false
	}

	nspbrSetAttrFromGet(ctx, data, getResponseData)

	return true
}
