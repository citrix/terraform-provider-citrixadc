package responderpolicylabel

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ResponderpolicylabelResourceModel describes the resource data model.
type ResponderpolicylabelResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Labelname       types.String `tfsdk:"labelname"`
	Newname         types.String `tfsdk:"newname"`
	Policylabeltype types.String `tfsdk:"policylabeltype"`
}

func (r *ResponderpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments to preserve information about this responder policy label.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder policy label\" or my responder policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"policylabeltype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("HTTP"),
				Description: "Type of responses sent by the policies bound to this policy label. Types are:\n* HTTP - HTTP responses.\n* OTHERTCP - NON-HTTP TCP responses.\n* SIP_UDP - SIP responses.\n* RADIUS - RADIUS responses.\n* MYSQL - SQL responses in MySQL format.\n* MSSQL - SQL responses in Microsoft SQL format.\n* NAT - NAT response.\n* MQTT - Trigger policies bind with MQTT type.\n* MQTT_JUMBO - Trigger policies bind with MQTT Jumbo type.",
			},
		},
	}
}

func responderpolicylabelGetThePayloadFromtheConfig(ctx context.Context, data *ResponderpolicylabelResourceModel) responder.Responderpolicylabel {
	tflog.Debug(ctx, "In responderpolicylabelGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderpolicylabel := responder.Responderpolicylabel{}
	if !data.Comment.IsNull() {
		responderpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() {
		responderpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Newname.IsNull() {
		responderpolicylabel.Newname = data.Newname.ValueString()
	}
	if !data.Policylabeltype.IsNull() {
		responderpolicylabel.Policylabeltype = data.Policylabeltype.ValueString()
	}

	return responderpolicylabel
}

func responderpolicylabelSetAttrFromGet(ctx context.Context, data *ResponderpolicylabelResourceModel, getResponseData map[string]interface{}) *ResponderpolicylabelResourceModel {
	tflog.Debug(ctx, "In responderpolicylabelSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["policylabeltype"]; ok && val != nil {
		data.Policylabeltype = types.StringValue(val.(string))
	} else {
		data.Policylabeltype = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Labelname.ValueString())

	return data
}
