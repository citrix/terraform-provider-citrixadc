package lbaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// LbactionResourceModel describes the resource data model.
type LbactionResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Comment types.String `tfsdk:"comment"`
	Name    types.String `tfsdk:"name"`
	Newname types.String `tfsdk:"newname"`
	Type    types.String `tfsdk:"type"`
	Value   types.List   `tfsdk:"value"`
}

func (r *LbactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this LB action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LB action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or 'my lb action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the LB action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my lb action\" or my lb action').",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of an LB action. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* SELECTIONORDER - services bound to vserver with order specified in value parameter is considerd for lb/gslb decision.",
			},
			"value": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The selection order list used during lb/gslb decision. Preference of services during lb/gslb decision is as follows - services corresponding to first order specified in the sequence is considered first, services corresponding to second order specified in the sequence is considered next and so on. For example, if -value 2 1 3 is specified here and service-1 bound to a vserver with order 1, service-2 bound to a vserver with order 2 and  service-3 bound to a vserver with order 3. Then preference of selecting services in LB decision is as follows: service-2, service-1, service-3.",
			},
		},
	}
}

func lbactionGetThePayloadFromtheConfig(ctx context.Context, data *LbactionResourceModel) lb.Lbaction {
	tflog.Debug(ctx, "In lbactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbaction := lb.Lbaction{}
	if !data.Comment.IsNull() {
		lbaction.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		lbaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		lbaction.Newname = data.Newname.ValueString()
	}
	if !data.Type.IsNull() {
		lbaction.Type = data.Type.ValueString()
	}

	return lbaction
}

func lbactionSetAttrFromGet(ctx context.Context, data *LbactionResourceModel, getResponseData map[string]interface{}) *LbactionResourceModel {
	tflog.Debug(ctx, "In lbactionSetAttrFromGet Function")

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
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
