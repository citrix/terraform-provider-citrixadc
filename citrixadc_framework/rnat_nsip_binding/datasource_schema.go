package rnat_nsip_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RnatNsipBindingDataSourceModel is the data-source-specific model, decoupled
// from the resource model. A data source is a pure read surface (Read only; no
// plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits. Every non-key attribute is Computed.
type RnatNsipBindingDataSourceModel struct {
	Id    types.String `tfsdk:"id"`
	Name  types.String `tfsdk:"name"`
	Natip types.String `tfsdk:"natip"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/rnat_nsip_binding.json). Never settable; populated
	// from GET, Null when the appliance omits them.
	Ownergroup types.String `tfsdk:"ownergroup"`
	Td         types.Int64  `tfsdk:"td"`
}

func RnatNsipBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the RNAT rule to which to bind NAT IPs.",
			},
			"natip": schema.StringAttribute{
				Required:    true,
				Description: "Any NetScaler-owned IPv4 address except the NSIP address. The NetScaler appliance replaces the source IP addresses of server-generated packets with the IP address specified. The IP address must be a public NetScaler-owned IP address. If you specify multiple addresses for this field, NATIP selection uses the round robin algorithm for each session. By specifying a range of IP addresses, you can specify all NetScaler-owned IP addresses, except the NSIP, that fall within the specified range.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"ownergroup": schema.StringAttribute{
				Computed:    true,
				Description: "The owner node group in a Cluster for this rnat rule.",
			},
			"td": schema.Int64Attribute{
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which the entity is configured.",
			},
		},
	}
}

// rnat_nsip_bindingDataSourceSetAttrFromGet projects a NITRO rnat_nsip_binding
// GET response onto the data-source model.
func rnat_nsip_bindingDataSourceSetAttrFromGet(ctx context.Context, data *RnatNsipBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In rnat_nsip_bindingDataSourceSetAttrFromGet Function")

	data.Name = utils.MapGetString(g, "name")
	data.Natip = utils.MapGetString(g, "natip")

	// Read-only (GET-only) attributes.
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	data.Td = utils.MapGetInt64(g, "td")

	// Set composite ID. Backward-compatible with SDK v2: identity is
	// "name,natip" (comma-separated key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("natip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Natip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
