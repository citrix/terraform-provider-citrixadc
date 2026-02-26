package sslvserver_ecccurve_binding

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

// SslvserverEcccurveBindingResourceModel describes the resource data model.
type SslvserverEcccurveBindingResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Ecccurvename types.String `tfsdk:"ecccurvename"`
	Vservername  types.String `tfsdk:"vservername"`
}

func (r *SslvserverEcccurveBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_ecccurve_binding resource.",
			},
			"ecccurvename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Named ECC curve bound to vserver/service.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_ecccurve_bindingGetThePayloadFromtheConfig(ctx context.Context, data *SslvserverEcccurveBindingResourceModel) ssl.Sslvserverecccurvebinding {
	tflog.Debug(ctx, "In sslvserver_ecccurve_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslvserver_ecccurve_binding := ssl.Sslvserverecccurvebinding{}
	if !data.Ecccurvename.IsNull() {
		sslvserver_ecccurve_binding.Ecccurvename = data.Ecccurvename.ValueString()
	}
	if !data.Vservername.IsNull() {
		sslvserver_ecccurve_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_ecccurve_binding
}

func sslvserver_ecccurve_bindingSetAttrFromGet(ctx context.Context, data *SslvserverEcccurveBindingResourceModel, getResponseData map[string]interface{}) *SslvserverEcccurveBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_ecccurve_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ecccurvename"]; ok && val != nil {
		data.Ecccurvename = types.StringValue(val.(string))
	} else {
		data.Ecccurvename = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ecccurvename:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Ecccurvename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
