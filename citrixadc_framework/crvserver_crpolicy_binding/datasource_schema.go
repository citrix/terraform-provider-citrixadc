package crvserver_crpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CrvserverCrpolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from CrvserverCrpolicyBindingResourceModel so the data source can
// expose read-only (GET-only) attributes the resource omits.
type CrvserverCrpolicyBindingDataSourceModel struct {
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

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/crvserver_crpolicy_binding.json). Never settable;
	// populated from GET and null when the appliance omits them.
	Hits         types.Int64 `tfsdk:"hits"`
	Pipolicyhits types.Int64 `tfsdk:"pipolicyhits"`
}

func CrvserverCrpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Name of the virtual server to which content is forwarded. Applicable only if the policy is a map policy and the cache redirection virtual server is of type REVERSE.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
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

// crvserver_crpolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// crvserver_crpolicy_binding GET response onto the data-source model. The
// shared utils.MapGet* helpers fill each attribute from the GET (or leave it
// Null when the GET omits it).
func crvserver_crpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CrvserverCrpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In crvserver_crpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Bindpoint = utils.MapGetString(g, "bindpoint")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Targetvserver = utils.MapGetString(g, "targetvserver")

	// Lookup keys: prefer the GET value, but preserve the configured value when
	// the appliance omits it from the binding response.
	if v := utils.MapGetString(g, "name"); !v.IsNull() {
		data.Name = v
	}
	if v := utils.MapGetString(g, "policyname"); !v.IsNull() {
		data.Policyname = v
	}

	// Read-only (GET-only) attributes.
	data.Hits = utils.MapGetInt64(g, "hits")
	data.Pipolicyhits = utils.MapGetInt64(g, "pipolicyhits")

	// Composite key -> id (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
