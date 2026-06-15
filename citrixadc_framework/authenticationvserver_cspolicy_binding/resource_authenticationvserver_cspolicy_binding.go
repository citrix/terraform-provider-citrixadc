package authenticationvserver_cspolicy_binding

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
var _ resource.Resource = &AuthenticationvserverCspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverCspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverCspolicyBindingResource)(nil)

func NewAuthenticationvserverCspolicyBindingResource() resource.Resource {
	return &AuthenticationvserverCspolicyBindingResource{}
}

// AuthenticationvserverCspolicyBindingResource defines the resource implementation.
type AuthenticationvserverCspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverCspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverCspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_cspolicy_binding"
}

func (r *AuthenticationvserverCspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverCspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverCspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_cspolicy_binding resource")
	authenticationvserver_cspolicy_binding := authenticationvserver_cspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource (NITRO add is PUT for this binding)
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_cspolicy_binding.Type(), &authenticationvserver_cspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_cspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_cspolicy_binding resource")

	// Set ID for the resource before reading state.
	// Legacy key order (name,policy) matches resource_id_mapping.json so imported
	// SDK v2 state stays parseable.
	data.Id = types.StringValue(authenticationvserver_cspolicy_bindingComposeId(&data))

	// Read the updated state back
	r.readAuthenticationvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverCspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_cspolicy_binding resource")

	r.readAuthenticationvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverCspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverCspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; every schema attribute is
	// RequiresReplace, so Update is a documented no-op. (Pattern 5)
	tflog.Debug(ctx, "Update is a no-op for authenticationvserver_cspolicy_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readAuthenticationvserverCspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverCspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverCspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_cspolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// ParseIdString handles BOTH the new "name:..,policy:.." format and the legacy
	// SDK v2 positional "name,policy" format.
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

	policy_value, ok := idMap["policy"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Attribute 'policy' not found in ID")
		return
	}

	// policy is the only disambiguating delete arg this binding requires.
	// URL-encode the value to handle slashy/special characters. (binding pattern (b))
	args := []string{fmt.Sprintf("policy:%s", url.QueryEscape(policy_value))}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_cspolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_cspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_cspolicy_binding binding")
}

// Helper function to read authenticationvserver_cspolicy_binding data from API
func (r *AuthenticationvserverCspolicyBindingResource) readAuthenticationvserverCspolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverCspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse name and policy from the ID.
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

	policy_Name, ok := idMap["policy"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'policy' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_cspolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_cspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_cspolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching policy.
	// The NITRO GET response for this binding only echoes name/policy/priority,
	// so policy is the only reliable discriminator (matching SDK v2 behaviour).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_Name {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "authenticationvserver_cspolicy_binding not found with the provided ID attributes")
		return
	}

	authenticationvserver_cspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
