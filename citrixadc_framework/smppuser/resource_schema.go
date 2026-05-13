package smppuser

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/smpp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SmppuserResourceModel describes the resource data model.
type SmppuserResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Username          types.String `tfsdk:"username"`
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
				Sensitive:   true,
				Description: "Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password for binding to the SMPP server. Must be the same as the password specified in the SMPP server.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"username": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the SMPP user. Must be the same as the user name specified in the SMPP server.",
			},
		},
	}
}

func smppuserGetThePayloadFromthePlan(ctx context.Context, data *SmppuserResourceModel) smpp.Smppuser {
	tflog.Debug(ctx, "In smppuserGetThePayloadFromthePlan Function")

	// Create API request body from the model
	smppuser := smpp.Smppuser{}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		smppuser.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		smppuser.Username = data.Username.ValueString()
	}

	return smppuser
}

func smppuserGetThePayloadFromtheConfig(ctx context.Context, data *SmppuserResourceModel, payload *smpp.Smppuser) {
	tflog.Debug(ctx, "In smppuserGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func smppuserSetAttrFromGet(ctx context.Context, data *SmppuserResourceModel, getResponseData map[string]interface{}) *SmppuserResourceModel {
	tflog.Debug(ctx, "In smppuserSetAttrFromGet Function")

	// Convert API response to model
	// password is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// password_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
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
