package authenticationprotecteduseraction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authentication"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AuthenticationprotecteduseractionResourceModel describes the resource data model.
type AuthenticationprotecteduseractionResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Maxconcurrentusers types.Int64  `tfsdk:"maxconcurrentusers"`
	Name               types.String `tfsdk:"name"`
	Realmstr           types.String `tfsdk:"realmstr"`
}

func (r *AuthenticationprotecteduseractionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationprotecteduseraction resource.",
			},
			"maxconcurrentusers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Max number of concurrent users allowed.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the action to configure.",
			},
			"realmstr": schema.StringAttribute{
				Required:    true,
				Description: "Kerberos Realm.",
			},
		},
	}
}

func authenticationprotecteduseractionGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationprotecteduseractionResourceModel) authentication.Authenticationprotecteduseraction {
	tflog.Debug(ctx, "In authenticationprotecteduseractionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationprotecteduseraction := authentication.Authenticationprotecteduseraction{}
	if !data.Maxconcurrentusers.IsNull() && !data.Maxconcurrentusers.IsUnknown() {
		authenticationprotecteduseraction.Maxconcurrentusers = utils.IntPtr(int(data.Maxconcurrentusers.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		authenticationprotecteduseraction.Name = data.Name.ValueString()
	}
	if !data.Realmstr.IsNull() && !data.Realmstr.IsUnknown() {
		authenticationprotecteduseraction.Realmstr = data.Realmstr.ValueString()
	}

	return authenticationprotecteduseraction
}

func authenticationprotecteduseractionSetAttrFromGet(ctx context.Context, data *AuthenticationprotecteduseractionResourceModel, getResponseData map[string]interface{}) *AuthenticationprotecteduseractionResourceModel {
	tflog.Debug(ctx, "In authenticationprotecteduseractionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["maxconcurrentusers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxconcurrentusers = types.Int64Value(intVal)
		}
	} else {
		data.Maxconcurrentusers = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["realmstr"]; ok && val != nil {
		data.Realmstr = types.StringValue(val.(string))
	} else {
		data.Realmstr = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
