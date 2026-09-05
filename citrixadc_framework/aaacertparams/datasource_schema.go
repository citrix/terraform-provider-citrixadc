package aaacertparams

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaacertparamsDataSourceModel is the data-source-specific model, decoupled from
// AaacertparamsResourceModel. A data source is a pure read surface, so it can
// expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AaacertparamsDataSourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupnamefield             types.String `tfsdk:"groupnamefield"`
	Usernamefield              types.String `tfsdk:"usernamefield"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/aaacertparams.json). Never settable; populated from GET.
	Twofactor types.String `tfsdk:"twofactor"`
}

func AaacertparamsDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that specifies the group, in the format <field>:<subfield>.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that contains the username, in the format <field>:<subfield>.",
			},

			// Read-only (GET-only) attributes surfaced by the data source. All Computed.
			"twofactor": schema.StringAttribute{
				Computed:    true,
				Description: "The state of the two-factor authentication. Possible values = ON, OFF.",
			},
		},
	}
}

// aaacertparamsDataSourceSetAttrFromGet projects a NITRO aaacertparams GET
// response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaacertparamsDataSourceSetAttrFromGet(ctx context.Context, data *AaacertparamsDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaacertparamsDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Defaultauthenticationgroup = utils.MapGetString(g, "defaultauthenticationgroup")
	data.Groupnamefield = utils.MapGetString(g, "groupnamefield")
	data.Usernamefield = utils.MapGetString(g, "usernamefield")

	// Read-only attributes.
	data.Twofactor = utils.MapGetString(g, "twofactor")

	// aaacertparams is a singleton; the ID is a fixed system-generated identifier.
	data.Id = types.StringValue("aaacertparams-config")
}
