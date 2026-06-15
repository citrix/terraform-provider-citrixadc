package authenticationvserver_authenticationpolicy_binding

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

// AuthenticationvserverAuthenticationpolicyBindingResourceModel describes the resource data model.
type AuthenticationvserverAuthenticationpolicyBindingResourceModel struct {
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

func (r *AuthenticationvserverAuthenticationpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_authenticationpolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bind point to which to bind the policy. Applies only to rewrite and cache policies. If you do not set this parameter, the policy is bound to REQ_DEFAULT or RES_DEFAULT, depending on whether the policy rule is a response-time or a request-time expression.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"groupextraction": schema.BoolAttribute{
				Optional: true,
				Computed: true,
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "On success invoke label.",
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
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
		},
	}
}

func authenticationvserver_authenticationpolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverAuthenticationpolicyBindingResourceModel) authentication.Authenticationvserverauthenticationpolicybinding {
	tflog.Debug(ctx, "In authenticationvserver_authenticationpolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_authenticationpolicy_binding := authentication.Authenticationvserverauthenticationpolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		authenticationvserver_authenticationpolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return authenticationvserver_authenticationpolicy_binding
}

// authenticationvserver_authenticationpolicy_bindingSetAttrFromGet is the resource-side
// state setter. The ID is composed exactly once in Create and is NOT recomputed here
// (Pattern 6). The Optional+Computed attributes are resolved to concrete values: the GET
// value when echoed, otherwise Null — this is required so that Terraform does not see an
// "unknown value after apply" for Computed attributes the binding GET does not return
// (bindpoint/groupextraction/nextfactor/secondary are not echoed by NITRO; this mirrors
// the SDK v2 d.Set(nil)->null behavior).
func authenticationvserver_authenticationpolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverAuthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationpolicy_bindingSetAttrFromGet Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	} else {
		data.Nextfactor = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	return data
}

// authenticationvserver_authenticationpolicy_bindingSetAttrFromGetForDatasource is the
// datasource-side state setter. The datasource has no prior plan/state to preserve, so it
// faithfully copies every field from the GET response and sets the ID itself (Pattern 7
// datasource split).
func authenticationvserver_authenticationpolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationvserverAuthenticationpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverAuthenticationpolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_authenticationpolicy_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	} else {
		data.Bindpoint = types.StringNull()
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	} else {
		data.Groupextraction = types.BoolNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["nextfactor"]; ok && val != nil {
		data.Nextfactor = types.StringValue(val.(string))
	} else {
		data.Nextfactor = types.StringNull()
	}
	if val, ok := getResponseData["policy"]; ok && val != nil {
		data.Policy = types.StringValue(val.(string))
	} else {
		data.Policy = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Set ID for the datasource - matches the legacy SDK v2 / Create composite order (name,policy)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(data.Name.ValueString())))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(data.Policy.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
