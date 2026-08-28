package aaaproxyparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// AaaproxyparamResourceModel describes the resource data model.
type AaaproxyparamResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Proxy                  types.String `tfsdk:"proxy"`
	Proxyauthorization     types.String `tfsdk:"proxyauthorization"`
	Proxyusername          types.String `tfsdk:"proxyusername"`
	Proxypassword          types.String `tfsdk:"proxypassword"`
	ProxypasswordWo        types.String `tfsdk:"proxypassword_wo"`
	ProxypasswordWoVersion types.Int64  `tfsdk:"proxypassword_wo_version"`
}

func (r *AaaproxyparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 2,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaaproxyparam resource.",
			},
			"proxy": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "IP address and Port of the proxy server to be used for HTTP access for this request. Configure in ipaddress:port format (a.b.c.d:e) or as a URL (http://a.b.c.d or http://a.b.c.d:8080).",
			},
			"proxyauthorization": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "This indicates whether Proxy-Authorization header will be sent or not. Possible values = disabled, basic",
			},
			"proxyusername": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Username that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password that will be sent as part of Basic Proxy-Authorization header.",
			},
			"proxypassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a proxypassword_wo update.",
			},
		},
	}
}

func aaaproxyparamGetThePayloadFromthePlan(ctx context.Context, data *AaaproxyparamResourceModel) aaa.Aaaproxyparam {
	tflog.Debug(ctx, "In aaaproxyparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	aaaproxyparam := aaa.Aaaproxyparam{}
	if !data.Proxy.IsNull() && !data.Proxy.IsUnknown() {
		aaaproxyparam.Proxy = data.Proxy.ValueString()
	}
	if !data.Proxyauthorization.IsNull() && !data.Proxyauthorization.IsUnknown() {
		aaaproxyparam.Proxyauthorization = data.Proxyauthorization.ValueString()
	}
	if !data.Proxyusername.IsNull() && !data.Proxyusername.IsUnknown() {
		aaaproxyparam.Proxyusername = data.Proxyusername.ValueString()
	}
	if !data.Proxypassword.IsNull() && !data.Proxypassword.IsUnknown() {
		aaaproxyparam.Proxypassword = data.Proxypassword.ValueString()
	}
	// Skip write-only attribute: proxypassword_wo
	// Skip version tracker attribute: proxypassword_wo_version

	return aaaproxyparam
}

func aaaproxyparamGetThePayloadFromtheConfig(ctx context.Context, data *AaaproxyparamResourceModel, payload *aaa.Aaaproxyparam) {
	tflog.Debug(ctx, "In aaaproxyparamGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: proxypassword_wo -> proxypassword
	if !data.ProxypasswordWo.IsNull() {
		proxypasswordWo := data.ProxypasswordWo.ValueString()
		if proxypasswordWo != "" {
			payload.Proxypassword = proxypasswordWo
		}
	}
}

func aaaproxyparamSetAttrFromGet(ctx context.Context, data *AaaproxyparamResourceModel, getResponseData map[string]interface{}) *AaaproxyparamResourceModel {
	tflog.Debug(ctx, "In aaaproxyparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["proxy"]; ok && val != nil {
		data.Proxy = types.StringValue(val.(string))
	} else {
		data.Proxy = types.StringNull()
	}
	if val, ok := getResponseData["proxyauthorization"]; ok && val != nil {
		data.Proxyauthorization = types.StringValue(val.(string))
	} else {
		data.Proxyauthorization = types.StringNull()
	}
	if val, ok := getResponseData["proxyusername"]; ok && val != nil {
		data.Proxyusername = types.StringValue(val.(string))
	} else {
		data.Proxyusername = types.StringNull()
	}
	// proxypassword is not returned by NITRO API in usable form (secret/ephemeral) - retain from config
	// proxypassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// proxypassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 1: No unique attributes - static ID (singleton)
	data.Id = types.StringValue("aaaproxyparam-config")

	return data
}
