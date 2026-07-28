package sslservice_sslcacertbundle_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslserviceSslcacertbundleBindingResourceModel describes the resource data model.
type SslserviceSslcacertbundleBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Cacertbundlename types.String `tfsdk:"cacertbundlename"`
	Servicename      types.String `tfsdk:"servicename"`
	Skipcacertbundle types.Bool   `tfsdk:"skipcacertbundle"`
}

func (r *SslserviceSslcacertbundleBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservice_sslcacertbundle_binding resource.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CA certbundle name bound to the service.",
			},
			"servicename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
			"skipcacertbundle": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The flag is used to indicate whether all CA_names in this particular CA certificate bundle needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
		},
	}
}

func sslservice_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslserviceSslcacertbundleBindingResourceModel) ssl.Sslservicesslcacertbundlebinding {
	tflog.Debug(ctx, "In sslservice_sslcacertbundle_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslservice_sslcacertbundle_binding := ssl.Sslservicesslcacertbundlebinding{}
	if !data.Cacertbundlename.IsNull() && !data.Cacertbundlename.IsUnknown() {
		sslservice_sslcacertbundle_binding.Cacertbundlename = data.Cacertbundlename.ValueString()
	}
	if !data.Servicename.IsNull() && !data.Servicename.IsUnknown() {
		sslservice_sslcacertbundle_binding.Servicename = data.Servicename.ValueString()
	}
	if !data.Skipcacertbundle.IsNull() && !data.Skipcacertbundle.IsUnknown() {
		sslservice_sslcacertbundle_binding.Skipcacertbundle = data.Skipcacertbundle.ValueBool()
	}

	return sslservice_sslcacertbundle_binding
}

func sslservice_sslcacertbundle_bindingSetAttrFromGet(ctx context.Context, data *SslserviceSslcacertbundleBindingResourceModel, getResponseData map[string]interface{}) *SslserviceSslcacertbundleBindingResourceModel {
	tflog.Debug(ctx, "In sslservice_sslcacertbundle_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertbundlename"]; ok && val != nil {
		data.Cacertbundlename = types.StringValue(val.(string))
	} else {
		data.Cacertbundlename = types.StringNull()
	}
	if val, ok := getResponseData["servicename"]; ok && val != nil {
		data.Servicename = types.StringValue(val.(string))
	} else {
		data.Servicename = types.StringNull()
	}
	if val, ok := getResponseData["skipcacertbundle"]; ok && val != nil {
		data.Skipcacertbundle = types.BoolValue(val.(bool))
	} else {
		data.Skipcacertbundle = types.BoolNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertbundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
