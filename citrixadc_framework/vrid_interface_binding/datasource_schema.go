package vrid_interface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VridInterfaceBindingDataSourceModel describes the DATASOURCE data model. It adds
// the read-only output fields (flags, vlan) returned by the GET endpoint.
type VridInterfaceBindingDataSourceModel struct {
	Id     types.String `tfsdk:"id"`
	VridId types.Int64  `tfsdk:"vrid_id"`
	Ifnum  types.String `tfsdk:"ifnum"`
	Flags  types.Int64  `tfsdk:"flags"`
	Vlan   types.Int64  `tfsdk:"vlan"`
}

func VridInterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer that uniquely identifies the VMAC address. The generic VMAC address is in the form of 00:00:5e:00:01:<VRID>. For example, if you add a VRID with a value of 60 and bind it to an interface, the resulting VMAC address is 00:00:5e:00:01:3c, where 3c is the hexadecimal representation of 60.",
			},
			"ifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces to bind to the VMAC, specified in (slot/port) notation (for example, 1/2).Use spaces to separate multiple entries.",
			},
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

// vrid_interface_bindingDataSourceSetAttrFromGet faithfully copies every field from
// the GET response (including the read-only flags and vlan) and composes the ID,
// since the datasource has no Create to seed those values.
func vrid_interface_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VridInterfaceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vrid_interface_bindingDataSourceSetAttrFromGet Function")

	// vrid_id maps to the NITRO integer key "id"; never null this Required key —
	// retain the config-supplied value when the GET omits it.
	if val, ok := g["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	// NITRO may echo ifnum as a scalar string or (on some firmware) a list. When the
	// firmware does NOT echo "ifnum" for vrid_interface_binding rows we deliberately
	// RETAIN the config-supplied ifnum (the datasource's lookup key) instead of
	// nulling it, so the composite ID and the ifnum output remain correct.
	if val, ok := g["ifnum"]; ok && val != nil {
		switch t := val.(type) {
		case string:
			data.Ifnum = types.StringValue(t)
		case []interface{}:
			if len(t) > 0 {
				if s, ok := t[0].(string); ok {
					data.Ifnum = types.StringValue(s)
				}
			}
		}
	}
	// Read-only NITRO output fields.
	data.Flags = utils.MapGetInt64(g, "flags")
	data.Vlan = utils.MapGetInt64(g, "vlan")

	// Set ID for the datasource. Composite key: vrid_id (NITRO id), ifnum.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
