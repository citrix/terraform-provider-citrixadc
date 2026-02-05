package gslbparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func GslbparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"automaticconfigsync": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "GSLB configuration will be synced automatically to remote gslb sites if enabled.",
			},
			"dropldnsreq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop LDNS requests if round-trip time (RTT) information is not available.",
			},
			"gslbconfigsyncmonitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, remote gslb site's rsync port will be monitored and site is considered for configuration sync only when the monitor is successful.",
			},
			"gslbsvcstatedelaytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amount of delay in updating the state of GSLB service to DOWN when MEP goes down.\n			This parameter is applicable only if monitors are not bound to GSLB services",
			},
			"gslbsyncinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time duartion (in seconds) for which the gslb sync process will wait before checking for config changes.",
			},
			"gslbsynclocfiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If disabled, Location files will not be synced to the remote sites as part of manual sync and automatic sync.",
			},
			"gslbsyncmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Mode in which configuration will be synced from master site to remote sites.",
			},
			"gslbsyncsaveconfigcommand": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If enabled, 'save ns config' command will be treated as other GSLB commands and synced to GSLB nodes when auto gslb sync option is enabled.",
			},
			"ldnsentrytimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, after which an inactive LDNS entry is removed.",
			},
			"ldnsmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The IPv4 network mask with which to create LDNS entries.",
			},
			"ldnsprobeorder": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Order in which monitors should be initiated to calculate RTT.",
			},
			"mepkeepalivetimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time duartion (in seconds) during which if no new packets received by Local gslb site from Remote gslb site then mark the MEP connection DOWN",
			},
			"rtttolerance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Tolerance, in milliseconds, for newly learned round-trip time (RTT) values. If the difference between the old RTT value and the newly computed RTT value is less than or equal to the specified tolerance value, the LDNS entry in the network metric table is not updated with the new RTT value. Prevents the exchange of metrics when variations in RTT values are negligible.",
			},
			"svcstatelearningtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time (in seconds) within which local or child site services remain in learning phase. GSLB site will enter the learning phase after reboot, HA failover, Cluster GSLB owner node changes or MEP being enabled on local node.  Backup parent (if configured) will selectively move the adopted children's GSLB services to learning phase when primary parent goes down. While a service is in learning period, remote site will not honour the state and stats got through MEP for that service. State can be learnt from health monitor if bound explicitly.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Action to perform when policy evaluation creates an UNDEF condition. Available settings function as follows:\n* NOLBACTION - Does not consider LB action in making LB decision.\n* RESET - Reset the request and notify the user, so that the user can resend the request.\n* DROP - Drop the request without sending a response to the user.",
			},
			"v6ldnsmasklen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Mask for creating LDNS entries for IPv6 source addresses. The mask is defined as the number of leading bits to consider, in the source IP address, when creating an LDNS entry.",
			},
		},
	}
}
