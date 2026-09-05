package bridgegroup_vlan_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BridgegroupVlanBindingDataSourceModel is the data-source-specific model,
// decoupled from BridgegroupVlanBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the lookup keys (as
// Computed/Required outputs) AND the read-only attributes that the resource
// deliberately omits.
type BridgegroupVlanBindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	BridgegroupId types.Int64  `tfsdk:"bridgegroup_id"` // Required lookup key
	Vlan          types.Int64  `tfsdk:"vlan"`           // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/bridgegroup_vlan_binding.json).
	Rnat types.Bool `tfsdk:"rnat"`
}

func BridgegroupVlanBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"vlan": schema.Int64Attribute{
				Required:    true,
				Description: "Names of all member VLANs.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"rnat": schema.BoolAttribute{
				Computed:    true,
				Description: "Temporary flag used for internal purpose.",
			},
		},
	}
}

// bridgegroup_vlan_bindingComposeIdForDatasource builds the composite resource
// ID for the data source using the legacy attribute order (bridgegroup_id,
// vlan) in the new key:value form.
func bridgegroup_vlan_bindingComposeIdForDatasource(data *BridgegroupVlanBindingDataSourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BridgegroupId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("vlan:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vlan.ValueInt64()))))
	return strings.Join(idParts, ",")
}

// bridgegroup_vlan_bindingDataSourceSetAttrFromGet projects a NITRO
// bridgegroup_vlan_binding GET response onto the data-source model.
func bridgegroup_vlan_bindingDataSourceSetAttrFromGet(ctx context.Context, data *BridgegroupVlanBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In bridgegroup_vlan_bindingDataSourceSetAttrFromGet Function")

	// The NITRO "id" field is the bridge group key.
	if val, ok := g["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
		}
	} else {
		data.BridgegroupId = types.Int64Null()
	}
	data.Vlan = utils.MapGetInt64(g, "vlan")

	// Read-only attributes.
	data.Rnat = utils.MapGetBool(g, "rnat")

	// Set the composite id for the datasource.
	data.Id = types.StringValue(bridgegroup_vlan_bindingComposeIdForDatasource(data))
}
