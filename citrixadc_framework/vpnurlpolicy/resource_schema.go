package vpnurlpolicy

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnurlpolicyResourceModel describes the resource data model.
type VpnurlpolicyResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Action    types.String `tfsdk:"action"`
	Comment   types.String `tfsdk:"comment"`
	Logaction types.String `tfsdk:"logaction"`
	Name      types.String `tfsdk:"name"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`
}

func (r *VpnurlpolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnurlpolicy resource.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "Action to be applied by the new urlPolicy if the rule criteria are met.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this policy.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of messagelog action to use when a request matches this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new urlPolicy.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the vpn urlPolicy.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl policy\" or 'my vpnurl policy').",
			},
			"rule": schema.StringAttribute{
				Required:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the NetScaler CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
		},
	}
}

func vpnurlpolicyGetThePayloadFromtheConfig(ctx context.Context, data *VpnurlpolicyResourceModel) vpn.Vpnurlpolicy {
	tflog.Debug(ctx, "In vpnurlpolicyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnurlpolicy := vpn.Vpnurlpolicy{}
	if !data.Action.IsNull() {
		vpnurlpolicy.Action = data.Action.ValueString()
	}
	if !data.Comment.IsNull() {
		vpnurlpolicy.Comment = data.Comment.ValueString()
	}
	if !data.Logaction.IsNull() {
		vpnurlpolicy.Logaction = data.Logaction.ValueString()
	}
	if !data.Name.IsNull() {
		vpnurlpolicy.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		vpnurlpolicy.Newname = data.Newname.ValueString()
	}
	if !data.Rule.IsNull() {
		vpnurlpolicy.Rule = data.Rule.ValueString()
	}

	return vpnurlpolicy
}

func vpnurlpolicySetAttrFromGet(ctx context.Context, data *VpnurlpolicyResourceModel, getResponseData map[string]interface{}) *VpnurlpolicyResourceModel {
	tflog.Debug(ctx, "In vpnurlpolicySetAttrFromGet Function")

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
