package gslbservice_dnsview_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/gslb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// GslbserviceDnsviewBindingResourceModel describes the resource data model.
type GslbserviceDnsviewBindingResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Servicename types.String `tfsdk:"servicename"`
	Viewip      types.String `tfsdk:"viewip"`
	Viewname    types.String `tfsdk:"viewname"`
}

func (r *GslbserviceDnsviewBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the gslbservice_dnsview_binding resource.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the GSLB service.",
			},
			"viewip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address to be used for the given view",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS view of the service. A DNS view is used in global server load balancing (GSLB) to return a predetermined IP address to a specific group of clients, which are identified by using a DNS policy.",
			},
		},
	}
}

func gslbservice_dnsview_bindingGetThePayloadFromtheConfig(ctx context.Context, data *GslbserviceDnsviewBindingResourceModel) gslb.Gslbservicednsviewbinding {
	tflog.Debug(ctx, "In gslbservice_dnsview_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	gslbservice_dnsview_binding := gslb.Gslbservicednsviewbinding{}
	if !data.Servicename.IsNull() {
		gslbservice_dnsview_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Viewip.IsNull() {
		gslbservice_dnsview_binding.Viewip = data.Viewip.ValueString()
	}
	if !data.Viewname.IsNull() {
		gslbservice_dnsview_binding.Viewname = data.Viewname.ValueString()
	}

	return gslbservice_dnsview_binding
}

func gslbservice_dnsview_bindingSetAttrFromGet(ctx context.Context, data *GslbserviceDnsviewBindingResourceModel, getResponseData map[string]interface{}) *GslbserviceDnsviewBindingResourceModel {
	tflog.Debug(ctx, "In gslbservice_dnsview_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["viewip"]; ok && val != nil {
		data.Viewip = types.StringValue(val.(string))
	} else {
		data.Viewip = types.StringNull()
	}
	if val, ok := getResponseData["viewname"]; ok && val != nil {
		data.Viewname = types.StringValue(val.(string))
	} else {
		data.Viewname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("viewname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Viewname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
