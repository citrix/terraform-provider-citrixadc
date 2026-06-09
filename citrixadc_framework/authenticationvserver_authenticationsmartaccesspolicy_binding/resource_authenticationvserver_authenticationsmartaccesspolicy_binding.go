package authenticationvserver_authenticationsmartaccesspolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationsmartaccesspolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource{}
}

// AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationsmartaccesspolicy_binding"
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationsmartaccesspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationsmartaccesspolicy_binding resource")
	authenticationvserver_authenticationsmartaccesspolicy_binding := authenticationvserver_authenticationsmartaccesspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(), &authenticationvserver_authenticationsmartaccesspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationsmartaccesspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationsmartaccesspolicy_binding resource")

	// Set ID for the resource before reading state. ID = name,policy
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationsmartaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationsmartaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationsmartaccesspolicy_binding resource")

	r.readAuthenticationvserverAuthenticationsmartaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationsmartaccesspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// Update is a no-op for this binding: NITRO exposes only add (bind) / delete (unbind),
	// there is no update endpoint, and every attribute uses RequiresReplace. (Pattern 5)
	tflog.Debug(ctx, "Update is a no-op for authenticationvserver_authenticationsmartaccesspolicy_binding; all attributes are RequiresReplace")

	// Read the current state back
	r.readAuthenticationvserverAuthenticationsmartaccesspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationsmartaccesspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationsmartaccesspolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgsMap.
	// URL/name key = name; delete args = policy (binding anchor) plus secondary/groupextraction
	// when set. (NITRO delete args also include bindpoint, which is absent from the model - see report.)
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
	if !data.Policy.IsNull() && data.Policy.ValueString() != "" {
		argsMap["policy"] = utils.UrlEncode(data.Policy.ValueString())
	}
	if !data.Secondary.IsNull() {
		argsMap["secondary"] = utils.UrlEncode(fmt.Sprintf("%v", data.Secondary.ValueBool()))
	}
	if !data.Groupextraction.IsNull() {
		argsMap["groupextraction"] = utils.UrlEncode(fmt.Sprintf("%v", data.Groupextraction.ValueBool()))
	}

	err = r.client.DeleteResourceWithArgsMap(service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(), name_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationsmartaccesspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationsmartaccesspolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationsmartaccesspolicy_binding data from API
func (r *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResource) readAuthenticationvserverAuthenticationsmartaccesspolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationsmartaccesspolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), nil, nil)
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
		ResourceType:             service.Authenticationvserver_authenticationsmartaccesspolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationsmartaccesspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationsmartaccesspolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// Filter on policy (the binding anchor); name is the parent URL key.
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

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationsmartaccesspolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationvserver_authenticationsmartaccesspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
