package auditsyslogglobal_auditsyslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuditsyslogglobalAuditsyslogpolicyBindingResourceModel describes the resource data model.
type AuditsyslogglobalAuditsyslogpolicyBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Feature        types.String `tfsdk:"feature"`
	Globalbindtype types.String `tfsdk:"globalbindtype"`
	Policyname     types.String `tfsdk:"policyname"`
	Priority       types.Int64  `tfsdk:"priority"`
}

func (r *AuditsyslogglobalAuditsyslogpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditsyslogglobal_auditsyslogpolicy_binding resource.",
			},
			"feature": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The feature to be checked while applying this config",
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SYSTEM_GLOBAL"),
				Description: "0",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit syslog policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func auditsyslogglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AuditsyslogglobalAuditsyslogpolicyBindingResourceModel) audit.Auditsyslogglobalauditsyslogpolicybinding {
	tflog.Debug(ctx, "In auditsyslogglobal_auditsyslogpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditsyslogglobal_auditsyslogpolicy_binding := audit.Auditsyslogglobalauditsyslogpolicybinding{}
	if !data.Feature.IsNull() {
		auditsyslogglobal_auditsyslogpolicy_binding.Feature = data.Feature.ValueString()
	}
	if !data.Globalbindtype.IsNull() {
		auditsyslogglobal_auditsyslogpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Policyname.IsNull() {
		auditsyslogglobal_auditsyslogpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		auditsyslogglobal_auditsyslogpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return auditsyslogglobal_auditsyslogpolicy_binding
}

func auditsyslogglobal_auditsyslogpolicy_bindingSetAttrFromGet(ctx context.Context, data *AuditsyslogglobalAuditsyslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuditsyslogglobalAuditsyslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In auditsyslogglobal_auditsyslogpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["feature"]; ok && val != nil {
		data.Feature = types.StringValue(val.(string))
	} else {
		data.Feature = types.StringNull()
	}
	if val, ok := getResponseData["globalbindtype"]; ok && val != nil {
		data.Globalbindtype = types.StringValue(val.(string))
	} else {
		data.Globalbindtype = types.StringNull()
	}
	if val, ok := getResponseData["policyname"]; ok && val != nil {
		data.Policyname = types.StringValue(val.(string))
	} else {
		data.Policyname = types.StringNull()
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else {
		data.Priority = types.Int64Null()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("globalbindtype:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Globalbindtype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
