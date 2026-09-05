package policymap

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicymapDataSourceModel is the data-source-specific model, decoupled from
// PolicymapResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (targetname). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type PolicymapDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Mappolicyname types.String `tfsdk:"mappolicyname"` // Required lookup key
	Sd            types.String `tfsdk:"sd"`
	Su            types.String `tfsdk:"su"`
	Td            types.String `tfsdk:"td"`
	Tu            types.String `tfsdk:"tu"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policymap.json). Never settable; populated from GET.
	Targetname types.String `tfsdk:"targetname"`
}

func PolicymapDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"mappolicyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the map policy. Must begin with a letter, number, or the underscore (_) character and must consist only of letters, numbers, and the hash (#), period (.), colon (:), space ( ), at (@), equals (=), hyphen (-), and underscore (_) characters.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my map\" or 'my map').",
			},
			"sd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Publicly known source domain name. This is the domain name with which a client request arrives at a reverse proxy virtual server for cache redirection. If you specify a source domain, you must specify a target domain.",
			},
			"su": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source URL. Specify all or part of the source URL, in the following format: /[[prefix] [*]] [.suffix].",
			},
			"td": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Target domain name sent to the server. The source domain name is replaced with this domain name.",
			},
			"tu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Target URL. Specify the target URL in the following format: /[[prefix] [*]][.suffix].",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"targetname": schema.StringAttribute{
				Computed:    true,
				Description: "The expression string.",
			},
		},
	}
}

// policymapDataSourceSetAttrFromGet projects a NITRO policymap GET response onto
// the data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func policymapDataSourceSetAttrFromGet(ctx context.Context, data *PolicymapDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policymapDataSourceSetAttrFromGet Function")

	if v, ok := g["mappolicyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Mappolicyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Sd = utils.MapGetString(g, "sd")
	data.Su = utils.MapGetString(g, "su")
	data.Td = utils.MapGetString(g, "td")
	data.Tu = utils.MapGetString(g, "tu")

	// Read-only metadata.
	data.Targetname = utils.MapGetString(g, "targetname")
}
