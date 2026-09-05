package nsextension

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsextensionDataSourceModel is the data-source-specific model, decoupled from
// NsextensionResourceModel. A data source is a pure read surface, so it exposes
// the read/write attributes (as Computed outputs) AND the read-only attributes
// the resource deliberately omits (type, functionhits, functionundefhits,
// functionhaltcount).
type NsextensionDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Comment        types.String `tfsdk:"comment"`
	Name           types.String `tfsdk:"name"`
	Overwrite      types.Bool   `tfsdk:"overwrite"`
	Src            types.String `tfsdk:"src"`
	Trace          types.String `tfsdk:"trace"`
	Tracefunctions types.String `tfsdk:"tracefunctions"`
	Tracevariables types.String `tfsdk:"tracevariables"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/nsextension.json). Never settable; populated from GET.
	Type              types.String `tfsdk:"type"`
	Functionhits      types.Int64  `tfsdk:"functionhits"`
	Functionundefhits types.Int64  `tfsdk:"functionundefhits"`
	Functionhaltcount types.Int64  `tfsdk:"functionhaltcount"`
}

func NsextensionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the extension object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the extension object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported extension.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
			"trace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables tracing to the NS log file of extension execution:\n   off   - turns off tracing (equivalent to unset ns extension <extension-name> -trace)\n   calls - traces extension function calls with arguments and function returns with the first return value\n   lines - traces the above plus line numbers for executed extension lines\n   all   - traces the above plus local variables changed by executed extension lines\nNote that the DEBUG log level must be enabled to see extension tracing.\nThis can be done by set audit syslogParams -loglevel ALL or -loglevel DEBUG.",
			},
			"tracefunctions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of extension functions to trace. By default, all extension functions are traced.",
			},
			"tracevariables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of variables (in traced extension functions) to trace. By default, all variables are traced.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "Type of the extension object.",
			},
			"functionhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of time function evaluates successfully.",
			},
			"functionundefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times error occured in evaluating extension function.",
			},
			"functionhaltcount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of time function evaluation is halted.",
			},
		},
	}
}

// nsextensionDataSourceSetAttrFromGet projects a NITRO nsextension GET response
// onto the data-source model. Attributes are filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func nsextensionDataSourceSetAttrFromGet(ctx context.Context, data *NsextensionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In nsextensionDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")
	data.Trace = utils.MapGetString(g, "trace")
	data.Tracefunctions = utils.MapGetString(g, "tracefunctions")
	data.Tracevariables = utils.MapGetString(g, "tracevariables")

	// Read-only attributes.
	data.Type = utils.MapGetString(g, "type")
	data.Functionhits = utils.MapGetInt64(g, "functionhits")
	data.Functionundefhits = utils.MapGetInt64(g, "functionundefhits")
	data.Functionhaltcount = utils.MapGetInt64(g, "functionhaltcount")
}
