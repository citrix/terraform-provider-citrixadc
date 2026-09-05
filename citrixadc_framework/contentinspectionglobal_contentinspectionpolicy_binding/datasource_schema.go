package contentinspectionglobal_contentinspectionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// ContentinspectionglobalContentinspectionpolicyBindingDataSourceModel is the
// data-source-specific model, decoupled from the resource model so the data
// source can expose read-only (GET-only) attributes the resource omits.
type ContentinspectionglobalContentinspectionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/contentinspectionglobal_contentinspectionpolicy_binding.json).
	// Never settable; populated from GET and null when the appliance omits them.
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func ContentinspectionglobalContentinspectionpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "Terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or evaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "* If labelType is policylabel, name of the policy label to invoke.\n* If labelType is reqvserver or resvserver, name of the virtual server to which to forward the request of response.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation. Available settings function as follows:\n* reqvserver - Forward the request to the specified request virtual server.\n* resvserver - Forward the response to the specified response virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the contentInspection policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "The bindpoint to which to policy is bound.",
			},

			// Read-only (GET-only) attributes surfaced by the data source.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound contentInspection policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "The number of policies bound to the bindpoint.",
			},
		},
	}
}

// contentinspectionglobal_contentinspectionpolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO contentinspectionglobal_contentinspectionpolicy_binding GET
// response onto the data-source model. The shared utils.MapGet* helpers fill
// each attribute from the GET (or leave it Null when the GET omits it).
func contentinspectionglobal_contentinspectionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *ContentinspectionglobalContentinspectionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In contentinspectionglobal_contentinspectionpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Priority = utils.MapGetInt64(g, "priority")

	// Lookup keys: prefer the GET value, but preserve the configured value when
	// the appliance omits it from the binding response.
	if v := utils.MapGetString(g, "policyname"); !v.IsNull() {
		data.Policyname = v
	}
	if v := utils.MapGetString(g, "type"); !v.IsNull() {
		data.Type = v
	}

	// Read-only (GET-only) attributes.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Composite key -> id (key:UrlEncode(value) pairs).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
