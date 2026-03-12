package vpnglobal_intranetip6_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetip6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range.",
			},
			"numaddr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}

func vpnglobal_intranetip6_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalIntranetip6BindingResourceModel) vpn.Vpnglobalintranetip6binding {
	tflog.Debug(ctx, "In vpnglobal_intranetip6_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_intranetip6_binding := vpn.Vpnglobalintranetip6binding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_intranetip6_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetip6.IsNull() {
		vpnglobal_intranetip6_binding.Intranetip6 = data.Intranetip6.ValueString()
	}
	if !data.Numaddr.IsNull() {
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
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip6:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Intranetip6.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
