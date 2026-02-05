package lsnrtspalgprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LsnrtspalgprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"rtspalgprofilename": schema.StringAttribute{
				Required:    true,
				Description: "The name of the RTSPALG Profile.",
			},
			"rtspidletimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Idle timeout for the rtsp sessions in seconds.",
			},
			"rtspportrange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "port for the RTSP",
			},
			"rtsptransportprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RTSP ALG Profile transport protocol type.",
			},
		},
	}
}
