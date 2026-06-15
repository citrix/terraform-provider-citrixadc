package authenticationpolicylabel_authenticationpolicy_binding

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
var _ resource.Resource = &AuthenticationpolicylabelAuthenticationpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationpolicylabelAuthenticationpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationpolicylabelAuthenticationpolicyBindingResource)(nil)

func NewAuthenticationpolicylabelAuthenticationpolicyBindingResource() resource.Resource {
	return &AuthenticationpolicylabelAuthenticationpolicyBindingResource{}
}

// AuthenticationpolicylabelAuthenticationpolicyBindingResource defines the resource implementation.
type AuthenticationpolicylabelAuthenticationpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationpolicylabel_authenticationpolicy_binding"
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationpolicylabel_authenticationpolicy_binding resource")
	authenticationpolicylabel_authenticationpolicy_binding := authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), &authenticationpolicylabel_authenticationpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationpolicylabel_authenticationpolicy_binding resource")

	// Set ID for the resource before reading state
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("labelname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Labelname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Read the updated state back
	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationpolicylabel_authenticationpolicy_binding resource")

	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationpolicylabel_authenticationpolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationpolicylabel_authenticationpolicy_binding := authenticationpolicylabel_authenticationpolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), &authenticationpolicylabel_authenticationpolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationpolicylabel_authenticationpolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationpolicylabel_authenticationpolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationpolicylabel_authenticationpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
	if err != nil {
		resp.Diagnostics.AddError("Parse Error", fmt.Sprintf("Unable to parse ID for delete: %s", err))
		return
	}

	labelname_value, ok := idMap["labelname"]
	if !ok {
		resp.Diagnostics.AddError("Parse Error", "Parent attribute 'labelname' not found in ID")
		return
	}

	var argsMap map[string]string = make(map[string]string)
	if val, ok := idMap["policyname"]; ok && val != "" {
		argsMap["policyname"] = val
	}
	// Match SDK v2 behavior: include priority as a delete arg to disambiguate
	// bindings of the same policy at different priorities (values are URL-encoded
	// internally by DeleteResourceWithArgsMap).
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		argsMap["priority"] = fmt.Sprintf("%v", data.Priority.ValueInt64())
	}

	err = r.client.DeleteResourceWithArgsMap(service.Authenticationpolicylabel_authenticationpolicy_binding.Type(), labelname_value, argsMap)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationpolicylabel_authenticationpolicy_binding binding")
}

// Helper function to read authenticationpolicylabel_authenticationpolicy_binding data from API
func (r *AuthenticationpolicylabelAuthenticationpolicyBindingResource) readAuthenticationpolicylabelAuthenticationpolicyBindingFromApi(ctx context.Context, data *AuthenticationpolicylabelAuthenticationpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Case 4: Array filter with parent ID - parse from ID
	idMap, _, err := utils.ParseIdString(data.Id.ValueString(), []string{"labelname", "policyname"}, nil)
	if err != nil {
		diags.AddError("Parse Error", fmt.Sprintf("Unable to parse ID: %s", err))
		return
	}

	labelname_Name, ok := idMap["labelname"]
	if !ok {
		diags.AddError("Parse Error", "ID attribute 'labelname' not found in ID string")
		return
	}

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Authenticationpolicylabel_authenticationpolicy_binding.Type(),
		ResourceName:             labelname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationpolicylabel_authenticationpolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationpolicylabel_authenticationpolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policyname
		if idVal, ok := idMap["policyname"]; ok {
			if val, ok := v["policyname"].(string); ok {
				if val != idVal {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		} else if _, ok := v["policyname"].(string); ok {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	//  Resource is missing
	if foundIndex == -1 {
		diags.AddError("Client Error", fmt.Sprintf("authenticationpolicylabel_authenticationpolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationpolicylabel_authenticationpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
