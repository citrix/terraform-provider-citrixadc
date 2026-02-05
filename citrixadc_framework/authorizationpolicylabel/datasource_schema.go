package authorizationpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthorizationpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authorization policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy label\" or 'authorization policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the auth policy label",
			},
		},
	}
}
