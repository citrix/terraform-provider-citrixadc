package crpolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrpolicyDataSourceModel is the data-source-specific model, decoupled from
// CrpolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only counter/metadata attributes that the resource
// deliberately omits (boundto, vstype, hits, priority, activepolicy, labelname,
// labeltype, builtin, feature, isdefault). Every non-key attribute is Computed.
type CrpolicyDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Logaction  types.String `tfsdk:"logaction"`
	Newname    types.String `tfsdk:"newname"`
	Policyname types.String `tfsdk:"policyname"` // Required lookup key
	Rule       types.String `tfsdk:"rule"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/crpolicy.json). Never settable.
	Boundto      types.String `tfsdk:"boundto"`
	Vstype       types.Int64  `tfsdk:"vstype"`
	Hits         types.Int64  `tfsdk:"hits"`
	Priority     types.Int64  `tfsdk:"priority"`
	Activepolicy types.Bool   `tfsdk:"activepolicy"`
	Labelname    types.String `tfsdk:"labelname"`
	Labeltype    types.String `tfsdk:"labeltype"`
	Builtin      types.List   `tfsdk:"builtin"`
	Feature      types.String `tfsdk:"feature"`
	Isdefault    types.Bool   `tfsdk:"isdefault"`
}

func CrpolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the built-in cache redirection action: CACHE/ORIGIN.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the cache redirection policy",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the content switching policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the cache redirection policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},

			// Read-only (GET-only) metadata surfaced by the data source (these are
			// intentionally NOT modeled on the resource). All Computed.
			"boundto": schema.StringAttribute{
				Computed:    true,
				Description: "Domain name.",
			},
			"vstype": schema.Int64Attribute{
				Computed:    true,
				Description: "Virtual server type.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of hits.",
			},
			"priority": schema.Int64Attribute{
				Computed:    true,
				Description: "priority of bound policy.",
			},
			"activepolicy": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates whether policy is bound or not.",
			},
			"labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "The invocation type. Possible values = reqvserver, resvserver, policylabel",
			},
			"builtin": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: "Flag to determine if the cr policy is built-in or not. Possible values = MODIFIABLE, DELETABLE, IMMUTABLE, PARTITION_ALL",
			},
			"feature": schema.StringAttribute{
				Computed:    true,
				Description: "The feature to be checked while applying this config.",
			},
			"isdefault": schema.BoolAttribute{
				Computed:    true,
				Description: "A value of true is returned if it is a default cr policy.",
			},
		},
	}
}

// crpolicyDataSourceSetAttrFromGet projects a NITRO crpolicy GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func crpolicyDataSourceSetAttrFromGet(ctx context.Context, data *CrpolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In crpolicyDataSourceSetAttrFromGet Function")

	if v, ok := g["policyname"]; ok && v != nil {
		data.Id = types.StringValue(utils.AnyToString(v))
		data.Policyname = types.StringValue(utils.AnyToString(v))
	}

	// Read/write attributes as read-back outputs.
	data.Action = utils.MapGetString(g, "action")
	data.Logaction = utils.MapGetString(g, "logaction")
	data.Rule = utils.MapGetString(g, "rule")

	// newname is a rename-only input the GET never returns -> Null.
	data.Newname = types.StringNull()

	// Read-only metadata.
	data.Boundto = utils.MapGetString(g, "boundto")
	data.Vstype = utils.MapGetInt64(g, "vstype")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Activepolicy = utils.MapGetBool(g, "activepolicy")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Builtin = utils.MapGetStringList(g, "builtin")
	data.Feature = utils.MapGetString(g, "feature")
	data.Isdefault = utils.MapGetBool(g, "isdefault")
}
