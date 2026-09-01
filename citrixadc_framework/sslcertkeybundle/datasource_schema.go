package sslcertkeybundle

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcertkeybundleDataSourceModel is the data-source-specific model, decoupled
// from SslcertkeybundleResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits. Every non-key attribute is Computed; the Framework's per-attribute
// model <-> schema reflection requires this model to have exactly the attributes
// the data-source schema declares, which is why it cannot reuse the resource
// model.
type SslcertkeybundleDataSourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Bundlefile         types.String `tfsdk:"bundlefile"`
	Certkeybundlename  types.String `tfsdk:"certkeybundlename"` // Required lookup key
	Passplain          types.String `tfsdk:"passplain"`
	PassplainWo        types.String `tfsdk:"passplain_wo"`
	PassplainWoVersion types.Int64  `tfsdk:"passplain_wo_version"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcertkeybundle.json). Never settable; populated from
	// GET.
	Certkeybundledigest types.String `tfsdk:"certkeybundledigest"`
}

func SslcertkeybundleDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bundlefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the X509 certificate bundle file that is used to form the certificate-key bundle. The certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The certificate bundle file consists of list of certificates and one key in PEM format.",
			},
			"certkeybundlename": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the cerKeyBundle. The name will be used to bind/unbind certkey bundle to vip. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},
			"passplain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Pass phrase used to encrypt the private-key. Required when certificate bundle file contains encrypted private-key in PEM format.",
			},
			"passplain_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a passplain_wo update.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"certkeybundledigest": schema.StringAttribute{
				Computed:    true,
				Description: "Stores the digest of certificate and key bundle file.",
			},
		},
	}
}

// sslcertkeybundleDataSourceSetAttrFromGet projects a NITRO sslcertkeybundle GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func sslcertkeybundleDataSourceSetAttrFromGet(ctx context.Context, data *SslcertkeybundleDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcertkeybundleDataSourceSetAttrFromGet Function")

	if v, ok := g["certkeybundlename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Certkeybundlename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attribute as read-back output.
	data.Bundlefile = utils.MapGetString(g, "bundlefile")

	// passplain / passplain_wo(+version) are write-only secret inputs the GET
	// never returns -> Null.
	data.Passplain = types.StringNull()
	data.PassplainWo = types.StringNull()
	data.PassplainWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Certkeybundledigest = utils.MapGetString(g, "certkeybundledigest")
}
