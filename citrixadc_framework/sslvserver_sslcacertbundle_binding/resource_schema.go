package sslvserver_sslcacertbundle_binding

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslvserverSslcacertbundleBindingResourceModel describes the resource data model.
type SslvserverSslcacertbundleBindingResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Cacertbundlename types.String `tfsdk:"cacertbundlename"`
	Skipcacertbundle types.Bool   `tfsdk:"skipcacertbundle"`
	Vservername      types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslcacertbundleBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslcacertbundle_binding resource.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "CA certbundle name bound to the vserver.",
			},
			"skipcacertbundle": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"vservername": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SSL virtual server.",
			},
		},
	}
}

func sslvserver_sslcacertbundle_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslvserverSslcacertbundleBindingResourceModel) ssl.Sslvserversslcacertbundlebinding {
	tflog.Debug(ctx, "In sslvserver_sslcacertbundle_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslvserver_sslcacertbundle_binding := ssl.Sslvserversslcacertbundlebinding{}
	if !data.Cacertbundlename.IsNull() && !data.Cacertbundlename.IsUnknown() {
		sslvserver_sslcacertbundle_binding.Cacertbundlename = data.Cacertbundlename.ValueString()
	}
	if !data.Skipcacertbundle.IsNull() && !data.Skipcacertbundle.IsUnknown() {
		sslvserver_sslcacertbundle_binding.Skipcacertbundle = data.Skipcacertbundle.ValueBool()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		sslvserver_sslcacertbundle_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslcacertbundle_binding
}

func sslvserver_sslcacertbundle_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslcacertbundleBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcacertbundleBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcacertbundle_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["cacertbundlename"]; ok && val != nil {
		data.Cacertbundlename = types.StringValue(val.(string))
	} else {
		data.Cacertbundlename = types.StringNull()
	}
	if val, ok := getResponseData["skipcacertbundle"]; ok && val != nil {
		data.Skipcacertbundle = types.BoolValue(val.(bool))
	} else {
		data.Skipcacertbundle = types.BoolNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// ID is set once in Create (resource) / Read (datasource); do not recompute here
	// to avoid wiping it when a key field is absent from the GET response (Pattern 6).

	return data
}
