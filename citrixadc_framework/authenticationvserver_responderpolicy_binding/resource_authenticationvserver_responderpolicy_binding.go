package authenticationvserver_responderpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverResponderpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverResponderpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverResponderpolicyBindingResource)(nil)

func NewAuthenticationvserverResponderpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverResponderpolicyBindingResource{}
}

// AuthenticationvserverResponderpolicyBindingResource defines the resource implementation.
type AuthenticationvserverResponderpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverResponderpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_responderpolicy_binding"
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverResponderpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_responderpolicy_binding resource")
	authenticationvserver_responderpolicy_binding := authenticationvserver_responderpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource (NITRO 'add' verb is PUT for this binding)
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_responderpolicy_binding.Type(), &authenticationvserver_responderpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_responderpolicy_binding resource")

	// Set ID for the resource before reading state.
	// Legacy SDK v2 ID format: "name,policy" (see resource_id_mapping.json).
	data.Id = types.StringValue(data.Name.ValueString() + "," + data.Policy.ValueString())

	// Read the updated state back
	r.readAuthenticationvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_responderpolicy_binding resource")

	r.readAuthenticationvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// No NITRO update endpoint exists for this binding; every attribute is RequiresReplace,
	// so Terraform never reaches Update with a real change. Read back and persist (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for authenticationvserver_responderpolicy_binding; all attributes are RequiresReplace")

	r.readAuthenticationvserverResponderpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverResponderpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverResponderpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_responderpolicy_binding resource")
	// Parse the ID (accepts both new key:value and legacy "name,policy" formats).
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

	// Build the delete args. policy is the disambiguator; the optional discriminators are
	// added only when set in state (mirrors the SDK v2 delete). URL-encode slashy/special values.
	args := make([]string, 0)
	if policy, ok := idMap["policy"]; ok && policy != "" {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(policy)))
	}
	if !data.Secondary.IsNull() && data.Secondary.ValueBool() {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(fmt.Sprintf("%v", data.Secondary.ValueBool()))))
	}
	if !data.Groupextraction.IsNull() && data.Groupextraction.ValueBool() {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))))
	}
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_responderpolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_responderpolicy_binding binding")
}

// Helper function to read authenticationvserver_responderpolicy_binding data from API
func (r *AuthenticationvserverResponderpolicyBindingResource) readAuthenticationvserverResponderpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverResponderpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Parse the ID (accepts both new key:value and legacy "name,policy" formats).
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
		ResourceType:             service.Authenticationvserver_responderpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_responderpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_responderpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the matching policy (matches SDK v2 behaviour).
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_value {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", "authenticationvserver_responderpolicy_binding not found with the provided ID attributes")
		return
	}

	authenticationvserver_responderpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
