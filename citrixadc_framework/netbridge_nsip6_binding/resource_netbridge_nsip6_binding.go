package netbridge_nsip6_binding

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
var _ resource.Resource = &NetbridgeNsip6BindingResource{}
var _ resource.ResourceWithConfigure = (*NetbridgeNsip6BindingResource)(nil)
var _ resource.ResourceWithImportState = (*NetbridgeNsip6BindingResource)(nil)

func NewNetbridgeNsip6BindingResource() resource.Resource {
	return &NetbridgeNsip6BindingResource{}
}

// NetbridgeNsip6BindingResource defines the resource implementation.
type NetbridgeNsip6BindingResource struct {
	client *service.NitroClient
}

func (r *NetbridgeNsip6BindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *NetbridgeNsip6BindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_netbridge_nsip6_binding"
}

func (r *NetbridgeNsip6BindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *NetbridgeNsip6BindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating netbridge_nsip6_binding resource")
	netbridge_nsip6_binding := netbridge_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Netbridge_nsip6_binding.Type(), &netbridge_nsip6_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create netbridge_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created netbridge_nsip6_binding resource")

	// Set ID for the resource before reading state.
	// Composite ID order matches resource_id_mapping.json ("name,ipaddress").
	// netmask is intentionally excluded (it is not part of the SDK v2 ID and is a
	// Computed/server-overridden value for IPv6 bindings).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ipaddress.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "netbridge_nsip6_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsip6BindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading netbridge_nsip6_binding resource")

	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *NetbridgeNsip6BindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state NetbridgeNsip6BindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating netbridge_nsip6_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		netbridge_nsip6_binding := netbridge_nsip6_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Netbridge_nsip6_binding.Type(), &netbridge_nsip6_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update netbridge_nsip6_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated netbridge_nsip6_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for netbridge_nsip6_binding resource, skipping update")
	}

	// Read the updated state back
	r.readNetbridgeNsip6BindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "netbridge_nsip6_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *NetbridgeNsip6BindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data NetbridgeNsip6BindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting netbridge_nsip6_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ipaddress"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	// The ipaddress value (an IPv6 subnet, e.g. "dea:97c5:d381:e72b::/64") contains
	// characters such as '/' that must be URL-encoded before being placed in the
	// delete ?args= query string. The NITRO client does NOT encode arg values, so we
	// encode here (mirrors the SDK v2 resource's url.PathEscape). netmask is excluded
	// from the delete args, matching the SDK v2 resource.
	args := make([]string, 0)
	if val, ok := idMap["ipaddress"]; ok && val != "" {
		args = append(args, fmt.Sprintf("ipaddress:%s", url.QueryEscape(val)))
	}

	err = r.client.DeleteResourceWithArgs(service.Netbridge_nsip6_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete netbridge_nsip6_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted netbridge_nsip6_binding binding")
}

// Helper function to read netbridge_nsip6_binding data from API
func (r *NetbridgeNsip6BindingResource) readNetbridgeNsip6BindingFromApi(ctx context.Context, data *NetbridgeNsip6BindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "ipaddress"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	name_Name, ok := idMap["name"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'name' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Netbridge_nsip6_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read netbridge_nsip6_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one matching the binding's ipaddress.
	// Identity is (name, ipaddress) per resource_id_mapping.json; netmask is not a
	// match key (it is server-derived for IPv6 bindings).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ipaddress
		if idVal, ok := idMap["ipaddress"]; ok {
			if val, ok := v["ipaddress"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	netbridge_nsip6_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
