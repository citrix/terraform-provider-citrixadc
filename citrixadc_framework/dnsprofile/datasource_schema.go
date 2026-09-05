package dnsprofile

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DnsprofileDataSourceModel is the data-source-specific model, decoupled from
// DnsprofileResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (referencecount). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type DnsprofileDataSourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Cacheecsresponses            types.String `tfsdk:"cacheecsresponses"`
	Cachenegativeresponses       types.String `tfsdk:"cachenegativeresponses"`
	Cacherecords                 types.String `tfsdk:"cacherecords"`
	Dnsanswerseclogging          types.String `tfsdk:"dnsanswerseclogging"`
	Dnserrorlogging              types.String `tfsdk:"dnserrorlogging"`
	Dnsextendedlogging           types.String `tfsdk:"dnsextendedlogging"`
	Dnsprofilename               types.String `tfsdk:"dnsprofilename"` // Required lookup key
	Dnsquerylogging              types.String `tfsdk:"dnsquerylogging"`
	Dropmultiqueryrequest        types.String `tfsdk:"dropmultiqueryrequest"`
	Insertecs                    types.String `tfsdk:"insertecs"`
	Maxcacheableecsprefixlength  types.Int64  `tfsdk:"maxcacheableecsprefixlength"`
	Maxcacheableecsprefixlength6 types.Int64  `tfsdk:"maxcacheableecsprefixlength6"`
	Recursiveresolution          types.String `tfsdk:"recursiveresolution"`
	Replaceecs                   types.String `tfsdk:"replaceecs"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/dnsprofile.json). Never settable; populated from GET.
	Referencecount types.Int64 `tfsdk:"referencecount"`
}

func DnsprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"cacheecsresponses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache DNS responses with EDNS Client Subnet(ECS) option in the DNS cache. When disabled, the appliance stops caching responses with ECS option. This is relevant to proxy configuration. Enabling/disabling support of ECS option when Citrix ADC is authoritative for a GSLB domain is supported using a knob in GSLB vserver. In all other modes, ECS option is ignored.",
			},
			"cachenegativeresponses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache negative responses in the DNS cache. When disabled, the appliance stops caching negative responses except referral records. This applies to all configurations - proxy, end resolver, and forwarder. However, cached responses are not flushed. The appliance does not serve negative responses from the cache until this parameter is enabled again.",
			},
			"cacherecords": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cache resource records in the DNS cache. Applies to resource records obtained through proxy configurations only. End resolver and forwarder configurations always cache records in the DNS cache, and you cannot disable this behavior. When you disable record caching, the appliance stops caching server responses. However, cached records are not flushed. The appliance does not serve requests from the cache until record caching is enabled again.",
			},
			"dnsanswerseclogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS answer section; if enabled, answer section in the response will be logged.",
			},
			"dnserrorlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS error logging; if enabled, whenever error is encountered in DNS module reason for the error will be logged.",
			},
			"dnsextendedlogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS extended logging; if enabled, authority and additional section in the response will be logged.",
			},
			"dnsprofilename": schema.StringAttribute{
				Required:    true,
				Description: "Name of the DNS profile",
			},
			"dnsquerylogging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS query logging; if enabled, DNS query information such as DNS query id, DNS query flags , DNS domain name and DNS query type will be logged",
			},
			"dropmultiqueryrequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Drop the DNS requests containing multiple queries. When enabled, DNS requests containing multiple queries will be dropped. In case of proxy configuration by default the DNS request containing multiple queries is forwarded to the backend and in case of ADNS and Resolver configuration NOCODE error response will be sent to the client.",
			},
			"insertecs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert ECS Option on DNS query",
			},
			"maxcacheableecsprefixlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum ecs prefix length that will be cached",
			},
			"maxcacheableecsprefixlength6": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum ecs prefix length that will be cached for IPv6 subnets",
			},
			"recursiveresolution": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DNS recursive resolution; if enabled, will do recursive resolution for DNS query when the profile is associated with ADNS service, CS Vserver and DNS action",
			},
			"replaceecs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Replace ECS Option on DNS query",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"referencecount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of entities using this profile.",
			},
		},
	}
}

// dnsprofileDataSourceSetAttrFromGet projects a NITRO dnsprofile GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func dnsprofileDataSourceSetAttrFromGet(ctx context.Context, data *DnsprofileDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In dnsprofileDataSourceSetAttrFromGet Function")

	if v, ok := g["dnsprofilename"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Dnsprofilename = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Cacheecsresponses = utils.MapGetString(g, "cacheecsresponses")
	data.Cachenegativeresponses = utils.MapGetString(g, "cachenegativeresponses")
	data.Cacherecords = utils.MapGetString(g, "cacherecords")
	data.Dnsanswerseclogging = utils.MapGetString(g, "dnsanswerseclogging")
	data.Dnserrorlogging = utils.MapGetString(g, "dnserrorlogging")
	data.Dnsextendedlogging = utils.MapGetString(g, "dnsextendedlogging")
	data.Dnsquerylogging = utils.MapGetString(g, "dnsquerylogging")
	data.Dropmultiqueryrequest = utils.MapGetString(g, "dropmultiqueryrequest")
	data.Insertecs = utils.MapGetString(g, "insertecs")
	data.Maxcacheableecsprefixlength = utils.MapGetInt64(g, "maxcacheableecsprefixlength")
	data.Maxcacheableecsprefixlength6 = utils.MapGetInt64(g, "maxcacheableecsprefixlength6")
	data.Recursiveresolution = utils.MapGetString(g, "recursiveresolution")
	data.Replaceecs = utils.MapGetString(g, "replaceecs")

	// Read-only metadata.
	data.Referencecount = utils.MapGetInt64(g, "referencecount")
}
