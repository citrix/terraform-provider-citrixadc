package policypatset_pattern_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicypatsetPatternBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"string": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String of characters that constitutes a pattern. For more information about the characters that can be used, refer to the character set parameter.\nNote: Minimum length for pattern sets used in rewrite actions of type REPLACE_ALL, DELETE_ALL, INSERT_AFTER_ALL, and INSERT_BEFORE_ALL, is three characters.",
			},
			"charset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Character set associated with the characters in the string.\nNote: UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\\xNN'. For example, the UTF-8 character '' can be encoded as '\\xC3\\xBC'.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this patset or a pattern bound to this patset.",
			},
			"feature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The feature to be checked while applying this config",
			},
			"index": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The index of the string associated with the patset.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the pattern set to which to bind the string.",
			},
		},
	}
}
