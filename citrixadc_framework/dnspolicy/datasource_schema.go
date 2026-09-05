package dnspolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnspolicyDataSourceModel is the data-source-specific model, decoupled from
// DnspolicyResourceModel. A data source is a pure read surface, so it exposes
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits (hits, undefhits,
// description, builtin, feature). Every non-key attribute is Computed.
type DnspolicyDataSourceModel struct {
	Id                types.String `tfsdk:"id"`
	Actionname        types.String `tfsdk:"actionname"`
	Cachebypass       types.String `tfsdk:"cachebypass"`
	Drop              types.String `tfsdk:"drop"`
	Logaction         types.String `tfsdk:"logaction"`
	Name              types.String `tfsdk:"name"` // Required lookup key
	Preferredlocation types.String `tfsdk:"preferredlocation"`
	Preferredloclist  types.List   `tfsdk:"preferredloclist"`
	Rule              types.String `tfsdk:"rule"`
	Viewname          types.String `tfsdk:"viewname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnspolicy.json). Never settable; populated from GET.
	Hits        types.Int64  `tfsdk:"hits"`
	Undefhits   types.Int64  `tfsdk:"undefhits"`
	Description types.String `tfsdk:"description"`
	Builtin     types.List   `tfsdk:"builtin"`
	Feature     types.String `tfsdk:"feature"`
}

func DnspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS action to perform when the rule evaluates to TRUE. The built in actions function as follows:\n* dns_default_act_Drop. Drop the DNS request.\n* dns_default_act_Cachebypass. Bypass the DNS cache and forward the request to the name server.\nYou can create custom actions by using the add dns action command in the CLI or the DNS > Actions > Create DNS Action dialog box in the Citrix ADC configuration utility.",
			},
			"cachebypass": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By pass dns cache for this.",
			},
			"drop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The dns packet must be dropped.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the messagelog action to use for requests that match this policy.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the DNS policy.",
			},
			"preferredlocation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The location used for the given policy. This is deprecated attribute. Please use -prefLocList",
			},
			"preferredloclist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The location list in priority order used for the given policy.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression against which DNS traffic is evaluated.\nNote:\n* On the command line interface, if the expression includes blank spaces, the entire expression must be enclosed in double quotation marks.\n* If the expression itself includes double quotation marks, you must escape the quotations by using the  character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.\nExample: CLIENT.UDP.DNS.DOMAIN.EQ(\"domainname\")",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The view name that must be used for the given policy.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy has been hit.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of Undef hits.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policy.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether DNS policy is default or not. Possible values: [ MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL ]",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// dnspolicyDataSourceSetAttrFromGet projects a NITRO dnspolicy GET response onto
// the data-source model. Attributes are filled from the GET (or left Null when
// the GET omits them) via the shared utils.MapGet* helpers.
func dnspolicyDataSourceSetAttrFromGet(ctx context.Context, data *DnspolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnspolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Actionname = utils.MapGetString(g, "actionname")
	data.Cachebypass = utils.MapGetString(g, "cachebypass")
	data.Drop = utils.MapGetString(g, "drop")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Preferredlocation = utils.MapGetString(g, "preferredlocation")
	data.Preferredloclist = utils.MapGetStringList(g, "preferredloclist")
	data.Rule = utils.MapGetString(g, "rule")
	data.Viewname = utils.MapGetString(g, "viewname")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Description = utils.MapGetString(g, "description")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
