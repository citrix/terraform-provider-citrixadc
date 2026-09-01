package botsignature

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotsignatureDataSourceModel is the data-source-specific model, decoupled from
// BotsignatureResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (response). Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type BotsignatureDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"` // Required lookup key
	Overwrite types.Bool   `tfsdk:"overwrite"`
	Src       types.String `tfsdk:"src"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/botsignature.json). Never settable; populated from GET.
	Response types.String `tfsdk:"response"`
}

func BotsignatureDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the signature file object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name to assign to the bot signature file object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported signature file.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Response returned by the appliance for the bot signature.",
			},
		},
	}
}

// botsignatureDataSourceSetAttrFromGet projects a NITRO botsignature GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func botsignatureDataSourceSetAttrFromGet(ctx context.Context, data *BotsignatureDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In botsignatureDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs. comment and overwrite are not
	// returned by the GET, so MapGet* resolves them to Null.
	data.Comment = utils.MapGetString(g, "comment")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Src = utils.MapGetString(g, "src")

	// Read-only metadata.
	data.Response = utils.MapGetString(g, "response")
}
