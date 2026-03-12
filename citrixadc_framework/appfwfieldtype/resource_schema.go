package appfwfieldtype

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwfieldtypeResourceModel describes the resource data model.
type AppfwfieldtypeResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Name       types.String `tfsdk:"name"`
	Nocharmaps types.Bool   `tfsdk:"nocharmaps"`
	Priority   types.Int64  `tfsdk:"priority"`
	Regex      types.String `tfsdk:"regex"`
}

func (r *AppfwfieldtypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwfieldtype resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment describing the type of field that this field type is intended to match.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the field type.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the field type is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my field type\" or 'my field type').",
			},
			"nocharmaps": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "will not show internal field types added as part of FieldFormat learn rules deployment",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.",
			},
			"regex": schema.StringAttribute{
				Required:    true,
				Description: "PCRE - format regular expression defining the characters and length allowed for this field type.",
			},
		},
	}
}

func appfwfieldtypeGetThePayloadFromtheConfig(ctx context.Context, data *AppfwfieldtypeResourceModel) appfw.Appfwfieldtype {
	tflog.Debug(ctx, "In appfwfieldtypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	appfwfieldtype := appfw.Appfwfieldtype{}
	if !data.Comment.IsNull() {
		appfwfieldtype.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		appfwfieldtype.Name = data.Name.ValueString()
	}
	if !data.Nocharmaps.IsNull() {
		appfwfieldtype.Nocharmaps = data.Nocharmaps.ValueBool()
	}
	if !data.Priority.IsNull() {
		appfwfieldtype.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Regex.IsNull() {
		appfwfieldtype.Regex = data.Regex.ValueString()
	}

	return appfwfieldtype
}

func appfwfieldtypeSetAttrFromGet(ctx context.Context, data *AppfwfieldtypeResourceModel, getResponseData map[string]interface{}) *AppfwfieldtypeResourceModel {
	tflog.Debug(ctx, "In appfwfieldtypeSetAttrFromGet Function")

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
	if val, ok := getResponseData["nocharmaps"]; ok && val != nil {
		data.Nocharmaps = types.BoolValue(val.(bool))
	} else {
		data.Nocharmaps = types.BoolNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["regex"]; ok && val != nil {
		data.Regex = types.StringValue(val.(string))
	} else {
		data.Regex = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
