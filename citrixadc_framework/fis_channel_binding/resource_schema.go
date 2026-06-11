package fis_channel_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FisChannelBindingResourceModel describes the resource data model.
type FisChannelBindingResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Ifnum     types.String `tfsdk:"ifnum"`
	Name      types.String `tfsdk:"name"`
	Ownernode types.Int64  `tfsdk:"ownernode"`
}

func (r *FisChannelBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the fis_channel_binding resource.",
			},
			"ifnum": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Interface to be bound to the FIS, specified in slot/port notation (for example, 1/3)",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the FIS to which you want to bind interfaces.",
			},
			"ownernode": schema.Int64Attribute{
				// Cluster/show-only attribute; not a bind arg. Optional (no Computed).
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "ID of the cluster node for which you are creating the FIS. Can be configured only through the cluster IP address.",
			},
		},
	}
}

func fis_channel_bindingGetThePayloadFromthePlan(ctx context.Context, data *FisChannelBindingResourceModel) network.Fischannelbinding {
	tflog.Debug(ctx, "In fis_channel_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model.
	// ownernode is a cluster/show-only attribute, NOT a bind arg; it is excluded from the payload.
	fis_channel_binding := network.Fischannelbinding{}
	if !data.Ifnum.IsNull() && !data.Ifnum.IsUnknown() {
		fis_channel_binding.Ifnum = data.Ifnum.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		fis_channel_binding.Name = data.Name.ValueString()
	}

	return fis_channel_binding
}

// fis_channel_bindingSetAttrFromGet is used by the resource Read/Create flow.
// It preserves the plan/state-supplied identity values (name, ifnum are RequiresReplace)
// and does NOT recompute the ID, which is set exactly once in Create.
func fis_channel_bindingSetAttrFromGet(ctx context.Context, data *FisChannelBindingResourceModel, getResponseData map[string]interface{}) *FisChannelBindingResourceModel {
	tflog.Debug(ctx, "In fis_channel_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	}

	return data
}

// fis_channel_bindingSetAttrFromGetForDatasource faithfully copies every field
// from the GET response and composes the ID, since the datasource has no Create to seed those values.
func fis_channel_bindingSetAttrFromGetForDatasource(ctx context.Context, data *FisChannelBindingResourceModel, getResponseData map[string]interface{}) *FisChannelBindingResourceModel {
	tflog.Debug(ctx, "In fis_channel_bindingSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["ifnum"]; ok && val != nil {
		data.Ifnum = types.StringValue(val.(string))
	} else {
		data.Ifnum = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["ownernode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ownernode = types.Int64Value(intVal)
		}
	} else {
		data.Ownernode = types.Int64Null()
	}

	// Set ID for the datasource
	// Composite key: name,ifnum
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("ifnum:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ifnum.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
