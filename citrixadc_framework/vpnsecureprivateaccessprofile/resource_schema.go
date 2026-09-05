package vpnsecureprivateaccessprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

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
// and removal is a silent no-op. Because these attributes revert to no value
// (absent from GET) after unset, marking the plan unknown also avoids a
// "provider produced inconsistent result" error, which a static Default would
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

// VpnsecureprivateaccessprofileResourceModel describes the resource data model.
type VpnsecureprivateaccessprofileResourceModel struct {
	Id                          types.String `tfsdk:"id"`
	Name                        types.String `tfsdk:"name"`
	Url                         types.String `tfsdk:"url"`
	Customerid                  types.String `tfsdk:"customerid"`
	Chromeenterprisepremiummode types.String `tfsdk:"chromeenterprisepremiummode"`
	Googlecustomerid            types.String `tfsdk:"googlecustomerid"`
	Googlesecuritygatewayid     types.String `tfsdk:"googlesecuritygatewayid"`
	Forceclienttype             types.String `tfsdk:"forceclienttype"`
	Sharedsecret                types.String `tfsdk:"sharedsecret"`
	SharedsecretWo              types.String `tfsdk:"sharedsecret_wo"`
	SharedsecretWoVersion       types.Int64  `tfsdk:"sharedsecret_wo_version"`
}

func (r *VpnsecureprivateaccessprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnsecureprivateaccessprofile resource.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "name of Secure Private Access profile.",
			},
			"url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Description: "Public URL for your Secure Private Access site or load balancing server.",
			},
			"customerid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					unsetOnRemoveStringModifier{},
				},
				Description: "Customer ID of the citrix cloud customer.",
			},
			"chromeenterprisepremiummode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Reverts to the non-empty NITRO default "OFF" (returned by GET) when
				// removed from config. A static Default (applied by the framework
				// whenever config is null, overriding the carried-forward prior state)
				// both triggers the unset (plan "OFF" != prior non-default value) and
				// keeps the post-unset plan idempotent (plan "OFF" == state "OFF").
				// The unsetOnRemoveStringModifier is unusable here: it forces the plan
				// unknown on every config-omit, producing a perpetual spurious diff.
				Default:     stringdefault.StaticString("OFF"),
				Description: "Secure Private Access Chrome Enterprise Premium mode of operation. Possible values = OFF, WITH_PARTNER_CONNECTOR, WITHOUT_PARTNER_CONNECTOR",
			},
			"googlecustomerid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					unsetOnRemoveStringModifier{},
				},
				Description: "Your organization's unique ID on Google's Admin console Profile settings.",
			},
			"googlesecuritygatewayid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					unsetOnRemoveStringModifier{},
				},
				Description: "The ID of the Google Secure Gateway.",
			},
			"forceclienttype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				// Reverts to the non-empty NITRO default "ON" (returned by GET) when
				// removed from config. See chromeenterprisepremiummode above for why a
				// static Default is used instead of unsetOnRemoveStringModifier.
				Default:     stringdefault.StaticString("ON"),
				Description: "Automatically configures the session for Citrix Secure Access client connectivity. Possible values = ON, OFF",
			},
			"sharedsecret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "Secure Private Access Shared Secret.",
			},
			"sharedsecret_wo": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				WriteOnly:   true,
				Description: "Secure Private Access Shared Secret.",
			},
			"sharedsecret_wo_version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "Increment this version to signal a sharedsecret_wo update.",
			},
		},
	}
}

func vpnsecureprivateaccessprofileGetThePayloadFromthePlan(ctx context.Context, data *VpnsecureprivateaccessprofileResourceModel) vpn.Vpnsecureprivateaccessprofile {
	tflog.Debug(ctx, "In vpnsecureprivateaccessprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnsecureprivateaccessprofile := vpn.Vpnsecureprivateaccessprofile{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnsecureprivateaccessprofile.Name = data.Name.ValueString()
	}
	if !data.Url.IsNull() && !data.Url.IsUnknown() {
		vpnsecureprivateaccessprofile.Url = data.Url.ValueString()
	}
	if !data.Customerid.IsNull() && !data.Customerid.IsUnknown() {
		vpnsecureprivateaccessprofile.Customerid = data.Customerid.ValueString()
	}
	if !data.Chromeenterprisepremiummode.IsNull() && !data.Chromeenterprisepremiummode.IsUnknown() {
		vpnsecureprivateaccessprofile.Chromeenterprisepremiummode = data.Chromeenterprisepremiummode.ValueString()
	}
	if !data.Googlecustomerid.IsNull() && !data.Googlecustomerid.IsUnknown() {
		vpnsecureprivateaccessprofile.Googlecustomerid = data.Googlecustomerid.ValueString()
	}
	if !data.Googlesecuritygatewayid.IsNull() && !data.Googlesecuritygatewayid.IsUnknown() {
		vpnsecureprivateaccessprofile.Googlesecuritygatewayid = data.Googlesecuritygatewayid.ValueString()
	}
	if !data.Forceclienttype.IsNull() && !data.Forceclienttype.IsUnknown() {
		vpnsecureprivateaccessprofile.Forceclienttype = data.Forceclienttype.ValueString()
	}
	if !data.Sharedsecret.IsNull() && !data.Sharedsecret.IsUnknown() {
		vpnsecureprivateaccessprofile.Sharedsecret = data.Sharedsecret.ValueString()
	}
	// Skip write-only attribute: sharedsecret_wo
	// Skip version tracker attribute: sharedsecret_wo_version

	return vpnsecureprivateaccessprofile
}

func vpnsecureprivateaccessprofileGetThePayloadFromtheConfig(ctx context.Context, data *VpnsecureprivateaccessprofileResourceModel, payload *vpn.Vpnsecureprivateaccessprofile) {
	tflog.Debug(ctx, "In vpnsecureprivateaccessprofileGetThePayloadFromtheConfig Function")

	// Add write-only attributes from config to the provided payload
	// Handle write-only secret attribute: sharedsecret_wo -> sharedsecret
	if !data.SharedsecretWo.IsNull() {
		sharedsecretWo := data.SharedsecretWo.ValueString()
		if sharedsecretWo != "" {
			payload.Sharedsecret = sharedsecretWo
		}
	}
}

func vpnsecureprivateaccessprofileSetAttrFromGet(ctx context.Context, data *VpnsecureprivateaccessprofileResourceModel, getResponseData map[string]interface{}) *VpnsecureprivateaccessprofileResourceModel {
	tflog.Debug(ctx, "In vpnsecureprivateaccessprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["url"]; ok && val != nil {
		data.Url = types.StringValue(val.(string))
	} else {
		data.Url = types.StringNull()
	}
	if val, ok := getResponseData["customerid"]; ok && val != nil {
		data.Customerid = types.StringValue(val.(string))
	} else {
		data.Customerid = types.StringNull()
	}
	if val, ok := getResponseData["chromeenterprisepremiummode"]; ok && val != nil {
		data.Chromeenterprisepremiummode = types.StringValue(val.(string))
	} else {
		data.Chromeenterprisepremiummode = types.StringNull()
	}
	if val, ok := getResponseData["googlecustomerid"]; ok && val != nil {
		data.Googlecustomerid = types.StringValue(val.(string))
	} else {
		data.Googlecustomerid = types.StringNull()
	}
	if val, ok := getResponseData["googlesecuritygatewayid"]; ok && val != nil {
		data.Googlesecuritygatewayid = types.StringValue(val.(string))
	} else {
		data.Googlesecuritygatewayid = types.StringNull()
	}
	if val, ok := getResponseData["forceclienttype"]; ok && val != nil {
		data.Forceclienttype = types.StringValue(val.(string))
	} else {
		data.Forceclienttype = types.StringNull()
	}
	// sharedsecret is not returned by NITRO API in usable form (secret/ephemeral) - retain from config
	// sharedsecret_wo is not returned by NITRO API (secret/ephemeral) - retain from config
	// sharedsecret_wo_version is not returned by NITRO API (secret/ephemeral) - retain from config

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
