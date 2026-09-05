package appfwsignatures

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AppfwsignaturesDataSourceModel is the data-source-specific model, decoupled
// from AppfwsignaturesResourceModel. A data source is a pure read surface (Read
// only; no plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits. The Framework's per-attribute model <->
// schema reflection requires this model to have exactly the attributes the
// data-source schema declares, which is why it cannot reuse the resource model.
type AppfwsignaturesDataSourceModel struct {
	Id                      types.String `tfsdk:"id"`
	Action                  types.List   `tfsdk:"action"`
	Autoenablenewsignatures types.String `tfsdk:"autoenablenewsignatures"`
	Category                types.String `tfsdk:"category"`
	Comment                 types.String `tfsdk:"comment"`
	Enabled                 types.String `tfsdk:"enabled"`
	Merge                   types.Bool   `tfsdk:"merge"`
	Mergedefault            types.Bool   `tfsdk:"mergedefault"`
	Name                    types.String `tfsdk:"name"` // Required lookup key
	Overwrite               types.Bool   `tfsdk:"overwrite"`
	Preservedefactions      types.Bool   `tfsdk:"preservedefactions"`
	Ruleid                  types.List   `tfsdk:"ruleid"`
	Sha1                    types.String `tfsdk:"sha1"`
	Src                     types.String `tfsdk:"src"`
	Vendortype              types.String `tfsdk:"vendortype"`
	Xslt                    types.String `tfsdk:"xslt"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/appfwsignatures.json). Never settable; populated from GET.
	Response         types.String `tfsdk:"response"`
	Encryptedversion types.Int64  `tfsdk:"encryptedversion"`
}

func AppfwsignaturesDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Signature action",
			},
			"autoenablenewsignatures": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable auto enable new signatures",
			},
			"category": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Signature category to be Enabled/Disabled",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the signatures object.",
			},
			"enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Flag used to enable/disable enable signature rule IDs/Signature Category",
			},
			"merge": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges the existing Signature with new signature rules",
			},
			"mergedefault": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Merges signature file with default signature file.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the signature object.",
			},
			"overwrite": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Overwrite any existing signatures object of the same name.",
			},
			"preservedefactions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "preserves def actions of signature rules",
			},
			"ruleid": schema.ListAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Signature rule IDs to be Enabled/Disabled",
			},
			"sha1": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File path for sha1 file to validate signature file",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL (protocol, host, path, and file name) for the location at which to store the imported signatures object.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access.",
			},
			"vendortype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Third party vendor type for which WAF signatures has to be generated.",
			},
			"xslt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "XSLT file source.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"response": schema.StringAttribute{
				Computed:    true,
				Description: "Signature response returned by the appliance.",
			},
			"encryptedversion": schema.Int64Attribute{
				Computed:    true,
				Description: "Encrypted signature version.",
			},
		},
	}
}

// appfwsignaturesDataSourceSetAttrFromGet projects a NITRO appfwsignatures GET
// response onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func appfwsignaturesDataSourceSetAttrFromGet(ctx context.Context, data *AppfwsignaturesDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In appfwsignaturesDataSourceSetAttrFromGet Function")

	if v, ok := g["name"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Name = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs (Null when the GET omits them).
	data.Autoenablenewsignatures = utils.MapGetString(g, "autoenablenewsignatures")
	data.Category = utils.MapGetString(g, "category")
	data.Comment = utils.MapGetString(g, "comment")
	data.Enabled = utils.MapGetString(g, "enabled")
	data.Merge = utils.MapGetBool(g, "merge")
	data.Mergedefault = utils.MapGetBool(g, "mergedefault")
	data.Overwrite = utils.MapGetBool(g, "overwrite")
	data.Preservedefactions = utils.MapGetBool(g, "preservedefactions")
	data.Sha1 = utils.MapGetString(g, "sha1")
	data.Src = utils.MapGetString(g, "src")
	data.Vendortype = utils.MapGetString(g, "vendortype")
	data.Xslt = utils.MapGetString(g, "xslt")
	data.Action = utils.MapGetStringList(g, "action")

	// ruleid is an Int64-element list write-only input the GET never returns.
	data.Ruleid = types.ListNull(types.Int64Type)

	// Read-only (GET-only) attributes.
	data.Response = utils.MapGetString(g, "response")
	data.Encryptedversion = utils.MapGetInt64(g, "encryptedversion")
}
