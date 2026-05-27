package apispec

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ApispecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encrypted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the encrypted API spec. Must be in NetScaler format",
			},
			"file": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the api spec file. The spec file should be present on the appliance's hard-disk drive or solid-state drive. Storing a spec file in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/apispec/ is the default path.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the spec. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the spec is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my spec\" or 'my spec').",
			},
			"skipvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disabling openapi spec validation while adding it",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the spec file. The three formats supported by the appliance are:\nPROTO \nOAS/Swagger\nGRAPHQL",
			},
		},
	}
}
