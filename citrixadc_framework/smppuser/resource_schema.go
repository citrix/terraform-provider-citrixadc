package smppuser

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/smpp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SmppuserResourceModel describes the resource data model.
type SmppuserResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func (r *SmppuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the smppuser resource.",
			},
			"password": schema.StringAttribute{
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

func smppuserGetThePayloadFromtheConfig(ctx context.Context, data *SmppuserResourceModel) smpp.Smppuser {
	tflog.Debug(ctx, "In smppuserGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	smppuser := smpp.Smppuser{}
	if !data.Password.IsNull() {
		smppuser.Password = data.Password.ValueString()
	}
	if !data.Username.IsNull() {
		smppuser.Username = data.Username.ValueString()
	}

	return smppuser
}

func smppuserSetAttrFromGet(ctx context.Context, data *SmppuserResourceModel, getResponseData map[string]interface{}) *SmppuserResourceModel {
	tflog.Debug(ctx, "In smppuserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Username.ValueString())

	return data
}
