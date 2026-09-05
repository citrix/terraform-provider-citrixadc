package subscriberprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscriberprofileDataSourceModel is the data-source-specific model, decoupled
// from SubscriberprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type SubscriberprofileDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Ip                  types.String `tfsdk:"ip"`
	Servicepath         types.String `tfsdk:"servicepath"`
	Subscriberrules     types.List   `tfsdk:"subscriberrules"`
	Subscriptionidtype  types.String `tfsdk:"subscriptionidtype"`
	Subscriptionidvalue types.String `tfsdk:"subscriptionidvalue"`
	Vlan                types.Int64  `tfsdk:"vlan"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/subscriberprofile.json). Never settable; populated from GET.
	Flags            types.Int64  `tfsdk:"flags"`
	Ttl              types.Int64  `tfsdk:"ttl"`
	Avpdisplaybuffer types.String `tfsdk:"avpdisplaybuffer"`
}

func SubscriberprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ip": schema.StringAttribute{
				Required:    true,
				Description: "Subscriber ip address",
			},
			"servicepath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the servicepath to be taken for this subscriber.",
			},
			"subscriberrules": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Rules configured for this subscriber. This is similar to rules received from PCRF for dynamic subscriber sessions.",
			},
			"subscriptionidtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id type",
			},
			"subscriptionidvalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscription-Id value",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "The vlan number on which the subscriber is located.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Subscriber Session flags.",
			},
			"ttl": schema.Int64Attribute{
				Computed:    true,
				Description: "Subscriber Session TTL.",
			},
			"avpdisplaybuffer": schema.StringAttribute{
				Computed:    true,
				Description: "Subscriber Attributes Display.",
			},
		},
	}
}

// subscriberprofileDataSourceSetAttrFromGet projects a NITRO subscriberprofile
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func subscriberprofileDataSourceSetAttrFromGet(ctx context.Context, data *SubscriberprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In subscriberprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["ip"]; ok && v != nil {
		data.Ip = types.StringValue(utils.AnyToString(v))
		data.Id = types.StringValue(utils.AnyToString(v))
	}

	data.Servicepath = utils.MapGetString(g, "servicepath")
	data.Subscriberrules = utils.MapGetStringList(g, "subscriberrules")
	data.Subscriptionidtype = utils.MapGetString(g, "subscriptionidtype")
	data.Subscriptionidvalue = utils.MapGetString(g, "subscriptionidvalue")
	data.Vlan = utils.MapGetInt64(g, "vlan")

	// Read-only attributes.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Avpdisplaybuffer = utils.MapGetString(g, "avpdisplaybuffer")
}
