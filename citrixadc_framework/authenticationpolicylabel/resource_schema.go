package authenticationpolicylabel

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
)

// AuthenticationpolicylabelResourceModel describes the resource data model.
type AuthenticationpolicylabelResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Comment     types.String `tfsdk:"comment"`
	Labelname   types.String `tfsdk:"labelname"`
	Loginschema types.String `tfsdk:"loginschema"`
	Newname     types.String `tfsdk:"newname"`
	Type        types.String `tfsdk:"type"`
}

func (r *AuthenticationpolicylabelResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the authenticationpolicylabel resource.",
			},
			"comment": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed + ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only attr; preserve state on unknown, replace only when user-configured value changes.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Any comments to preserve information about this authentication policy label.",
			},
			"labelname": schema.StringAttribute{
				// SDK v2 parity: Required + ForceNew (primary key).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the new authentication policy label.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my authentication policy label\" or 'authentication policy label').",
			},
			"loginschema": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed + ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: updatable attr; preserve state on unknown, no replace needed.
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Login schema associated with authentication policy label. Login schema defines the UI rendering by providing customization option of the fields. If user intervention is not needed for a given factor such as group extraction, a loginSchema whose authentication schema is \"noschema\" should be used.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). Changing it must
				// NOT force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
				Description: "The new name of the auth policy label.",
			},
			"type": schema.StringAttribute{
				// SDK v2 parity: Optional + Computed + ForceNew. The NITRO server
				// applies the default (AAATM_REQ) and echoes it in GET, so it is Computed
				// rather than carrying a schema Default (which would panic without Computed
				// and would not match the SDK v2 server-default behavior).
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// GH #1436: create-only attr; preserve state on unknown, replace only when user-configured value changes.
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of feature (aaatm or rba) against which to match the policies bound to this policy label.",
			},
		},
	}
}

func authenticationpolicylabelGetThePayloadFromthePlan(ctx context.Context, data *AuthenticationpolicylabelResourceModel) authentication.Authenticationpolicylabel {
	tflog.Debug(ctx, "In authenticationpolicylabelGetThePayloadFromthePlan Function")

	// Create API request body from the model
	authenticationpolicylabel := authentication.Authenticationpolicylabel{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		authenticationpolicylabel.Comment = data.Comment.ValueString()
	}
	if !data.Labelname.IsNull() && !data.Labelname.IsUnknown() {
		authenticationpolicylabel.Labelname = data.Labelname.ValueString()
	}
	if !data.Loginschema.IsNull() && !data.Loginschema.IsUnknown() {
		authenticationpolicylabel.Loginschema = data.Loginschema.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add payload, so it is deliberately excluded from the create POST body.
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		authenticationpolicylabel.Type = data.Type.ValueString()
	}

	return authenticationpolicylabel
}

func authenticationpolicylabelSetAttrFromGet(ctx context.Context, data *AuthenticationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authenticationpolicylabelSetAttrFromGet Function")

	// Convert API response to model.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
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
	if val, ok := getResponseData["loginschema"]; ok && val != nil {
		data.Loginschema = types.StringValue(val.(string))
	} else {
		data.Loginschema = types.StringNull()
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	return data
}

// authenticationpolicylabelSetAttrFromGetForDatasource faithfully copies every
// field from the GET response. The datasource has no prior plan/state to preserve,
// so it must populate the model directly from the API response and set the ID itself.
func authenticationpolicylabelSetAttrFromGetForDatasource(ctx context.Context, data *AuthenticationpolicylabelResourceModel, getResponseData map[string]interface{}) *AuthenticationpolicylabelResourceModel {
	tflog.Debug(ctx, "In authenticationpolicylabelSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["labelname"]; ok && val != nil {
		data.Labelname = types.StringValue(val.(string))
	} else {
		data.Labelname = types.StringNull()
	}
	if val, ok := getResponseData["loginschema"]; ok && val != nil {
		data.Loginschema = types.StringValue(val.(string))
	} else {
		data.Loginschema = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Labelname.ValueString()))

	return data
}
