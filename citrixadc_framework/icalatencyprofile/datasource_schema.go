package icalatencyprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcalatencyprofileDataSourceModel is the data-source-specific model, decoupled
// from IcalatencyprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (refcnt, builtin, feature, isdefault). Every non-key attribute is Computed;
// the Framework's per-attribute model <-> schema reflection requires this model
// to have exactly the attributes the data-source schema declares, which is why
// it cannot reuse the resource model.
type IcalatencyprofileDataSourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"` // Required lookup key
	L7latencymaxnotifycount  types.Int64  `tfsdk:"l7latencymaxnotifycount"`
	L7latencymonitoring      types.String `tfsdk:"l7latencymonitoring"`
	L7latencynotifyinterval  types.Int64  `tfsdk:"l7latencynotifyinterval"`
	L7latencythresholdfactor types.Int64  `tfsdk:"l7latencythresholdfactor"`
	L7latencywaittime        types.Int64  `tfsdk:"l7latencywaittime"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/icalatencyprofile.json). Never settable; populated from GET.
	Refcnt    types.Int64  `tfsdk:"refcnt"`
	Builtin   types.List   `tfsdk:"builtin"`
	Feature   types.String `tfsdk:"feature"`
	Isdefault types.Bool   `tfsdk:"isdefault"`
}

func IcalatencyprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"l7latencymaxnotifycount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Max notify Count. This is the upper limit on the number of notifications sent to the Insight Center within an interval where the Latency is above the threshold.",
			},
			"l7latencymonitoring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/Disable L7 Latency monitoring for L7 latency notifications",
			},
			"l7latencynotifyinterval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Notify Interval. This is the interval at which the Citrix ADC sends out notifications to the Insight Center after the wait time has passed.",
			},
			"l7latencythresholdfactor": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency threshold factor. This is the factor by which the active latency should be greater than the minimum observed value to determine that the latency is high and may need to be reported",
			},
			"l7latencywaittime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "L7 Latency Wait time. This is the time for which the Citrix ADC waits after the threshold is exceeded before it sends out a Notification to the Insight Center.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the ICA latencyprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA latency profile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica l7latencyprofile\" or 'my ica l7latencyprofile').",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this l7latencyprofile.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that the ICA latencyprofile is a built-in (SYSTEM INTERNAL) type. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default l7latencyprofile.",
			},
		},
	}
}

// icalatencyprofileDataSourceSetAttrFromGet projects a NITRO icalatencyprofile
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func icalatencyprofileDataSourceSetAttrFromGet(ctx context.Context, data *IcalatencyprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In icalatencyprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.L7latencymaxnotifycount = utils.MapGetInt64(g, "l7latencymaxnotifycount")
	data.L7latencymonitoring = utils.MapGetString(g, "l7latencymonitoring")
	data.L7latencynotifyinterval = utils.MapGetInt64(g, "l7latencynotifyinterval")
	data.L7latencythresholdfactor = utils.MapGetInt64(g, "l7latencythresholdfactor")
	data.L7latencywaittime = utils.MapGetInt64(g, "l7latencywaittime")

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
