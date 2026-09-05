package aaacertparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// AaacertparamsResourceModel describes the resource data model.
type AaacertparamsResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupnamefield             types.String `tfsdk:"groupnamefield"`
	Usernamefield              types.String `tfsdk:"usernamefield"`
}

func (r *AaacertparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaacertparams resource.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Client certificate field that specifies the group, in the format <field>:<subfield>.",
			},
			"usernamefield": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Client certificate field that contains the username, in the format <field>:<subfield>.",
			},
		},
	}
}

func aaacertparamsGetThePayloadFromtheConfig(ctx context.Context, data *AaacertparamsResourceModel) aaa.Aaacertparams {
	tflog.Debug(ctx, "In aaacertparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaacertparams := aaa.Aaacertparams{}
	if !data.Defaultauthenticationgroup.IsNull() && !data.Defaultauthenticationgroup.IsUnknown() {
		aaacertparams.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupnamefield.IsNull() && !data.Groupnamefield.IsUnknown() {
		aaacertparams.Groupnamefield = data.Groupnamefield.ValueString()
	}
	if !data.Usernamefield.IsNull() && !data.Usernamefield.IsUnknown() {
		aaacertparams.Usernamefield = data.Usernamefield.ValueString()
	}

	return aaacertparams
}

func aaacertparamsSetAttrFromGet(ctx context.Context, data *AaacertparamsResourceModel, getResponseData map[string]interface{}) *AaacertparamsResourceModel {
	tflog.Debug(ctx, "In aaacertparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["groupnamefield"]; ok && val != nil {
		data.Groupnamefield = types.StringValue(val.(string))
	} else {
		data.Groupnamefield = types.StringNull()
	}
	if val, ok := getResponseData["usernamefield"]; ok && val != nil {
		data.Usernamefield = types.StringValue(val.(string))
	} else {
		data.Usernamefield = types.StringNull()
	}

	// Set ID for the resource
	data.Id = types.StringValue("aaacertparams-config")

	return data
}
