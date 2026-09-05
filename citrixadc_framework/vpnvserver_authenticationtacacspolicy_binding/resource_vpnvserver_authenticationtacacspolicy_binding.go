package vpnvserver_authenticationtacacspolicy_binding

import (
	"context"
	"fmt"
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
var _ resource.Resource = &VpnvserverAuthenticationtacacspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationtacacspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationtacacspolicyBindingResource)(nil)

func NewVpnvserverAuthenticationtacacspolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationtacacspolicyBindingResource{}
}

// VpnvserverAuthenticationtacacspolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationtacacspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationtacacspolicy_binding"
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationtacacspolicy_binding resource")
	vpnvserver_authenticationtacacspolicy_binding := vpnvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), &vpnvserver_authenticationtacacspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_authenticationtacacspolicy_binding resource")

	// Set ID for the resource before reading state.
	// ID is composed of name,policy per resource_id_mapping.json (matches SDK v2
	// d.SetId(name + "," + policy)). bindpoint is NOT part of the identity.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_authenticationtacacspolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationtacacspolicy_binding resource")

	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver_authenticationtacacspolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnvserver_authenticationtacacspolicy_binding := vpnvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), &vpnvserver_authenticationtacacspolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnvserver_authenticationtacacspolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver_authenticationtacacspolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_authenticationtacacspolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationtacacspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationtacacspolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policy"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	name_value, ok := idMap["name"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'name' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policy"]; ok && val != "" {
		argsMap["policy"] = val
	}
	// bindpoint is not part of the ID; source it from state to disambiguate the
	// binding on delete (matches SDK v2 delete which passed bindpoint when set).
	// DeleteResourceWithArgsMap URL-encodes the arg values internally.
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = data.Bindpoint.ValueString()
	}
	if !data.Secondary.IsNull() && data.Secondary.ValueBool() {
		argsMap["secondary"] = "true"
	}
	if !data.Groupextraction.IsNull() && data.Groupextraction.ValueBool() {
		argsMap["groupextraction"] = "true"
	}

	err = r.client.DeleteResourceWithArgsMap(service.Vpnvserver_authenticationtacacspolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_authenticationtacacspolicy_binding binding")
}

// Helper function to read vpnvserver_authenticationtacacspolicy_binding data from API
func (r *VpnvserverAuthenticationtacacspolicyBindingResource) readVpnvserverAuthenticationtacacspolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationtacacspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"name", "policy"}, nil)
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
		ResourceType:             service.Vpnvserver_authenticationtacacspolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
		// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right id.
	// The binding identity is (name, policy); bindpoint is not part of the ID
	// and is not returned by GET, so match on policy only (matches SDK v2 read).
	foundIndex := -1
	for i, v := range dataArr {
		if idVal, ok := idMap["policy"]; ok {
			if val, ok := v["policy"].(string); ok && val == idVal {
				foundIndex = i
				break
			}
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	vpnvserver_authenticationtacacspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
