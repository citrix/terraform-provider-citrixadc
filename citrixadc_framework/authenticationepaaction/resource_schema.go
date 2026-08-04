package authenticationepaaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthenticationepaactionResourceModel describes the resource data model.
type AuthenticationepaactionResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Csecexpr        types.String `tfsdk:"csecexpr"`
	Defaultepagroup types.String `tfsdk:"defaultepagroup"`
	Deletefiles     types.String `tfsdk:"deletefiles"`
	Deviceposture   types.String `tfsdk:"deviceposture"`
	Killprocess     types.String `tfsdk:"killprocess"`
	Name            types.String `tfsdk:"name"`
	Quarantinegroup types.String `tfsdk:"quarantinegroup"`
}

func (r *AuthenticationepaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationepaaction resource.",
			},
			"csecexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "it holds the ClientSecurityExpression to be sent to the client",
			},
			"defaultepagroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the EPA check succeeds.",
			},
			"deletefiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the path(s) and name(s) of the files to be deleted by the endpoint analysis (EPA) tool. Multiple files to be delimited by comma",
			},
			"deviceposture": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to enable/disable device posture service scan",
			},
			"killprocess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "String specifying the name of a process to be terminated by the endpoint analysis (EPA) tool. Multiple processes to be delimited by comma",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the epa action. Must begin with a\n	    letter, number, or the underscore character (_), and must consist\n	    only of letters, numbers, and the hyphen (-), period (.) pound\n	    (#), space ( ), at (@), equals (=), colon (:), and underscore\n		    characters. Cannot be changed after epa action is created.The following requirement applies only to the Citrix ADC CLI:If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my aaa action\" or 'my aaa action').",
			},
			"quarantinegroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the quarantine group that is chosen when the EPA check fails\nif configured.",
			},
		},
	}
}

func authenticationepaactionGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationepaactionResourceModel) authentication.Authenticationepaaction {
	tflog.Debug(ctx, "In authenticationepaactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationepaaction := authentication.Authenticationepaaction{}
	if !data.Csecexpr.IsNull() && !data.Csecexpr.IsUnknown() {
		authenticationepaaction.Csecexpr = data.Csecexpr.ValueString()
	}
	if !data.Defaultepagroup.IsNull() && !data.Defaultepagroup.IsUnknown() {
		authenticationepaaction.Defaultepagroup = data.Defaultepagroup.ValueString()
	}
	if !data.Deletefiles.IsNull() && !data.Deletefiles.IsUnknown() {
		authenticationepaaction.Deletefiles = data.Deletefiles.ValueString()
	}
	if !data.Deviceposture.IsNull() && !data.Deviceposture.IsUnknown() {
		authenticationepaaction.Deviceposture = data.Deviceposture.ValueString()
	}
	if !data.Killprocess.IsNull() && !data.Killprocess.IsUnknown() {
		authenticationepaaction.Killprocess = data.Killprocess.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationepaaction.Name = data.Name.ValueString()
	}
	if !data.Quarantinegroup.IsNull() && !data.Quarantinegroup.IsUnknown() {
		authenticationepaaction.Quarantinegroup = data.Quarantinegroup.ValueString()
	}

	return authenticationepaaction
}

func authenticationepaactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AuthenticationepaactionResourceModel) authentication.Authenticationepaaction {
	tflog.Debug(ctx, "In authenticationepaactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields
	authenticationepaaction := authentication.Authenticationepaaction{}
	if !data.Csecexpr.IsNull() && !data.Csecexpr.IsUnknown() {
		authenticationepaaction.Csecexpr = data.Csecexpr.ValueString()
	}
	if !data.Defaultepagroup.IsNull() && !data.Defaultepagroup.IsUnknown() {
		authenticationepaaction.Defaultepagroup = data.Defaultepagroup.ValueString()
	}
	if !data.Deletefiles.IsNull() && !data.Deletefiles.IsUnknown() {
		authenticationepaaction.Deletefiles = data.Deletefiles.ValueString()
	}
	if !data.Deviceposture.IsNull() && !data.Deviceposture.IsUnknown() {
		authenticationepaaction.Deviceposture = data.Deviceposture.ValueString()
	}
	if !data.Killprocess.IsNull() && !data.Killprocess.IsUnknown() {
		authenticationepaaction.Killprocess = data.Killprocess.ValueString()
	}
	// name is the primary key and is required in the update payload
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationepaaction.Name = data.Name.ValueString()
	}
	if !data.Quarantinegroup.IsNull() && !data.Quarantinegroup.IsUnknown() {
		authenticationepaaction.Quarantinegroup = data.Quarantinegroup.ValueString()
	}

	return authenticationepaaction
}

func authenticationepaactionSetAttrFromGet(ctx context.Context, data *AuthenticationepaactionResourceModel, getResponseData map[string]interface{}) *AuthenticationepaactionResourceModel {
	tflog.Debug(ctx, "In authenticationepaactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["csecexpr"]; ok && val != nil {
		data.Csecexpr = types.StringValue(val.(string))
	} else {
		data.Csecexpr = types.StringNull()
	}
	if val, ok := getResponseData["defaultepagroup"]; ok && val != nil {
		data.Defaultepagroup = types.StringValue(val.(string))
	} else {
		data.Defaultepagroup = types.StringNull()
	}
	if val, ok := getResponseData["deletefiles"]; ok && val != nil {
		data.Deletefiles = types.StringValue(val.(string))
	} else {
		data.Deletefiles = types.StringNull()
	}
	if val, ok := getResponseData["deviceposture"]; ok && val != nil {
		data.Deviceposture = types.StringValue(val.(string))
	} else {
		data.Deviceposture = types.StringNull()
	}
	if val, ok := getResponseData["killprocess"]; ok && val != nil {
		data.Killprocess = types.StringValue(val.(string))
	} else {
		data.Killprocess = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["quarantinegroup"]; ok && val != nil {
		data.Quarantinegroup = types.StringValue(val.(string))
	} else {
		data.Quarantinegroup = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
