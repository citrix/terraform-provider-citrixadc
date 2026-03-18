package lsnappsprofile_lsnappsattributes_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnappsprofileLsnappsattributesBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"appsattributesname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the LSN application port ATTRIBUTES command to bind to the specified LSN Appsprofile. Properties of the Appsprofile will be applicable to this APPSATTRIBUTES",
			},
			"appsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name for the LSN application profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the LSN application profile is created. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"lsn application profile1\" or 'lsn application profile1').",
			},
		},
	}
}
