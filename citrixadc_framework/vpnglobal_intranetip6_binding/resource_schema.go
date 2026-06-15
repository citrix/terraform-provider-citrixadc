package vpnglobal_intranetip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// VpnglobalIntranetip6BindingResourceModel describes the resource data model.
type VpnglobalIntranetip6BindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetip6            types.String `tfsdk:"intranetip6"`
	Numaddr                types.Int64  `tfsdk:"numaddr"`
}

func (r *VpnglobalIntranetip6BindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_intranetip6_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetip6": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The intranet ip address or range.",
			},
			"numaddr": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}

func vpnglobal_intranetip6_bindingGetThePayloadFromthePlan(ctx context.Context, data *VpnglobalIntranetip6BindingResourceModel) vpn.Vpnglobalintranetip6binding {
	tflog.Debug(ctx, "In vpnglobal_intranetip6_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnglobal_intranetip6_binding := vpn.Vpnglobalintranetip6binding{}
	if !data.Gotopriorityexpression.IsNull() && !data.Gotopriorityexpression.IsUnknown() {
		vpnglobal_intranetip6_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetip6.IsNull() && !data.Intranetip6.IsUnknown() {
		vpnglobal_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Numaddr.IsNull() && !data.Numaddr.IsUnknown() {
		vpnglobal_intranetip6_binding.Numaddr = utils.IntPtr(int(data.Numaddr.ValueInt64()))
	}

	return vpnglobal_intranetip6_binding
}

func vpnglobal_intranetip6_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalIntranetip6BindingResourceModel, getResponseData map[string]interface{}) *VpnglobalIntranetip6BindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_intranetip6_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetip6"]; ok && val != nil {
		data.Intranetip6 = types.StringValue(val.(string))
	} else {
		data.Intranetip6 = types.StringNull()
	}
	if val, ok := getResponseData["numaddr"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Numaddr = types.Int64Value(intVal)
		}
	} else {
		data.Numaddr = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip6.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("numaddr:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Numaddr.ValueInt64()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
