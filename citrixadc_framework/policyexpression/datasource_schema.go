package policyexpression

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicyexpressionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"clientsecuritymessage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Message to display if the expression fails. Allowed for classic end-point check expressions only.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the expression. Displayed upon viewing the policy expression.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique name for the expression. Not case sensitive. Must begin with an ASCII letter or underscore (_) character, and must consist only of ASCII alphanumeric or underscore characters. Must not begin with 're' or 'xp' or be a word reserved for use as an expression qualifier prefix (such as HTTP) or enumeration value (such as ASCII). Must not be the name of an existing named expression, pattern set, dataset, stringmap, or HTTP callout.",
			},
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression string. For example: http.req.body(100).contains(\"this\").",
			},
		},
	}
}
