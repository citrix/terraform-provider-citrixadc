package aaasession

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaasessionResourceModel describes the resource data model.
// aaasession is an action-only resource (kill via ?action=kill). All attributes
// are optional kill filters. nodeid is a GET-only cluster filter (Pattern 15)
// and is therefore excluded from the kill payload (see the payload builder).
type AaasessionResourceModel struct {
	Id         types.String `tfsdk:"id"`
	All        types.Bool   `tfsdk:"all"`
	Groupname  types.String `tfsdk:"groupname"`
	Iip        types.String `tfsdk:"iip"`
	Netmask    types.String `tfsdk:"netmask"`
	Nodeid     types.Int64  `tfsdk:"nodeid"`
	Sessionkey types.String `tfsdk:"sessionkey"`
	Username   types.String `tfsdk:"username"`
}

func (r *AaasessionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaasession resource.",
			},
			// All kill arguments are optional filters. Read is a no-op (kill is an
			// action, the killed session is not a persistent managed object), so
			// these must NOT be Computed or Terraform reports an unknown value
			// after apply (Pattern 13 schema-flag implication).
			"all": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Terminate all active AAA-TM/VPN sessions.",
			},
			"groupname": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AAA group.",
			},
			"iip": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "IP address or the first address in the intranet IP range.",
			},
			"netmask": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Subnet mask for the intranet IP range.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"sessionkey": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Show aaa session associated with given session key",
			},
			"username": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the AAA user.",
			},
		},
	}
}

// aaasessionGetThePayloadFromthePlan builds the body for the kill action.
// Build a map containing ONLY the kill arguments that are set so the action body
// never includes invalid arguments. nodeid is a GET-only cluster filter and is
// intentionally excluded from the kill payload (Pattern 15).
func aaasessionGetThePayloadFromthePlan(ctx context.Context, data *AaasessionResourceModel) map[string]interface{} {
	tflog.Debug(ctx, "In aaasessionGetThePayloadFromthePlan Function")

	aaasession := map[string]interface{}{}
	if !data.All.IsNull() && !data.All.IsUnknown() {
		aaasession["all"] = data.All.ValueBool()
	}
	if !data.Groupname.IsNull() && !data.Groupname.IsUnknown() {
		aaasession["groupname"] = data.Groupname.ValueString()
	}
	if !data.Iip.IsNull() && !data.Iip.IsUnknown() {
		aaasession["iip"] = data.Iip.ValueString()
	}
	if !data.Netmask.IsNull() && !data.Netmask.IsUnknown() {
		aaasession["netmask"] = data.Netmask.ValueString()
	}
	if !data.Sessionkey.IsNull() && !data.Sessionkey.IsUnknown() {
		aaasession["sessionkey"] = data.Sessionkey.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		aaasession["username"] = data.Username.ValueString()
	}

	return aaasession
}

// aaasessionSetAttrFromGetForDatasource faithfully copies the GET (get all)
// response into the model for the read-only datasource. It exposes both the
// filter args and the session output (read-only) fields. The resource itself
// never calls this (its Read is a no-op).
func aaasessionSetAttrFromGetForDatasource(ctx context.Context, data *AaasessionResourceModel, getResponseData map[string]interface{}) *AaasessionResourceModel {
	tflog.Debug(ctx, "In aaasessionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["all"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.All = types.BoolValue(b)
		}
	}
	if val, ok := getResponseData["groupname"]; ok && val != nil {
		data.Groupname = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["iip"]; ok && val != nil {
		data.Iip = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["netmask"]; ok && val != nil {
		data.Netmask = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["sessionkey"]; ok && val != nil {
		data.Sessionkey = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	}

	return data
}
