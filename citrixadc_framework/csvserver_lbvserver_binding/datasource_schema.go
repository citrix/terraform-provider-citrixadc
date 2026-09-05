package csvserver_lbvserver_binding

import (
	"context"
	"fmt"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverLbvserverBindingDataSourceModel is the data-source-specific model,
// decoupled from CsvserverLbvserverBindingResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares.
type CsvserverLbvserverBindingDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Lbvserver     types.String `tfsdk:"lbvserver"`
	Name          types.String `tfsdk:"name"`
	Targetvserver types.String `tfsdk:"targetvserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver_lbvserver_binding.json). Never settable;
	// populated from GET.
	Vserverid    types.String `tfsdk:"vserverid"`
	Hits         types.Int64  `tfsdk:"hits"`
	Cookieipport types.String `tfsdk:"cookieipport"`
}

func CsvserverLbvserverBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"lbvserver": schema.StringAttribute{
				Required:    true,
				Description: "Name of the default lb vserver bound. Use this param for Default binding only. For Example: bind cs vserver cs1 -lbvserver lb1",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The virtual server name (created with the add lb vserver command) to which content will be switched.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"vserverid": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver Id of vserver.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"cookieipport": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver id of the lb vserver that is inserted into the set-cookie HTTP header.",
			},
		},
	}
}

// csvserver_lbvserver_bindingDataSourceSetAttrFromGet projects a NITRO
// csvserver_lbvserver_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func csvserver_lbvserver_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverLbvserverBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserver_lbvserver_bindingDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Lbvserver = utils.MapGetString(g, "lbvserver")
	data.Name = utils.MapGetString(g, "name")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")

	// Read-only attributes.
	data.Vserverid = utils.MapGetString(g, "vserverid")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Cookieipport = utils.MapGetString(g, "cookieipport")

	// Composite binding ID: "<name>,<lbvserver>".
	data.Id = types.StringValue(fmt.Sprintf("%s,%s", data.Name.ValueString(), data.Lbvserver.ValueString()))
}
