package cacheglobal_cachepolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// CacheglobalCachepolicyBindingDataSourceModel is the data-source-specific
// model, decoupled from CacheglobalCachepolicyBindingResourceModel. A data
// source is a pure read surface, so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits.
type CacheglobalCachepolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policy                 types.String `tfsdk:"policy"` // Required lookup key
	Precededefrules        types.String `tfsdk:"precededefrules"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/cacheglobal_cachepolicy_binding.json).
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func CacheglobalCachepolicyBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "0",
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the priority of the next policy which will get evaluated if the current policy rule evaluates to TRUE.",
			},
			"invoke": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Invoke policies bound to a virtual server or a user-defined policy label. After the invoked policies are evaluated, the flow returns to the policy with the next priority. Applicable only to default-syntax policies.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the label to invoke if the current policy rule evaluates to TRUE. (To invoke a label associated with a virtual server, specify the name of the virtual server.)",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of policy label to invoke.",
			},
			"policy": schema.StringAttribute{
				Required:    true,
				Description: "Name of the cache policy.",
			},
			"precededefrules": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify whether this policy should be evaluated.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "The bind point to which policy is bound. When you specify the type, detailed information about that bind point appears.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound cache policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// cacheglobal_cachepolicy_bindingComposeIdForDatasource builds the composite
// resource ID for the data source using the legacy attribute order (policy,
// type) in the new key:value form.
func cacheglobal_cachepolicy_bindingComposeIdForDatasource(data *CacheglobalCachepolicyBindingDataSourceModel) string {
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policy:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policy.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	return strings.Join(idParts, ",")
}

// cacheglobal_cachepolicy_bindingDataSourceSetAttrFromGet projects a NITRO
// cacheglobal_cachepolicy_binding GET response onto the data-source model.
func cacheglobal_cachepolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *CacheglobalCachepolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In cacheglobal_cachepolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policy = utils.MapGetString(g, "policy")
	data.Precededefrules = utils.MapGetString(g, "precededefrules")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only attributes.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set the composite id for the datasource.
	data.Id = types.StringValue(cacheglobal_cachepolicy_bindingComposeIdForDatasource(data))
}
