package nsconsoleloginprompt

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsconsoleloginpromptResourceModel describes the resource data model.
type NsconsoleloginpromptResourceModel struct {
	Id           types.String `tfsdk:"id"`
	Promptstring types.String `tfsdk:"promptstring"`
}

func (r *NsconsoleloginpromptResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsconsoleloginprompt resource.",
			},
			"promptstring": schema.StringAttribute{
				// SDK v2 backward-compat: promptstring was Required (not Computed) in the
				// legacy resource. Keep Required to preserve the existing user contract.
				Required:    true,
				Description: "Console login prompt string",
			},
		},
	}
}

func nsconsoleloginpromptGetThePayloadFromtheConfig(ctx context.Context, data *NsconsoleloginpromptResourceModel) ns.Nsconsoleloginprompt {
	tflog.Debug(ctx, "In nsconsoleloginpromptGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsconsoleloginprompt := ns.Nsconsoleloginprompt{}
	if !data.Promptstring.IsNull() && !data.Promptstring.IsUnknown() {
		nsconsoleloginprompt.Promptstring = data.Promptstring.ValueString()
	}

	return nsconsoleloginprompt
}

func nsconsoleloginpromptSetAttrFromGet(ctx context.Context, data *NsconsoleloginpromptResourceModel, getResponseData map[string]interface{}) *NsconsoleloginpromptResourceModel {
	tflog.Debug(ctx, "In nsconsoleloginpromptSetAttrFromGet Function")

	// Convert API response to model.
	// Omit-on-default guard: NITRO may omit promptstring from GET when it is at its
	// default/empty value. Do NOT clobber a known configured value in that case;
	// only null it when it is unknown (never realistic for this Required attr).
	if val, ok := getResponseData["promptstring"]; ok && val != nil {
		data.Promptstring = types.StringValue(val.(string))
	} else if data.Promptstring.IsUnknown() {
		data.Promptstring = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsconsoleloginprompt-config")

	return data
}
