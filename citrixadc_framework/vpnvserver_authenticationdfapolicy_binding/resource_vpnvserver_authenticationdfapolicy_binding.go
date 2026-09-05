package vpnvserver_authenticationdfapolicy_binding

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
var _ resource.Resource = &VpnvserverAuthenticationdfapolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationdfapolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationdfapolicyBindingResource)(nil)

func NewVpnvserverAuthenticationdfapolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationdfapolicyBindingResource{}
}

// VpnvserverAuthenticationdfapolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationdfapolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationdfapolicy_binding"
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationdfapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationdfapolicy_binding resource")
	vpnvserver_authenticationdfapolicy_binding := vpnvserver_authenticationdfapolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationdfapolicy_binding.Type(), &vpnvserver_authenticationdfapolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationdfapolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_authenticationdfapolicy_binding resource")

	// Set ID for the resource before reading state (legacy format: name,policy)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readVpnvserverAuthenticationdfapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_authenticationdfapolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationdfapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationdfapolicy_binding resource")

	r.readVpnvserverAuthenticationdfapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
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

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverAuthenticationdfapolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver_authenticationdfapolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnvserver_authenticationdfapolicy_binding := vpnvserver_authenticationdfapolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationdfapolicy_binding.Type(), &vpnvserver_authenticationdfapolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationdfapolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnvserver_authenticationdfapolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver_authenticationdfapolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnvserverAuthenticationdfapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "vpnvserver_authenticationdfapolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationdfapolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationdfapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationdfapolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parent (name) and policy come from the ID; the remaining disambiguating args
	// (secondary, groupextraction, bindpoint) come from prior state, mirroring the
	// SDK v2 delete contract. Values are URL-encoded for slashy/special content.
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

	args := make([]string, 0)
	if val, ok := idMap["policy"]; ok && val != "" {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(val)))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%t", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%t", data.Groupextraction.ValueBool()))))
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Vpnvserver_authenticationdfapolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_authenticationdfapolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_authenticationdfapolicy_binding binding")
}

// Helper function to read vpnvserver_authenticationdfapolicy_binding data from API
func (r *VpnvserverAuthenticationdfapolicyBindingResource) readVpnvserverAuthenticationdfapolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationdfapolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Vpnvserver_authenticationdfapolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationdfapolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the one with the right policy (the binding key,
	// matching the SDK v2 contract which filtered solely on policy under parent name).
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policy
		if idVal, ok := idMap["policy"]; ok {
			if val, ok := v["policy"].(string); ok {
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

	vpnvserver_authenticationdfapolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
