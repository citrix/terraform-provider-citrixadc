package authenticationpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationpolicylabelDataSourceModel is the data-source-specific model,
// decoupled from AuthenticationpolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits (numpol,
// hits, policyname, priority, gotopriorityexpression, flowtype, description).
// Every non-key attribute is Computed.
type AuthenticationpolicylabelDataSourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Labelname   types.String `tfsdk:"labelname"`
	Loginschema types.String `tfsdk:"loginschema"`
	Newname     types.String `tfsdk:"newname"`
	Type        types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/authenticationpolicylabel.json). Never settable;
	// populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Flowtype               types.Int64  `tfsdk:"flowtype"`
	Description            types.String `tfsdk:"description"`
}

func AuthenticationpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about this authentication policy label.",
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authentication policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy label\" or 'authentication policy label').",
			},
			"loginschema": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is \"noschema\" should be used.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the auth policy label",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of feature (aaatm or rba) against which to match the policies bound to this policy label.",
			},

			// Read-only (GET-only) metadata surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of polices bound to label.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of times policy label was invoked.",
			},
			"policyname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the authentication policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowtype of the bound authentication policy.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policylabel.",
			},
		},
	}
}

// authenticationpolicylabelDataSourceSetAttrFromGet projects a NITRO
// authenticationpolicylabel GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them). The shared utils.MapGet*
// helpers implement that projection.
func authenticationpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *AuthenticationpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authenticationpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	data.Comment = utils.MapGetString(g, "comment")
	data.Loginschema = utils.MapGetString(g, "loginschema")
	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()
	data.Type = utils.MapGetString(g, "type")

	// Read-only metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Description = utils.MapGetString(g, "description")
}
