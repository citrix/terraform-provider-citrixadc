package lbparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LBParameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"allowboundsvcremoval": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"computedadccookieattribute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"consolidatedlconn": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"cookiepassphrase": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"dbsttl": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"dropmqttjumbomessage": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"httponlycookieflag": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"literaladccookieattribute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"maxpipelinenat": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"monitorconnectionclose": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"monitorskipmaxclient": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"preferdirectroute": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"proximityfromself": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"retainservicestate": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"startuprrfactor": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"storemqttclientidandusername": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"sessionsthreshold": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
			"undefaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useencryptedpersistencecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"useportforhashlb": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"usesecuredpersistencecookie": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"vserverspecificmac": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"lbhashalgorithm": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"lbhashfingers": schema.Int64Attribute{
				Optional: true,
				Computed: true,
			},
		},
	}
}
