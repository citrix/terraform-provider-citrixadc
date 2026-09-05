package subscriberparam

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SubscriberparamDataSourceModel is the data-source-specific model, decoupled
// from SubscriberparamResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// The Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares.
type SubscriberparamDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Idleaction           types.String `tfsdk:"idleaction"`
	Idlettl              types.Int64  `tfsdk:"idlettl"`
	Interfacetype        types.String `tfsdk:"interfacetype"`
	Ipv6prefixlookuplist types.List   `tfsdk:"ipv6prefixlookuplist"`
	Keytype              types.String `tfsdk:"keytype"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/subscriberparam.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func SubscriberparamDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"idleaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Once idleTTL exprires on a subscriber session, Citrix ADC will take an idle action on that session. idleAction could be chosen from one of these ==>\n1. ccrTerminate: (default) send CCR-T to inform PCRF about session termination and delete the session.  \n2. delete: Just delete the subscriber session without informing PCRF.\n3. ccrUpdate: Do not delete the session and instead send a CCR-U to PCRF requesting for an updated session. !",
			},
			"idlettl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "q!Idle Timeout, in seconds, after which Citrix ADC will take an idleAction on a subscriber session (refer to 'idleAction' arguement in 'set subscriber param' for more details on idleAction). Any data-plane or control plane activity updates the idleTimeout on subscriber session. idleAction could be to 'just delete the session' or 'delete and CCR-T' (if PCRF is configured) or 'do not delete but send a CCR-U'. \nZero value disables the idle timeout. !",
			},
			"interfacetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Subscriber Interface refers to Citrix ADC interaction with control plane protocols, RADIUS and GX.\nTypes of subscriber interface: NONE, RadiusOnly, RadiusAndGx, GxOnly.\nNONE: Only static subscribers can be configured.\nRadiusOnly: GX interface is absent. Subscriber information is obtained through RADIUS Accounting messages.\nRadiusAndGx: Subscriber ID obtained through RADIUS Accounting is used to query PCRF. Subscriber information is obtained from both RADIUS and PCRF.\nGxOnly: RADIUS interface is absent. Subscriber information is queried using Subscriber IP or IP+VLAN.",
			},
			"ipv6prefixlookuplist": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "The ipv6PrefixLookupList should consist of all the ipv6 prefix lengths assigned to the UE's'",
			},
			"keytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of subscriber key type IP or IPANDVLAN. IPANDVLAN option can be used only when the interfaceType is set to gxOnly.\nChanging the lookup method should result to the subscriber session database being flushed.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether the configuration is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// subscriberparamDataSourceSetAttrFromGet projects a NITRO subscriberparam GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func subscriberparamDataSourceSetAttrFromGet(ctx context.Context, data *SubscriberparamDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In subscriberparamDataSourceSetAttrFromGet Function")

	data.Idleaction = utils.MapGetString(g, "idleaction")
	data.Idlettl = utils.MapGetInt64(g, "idlettl")
	data.Interfacetype = utils.MapGetString(g, "interfacetype")
	data.Keytype = utils.MapGetString(g, "keytype")

	// ipv6prefixlookuplist is an Int64-typed list on the schema; convert explicitly.
	if val, ok := g["ipv6prefixlookuplist"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			int64List := make([]int64, 0, len(sliceVal))
			for _, item := range sliceVal {
				if intVal, err := utils.ConvertToInt64(item); err == nil {
					int64List = append(int64List, intVal)
				}
			}
			listValue, _ := types.ListValueFrom(ctx, types.Int64Type, int64List)
			data.Ipv6prefixlookuplist = listValue
		} else {
			data.Ipv6prefixlookuplist = types.ListNull(types.Int64Type)
		}
	} else {
		data.Ipv6prefixlookuplist = types.ListNull(types.Int64Type)
	}

	// Read-only attributes.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")

	// Singleton resource: static ID.
	data.Id = types.StringValue("subscriberparam-config")
}
