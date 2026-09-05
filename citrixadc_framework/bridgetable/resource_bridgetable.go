package bridgetable

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &BridgetableResource{}
var _ resource.ResourceWithConfigure = (*BridgetableResource)(nil)
var _ resource.ResourceWithImportState = (*BridgetableResource)(nil)

func NewBridgetableResource() resource.Resource {
	return &BridgetableResource{}
}

// BridgetableResource defines the resource implementation.
type BridgetableResource struct {
	client *service.NitroClient
}

func (r *BridgetableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *BridgetableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_bridgetable"
}

func (r *BridgetableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// bridgetableName builds the backward-compatible composite name/ID "mac,vxlan,vtep".
func bridgetableName(data *BridgetableResourceModel) string {
	return fmt.Sprintf("%s,%d,%s", data.Mac.ValueString(), data.Vxlan.ValueInt64(), data.Vtep.ValueString())
}

func (r *BridgetableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data BridgetableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating bridgetable resource")

	// Add the bridge table entry (bridgeage excluded - applied separately below).
	bridgetable := bridgetableGetThePayloadFromthePlan(ctx, &data)
	name := bridgetableName(&data)
	_, err := r.client.AddResource(service.Bridgetable.Type(), name, &bridgetable)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create bridgetable, got error: %s", err))
		return
	}

	// bridgeage is a table-wide setting applied via an unnamed update.
	if !data.Bridgeage.IsNull() && !data.Bridgeage.IsUnknown() {
		bridgeagePayload := bridgetableGetTheBridgeagePayload(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Bridgetable.Type(), &bridgeagePayload); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to set bridgeage for bridgetable, got error: %s", err))
			return
		}
	}

	// Backward-compatible composite ID.
	data.Id = types.StringValue(name)

	tflog.Trace(ctx, "Created bridgetable resource")

	// Read the updated state back
	if !r.readBridgetableFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "bridgetable not found immediately after create")
		}
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgetableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data BridgetableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading bridgetable resource")

	found := r.readBridgetableFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *BridgetableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, config, state BridgetableResourceModel

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

	tflog.Debug(ctx, "Updating bridgetable resource")

	hasChange := false
	attributesToUnset := []string{}

	// Only bridgeage is updateable in place; all other attributes are RequiresReplace.
	if !data.Bridgeage.Equal(state.Bridgeage) {
		tflog.Debug(ctx, "bridgeage has changed for bridgetable")
		if config.Bridgeage.IsNull() { // removed from config -> unset it
			attributesToUnset = append(attributesToUnset, "bridgeage")
		} else {
			hasChange = true
		}
	}

	if hasChange {
		bridgeagePayload := bridgetableGetTheBridgeagePayload(ctx, &data)
		if err := r.client.UpdateUnnamedResource(service.Bridgetable.Type(), &bridgeagePayload); err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update bridgetable, got error: %s", err))
			return
		}
		tflog.Trace(ctx, "Updated bridgetable resource")
	} else {
		tflog.Debug(ctx, "No changes detected for bridgetable resource, skipping update")
	}

	// Unset attributes removed from config so the appliance reverts them to
	// their defaults. bridgeage is a table-wide setting, so the unset carries no
	// identifying key (matches the NITRO unset payload: {"bridgeage":true}).
	unsetIdPayload := map[string]interface{}{}
	if err := utils.ExecuteUnset(r.client, service.Bridgetable.Type(), unsetIdPayload, attributesToUnset); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to unset bridgetable attributes, got error: %s", err))
		return
	}

	// Read the updated state back
	if !r.readBridgetableFromApi(ctx, &data, &resp.Diagnostics) {
		if !resp.Diagnostics.HasError() {
			resp.Diagnostics.AddError("Client Error", "bridgetable not found immediately after update")
		}
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *BridgetableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data BridgetableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting bridgetable resource")

	// Delete requires the identifying args (mac, vxlan, vtep, devicevlan). mac and
	// vtep are pre-escaped because a MAC address contains ':' which would otherwise
	// break the "key:value" arg parsing (matches SDK v2).
	argsMap := make(map[string]string)
	argsMap["mac"] = url.QueryEscape(data.Mac.ValueString())
	argsMap["vtep"] = url.QueryEscape(data.Vtep.ValueString())
	argsMap["vxlan"] = strconv.Itoa(int(data.Vxlan.ValueInt64()))
	argsMap["devicevlan"] = strconv.Itoa(int(data.Devicevlan.ValueInt64()))

	if err := r.client.DeleteResourceWithArgsMap(service.Bridgetable.Type(), "", argsMap); err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete bridgetable, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted bridgetable resource")
}

// readBridgetableFromApi finds the bridgetable entry matching the resource's
// identity (mac, vxlan, vtep) and maps it onto the model. Returns false when the
// entry no longer exists so the caller can remove it from state.
func (r *BridgetableResource) readBridgetableFromApi(ctx context.Context, data *BridgetableResourceModel, diags *diag.Diagnostics) bool {
	findParams := service.FindParams{
		ResourceType: service.Bridgetable.Type(),
	}
	dataArray, err := r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read bridgetable, got error: %s", err))
		return false
	}
	if len(dataArray) == 0 {
		return false
	}

	// Derive the identity keys from the ID ("mac,vxlan,vtep") so Read works after
	// import (where only the ID is populated) as well as on refresh.
	idSlice := strings.SplitN(data.Id.ValueString(), ",", 3)
	if len(idSlice) != 3 {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse bridgetable ID %q, expected \"mac,vxlan,vtep\"", data.Id.ValueString()))
		return false
	}
	mac := idSlice[0]
	vxlan := idSlice[1]
	vtep := idSlice[2]

	foundIndex := -1
	for i, entry := range dataArray {
		match := true
		if fmt.Sprintf("%v", entry["mac"]) != mac {
			match = false
		}
		if fmt.Sprintf("%v", entry["vxlan"]) != vxlan {
			match = false
		}
		if fmt.Sprintf("%v", entry["vtep"]) != vtep {
			match = false
		}
		if match {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		tflog.Warn(ctx, fmt.Sprintf("bridgetable %s not found in array", data.Id.ValueString()))
		return false
	}

	bridgetableSetAttrFromGet(ctx, data, dataArray[foundIndex])

	return true
}
