package authenticationvserver_authenticationnegotiatepolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationnegotiatepolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationnegotiatepolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationnegotiatepolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationnegotiatepolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationnegotiatepolicyBindingResource{}
}

// AuthenticationvserverAuthenticationnegotiatepolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationnegotiatepolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationnegotiatepolicy_binding"
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationnegotiatepolicy_binding resource")
	authenticationvserver_authenticationnegotiatepolicy_binding := authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), &authenticationvserver_authenticationnegotiatepolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// Set ID for the resource before reading state.
	// Composite key:UrlEncode(value) pairs in the legacy SDK v2 order (name,policy)
	// per resource_id_mapping.json. name+policy uniquely identify the binding.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationnegotiatepolicy_binding resource")

	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationnegotiatepolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationvserver_authenticationnegotiatepolicy_binding := authenticationvserver_authenticationnegotiatepolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), &authenticationvserver_authenticationnegotiatepolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationvserver_authenticationnegotiatepolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationvserver_authenticationnegotiatepolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationnegotiatepolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// Parent (name) and policy come from the ID; the remaining delete args
	// (secondary, groupextraction, bindpoint) come from state, mirroring the
	// SDK v2 delete-arg set. DeleteResourceWithArgsMap URL-encodes values internally.
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
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		argsMap["secondary"] = fmt.Sprintf("%v", data.Secondary.ValueBool())
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		argsMap["groupextraction"] = fmt.Sprintf("%v", data.Groupextraction.ValueBool())
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		argsMap["bindpoint"] = data.Bindpoint.ValueString()
	}

	err = r.client.DeleteResourceWithArgsMap(service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationnegotiatepolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationnegotiatepolicy_binding data from API
func (r *AuthenticationvserverAuthenticationnegotiatepolicyBindingResource) readAuthenticationvserverAuthenticationnegotiatepolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationnegotiatepolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Authenticationvserver_authenticationnegotiatepolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationnegotiatepolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationnegotiatepolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the binding matching the policy from the ID.
	// name+policy uniquely identify the binding (same as SDK v2 Read).
	policy_value := idMap["policy"]
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_value {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationnegotiatepolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationvserver_authenticationnegotiatepolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
