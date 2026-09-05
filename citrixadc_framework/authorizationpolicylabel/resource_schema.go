package authorizationpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/authorization"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuthorizationpolicylabelResourceModel describes the resource data model.
type AuthorizationpolicylabelResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Labelname types.String `tfsdk:"labelname"`
	Newname   types.String `tfsdk:"newname"`
}

func (r *AuthorizationpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authorizationpolicylabel resource.",
			},
			"labelname": schema.StringAttribute{
				Required: true,
				// SDK v2 marked labelname ForceNew -> RequiresReplace. The primary key
				// cannot be changed in place (an in-place name change is expressed via
				// the separate newname / rename action, not by editing labelname).
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new authorization policy label. \nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the authorization policy is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authorization policy label\" or 'authorization policy label').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				// newname is the rename trigger (NITRO ?action=rename). Changing it
				// must NOT force replacement - it drives an in-place rename via Update.
				// Not Computed: it is a pure user input, never echoed back by GET, so
				// making it Computed would leave it unknown after apply.
				Description: "The new name of the auth policy label",
			},
		},
	}
}

func authorizationpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *AuthorizationpolicylabelResourceModel) authorization.Authorizationpolicylabel {
	tflog.Debug(ctx, "In authorizationpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authorizationpolicylabel := authorization.Authorizationpolicylabel{}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		authorizationpolicylabel.Labelname = data.Labelname.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.

	return authorizationpolicylabel
}

func authorizationpolicylabelSetAttrFromGet(ctx context.Context, data *AuthorizationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authorizationpolicylabelSetAttrFromGet Function")

	// labelname is the user-facing key. Once a rename has happened (via newname),
	// the live object name (tracked by data.Id) diverges from the configured
	// labelname, and GET returns the live (new) name. Overwriting labelname from
	// GET would clobber the user's configured value and trigger a spurious
	// RequiresReplace diff. So only adopt the GET value when we don't already have
	// one (e.g. on import, where state carries only the ID); otherwise preserve.
	if data.Labelname.IsNull() || data.Labelname.IsUnknown() || data.Labelname.ValueString() == "" {
		if val, ok := getResponseData["labelname"]; ok && val != nil {
			data.Labelname = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.

	return data
}

// authorizationpolicylabelSetAttrFromGetForDatasource faithfully copies every field
// from the GET response. The datasource has no prior plan/state to preserve, so it
// must populate the model directly from the API response and set the ID itself.
func authorizationpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *AuthorizationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthorizationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authorizationpolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
