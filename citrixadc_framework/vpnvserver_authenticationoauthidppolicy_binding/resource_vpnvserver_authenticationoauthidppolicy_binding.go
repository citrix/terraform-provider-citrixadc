package vpnvserver_authenticationoauthidppolicy_binding

import (
	"context"
	"fmt"
	"net/url"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnvserverAuthenticationoauthidppolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationoauthidppolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationoauthidppolicyBindingResource)(nil)

func NewVpnvserverAuthenticationoauthidppolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationoauthidppolicyBindingResource{}
}

// VpnvserverAuthenticationoauthidppolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationoauthidppolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationoauthidppolicy_binding"
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationoauthidppolicy_binding resource")
	vpnvserver_authenticationoauthidppolicy_binding := vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), &vpnvserver_authenticationoauthidppolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_authenticationoauthidppolicy_binding resource")

	// Set ID for the resource before reading state.
	// Composite of legacy unique attrs (name,policy) matching resource_id_mapping.json.
	data.Id = types.StringValue(vpnvserver_authenticationoauthidppolicy_bindingComposeId(&data))

	// Read the updated state back
	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationoauthidppolicy_binding resource")

	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver_authenticationoauthidppolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnvserver_authenticationoauthidppolicy_binding := vpnvserver_authenticationoauthidppolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), &vpnvserver_authenticationoauthidppolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnvserver_authenticationoauthidppolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver_authenticationoauthidppolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationoauthidppolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationoauthidppolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parse name from the ID (handles both new key:value and legacy comma formats).
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

	// Delete args come from the model (bindpoint/secondary/groupextraction are not
	// echoed by GET so they are only available from state, not from the ID).
	// URL-encode values so slashy/special characters survive in the query string.
	args := make([]string, 0)
	if val, ok := idMap["policy"]; ok && val != "" {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(val)))
	} else if !data.Policy.IsNull() {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(data.Policy.ValueString())))
	}
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}
	if !data.Secondary.IsNull() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%t", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%t", data.Groupextraction.ValueBool()))))
	}

	err = r.client.DeleteResourceWithArgs(service.Vpnvserver_authenticationoauthidppolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_authenticationoauthidppolicy_binding binding")
}

// Helper function to read vpnvserver_authenticationoauthidppolicy_binding data from API
func (r *VpnvserverAuthenticationoauthidppolicyBindingResource) readVpnvserverAuthenticationoauthidppolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationoauthidppolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Vpnvserver_authenticationoauthidppolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationoauthidppolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnvserver_authenticationoauthidppolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one matching the policy (the binding's
	// disambiguator within the parent vpnvserver name).
	policyId, hasPolicy := idMap["policy"]
	foundIndex := -1
	for i, v := range dataArr {
		if hasPolicy {
			if val, ok := v["policy"].(string); ok && val == policyId {
				foundIndex = i
				break
			}
		} else {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("vpnvserver_authenticationoauthidppolicy_binding not found with the provided ID attributes"))
		return
	}

	vpnvserver_authenticationoauthidppolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
