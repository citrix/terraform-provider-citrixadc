package dbuser

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func DbuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the database user. Must be the same as the user name specified in the database.",
			},
		},
	}
}

type DbuserDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func dbuserDataSourceSetAttrFromGet(ctx context.Context, data *DbuserDataSourceModel, getResponseData map[string]interface{}) *DbuserDataSourceModel {
	tflog.Debug(ctx, "In dbuserDataSourceSetAttrFromGet Function")

	// Convert API response to model
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Username.ValueString()))

	return data
}
