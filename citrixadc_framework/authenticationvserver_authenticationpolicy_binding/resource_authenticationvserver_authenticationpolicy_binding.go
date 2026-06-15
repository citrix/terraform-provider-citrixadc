package authenticationvserver_authenticationpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationpolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationpolicyBindingResource{}
}

// AuthenticationvserverAuthenticationpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationpolicy_binding"
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationpolicy_binding resource")
	authenticationvserver_authenticationpolicy_binding := authenticationvserver_authenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), &authenticationvserver_authenticationpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationpolicy_binding resource")

	// Set ID for the resource before reading state.
	// Composite key order matches the legacy SDK v2 ID (name,policy) recorded in
	// resource_id_mapping.json so imported SDK v2 state IDs resolve via ParseIdString.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationpolicy_binding resource")

	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationvserver_authenticationpolicy_binding := authenticationvserver_authenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationpolicy_binding.Type(), &authenticationvserver_authenticationpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationvserver_authenticationpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationvserver_authenticationpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationpolicy_binding resource")
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

	// Build delete args matching the SDK v2 contract: policy is required; secondary,
	// groupextraction and bindpoint are passed when known (from state) to disambiguate
	// the binding. DeleteResourceWithArgsMap URL-encodes the values.
	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policy"]; ok && val != "" {
		argsMap["policy"] = val
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		argsMap["secondary"] = fmt.Sprintf("%v", data.Secondary.ValueBool())
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		argsMap["groupextraction"] = fmt.Sprintf("%v", data.Groupextraction.ValueBool())
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = data.Bindpoint.ValueString()
	}

	err = r.client.DeleteResourceWithArgsMap(service.Authenticationvserver_authenticationpolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationpolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationpolicy_binding data from API
func (r *AuthenticationvserverAuthenticationpolicyBindingResource) readAuthenticationvserverAuthenticationpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Authenticationvserver_authenticationpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// The composite ID is (name,policy); the binding is uniquely identified within the
	// parent vserver by its policy name (mirrors the SDK v2 read).
	foundIndex := -1
	policyVal := idMap["policy"]
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policyVal {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationpolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationvserver_authenticationpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
