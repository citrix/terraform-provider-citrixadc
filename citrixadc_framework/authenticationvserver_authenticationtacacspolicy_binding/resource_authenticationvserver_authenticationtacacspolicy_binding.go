package authenticationvserver_authenticationtacacspolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuthenticationtacacspolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuthenticationtacacspolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuthenticationtacacspolicyBindingResource)(nil)

func NewAuthenticationvserverAuthenticationtacacspolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuthenticationtacacspolicyBindingResource{}
}

// AuthenticationvserverAuthenticationtacacspolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuthenticationtacacspolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_authenticationtacacspolicy_binding"
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_authenticationtacacspolicy_binding resource")
	authenticationvserver_authenticationtacacspolicy_binding := authenticationvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - use UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationtacacspolicy_binding.Type(), &authenticationvserver_authenticationtacacspolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_authenticationtacacspolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(authenticationvserver_authenticationtacacspolicy_bindingComposeId(&data))

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_authenticationtacacspolicy_binding resource")

	r.readAuthenticationvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	tflog.Debug(ctx, "Updating authenticationvserver_authenticationtacacspolicy_binding resource")

	// Check if there are any changes in updateable attributes
	hasChange := false

	if hasChange {
		// Create API request body from the model
		authenticationvserver_authenticationtacacspolicy_binding := authenticationvserver_authenticationtacacspolicy_bindingGetThePayloadFromthePlan(ctx, &data)
		// Make API call
		// Binding resource - use UpdateUnnamedResource
		err := r.client.UpdateUnnamedResource(service.Authenticationvserver_authenticationtacacspolicy_binding.Type(), &authenticationvserver_authenticationtacacspolicy_binding)
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to update authenticationvserver_authenticationtacacspolicy_binding, got error: %s", err))
			return
		}

		tflog.Trace(ctx, "Updated authenticationvserver_authenticationtacacspolicy_binding resource")
	} else {
		tflog.Debug(ctx, "No changes detected for authenticationvserver_authenticationtacacspolicy_binding resource, skipping update")
	}

	// Read the updated state back
	r.readAuthenticationvserverAuthenticationtacacspolicyBindingFromApi(ctx, &data, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_authenticationtacacspolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// The parent (name) is the resource name; the bound policy and the remaining unique/disambiguating
	// attributes are passed as delete args. ParseIdString handles both the new key:value ID format and
	// the legacy SDK v2 "name,policy" comma format (returns URL-decoded values), so we URL-encode each
	// arg value ourselves (DeleteResourceWithArgs/Map do NOT encode arg values) to handle slashy/special
	// values, mirroring the SDK v2 resource.
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
	if val, ok := idMap["policy"]; ok && val != "" {
		args = append(args, fmt.Sprintf("policy:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["secondary"]; ok && val != "" {
		args = append(args, fmt.Sprintf("secondary:%s", url.QueryEscape(val)))
	}
	if val, ok := idMap["groupextraction"]; ok && val != "" {
		args = append(args, fmt.Sprintf("groupextraction:%s", url.QueryEscape(val)))
	}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_authenticationtacacspolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_authenticationtacacspolicy_binding binding")
}

// Helper function to read authenticationvserver_authenticationtacacspolicy_binding data from API
func (r *AuthenticationvserverAuthenticationtacacspolicyBindingResource) readAuthenticationvserverAuthenticationtacacspolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuthenticationtacacspolicyBindingResourceModel, diags *diag.Diagnostics) {

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
		ResourceType:             service.Authenticationvserver_authenticationtacacspolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_authenticationtacacspolicy_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		diags.AddError("Client Error", "authenticationvserver_authenticationtacacspolicy_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id.
	// policy is the primary (string) disambiguator under a given parent name. The boolean
	// ID components (groupextraction, secondary) may be ABSENT from the GET response when
	// they are at their NITRO default (false), so an absent boolean field is treated as
	// false rather than as a non-match.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check policy (strict)
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

		// Check groupextraction (absent => default false)
		if idVal, ok := idMap["groupextraction"]; ok {
			idValBool, _ := strconv.ParseBool(idVal)
			respBool := false
			if val, ok := v["groupextraction"].(bool); ok {
				respBool = val
			}
			if respBool != idValBool {
				match = false
				continue
			}
		}

		// Check secondary (absent => default false)
		if idVal, ok := idMap["secondary"]; ok {
			idValBool, _ := strconv.ParseBool(idVal)
			respBool := false
			if val, ok := v["secondary"].(bool); ok {
				respBool = val
			}
			if respBool != idValBool {
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
		diags.AddError("Client Error", fmt.Sprintf("authenticationvserver_authenticationtacacspolicy_binding not found with the provided ID attributes"))
		return
	}

	authenticationvserver_authenticationtacacspolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
