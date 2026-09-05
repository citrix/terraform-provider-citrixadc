package bridgegroup_nsip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BridgegroupNsip6BindingDataSourceModel is the data-source-specific model,
// decoupled from BridgegroupNsip6BindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the read/write
// attributes (as Computed outputs) AND the read-only attributes that the
// resource deliberately omits. Every non-key attribute is Computed.
type BridgegroupNsip6BindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	BridgegroupId types.Int64  `tfsdk:"bridgegroup_id"` // Required lookup key
	Ipaddress     types.String `tfsdk:"ipaddress"`      // Required lookup key

	// Read/write attributes surfaced here as Computed outputs.
	Netmask    types.String `tfsdk:"netmask"`
	Ownergroup types.String `tfsdk:"ownergroup"`
	Td         types.Int64  `tfsdk:"td"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/bridgegroup_nsip6_binding.json).
	Rnat types.Bool `tfsdk:"rnat"`
}

func BridgegroupNsip6BindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bridgegroup_id": schema.Int64Attribute{
				Required:    true,
				Description: "The integer that uniquely identifies the bridge group.",
			},
			"ipaddress": schema.StringAttribute{
				Required:    true,
				Description: "The IP address assigned to the  bridge group.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A subnet mask associated with the network address.",
			},
			"ownergroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The owner node group in a Cluster for this vlan.",
			},
			"td": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Integer value that uniquely identifies the traffic domain in which you want to configure the entity. If you do not specify an ID, the entity becomes part of the default traffic domain, which has an ID of 0.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"rnat": schema.BoolAttribute{
				Computed:    true,
				Description: "Temporary flag used for internal purpose.",
			},
		},
	}
}

// bridgegroup_nsip6_bindingComposeIdForDatasource builds the composite resource
// ID for the data source using the legacy attribute order (bridgegroup_id,
// ipaddress) in the new key:value form.
func bridgegroup_nsip6_bindingComposeIdForDatasource(data *BridgegroupNsip6BindingDataSourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("bridgegroup_id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.BridgegroupId.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("ipaddress:%s", utils.UrlEncode(data.Ipaddress.ValueString())))
	return strings.Join(idParts, ",")
}

// bridgegroup_nsip6_bindingDataSourceSetAttrFromGet projects a NITRO
// bridgegroup_nsip6_binding GET response onto the data-source model.
func bridgegroup_nsip6_bindingDataSourceSetAttrFromGet(ctx context.Context, data *BridgegroupNsip6BindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In bridgegroup_nsip6_bindingDataSourceSetAttrFromGet Function")

	if val, ok := g["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.BridgegroupId = types.Int64Value(intVal)
		}
	} else {
		data.BridgegroupId = types.Int64Null()
	}
	data.Ipaddress = utils.MapGetString(g, "ipaddress")
	data.Netmask = utils.MapGetString(g, "netmask")
	data.Ownergroup = utils.MapGetString(g, "ownergroup")
	// td is a config-supplied key; NITRO omits it for the default traffic
	// domain (0), so preserve the configured value instead of nulling it.
	if tdv, tdok := g["td"]; tdok && tdv != nil {
		if iv, err := utils.ConvertToInt64(tdv); err == nil {
			data.Td = types.Int64Value(iv)
		}
	}

	// Read-only attributes.
	data.Rnat = utils.MapGetBool(g, "rnat")

	// Set the composite id for the datasource.
	data.Id = types.StringValue(bridgegroup_nsip6_bindingComposeIdForDatasource(data))
}
