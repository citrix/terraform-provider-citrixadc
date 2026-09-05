package nspartition

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NspartitionDataSourceModel is the data-source-specific model, decoupled from
// NspartitionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (partitionid, partitiontype, pmacinternal). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type NspartitionDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Force         types.Bool   `tfsdk:"force"`
	Maxbandwidth  types.Int64  `tfsdk:"maxbandwidth"`
	Maxconn       types.Int64  `tfsdk:"maxconn"`
	Maxmemlimit   types.Int64  `tfsdk:"maxmemlimit"`
	Minbandwidth  types.Int64  `tfsdk:"minbandwidth"`
	Partitionmac  types.String `tfsdk:"partitionmac"`
	Partitionname types.String `tfsdk:"partitionname"` // Required lookup key
	Save          types.Bool   `tfsdk:"save"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/nspartition.json). Never settable; populated from GET.
	Partitionid   types.Int64  `tfsdk:"partitionid"`
	Partitiontype types.String `tfsdk:"partitiontype"`
	Pmacinternal  types.Bool   `tfsdk:"pmacinternal"`
}

func NspartitionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"force": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will not be saved",
			},
			"maxbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits.",
			},
			"maxconn": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent connections that can be open in the partition. A zero value indicates no limit on number of open connections.",
			},
			"maxmemlimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum memory, in megabytes, allocated to the partition.  A zero value indicates the memory is unlimited on the partition and it can consume up to the system limits.",
			},
			"minbandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum bandwidth, in Kbps, that the partition can consume. A zero value indicates the bandwidth is unrestricted on the partition and it can consume up to the system limits",
			},
			"partitionmac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Special MAC address for the partition which is used for communication over shared vlans in this partition. If not specified, the MAC address is auto-generated.",
			},
			"partitionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the Partition. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"save": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switches to new admin partition without prompt for saving configuration. Configuration will be saved",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"partitionid": schema.Int64Attribute{
				Computed:    true,
				Description: "Partition Id.",
			},
			"partitiontype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of the Partition. Possible values = Default Partition, Current Partition.",
			},
			"pmacinternal": schema.BoolAttribute{
				Computed:    true,
				Description: "Partition MAC is generated internally.",
			},
		},
	}
}

// nspartitionDataSourceSetAttrFromGet projects a NITRO nspartition GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func nspartitionDataSourceSetAttrFromGet(ctx context.Context, data *NspartitionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nspartitionDataSourceSetAttrFromGet Function")

	if v, ok := g["partitionname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Partitionname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Maxbandwidth = utils.MapGetInt64(g, "maxbandwidth")
	data.Maxconn = utils.MapGetInt64(g, "maxconn")
	data.Maxmemlimit = utils.MapGetInt64(g, "maxmemlimit")
	data.Minbandwidth = utils.MapGetInt64(g, "minbandwidth")
	data.Partitionmac = utils.MapGetString(g, "partitionmac")

	// force / save are action-only inputs the GET never returns -> Null.
	data.Force = types.BoolNull()
	data.Save = types.BoolNull()

	// Read-only metadata.
	data.Partitionid = utils.MapGetInt64(g, "partitionid")
	data.Partitiontype = utils.MapGetString(g, "partitiontype")
	data.Pmacinternal = utils.MapGetBool(g, "pmacinternal")
}
