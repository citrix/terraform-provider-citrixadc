package authenticationadfsproxyprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func AuthenticationadfsproxyprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"certkeyname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSL certificate of the proxy that is registered at adfs server for trust.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the adfs proxy profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my push service\" or 'my push service').",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Description: "This is the password of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a password_wo update.",
			},
			"serverurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified url of the adfs server.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the name of an account in directory that would be used to authenticate trust request from ADC acting as a proxy.",
			},
		},
	}
}
