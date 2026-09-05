package csvserver_cspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CsvserverCspolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from CsvserverCspolicyBindingResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes that the resource deliberately omits.
type CsvserverCspolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetlbvserver        types.String `tfsdk:"targetlbvserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/csvserver_cspolicy_binding.json). Never settable;
	// populated from GET.
	Rule         types.String `tfsdk:"rule"`
	Cookieipport types.String `tfsdk:"cookieipport"`
	Hits         types.Int64  `tfsdk:"hits"`
	Pipolicyhits types.Int64  `tfsdk:"pipolicyhits"`
	Vserverid    types.String `tfsdk:"vserverid"`
}

func CsvserverCspolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"bindpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The bindpoint to which the policy is bound.",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke flag.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label invoked.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The invocation type.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the content switching virtual server to which the content switching policy applies.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Policies bound to this vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority for the policy.",
			},
			"targetlbvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "target vserver name.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"rule": schema.StringAttribute{
				Computed:    true,
				Description: "Rule.",
			},
			"cookieipport": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver id of the lb vserver that is inserted into the set-cookie HTTP header.",
			},
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"pipolicyhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"vserverid": schema.StringAttribute{
				Computed:    true,
				Description: "Vserver Id of vserver.",
			},
		},
	}
}

// csvserver_cspolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// csvserver_cspolicy_binding GET response onto the data-source model via the
// shared utils.MapGet* helpers.
func csvserver_cspolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CsvserverCspolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In csvserver_cspolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Name = utils.MapGetString(g, "name")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Targetlbvserver = utils.MapGetString(g, "targetlbvserver")

	// Read-only attributes.
	data.Rule = utils.MapGetString(g, "rule")
	data.Cookieipport = utils.MapGetString(g, "cookieipport")
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Pipolicyhits = utils.MapGetInt64(g, "pipolicyhits")
	data.Vserverid = utils.MapGetString(g, "vserverid")

	// Set ID for the data source (composite key: name,policyname).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
