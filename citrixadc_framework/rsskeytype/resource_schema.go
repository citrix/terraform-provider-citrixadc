package rsskeytype

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/network"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// RsskeytypeResourceModel describes the resource data model.
type RsskeytypeResourceModel struct {
	Id      types.String `tfsdk:"id"`
	Rsstype types.String `tfsdk:"rsstype"`
}

func (r *RsskeytypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rsskeytype resource.",
			},
			"rsstype": schema.StringAttribute{
				Required:    true,
				Default:     stringdefault.StaticString("ASYMMETRIC"),
				Description: "Type of RSS key, possible values are SYMMETRIC and ASYMMETRIC.",
			},
		},
	}
}

func rsskeytypeGetThePayloadFromtheConfig(ctx context.Context, data *RsskeytypeResourceModel) network.Rsskeytype {
	tflog.Debug(ctx, "In rsskeytypeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rsskeytype := network.Rsskeytype{}
	if !data.Rsstype.IsNull() {
		rsskeytype.Rsstype = data.Rsstype.ValueString()
	}

	return rsskeytype
}

func rsskeytypeSetAttrFromGet(ctx context.Context, data *RsskeytypeResourceModel, getResponseData map[string]interface{}) *RsskeytypeResourceModel {
	tflog.Debug(ctx, "In rsskeytypeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["rsstype"]; ok && val != nil {
		data.Rsstype = types.StringValue(val.(string))
	} else {
		data.Rsstype = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("rsskeytype-config")

	return data
}
