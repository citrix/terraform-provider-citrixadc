package authenticationvserver_cachepolicy_binding

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

// AuthenticationvserverCachepolicyBindingResourceModel describes the resource data model.
type AuthenticationvserverCachepolicyBindingResourceModel struct {
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

func (r *AuthenticationvserverCachepolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationvserver_cachepolicy_binding resource.",
			},
			"bindpoint": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Bindpoint to which the policy is bound.",
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
				// Not echoed by NITRO GET; drop Computed so it resolves to null when
				// unset instead of staying unknown after apply (Pattern 13).
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
				// Not echoed by NITRO GET; drop Computed so it resolves to null when
				// unset instead of staying unknown after apply (Pattern 13).
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
				// Not echoed by NITRO GET; drop Computed so it resolves to null when
				// unset instead of staying unknown after apply (Pattern 13).
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only while bindind classic authentication policy as advance authentication policy use nFactor",
			},
		},
	}
}

func authenticationvserver_cachepolicy_bindingGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationvserverCachepolicyBindingResourceModel) authentication.Authenticationvservercachepolicybinding {
	tflog.Debug(ctx, "In authenticationvserver_cachepolicy_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationvserver_cachepolicy_binding := authentication.Authenticationvservercachepolicybinding{}
	if !data.Bindpoint.IsNull() && !data.Bindpoint.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Bindpoint = data.Bindpoint.ValueString()
	}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Groupextraction.IsNull() && !data.Groupextraction.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Groupextraction = data.Groupextraction.ValueBool()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Name = data.Name.ValueString()
	}
	if !data.Nextfactor.IsNull() && !data.Nextfactor.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Nextfactor = data.Nextfactor.ValueString()
	}
	if !data.Policy.IsNull() && !data.Policy.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Policy = data.Policy.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Secondary.IsNull() && !data.Secondary.IsUnknown() {
		authenticationvserver_cachepolicy_binding.Secondary = data.Secondary.ValueBool()
	}

	return authenticationvserver_cachepolicy_binding
}

// authenticationvserver_cachepolicy_bindingSetAttrFromGet is the RESOURCE setter.
// It preserves user-configured plan/state values for attributes the NITRO GET does NOT
// echo back (secondary, groupextraction, nextfactor) and adopts only the fields that
// NITRO actually returns, to avoid "inconsistent result after apply" diffs. The ID is
// set once in Create (name:policy) and must NOT be recomputed here (Pattern 6).
func authenticationvserver_cachepolicy_bindingSetAttrFromGet(ctx context.Context, data *AuthenticationvserverCachepolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverCachepolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_cachepolicy_bindingSetAttrFromGet Function")

	// Convert API response to model (resource: only adopt fields the GET echoes)
	if val, ok := getResponseData["bindpoint"]; ok && val != nil {
		data.Bindpoint = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	}
	// groupextraction is not echoed by NITRO GET - preserve the plan/state value (Pattern 7)
	if val, ok := getResponseData["groupextraction"]; ok && val != nil {
		data.Groupextraction = types.BoolValue(val.(bool))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	// nextfactor is not echoed by NITRO GET - preserve the plan/state value (Pattern 7)
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
	// secondary is not echoed by NITRO GET - preserve the plan/state value (Pattern 7)
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	}

	// ID is set once in Create as name:policy - do not recompute here (Pattern 6).
	return data
}

// authenticationvserver_cachepolicy_bindingSetAttrFromGetForDatasource is the DATASOURCE
// setter. The datasource has no prior state to preserve, so it faithfully copies every
// field the GET response provides and sets its own ID (Pattern 7 datasource split).
func authenticationvserver_cachepolicy_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationvserverCachepolicyBindingResourceModel, getResponseData map[string]interface{}) *AuthenticationvserverCachepolicyBindingResourceModel {
	tflog.Debug(ctx, "In authenticationvserver_cachepolicy_bindingSetAttrFromGetForDatasource Function")

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
		} else {
			data.Priority = types.Int64Null()
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["secondary"]; ok && val != nil {
		data.Secondary = types.BoolValue(val.(bool))
	} else {
		data.Secondary = types.BoolNull()
	}

	// Datasource has no Create - compose the ID here (name:policy, matching the resource).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
