package sslservicegroup_ecccurve_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslservicegroupEcccurveBindingResourceModel describes the resource data model.
type SslservicegroupEcccurveBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Ecccurvename     types.String `tfsdk:"ecccurvename"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *SslservicegroupEcccurveBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup_ecccurve_binding resource.",
			},
			"ecccurvename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Named ECC curve bound to servicegroup.",
			},
			"servicegroupname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}

func sslservicegroup_ecccurve_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslservicegroupEcccurveBindingResourceModel) ssl.Sslservicegroupecccurvebinding {
	tflog.Debug(ctx, "In sslservicegroup_ecccurve_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservicegroup_ecccurve_binding := ssl.Sslservicegroupecccurvebinding{}
	if !data.Ecccurvename.IsNull() {
		sslservicegroup_ecccurve_binding.Ecccurvename = data.Ecccurvename.ValueString()
	}
	if !data.Servicegroupname.IsNull() {
		sslservicegroup_ecccurve_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return sslservicegroup_ecccurve_binding
}

func sslservicegroup_ecccurve_bindingSetAttrFromGet(ctx context.Context, data *SslservicegroupEcccurveBindingResourceModel, getResponseData map[string]interface{}) *SslservicegroupEcccurveBindingResourceModel {
	tflog.Debug(ctx, "In sslservicegroup_ecccurve_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ecccurvename"]; ok && val != nil {
		data.Ecccurvename = types.StringValue(val.(string))
	} else {
		data.Ecccurvename = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ecccurvename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ecccurvename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
