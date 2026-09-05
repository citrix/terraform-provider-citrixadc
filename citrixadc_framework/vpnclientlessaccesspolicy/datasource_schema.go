package vpnclientlessaccesspolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnclientlessaccesspolicyDataSourceModel is the data-source-specific model,
// decoupled from VpnclientlessaccesspolicyResourceModel. A data source is a pure
// read surface, so it exposes the FULL GET projection: the read/write attributes
// (as Computed outputs) AND the read-only counters/metadata the resource
// deliberately omits (undefaction, hits, undefhits, description, isdefault,
// builtin, feature). Every non-key attribute is Computed.
type VpnclientlessaccesspolicyDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Profilename types.String `tfsdk:"profilename"`
	Rule        types.String `tfsdk:"rule"`

	// Read-only (GET-only) attributes from zion73x_readonly. Never settable.
	Undefaction types.String `tfsdk:"undefaction"`
	Hits        types.Int64  `tfsdk:"hits"`
	Undefhits   types.Int64  `tfsdk:"undefhits"`
	Description types.String `tfsdk:"description"`
	Isdefault   types.Bool   `tfsdk:"isdefault"`
	Builtin     types.List   `tfsdk:"builtin"`
	Feature     types.String `tfsdk:"feature"`
}

func VpnclientlessaccesspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the new clientless access policy.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the profile to invoke for the clientless access.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, specifying the traffic that matches the policy.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"undefaction": schema.StringAttribute{
				Computed:    true,
				Description: "The UNDEF action.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy evaluated to true.",
			},
			"undefhits": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of times the policy evaluation resulted in undefined processing.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the clientless access policy.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default vpnclientlessrwpolicy.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if vpn clientless rewrite policy is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// vpnclientlessaccesspolicyDataSourceSetAttrFromGet projects a NITRO
// vpnclientlessaccesspolicy GET response onto the data-source model. Every
// attribute is filled from the GET (or left Null when the GET omits it) via the
// shared utils.MapGet* helpers.
func vpnclientlessaccesspolicyDataSourceSetAttrFromGet(ctx context.Context, data *VpnclientlessaccesspolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In vpnclientlessaccesspolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	data.Profilename = utils.MapGetString(g, "profilename")
	data.Rule = utils.MapGetString(g, "rule")

	// Read-only attributes.
	data.Undefaction = utils.MapGetString(g, "undefaction")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Undefhits = utils.MapGetInt64(g, "undefhits")
	data.Description = utils.MapGetString(g, "description")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
