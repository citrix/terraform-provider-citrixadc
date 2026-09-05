package smppuser

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func SmppuserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SMPP user. Must be the same as the user name specified in the SMPP server.",
			},
		},
	}
}

type SmppuserDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func smppuserDataSourceSetAttrFromGet(ctx context.Context, data *SmppuserDataSourceModel, getResponseData map[string]interface{}) *SmppuserDataSourceModel {
	tflog.Debug(ctx, "In smppuserDataSourceSetAttrFromGet Function")

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
