package policystringmap_pattern_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/policy"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// PolicystringmapPatternBindingResourceModel describes the resource data model.
type PolicystringmapPatternBindingResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Key     types.String `tfsdk:"key"`
	Name    types.String `tfsdk:"name"`
	Value   types.String `tfsdk:"value"`
}

func (r *PolicystringmapPatternBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the policystringmap_pattern_binding resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with the string map or key-value pair bound to this string map.",
			},
			"key": schema.StringAttribute{
				Required:    true,
				Description: "Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8.  UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\\xNN'. For example, the UTF-8 character '' can be encoded as '\\xC3\\xBC'.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the string map to which to bind the key-value pair.",
			},
			"value": schema.StringAttribute{
				Required:    true,
				Description: "Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.",
			},
		},
	}
}

func policystringmap_pattern_bindingGetThePayloadFromtheConfig(ctx context.Context, data *PolicystringmapPatternBindingResourceModel) policy.Policystringmappatternbinding {
	tflog.Debug(ctx, "In policystringmap_pattern_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	policystringmap_pattern_binding := policy.Policystringmappatternbinding{}
	if !data.Comment.IsNull() {
		policystringmap_pattern_binding.Comment = data.Comment.ValueString()
	}
	if !data.Key.IsNull() {
		policystringmap_pattern_binding.Key = data.Key.ValueString()
	}
	if !data.Name.IsNull() {
		policystringmap_pattern_binding.Name = data.Name.ValueString()
	}
	if !data.Value.IsNull() {
		policystringmap_pattern_binding.Value = data.Value.ValueString()
	}

	return policystringmap_pattern_binding
}

func policystringmap_pattern_bindingSetAttrFromGet(ctx context.Context, data *PolicystringmapPatternBindingResourceModel, getResponseData map[string]interface{}) *PolicystringmapPatternBindingResourceModel {
	tflog.Debug(ctx, "In policystringmap_pattern_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["key"]; ok && val != nil {
		data.Key = types.StringValue(val.(string))
	} else {
		data.Key = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["value"]; ok && val != nil {
		data.Value = types.StringValue(val.(string))
	} else {
		data.Value = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("key:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Key.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
