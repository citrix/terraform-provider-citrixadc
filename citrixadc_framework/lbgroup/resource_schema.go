package lbgroup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/lb"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// LbgroupResourceModel describes the resource data model.
type LbgroupResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Backuppersistencetimeout types.Int64  `tfsdk:"backuppersistencetimeout"`
	Cookiedomain             types.String `tfsdk:"cookiedomain"`
	Cookiename               types.String `tfsdk:"cookiename"`
	Mastervserver            types.String `tfsdk:"mastervserver"`
	Name                     types.String `tfsdk:"name"`
	Newname                  types.String `tfsdk:"newname"`
	Persistencebackup        types.String `tfsdk:"persistencebackup"`
	Persistencetype          types.String `tfsdk:"persistencetype"`
	Persistmask              types.String `tfsdk:"persistmask"`
	Rule                     types.String `tfsdk:"rule"`
	Timeout                  types.Int64  `tfsdk:"timeout"`
	Usevserverpersistency    types.String `tfsdk:"usevserverpersistency"`
	V6persistmasklen         types.Int64  `tfsdk:"v6persistmasklen"`
}

func (r *LbgroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the lbgroup resource.",
			},
			"backuppersistencetimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Time period, in minutes, for which backup persistence is in effect.",
			},
			"cookiedomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Domain attribute for the HTTP cookie.",
			},
			"cookiename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use this parameter to specify the cookie name for COOKIE peristence type. It specifies the name of cookie with a maximum of 32 characters. If not specified, cookie name is internally generated.",
			},
			"mastervserver": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When USE_VSERVER_PERSISTENCE is enabled, one can use this setting to designate a member vserver as master which is responsible to create the persistence sessions",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the load balancing virtual server group.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the load balancing virtual server group.",
			},
			"persistencebackup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of backup persistence for the group.",
			},
			"persistencetype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of persistence for the group. Available settings function as follows:\n* SOURCEIP - Create persistence sessions based on the client IP.\n* COOKIEINSERT - Create persistence sessions based on a cookie in client requests. The cookie is inserted by a Set-Cookie directive from the server, in its first response to a client.\n* RULE - Create persistence sessions based on a user defined rule.\n* NONE - Disable persistence for the group.",
			},
			"persistmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Persistence mask to apply to source IPv4 addresses when creating source IP based persistence sessions.",
			},
			"rule": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("None"),
				Description: "Expression, or name of a named expression, against which traffic is evaluated.\n\nThe following requirements apply only to the Citrix ADC CLI:\n* If the expression includes one or more spaces, enclose the entire expression in double quotation marks.\n* If the expression itself includes double quotation marks, escape the quotations by using the \\ character.\n* Alternatively, you can use single quotation marks to enclose the rule, in which case you do not have to escape the double quotation marks.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(2),
				Description: "Time period for which a persistence session is in effect.",
			},
			"usevserverpersistency": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Use this parameter to enable vserver level persistence on group members. This allows member vservers to have their own persistence, but need to be compatible with other members persistence rules. When this setting is enabled persistence sessions created by any of the members can be shared by other member vservers.",
			},
			"v6persistmasklen": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(128),
				Description: "Persistence mask to apply to source IPv6 addresses when creating source IP based persistence sessions.",
			},
		},
	}
}

func lbgroupGetThePayloadFromtheConfig(ctx context.Context, data *LbgroupResourceModel) lb.Lbgroup {
	tflog.Debug(ctx, "In lbgroupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	lbgroup := lb.Lbgroup{}
	if !data.Backuppersistencetimeout.IsNull() {
		lbgroup.Backuppersistencetimeout = utils.IntPtr(int(data.Backuppersistencetimeout.ValueInt64()))
	}
	if !data.Cookiedomain.IsNull() {
		lbgroup.Cookiedomain = data.Cookiedomain.ValueString()
	}
	if !data.Cookiename.IsNull() {
		lbgroup.Cookiename = data.Cookiename.ValueString()
	}
	if !data.Mastervserver.IsNull() {
		lbgroup.Mastervserver = data.Mastervserver.ValueString()
	}
	if !data.Name.IsNull() {
		lbgroup.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		lbgroup.Newname = data.Newname.ValueString()
	}
	if !data.Persistencebackup.IsNull() {
		lbgroup.Persistencebackup = data.Persistencebackup.ValueString()
	}
	if !data.Persistencetype.IsNull() {
		lbgroup.Persistencetype = data.Persistencetype.ValueString()
	}
	if !data.Persistmask.IsNull() {
		lbgroup.Persistmask = data.Persistmask.ValueString()
	}
	if !data.Rule.IsNull() {
		lbgroup.Rule = data.Rule.ValueString()
	}
	if !data.Timeout.IsNull() {
		lbgroup.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Usevserverpersistency.IsNull() {
		lbgroup.Usevserverpersistency = data.Usevserverpersistency.ValueString()
	}
	if !data.V6persistmasklen.IsNull() {
		lbgroup.V6persistmasklen = utils.IntPtr(int(data.V6persistmasklen.ValueInt64()))
	}

	return lbgroup
}

func lbgroupSetAttrFromGet(ctx context.Context, data *LbgroupResourceModel, getResponseData map[string]interface{}) *LbgroupResourceModel {
	tflog.Debug(ctx, "In lbgroupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["backuppersistencetimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Backuppersistencetimeout = types.Int64Value(intVal)
		}
	} else {
		data.Backuppersistencetimeout = types.Int64Null()
	}
	if val, ok := getResponseData["cookiedomain"]; ok && val != nil {
		data.Cookiedomain = types.StringValue(val.(string))
	} else {
		data.Cookiedomain = types.StringNull()
	}
	if val, ok := getResponseData["cookiename"]; ok && val != nil {
		data.Cookiename = types.StringValue(val.(string))
	} else {
		data.Cookiename = types.StringNull()
	}
	if val, ok := getResponseData["mastervserver"]; ok && val != nil {
		data.Mastervserver = types.StringValue(val.(string))
	} else {
		data.Mastervserver = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["persistencebackup"]; ok && val != nil {
		data.Persistencebackup = types.StringValue(val.(string))
	} else {
		data.Persistencebackup = types.StringNull()
	}
	if val, ok := getResponseData["persistencetype"]; ok && val != nil {
		data.Persistencetype = types.StringValue(val.(string))
	} else {
		data.Persistencetype = types.StringNull()
	}
	if val, ok := getResponseData["persistmask"]; ok && val != nil {
		data.Persistmask = types.StringValue(val.(string))
	} else {
		data.Persistmask = types.StringNull()
	}
	if val, ok := getResponseData["rule"]; ok && val != nil {
		data.Rule = types.StringValue(val.(string))
	} else {
		data.Rule = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["usevserverpersistency"]; ok && val != nil {
		data.Usevserverpersistency = types.StringValue(val.(string))
	} else {
		data.Usevserverpersistency = types.StringNull()
	}
	if val, ok := getResponseData["v6persistmasklen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.V6persistmasklen = types.Int64Value(intVal)
		}
	} else {
		data.V6persistmasklen = types.Int64Null()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
