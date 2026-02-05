package authenticationauthnprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationauthnprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"authenticationdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain for which TM cookie must to be set. If unspecified, cookie will be set for FQDN.",
			},
			"authenticationhost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname of the authentication vserver to which user must be redirected for authentication.",
			},
			"authenticationlevel": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication weight or level of the vserver to which this will bound. This is used to order TM vservers based on the protection required. A session that is created by authenticating against TM vserver at given level cannot be used to access TM vserver at a higher level.",
			},
			"authnvsname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the authentication vserver at which authentication should be done.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the authentication profile.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the RADIUS action is added.",
			},
		},
	}
}
