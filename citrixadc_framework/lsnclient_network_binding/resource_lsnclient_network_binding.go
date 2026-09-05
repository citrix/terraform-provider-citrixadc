package lsnclient_network_binding

import (
	"context"
	"fmt"
	"net/url"
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
var _ resource.Resource = &LsnclientNetworkBindingResource{}
var _ resource.ResourceWithConfigure = (*LsnclientNetworkBindingResource)(nil)
var _ resource.ResourceWithImportState = (*LsnclientNetworkBindingResource)(nil)

func NewLsnclientNetworkBindingResource() resource.Resource {
	return &LsnclientNetworkBindingResource{}
}

// LsnclientNetworkBindingResource defines the resource implementation.
type LsnclientNetworkBindingResource struct {
	client *service.NitroClient
}

func (r *LsnclientNetworkBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *LsnclientNetworkBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lsnclient_network_binding"
}

func (r *LsnclientNetworkBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *LsnclientNetworkBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data LsnclientNetworkBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating lsnclient_network_binding resource")
	lsnclient_network_binding := lsnclient_network_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Lsnclient_network_binding.Type(), &lsnclient_network_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create lsnclient_network_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created lsnclient_network_binding resource")

	// Set ID for the resource before reading state.
	// Use the legacy identity keys only (clientname,network) so the ID stays
	// backward-compatible with the SDK v2 "clientname,network" format and with
	// resource_id_mapping.json. netmask/td are not part of the binding identity
	// (and td is not even echoed by the GET response).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("clientname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Clientname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("network:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Network.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readLsnclientNetworkBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNetworkBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data LsnclientNetworkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading lsnclient_network_binding resource")

	r.readLsnclientNetworkBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNetworkBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state LsnclientNetworkBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating lsnclient_network_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		lsnclient_network_binding := lsnclient_network_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Lsnclient_network_binding.Type(), &lsnclient_network_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update lsnclient_network_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated lsnclient_network_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for lsnclient_network_binding resource, skipping update")
	}

	// Read the updated state back
	r.readLsnclientNetworkBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "lsnclient_network_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *LsnclientNetworkBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data LsnclientNetworkBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting lsnclient_network_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parse the (clientname,network) identity from the ID; the parent is clientname.
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "network"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	clientname_value, ok := idMap["clientname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'clientname' not found in ID")
		return
	}

	network_value, ok := idMap["network"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'network' not found in ID")
		return
	}

	// Mirror the SDK v2 delete args: network (required), plus netmask/td when set.
	// URL-encode values so slashy/special characters are transmitted correctly.
	args := make([]string, 0)
	args = append(args, fmt.Sprintf("network:%s", url.QueryEscape(network_value)))
	if !data.Netmask.IsNull() && data.Netmask.ValueString() != "" {
		args = append(args, fmt.Sprintf("netmask:%s", url.QueryEscape(data.Netmask.ValueString())))
	}
	if !data.Td.IsNull() {
		args = append(args, fmt.Sprintf("td:%s", url.QueryEscape(fmt.Sprintf("%v", data.Td.ValueInt64()))))
	}

	err = r.client.DeleteResourceWithArgs(service.Lsnclient_network_binding.Type(), clientname_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete lsnclient_network_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted lsnclient_network_binding binding")
}

// Helper function to read lsnclient_network_binding data from API
func (r *LsnclientNetworkBindingResource) readLsnclientNetworkBindingFromApi(ctx context.Context, data *LsnclientNetworkBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"clientname", "network"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	clientname_Name, ok := idMap["clientname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'clientname' not found in ID string")
		return
	}

	network_Name, ok := idMap["network"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'network' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Lsnclient_network_binding.Type(),
		ResourceName:             clientname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read lsnclient_network_binding, got error: %s", err))
		return
	}

	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the binding for this network.
	// Match only on `network` (the identity within a clientname), mirroring the
	// SDK v2 resource: the GET response does not echo `td`, so matching on td
	// would reject every record.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["network"].(string); ok && val == network_Name {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		// Binding not present in the returned set: signal removal via a null Id (see above).
		data.Id = types.StringNull()
		return
	}

	lsnclient_network_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
