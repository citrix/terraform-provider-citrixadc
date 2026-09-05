package ip6tunnel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ip6tunnelDataSourceModel is the data-source-specific model, decoupled from
// Ip6tunnelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type Ip6tunnelDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Local      types.String `tfsdk:"local"`
	Name       types.String `tfsdk:"name"` // Required lookup key
	Ownergroup types.String `tfsdk:"ownergroup"`
	Remote     types.String `tfsdk:"remote"` // Required lookup filter

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/ip6tunnel.json). Never settable; populated from GET.
	Remoteip types.String `tfsdk:"remoteip"`
	Type     types.Int64  `tfsdk:"type"`
	Encapip  types.String `tfsdk:"encapip"`
}

func Ip6tunnelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"local": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An IPv6 address of the local Citrix ADC used to set up the tunnel.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the IPv6 Tunnel. Cannot be changed after the service group is created. Must begin with a number or letter, and can consist of letters, numbers, and the @ _ - . (period) : (colon) # and space ( ) characters.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for the tunnel.",
			},
			"remote": schema.StringAttribute{
				Required:    true,
				Description: "An IPv6 address of the remote Citrix ADC used to set up the tunnel.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed; null when the
			// appliance omits them.
			"remoteip": schema.StringAttribute{
				Computed:    true,
				Description: "The remote IP address or subnet of the tunnel.",
			},
			"type": schema.Int64Attribute{
				Computed:    true,
				Description: "The type of this tunnel.",
			},
			"encapip": schema.StringAttribute{
				Computed:    true,
				Description: "The effective local IP address of the tunnel. Used as the source of the encapsulated packets.",
			},
		},
	}
}

// ip6tunnelDataSourceSetAttrFromGet projects a NITRO ip6tunnel GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func ip6tunnelDataSourceSetAttrFromGet(ctx context.Context, data *Ip6tunnelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In ip6tunnelDataSourceSetAttrFromGet Function")

	// Named resource keyed on "name"; the ID is the plain name value.
	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Local = utils.MapGetString(g, "local")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")

	// NITRO echoes the configured "remote" write property back as the read-only
	// "remoteip" property, so map state's "remote" from "remoteip" first, falling
	// back to "remote".
	if v, ok := g["remoteip"]; ok && v != nil {
		data.Remote = types.StringValue(utils.AnyToString(v))
	} else {
		data.Remote = utils.MapGetString(g, "remote")
	}

	// Read-only metadata.
	data.Remoteip = utils.MapGetString(g, "remoteip")
	data.Type = utils.MapGetInt64(g, "type")
	data.Encapip = utils.MapGetString(g, "encapip")
}
