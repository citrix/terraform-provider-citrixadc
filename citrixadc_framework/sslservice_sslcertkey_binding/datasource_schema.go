package sslservice_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslserviceSslcertkeyBindingDataSourceModel is the data-source-specific model. A
// data source is a pure read surface, so in addition to the read/write attributes
// (surfaced as Computed outputs) it exposes the read-only (GET-only) NITRO
// attributes that the resource intentionally omits.
type SslserviceSslcertkeyBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ca          types.Bool   `tfsdk:"ca"`
	Certkeyname types.String `tfsdk:"certkeyname"`
	Crlcheck    types.String `tfsdk:"crlcheck"`
	Ocspcheck   types.String `tfsdk:"ocspcheck"`
	Servicename types.String `tfsdk:"servicename"`
	Skipcaname  types.Bool   `tfsdk:"skipcaname"`
	Snicert     types.Bool   `tfsdk:"snicert"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslservice_sslcertkey_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Cleartextport types.Int64 `tfsdk:"cleartextport"`
}

func SslserviceSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ca": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "The certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rule to use for the OCSP responder associated with the CA certificate during client authentication. If MANDATORY is specified, deny all SSL clients if the OCSP check fails because of connectivity issues with the remote OCSP server, or any other reason that prevents the OCSP check. With the OPTIONAL setting, allow SSL clients even if the OCSP check fails except when the client certificate is revoked.",
			},
			"servicename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL service for which to set advanced configuration.",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting      for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"cleartextport": schema.Int64Attribute{
				Computed:    true,
				Description: "The clearTextPort settings.",
			},
		},
	}
}

// sslservice_sslcertkey_bindingDataSourceSetAttrFromGet projects a NITRO
// sslservice_sslcertkey_binding GET response onto the data-source model and sets
// the composite ID. Attributes are filled from the GET (or left Null when the GET
// omits them) via the shared utils.MapGet* helpers.
func sslservice_sslcertkey_bindingDataSourceSetAttrFromGet(ctx context.Context, data *SslserviceSslcertkeyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslservice_sslcertkey_bindingDataSourceSetAttrFromGet Function")

	data.Ca = utils.MapGetBool(g, "ca")
	data.Certkeyname = utils.MapGetString(g, "certkeyname")
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Servicename = utils.MapGetString(g, "servicename")
	data.Skipcaname = utils.MapGetBool(g, "skipcaname")
	data.Snicert = utils.MapGetBool(g, "snicert")

	// Read-only (GET-only) attributes.
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")

	// Datasource has no Create — set the composite ID here.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("crlcheck:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Crlcheck.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("servicename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Servicename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
