package policyurlset

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// PolicyurlsetDataSourceModel is the data-source-specific model, decoupled from
// PolicyurlsetResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits
// (patterncount). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type PolicyurlsetDataSourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"` // Required lookup key
	Canaryurl           types.String `tfsdk:"canaryurl"`
	Comment             types.String `tfsdk:"comment"`
	Delimiter           types.String `tfsdk:"delimiter"`
	Imported            types.Bool   `tfsdk:"imported"`
	Interval            types.Int64  `tfsdk:"interval"`
	Matchedid           types.Int64  `tfsdk:"matchedid"`
	Overwrite           types.Bool   `tfsdk:"overwrite"`
	Privateset          types.Bool   `tfsdk:"privateset"`
	Rowseparator        types.String `tfsdk:"rowseparator"`
	Subdomainexactmatch types.Bool   `tfsdk:"subdomainexactmatch"`
	Url                 types.String `tfsdk:"url"`
	UrlWo               types.String `tfsdk:"url_wo"`
	UrlWoVersion        types.Int64  `tfsdk:"url_wo_version"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/policyurlset.json). Never settable; populated from GET.
	Patterncount types.Int64 `tfsdk:"patterncount"`
}

func PolicyurlsetDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"canaryurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add this URL to this urlset. Used for testing when contents of urlset is kept confidential.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this url set.",
			},
			"delimiter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CSV file record delimiter.",
			},
			"imported": schema.BoolAttribute{
				Computed:    true,
				Description: "when set, display shows all imported urlsets.",
			},
			"interval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The interval, in seconds, rounded down to the nearest 15 minutes, at which the update of urlset occurs.",
			},
			"matchedid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "An ID that would be sent to AppFlow to indicate which URLSet was the last one that matched the requested URL.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name of the url set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters. Must not be the name of an existing named expression, pattern set, dataset, string map, or HTTP callout.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file.",
			},
			"privateset": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prevent this urlset from being exported.",
			},
			"rowseparator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CSV file row separator.",
			},
			"subdomainexactmatch": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Force exact subdomain matching, ex. given an entry 'google.com' in the urlset, a request to 'news.google.com' won't match, if subdomainExactMatch is set.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path and file name) from where the CSV (comma separated file) file will be imported or exported. Each record/line will one entry within the urlset. The first field contains the URL pattern, subsequent fields contains the metadata, if available. HTTP, HTTPS and FTP protocols are supported. NOTE: The operation fails if the destination HTTPS server requires client certificate authentication for access.",
			},
			"url_wo": schema.StringAttribute{
				Optional:    true,
				Description: "URL (protocol, host, path and file name) from where the CSV (comma separated file) file will be imported or exported. Each record/line will one entry within the urlset. The first field contains the URL pattern, subsequent fields contains the metadata, if available. HTTP, HTTPS and FTP protocols are supported. NOTE: The operation fails if the destination HTTPS server requires client certificate authentication for access.",
			},
			"url_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Increment this version to signal a url_wo update.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"patterncount": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of patterns in this urlset.",
			},
		},
	}
}

// policyurlsetDataSourceSetAttrFromGet projects a NITRO policyurlset GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func policyurlsetDataSourceSetAttrFromGet(ctx context.Context, data *PolicyurlsetDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In policyurlsetDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Canaryurl = utils.MapGetString(g, "canaryurl")
	data.Comment = utils.MapGetString(g, "comment")
	data.Delimiter = utils.MapGetString(g, "delimiter")
	data.Imported = utils.MapGetBool(g, "imported")
	data.Interval = utils.MapGetInt64(g, "interval")
	data.Matchedid = utils.MapGetInt64(g, "matchedid")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Privateset = utils.MapGetBool(g, "privateset")
	data.Rowseparator = utils.MapGetString(g, "rowseparator")
	data.Subdomainexactmatch = utils.MapGetBool(g, "subdomainexactmatch")

	// url / url_wo / url_wo_version are write-only secret / action-only inputs the
	// GET never returns -> Null.
	data.Url = types.StringNull()
	data.UrlWo = types.StringNull()
	data.UrlWoVersion = types.Int64Null()

	// Read-only metadata.
	data.Patterncount = utils.MapGetInt64(g, "patterncount")
}
