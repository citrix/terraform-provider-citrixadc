package locationparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func LocationparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"context": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Context for describing locations. In geographic context, qualifier labels are assigned by default in the following sequence: Continent.Country.Region.City.ISP.Organization. In custom context, the qualifiers labels can have any meaning that you designate.",
			},
			"matchwildcardtoany": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Indicates whether wildcard qualifiers should match any other\nqualifier including non-wildcard while evaluating\nlocation based expressions.\nPossible values: Yes, No, Expression.\n    Yes - Wildcard qualifiers match any other qualifiers.\n    No  - Wildcard qualifiers do not match non-wildcard\n          qualifiers, but match other wildcard qualifiers.\n    Expression - Wildcard qualifiers in an expression\n          match any qualifier in an LDNS location,\n          wildcard qualifiers in the LDNS location do not match\n          non-wildcard qualifiers in an expression",
			},
			"q1label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the first qualifier. Can be specified for custom context only.",
			},
			"q2label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the second qualifier. Can be specified for custom context only.",
			},
			"q3label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the third qualifier. Can be specified for custom context only.",
			},
			"q4label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fourth qualifier. Can be specified for custom context only.",
			},
			"q5label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the fifth qualifier. Can be specified for custom context only.",
			},
			"q6label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Label specifying the meaning of the sixth qualifier. Can be specified for custom context only.",
			},
		},
	}
}
