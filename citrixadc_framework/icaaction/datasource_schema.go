package icaaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func IcaactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"accessprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica accessprofile to be associated with this action.",
			},
			"latencyprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica latencyprofile to be associated with this action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica action\" or 'my ica action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the ICA action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#),period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks ( for example, \"my ica action\" or 'my ica action').",
			},
		},
	}
}
