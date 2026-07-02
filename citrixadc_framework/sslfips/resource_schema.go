package sslfips

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslfipsResourceModel describes the resource data model.
type SslfipsResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	Fipsfw                 types.String `tfsdk:"fipsfw"`
	Hsmlabel               types.String `tfsdk:"hsmlabel"`
	Inithsm                types.String `tfsdk:"inithsm"`
	Oldsopassword          types.String `tfsdk:"oldsopassword"`
	OldsopasswordWo        types.String `tfsdk:"oldsopassword_wo"`
	OldsopasswordWoVersion types.Int64  `tfsdk:"oldsopassword_wo_version"`
	Sopassword             types.String `tfsdk:"sopassword"`
	SopasswordWo           types.String `tfsdk:"sopassword_wo"`
	SopasswordWoVersion    types.Int64  `tfsdk:"sopassword_wo_version"`
	Userpassword           types.String `tfsdk:"userpassword"`
	UserpasswordWo         types.String `tfsdk:"userpassword_wo"`
	UserpasswordWoVersion  types.Int64  `tfsdk:"userpassword_wo_version"`
}

func (r *SslfipsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfips resource.",
			},
			"fipsfw": schema.StringAttribute{
				// fipsfw is mandatory only for the firmware `change` action, not for the
				// HSM-init set (Pattern 8 note) - keep Optional, drop Computed.
				Optional:    true,
				Description: "Path to the FIPS firmware file.",
			},
			"hsmlabel": schema.StringAttribute{
				Optional:    true,
				Description: "Label to identify the Hardware Security Module (HSM).",
			},
			"inithsm": schema.StringAttribute{
				Required:    true,
				Description: "FIPS initialization level. The appliance currently supports Level-2 (FIPS 140-2).",
			},
			"oldsopassword": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Old password for the security officer.",
			},
			"oldsopassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Old password for the security officer.",
			},
			"oldsopassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a oldsopassword_wo update.",
			},
			"sopassword": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Security officer password that will be in effect after you have configured the HSM.",
			},
			"sopassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Security officer password that will be in effect after you have configured the HSM.",
			},
			"sopassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a sopassword_wo update.",
			},
			"userpassword": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "The Hardware Security Module's (HSM) User password.",
			},
			"userpassword_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "The Hardware Security Module's (HSM) User password.",
			},
			"userpassword_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a userpassword_wo update.",
			},
		},
	}
}

func sslfipsGetThePayloadFromthePlan(ctx context.Context, data *SslfipsResourceModel) ssl.Sslfips {
	tflog.Debug(ctx, "In sslfipsGetThePayloadFromthePlan Function")

	// Create API request body from the model
	sslfips := ssl.Sslfips{}
	if !data.Fipsfw.IsNull() && !data.Fipsfw.IsUnknown() {
		sslfips.Fipsfw = data.Fipsfw.ValueString()
	}
	if !data.Hsmlabel.IsNull() && !data.Hsmlabel.IsUnknown() {
		sslfips.Hsmlabel = data.Hsmlabel.ValueString()
	}
	if !data.Inithsm.IsNull() && !data.Inithsm.IsUnknown() {
		sslfips.Inithsm = data.Inithsm.ValueString()
	}
	if !data.Oldsopassword.IsNull() && !data.Oldsopassword.IsUnknown() {
		sslfips.Oldsopassword = data.Oldsopassword.ValueString()
	}
	// Skip write-only attribute: oldsopassword_wo
	// Skip version tracker attribute: oldsopassword_wo_version
	if !data.Sopassword.IsNull() && !data.Sopassword.IsUnknown() {
		sslfips.Sopassword = data.Sopassword.ValueString()
	}
	// Skip write-only attribute: sopassword_wo
	// Skip version tracker attribute: sopassword_wo_version
	if !data.Userpassword.IsNull() && !data.Userpassword.IsUnknown() {
		sslfips.Userpassword = data.Userpassword.ValueString()
	}
	// Skip write-only attribute: userpassword_wo
	// Skip version tracker attribute: userpassword_wo_version

	return sslfips
}

func sslfipsGetThePayloadFromtheConfig(ctx context.Context, data *SslfipsResourceModel, payload *ssl.Sslfips) {
	tflog.Debug(ctx, "In sslfipsGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: oldsopassword_wo -> oldsopassword
	if !data.OldsopasswordWo.IsNull() {
		oldsopasswordWo := data.OldsopasswordWo.ValueString()
		if oldsopasswordWo != "" {
			payload.Oldsopassword = oldsopasswordWo
		}
	}
	// Handle write-only secret attribute: sopassword_wo -> sopassword
	if !data.SopasswordWo.IsNull() {
		sopasswordWo := data.SopasswordWo.ValueString()
		if sopasswordWo != "" {
			payload.Sopassword = sopasswordWo
		}
	}
	// Handle write-only secret attribute: userpassword_wo -> userpassword
	if !data.UserpasswordWo.IsNull() {
		userpasswordWo := data.UserpasswordWo.ValueString()
		if userpasswordWo != "" {
			payload.Userpassword = userpasswordWo
		}
	}
}

func sslfipsSetAttrFromGet(ctx context.Context, data *SslfipsResourceModel, getResponseData map[string]interface{}) *SslfipsResourceModel {
	tflog.Debug(ctx, "In sslfipsSetAttrFromGet Function")

	// Convert API response to model
	// fipsfw is a write-only firmware path argument; the GET response does not echo
	// it back. Preserve the existing plan/state value to avoid a perpetual diff (Pattern 7).
	if val, ok := getResponseData["fipsfw"]; ok && val != nil {
		data.Fipsfw = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["hsmlabel"]; ok && val != nil {
		data.Hsmlabel = types.StringValue(val.(string))
	} else {
		data.Hsmlabel = types.StringNull()
	}
	if val, ok := getResponseData["inithsm"]; ok && val != nil {
		data.Inithsm = types.StringValue(val.(string))
	} else {
		data.Inithsm = types.StringNull()
	}
	// oldsopassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// oldsopassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// oldsopassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	// sopassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// sopassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// sopassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config
	// userpassword is not returned by NITRO API (secret/ephemeral) - retain from config
	// userpassword_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// userpassword_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("sslfips-config")

	return data
}
