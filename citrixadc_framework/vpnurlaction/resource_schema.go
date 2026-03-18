package vpnurlaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/vpn"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// VpnurlactionResourceModel describes the resource data model.
type VpnurlactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Actualurl        types.String `tfsdk:"actualurl"`
	Applicationtype  types.String `tfsdk:"applicationtype"`
	Clientlessaccess types.String `tfsdk:"clientlessaccess"`
	Comment          types.String `tfsdk:"comment"`
	Iconurl          types.String `tfsdk:"iconurl"`
	Linkname         types.String `tfsdk:"linkname"`
	Name             types.String `tfsdk:"name"`
	Newname          types.String `tfsdk:"newname"`
	Samlssoprofile   types.String `tfsdk:"samlssoprofile"`
	Ssotype          types.String `tfsdk:"ssotype"`
	Vservername      types.String `tfsdk:"vservername"`
}

func (r *VpnurlactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the vpnurlaction resource.",
			},
			"actualurl": schema.StringAttribute{
				Required:    true,
				Description: "Web address for the bookmark link.",
			},
			"applicationtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The type of application this VPN URL represents. Possible values are CVPN/SaaS/VPN",
			},
			"clientlessaccess": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If clientless access to the resource hosting the link is allowed, also use clientless access for the bookmarked web address in the Secure Client Access based session. Allows single sign-on and other HTTP processing on NetScaler Gateway for HTTPS resources.",
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the bookmark link.",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the vpn urlAction.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the NetScaler CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my vpnurl action\" or 'my vpnurl action').",
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
			"vservername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the associated vserver to handle selfAuth SSO",
			},
		},
	}
}

func vpnurlactionGetThePayloadFromtheConfig(ctx context.Context, data *VpnurlactionResourceModel) vpn.Vpnurlaction {
	tflog.Debug(ctx, "In vpnurlactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	vpnurlaction := vpn.Vpnurlaction{}
	if !data.Actualurl.IsNull() {
		vpnurlaction.Actualurl = data.Actualurl.ValueString()
	}
	if !data.Applicationtype.IsNull() {
		vpnurlaction.Applicationtype = data.Applicationtype.ValueString()
	}
	if !data.Clientlessaccess.IsNull() {
		vpnurlaction.Clientlessaccess = data.Clientlessaccess.ValueString()
	}
	if !data.Comment.IsNull() {
		vpnurlaction.Comment = data.Comment.ValueString()
	}
	if !data.Iconurl.IsNull() {
		vpnurlaction.Iconurl = data.Iconurl.ValueString()
	}
	if !data.Linkname.IsNull() {
		vpnurlaction.Linkname = data.Linkname.ValueString()
	}
	if !data.Name.IsNull() {
		vpnurlaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		vpnurlaction.Newname = data.Newname.ValueString()
	}
	if !data.Samlssoprofile.IsNull() {
		vpnurlaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Ssotype.IsNull() {
		vpnurlaction.Ssotype = data.Ssotype.ValueString()
	}
	if !data.Vservername.IsNull() {
		vpnurlaction.Vservername = data.Vservername.ValueString()
	}

	return vpnurlaction
}

func vpnurlactionSetAttrFromGet(ctx context.Context, data *VpnurlactionResourceModel, getResponseData map[string]interface{}) *VpnurlactionResourceModel {
	tflog.Debug(ctx, "In vpnurlactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["actualurl"]; ok && val != nil {
		data.Actualurl = types.StringValue(val.(string))
	} else {
		data.Actualurl = types.StringNull()
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
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else {
		data.Vservername = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
