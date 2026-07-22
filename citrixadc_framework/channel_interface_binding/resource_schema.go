package channel_interface_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// ChannelInterfaceBindingResourceModel describes the resource data model.
//
// The NITRO parent key for this binding is literally named "id" (the LA channel
// name). Because Terraform reserves the "id" attribute for the synthetic resource
// identifier, the parent key is exposed here as "channelid" while the composite
// resource ID string still uses the "id" key (id:<channel>,ifnum:<intf>).
type ChannelInterfaceBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Channelid types.String `tfsdk:"channelid"`
	Ifnum     types.List   `tfsdk:"ifnum"`
}

func (r *ChannelInterfaceBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the channel_interface_binding resource.",
			},
			"channelid": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "ID of the LA channel or the cluster LA channel to which you want to bind interfaces. Specify an LA channel in LA/x notation, where x can range from 1 to 8 or a cluster LA channel in CLA/x notation or  Link redundant channel in LR/x notation , where x can range from 1 to 4.",
			},
			"ifnum": schema.ListAttribute{
				ElementType: types.StringType,
				Required:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Interfaces to be bound to the LA channel of a Citrix ADC or to the LA channel of a cluster configuration.\nFor an LA channel of a Citrix ADC, specify an interface in C/U notation (for example, 1/3).\nFor an LA channel of a cluster configuration, specify an interface in N/C/U notation (for example, 2/1/3).\nwhere C can take one of the following values:\n* 0 - Indicates a management interface.\n* 1 - Indicates a 1 Gbps port.\n* 10 - Indicates a 10 Gbps port.\nU is a unique integer for representing an interface in a particular port group.\nN is the ID of the node to which an interface belongs in a cluster configuration.\nUse spaces to separate multiple entries.",
			},
		},
	}
}

func channel_interface_bindingGetThePayloadFromthePlan(ctx context.Context, data *ChannelInterfaceBindingResourceModel) network.Channelinterfacebinding {
	tflog.Debug(ctx, "In channel_interface_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// NOTE: svmcmd is an internal/read-only "source of cmd" field and is NOT a CLI
	// bind argument (Pattern 15) - it is intentionally excluded from the payload.
	channel_interface_binding := network.Channelinterfacebinding{}
	if !data.Channelid.IsNull() && !data.Channelid.IsUnknown() {
		channel_interface_binding.Id = data.Channelid.ValueString()
	}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		var ifnumList []string
		data.Ifnum.ElementsAs(ctx, &ifnumList, false)
		channel_interface_binding.Ifnum = ifnumList
	}

	return channel_interface_binding
}

// channel_interface_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied values (channelid, ifnum are RequiresReplace
// identity attrs) and does NOT recompute the ID, which is set exactly once in Create.
func channel_interface_bindingSetAttrFromGet(ctx context.Context, data *ChannelInterfaceBindingResourceModel, getResponseData map[string]interface{}) *ChannelInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In channel_interface_bindingSetAttrFromGet Function")

	// Convert API response to model. The channelid (parent key) and ifnum are
	// RequiresReplace identity attrs supplied by the plan; preserve them as-is.
	if val, ok := getResponseData["id"]; ok && val != nil {
		if data.Channelid.IsNull() || data.Channelid.ValueString() == "" {
			data.Channelid = types.StringValue(val.(string))
		}
	}

	return data
}

// channel_interface_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to
// seed those values.
func channel_interface_bindingSetAttrFromGetForDatasource(ctx context.Context, data *ChannelInterfaceBindingResourceModel, getResponseData map[string]interface{}) *ChannelInterfaceBindingResourceModel {
	tflog.Debug(ctx, "In channel_interface_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["id"]; ok && val != nil {
		data.Channelid = types.StringValue(val.(string))
	} else {
		data.Channelid = types.StringNull()
	}
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
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

	// Set ID for the datasource. Composite key id:<channel>,ifnum:<intf>.
	data.Id = types.StringValue(channel_interface_bindingComposeId(data.Channelid.ValueString(), datasourceFirstIfnum(ctx, data)))

	return data
}

// datasourceFirstIfnum returns the first configured/echoed interface, used to
// compose the single-interface composite ID for the datasource.
func datasourceFirstIfnum(ctx context.Context, data *ChannelInterfaceBindingResourceModel) string {
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
