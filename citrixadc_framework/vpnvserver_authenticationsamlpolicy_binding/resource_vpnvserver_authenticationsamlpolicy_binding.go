package vpnvserver_authenticationsamlpolicy_binding

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &VpnvserverAuthenticationsamlpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*VpnvserverAuthenticationsamlpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*VpnvserverAuthenticationsamlpolicyBindingResource)(nil)

func NewVpnvserverAuthenticationsamlpolicyBindingResource() resource.Resource {
	return &VpnvserverAuthenticationsamlpolicyBindingResource{}
}

// VpnvserverAuthenticationsamlpolicyBindingResource defines the resource implementation.
type VpnvserverAuthenticationsamlpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpnvserver_authenticationsamlpolicy_binding"
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data VpnvserverAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating vpnvserver_authenticationsamlpolicy_binding resource")
	vpnvserver_authenticationsamlpolicy_binding := vpnvserver_authenticationsamlpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationsamlpolicy_binding.Type(), &vpnvserver_authenticationsamlpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create vpnvserver_authenticationsamlpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created vpnvserver_authenticationsamlpolicy_binding resource")

	// Set ID for the resource before reading state (legacy order: name,policy)
	data.Id = types.StringValue(vpnvserver_authenticationsamlpolicy_bindingComposeId(data.Name.ValueString(), data.Policy.ValueString()))

	// Read the updated state back
	r.readVpnvserverAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data VpnvserverAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading vpnvserver_authenticationsamlpolicy_binding resource")

	r.readVpnvserverAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state VpnvserverAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating vpnvserver_authenticationsamlpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		vpnvserver_authenticationsamlpolicy_binding := vpnvserver_authenticationsamlpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Vpnvserver_authenticationsamlpolicy_binding.Type(), &vpnvserver_authenticationsamlpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update vpnvserver_authenticationsamlpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated vpnvserver_authenticationsamlpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for vpnvserver_authenticationsamlpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readVpnvserverAuthenticationsamlpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VpnvserverAuthenticationsamlpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VpnvserverAuthenticationsamlpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting vpnvserver_authenticationsamlpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ID is legacy order name,policy; parse to recover the parent name and policy.
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

	// Build delete args matching the NITRO delete endpoint:
	// args=policy:<>,secondary:<>,groupextraction:<>,bindpoint:<>.
	// URL-encode each value (bindpoint and others may contain slashy/special characters).
	args := make([]string, 0)
	if val, ok := idMap["policy"]; ok && val != "" {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(val)))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(strconv.FormatBool(data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(strconv.FormatBool(data.Groupextraction.ValueBool()))))
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Vpnvserver_authenticationsamlpolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete vpnvserver_authenticationsamlpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted vpnvserver_authenticationsamlpolicy_binding binding")
}

// Helper function to read vpnvserver_authenticationsamlpolicy_binding data from API
func (r *VpnvserverAuthenticationsamlpolicyBindingResource) readVpnvserverAuthenticationsamlpolicyBindingFromApi(ctx context.Context, data *VpnvserverAuthenticationsamlpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Vpnvserver_authenticationsamlpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read vpnvserver_authenticationsamlpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "vpnvserver_authenticationsamlpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching policy (the bound entity).
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
		diags.AddError("Client Error", fmt.Sprintf("vpnvserver_authenticationsamlpolicy_binding not found with the provided ID attributes"))
		return
	}

	vpnvserver_authenticationsamlpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
