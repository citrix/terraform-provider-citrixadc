package cmppolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cmp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CmppolicyResourceModel describes the resource data model.
type CmppolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Resaction types.String `tfsdk:"resaction"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *CmppolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the cmppolicy resource.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the HTTP compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nCan be changed after the policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the compression policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my cmp policy\" or 'my cmp policy').",
			},
			"resaction": schema.StringAttribute{
				Required:    true,
				Description: "The built-in or user-defined compression action to apply to the response when the policy matches a request or response.",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression that determines which HTTP requests or responses match the compression policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func cmppolicyGetThePayloadFromtheConfig(ctx context.Context, data *CmppolicyResourceModel) cmp.Cmppolicy {
	tflog.Debug(ctx, "In cmppolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	cmppolicy := cmp.Cmppolicy{}
	if !data.Name.IsNull() {
		cmppolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		cmppolicy.Newname = data.Newname.ValueString()
	}
	if !data.Resaction.IsNull() {
		cmppolicy.Resaction = data.Resaction.ValueString()
	}
	if !data.Rule.IsNull() {
		cmppolicy.Rule = data.Rule.ValueString()
	}

	return cmppolicy
}

func cmppolicySetAttrFromGet(ctx context.Context, data *CmppolicyResourceModel, getResponseData map[string]interface{}) *CmppolicyResourceModel {
	tflog.Debug(ctx, "In cmppolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["resaction"]; ok && val != nil {
		data.Resaction = types.StringValue(val.(string))
	} else {
		data.Resaction = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
