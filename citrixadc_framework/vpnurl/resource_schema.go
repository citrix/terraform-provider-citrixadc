package vpnurl

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

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
				Description: "To store the template details in the json format.",
			},
			"applicationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN",
			},
			"clientlessaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on Citrix Gateway for HTTPS resources.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments associated with the bookmark link.",
			},
			"iconurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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
				Description: "Single sign on type for unified gateway",
			},
			"urlname": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bookmark link.",
			},
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the associated LB/CS vserver",
			},
		},
	}
}

func vpnurlGetThePayloadFromtheConfig(ctx context.Context, data *VpnurlResourceModel) vpn.Vpnurl {
	tflog.Debug(ctx, "In vpnurlGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnurl := vpn.Vpnurl{}
	if !data.Actualurl.IsNull() {
		vpnurl.Actualurl = data.Actualurl.ValueString()
	}
	if !data.Appjson.IsNull() {
		vpnurl.Appjson = data.Appjson.ValueString()
	}
	if !data.Applicationtype.IsNull() {
		vpnurl.Applicationtype = data.Applicationtype.ValueString()
	}
	if !data.Clientlessaccess.IsNull() {
		vpnurl.Clientlessaccess = data.Clientlessaccess.ValueString()
	}
	if !data.Comment.IsNull() {
		vpnurl.Comment = data.Comment.ValueString()
	}
	if !data.Iconurl.IsNull() {
		vpnurl.Iconurl = data.Iconurl.ValueString()
	}
	if !data.Linkname.IsNull() {
		vpnurl.Linkname = data.Linkname.ValueString()
	}
	if !data.Samlssoprofile.IsNull() {
		vpnurl.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Ssotype.IsNull() {
		vpnurl.Ssotype = data.Ssotype.ValueString()
	}
	if !data.Urlname.IsNull() {
		vpnurl.Urlname = data.Urlname.ValueString()
	}
	if !data.Vservername.IsNull() {
		vpnurl.Vservername = data.Vservername.ValueString()
	}

	return vpnurl
}

func vpnurlSetAttrFromGet(ctx context.Context, data *VpnurlResourceModel, getResponseData map[string]interface{}) *VpnurlResourceModel {
	tflog.Debug(ctx, "In vpnurlSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actualurl"]; ok && val != nil {
		data.Actualurl = types.StringValue(val.(string))
	} else {
		data.Actualurl = types.StringNull()
	}
	if val, ok := getResponseData["appjson"]; ok && val != nil {
		data.Appjson = types.StringValue(val.(string))
	} else {
		data.Appjson = types.StringNull()
	}
	if val, ok := getResponseData["applicationtype"]; ok && val != nil {
		data.Applicationtype = types.StringValue(val.(string))
	} else {
		data.Applicationtype = types.StringNull()
	}
	if val, ok := getResponseData["clientlessaccess"]; ok && val != nil {
		data.Clientlessaccess = types.StringValue(val.(string))
	} else {
		data.Clientlessaccess = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["iconurl"]; ok && val != nil {
		data.Iconurl = types.StringValue(val.(string))
	} else {
		data.Iconurl = types.StringNull()
	}
	if val, ok := getResponseData["linkname"]; ok && val != nil {
		data.Linkname = types.StringValue(val.(string))
	} else {
		data.Linkname = types.StringNull()
	}
	if val, ok := getResponseData["samlssoprofile"]; ok && val != nil {
		data.Samlssoprofile = types.StringValue(val.(string))
	} else {
		data.Samlssoprofile = types.StringNull()
	}
	if val, ok := getResponseData["ssotype"]; ok && val != nil {
		data.Ssotype = types.StringValue(val.(string))
	} else {
		data.Ssotype = types.StringNull()
	}
	if val, ok := getResponseData["urlname"]; ok && val != nil {
		data.Urlname = types.StringValue(val.(string))
	} else {
		data.Urlname = types.StringNull()
	}
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Urlname.ValueString())

	return data
}
