package mcpprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/basic"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. Because these attributes revert to their NITRO
// default (or absence from GET) after unset, marking the plan unknown also avoids
// a "provider produced inconsistent result" error, which a static Default would
// trigger.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// McpprofileResourceModel describes the resource data model.
type McpprofileResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Comment                     types.String `tfsdk:"comment"`
	Hostreplacement             types.String `tfsdk:"hostreplacement"`
	Insertheaderinclientrequest types.String `tfsdk:"insertheaderinclientrequest"`
	Name                        types.String `tfsdk:"name"`
	Profiletype                 types.String `tfsdk:"profiletype"`
	Protocolversion             types.String `tfsdk:"protocolversion"`
	Proxymode                   types.String `tfsdk:"proxymode"`
	Tokenorapi                  types.String `tfsdk:"tokenorapi"`
	TokenorapiWo                types.String `tfsdk:"tokenorapi_wo"`
	TokenorapiWoVersion         types.Int64  `tfsdk:"tokenorapi_wo_version"`
	Urlreplacement              types.String `tfsdk:"urlreplacement"`
}

func (r *McpprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the mcpprofile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the mcp profile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"proxymode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Reverts to the non-empty NITRO default "FORWARD" (returned by GET)
				// when removed from config. A static Default (applied by the framework
				// whenever config is null, overriding carried-forward prior state, and
				// preserved through MarkComputedNilsAsUnknown) both triggers the unset
				// (plan "FORWARD" != prior "REVERSE") and keeps the post-unset plan
				// idempotent. unsetOnRemoveStringModifier is unusable here: it forces
				// the plan unknown on every config-omit -> perpetual spurious diff.
				Default:     stringdefault.StaticString("FORWARD"),
				Description: "Proxy mode for the MCP profile. FORWARD mode replaces Host and URL in backend requests. REVERSE mode passes requests as-is. Possible values = FORWARD, REVERSE",
			},
			"profiletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					// Create-only: profiletype is not part of the NITRO update payload,
					// so a change forces resource re-creation.
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of MCP profile. Frontend profiles apply to the entity that receives requests from a client. Backend profiles apply to the entity that sends client requests to a server. Possible values = BACKEND, FRONTEND",
			},
			"hostreplacement": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Context-dependent default: ENABLED under proxymode=FORWARD, DISABLED
				// (and immutable) under REVERSE. A static Default is therefore wrong,
				// and UseStateForUnknown would carry the stale value forward when
				// proxymode changes -> "provider produced inconsistent result". With no
				// plan modifier the framework keeps the prior value when nothing changed
				// (idempotent) and re-marks it "known after apply" when the plan differs
				// (e.g. proxymode changed), letting the appliance recompute it.
				Description: "Whether the Host header should be replaced with the backend MCP server FQDN in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED by default. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED. Possible values = ENABLED, DISABLED",
			},
			"urlreplacement": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Context-dependent default (see hostreplacement): no plan modifier so
				// the appliance recomputes it when proxymode changes.
				Description: "Whether the URL should be replaced with the backend MCP server URL in FORWARD proxy mode. If mcpProxyMode is FORWARD, this parameter is ENABLED by default. If mcpProxyMode is REVERSE, this parameter is DISABLED and cannot be ENABLED. Possible values = ENABLED, DISABLED",
			},
			"protocolversion": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					// Firmware-versioned server default (not proxymode-dependent), so
					// UseStateForUnknown safely keeps the read value stable. No static
					// Default (it varies by firmware) and no unsetOnRemove (that forced
					// a perpetual spurious "known after apply").
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "MCP protocol version to advertise during monitoring of a mcp server.",
			},
			// tokenorapi is a secret (a NITRO password_key). It is modelled as a plain
			// Optional+Sensitive attribute (the backward-compatible plaintext path)
			// alongside the write-only ephemeral path (tokenorapi_wo +
			// tokenorapi_wo_version). NITRO returns it only in an encrypted form that
			// never matches the configured plaintext, so it is never read back into
			// state (retained from config in SetAttrFromGet) and is not part of the
			// unset flow (secrets use the _wo rotation path instead).
			"tokenorapi": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "If you like to insert Bearer or API token, configure this parameter with full header.",
			},
			"tokenorapi_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "If you like to insert Bearer or API token, configure this parameter with full header. (write-only ephemeral variant of tokenorapi)",
			},
			"tokenorapi_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a tokenorapi_wo update.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					unsetOnRemoveStringModifier{},
				},
				Description: "Any information about the MCP profile.",
			},
			"insertheaderinclientrequest": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Reverts to the non-empty NITRO default "DISABLED" (returned by GET)
				// when removed from config; not proxymode-dependent. Static Default
				// triggers the unset and stays idempotent (see proxymode above).
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Whether mcp_token_or_api configuration will be used for MCP requests coming from client. Possible values = ENABLED, DISABLED",
			},
		},
	}
}

func mcpprofileGetThePayloadFromthePlan(ctx context.Context, data *McpprofileResourceModel) basic.Mcpprofile {
	tflog.Debug(ctx, "In mcpprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	mcpprofile := basic.Mcpprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		mcpprofile.Name = data.Name.ValueString()
	}
	if !data.Proxymode.IsNull() && !data.Proxymode.IsUnknown() {
		mcpprofile.Proxymode = data.Proxymode.ValueString()
	}
	if !data.Profiletype.IsNull() && !data.Profiletype.IsUnknown() {
		mcpprofile.Profiletype = data.Profiletype.ValueString()
	}
	if !data.Hostreplacement.IsNull() && !data.Hostreplacement.IsUnknown() {
		mcpprofile.Hostreplacement = data.Hostreplacement.ValueString()
	}
	if !data.Urlreplacement.IsNull() && !data.Urlreplacement.IsUnknown() {
		mcpprofile.Urlreplacement = data.Urlreplacement.ValueString()
	}
	if !data.Protocolversion.IsNull() && !data.Protocolversion.IsUnknown() {
		mcpprofile.Protocolversion = data.Protocolversion.ValueString()
	}
	if !data.Tokenorapi.IsNull() && !data.Tokenorapi.IsUnknown() {
		mcpprofile.Tokenorapi = data.Tokenorapi.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		mcpprofile.Comment = data.Comment.ValueString()
	}
	if !data.Insertheaderinclientrequest.IsNull() && !data.Insertheaderinclientrequest.IsUnknown() {
		mcpprofile.Insertheaderinclientrequest = data.Insertheaderinclientrequest.ValueString()
	}

	return mcpprofile
}

// mcpprofileGetThePayloadFromtheConfig overlays write-only attributes (which are
// nulled out of the plan) from config onto the payload. Handles the write-only
// secret tokenorapi_wo -> tokenorapi.
func mcpprofileGetThePayloadFromtheConfig(ctx context.Context, data *McpprofileResourceModel, payload *basic.Mcpprofile) {
	tflog.Debug(ctx, "In mcpprofileGetThePayloadFromtheConfig Function")

	if !data.TokenorapiWo.IsNull() {
		tokenorapiWo := data.TokenorapiWo.ValueString()
		if tokenorapiWo != "" {
			payload.Tokenorapi = tokenorapiWo
		}
	}
}

func mcpprofileSetAttrFromGet(ctx context.Context, data *McpprofileResourceModel, getResponseData map[string]interface{}) *McpprofileResourceModel {
	tflog.Debug(ctx, "In mcpprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["proxymode"]; ok && val != nil {
		data.Proxymode = types.StringValue(val.(string))
	} else {
		data.Proxymode = types.StringNull()
	}
	if val, ok := getResponseData["profiletype"]; ok && val != nil {
		data.Profiletype = types.StringValue(val.(string))
	} else {
		data.Profiletype = types.StringNull()
	}
	if val, ok := getResponseData["hostreplacement"]; ok && val != nil {
		data.Hostreplacement = types.StringValue(val.(string))
	} else {
		data.Hostreplacement = types.StringNull()
	}
	if val, ok := getResponseData["urlreplacement"]; ok && val != nil {
		data.Urlreplacement = types.StringValue(val.(string))
	} else {
		data.Urlreplacement = types.StringNull()
	}
	if val, ok := getResponseData["protocolversion"]; ok && val != nil {
		data.Protocolversion = types.StringValue(val.(string))
	} else {
		data.Protocolversion = types.StringNull()
	}
	// tokenorapi is a secret returned by NITRO only in an encrypted form that does
	// not match the plaintext supplied in configuration. When the user configured a
	// plaintext value, retain it (avoids a "provider produced inconsistent result
	// after apply"); otherwise leave it null (the value may have been supplied via
	// the write-only tokenorapi_wo path). Never overwrite with the encrypted GET
	// value.
	if !data.Tokenorapi.IsNull() && !data.Tokenorapi.IsUnknown() {
		// retain the configured plaintext value
	} else {
		data.Tokenorapi = types.StringNull()
	}
	// tokenorapi_wo / tokenorapi_wo_version are write-only/ephemeral and are never
	// returned by NITRO; retain them from config/state.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["insertheaderinclientrequest"]; ok && val != nil {
		data.Insertheaderinclientrequest = types.StringValue(val.(string))
	} else {
		data.Insertheaderinclientrequest = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
