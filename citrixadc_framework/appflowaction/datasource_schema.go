package appflowaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func AppflowactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"botinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the bot insight records to the configured collectors.",
			},
			"ciinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the ContentInspection Insight records to the configured collectors.",
			},
			"clientsidemeasurements": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will collect the time required to load and render the mainpage on the client.",
			},
			"collectors": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Name(s) of collector(s) to be associated with the AppFlow action.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments about this action.  In the CLI, if including spaces between words, enclose the comment in quotation marks. (The quotation marks are not required in the configuration utility.)",
			},
			"distributionalgorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will distribute records among the collectors. Else, all records will be sent to all the collectors.",
			},
			"metricslog": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If only the stats records are to be exported, turn on this option.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the AppFlow action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at\n(@), equals (=), and hyphen (-) characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my appflow action\" or 'my appflow action').",
			},
			"pagetracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will start tracking the page for waterfall chart by inserting a NS_ESNS cookie in the response.",
			},
			"securityinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the security insight records to the configured collectors.",
			},
			"transactionlog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Log ANOMALOUS or ALL transactions",
			},
			"videoanalytics": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the videoinsight records to the configured collectors.",
			},
			"webinsight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "On enabling this option, the Citrix ADC will send the webinsight records to the configured collectors.",
			},
		},
	}
}
