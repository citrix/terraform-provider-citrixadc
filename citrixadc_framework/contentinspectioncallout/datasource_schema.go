package contentinspectioncallout

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ContentinspectioncalloutDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this Content Inspection callout.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the Content Inspection callout. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or callout.",
			},
			"profilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Content Inspection profile. The type of the configured profile must match the type specified using -type argument.",
			},
			"resultexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression that extracts the callout results from the response sent by the CI callout agent. Must be a response based expression, that is, it must begin with ICAP.RES. The operations in this expression must match the return type. For example, if you configure a return type of TEXT, the result expression must be a text based expression, as in the following example: icap.res.header(\"ISTag\")",
			},
			"returntype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of data that the target callout agent returns in response to the callout.\nAvailable settings function as follows:\n* TEXT - Treat the returned value as a text string.\n* NUM - Treat the returned value as a number.\n* BOOL - Treat the returned value as a Boolean value.\nNote: You cannot change the return type after it is set.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of Content Inspection server. Mutually exclusive with the server name parameter.",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the load balancing or content switching virtual server or service to which the Content Inspection request is issued. Mutually exclusive with server IP address and port parameters. The service type must be TCP or SSL_TCP. If there are vservers and services with the same name, then vserver is selected.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Port of the Content Inspection server.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the Content Inspection callout. It must be one of the following:\n* ICAP - Sends ICAP request to the configured ICAP server.",
			},
		},
	}
}
