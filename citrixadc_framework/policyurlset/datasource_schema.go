package policyurlset

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
