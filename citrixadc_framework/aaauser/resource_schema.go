package aaauser

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaauserResourceModel describes the resource data model.
type AaauserResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Loggedin types.Bool   `tfsdk:"loggedin"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

func (r *AaauserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaauser resource.",
			},
			"loggedin": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Show whether the user is logged in or not.",
			},
			"password": schema.StringAttribute{
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

func aaauserGetThePayloadFromtheConfig(ctx context.Context, data *AaauserResourceModel) aaa.Aaauser {
	tflog.Debug(ctx, "In aaauserGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaauser := aaa.Aaauser{}
	if !data.Loggedin.IsNull() {
		aaauser.Loggedin = data.Loggedin.ValueBool()
	}
	if !data.Password.IsNull() {
		aaauser.Password = data.Password.ValueString()
	}
	if !data.Username.IsNull() {
		aaauser.Username = data.Username.ValueString()
	}

	return aaauser
}

func aaauserSetAttrFromGet(ctx context.Context, data *AaauserResourceModel, getResponseData map[string]interface{}) *AaauserResourceModel {
	tflog.Debug(ctx, "In aaauserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["loggedin"]; ok && val != nil {
		data.Loggedin = types.BoolValue(val.(bool))
	} else {
		data.Loggedin = types.BoolNull()
	}
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
