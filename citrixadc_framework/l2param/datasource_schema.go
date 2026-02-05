package l2param

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func L2paramDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bdggrpproxyarp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set/reset proxy ARP in bridge group deployment",
			},
			"bdgsetting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bridging settings for C2C behavior. If enabled, each PE will learn MAC entries independently. Otherwise, when L2 mode is ON, learned MAC entries on a PE will be broadcasted to all other PEs.",
			},
			"bridgeagetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time-out value for the bridge table entries, in seconds. The new value applies only to the entries that are dynamically learned after the new value is set. Previously existing bridge table entries expire after the previously configured time-out value.",
			},
			"garponvridintf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send GARP messagess on VRID-configured interfaces upon failover",
			},
			"garpreply": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set/reset REPLY form of GARP",
			},
			"macmodefwdmypkt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allows MAC mode vserver to pick and forward the packets even if it is destined to Citrix ADC owned VIP.",
			},
			"maxbridgecollision": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bridge collision for loop detection",
			},
			"mbfinstlearning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable instant learning of MAC changes in MBF mode.",
			},
			"mbfpeermacupdate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "When mbf_instant_learning is enabled, learn any changes in peer's MAC after this time interval, which is in 10ms ticks.",
			},
			"proxyarp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Proxies the ARP as Citrix ADC MAC for FreeBSD.",
			},
			"returntoethernetsender": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Return to ethernet sender.",
			},
			"rstintfonhafo": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable the reset interface upon HA failover.",
			},
			"skipproxyingbsdtraffic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control source parameters (IP and Port) for FreeBSD initiated traffic. If Enabled, source parameters are retained. Else proxy the source parameters based on next hop.",
			},
			"stopmacmoveupdate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Stop Update of server mac change to NAT sessions.",
			},
			"usemymac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Citrix ADC MAC for all outgoing packets.",
			},
			"usenetprofilebsdtraffic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control source parameters (IP and Port) for FreeBSD initiated traffic. If enabled proxy the source parameters based on netprofile source ip. If netprofile does not have ip configured, then it will continue to use NSIP as earlier.",
			},
		},
	}
}
