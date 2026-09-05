package ssldtlsprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SsldtlsprofileDataSourceModel is the data-source-specific model, decoupled
// from SsldtlsprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, feature). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type SsldtlsprofileDataSourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Helloverifyrequest   types.String `tfsdk:"helloverifyrequest"`
	Initialretrytimeout  types.Int64  `tfsdk:"initialretrytimeout"`
	Maxbadmacignorecount types.Int64  `tfsdk:"maxbadmacignorecount"`
	Maxholdqlen          types.Int64  `tfsdk:"maxholdqlen"`
	Maxpacketsize        types.Int64  `tfsdk:"maxpacketsize"`
	Maxrecordsize        types.Int64  `tfsdk:"maxrecordsize"`
	Maxretrytime         types.Int64  `tfsdk:"maxretrytime"`
	Name                 types.String `tfsdk:"name"` // Required lookup key
	Pmtudiscovery        types.String `tfsdk:"pmtudiscovery"`
	Terminatesession     types.String `tfsdk:"terminatesession"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/ssldtlsprofile.json). Never settable; populated from GET.
	Builtin types.List   `tfsdk:"builtin"`
	Feature types.String `tfsdk:"feature"`
}

func SsldtlsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"helloverifyrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send a Hello Verify request to validate the client.",
			},
			"initialretrytimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial time out value to retransmit the last flight sent from the NetScaler.",
			},
			"maxbadmacignorecount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of bad MAC errors to ignore for a connection prior disconnect. Disabling parameter terminateSession terminates session immediately when bad MAC is detected in the connection.",
			},
			"maxholdqlen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of datagrams that can be queued at DTLS layer for processing",
			},
			"maxpacketsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of packets to reassemble. This value helps protect against a fragmented packet attack.",
			},
			"maxrecordsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum size of records that can be sent if PMTU is disabled.",
			},
			"maxretrytime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Wait for the specified time, in seconds, before resending the request.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DTLS profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@),equals sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.",
			},
			"pmtudiscovery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source for the maximum record size value. If ENABLED, the value is taken from the PMTU table. If DISABLED, the value is taken from the profile.",
			},
			"terminatesession": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Terminate the session if the message authentication code (MAC) of the client and server do not match.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether dtls profile is built-in or not.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// ssldtlsprofileDataSourceSetAttrFromGet projects a NITRO ssldtlsprofile GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func ssldtlsprofileDataSourceSetAttrFromGet(ctx context.Context, data *SsldtlsprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In ssldtlsprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Helloverifyrequest = utils.MapGetString(g, "helloverifyrequest")
	data.Initialretrytimeout = utils.MapGetInt64(g, "initialretrytimeout")
	data.Maxbadmacignorecount = utils.MapGetInt64(g, "maxbadmacignorecount")
	data.Maxholdqlen = utils.MapGetInt64(g, "maxholdqlen")
	data.Maxpacketsize = utils.MapGetInt64(g, "maxpacketsize")
	data.Maxrecordsize = utils.MapGetInt64(g, "maxrecordsize")
	data.Maxretrytime = utils.MapGetInt64(g, "maxretrytime")
	data.Pmtudiscovery = utils.MapGetString(g, "pmtudiscovery")
	data.Terminatesession = utils.MapGetString(g, "terminatesession")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
