package cachepolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cache"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			"action": schema.StringAttribute{
				Required:    true,
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
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the cache policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the policy is created.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression against which the traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"storeingroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the content group in which to store the object when the final result of policy evaluation is CACHE. The content group must exist before being mentioned here. Use the \"show cache contentgroup\" command to view the list of existing content groups.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to be performed when the result of rule evaluation is undefined.",
			},
		},
	}
}

func cachepolicyGetThePayloadFromtheConfig(ctx context.Context, data *CachepolicyResourceModel) cache.Cachepolicy {
	tflog.Debug(ctx, "In cachepolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cachepolicy := cache.Cachepolicy{}
	if !data.Action.IsNull() {
		cachepolicy.Action = data.Action.ValueString()
	}
	if !data.Newname.IsNull() {
		cachepolicy.Newname = data.Newname.ValueString()
	}
	if !data.Policyname.IsNull() {
		cachepolicy.Policyname = data.Policyname.ValueString()
	}
	if !data.Rule.IsNull() {
		cachepolicy.Rule = data.Rule.ValueString()
	}
	if !data.Storeingroup.IsNull() {
		cachepolicy.Storeingroup = data.Storeingroup.ValueString()
	}
	if !data.Undefaction.IsNull() {
		cachepolicy.Undefaction = data.Undefaction.ValueString()
	}

	return cachepolicy
}

func cachepolicySetAttrFromGet(ctx context.Context, data *CachepolicyResourceModel, getResponseData map[string]interface{}) *CachepolicyResourceModel {
	tflog.Debug(ctx, "In cachepolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Policyname.ValueString())

	return data
}
