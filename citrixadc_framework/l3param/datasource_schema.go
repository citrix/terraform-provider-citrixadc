package l3param

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func L3paramDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"acllogtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to tune acl logging time",
			},
			"allowclasseipv4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable IPv4 Class E address clients",
			},
			"dropdfflag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dropping the IP DF flag.",
			},
			"dropipfragments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable dropping of IP fragments.",
			},
			"dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Dynamic routing on partition. This configuration is not applicable to default partition",
			},
			"externalloopback": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable external loopback.",
			},
			"forwardicmpfragments": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable forwarding of ICMP fragments.",
			},
			"icmpgenratethreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NS generated ICMP pkts per 10ms rate threshold",
			},
			"implicitaclallow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Do not apply ACLs for internal ports",
			},
			"implicitpbr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable Policy Based Routing for control packets",
			},
			"ipv6dynamicrouting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable IPv6 Dynamic routing",
			},
			"miproundrobin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable round robin usage of mapped IPs.",
			},
			"overridernat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "USNIP/USIP settings override RNAT settings for configured\n              service/virtual server traffic..",
			},
			"srcnat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Perform NAT if only the source is in the private network",
			},
			"tnlpmtuwoconn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable learning PMTU of IP tunnel when ICMP error does not contain connection information.",
			},
			"usipserverstraypkt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable detection of stray server side pkts in USIP mode.",
			},
		},
	}
}
