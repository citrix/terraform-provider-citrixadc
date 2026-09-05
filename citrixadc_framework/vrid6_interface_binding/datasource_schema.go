package vrid6_interface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Vrid6InterfaceBindingDataSourceModel describes the datasource data model.
// In addition to the resource identity attributes it exposes the read-only
// GET-response fields (flags, vlan) which are not part of the write payload.
type Vrid6InterfaceBindingDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"` // Required lookup key (NITRO key "id")
	Ifnum  types.String `tfsdk:"ifnum"`   // Required lookup key
	Flags  types.Int64  `tfsdk:"flags"`
	Vlan   types.Int64  `tfsdk:"vlan"`
}

func Vrid6InterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces to bind to the VMAC6, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
			"vlan": schema.Int64Attribute{
				Computed:    true,
				Description: "The VLAN in which this VRID resides.",
			},
		},
	}
}

// vrid6_interface_bindingDataSourceSetAttrFromGet projects a NITRO
// vrid6_interface_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func vrid6_interface_bindingDataSourceSetAttrFromGet(ctx context.Context, data *Vrid6InterfaceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vrid6_interface_bindingDataSourceSetAttrFromGet Function")

	// vrid_id (NITRO key "id") and ifnum are the lookup keys. The firmware does
	// NOT reliably echo them for these rows, so only overwrite when present and
	// otherwise RETAIN the config-supplied values (needed for a correct
	// composite ID and ifnum output).
	if v, ok := g["id"]; ok && v != nil {
		if iv, err := utils.ConvertToInt64(v); err == nil {
			data.VridId = types.Int64Value(iv)
		}
	}
	if v, ok := g["ifnum"]; ok && v != nil {
		data.Ifnum = types.StringValue(utils.AnyToString(v))
	}

	// Read-only (GET-only) attributes.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Vlan = utils.MapGetInt64(g, "vlan")

	// Set the composite ID (key order: id,ifnum).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
