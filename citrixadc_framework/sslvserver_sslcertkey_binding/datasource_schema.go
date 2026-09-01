package sslvserver_sslcertkey_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslvserverSslcertkeyBindingDataSourceModel is the data-source-specific model. A
// data source is a pure read surface, so in addition to the read/write attributes
// (surfaced as Computed outputs) it exposes the read-only (GET-only) NITRO
// attributes that the resource intentionally omits.
type SslvserverSslcertkeyBindingDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Ca          types.Bool   `tfsdk:"ca"`
	Certkeyname types.String `tfsdk:"certkeyname"`
	Crlcheck    types.String `tfsdk:"crlcheck"`
	Ocspcheck   types.String `tfsdk:"ocspcheck"`
	Skipcaname  types.Bool   `tfsdk:"skipcaname"`
	Snicert     types.Bool   `tfsdk:"snicert"`
	Vservername types.String `tfsdk:"vservername"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslvserver_sslcertkey_binding.json). Never settable;
	// populated from GET, null when the appliance omits them.
	Cleartextport types.Int64 `tfsdk:"cleartextport"`
}

func SslvserverSslcertkeyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ca": schema.BoolAttribute{
				Required:    true,
				Description: "CA certificate.",
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "The name of the certificate key pair binding.",
			},
			"crlcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the CRL check parameter. (Mandatory/Optional)",
			},
			"ocspcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The state of the OCSP check parameter. (Mandatory/Optional)",
			},
			"skipcaname": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The flag is used to indicate whether this particular CA certificate's CA_Name needs to be sent to the SSL client while requesting for client certificate in a SSL handshake",
			},
			"snicert": schema.BoolAttribute{
				Required:    true,
				Description: "The name of the CertKey. Use this option to bind Certkey(s) which will be used in SNI processing.",
			},
			"vservername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL virtual server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"cleartextport": schema.Int64Attribute{
				Computed:    true,
				Description: "Port on which clear-text data is sent by the appliance to the server. Do not specify this parameter for SSL offloading with end-to-end encryption.",
			},
		},
	}
}

// sslvserver_sslcertkey_bindingDataSourceSetAttrFromGet projects a NITRO
// sslvserver_sslcertkey_binding GET response onto the data-source model and sets
// the composite ID. Attributes are filled from the GET (or left Null when the GET
// omits them) via the shared utils.MapGet* helpers.
func sslvserver_sslcertkey_bindingDataSourceSetAttrFromGet(ctx context.Context, data *SslvserverSslcertkeyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslvserver_sslcertkey_bindingDataSourceSetAttrFromGet Function")

	data.Ca = utils.MapGetBool(g, "ca")
	data.Certkeyname = utils.MapGetString(g, "certkeyname")
	data.Crlcheck = utils.MapGetString(g, "crlcheck")
	data.Ocspcheck = utils.MapGetString(g, "ocspcheck")
	data.Skipcaname = utils.MapGetBool(g, "skipcaname")
	data.Snicert = utils.MapGetBool(g, "snicert")
	data.Vservername = utils.MapGetString(g, "vservername")

	// Read-only (GET-only) attributes.
	data.Cleartextport = utils.MapGetInt64(g, "cleartextport")

	// Datasource has no Create — set the composite ID here (legacy
	// resource_id_mapping.json order).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("certkeyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("snicert:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Snicert.ValueBool()))))
	idParts = append(idParts, fmt.Sprintf("ca:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ca.ValueBool()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
