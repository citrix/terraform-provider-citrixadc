package tunneltrafficpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/tunnel"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// TunneltrafficpolicyResourceModel describes the resource data model.
type TunneltrafficpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *TunneltrafficpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the tunneltrafficpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Name of the built-in compression action to associate with the policy.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the tunnel traffic policy.\nMust begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy)'.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the tunnel traffic policy. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), e\nquals (=), and hyphen (-) characters.\nChoose a name that reflects the function that the policy performs.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my tunnel policy\" or 'my tunnel policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n*  If the expression itself includes double quotation marks, you must escape the quotations by using the \\ character. \n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func tunneltrafficpolicyGetThePayloadFromtheConfig(ctx context.Context, data *TunneltrafficpolicyResourceModel) tunnel.Tunneltrafficpolicy {
	tflog.Debug(ctx, "In tunneltrafficpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	tunneltrafficpolicy := tunnel.Tunneltrafficpolicy{}
	if !data.Action.IsNull() {
		tunneltrafficpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		tunneltrafficpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		tunneltrafficpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		tunneltrafficpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		tunneltrafficpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		tunneltrafficpolicy.Rule = data.Rule.ValueString()
	}

	return tunneltrafficpolicy
}

func tunneltrafficpolicySetAttrFromGet(ctx context.Context, data *TunneltrafficpolicyResourceModel, getResponseData map[string]interface{}) *TunneltrafficpolicyResourceModel {
	tflog.Debug(ctx, "In tunneltrafficpolicySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		data.Action = types.StringValue(val.(string))
	} else {
		data.Action = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
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
