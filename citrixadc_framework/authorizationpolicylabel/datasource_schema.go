package authorizationpolicylabel

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthorizationpolicylabelDataSourceModel is the data-source-specific model,
// decoupled from AuthorizationpolicylabelResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only runtime/binding attributes that the resource
// deliberately omits. Every non-key attribute is Computed; the Framework's
// per-attribute model <-> schema reflection requires this model to have exactly
// the attributes the data-source schema declares, which is why it cannot reuse
// the resource model.
type AuthorizationpolicylabelDataSourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`

	// Read-only (GET-only) runtime/binding attributes from the NITRO doc
	// read-only set (zion73x_readonly/authorizationpolicylabel.json). Never
	// settable; populated from GET.
	Numpol                 types.Int64  `tfsdk:"numpol"`
	Hits                   types.Int64  `tfsdk:"hits"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Labeltype              types.String `tfsdk:"labeltype"`
	InvokeLabelname        types.String `tfsdk:"invoke_labelname"`
	Flowtype               types.Int64  `tfsdk:"flowtype"`
	Description            types.String `tfsdk:"description"`
}

func AuthorizationpolicylabelDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"labelname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the new authorization policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy label\" or 'authorization policy label').",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the auth policy label",
			},

			// Read-only (GET-only) runtime/binding attributes surfaced by the data
			// source (these are intentionally NOT modeled on the resource). All Computed.
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
				Description: "Name of the authorization policy to bind to the policy label.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "Type of invocation (reqvserver, resvserver, policylabel).",
			},
			"invoke_labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the policy label to invoke if the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is set to Policy Label.",
			},
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "Flowtype of the bound authorization policy.",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Description of the policylabel.",
			},
		},
	}
}

// authorizationpolicylabelDataSourceSetAttrFromGet projects a NITRO
// authorizationpolicylabel GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled from
// the GET (or left Null when the GET omits them) — no unknown->null resolution or
// plan preservation is required. The shared utils.MapGet* helpers implement that
// projection.
func authorizationpolicylabelDataSourceSetAttrFromGet(ctx context.Context, data *AuthorizationpolicylabelDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In authorizationpolicylabelDataSourceSetAttrFromGet Function")

	if v, ok := g["labelname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Labelname = types.StringValue(utils.AnyToString(v))
	}

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only runtime/binding attributes.
	data.Numpol = utils.MapGetInt64(g, "numpol")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.InvokeLabelname = utils.MapGetString(g, "invoke_labelname")
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Description = utils.MapGetString(g, "description")
}
