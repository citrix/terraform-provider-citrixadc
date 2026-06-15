package authenticationvserver_authenticationcertpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationcertpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationcertpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationcertpolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationcertpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationcertpolicyBindingResource{}
}

// AuthenticationvserverAuthenticationcertpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationcertpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationcertpolicy_binding"
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationcertpolicy_binding resource")
	authenticationvserver_authenticationcertpolicy_binding := authenticationvserver_authenticationcertpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationcertpolicy_binding.Type(), &authenticationvserver_authenticationcertpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationcertpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationcertpolicy_binding resource")

	// Set ID for the resource before reading state.
	// ID identity matches the SDK v2 contract and resource_id_mapping.json: name,policy.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationcertpolicy_binding resource")

	r.readAuthenticationvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationcertpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationvserver_authenticationcertpolicy_binding := authenticationvserver_authenticationcertpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationcertpolicy_binding.Type(), &authenticationvserver_authenticationcertpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationcertpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationvserver_authenticationcertpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationvserver_authenticationcertpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationcertpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationcertpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationcertpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parse the parent name (and policy) from the ID (handles both the new
	// name:...,policy:... format and the legacy SDK v2 name,policy format).
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

	// Build disambiguating delete args from state, mirroring the SDK v2 resource.
	// Values are URL-encoded (pattern (b)) so slashy/special bindpoint values are
	// transmitted correctly.
	args := make([]string, 0)
	policy_value := idMap["policy"]
	if policy_value == "" && !data.Policy.IsNull() {
		policy_value = data.Policy.ValueString()
	}
	args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(policy_value)))
	if !data.Secondary.IsNull() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%v", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))))
	}
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_authenticationcertpolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationcertpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationcertpolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationcertpolicy_binding data from API
func (r *AuthenticationvserverAuthenticationcertpolicyBindingResource) readAuthenticationvserverAuthenticationcertpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationcertpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Authenticationvserver_authenticationcertpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationcertpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationcertpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the binding with the matching policy.
	// The binding's identity within a parent vserver is (name, policy); other
	// attributes (groupextraction, secondary, bindpoint) are not part of the ID.
	policy_value, ok := idMap["policy"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'policy' not found in ID string")
		return
	}
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_value {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationcertpolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationvserver_authenticationcertpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
