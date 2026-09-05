package channel_interface_binding

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ChannelInterfaceBindingDataSourceModel is the data-source-specific model,
// decoupled from ChannelInterfaceBindingResourceModel. A data source is a pure
// read surface, so it can expose the FULL GET projection: the lookup keys AND
// the read-only member-interface metadata that the resource deliberately omits.
type ChannelInterfaceBindingDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Channelid types.String `tfsdk:"channelid"` // Required lookup key
	Ifnum     types.List   `tfsdk:"ifnum"`     // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/channel_interface_binding.json).
	Slavespeed   types.Int64  `tfsdk:"slavespeed"`
	Slaveflowctl types.Int64  `tfsdk:"slaveflowctl"`
	Lamode       types.String `tfsdk:"lamode"`
	Slavestate   types.Int64  `tfsdk:"slavestate"`
	Lractiveintf types.Int64  `tfsdk:"lractiveintf"`
	Slaveduplex  types.Int64  `tfsdk:"slaveduplex"`
	Slavetime    types.Int64  `tfsdk:"slavetime"`
	Slavemedia   types.Int64  `tfsdk:"slavemedia"`
}

func ChannelInterfaceBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"channelid": schema.StringAttribute{
				Required:    true,
				Description: "ID of the LA channel or the cluster LA channel to which you want to bind interfaces. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or a cluster LA channel in CLA/x notation or  Link redundant channel in LR/x notation , where x can range from 1 to 4.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.\nFor an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).\nFor an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).\nwhere C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\nU is a unique integer for representing an interface in a particular port group.\nN is the ID of the node to which an interface belongs in a cluster configuration.\nUse spaces to separate multiple entries.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"slavespeed": schema.Int64Attribute{
				Computed:    true,
				Description: "Speed of the member interfaces.",
			},
			"slaveflowctl": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowcontrol of the member interfaces.",
			},
			"lamode": schema.StringAttribute{
				Computed:    true,
				Description: "The  mode(AUTO/MANNUAL) for the LA channel. Possible values = MANUAL, AUTO",
			},
			"slavestate": schema.Int64Attribute{
				Computed:    true,
				Description: "State of the member interfaces.",
			},
			"lractiveintf": schema.Int64Attribute{
				Computed:    true,
				Description: "LR set member interface state(active/inactive).",
			},
			"slaveduplex": schema.Int64Attribute{
				Computed:    true,
				Description: "Duplex of the member interfaces.",
			},
			"slavetime": schema.Int64Attribute{
				Computed:    true,
				Description: "UP time of the member interfaces.",
			},
			"slavemedia": schema.Int64Attribute{
				Computed:    true,
				Description: "Media type of the member interfaces.",
			},
		},
	}
}

// channelInterfaceDatasourceFirstIfnum returns the first configured/echoed
// interface, used to compose the single-interface composite ID for the
// datasource and to match the aggregate GET rows.
func channelInterfaceDatasourceFirstIfnum(ctx context.Context, data *ChannelInterfaceBindingDataSourceModel) string {
	if data.Ifnum.IsNull() || data.Ifnum.IsUnknown() {
		return ""
	}
	var list []string
	data.Ifnum.ElementsAs(ctx, &list, false)
	if len(list) == 0 {
		return ""
	}
	return list[0]
}

// channel_interface_bindingDataSourceSetAttrFromGet projects a NITRO
// channel_interface_binding GET response onto the data-source model.
func channel_interface_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ChannelInterfaceBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In channel_interface_bindingDataSourceSetAttrFromGet Function")

	data.Channelid = utils.MapGetString(g, "id")

	if val, ok := g["ifnum"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			stringList := utils.ToStringList(v)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Ifnum = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Ifnum = listValue
		default:
			data.Ifnum = types.ListNull(types.StringType)
		}
	} else {
		data.Ifnum = types.ListNull(types.StringType)
	}

	// Read-only attributes.
	data.Slavespeed = utils.MapGetInt64(g, "slavespeed")
	data.Slaveflowctl = utils.MapGetInt64(g, "slaveflowctl")
	data.Lamode = utils.MapGetString(g, "lamode")
	data.Slavestate = utils.MapGetInt64(g, "slavestate")
	data.Lractiveintf = utils.MapGetInt64(g, "lractiveintf")
	data.Slaveduplex = utils.MapGetInt64(g, "slaveduplex")
	data.Slavetime = utils.MapGetInt64(g, "slavetime")
	data.Slavemedia = utils.MapGetInt64(g, "slavemedia")

	// Set the composite id (id:<channel>,ifnum:<intf>) for the datasource.
	data.Id = types.StringValue(channel_interface_bindingComposeId(data.Channelid.ValueString(), channelInterfaceDatasourceFirstIfnum(ctx, data)))
}
