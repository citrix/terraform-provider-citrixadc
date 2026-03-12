package policymap

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

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
		},
	}
}
