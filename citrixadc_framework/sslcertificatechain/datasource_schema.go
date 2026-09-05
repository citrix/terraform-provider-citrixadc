package sslcertificatechain

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertificatechainDataSourceModel is the data-source-specific model,
// decoupled from SslcertificatechainResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the lookup key (as a Computed output on
// id) AND the read-only chain metadata attributes that the resource deliberately
// omits. Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type SslcertificatechainDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Certkeyname types.String `tfsdk:"certkeyname"` // Required lookup key

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcertificatechain.json). Never settable; populated
	// from GET.
	Chainlinked        types.List   `tfsdk:"chainlinked"`
	Chainpossiblelinks types.List   `tfsdk:"chainpossiblelinks"`
	Chainissuer        types.String `tfsdk:"chainissuer"`
	Chaincomplete      types.Int64  `tfsdk:"chaincomplete"`
}

func SslcertificatechainDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"certkeyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the certificate-key pair.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"chainlinked": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Certkeys which are currently in SSL certificate chain.",
			},
			"chainpossiblelinks": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Certkeys which can be in SSL certificate chain.",
			},
			"chainissuer": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the issuer.",
			},
			"chaincomplete": schema.Int64Attribute{
				Computed:    true,
				Description: "Is set to 1 if ssl certificate chain is complete.",
			},
		},
	}
}

// sslcertificatechainDataSourceSetAttrFromGet projects a NITRO
// sslcertificatechain GET response onto the data-source model. Because a data
// source has no plan/apply reconciliation, attributes are simply filled from the
// GET (or left Null when the GET omits them). The shared utils.MapGet* helpers
// implement that projection.
func sslcertificatechainDataSourceSetAttrFromGet(ctx context.Context, data *SslcertificatechainDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcertificatechainDataSourceSetAttrFromGet Function")

	if v, ok := g["certkeyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Certkeyname = types.StringValue(utils.AnyToString(v))
	}

	// Read-only chain metadata.
	data.Chainlinked = utils.MapGetStringList(g, "chainlinked")
	data.Chainpossiblelinks = utils.MapGetStringList(g, "chainpossiblelinks")
	data.Chainissuer = utils.MapGetString(g, "chainissuer")
	data.Chaincomplete = utils.MapGetInt64(g, "chaincomplete")
}
