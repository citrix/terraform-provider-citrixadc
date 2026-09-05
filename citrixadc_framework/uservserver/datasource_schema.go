package uservserver

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// UservserverDataSourceModel is the data-source-specific model, decoupled from
// UservserverResourceModel. A data source is a pure read surface (Read only; no
// plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes the
// resource deliberately omits (curstate, state-change timers, ...). Every
// non-key attribute is Computed.
type UservserverDataSourceModel struct {
	Id   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Params       types.String `tfsdk:"params"`
	Comment      types.String `tfsdk:"comment"`
	Defaultlb    types.String `tfsdk:"defaultlb"`
	Ipaddress    types.String `tfsdk:"ipaddress"`
	Port         types.Int64  `tfsdk:"port"`
	State        types.String `tfsdk:"state"`
	Userprotocol types.String `tfsdk:"userprotocol"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/uservserver.json). Never settable; populated from GET.
	Curstate                  types.String `tfsdk:"curstate"`
	Value                     types.String `tfsdk:"value"`
	Statechangetimesec        types.String `tfsdk:"statechangetimesec"`
	Statechangetimemsec       types.Int64  `tfsdk:"statechangetimemsec"`
	Tickssincelaststatechange types.Int64  `tfsdk:"tickssincelaststatechange"`
	Nodefaultbindings         types.String `tfsdk:"nodefaultbindings"`
}

func UservserverDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"params": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the protocol.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments that you might want to associate with the virtual server.",
			},
			"defaultlb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the default Load Balancing virtual server used for load balancing of services. The protocol type of default Load Balancing virtual server should be a user type.",
			},
			"ipaddress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address to assign to the virtual server.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the virtual server. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vserver\" or 'my vserver').",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port number for the virtual server.",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial state of the user vserver.",
			},
			"userprotocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User protocol uesd by the service.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"curstate": schema.StringAttribute{
				Computed:    true,
				Description: "Current user vserver state.",
			},
			"value": schema.StringAttribute{
				Computed:    true,
				Description: "SSL status.",
			},
			"statechangetimesec": schema.StringAttribute{
				Computed:    true,
				Description: "Time when last state change happened. Seconds part.",
			},
			"statechangetimemsec": schema.Int64Attribute{
				Computed:    true,
				Description: "Time at which last state change happened. Milliseconds part.",
			},
			"tickssincelaststatechange": schema.Int64Attribute{
				Computed:    true,
				Description: "Time in 10 millisecond ticks since the last state change.",
			},
			"nodefaultbindings": schema.StringAttribute{
				Computed:    true,
				Description: "to determine if the configuration will have default ssl CIPHER and ECC curve bindings.",
			},
		},
	}
}

// uservserverDataSourceSetAttrFromGet projects a NITRO uservserver GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func uservserverDataSourceSetAttrFromGet(ctx context.Context, data *UservserverDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In uservserverDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	// "Params" is the NITRO json key (capitalised); fall back to "params".
	if v, ok := g["Params"]; ok && v != nil {
		data.Params = types.StringValue(utils.AnyToString(v))
	} else {
		data.Params = utils.MapGetString(g, "params")
	}
	data.Comment = utils.MapGetString(g, "comment")
	data.Defaultlb = utils.MapGetString(g, "defaultlb")
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Port = utils.MapGetInt64(g, "port")
	data.State = utils.MapGetString(g, "state")
	data.Userprotocol = utils.MapGetString(g, "userprotocol")

	// Read-only metadata.
	data.Curstate = utils.MapGetString(g, "curstate")
	data.Value = utils.MapGetString(g, "value")
	data.Statechangetimesec = utils.MapGetString(g, "statechangetimesec")
	data.Statechangetimemsec = utils.MapGetInt64(g, "statechangetimemsec")
	data.Tickssincelaststatechange = utils.MapGetInt64(g, "tickssincelaststatechange")
	data.Nodefaultbindings = utils.MapGetString(g, "nodefaultbindings")
}
