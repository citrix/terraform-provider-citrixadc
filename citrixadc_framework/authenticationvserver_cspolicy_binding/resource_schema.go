package authenticationvserver_cspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationvserverCspolicyBindingResourceModel describes the resource data model.
type AuthenticationvserverCspolicyBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupextraction        types.Bool   `tfsdk:"groupextraction"`
	Name                   types.String `tfsdk:"name"`
	Nextfactor             types.String `tfsdk:"nextfactor"`
	Policy                 types.String `tfsdk:"policy"`
	Priority               types.Int64  `tfsdk:"priority"`
	Secondary              types.Bool   `tfsdk:"secondary"`
}

func (r *AuthenticationvserverCspolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_cspolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the authentication virtual server to which to bind the policy.",
			},
			"nextfactor": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only while binding advance authentication policy as classic authentication policy does not support nFactor",
			},
			"policy": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the policy, if any, bound to the authentication vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The priority, if any, of the vpn vserver policy.",
			},
			"secondary": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
		},
	}
}

func authenticationvserver_cspolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverCspolicyBindingResourceModel) authentication.Authenticationvservercspolicybinding {
	tflog.Debug(ctx, "In authenticationvserver_cspolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_cspolicy_binding := authentication.Authenticationvservercspolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		authenticationvserver_cspolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		authenticationvserver_cspolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		authenticationvserver_cspolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_cspolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		authenticationvserver_cspolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		authenticationvserver_cspolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		authenticationvserver_cspolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		authenticationvserver_cspolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return authenticationvserver_cspolicy_binding
}

// authenticationvserver_cspolicy_bindingComposeId builds the composite ID using the
// SDK v2 / legacy key order (name,policy) so imported legacy state stays parseable
// and the ID stays backward compatible with resource_id_mapping.json.
func authenticationvserver_cspolicy_bindingComposeId(data *AuthenticationvserverCspolicyBindingResourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	return strings.Join(idParts, ",")
}

// authenticationvserver_cspolicy_bindingSetAttrFromGet is the resource-side state
// setter. The NITRO GET response for this binding only echoes name, policy and
// priority (priority as a string); bindpoint, gotopriorityexpression,
// groupextraction, nextfactor and secondary are write-only inputs that are never
// returned. Setting them to null here (the codegen default) clobbers the user's
// configured values and produces "inconsistent result after apply". So we preserve
// the existing plan/state for non-echoed fields and only adopt the values NITRO
// actually returns. (Pattern 7)
func authenticationvserver_cspolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_cspolicy_bindingSetAttrFromGet Function")

	// Echoed fields - adopt from GET response.
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}

	// Non-echoed fields (bindpoint, gotopriorityexpression, groupextraction,
	// nextfactor, secondary) are NOT returned by NITRO GET - preserve the
	// existing plan/state value rather than nulling it.

	// Set ID for the resource (legacy key order: name,policy)
	data.Id = types.StringValue(authenticationvserver_cspolicy_bindingComposeId(data))

	return data
}

// authenticationvserver_cspolicy_bindingSetAttrFromGetForDatasource faithfully
// copies every field the NITRO GET response provides. Datasources have no prior
// plan/state to preserve, so unlike the resource setter this simply adopts whatever
// GET returns and leaves the rest at config/null. It also sets data.Id since the
// datasource has no Create. (Pattern 7 datasource split)
func authenticationvserver_cspolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationvserverCspolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverCspolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_cspolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	// Set ID for the datasource (legacy key order: name,policy)
	data.Id = types.StringValue(authenticationvserver_cspolicy_bindingComposeId(data))

	return data
}
