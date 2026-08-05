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
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the assignment. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the assignment is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my assignment\" or my assignment).",
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
	if !data.Add.IsNull() && !data.Add.IsUnknown() {
		nsassignment.Add = data.Add.ValueString()
	}
	if !data.Append.IsNull() && !data.Append.IsUnknown() {
		nsassignment.Append = data.Append.ValueString()
	}
	if !data.Clear.IsNull() && !data.Clear.IsUnknown() {
		nsassignment.Clear = data.Clear.ValueBool()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nsassignment.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsassignment.Name = data.Name.ValueString()
	}
	if !data.Set.IsNull() && !data.Set.IsUnknown() {
		nsassignment.Set = data.Set.ValueString()
	}
	if !data.Sub.IsNull() && !data.Sub.IsUnknown() {
		nsassignment.Sub = data.Sub.ValueString()
	}
	if !data.Variable.IsNull() && !data.Variable.IsUnknown() {
		nsassignment.Variable = data.Variable.ValueString()
	}

	return nsassignment
}

func nsassignmentSetAttrFromGet(ctx context.Context, data *NsassignmentResourceModel, getResponseData map[string]interface{}) *NsassignmentResourceModel {
	tflog.Debug(ctx, "In nsassignmentSetAttrFromGet Function")

	// Convert API response to model.
	// NITRO omits some values from GET when they equal the NITRO default
	// (e.g. an empty string or a false boolean). To avoid clobbering a value
	// the user explicitly configured (omit-on-default trap), the else-branches
	// only null the attribute when the current value is unknown; otherwise the
	// previously-known (plan/state/config) value is preserved.

	// NITRO returns the "add" field using the JSON key "Add" (capitalised in the
	// nitro-go struct tag); fall back to the lowercase form for robustness.
	addVal, addOk := getResponseData["Add"]
	if !addOk || addVal == nil {
		addVal, addOk = getResponseData["add"]
	}
	if addOk && addVal != nil {
		data.Add = types.StringValue(addVal.(string))
	} else if data.Add.IsUnknown() {
		data.Add = types.StringNull()
	}

	if val, ok := getResponseData["append"]; ok && val != nil {
		data.Append = types.StringValue(val.(string))
	} else if data.Append.IsUnknown() {
		data.Append = types.StringNull()
	}
	if val, ok := getResponseData["clear"]; ok && val != nil {
		switch v := val.(type) {
		case bool:
			data.Clear = types.BoolValue(v)
		case string:
			data.Clear = types.BoolValue(v == "true" || v == "True")
		default:
			if data.Clear.IsUnknown() {
				data.Clear = types.BoolNull()
			}
		}
	} else if data.Clear.IsUnknown() {
		data.Clear = types.BoolNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["set"]; ok && val != nil {
		data.Set = types.StringValue(val.(string))
	} else if data.Set.IsUnknown() {
		data.Set = types.StringNull()
	}
	if val, ok := getResponseData["sub"]; ok && val != nil {
		data.Sub = types.StringValue(val.(string))
	} else if data.Sub.IsUnknown() {
		data.Sub = types.StringNull()
	}
	if val, ok := getResponseData["variable"]; ok && val != nil {
		data.Variable = types.StringValue(val.(string))
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
