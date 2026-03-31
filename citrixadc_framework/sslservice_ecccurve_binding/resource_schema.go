package sslservice_ecccurve_binding

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

// SslserviceEcccurveBindingResourceModel describes the resource data model.
type SslserviceEcccurveBindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Ecccurvename types.String `tfsdk:"ecccurvename"`
	Servicename  types.String `tfsdk:"servicename"`
}

func (r *SslserviceEcccurveBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice_ecccurve_binding resource.",
			},
			"ecccurvename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Named ECC curve bound to service/vserver.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
		},
	}
}

func sslservice_ecccurve_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslserviceEcccurveBindingResourceModel) ssl.Sslserviceecccurvebinding {
	tflog.Debug(ctx, "In sslservice_ecccurve_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslservice_ecccurve_binding := ssl.Sslserviceecccurvebinding{}
	if !data.Ecccurvename.IsNull() {
		sslservice_ecccurve_binding.Ecccurvename = data.Ecccurvename.ValueString()
	}
	if !data.Servicename.IsNull() {
		sslservice_ecccurve_binding.Servicename = data.Servicename.ValueString()
	}

	return sslservice_ecccurve_binding
}

func sslservice_ecccurve_bindingSetAttrFromGet(ctx context.Context, data *SslserviceEcccurveBindingResourceModel, getResponseData map[string]interface{}) *SslserviceEcccurveBindingResourceModel {
	tflog.Debug(ctx, "In sslservice_ecccurve_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ecccurvename"]; ok && val != nil {
		data.Ecccurvename = types.StringValue(val.(string))
	} else {
		data.Ecccurvename = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ecccurvename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ecccurvename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
