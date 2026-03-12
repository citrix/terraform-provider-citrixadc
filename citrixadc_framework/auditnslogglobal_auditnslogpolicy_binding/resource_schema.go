package auditnslogglobal_auditnslogpolicy_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuditnslogglobalAuditnslogpolicyBindingResourceModel describes the resource data model.
type AuditnslogglobalAuditnslogpolicyBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Globalbindtype types.String `tfsdk:"globalbindtype"`
	Policyname     types.String `tfsdk:"policyname"`
	Priority       types.Int64  `tfsdk:"priority"`
}

func (r *AuditnslogglobalAuditnslogpolicyBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditnslogglobal_auditnslogpolicy_binding resource.",
			},
			"globalbindtype": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("SYSTEM_GLOBAL"),
				Description: "0",
			},
			"policyname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the audit nslog policy.",
			},
			"priority": schema.Int64Attribute{
				Required:    true,
				Description: "Specifies the priority of the policy.",
			},
		},
	}
}

func auditnslogglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig(ctx context.Context, data *AuditnslogglobalAuditnslogpolicyBindingResourceModel) audit.Auditnslogglobalauditnslogpolicybinding {
	tflog.Debug(ctx, "In auditnslogglobal_auditnslogpolicy_bindingGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	auditnslogglobal_auditnslogpolicy_binding := audit.Auditnslogglobalauditnslogpolicybinding{}
	if !data.Globalbindtype.IsNull() {
		auditnslogglobal_auditnslogpolicy_binding.Globalbindtype = data.Globalbindtype.ValueString()
	}
	if !data.Policyname.IsNull() {
		auditnslogglobal_auditnslogpolicy_binding.Policyname = data.Policyname.ValueString()
	}
	if !data.Priority.IsNull() {
		auditnslogglobal_auditnslogpolicy_binding.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return auditnslogglobal_auditnslogpolicy_binding
}

func auditnslogglobal_auditnslogpolicy_bindingSetAttrFromGet(ctx context.Context, data *AuditnslogglobalAuditnslogpolicyBindingResourceModel, getResponseData map[string]interface{}) *AuditnslogglobalAuditnslogpolicyBindingResourceModel {
	tflog.Debug(ctx, "In auditnslogglobal_auditnslogpolicy_bindingSetAttrFromGet Function")

	// Convert API response to model
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
	// Case 3: Multiple unique attributes - comma-separated key:base64(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("globalbindtype:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Globalbindtype.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("policyname:%s", utils.EncodeToBase64(fmt.Sprintf("%v", data.Policyname.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
