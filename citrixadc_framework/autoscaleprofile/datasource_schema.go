package autoscaleprofile

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func AutoscaleprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"apikey": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "api key for authentication with service",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "AutoScale profile name.",
			},
			"sharedsecret": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "shared secret for authentication with service",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of profile.",
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL providing the service",
			},
		},
	}
}

type AutoscaleprofileDataSourceModel struct {
	Id           types.String `tfsdk:"id"`
	Apikey       types.String `tfsdk:"apikey"`
	Name         types.String `tfsdk:"name"`
	Sharedsecret types.String `tfsdk:"sharedsecret"`
	Type         types.String `tfsdk:"type"`
	Url          types.String `tfsdk:"url"`
}

func autoscaleprofileDataSourceSetAttrFromGet(ctx context.Context, data *AutoscaleprofileDataSourceModel, getResponseData map[string]interface{}) *AutoscaleprofileDataSourceModel {
	tflog.Debug(ctx, "In autoscaleprofileDataSourceSetAttrFromGet Function")

	// Convert API response to model
	// apikey is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// sharedsecret is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
