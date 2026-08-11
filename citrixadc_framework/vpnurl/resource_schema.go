package vpnurl

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset — mirroring the
// SDK v2 unset-on-remove contract. Without it an Optional+Computed attribute is
// "sticky": the prior value is carried forward and removal is a silent no-op.
// It intentionally does nothing when the config still carries a value, on create
// (no prior state), or when the prior value is already empty (avoids churn).
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

// VpnurlResourceModel describes the resource data model.
type VpnurlResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Actualurl        types.String `tfsdk:"actualurl"`
	Appjson          types.String `tfsdk:"appjson"`
	Applicationtype  types.String `tfsdk:"applicationtype"`
	Clientlessaccess types.String `tfsdk:"clientlessaccess"`
	Comment          types.String `tfsdk:"comment"`
	Iconurl          types.String `tfsdk:"iconurl"`
	Linkname         types.String `tfsdk:"linkname"`
	Samlssoprofile   types.String `tfsdk:"samlssoprofile"`
	Ssotype          types.String `tfsdk:"ssotype"`
	Urlname          types.String `tfsdk:"urlname"`
	Vservername      types.String `tfsdk:"vservername"`
}

func (r *VpnurlResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnurl resource.",
			},
			"actualurl": schema.StringAttribute{
				Required:    true,
				Description: "Web address for the bookmark link.",
			},
			"appjson": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "To store the template details in the json format.",
			},
			"applicationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN",
			},
			"clientlessaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("OFF"),
				Description: "If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Any comments associated with the bookmark link.",
			},
			"iconurl": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "URL to fetch icon file for displaying this resource.",
			},
			"linkname": schema.StringAttribute{
				Required:    true,
				Description: "Description of the bookmark link. The description appears in the Access Interface.",
			},
			"samlssoprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Profile to be used for doing SAML SSO",
			},
			"ssotype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Single sign on type for unified gateway",
			},
			"urlname": schema.StringAttribute{
				Required: true,
				// SDK v2 marked urlname as ForceNew -> preserve as RequiresReplace for backward compatibility.
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the bookmark link.",
			},
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString(""),
				Description: "Name of the associated LB/CS vserver",
			},
		},
	}
}

func vpnurlGetThePayloadFromtheConfig(ctx context.Context, data *VpnurlResourceModel) vpn.Vpnurl {
	tflog.Debug(ctx, "In vpnurlGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnurl := vpn.Vpnurl{}
	if !data.Actualurl.IsNull() && !data.Actualurl.IsUnknown() {
		vpnurl.Actualurl = data.Actualurl.ValueString()
	}
	if !data.Appjson.IsNull() && !data.Appjson.IsUnknown() {
		vpnurl.Appjson = data.Appjson.ValueString()
	}
	if !data.Applicationtype.IsNull() && !data.Applicationtype.IsUnknown() {
		vpnurl.Applicationtype = data.Applicationtype.ValueString()
	}
	if !data.Clientlessaccess.IsNull() && !data.Clientlessaccess.IsUnknown() {
		vpnurl.Clientlessaccess = data.Clientlessaccess.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		vpnurl.Comment = data.Comment.ValueString()
	}
	if !data.Iconurl.IsNull() && !data.Iconurl.IsUnknown() {
		vpnurl.Iconurl = data.Iconurl.ValueString()
	}
	if !data.Linkname.IsNull() && !data.Linkname.IsUnknown() {
		vpnurl.Linkname = data.Linkname.ValueString()
	}
	if !data.Samlssoprofile.IsNull() && !data.Samlssoprofile.IsUnknown() {
		vpnurl.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Ssotype.IsNull() && !data.Ssotype.IsUnknown() {
		vpnurl.Ssotype = data.Ssotype.ValueString()
	}
	if !data.Urlname.IsNull() && !data.Urlname.IsUnknown() {
		vpnurl.Urlname = data.Urlname.ValueString()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		vpnurl.Vservername = data.Vservername.ValueString()
	}

	return vpnurl
}

func vpnurlSetAttrFromGet(ctx context.Context, data *VpnurlResourceModel, getResponseData map[string]interface{}) *VpnurlResourceModel {
	tflog.Debug(ctx, "In vpnurlSetAttrFromGet Function")

	// Convert API response to model.
	// Guard the else-branches so we only null a value when it is unknown (never
	// clobber a known configured value that NITRO happens to omit from GET).
	if val, ok := getResponseData["actualurl"]; ok && val != nil {
		data.Actualurl = types.StringValue(val.(string))
	} else if data.Actualurl.IsUnknown() {
		data.Actualurl = types.StringNull()
	}
	if val, ok := getResponseData["appjson"]; ok && val != nil {
		data.Appjson = types.StringValue(val.(string))
	} else if data.Appjson.IsUnknown() {
		data.Appjson = types.StringNull()
	}
	if val, ok := getResponseData["applicationtype"]; ok && val != nil {
		data.Applicationtype = types.StringValue(val.(string))
	} else if data.Applicationtype.IsUnknown() {
		data.Applicationtype = types.StringNull()
	}
	if val, ok := getResponseData["clientlessaccess"]; ok && val != nil {
		data.Clientlessaccess = types.StringValue(val.(string))
	} else if data.Clientlessaccess.IsUnknown() {
		data.Clientlessaccess = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["iconurl"]; ok && val != nil {
		data.Iconurl = types.StringValue(val.(string))
	} else if data.Iconurl.IsUnknown() {
		data.Iconurl = types.StringNull()
	}
	if val, ok := getResponseData["linkname"]; ok && val != nil {
		data.Linkname = types.StringValue(val.(string))
	} else if data.Linkname.IsUnknown() {
		data.Linkname = types.StringNull()
	}
	if val, ok := getResponseData["samlssoprofile"]; ok && val != nil {
		data.Samlssoprofile = types.StringValue(val.(string))
	} else if data.Samlssoprofile.IsUnknown() {
		data.Samlssoprofile = types.StringNull()
	}
	if val, ok := getResponseData["ssotype"]; ok && val != nil {
		data.Ssotype = types.StringValue(val.(string))
	} else if data.Ssotype.IsUnknown() {
		data.Ssotype = types.StringNull()
	}
	if val, ok := getResponseData["urlname"]; ok && val != nil {
		data.Urlname = types.StringValue(val.(string))
	} else if data.Urlname.IsUnknown() {
		data.Urlname = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else if data.Vservername.IsUnknown() {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Urlname.ValueString())

	return data
}
