package subscribersessions

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// SubscribersessionsDataSourceModel is the datasource-only model. It
// intentionally does NOT reuse the resource model (SubscribersessionsResource
// Model): the resource is an action-only clear with a minimal schema (id, ip,
// vlan), while the datasource is backed by get(all) and exposes the read-only
// subscriber-session telemetry.
//
// Filters mirror the get(all) args: ip (String), vlan (Double/int), nodeid
// (Double/int, cluster-only GET filter). Per the NITRO doc get response,
// flags/ttl/idlettl are Double and map to types.Int64 (populated via
// utils.ConvertToInt64, which tolerates the float64 JSON wire format).
// subscriptionidtype/subscriptionidvalue/subscriberrules/avpdisplaybuffer/
// servicepath are strings. subscriberrules is documented as String[] but the
// vendored NITRO struct types it as string; it is copied only when the get(all)
// response yields a string value.
type SubscribersessionsDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	Ip     types.String `tfsdk:"ip"`
	Vlan   types.Int64  `tfsdk:"vlan"`
	Nodeid types.Int64  `tfsdk:"nodeid"`

	// Read-only subscriber session telemetry (Computed) from get(all).
	Subscriptionidtype  types.String `tfsdk:"subscriptionidtype"`
	Subscriptionidvalue types.String `tfsdk:"subscriptionidvalue"`
	Subscriberrules     types.String `tfsdk:"subscriberrules"`
	Flags               types.Int64  `tfsdk:"flags"`
	Ttl                 types.Int64  `tfsdk:"ttl"`
	Idlettl             types.Int64  `tfsdk:"idlettl"`
	Avpdisplaybuffer    types.String `tfsdk:"avpdisplaybuffer"`
	Servicepath         types.String `tfsdk:"servicepath"`
}

func SubscribersessionsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			// Optional get(all) filters.
			"ip": schema.StringAttribute{
				Optional:    true,
				Description: "Subscriber IP Address.",
			},
			"vlan": schema.Int64Attribute{
				Optional:    true,
				Description: "The vlan number on which the subscriber is located.",
			},
			"nodeid": schema.Int64Attribute{
				Optional:    true,
				Description: "Unique number that identifies the cluster node.",
			},

			// Read-only subscriber session telemetry from get(all).
			"subscriptionidtype": schema.StringAttribute{
				Computed:    true,
				Description: "Subscription-Id type. Possible values = E164, IMSI, SIP_URI, NAI, PRIVATE.",
			},
			"subscriptionidvalue": schema.StringAttribute{
				Computed:    true,
				Description: "Subscription-Id value.",
			},
			"subscriberrules": schema.StringAttribute{
				Computed:    true,
				Description: "Rules stored in this session for this subscriber.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Subscriber Session flags.",
			},
			"ttl": schema.Int64Attribute{
				Computed:    true,
				Description: "Subscriber Session revalidation Timeout remaining.",
			},
			"idlettl": schema.Int64Attribute{
				Computed:    true,
				Description: "Subscriber Session Activity Timeout remaining.",
			},
			"avpdisplaybuffer": schema.StringAttribute{
				Computed:    true,
				Description: "Subscriber Attributes Display.",
			},
			"servicepath": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the servicepath to be taken for this subscriber.",
			},
		},
	}
}
