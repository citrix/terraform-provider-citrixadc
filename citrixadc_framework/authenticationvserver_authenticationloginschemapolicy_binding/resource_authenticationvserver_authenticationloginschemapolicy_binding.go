package authenticationvserver_authenticationloginschemapolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationloginschemapolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationloginschemapolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationloginschemapolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationloginschemapolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationloginschemapolicyBindingResource{}
}

// AuthenticationvserverAuthenticationloginschemapolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationloginschemapolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationloginschemapolicy_binding"
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationloginschemapolicy_binding resource")
	authenticationvserver_authenticationloginschemapolicy_binding := authenticationvserver_authenticationloginschemapolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), &authenticationvserver_authenticationloginschemapolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationloginschemapolicy_binding resource")

	// Set ID for the resource before reading state.
	// Backward-compatible composite ID matching SDK v2 ("name,policy") expressed in the
	// new "name:value,policy:value" form (resource_id_mapping.json: "name,policy").
	data.Id = types.StringValue(authenticationvserver_authenticationloginschemapolicy_bindingComposeId(&data))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationloginschemapolicy_binding resource")

	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationloginschemapolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationvserver_authenticationloginschemapolicy_binding := authenticationvserver_authenticationloginschemapolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), &authenticationvserver_authenticationloginschemapolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationvserver_authenticationloginschemapolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationvserver_authenticationloginschemapolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationloginschemapolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// Parse name from the (legacy- or new-format) ID; the disambiguating delete args
	// (policy/secondary/groupextraction/bindpoint) come from the model state, matching
	// the SDK v2 resource. Slashy/special values are URL-encoded.
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
	policy_value, ok := idMap["policy"]
	if !ok || policy_value == "" {
		policy_value = data.Policy.ValueString()
	}
	args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(policy_value)))
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%v", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))))
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationloginschemapolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationloginschemapolicy_binding data from API
func (r *AuthenticationvserverAuthenticationloginschemapolicyBindingResource) readAuthenticationvserverAuthenticationloginschemapolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationloginschemapolicyBindingResourceModel, diags *diag.Diagnostics) {

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

	policy_value, ok := idMap["policy"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'policy' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_authenticationloginschemapolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationloginschemapolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationloginschemapolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one whose policy matches the ID (matches the
	// SDK v2 read filter, which keyed only on policy under the parent name).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_value {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "authenticationvserver_authenticationloginschemapolicy_binding not found with the provided ID attributes")
		return
	}

	authenticationvserver_authenticationloginschemapolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
