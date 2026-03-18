package nsassignment

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsassignmentResourceModel describes the resource data model.
type NsassignmentResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Add      types.String `tfsdk:"add"`
	Append   types.String `tfsdk:"append"`
	Clear    types.Bool   `tfsdk:"clear"`
	Comment  types.String `tfsdk:"comment"`
	Name     types.String `tfsdk:"name"`
	Newname  types.String `tfsdk:"newname"`
	Set      types.String `tfsdk:"set"`
	Sub      types.String `tfsdk:"sub"`
	Variable types.String `tfsdk:"variable"`
}

func (r *NsassignmentResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsassignment resource.",
			},
			"add": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and added to the left hand variable.",
			},
			"append": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and appended to the left hand variable.",
			},
			"clear": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Clear the variable value. Deallocates a text value, and for a map, the text key.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Can be used to preserve information about this rewrite action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the assignment is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my assignment\" or my assignment).",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the assignment.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the rewrite policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my assignment\" or my assignment).",
			},
			"set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and assigned to the left hand variable.",
			},
			"sub": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Right hand side of the assignment. The expression is evaluated and subtracted from the left hand variable.",
			},
			"variable": schema.StringAttribute{
				Required:    true,
				Description: "Left hand side of the assigment, of the form $variable-name (for a singleton variabled) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry",
			},
		},
	}
}

func nsassignmentGetThePayloadFromtheConfig(ctx context.Context, data *NsassignmentResourceModel) ns.Nsassignment {
	tflog.Debug(ctx, "In nsassignmentGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsassignment := ns.Nsassignment{}
	if !data.Add.IsNull() {
		nsassignment.Add = data.Add.ValueString()
	}
	if !data.Append.IsNull() {
		nsassignment.Append = data.Append.ValueString()
	}
	if !data.Clear.IsNull() {
		nsassignment.Clear = data.Clear.ValueBool()
	}
	if !data.Comment.IsNull() {
		nsassignment.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() {
		nsassignment.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		nsassignment.Newname = data.Newname.ValueString()
	}
	if !data.Set.IsNull() {
		nsassignment.Set = data.Set.ValueString()
	}
	if !data.Sub.IsNull() {
		nsassignment.Sub = data.Sub.ValueString()
	}
	if !data.Variable.IsNull() {
		nsassignment.Variable = data.Variable.ValueString()
	}

	return nsassignment
}

func nsassignmentSetAttrFromGet(ctx context.Context, data *NsassignmentResourceModel, getResponseData map[string]interface{}) *NsassignmentResourceModel {
	tflog.Debug(ctx, "In nsassignmentSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["Add"]; ok && val != nil {
		data.Add = types.StringValue(val.(string))
	} else {
		data.Add = types.StringNull()
	}
	if val, ok := getResponseData["append"]; ok && val != nil {
		data.Append = types.StringValue(val.(string))
	} else {
		data.Append = types.StringNull()
	}
	if val, ok := getResponseData["clear"]; ok && val != nil {
		data.Clear = types.BoolValue(val.(bool))
	} else {
		data.Clear = types.BoolNull()
	}
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
	if val, ok := getResponseData["set"]; ok && val != nil {
		data.Set = types.StringValue(val.(string))
	} else {
		data.Set = types.StringNull()
	}
	if val, ok := getResponseData["sub"]; ok && val != nil {
		data.Sub = types.StringValue(val.(string))
	} else {
		data.Sub = types.StringNull()
	}
	if val, ok := getResponseData["variable"]; ok && val != nil {
		data.Variable = types.StringValue(val.(string))
	} else {
		data.Variable = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
