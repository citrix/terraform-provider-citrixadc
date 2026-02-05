package dbdbprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/db"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// DbdbprofileResourceModel describes the resource data model.
type DbdbprofileResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Conmultiplex           types.String `tfsdk:"conmultiplex"`
	Enablecachingconmuxoff types.String `tfsdk:"enablecachingconmuxoff"`
	Interpretquery         types.String `tfsdk:"interpretquery"`
	Kcdaccount             types.String `tfsdk:"kcdaccount"`
	Name                   types.String `tfsdk:"name"`
	Stickiness             types.String `tfsdk:"stickiness"`
}

func (r *DbdbprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the dbdbprofile resource.",
			},
			"conmultiplex": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Use the same server-side connection for multiple client-side requests. Default is enabled.",
			},
			"enablecachingconmuxoff": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable caching when connection multiplexing is OFF.",
			},
			"interpretquery": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("True"),
				Description: "If ENABLED, inspect the query and update the connection information, if required. If DISABLED, forward the query to the server.",
			},
			"kcdaccount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the KCD account that is used for Windows authentication.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the database profile. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at sign (@), equal sign (=), and hyphen (-) characters. Cannot be changed after the profile is created.\n	    CLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my profile\" or 'my profile').",
			},
			"stickiness": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If the queries are related to each other, forward to the same backend server.",
			},
		},
	}
}

func dbdbprofileGetThePayloadFromtheConfig(ctx context.Context, data *DbdbprofileResourceModel) db.Dbdbprofile {
	tflog.Debug(ctx, "In dbdbprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	dbdbprofile := db.Dbdbprofile{}
	if !data.Conmultiplex.IsNull() {
		dbdbprofile.Conmultiplex = data.Conmultiplex.ValueString()
	}
	if !data.Enablecachingconmuxoff.IsNull() {
		dbdbprofile.Enablecachingconmuxoff = data.Enablecachingconmuxoff.ValueString()
	}
	if !data.Interpretquery.IsNull() {
		dbdbprofile.Interpretquery = data.Interpretquery.ValueString()
	}
	if !data.Kcdaccount.IsNull() {
		dbdbprofile.Kcdaccount = data.Kcdaccount.ValueString()
	}
	if !data.Name.IsNull() {
		dbdbprofile.Name = data.Name.ValueString()
	}
	if !data.Stickiness.IsNull() {
		dbdbprofile.Stickiness = data.Stickiness.ValueString()
	}

	return dbdbprofile
}

func dbdbprofileSetAttrFromGet(ctx context.Context, data *DbdbprofileResourceModel, getResponseData map[string]interface{}) *DbdbprofileResourceModel {
	tflog.Debug(ctx, "In dbdbprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["conmultiplex"]; ok && val != nil {
		data.Conmultiplex = types.StringValue(val.(string))
	} else {
		data.Conmultiplex = types.StringNull()
	}
	if val, ok := getResponseData["enablecachingconmuxoff"]; ok && val != nil {
		data.Enablecachingconmuxoff = types.StringValue(val.(string))
	} else {
		data.Enablecachingconmuxoff = types.StringNull()
	}
	if val, ok := getResponseData["interpretquery"]; ok && val != nil {
		data.Interpretquery = types.StringValue(val.(string))
	} else {
		data.Interpretquery = types.StringNull()
	}
	if val, ok := getResponseData["kcdaccount"]; ok && val != nil {
		data.Kcdaccount = types.StringValue(val.(string))
	} else {
		data.Kcdaccount = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["stickiness"]; ok && val != nil {
		data.Stickiness = types.StringValue(val.(string))
	} else {
		data.Stickiness = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
