package snmpgroup

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpgroupDataSourceModel is the data-source-specific model, decoupled from
// SnmpgroupResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type SnmpgroupDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`          // Required lookup key
	Readviewname  types.String `tfsdk:"readviewname"`  // Optional+Computed
	Securitylevel types.String `tfsdk:"securitylevel"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/snmpgroup.json). Never settable; populated from GET.
	Storagetype types.String `tfsdk:"storagetype"`
	Status      types.String `tfsdk:"status"`
}

func SnmpgroupDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SNMPv3 group. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the SNMPv3 group. \n            \nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"readviewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the configured SNMPv3 view that you want to bind to this SNMPv3 group. An SNMPv3 user bound to this group can access the subtrees that are bound to this SNMPv3 view as type INCLUDED, but cannot access the ones that are type EXCLUDED. If the Citrix ADC has multiple SNMPv3 view entries with the same name, all such entries are associated with the SNMPv3 group.",
			},
			"securitylevel": schema.StringAttribute{
				Required:    true,
				Description: "Security level required for communication between the Citrix ADC and the SNMPv3 users who belong to the group. Specify one of the following options:\nnoAuthNoPriv. Require neither authentication nor encryption.\nauthNoPriv. Require authentication but no encryption.\nauthPriv. Require authentication and encryption.\nNote: If you specify authentication, you must specify an encryption algorithm when you assign an SNMPv3 user to the group. If you also specify encryption, you must assign both an authentication and an encryption algorithm for each group member.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"storagetype": schema.StringAttribute{
				Computed:    true,
				Description: "The storage type for this group. Possible values = volatile, nonVolatile",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The status of this group. Possible values = active",
			},
		},
	}
}

// snmpgroupDataSourceSetAttrFromGet projects a NITRO snmpgroup GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func snmpgroupDataSourceSetAttrFromGet(ctx context.Context, data *SnmpgroupDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpgroupDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Readviewname = utils.MapGetString(g, "readviewname")
	data.Securitylevel = utils.MapGetString(g, "securitylevel")

	// Read-only (GET-only) metadata.
	data.Storagetype = utils.MapGetString(g, "storagetype")
	data.Status = utils.MapGetString(g, "status")
}
