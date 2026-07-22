package sslcertificatechain

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertificatechainResourceModel describes the resource data model.
type SslcertificatechainResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Certkeyname types.String `tfsdk:"certkeyname"`
}

func (r *SslcertificatechainResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcertificatechain resource.",
			},
			"certkeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the certificate-key pair.",
			},
		},
	}
}

func sslcertificatechainGetThePayloadFromthePlan(ctx context.Context, data *SslcertificatechainResourceModel) ssl.Sslcertificatechain {
	tflog.Debug(ctx, "In sslcertificatechainGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslcertificatechain := ssl.Sslcertificatechain{}
	if !data.Certkeyname.IsNull() && !data.Certkeyname.IsUnknown() {
		sslcertificatechain.Certkeyname = data.Certkeyname.ValueString()
	}

	return sslcertificatechain
}

func sslcertificatechainSetAttrFromGet(ctx context.Context, data *SslcertificatechainResourceModel, getResponseData map[string]interface{}) *SslcertificatechainResourceModel {
	tflog.Debug(ctx, "In sslcertificatechainSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	} else {
		data.Certkeyname = types.StringNull()
	}

	// ID is set once in Create; do not recompute here (Pattern 6).

	return data
}

func sslcertificatechainSetAttrFromGetForDatasource(ctx context.Context, data *SslcertificatechainResourceModel, getResponseData map[string]interface{}) *SslcertificatechainResourceModel {
	tflog.Debug(ctx, "In sslcertificatechainSetAttrFromGetForDatasource Function")

	// Convert API response to model
	if val, ok := getResponseData["certkeyname"]; ok && val != nil {
		data.Certkeyname = types.StringValue(val.(string))
	} else {
		data.Certkeyname = types.StringNull()
	}

	// Datasource has no Create; set ID here.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Certkeyname.ValueString()))

	return data
}
