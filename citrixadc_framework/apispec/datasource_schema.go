package apispec

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ApispecDataSourceModel is the data-source-specific model, decoupled from
// ApispecResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only metadata attributes that the resource deliberately
// omits (builtin, ready, nsappversion). Every non-key attribute is Computed; the
// Framework's per-attribute model <-> schema reflection requires this model to
// have exactly the attributes the data-source schema declares, which is why it
// cannot reuse the resource model.
type ApispecDataSourceModel struct {
	Id             types.String `tfsdk:"id"`
	Encrypted      types.Bool   `tfsdk:"encrypted"`
	File           types.String `tfsdk:"file"`
	Name           types.String `tfsdk:"name"`
	Skipvalidation types.String `tfsdk:"skipvalidation"`
	Type           types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/apispec.json). Never settable; populated from GET.
	Builtin      types.List   `tfsdk:"builtin"`
	Ready        types.String `tfsdk:"ready"`
	Nsappversion types.String `tfsdk:"nsappversion"`
}

func ApispecDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encrypted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the encrypted API spec. Must be in NetScaler format",
			},
			"file": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of and, optionally, path to the api spec file. The spec file should be present on the appliance's hard-disk drive or solid-state drive. Storing a spec file in any location other than the default might cause inconsistency in a high availability setup. /nsconfig/apispec/ is the default path.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the spec. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the spec is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my spec\" or 'my spec').",
			},
			"skipvalidation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disabling openapi spec validation while adding it",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input format of the spec file. The three formats supported by the appliance are:\nPROTO \nOAS/Swagger\nGRAPHQL",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Indicates that a variable is a built-in (SYSTEM INTERNAL) type.",
			},
			"ready": schema.StringAttribute{
				Computed:    true,
				Description: "Indicates whether the api spec is ready, that is whether the internal registry is created or not.",
			},
			"nsappversion": schema.StringAttribute{
				Computed:    true,
				Description: "NS App Version of the api spec file.",
			},
		},
	}
}

// apispecDataSourceSetAttrFromGet projects a NITRO apispec GET response onto the
// data-source model. Because a data source has no plan/apply reconciliation,
// attributes are simply filled from the GET (or left Null when the GET omits
// them). The shared utils.MapGet* helpers implement that projection.
func apispecDataSourceSetAttrFromGet(ctx context.Context, data *ApispecDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In apispecDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Encrypted = utils.MapGetBool(g, "encrypted")
	data.File = utils.MapGetString(g, "file")
	data.Skipvalidation = utils.MapGetString(g, "skipvalidation")
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Ready = utils.MapGetString(g, "ready")
	data.Nsappversion = utils.MapGetString(g, "nsappversion")
}
