package nslicenseserver

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NslicenseserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"deviceprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device profile is created on ADM and contains the user name and password of the instance(s). ADM will use this info to add the NS for registration",
			},
			"forceupdateip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If this flag is used while adding the licenseserver, existing config will be overwritten. Use this flag only if you are sure that the new licenseserver has the required capacity.",
			},
			"licensemode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This paramter indicates type of license customer interested while configuring add/set licenseserver",
			},
			"licenseserverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of the License server. Either licenseserverip or servername must be specified.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "License server port.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully qualified domain name of the License server. Either licenseserverip or servername must be specified.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
		},
	}
}
