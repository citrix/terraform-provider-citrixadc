package sslcertkeybundle

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertkeybundleResourceModel describes the resource data model.
type SslcertkeybundleResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Bundlefile         types.String `tfsdk:"bundlefile"`
	Certkeybundlename  types.String `tfsdk:"certkeybundlename"`
	Passplain          types.String `tfsdk:"passplain"`
	PassplainWo        types.String `tfsdk:"passplain_wo"`
	PassplainWoVersion types.Int64  `tfsdk:"passplain_wo_version"`
}

func (r *SslcertkeybundleResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertkeybundle resource.",
			},
			"bundlefile": schema.StringAttribute{
				Required:    true,
				Description: "Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. The certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The certificate bundle file consists of list of certificates and one key in PEM format.",
			},
			"certkeybundlename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name given to the cerKeyBundle. The name will be used to bind/unbind certkey bundle to vip. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"passplain": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a passplain_wo update.",
			},
		},
	}
}

func sslcertkeybundleGetThePayloadFromthePlan(ctx context.Context, data *SslcertkeybundleResourceModel) ssl.Sslcertkeybundle {
	tflog.Debug(ctx, "In sslcertkeybundleGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertkeybundle := ssl.Sslcertkeybundle{}
	if !data.Bundlefile.IsNull() && !data.Bundlefile.IsUnknown() {
		sslcertkeybundle.Bundlefile = data.Bundlefile.ValueString()
	}
	if !data.Certkeybundlename.IsNull() && !data.Certkeybundlename.IsUnknown() {
		sslcertkeybundle.Certkeybundlename = data.Certkeybundlename.ValueString()
	}
	if !data.Passplain.IsNull() && !data.Passplain.IsUnknown() {
		sslcertkeybundle.Passplain = data.Passplain.ValueString()
	}
	// Skip write-only attribute: passplain_wo
	// Skip version tracker attribute: passplain_wo_version

	return sslcertkeybundle
}

func sslcertkeybundleGetThePayloadFromtheConfig(ctx context.Context, data *SslcertkeybundleResourceModel, payload *ssl.Sslcertkeybundle) {
	tflog.Debug(ctx, "In sslcertkeybundleGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: passplain_wo -> passplain
	if !data.PassplainWo.IsNull() {
		passplainWo := data.PassplainWo.ValueString()
		if passplainWo != "" {
			payload.Passplain = passplainWo
		}
	}
}

func sslcertkeybundleSetAttrFromGet(ctx context.Context, data *SslcertkeybundleResourceModel, getResponseData map[string]interface{}) *SslcertkeybundleResourceModel {
	tflog.Debug(ctx, "In sslcertkeybundleSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bundlefile"]; ok && val != nil {
		data.Bundlefile = types.StringValue(val.(string))
	} else {
		data.Bundlefile = types.StringNull()
	}
	if val, ok := getResponseData["certkeybundlename"]; ok && val != nil {
		data.Certkeybundlename = types.StringValue(val.(string))
	} else {
		data.Certkeybundlename = types.StringNull()
	}
	// passplain is not returned by NITRO API (secret/ephemeral) - retain from config
	// passplain_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// passplain_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// ID is set once in Create; do not recompute here (Pattern 6).

	return data
}

func sslcertkeybundleSetAttrFromGetForDatasource(ctx context.Context, data *SslcertkeybundleResourceModel, getResponseData map[string]interface{}) *SslcertkeybundleResourceModel {
	tflog.Debug(ctx, "In sslcertkeybundleSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["bundlefile"]; ok && val != nil {
		data.Bundlefile = types.StringValue(val.(string))
	} else {
		data.Bundlefile = types.StringNull()
	}
	if val, ok := getResponseData["certkeybundlename"]; ok && val != nil {
		data.Certkeybundlename = types.StringValue(val.(string))
	} else {
		data.Certkeybundlename = types.StringNull()
	}
	// passplain is a secret and is never returned by NITRO GET.

	// Datasource has no Create; set ID here.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Certkeybundlename.ValueString()))

	return data
}
