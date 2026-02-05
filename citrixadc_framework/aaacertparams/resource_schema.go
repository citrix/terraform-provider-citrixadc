package aaacertparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

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
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupnamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that specifies the group, in the format <field>:<subfield>.",
			},
			"usernamefield": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client certificate field that contains the username, in the format <field>:<subfield>.",
			},
		},
	}
}

func aaacertparamsGetThePayloadFromtheConfig(ctx context.Context, data *AaacertparamsResourceModel) aaa.Aaacertparams {
	tflog.Debug(ctx, "In aaacertparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaacertparams := aaa.Aaacertparams{}
	if !data.Defaultauthenticationgroup.IsNull() {
		aaacertparams.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupnamefield.IsNull() {
		aaacertparams.Groupnamefield = data.Groupnamefield.ValueString()
	}
	if !data.Usernamefield.IsNull() {
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
