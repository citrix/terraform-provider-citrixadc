package authenticationvserver_auditsyslogpolicy_binding

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
var _ resource.Resource = &AuthenticationvserverAuditsyslogpolicyBindingResource{}
var _ resource.ResourceWithConfigure = (*AuthenticationvserverAuditsyslogpolicyBindingResource)(nil)
var _ resource.ResourceWithImportState = (*AuthenticationvserverAuditsyslogpolicyBindingResource)(nil)

func NewAuthenticationvserverAuditsyslogpolicyBindingResource() resource.Resource {
	return &AuthenticationvserverAuditsyslogpolicyBindingResource{}
}

// AuthenticationvserverAuditsyslogpolicyBindingResource defines the resource implementation.
type AuthenticationvserverAuditsyslogpolicyBindingResource struct {
	client *service.NitroClient
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_authenticationvserver_auditsyslogpolicy_binding"
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// composeId builds the backward-compatible composite ID. The positional order
// (name,policy) matches resource_id_mapping.json so legacy SDK v2 imports still resolve.
func composeId(data *AuthenticationvserverAuditsyslogpolicyBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	return strings.Join(idParts, ",")
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Creating authenticationvserver_auditsyslogpolicy_binding resource")
	authenticationvserver_auditsyslogpolicy_binding := authenticationvserver_auditsyslogpolicy_bindingGetThePayloadFromthePlan(ctx, &data)

	// Make API call
	// Binding resource - NITRO add verb is PUT for this binding -> UpdateUnnamedResource
	err := r.client.UpdateUnnamedResource(service.Authenticationvserver_auditsyslogpolicy_binding.Type(), &authenticationvserver_auditsyslogpolicy_binding)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Created authenticationvserver_auditsyslogpolicy_binding resource")

	// Set ID for the resource before reading state
	data.Id = types.StringValue(composeId(&data))

	// Read the updated state back
	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "authenticationvserver_auditsyslogpolicy_binding not found on the ADC immediately after create")
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Reading authenticationvserver_auditsyslogpolicy_binding resource")

	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// Binding is gone on the ADC (readFromApi nulled the Id): drop it from state so a
	// subsequent apply recreates it, matching the SDK v2 provider's behaviour.
	if data.Id.IsNull() {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data, state AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state to preserve ID
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Preserve ID from prior state
	data.Id = state.Id

	// All attributes are RequiresReplace; there is no NITRO update endpoint for this
	// binding, so Update is a documented no-op (Pattern 5).
	tflog.Debug(ctx, "Update is a no-op for authenticationvserver_auditsyslogpolicy_binding; all attributes are RequiresReplace")

	// Read the updated state back
	r.readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx, &data, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Id.IsNull() {
		resp.Diagnostics.AddError("Client Error", "authenticationvserver_auditsyslogpolicy_binding not found on the ADC immediately after update")
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data AuthenticationvserverAuditsyslogpolicyBindingResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Deleting authenticationvserver_auditsyslogpolicy_binding resource")
	// Binding with parent - delete using DeleteResourceWithArgs.
	// legacyOrder name,policy matches resource_id_mapping.json so legacy IDs parse too.
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
	// secondary/groupextraction/bindpoint are part of the binding identity on the ADC.
	// Include them in the delete args when set (URL-encoded for slashy/special values).
	if !data.Bindpoint.IsNull() && data.Bindpoint.ValueString() != "" {
		args = append(args, fmt.Sprintf("bindpoint:%s", url.QueryEscape(data.Bindpoint.ValueString())))
	}
	if !data.Secondary.IsNull() && data.Secondary.ValueBool() {
		args = append(args, fmt.Sprintf("secondary:%v", data.Secondary.ValueBool()))
	}
	if !data.Groupextraction.IsNull() && data.Groupextraction.ValueBool() {
		args = append(args, fmt.Sprintf("groupextraction:%v", data.Groupextraction.ValueBool()))
	}

	err = r.client.DeleteResourceWithArgs(service.Authenticationvserver_auditsyslogpolicy_binding.Type(), name_value, args)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to delete authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Deleted authenticationvserver_auditsyslogpolicy_binding binding")
}

// Helper function to read authenticationvserver_auditsyslogpolicy_binding data from API
func (r *AuthenticationvserverAuditsyslogpolicyBindingResource) readAuthenticationvserverAuditsyslogpolicyBindingFromApi(ctx context.Context, data *AuthenticationvserverAuditsyslogpolicyBindingResourceModel, diags *diag.Diagnostics) {

	// Array filter with parent ID - parse from ID. legacyOrder name,policy.
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
	policy_value := idMap["policy"]

	var dataArr []map[string]interface{}

	findParams := service.FindParams{
		ResourceType:             service.Authenticationvserver_auditsyslogpolicy_binding.Type(),
		ResourceName:             name_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = r.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read authenticationvserver_auditsyslogpolicy_binding, got error: %s", err))
		return
	}

	// Binding (or its parent) no longer exists on the ADC. Signal removal via a null Id
	// (matches SDK v2 d.SetId("")) so the Read caller drops it from state instead of erroring.
	if len(dataArr) == 0 {
		data.Id = types.StringNull()
		return
	}

	// Iterate through results to find the binding matching policy.
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["policy"].(string); ok && val == policy_value {
			foundIndex = i
			break
		}
	}

	// Binding not present in the returned set: signal removal via a null Id (see above).
	if foundIndex == -1 {
		data.Id = types.StringNull()
		return
	}

	authenticationvserver_auditsyslogpolicy_bindingSetAttrFromGet(ctx, data, dataArr[foundIndex])
}
