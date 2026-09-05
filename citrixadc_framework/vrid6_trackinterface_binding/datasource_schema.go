package vrid6_trackinterface_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Vrid6TrackinterfaceBindingDataSourceModel describes the datasource data model.
// In addition to the resource identity attributes it exposes the read-only
// GET-response field (flags). trackinterface bindings have no vlan field.
type Vrid6TrackinterfaceBindingDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	VridId     types.Int64  `tfsdk:"vrid_id"`
	Trackifnum types.String `tfsdk:"trackifnum"`
	Flags      types.Int64  `tfsdk:"flags"`
}

func Vrid6TrackinterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"vrid_id": schema.Int64Attribute{
				Required:    true,
				Description: "Integer value that uniquely identifies a VMAC6 address.",
			},
			"trackifnum": schema.StringAttribute{
				Required:    true,
				Description: "Interfaces which need to be tracked for this vrID.",
			},
			"flags": schema.Int64Attribute{
				Computed:    true,
				Description: "Flags.",
			},
		},
	}
}

// vrid6_trackinterface_bindingDataSourceSetAttrFromGet faithfully copies every field
// from the GET response (including the read-only flags) and composes the ID, since
// the datasource has no Create to seed those values. Note: trackinterface bindings
// have no vlan field.
func vrid6_trackinterface_bindingDataSourceSetAttrFromGet(ctx context.Context, data *Vrid6TrackinterfaceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vrid6_trackinterface_bindingDataSourceSetAttrFromGet Function")

	// vrid_id maps to the NITRO integer key "id"; never null this Required key —
	// retain the config-supplied value when the GET omits it.
	if val, ok := g["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.VridId = types.Int64Value(intVal)
		}
	}
	// trackifnum is the datasource lookup key; retain the config-supplied value when
	// the GET omits it instead of nulling it.
	if val, ok := g["trackifnum"]; ok && val != nil {
		data.Trackifnum = types.StringValue(utils.AnyToString(val))
	}
	// Read-only NITRO output field.
	data.Flags = utils.MapGetInt64(g, "flags")

	// Set ID for the datasource.
	// Composite key: id,trackifnum (key:UrlEncode(value) pairs)
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.VridId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("trackifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Trackifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
