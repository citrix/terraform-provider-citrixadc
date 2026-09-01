package snmpview

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpviewDataSourceModel is the data-source-specific model, decoupled from
// SnmpviewResourceModel. A data source is a pure read surface, so it exposes the
// read/write attributes (as Computed outputs) AND the read-only attributes the
// GET returns (storagetype, status) that the resource deliberately omits.
type SnmpviewDataSourceModel struct {
	Id      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Subtree types.String `tfsdk:"subtree"`
	Type    types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/snmpview.json). Never settable; populated from GET.
	Storagetype types.String `tfsdk:"storagetype"`
	Status      types.String `tfsdk:"status"`
}

func SnmpviewDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the SNMPv3 view. Can consist of 1 to 31 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a name that helps identify the SNMPv3 view.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my view\" or 'my view').",
			},
			"subtree": schema.StringAttribute{
				Required:    true,
				Description: "A particular branch (subtree) of the MIB tree that you want to associate with this SNMPv3 view. You must specify the subtree as an SNMP OID.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Include or exclude the subtree, specified by the subtree parameter, in or from this view. This setting can be useful when you have included a subtree, such as A, in an SNMPv3 view and you want to exclude a specific subtree of A, such as B, from the SNMPv3 view.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"storagetype": schema.StringAttribute{
				Computed:    true,
				Description: "The storage type for this view. Possible values: [ volatile, nonVolatile ].",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "The status of this view. Possible values: [ active ].",
			},
		},
	}
}

// snmpviewDataSourceSetAttrFromGet projects a NITRO snmpview GET response onto
// the data-source model. Attributes are filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func snmpviewDataSourceSetAttrFromGet(ctx context.Context, data *SnmpviewDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpviewDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}
	data.Subtree = utils.MapGetString(g, "subtree")
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Storagetype = utils.MapGetString(g, "storagetype")
	data.Status = utils.MapGetString(g, "status")
}
