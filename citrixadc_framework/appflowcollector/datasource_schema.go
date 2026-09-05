package appflowcollector

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppflowcollectorDataSourceModel is the data-source-specific model, decoupled
// from AppflowcollectorResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AppflowcollectorDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Ipaddress  types.String `tfsdk:"ipaddress"`
	Name       types.String `tfsdk:"name"` // Required lookup key
	Netprofile types.String `tfsdk:"netprofile"`
	Newname    types.String `tfsdk:"newname"`
	Port       types.Int64  `tfsdk:"port"`
	Transport  types.String `tfsdk:"transport"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appflowcollector.json). Never settable; populated from GET.
	State types.String `tfsdk:"state"`
}

func AppflowcollectorDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 address of the collector.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n Only four collectors can be configured.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow collector\" or 'my appflow collector').",
			},
			"netprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Netprofile to associate with the collector. The IP address defined in the profile is used as the source IP address for AppFlow traffic for this collector.  If you do not set this parameter, the Citrix ADC IP (NSIP) address is used as the source IP address.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the collector. Must begin with an ASCII alphabetic or underscore (_) character, and must\ncontain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow coll\" or 'my appflow coll').",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port on which the collector listens.",
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of collector: either logstream or ipfix or rest.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "Appflow collector state. Possible values: [ UP, DOWN, UNKNOWN, BUSY, OUT OF SERVICE, GOING OUT OF SERVICE, DOWN WHEN GOING OUT OF SERVICE, NS_EMPTY_STR, Unknown, DISABLED ].",
			},
		},
	}
}

// appflowcollectorDataSourceSetAttrFromGet projects a NITRO appflowcollector GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that projection.
func appflowcollectorDataSourceSetAttrFromGet(ctx context.Context, data *AppflowcollectorDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appflowcollectorDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Netprofile = utils.MapGetString(g, "netprofile")
	data.Port = utils.MapGetInt64(g, "port")
	data.Transport = utils.MapGetString(g, "transport")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only attributes.
	data.State = utils.MapGetString(g, "state")
}
