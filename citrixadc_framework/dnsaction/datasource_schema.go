package dnsaction

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsactionDataSourceModel is the data-source-specific model, decoupled from
// DnsactionResourceModel. A data source is a pure read surface, so it can expose
// the full GET projection: the read/write attributes (as Computed outputs) AND
// the read-only attributes the resource deliberately omits.
type DnsactionDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Actionname       types.String `tfsdk:"actionname"`
	Actiontype       types.String `tfsdk:"actiontype"`
	Dnsprofilename   types.String `tfsdk:"dnsprofilename"`
	Ipaddress        types.List   `tfsdk:"ipaddress"`
	Preferredloclist types.List   `tfsdk:"preferredloclist"`
	Ttl              types.Int64  `tfsdk:"ttl"`
	Viewname         types.String `tfsdk:"viewname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/dnsaction.json). Never settable; populated from GET.
	Drop        types.String `tfsdk:"drop"`
	Cachebypass types.String `tfsdk:"cachebypass"`
	Builtin     types.List   `tfsdk:"builtin"`
	Feature     types.String `tfsdk:"feature"`
}

func DnsactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"actionname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the dns action.",
			},
			"actiontype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of DNS action that is being configured.",
			},
			"dnsprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the DNS profile to be associated with the transaction for which the action is chosen",
			},
			"ipaddress": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "List of IP address to be returned in case of rewrite_response actiontype. They can be of IPV4 or IPV6 type.\n        In case of set command We will remove all the IP address previously present in the action and will add new once given in set dns action command.",
			},
			"preferredloclist": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The location list in priority order used for the given action.",
			},
			"ttl": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time to live, in seconds.",
			},
			"viewname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The view name that must be used for the given action.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"drop": schema.StringAttribute{
				Computed:    true,
				Description: "The dns packet must be dropped. Possible values = YES, NO.",
			},
			"cachebypass": schema.StringAttribute{
				Computed:    true,
				Description: "By pass dns cache for this. Possible values = YES, NO.",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine whether DNS action is default or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL.",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
		},
	}
}

// dnsactionDataSourceSetAttrFromGet projects a NITRO dnsaction GET response onto
// the data-source model. Attributes are simply filled from the GET (or left Null
// when the GET omits them) via the shared utils.MapGet* helpers.
func dnsactionDataSourceSetAttrFromGet(ctx context.Context, data *DnsactionDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsactionDataSourceSetAttrFromGet Function")

	if v, ok := g["actionname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Actionname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Actiontype = utils.MapGetString(g, "actiontype")
	data.Dnsprofilename = utils.MapGetString(g, "dnsprofilename")
	data.Ipaddress = utils.MapGetStringList(g, "ipaddress")
	data.Preferredloclist = utils.MapGetStringList(g, "preferredloclist")
	data.Ttl = utils.MapGetInt64(g, "ttl")
	data.Viewname = utils.MapGetString(g, "viewname")

	// Read-only attributes.
	data.Drop = utils.MapGetString(g, "drop")
	data.Cachebypass = utils.MapGetString(g, "cachebypass")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
}
