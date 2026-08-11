package cachepolicy

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/cache"
	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CachepolicyResourceModel describes the resource data model.
type CachepolicyResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Action       types.String `tfsdk:"action"`
	Invalgroups  types.List   `tfsdk:"invalgroups"`
	Invalobjects types.List   `tfsdk:"invalobjects"`
	Newname      types.String `tfsdk:"newname"`
	Policyname   types.String `tfsdk:"policyname"`
	Rule         types.String `tfsdk:"rule"`
	Storeingroup types.String `tfsdk:"storeingroup"`
	Undefaction  types.String `tfsdk:"undefaction"`
}

func (r *CachepolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cachepolicy resource.",
			},
			// SDK v2 backward-compat: action was Optional+Computed, not Required.
			// NITRO marks it mandatory for the add operation, but keeping it
			// Optional+Computed preserves the existing user-facing contract (a
			// missing value fails at NITRO create time, exactly as it did in v2).
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to apply to content that matches the policy.\n* CACHE or MAY_CACHE action - positive cachability policy\n* NOCACHE or MAY_NOCACHE action - negative cachability policy\n* INVAL action - Dynamic Invalidation Policy",
			},
			"invalgroups": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Content group(s) to be invalidated when the INVAL action is applied. Maximum number of content groups that can be specified is 16.",
			},
			"invalobjects": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Content groups(s) in which the objects will be invalidated if the action is INVAL.",
			},
			// newname is the rename trigger (NITRO ?action=rename). It is a pure
			// user input, never echoed back by GET, so it must NOT be Computed
			// (Computed causes known-after-apply churn / inconsistent-result) and
			// must NOT carry RequiresReplace (that would recreate the resource
			// instead of driving the in-place rename handled in Update).
			"newname": schema.StringAttribute{
				Optional:    true,
				Description: "New name for the cache policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			// SDK v2 backward-compat: policyname was Required + ForceNew.
			"policyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.",
			},
			// SDK v2 backward-compat: rule was Optional+Computed, not Required.
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which the traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			// Optional+Computed with an explicit NITRO default so removing it from
			// config produces a plan diff (drives the unset in Update). "DEFAULT" is
			// the value NITRO reports when storeingroup is unset.
			"storeingroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DEFAULT"),
				Description: "Name of the content group in which to store the object when the final result of policy evaluation is CACHE. The content group must exist before being mentioned here. Use the \"show cache contentgroup\" command to view the list of existing content groups.",
			},
			// undefaction is NOT unset-wired: its NITRO revert value ("Use Global") is
			// not a valid input value (NITRO rejects it in the update payload), so it
			// can be neither a stable Default nor round-tripped, unlike storeingroup.
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed when the result of rule evaluation is undefined.",
			},
		},
	}
}

// cachepolicyGetThePayloadFromthePlan builds the NITRO add/update body from the
// plan. newname is deliberately excluded - it is a rename-only argument handled
// via ?action=rename in Update, not part of the add/update POST/PUT body.
func cachepolicyGetThePayloadFromthePlan(ctx context.Context, data *CachepolicyResourceModel) cache.Cachepolicy {
	tflog.Debug(ctx, "In cachepolicyGetThePayloadFromthePlan Function")

	cachepolicy := cache.Cachepolicy{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		cachepolicy.Action = data.Action.ValueString()
	}
	if !data.Invalgroups.IsNull() && !data.Invalgroups.IsUnknown() {
		var invalgroupsList []string
		data.Invalgroups.ElementsAs(ctx, &invalgroupsList, false)
		cachepolicy.Invalgroups = invalgroupsList
	}
	if !data.Invalobjects.IsNull() && !data.Invalobjects.IsUnknown() {
		var invalobjectsList []string
		data.Invalobjects.ElementsAs(ctx, &invalobjectsList, false)
		cachepolicy.Invalobjects = invalobjectsList
	}
	// newname is rename-only (NITRO ?action=rename) and is excluded from the payload.
	if !data.Policyname.IsNull() && !data.Policyname.IsUnknown() {
		cachepolicy.Policyname = data.Policyname.ValueString()
	}
	if !data.Rule.IsNull() && !data.Rule.IsUnknown() {
		cachepolicy.Rule = data.Rule.ValueString()
	}
	if !data.Storeingroup.IsNull() && !data.Storeingroup.IsUnknown() {
		cachepolicy.Storeingroup = data.Storeingroup.ValueString()
	}
	if !data.Undefaction.IsNull() && !data.Undefaction.IsUnknown() {
		cachepolicy.Undefaction = data.Undefaction.ValueString()
	}

	return cachepolicy
}

// cachepolicySetAttrFromGet is the RESOURCE state setter. It preserves the
// user-facing key (policyname) and the rename-only newname so that a live rename
// (where the object name diverges from the configured policyname) does not
// clobber the plan/state values. The resource ID is managed by the caller
// (Create/Update/import), so this function never touches data.Id.
func cachepolicySetAttrFromGet(ctx context.Context, data *CachepolicyResourceModel, getResponseData map[string]interface{}) *CachepolicyResourceModel {
	tflog.Debug(ctx, "In cachepolicySetAttrFromGet Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["invalgroups"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(sliceVal))
			data.Invalgroups = listValue
		} else {
			data.Invalgroups = types.ListNull(types.StringType)
		}
	} else {
		data.Invalgroups = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["invalobjects"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(sliceVal))
			data.Invalobjects = listValue
		} else {
			data.Invalobjects = types.ListNull(types.StringType)
		}
	} else {
		data.Invalobjects = types.ListNull(types.StringType)
	}
	// policyname is the user-facing key. After a rename (via newname) the live
	// object name (tracked by data.Id) diverges from the configured policyname,
	// and GET returns the live (new) name. Only adopt the GET value when we do
	// not already have one (e.g. import, where state carries only the ID);
	// otherwise preserve the configured value to avoid a spurious replace diff.
	if data.Policyname.IsNull() || data.Policyname.IsUnknown() || data.Policyname.ValueString() == "" {
		if val, ok := getResponseData["policyname"]; ok && val != nil {
			data.Policyname = types.StringValue(val.(string))
		}
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["storeingroup"]; ok && val != nil {
		data.Storeingroup = types.StringValue(val.(string))
	} else {
		data.Storeingroup = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.

	return data
}

// cachepolicySetAttrFromGetForDatasource is the DATASOURCE state setter. The
// datasource has no prior plan/state to preserve, so it copies every field
// directly from the GET response and sets its own ID.
func cachepolicySetAttrFromGetForDatasource(ctx context.Context, data *CachepolicyResourceModel, getResponseData map[string]interface{}) *CachepolicyResourceModel {
	tflog.Debug(ctx, "In cachepolicySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["invalgroups"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(sliceVal))
			data.Invalgroups = listValue
		} else {
			data.Invalgroups = types.ListNull(types.StringType)
		}
	} else {
		data.Invalgroups = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["invalobjects"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(sliceVal))
			data.Invalobjects = listValue
		} else {
			data.Invalobjects = types.ListNull(types.StringType)
		}
	} else {
		data.Invalobjects = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["storeingroup"]; ok && val != nil {
		data.Storeingroup = types.StringValue(val.(string))
	} else {
		data.Storeingroup = types.StringNull()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Policyname.ValueString()))

	return data
}
