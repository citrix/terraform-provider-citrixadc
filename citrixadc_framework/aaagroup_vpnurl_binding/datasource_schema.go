package aaagroup_vpnurl_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaagroupVpnurlBindingDataSourceModel is the data-source-specific model,
// decoupled from the resource model. A data source is a pure read surface (Read
// only; no plan/apply lifecycle), so it can expose the FULL GET projection: the
// read/write attributes (as Computed outputs) AND the read-only attributes that
// the resource deliberately omits (e.g. acttype). Every non-key attribute is
// Computed.
type AaagroupVpnurlBindingDataSourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Gotopriorityexpression types.String `tfsdk:"gotopriorityexpression"`
	Groupname              types.String `tfsdk:"groupname"`
	Urlname                types.String `tfsdk:"urlname"`

	// Read-only (GET-only) attributes from the NITRO doc read-only set
	// (zion73x_readonly/aaagroup_vpnurl_binding.json).
	Acttype types.Int64 `tfsdk:"acttype"`
}

func AaagroupVpnurlBindingDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"gotopriorityexpression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression or other value specifying the next policy to evaluate if the current policy evaluates to TRUE.  Specify one of the following values:\n* NEXT - Evaluate the policy with the next higher priority number.\n* END - End policy evaluation.\n* USE_INVOCATION_RESULT - Applicable if this policy invokes another policy label. If the final goto in the invoked policy label has a value of END, the evaluation stops. If the final goto is anything other than END, the current policy label performs a NEXT.\n* An expression that evaluates to a number.\nIf you specify an expression, the number to which it evaluates determines the next policy to evaluate, as follows:\n*  If the expression evaluates to a higher numbered priority, the policy with that priority is evaluated next.\n* If the expression evaluates to the priority of the current policy, the policy with the next higher numbered priority is evaluated next.\n* If the expression evaluates to a number that is larger than the largest numbered priority, policy evaluation ends.\nAn UNDEF event is triggered if:\n* The expression is invalid.\n* The expression evaluates to a priority number that is numerically lower than the current policy's priority.\n* The expression evaluates to a priority number that is between the current policy's priority number (say, 30) and the highest priority number (say, 100), but does not match any configured priority number (for example, the expression evaluates to the number 85). This example assumes that the priority number increments by 10 for every successive policy, and therefore a priority number of 85 does not exist in the policy label.",
			},
			"groupname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the group that you are binding.",
			},
			"urlname": schema.StringAttribute{
				Required:    true,
				Description: "The intranet url",
			},

			// Read-only (GET-only) attribute surfaced only by the data source.
			"acttype": schema.Int64Attribute{
				Computed:    true,
				Description: "Action type of the binding. Read-only value returned by the appliance.",
			},
		},
	}
}

// aaagroup_vpnurl_bindingDataSourceSetAttrFromGet projects a NITRO GET response
// onto the data-source model. Because a data source has no plan/apply
// reconciliation, attributes are simply filled from the GET (or left Null when the
// GET omits them). The shared utils.MapGet* helpers implement that projection.
func aaagroup_vpnurl_bindingDataSourceSetAttrFromGet(ctx context.Context, data *AaagroupVpnurlBindingDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaagroup_vpnurl_bindingDataSourceSetAttrFromGet Function")

	data.Gotopriorityexpression = utils.MapGetString(g, "gotopriorityexpression")
	data.Groupname = utils.MapGetString(g, "groupname")
	data.Urlname = utils.MapGetString(g, "urlname")

	// Read-only (GET-only) attribute.
	data.Acttype = utils.MapGetInt64(g, "acttype")

	// Composite ID (groupname + urlname).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("groupname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Groupname.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("urlname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Urlname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))
}
