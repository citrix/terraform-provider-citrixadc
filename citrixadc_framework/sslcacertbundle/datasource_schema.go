package sslcacertbundle

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslcacertbundleDataSourceModel is the data-source-specific model, decoupled
// from SslcacertbundleResourceModel. A data source is a pure read surface, so it
// exposes the read/write attributes (as Computed outputs) AND the read-only
// attributes the GET returns (servername, cacertbundledigest) that the resource
// deliberately omits.
type SslcacertbundleDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Bundlefile       types.String `tfsdk:"bundlefile"`
	Cacertbundlename types.String `tfsdk:"cacertbundlename"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/sslcacertbundle.json). Never settable; populated from GET.
	Servername         types.String `tfsdk:"servername"`
	Cacertbundledigest types.String `tfsdk:"cacertbundledigest"`
}

func SslcacertbundleDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bundlefile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the X509 CA certificate bundle file that is used to form cacertbundle entity. The CA certificate bundle file should be present on the appliance's hard-disk drive or solid-state drive. /nsconfig/ssl/ is the default path. The CA certificate bundle file consists of list of certificates.",
			},
			"cacertbundlename": schema.StringAttribute{
				Required:    true,
				Description: "Name given to the CA certbundle. The name will be used for bind/unbind/update operations. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. The following requirement applies only to the Citrix ADC CLI: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my file\" or 'my file').",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"servername": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver/Service/Servicegroup name to which the cacertbundle is bound.",
			},
			"cacertbundledigest": schema.StringAttribute{
				Computed:    true,
				Description: "Stores the digest of a CA certificate bundle file.",
			},
		},
	}
}

// sslcacertbundleDataSourceSetAttrFromGet projects a NITRO sslcacertbundle GET
// response onto the data-source model. Attributes are filled from the GET (or
// left Null when the GET omits them) via the shared utils.MapGet* helpers.
func sslcacertbundleDataSourceSetAttrFromGet(ctx context.Context, data *SslcacertbundleDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslcacertbundleDataSourceSetAttrFromGet Function")

	if v, ok := g["cacertbundlename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Cacertbundlename = types.StringValue(utils.AnyToString(v))
	}
	data.Bundlefile = utils.MapGetString(g, "bundlefile")

	// Read-only metadata.
	data.Servername = utils.MapGetString(g, "servername")
	data.Cacertbundledigest = utils.MapGetString(g, "cacertbundledigest")
}
