package aaauser

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func AaauserDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"password": schema.StringAttribute{
				Sensitive:   true,
				Optional:    true,
				Computed:    true,
				Description: "Password with which the user logs on. Required for any user account that does not exist on an external authentication server.\nIf you are not using an external authentication server, all user accounts must have a password. If you are using an external authentication server, you must provide a password for local user accounts that do not exist on the authentication server.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the user is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or\nsingle quotation marks (for example, \"my aaa user\" or \"my aaa user\").",
			},
		},
	}
}

type AaauserDataSourceModel struct {
	Id       types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func aaauserDataSourceSetAttrFromGet(ctx context.Context, data *AaauserDataSourceModel, getResponseData map[string]interface{}) *AaauserDataSourceModel {
	tflog.Debug(ctx, "In aaauserDataSourceSetAttrFromGet Function")

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
