package aaagroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaagroupResourceModel describes the resource data model.
type AaagroupResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Groupname types.String `tfsdk:"groupname"`
	Loggedin  types.Bool   `tfsdk:"loggedin"`
	Weight    types.Int64  `tfsdk:"weight"`
}

func (r *AaagroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaagroup resource.",
			},
			"groupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the group. Must begin with a letter, number, or the underscore character (_), and must consist only of letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore  characters. Cannot be changed after the group is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my aaa group\" or 'my aaa group').",
			},
			"loggedin": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display only the group members who are currently logged in. If there are large number of sessions, this command may provide partial details.",
			},
			"weight": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Weight of this group with respect to other configured aaa groups (lower the number higher the weight)",
			},
		},
	}
}

func aaagroupGetThePayloadFromtheConfig(ctx context.Context, data *AaagroupResourceModel) aaa.Aaagroup {
	tflog.Debug(ctx, "In aaagroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaagroup := aaa.Aaagroup{}
	if !data.Groupname.IsNull() {
		aaagroup.Groupname = data.Groupname.ValueString()
	}
	if !data.Loggedin.IsNull() {
		aaagroup.Loggedin = data.Loggedin.ValueBool()
	}
	if !data.Weight.IsNull() {
		aaagroup.Weight = utils.IntPtr(int(data.Weight.ValueInt64()))
	}

	return aaagroup
}

func aaagroupSetAttrFromGet(ctx context.Context, data *AaagroupResourceModel, getResponseData map[string]interface{}) *AaagroupResourceModel {
	tflog.Debug(ctx, "In aaagroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	} else {
		data.Groupname = types.StringNull()
	}
	if val, ok := getResponseData["loggedin"]; ok && val != nil {
		data.Loggedin = types.BoolValue(val.(bool))
	} else {
		data.Loggedin = types.BoolNull()
	}
	if val, ok := getResponseData["weight"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Weight = types.Int64Value(intVal)
		}
	} else {
		data.Weight = types.Int64Null()
	}

	// Set ID for the resource
	data.Id = types.StringValue(data.Groupname.ValueString())

	return data
}
