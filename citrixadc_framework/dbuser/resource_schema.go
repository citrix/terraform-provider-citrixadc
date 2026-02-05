package dbuser

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DbuserResourceModel describes the resource data model.
type DbuserResourceModel struct {
	Id       types.String `tfsdk:"id"`
	Loggedin types.Bool   `tfsdk:"loggedin"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
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

func dbuserGetThePayloadFromtheConfig(ctx context.Context, data *DbuserResourceModel) db.Dbuser {
	tflog.Debug(ctx, "In dbuserGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dbuser := db.Dbuser{}
	if !data.Loggedin.IsNull() {
		dbuser.Loggedin = data.Loggedin.ValueBool()
	}
	if !data.Password.IsNull() {
		dbuser.Password = data.Password.ValueString()
	}
	if !data.Username.IsNull() {
		dbuser.Username = data.Username.ValueString()
	}

	return dbuser
}

func dbuserSetAttrFromGet(ctx context.Context, data *DbuserResourceModel, getResponseData map[string]interface{}) *DbuserResourceModel {
	tflog.Debug(ctx, "In dbuserSetAttrFromGet Function")

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
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(data.Username.ValueString())

	return data
}
