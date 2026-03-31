package vpnglobal_intranetip_binding

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

// VpnglobalIntranetipBindingResourceModel describes the resource data model.
type VpnglobalIntranetipBindingResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Intranetip             types.String `tfsdk:"intranetip"`
	Netmask                types.String `tfsdk:"netmask"`
}

func (r *VpnglobalIntranetipBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnglobal_intranetip_binding resource.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applicable only to advance vpn session policy. An expression or other value specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"intranetip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range.",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The intranet ip address or range's netmask.",
			},
		},
	}
}

func vpnglobal_intranetip_bindingGetThePayloadFromtheConfig(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel) vpn.Vpnglobalintranetipbinding {
	tflog.Debug(ctx, "In vpnglobal_intranetip_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnglobal_intranetip_binding := vpn.Vpnglobalintranetipbinding{}
	if !data.Gotopriorityexpression.IsNull() {
		vpnglobal_intranetip_binding.Gotopriorityexpression = data.Gotopriorityexpression.ValueString()
	}
	if !data.Intranetip.IsNull() {
		vpnglobal_intranetip_binding.Intranetip = data.Intranetip.ValueString()
	}
	if !data.Netmask.IsNull() {
		vpnglobal_intranetip_binding.Netmask = data.Netmask.ValueString()
	}

	return vpnglobal_intranetip_binding
}

func vpnglobal_intranetip_bindingSetAttrFromGet(ctx context.Context, data *VpnglobalIntranetipBindingResourceModel, getResponseData map[string]interface{}) *VpnglobalIntranetipBindingResourceModel {
	tflog.Debug(ctx, "In vpnglobal_intranetip_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["gotopriorityexpression"]; ok && val != nil {
		data.Gotopriorityexpression = types.StringValue(val.(string))
	} else {
		data.Gotopriorityexpression = types.StringNull()
	}
	if val, ok := getResponseData["intranetip"]; ok && val != nil {
		data.Intranetip = types.StringValue(val.(string))
	} else {
		data.Intranetip = types.StringNull()
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	} else {
		data.Netmask = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("intranetip:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Intranetip.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
