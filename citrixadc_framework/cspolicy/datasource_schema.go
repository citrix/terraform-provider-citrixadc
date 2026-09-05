package cspolicy

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CspolicyDataSourceModel is the data-source-specific model, decoupled from
// CspolicyResourceModel.
//
// A data source is a pure read surface (Read only; no plan/apply lifecycle), so
// it can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes that the resource deliberately omits
// (vstype, hits, labelname, labeltype, activepolicy, boundto). Every non-key
// attribute is Computed; the Framework's per-attribute model <-> schema
// reflection requires this model to have exactly the attributes the data-source
// schema declares, which is why it cannot reuse the resource model.
type CspolicyDataSourceModel struct {
	Id         types.String `tfsdk:"id"`
	Policyname types.String `tfsdk:"policyname"` // Required lookup key

	// Read/write attributes, surfaced here as Computed outputs.
	Action    types.String `tfsdk:"action"`
	Logaction types.String `tfsdk:"logaction"`
	Newname   types.String `tfsdk:"newname"`
	Rule      types.String `tfsdk:"rule"`

	// Provider-side binding helpers (not NITRO cspolicy fields).
	Csvserver       types.String `tfsdk:"csvserver"`
	Targetlbvserver types.String `tfsdk:"targetlbvserver"`
	ForcenewIdSet   types.Set    `tfsdk:"forcenew_id_set"`
	Priority        types.Int64  `tfsdk:"priority"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cspolicy.json). Never settable; populated from GET.
	// (priority from the read-only set is already modeled above.)
	Vstype       types.Int64  `tfsdk:"vstype"`
	Hits         types.Int64  `tfsdk:"hits"`
	Labelname    types.String `tfsdk:"labelname"`
	Labeltype    types.String `tfsdk:"labeltype"`
	Activepolicy types.Bool   `tfsdk:"activepolicy"`
	Boundto      types.String `tfsdk:"boundto"`
}

func CspolicyDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content switching action that names the target load balancing virtual server to which the traffic is switched.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The log action associated with the content switching policy",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The new name of the content switching policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the content switching policy. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after a policy is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my policy\" or 'my policy').",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\nThe following requirements apply only to the Citrix ADC CLI:\n*  If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n*  If the expression itself includes double quotation marks, escape the quotations by using the  character.\n*  Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"csvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The content switching vserver to which the cspolicy should be bound.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The target load balancing vserver for the csvserver policy binding.",
			},
			"forcenew_id_set": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "Auxiliary set attribute used to force recreation of the resource.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for the csvserver policy binding.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"vstype": schema.Int64Attribute{
				Computed:    true,
				Description: "Virtual server type.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Total number of hits.",
			},
			"labelname": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				Computed:    true,
				Description: "The invocation type. Possible values = reqvserver, resvserver, policylabel.",
			},
			"activepolicy": schema.BoolAttribute{
				Computed:    true,
				Description: "Indicates whether policy is bound or not.",
			},
			"boundto": schema.StringAttribute{
				Computed:    true,
				Description: "Location where policy is bound.",
			},
		},
	}
}

// cspolicyDataSourceSetAttrFromGet projects a NITRO cspolicy GET response onto
// the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when
// the GET omits them). The shared utils.MapGet* helpers implement that
// projection.
func cspolicyDataSourceSetAttrFromGet(ctx context.Context, data *CspolicyDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cspolicyDataSourceSetAttrFromGet Function")

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

	// Provider-side binding helpers are not returned by the single-resource GET.
	data.Csvserver = utils.MapGetString(g, "csvserver")
	data.Targetlbvserver = utils.MapGetString(g, "targetlbvserver")
	data.ForcenewIdSet = types.SetNull(types.StringType)
	data.Priority = utils.MapGetInt64(g, "priority")

	// Read-only attributes.
	data.Vstype = utils.MapGetInt64(g, "vstype")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Activepolicy = utils.MapGetBool(g, "activepolicy")
	data.Boundto = utils.MapGetString(g, "boundto")
}
