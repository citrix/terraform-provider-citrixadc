package sslhsmkey

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslhsmkeyDataSourceModel is the data-source-specific model, decoupled from
// SslhsmkeyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits.
// Every non-key attribute is Computed; the Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type SslhsmkeyDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Hsmkeyname types.String `tfsdk:"hsmkeyname"` // Required lookup key
	Hsmtype    types.String `tfsdk:"hsmtype"`
	Key        types.String `tfsdk:"key"`
	Keystore   types.String `tfsdk:"keystore"`
	Password   types.String `tfsdk:"password"`
	Serialnum  types.String `tfsdk:"serialnum"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/sslhsmkey.json). Never settable; populated from GET.
	State types.String `tfsdk:"state"`
}

func SslhsmkeyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"hsmkeyname": schema.StringAttribute{
				Required:    true,
				Description: "0",
			},
			"hsmtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of HSM.",
			},
			"key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.",
			},
			"keystore": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for a partition. Applies only to SafeNet HSM.",
			},
			"serialnum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Serial number of the partition on which the key is present. Applies only to SafeNet HSM.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "Current state of key. Possible values = Created, Access Token Unavailable, Unauthorized, Does not exist, Unreachable, Marked down, Key operations successful, Key operations failed, Key operation throttled.",
			},
		},
	}
}

// sslhsmkeyDataSourceSetAttrFromGet projects a NITRO sslhsmkey GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them) via the shared utils.MapGet* helpers.
func sslhsmkeyDataSourceSetAttrFromGet(ctx context.Context, data *SslhsmkeyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In sslhsmkeyDataSourceSetAttrFromGet Function")

	if v, ok := g["hsmkeyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Hsmkeyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Hsmtype = utils.MapGetString(g, "hsmtype")
	data.Key = utils.MapGetString(g, "key")
	data.Keystore = utils.MapGetString(g, "keystore")
	data.Serialnum = utils.MapGetString(g, "serialnum")

	// password is a secret input the GET never returns -> Null.
	data.Password = types.StringNull()

	// Read-only attributes.
	data.State = utils.MapGetString(g, "state")
}
