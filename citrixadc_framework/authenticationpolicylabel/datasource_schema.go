package authenticationpolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this authentication policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authentication policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy label\" or 'authentication policy label').",
			},
			"loginschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is \"noschema\" should be used.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the auth policy label",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of feature (aaatm or rba) against which to match the policies bound to this policy label.",
			},
		},
	}
}
