package auditmessageaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditmessageactionDataSourceModel is the data-source-specific model, decoupled
// from AuditmessageactionResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (loglevel1, hits, undefhits, referencecount). Every non-key attribute is
// Computed; the Framework's per-attribute model <-> schema reflection requires
// this model to have exactly the attributes the data-source schema declares,
// which is why it cannot reuse the resource model.
type AuditmessageactionDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Bypasssafetycheck types.String `tfsdk:"bypasssafetycheck"`
	Loglevel          types.String `tfsdk:"loglevel"`
	Logtonewnslog     types.String `tfsdk:"logtonewnslog"`
	Name              types.String `tfsdk:"name"` // Required lookup key
	Stringbuilderexpr types.String `tfsdk:"stringbuilderexpr"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/auditmessageaction.json). Never settable; populated from GET.
	Loglevel1      types.String `tfsdk:"loglevel1"`
	Hits           types.Int64  `tfsdk:"hits"`
	Undefhits      types.Int64  `tfsdk:"undefhits"`
	Referencecount types.Int64  `tfsdk:"referencecount"`
}

func AuditmessageactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bypasssafetycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check and allow unsafe expressions.",
			},
			"loglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the severity level of the log message being generated..\nThe following loglevels are valid:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"logtonewnslog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the message to the new nslog.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit message action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the message action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my message action\" or 'my message action').",
			},
			"stringbuilderexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default-syntax expression that defines the format and content of the log message.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"loglevel1": schema.StringAttribute{
				Computed:    true,
				Description: "Resolved audit log level as returned by the appliance. Possible values: ALL, EMERGENCY, ALERT, CRITICAL, ERROR, WARNING, NOTICE, INFORMATIONAL, DEBUG, NONE.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action has been taken.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the action resulted in UNDEF.",
			},
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of references to the action.",
			},
		},
	}
}

// auditmessageactionDataSourceSetAttrFromGet projects a NITRO auditmessageaction
// GET response onto the data-source model. Because a data source has no
// plan/apply reconciliation, attributes are simply filled from the GET (or left
// Null when the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func auditmessageactionDataSourceSetAttrFromGet(ctx context.Context, data *AuditmessageactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In auditmessageactionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Bypasssafetycheck = utils.MapGetString(g, "bypasssafetycheck")
	// NITRO GET returns the resolved log level in the read-only "loglevel1"
	// field; preserve the existing loglevel<-loglevel1 mapping for backward
	// compatibility.
	data.Loglevel = utils.MapGetString(g, "loglevel1")
	data.Logtonewnslog = utils.MapGetString(g, "logtonewnslog")
	data.Stringbuilderexpr = utils.MapGetString(g, "stringbuilderexpr")

	// Read-only metadata.
	data.Loglevel1 = utils.MapGetString(g, "loglevel1")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
}
