package rewriteparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rewrite"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RewriteparamResourceModel describes the resource data model.
type RewriteparamResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Timeout     types.Int64  `tfsdk:"timeout"`
	Undefaction types.String `tfsdk:"undefaction"`
}

func (r *RewriteparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rewriteparam resource.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3900),
				Description: "Maximum time in milliseconds to allow for processing all the policies and their selected actions without interruption. If the timeout is reached then the evaluation causes an UNDEF to be raised and no further processing is performed. Note that some rewrites may have already been performed.",
			},
			"undefaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("NOREWRITE"),
				Description: "Action to perform if the result of policy evaluation is undefined (UNDEF). An UNDEF event indicates an error condition in evaluating the expression.\nAvailable settings function as follows:\n* NOREWRITE - Do not modify the message.\n* RESET - Reset the connection and notify the user's browser, so that the user can resend the request.\n* DROP - Drop the message without sending a response to the user.",
			},
		},
	}
}

func rewriteparamGetThePayloadFromtheConfig(ctx context.Context, data *RewriteparamResourceModel) rewrite.Rewriteparam {
	tflog.Debug(ctx, "In rewriteparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rewriteparam := rewrite.Rewriteparam{}
	if !data.Timeout.IsNull() {
		rewriteparam.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Undefaction.IsNull() {
		rewriteparam.Undefaction = data.Undefaction.ValueString()
	}

	return rewriteparam
}

func rewriteparamSetAttrFromGet(ctx context.Context, data *RewriteparamResourceModel, getResponseData map[string]interface{}) *RewriteparamResourceModel {
	tflog.Debug(ctx, "In rewriteparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["undefaction"]; ok && val != nil {
		data.Undefaction = types.StringValue(val.(string))
	} else {
		data.Undefaction = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("rewriteparam-config")

	return data
}
