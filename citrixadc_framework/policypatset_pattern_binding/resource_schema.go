package policypatset_pattern_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicypatsetPatternBindingResourceModel describes the resource data model.
type PolicypatsetPatternBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	String  types.String `tfsdk:"string"`
	Charset types.String `tfsdk:"charset"`
	Comment types.String `tfsdk:"comment"`
	Feature types.String `tfsdk:"feature"`
	Index   types.Int64  `tfsdk:"index"`
	Name    types.String `tfsdk:"name"`
}

func (r *PolicypatsetPatternBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policypatset_pattern_binding resource.",
			},
			"string": schema.StringAttribute{
				Required:    true,
				Description: "String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter.\nNote: Minimum length for pattern sets used in rewrite actions of type REPLACE_ALL, DELETE_ALL, INSERT_AFTER_ALL, and INSERT_BEFORE_ALL, is three characters.",
			},
			"charset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Character set associated with the characters in the string.\nNote: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\\xNN'. For example, the UTF-8 character '' can be encoded as '\\xC3\\xBC'.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this patset or a pattern bound to this patset.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index of the string associated with the patset.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the pattern set to which to bind the string.",
			},
		},
	}
}

func policypatset_pattern_bindingGetThePayloadFromtheConfig(ctx context.Context, data *PolicypatsetPatternBindingResourceModel) policy.Policypatsetpatternbinding {
	tflog.Debug(ctx, "In policypatset_pattern_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policypatset_pattern_binding := policy.Policypatsetpatternbinding{}
	if !data.String.IsNull() {
		policypatset_pattern_binding.String = data.String.ValueString()
	}
	if !data.Charset.IsNull() {
		policypatset_pattern_binding.Charset = data.Charset.ValueString()
	}
	if !data.Comment.IsNull() {
		policypatset_pattern_binding.Comment = data.Comment.ValueString()
	}
	if !data.Feature.IsNull() {
		policypatset_pattern_binding.Feature = data.Feature.ValueString()
	}
	if !data.Index.IsNull() {
		policypatset_pattern_binding.Index = utils.IntPtr(int(data.Index.ValueInt64()))
	}
	if !data.Name.IsNull() {
		policypatset_pattern_binding.Name = data.Name.ValueString()
	}

	return policypatset_pattern_binding
}

func policypatset_pattern_bindingSetAttrFromGet(ctx context.Context, data *PolicypatsetPatternBindingResourceModel, getResponseData map[string]interface{}) *PolicypatsetPatternBindingResourceModel {
	tflog.Debug(ctx, "In policypatset_pattern_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["String"]; ok && val != nil {
		data.String = types.StringValue(val.(string))
	} else {
		data.String = types.StringNull()
	}
	if val, ok := getResponseData["charset"]; ok && val != nil {
		data.Charset = types.StringValue(val.(string))
	} else {
		data.Charset = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
	if val, ok := getResponseData["index"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Index = types.Int64Value(intVal)
		}
	} else {
		data.Index = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
