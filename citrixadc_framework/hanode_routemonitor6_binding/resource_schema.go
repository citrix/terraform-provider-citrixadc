package hanode_routemonitor6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ha"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// HanodeRoutemonitor6BindingResourceModel describes the resource data model.
type HanodeRoutemonitor6BindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Id           types.Int64  `tfsdk:"id"`
	Netmask      types.String `tfsdk:"netmask"`
	Routemonitor types.String `tfsdk:"routemonitor"`
}

func (r *HanodeRoutemonitor6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the hanode_routemonitor6_binding resource.",
			},
			"id": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number that uniquely identifies the local node. The ID of the local node is always 0.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The netmask.",
			},
			"routemonitor": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The IP address (IPv4 or IPv6).",
			},
		},
	}
}

func hanode_routemonitor6_bindingGetThePayloadFromthePlan(ctx context.Context, data *HanodeRoutemonitor6BindingResourceModel) ha.Hanoderoutemonitor6binding {
	tflog.Debug(ctx, "In hanode_routemonitor6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	hanode_routemonitor6_binding := ha.Hanoderoutemonitor6binding{}
	if !data.Id.IsNull() && !data.Id.IsUnknown() {
		hanode_routemonitor6_binding.Id = utils.IntPtr(int(data.Id.ValueInt64()))
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		hanode_routemonitor6_binding.Netmask = data.Netmask.ValueString()
	}
	if !data.Routemonitor.IsNull() && !data.Routemonitor.IsUnknown() {
		hanode_routemonitor6_binding.Routemonitor = data.Routemonitor.ValueString()
	}

	return hanode_routemonitor6_binding
}

func hanode_routemonitor6_bindingSetAttrFromGet(ctx context.Context, data *HanodeRoutemonitor6BindingResourceModel, getResponseData map[string]interface{}) *HanodeRoutemonitor6BindingResourceModel {
	tflog.Debug(ctx, "In hanode_routemonitor6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["id"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Id = types.Int64Value(intVal)
		}
	} else {
		data.Id = types.Int64Null()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}
	if val, ok := getResponseData["routemonitor"]; ok && val != nil {
		data.Routemonitor = types.StringValue(val.(string))
	} else {
		data.Routemonitor = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("id:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Id.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("netmask:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Netmask.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("routemonitor:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Routemonitor.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
