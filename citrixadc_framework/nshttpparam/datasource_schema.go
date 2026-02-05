package nshttpparam

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func NshttpparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reuse server connections for requests from more than one client connections.",
			},
			"dropinvalreqs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop invalid HTTP requests or responses.",
			},
			"http2serverside": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable HTTP/2 on server side",
			},
			"ignoreconnectcodingscheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Ignore Coding scheme in CONNECT request.",
			},
			"insnssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable or disable Citrix ADC server header insertion for Citrix ADC generated HTTP responses.",
			},
			"logerrresp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Server header value to be inserted.",
			},
			"markconnreqinval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark CONNECT requests as invalid.",
			},
			"markhttp09inval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mark HTTP/0.9 requests as invalid.",
			},
			"maxreusepool": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum limit on the number of connections, from the Citrix ADC to a particular server that are kept in the reuse pool. This setting is helpful for optimal memory utilization and for reducing the idle connections to the server just after the peak time.",
			},
			"nssrvrhdr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The server header value to be inserted. If no explicit header is specified then NSBUILD.RELEASE is used as default server header.",
			},
		},
	}
}
