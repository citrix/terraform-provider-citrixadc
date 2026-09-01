package crvserver_cspolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrvserverCspolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from CrvserverCspolicyBindingResourceModel. A data source is a pure
// read surface (Read only; no plan/apply lifecycle), so it can expose the FULL
// GET projection: the read/write attributes (as Computed outputs) AND the
// read-only attributes that the resource deliberately omits.
type CrvserverCspolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Bindpoint              types.String `tfsdk:"bindpoint"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Name                   types.String `tfsdk:"name"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Targetvserver          types.String `tfsdk:"targetvserver"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/crvserver_cspolicy_binding.json). Never settable;
	// populated from GET.
	Hits         types.Int64 `tfsdk:"hits"`
	Pipolicyhits types.Int64 `tfsdk:"pipolicyhits"`
}

func CrvserverCspolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Name of the cache redirection virtual server to which to bind the cache redirection policy.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Policies bound to this vserver.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "The priority for the policy.",
			},
			"targetvserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The CSW target server names.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (these are intentionally NOT modeled on the resource). All Computed.
			"hits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
			"pipolicyhits": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of hits.",
			},
		},
	}
}

// crvserver_cspolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// crvserver_cspolicy_binding GET response onto the data-source model. Because a
// data source has no plan/apply reconciliation, attributes are simply filled
// from the GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func crvserver_cspolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CrvserverCspolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In crvserver_cspolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Name = utils.MapGetString(g, "name")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")

	// Read-only attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Pipolicyhits = utils.MapGetInt64(g, "pipolicyhits")

	// Set ID for the data source (composite key: name,policyname).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
