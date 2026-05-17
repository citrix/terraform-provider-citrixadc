package dbuser

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DbuserResourceModel describes the resource data model.
type DbuserResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Loggedin          types.Bool   `tfsdk:"loggedin"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Username          types.String `tfsdk:"username"`
}

func (r *DbuserResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dbuser resource.",
			},
			"loggedin": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Display the names of all database users currently logged on to the Citrix ADC.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
			},
			"password_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Password for logging on to the database. Must be the same as the password specified in the database.",
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
				Description: "Name of the database user. Must be the same as the user name specified in the database.",
			},
		},
	}
}

func dbuserGetThePayloadFromthePlan(ctx context.Context, data *DbuserResourceModel) db.Dbuser {
	tflog.Debug(ctx, "In dbuserGetThePayloadFromthePlan Function")

	// Create API request body from the model
	dbuser := db.Dbuser{}
	if !data.Loggedin.IsNull() && !data.Loggedin.IsUnknown() {
		dbuser.Loggedin = data.Loggedin.ValueBool()
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		dbuser.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		dbuser.Username = data.Username.ValueString()
	}

	return dbuser
}

func dbuserGetThePayloadFromtheConfig(ctx context.Context, data *DbuserResourceModel, payload *db.Dbuser) {
	tflog.Debug(ctx, "In dbuserGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
}

func dbuserSetAttrFromGet(ctx context.Context, data *DbuserResourceModel, getResponseData map[string]interface{}) *DbuserResourceModel {
	tflog.Debug(ctx, "In dbuserSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["loggedin"]; ok && val != nil {
		data.Loggedin = types.BoolValue(val.(bool))
	} else {
		data.Loggedin = types.BoolNull()
	}
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
