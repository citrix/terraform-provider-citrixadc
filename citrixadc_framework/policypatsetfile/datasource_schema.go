package policypatsetfile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func PolicypatsetfileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"charset": schema.StringAttribute{
				Computed:    true,
				Description: "Character set associated with the characters in the string.",
			},
			"comment": schema.StringAttribute{
				Computed:    true,
				Description: "Any comments to preserve information about this patsetfile.",
			},
			"delimiter": schema.StringAttribute{
				Computed:    true,
				Description: "patset file patterns delimiter.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the imported patset file. Unique name of the pattern set. Not case sensitive. Must begin with an ASCII letter or underscore (_) character and must contain only alphanumeric and underscore characters.",
			},
			"overwrite": schema.BoolAttribute{
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Computed:    true,
				Description: "URL in protocol, host, path, and file name format from where the patset file will be imported. If file is already present, then it can be imported using local keyword (import patsetfile local:filename patsetfile1)\n                      NOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access",
			},
		},
	}
}
