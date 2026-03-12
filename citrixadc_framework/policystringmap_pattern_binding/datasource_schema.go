package policystringmap_pattern_binding

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicystringmapPatternBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comments associated with the string map or key-value pair bound to this string map.",
			},
			"key": schema.StringAttribute{
				Required:    true,
				Description: "Character string constituting the key to be bound to the string map. The key is matched against the data processed by the operation that uses the string map. The default character set is ASCII. UTF-8 characters can be included if the character set is UTF-8.  UTF-8 characters can be entered directly (if the UI supports it) or can be encoded as a sequence of hexadecimal bytes '\\xNN'. For example, the UTF-8 character '' can be encoded as '\\xC3\\xBC'.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the string map to which to bind the key-value pair.",
			},
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Character string constituting the value associated with the key. This value is returned when processed data matches the associated key. Refer to the key parameter for details of the value character set.",
			},
		},
	}
}
