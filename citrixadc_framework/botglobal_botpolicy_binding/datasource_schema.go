package botglobal_botpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// BotglobalBotpolicyBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface, so it
// exposes the full GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits. Every
// non-key attribute is Computed.
type BotglobalBotpolicyBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Globalbindtype         types.String `tfsdk:"globalbindtype"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Invoke                 types.Bool   `tfsdk:"invoke"`
	Labelname              types.String `tfsdk:"labelname"`
	Labeltype              types.String `tfsdk:"labeltype"`
	Policyname             types.String `tfsdk:"policyname"` // Required lookup key
	Priority               types.Int64  `tfsdk:"priority"`
	Type                   types.String `tfsdk:"type"` // Required lookup key

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/botglobal_botpolicy_binding.json).
	Flowtype types.Int64 `tfsdk:"flowtype"`
	Numpol   types.Int64 `tfsdk:"numpol"`
}

func BotglobalBotpolicyBindingDataSourceSchema() schema.Schema {
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
				Description: "If the current policy evaluates to TRUE, terminate evaluation of policies bound to the current policy label, and then forward the request to the specified virtual server.",
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
				Description: "Name of the bot policy.",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the priority of the policy.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Specifies the bind point whose policies you want to display. Available settings function as follows:\n* REQ_OVERRIDE - Request override. Binds the policy to the priority request queue.\n* REQ_DEFAULT - Binds the policy to the default request queue.",
			},

			// Read-only (GET-only) attributes surfaced by the data source
			// (intentionally NOT modeled on the resource). All Computed.
			"flowtype": schema.Int64Attribute{
				Computed:    true,
				Description: "flowtype of the bound bot policy.",
			},
			"numpol": schema.Int64Attribute{
				Computed:    true,
				Description: "number of polices bound to label.",
			},
		},
	}
}

// botglobal_botpolicy_bindingDataSourceSetAttrFromGet projects a NITRO GET
// response onto the data-source model. Attributes are simply filled from the GET
// (or left Null when the GET omits them) via the shared utils.MapGet* helpers.
func botglobal_botpolicy_bindingDataSourceSetAttrFromGet(ctx context.Context, data *BotglobalBotpolicyBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In botglobal_botpolicy_bindingDataSourceSetAttrFromGet Function")

	data.Globalbindtype = utils.MapGetString(g, "globalbindtype")
	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Invoke = utils.MapGetBool(g, "invoke")
	data.Labelname = utils.MapGetString(g, "labelname")
	data.Labeltype = utils.MapGetString(g, "labeltype")
	data.Policyname = utils.MapGetString(g, "policyname")
	data.Priority = utils.MapGetInt64(g, "priority")
	data.Type = utils.MapGetString(g, "type")

	// Read-only (GET-only) attributes.
	data.Flowtype = utils.MapGetInt64(g, "flowtype")
	data.Numpol = utils.MapGetInt64(g, "numpol")

	// Compose the ID for the datasource (no Create step). ID = policyname,type
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("type:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Type.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
