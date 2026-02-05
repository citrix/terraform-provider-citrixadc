package csaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/cs"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsactionResourceModel describes the resource data model.
type CsactionResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	Name              types.String `tfsdk:"name"`
	Newname           types.String `tfsdk:"newname"`
	Targetlbvserver   types.String `tfsdk:"targetlbvserver"`
	Targetvserver     types.String `tfsdk:"targetvserver"`
	Targetvserverexpr types.String `tfsdk:"targetvserverexpr"`
}

func (r *CsactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the csaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with this cs action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Can be changed after the content switching action is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my action\" or 'my action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the content switching action. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing virtual server to which the content is switched.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the VPN, GSLB or Authentication virtual server to which the content is switched.",
			},
			"targetvserverexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Information about this content switching action.",
			},
		},
	}
}

func csactionGetThePayloadFromtheConfig(ctx context.Context, data *CsactionResourceModel) cs.Csaction {
	tflog.Debug(ctx, "In csactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	csaction := cs.Csaction{}
	if !data.Comment.IsNull() {
		csaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		csaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		csaction.Newname = data.Newname.ValueString()
	}
	if !data.Targetlbvserver.IsNull() {
		csaction.Targetlbvserver = data.Targetlbvserver.ValueString()
	}
	if !data.Targetvserver.IsNull() {
		csaction.Targetvserver = data.Targetvserver.ValueString()
	}
	if !data.Targetvserverexpr.IsNull() {
		csaction.Targetvserverexpr = data.Targetvserverexpr.ValueString()
	}

	return csaction
}

func csactionSetAttrFromGet(ctx context.Context, data *CsactionResourceModel, getResponseData map[string]interface{}) *CsactionResourceModel {
	tflog.Debug(ctx, "In csactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
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
	if val, ok := getResponseData["targetlbvserver"]; ok && val != nil {
		data.Targetlbvserver = types.StringValue(val.(string))
	} else {
		data.Targetlbvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserver"]; ok && val != nil {
		data.Targetvserver = types.StringValue(val.(string))
	} else {
		data.Targetvserver = types.StringNull()
	}
	if val, ok := getResponseData["targetvserverexpr"]; ok && val != nil {
		data.Targetvserverexpr = types.StringValue(val.(string))
	} else {
		data.Targetvserverexpr = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
