package dnspolicy

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

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
		},
	}
}
