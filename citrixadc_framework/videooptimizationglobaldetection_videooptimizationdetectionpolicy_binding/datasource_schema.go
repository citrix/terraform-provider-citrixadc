package videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingDataSourceModel
// is the data-source-specific model, decoupled from the resource model. A data
// source is a pure read surface (Read only; no plan/apply lifecycle), so it can
// expose the full GET projection: the read/write attributes (as Computed outputs)
// AND the read-only attributes that the resource deliberately omits (numpol).
type VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"`
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"`

	// Read-only (GET-only) metadata from the NITRO doc read-only set
	// (zion73x_readonly/videooptimizationglobaldetection_videooptimizationdetectionpolicy_binding.json).
	Numpol types.Int64 `tfsdk:"numpol"`
}

func VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server or\nevaluate the specified policy label.",
			},
			"labelname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the policy label to invoke. If the current policy evaluates to TRUE, the invoke parameter is set, and Label Type is policylabel.",
			},
			"labeltype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of invocation, Available settings function as follows:\n* vserver - Forward the request to the specified virtual server.\n* policylabel - Invoke the specified policy label.",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the videooptimization detection policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the bind point whose policies you want to display.",
			},

			// Read-only (GET-only) metadata surfaced by the data source. All Computed.
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "number of polices bound.",
			},
		},
	}
}

// videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSourceSetAttrFromGet
// projects a NITRO GET response onto the data-source model using the shared
// utils.MapGet* helpers.
func videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *VideooptimizationglobaldetectionVideooptimizationdetectionpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In videooptimizationglobaldetection_videooptimizationdetectionpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) metadata.
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Set composite ID (policyname,priority,type) preserving the resource ID format.
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(data.Policyname.ValueString())))
	idParts = append(idParts, fmt.Sprintf("priority:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Priority.ValueInt64()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(data.Type.ValueString())))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
