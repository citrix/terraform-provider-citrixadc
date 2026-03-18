package nsassignment

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NsassignmentDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Optional:    true,
				Computed:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Left hand side of the assigment, of the form $variable-name (for a singleton variabled) or $variable-name[key-expression], where key-expression is an expression that evaluates to a text string and provides the key to select a map entry",
			},
		},
	}
}
