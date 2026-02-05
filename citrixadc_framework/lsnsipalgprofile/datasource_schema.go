package lsnsipalgprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnsipalgprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"datasessionidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle timeout for the data channel sessions in seconds.",
			},
			"opencontactpinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE ContactPinhole creation.",
			},
			"openrecordroutepinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE RecordRoutePinhole creation.",
			},
			"openregisterpinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE RegisterPinhole creation.",
			},
			"openroutepinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE RoutePinhole creation.",
			},
			"openviapinhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE ViaPinhole creation.",
			},
			"registrationtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP registration timeout in seconds.",
			},
			"rport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ENABLE/DISABLE rport.",
			},
			"sipalgprofilename": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SIPALG Profile.",
			},
			"sipdstportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Destination port range for SIP_UDP and SIP_TCP.",
			},
			"sipsessiontimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP control channel session timeout in seconds.",
			},
			"sipsrcportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source port range for SIP_UDP and SIP_TCP.",
			},
			"siptransportprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SIP ALG Profile transport protocol type.",
			},
		},
	}
}
