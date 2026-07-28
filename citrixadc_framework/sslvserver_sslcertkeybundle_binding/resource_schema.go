package sslvserver_sslcertkeybundle_binding

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

// SslvserverSslcertkeybundleBindingResourceModel describes the resource data model.
type SslvserverSslcertkeybundleBindingResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Certkeybundlename types.String `tfsdk:"certkeybundlename"`
	Snicertkeybundle  types.Bool   `tfsdk:"snicertkeybundle"`
	Vservername       types.String `tfsdk:"vservername"`
}

func (r *SslvserverSslcertkeybundleBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslvserver_sslcertkeybundle_binding resource.",
			},
			"certkeybundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Certkeybundle name bound to the vserver.",
			},
			"snicertkeybundle": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to bind certkeybundle which will be used in SNI processing.",
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

func sslvserver_sslcertkeybundle_bindingGetThePayloadFromthePlan(ctx context.Context, data *SslvserverSslcertkeybundleBindingResourceModel) ssl.Sslvserversslcertkeybundlebinding {
	tflog.Debug(ctx, "In sslvserver_sslcertkeybundle_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslvserver_sslcertkeybundle_binding := ssl.Sslvserversslcertkeybundlebinding{}
	if !data.Certkeybundlename.IsNull() && !data.Certkeybundlename.IsUnknown() {
		sslvserver_sslcertkeybundle_binding.Certkeybundlename = data.Certkeybundlename.ValueString()
	}
	if !data.Snicertkeybundle.IsNull() && !data.Snicertkeybundle.IsUnknown() {
		sslvserver_sslcertkeybundle_binding.Snicertkeybundle = data.Snicertkeybundle.ValueBool()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		sslvserver_sslcertkeybundle_binding.Vservername = data.Vservername.ValueString()
	}

	return sslvserver_sslcertkeybundle_binding
}

func sslvserver_sslcertkeybundle_bindingSetAttrFromGet(ctx context.Context, data *SslvserverSslcertkeybundleBindingResourceModel, getResponseData map[string]interface{}) *SslvserverSslcertkeybundleBindingResourceModel {
	tflog.Debug(ctx, "In sslvserver_sslcertkeybundle_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["certkeybundlename"]; ok && val != nil {
		data.Certkeybundlename = types.StringValue(val.(string))
	} else {
		data.Certkeybundlename = types.StringNull()
	}
	if val, ok := getResponseData["snicertkeybundle"]; ok && val != nil {
		data.Snicertkeybundle = types.BoolValue(val.(bool))
	} else {
		data.Snicertkeybundle = types.BoolNull()
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
