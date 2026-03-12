package appfwfieldtype

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AppfwfieldtypeDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
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
				Optional:    true,
				Computed:    true,
				Description: "will not show internal field types added as part of FieldFormat learn rules deployment",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer specifying the priority of the field type. A lower number specifies a higher priority. Field types are checked in the order of their priority numbers.",
			},
			"regex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "PCRE - format regular expression defining the characters and length allowed for this field type.",
			},
		},
	}
}
