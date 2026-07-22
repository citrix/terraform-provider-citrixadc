package sslservicegroup_sslcacertbundle_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslservicegroupSslcacertbundleBindingResourceModel describes the resource data model.
type SslservicegroupSslcacertbundleBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Cacertbundlename types.String `tfsdk:"cacertbundlename"`
	Servicegroupname types.String `tfsdk:"servicegroupname"`
}

func (r *SslservicegroupSslcacertbundleBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslservicegroup_sslcacertbundle_binding resource.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CA certbundle name bound to the servicegroup.",
			},
			"servicegroupname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The name of the SSL service to which the SSL policy needs to be bound.",
			},
		},
	}
}

func sslservicegroup_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslservicegroupSslcacertbundleBindingResourceModel) ssl.Sslservicegroupsslcacertbundlebinding {
	tflog.Debug(ctx, "In sslservicegroup_sslcacertbundle_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslservicegroup_sslcacertbundle_binding := ssl.Sslservicegroupsslcacertbundlebinding{}
	if !data.Cacertbundlename.IsNull() && !data.Cacertbundlename.IsUnknown() {
		sslservicegroup_sslcacertbundle_binding.Cacertbundlename = data.Cacertbundlename.ValueString()
	}
	if !data.Servicegroupname.IsNull() && !data.Servicegroupname.IsUnknown() {
		sslservicegroup_sslcacertbundle_binding.Servicegroupname = data.Servicegroupname.ValueString()
	}

	return sslservicegroup_sslcacertbundle_binding
}

func sslservicegroup_sslcacertbundle_bindingSetAttrFromGet(ctx context.Context, data *SslservicegroupSslcacertbundleBindingResourceModel, getResponseData map[string]interface{}) *SslservicegroupSslcacertbundleBindingResourceModel {
	tflog.Debug(ctx, "In sslservicegroup_sslcacertbundle_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertbundlename"]; ok && val != nil {
		data.Cacertbundlename = types.StringValue(val.(string))
	} else {
		data.Cacertbundlename = types.StringNull()
	}
	if val, ok := getResponseData["servicegroupname"]; ok && val != nil {
		data.Servicegroupname = types.StringValue(val.(string))
	} else {
		data.Servicegroupname = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertbundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicegroupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicegroupname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
