package dbdbprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DbdbprofileDataSourceModel is the data-source-specific model, decoupled from
// DbdbprofileResourceModel. A data source is a pure read surface, so it can
// expose the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type DbdbprofileDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Conmultiplex           types.String `tfsdk:"conmultiplex"`
	Enablecachingconmuxoff types.String `tfsdk:"enablecachingconmuxoff"`
	Interpretquery         types.String `tfsdk:"interpretquery"`
	Kcdaccount             types.String `tfsdk:"kcdaccount"`
	Name                   types.String `tfsdk:"name"`
	Stickiness             types.String `tfsdk:"stickiness"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dbdbprofile.json). Never settable; populated from GET.
	Refcnt types.Int64 `tfsdk:"refcnt"`
}

func DbdbprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use the same server-side connection for multiple client-side requests. Default is enabled.",
			},
			"enablecachingconmuxoff": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable caching when connection multiplexing is OFF.",
			},
			"interpretquery": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If ENABLED, inspect the query and update the connection information, if required. If DISABLED, forward the query to the server.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the KCD account that is used for Windows authentication.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the database profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"stickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the queries are related to each other, forward to the same backend server.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"refcnt": schema.Int64Attribute{
				Computed:    true,
				Description: "Profile Reference Count.",
			},
		},
	}
}

// dbdbprofileDataSourceSetAttrFromGet projects a NITRO dbdbprofile GET response
// onto the data-source model. Attributes are simply filled from the GET (or left
// Null when the GET omits them) via the shared utils.MapGet* helpers.
func dbdbprofileDataSourceSetAttrFromGet(ctx context.Context, data *DbdbprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dbdbprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Conmultiplex = utils.MapGetString(g, "conmultiplex")
	data.Enablecachingconmuxoff = utils.MapGetString(g, "enablecachingconmuxoff")
	data.Interpretquery = utils.MapGetString(g, "interpretquery")
	data.Kcdaccount = utils.MapGetString(g, "kcdaccount")
	data.Stickiness = utils.MapGetString(g, "stickiness")

	// Read-only attributes.
	data.Refcnt = utils.MapGetInt64(g, "refcnt")
}
