package sslwrapkey

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslwrapkeyResourceModel describes the resource data model.
type SslwrapkeyResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Password          types.String `tfsdk:"password"`
	PasswordWo        types.String `tfsdk:"password_wo"`
	PasswordWoVersion types.Int64  `tfsdk:"password_wo_version"`
	Salt              types.String `tfsdk:"salt"`
	SaltWo            types.String `tfsdk:"salt_wo"`
	SaltWoVersion     types.Int64  `tfsdk:"salt_wo_version"`
	Wrapkeyname       types.String `tfsdk:"wrapkeyname"`
}

func (r *SslwrapkeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslwrapkey resource.",
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password string for the wrap key.",
			},
			"password_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password string for the wrap key.",
			},
			"password_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a password_wo update.",
			},
			"salt": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Salt string for the wrap key.",
			},
			"salt_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Salt string for the wrap key.",
			},
			"salt_wo_version": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a salt_wo update.",
			},
			"wrapkeyname": schema.StringAttribute{
				// Key attribute - mandatory (Pattern 8: tfdata wrongly marks optional).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the wrap key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the wrap key is created.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my key\" or 'my key').",
			},
		},
	}
}

func sslwrapkeyGetThePayloadFromthePlan(ctx context.Context, data *SslwrapkeyResourceModel) ssl.Sslwrapkey {
	tflog.Debug(ctx, "In sslwrapkeyGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslwrapkey := ssl.Sslwrapkey{}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		sslwrapkey.Password = data.Password.ValueString()
	}
	// Skip write-only attribute: password_wo
	// Skip version tracker attribute: password_wo_version
	if !data.Salt.IsNull() && !data.Salt.IsUnknown() {
		sslwrapkey.Salt = data.Salt.ValueString()
	}
	// Skip write-only attribute: salt_wo
	// Skip version tracker attribute: salt_wo_version
	if !data.Wrapkeyname.IsNull() && !data.Wrapkeyname.IsUnknown() {
		sslwrapkey.Wrapkeyname = data.Wrapkeyname.ValueString()
	}

	return sslwrapkey
}

func sslwrapkeyGetThePayloadFromtheConfig(ctx context.Context, data *SslwrapkeyResourceModel, payload *ssl.Sslwrapkey) {
	tflog.Debug(ctx, "In sslwrapkeyGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: password_wo -> password
	if !data.PasswordWo.IsNull() {
		passwordWo := data.PasswordWo.ValueString()
		if passwordWo != "" {
			payload.Password = passwordWo
		}
	}
	// Handle write-only secret attribute: salt_wo -> salt
	if !data.SaltWo.IsNull() {
		saltWo := data.SaltWo.ValueString()
		if saltWo != "" {
			payload.Salt = saltWo
		}
	}
}

func sslwrapkeySetAttrFromGet(ctx context.Context, data *SslwrapkeyResourceModel, getResponseData map[string]interface{}) *SslwrapkeyResourceModel {
	tflog.Debug(ctx, "In sslwrapkeySetAttrFromGet Function")

	// Convert API response to model.
	// password and salt are write-only secrets - the GET response returns only
	// wrapkeyname, so the secrets are preserved from plan/state and never touched.
	if val, ok := getResponseData["wrapkeyname"]; ok && val != nil {
		data.Wrapkeyname = types.StringValue(val.(string))
	}
	// ID is set once in Create; do not recompute it here (Pattern 6).

	return data
}

// sslwrapkeySetAttrFromGetForDatasource faithfully copies the GET response into the
// model for the datasource path (which has no prior plan/state to preserve) and sets
// the datasource ID.
func sslwrapkeySetAttrFromGetForDatasource(ctx context.Context, data *SslwrapkeyResourceModel, getResponseData map[string]interface{}) *SslwrapkeyResourceModel {
	tflog.Debug(ctx, "In sslwrapkeySetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["wrapkeyname"]; ok && val != nil {
		data.Wrapkeyname = types.StringValue(val.(string))
	} else {
		data.Wrapkeyname = types.StringNull()
	}

	// Set ID for the datasource (no Create runs for a datasource).
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Wrapkeyname.ValueString()))

	return data
}
