package snmpmib

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SnmpmibDataSourceModel is the data-source-specific model, decoupled from
// SnmpmibResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type SnmpmibDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Contact   types.String `tfsdk:"contact"`
	Customid  types.String `tfsdk:"customid"`
	Location  types.String `tfsdk:"location"`
	Name      types.String `tfsdk:"name"`
	Ownernode types.Int64  `tfsdk:"ownernode"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/snmpmib.json). Never settable; populated from GET.
	Sysdesc     types.String `tfsdk:"sysdesc"`
	Sysuptime   types.Int64  `tfsdk:"sysuptime"`
	Sysservices types.Int64  `tfsdk:"sysservices"`
	Sysoid      types.String `tfsdk:"sysoid"`
}

func SnmpmibDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"contact": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the administrator for this Citrix ADC. Along with the name, you can include information on how to contact this person, such as a phone number or an email address. Can consist of 1 to 127 characters that include uppercase and  lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the information includes one or more spaces, enclose it in double or single quotation marks (for example, \"my contact\" or 'my contact').",
			},
			"customid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom identification number for the Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters. You should choose a custom identification that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the ID includes one or more spaces, enclose it in double or single quotation marks (for example, \"my ID\" or 'my ID').",
			},
			"location": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical location of the Citrix ADC. For example, you can specify building name, lab number, and rack number. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the location includes one or more spaces, enclose it in double or single quotation marks (for example, \"my location\" or 'my location').",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name for this Citrix ADC. Can consist of 1 to 127 characters that include uppercase and lowercase letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at sign (@), equals (=), colon (:), and underscore (_) characters.  You should choose a name that helps identify the Citrix ADC appliance.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose it in double or single quotation marks (for example, \"my name\" or 'my name').",
			},
			"ownernode": schema.Int64Attribute{
				Required:    true,
				Description: "ID of the cluster node for which we are setting the mib. This is a mandatory argument to set snmp mib on CLIP.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"sysdesc": schema.StringAttribute{
				Computed:    true,
				Description: "The description of the system.",
			},
			"sysuptime": schema.Int64Attribute{
				Computed:    true,
				Description: "The UP time of the system in 100th of a second.",
			},
			"sysservices": schema.Int64Attribute{
				Computed:    true,
				Description: "The services offered by the system.",
			},
			"sysoid": schema.StringAttribute{
				Computed:    true,
				Description: "The OID of the system's management system.",
			},
		},
	}
}

// snmpmibDataSourceSetAttrFromGet projects a NITRO snmpmib GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func snmpmibDataSourceSetAttrFromGet(ctx context.Context, data *SnmpmibDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In snmpmibDataSourceSetAttrFromGet Function")

	data.Contact = utils.MapGetString(g, "contact")
	data.Customid = utils.MapGetString(g, "customid")
	data.Location = utils.MapGetString(g, "location")
	data.Name = utils.MapGetString(g, "name")

	// ownernode is the Required lookup key supplied via config; only overwrite it
	// when the GET echoes it so a config value the appliance omits is preserved.
	if v, ok := g["ownernode"]; ok && v != nil {
		data.Ownernode = utils.MapGetInt64(g, "ownernode")
	}

	// Read-only (GET-only) metadata.
	data.Sysdesc = utils.MapGetString(g, "sysdesc")
	data.Sysuptime = utils.MapGetInt64(g, "sysuptime")
	data.Sysservices = utils.MapGetInt64(g, "sysservices")
	data.Sysoid = utils.MapGetString(g, "sysoid")

	// snmpmib is a singleton (unnamed) resource; the ownernode is its unique
	// addressing attribute, so use it as the stable state ID (matches the resource).
	data.Id = types.StringValue(fmt.Sprintf("%d", data.Ownernode.ValueInt64()))
}
