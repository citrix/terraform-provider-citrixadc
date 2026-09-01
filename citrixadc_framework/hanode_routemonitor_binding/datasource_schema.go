package hanode_routemonitor_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// HanodeRoutemonitorBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface, so it
// can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed.
type HanodeRoutemonitorBindingDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	HanodeId     types.Int64  `tfsdk:"hanode_id"`
	Netmask      types.String `tfsdk:"netmask"`
	Routemonitor types.String `tfsdk:"routemonitor"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/hanode_routemonitor_binding.json). Never settable;
	// populated from GET; null when the appliance omits them.
	Routemonitorstate types.String `tfsdk:"routemonitorstate"`
	Flags             types.Int64  `tfsdk:"flags"`
}

func HanodeRoutemonitorBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hanode_id": schema.Int64Attribute{
				Required:    true,
				Description: "Number that uniquely identifies the local node. The ID of the local node is always 0.",
			},
			"netmask": schema.StringAttribute{
				Required:    true,
				Description: "The netmask.",
			},
			"routemonitor": schema.StringAttribute{
				Required:    true,
				Description: "The IP address (IPv4 or IPv6).",
			},

			// Read-only (GET-only) attributes surfaced by the data source (these
			// are intentionally NOT modeled on the resource). All Computed.
			"routemonitorstate": schema.StringAttribute{
				Computed:    true,
				Description: "State for the route monitor. Possible values = UP, DOWN.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "The flags for this entry.",
			},
		},
	}
}

// hanode_routemonitor_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers. The composite ID is
// built exactly as the resource Create emits it.
func hanode_routemonitor_bindingDataSourceSetAttrFromGet(ctx context.Context, data *HanodeRoutemonitorBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In hanode_routemonitor_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs. NITRO returns the hanode id
	// under the "id" key.
	data.HanodeId = utils.MapGetInt64(g, "id")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Routemonitor = utils.MapGetString(g, "routemonitor")

	// Read-only (GET-only) attributes.
	data.Routemonitorstate = utils.MapGetString(g, "routemonitorstate")
	data.Flags = utils.MapGetInt64(g, "flags")

	// Composite ID: hanode_id,routemonitor (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("hanode_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.HanodeId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
