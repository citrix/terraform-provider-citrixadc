package cachepolicylabel

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CachepolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"evaluates": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When to evaluate policies bound to this label: request-time or response-time.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Can be changed after the label is created.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the cache-policy label. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
		},
	}
}
