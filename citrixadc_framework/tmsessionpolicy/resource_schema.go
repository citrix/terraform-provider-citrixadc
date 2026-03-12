package tmsessionpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tm"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TmsessionpolicyResourceModel describes the resource data model.
type TmsessionpolicyResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Action types.String `tfsdk:"action"`
	Name   types.String `tfsdk:"name"`
	Rule   types.String `tfsdk:"rule"`
}

func (r *TmsessionpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tmsessionpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to be applied to connections that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the session policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a session policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, against which traffic is evaluated. Both classic and advance expressions are supported in default partition but only advance expressions in non-default partition.\n\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func tmsessionpolicyGetThePayloadFromtheConfig(ctx context.Context, data *TmsessionpolicyResourceModel) tm.Tmsessionpolicy {
	tflog.Debug(ctx, "In tmsessionpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tmsessionpolicy := tm.Tmsessionpolicy{}
	if !data.Action.IsNull() {
		tmsessionpolicy.Action = data.Action.ValueString()
	}
	if !data.Name.IsNull() {
		tmsessionpolicy.Name = data.Name.ValueString()
	}
	if !data.Rule.IsNull() {
		tmsessionpolicy.Rule = data.Rule.ValueString()
	}

	return tmsessionpolicy
}

func tmsessionpolicySetAttrFromGet(ctx context.Context, data *TmsessionpolicyResourceModel, getResponseData map[string]interface{}) *TmsessionpolicyResourceModel {
	tflog.Debug(ctx, "In tmsessionpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
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
