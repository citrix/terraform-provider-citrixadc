package vpnurlaction

import (
	"context"
	"fmt"

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
				// SDK v2: Required (Computed:false)
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
				// SDK v2: Required (Computed:false)
				Required:    true,
				Description: "Description of the bookmark link. The description appears in the Access Interface.",
			},
			"name": schema.StringAttribute{
				// SDK v2: Required + ForceNew -> Required + RequiresReplace. This is
				// the resource key (is_get_id/is_delete_id/x-unique-attr) and drives
				// the single_unique ID (plain value).
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the bookmark link.",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It must NOT
				// force replacement - it drives an in-place rename via Update. Not
				// Computed: it is a pure user input, never echoed back by GET.
				Optional:    true,
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

func vpnurlactionGetThePayloadFromthePlan(ctx context.Context, data *VpnurlactionResourceModel) vpn.Vpnurlaction {
	tflog.Debug(ctx, "In vpnurlactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	vpnurlaction := vpn.Vpnurlaction{}
	if !data.Actualurl.IsNull() && !data.Actualurl.IsUnknown() {
		vpnurlaction.Actualurl = data.Actualurl.ValueString()
	}
	if !data.Applicationtype.IsNull() && !data.Applicationtype.IsUnknown() {
		vpnurlaction.Applicationtype = data.Applicationtype.ValueString()
	}
	if !data.Clientlessaccess.IsNull() && !data.Clientlessaccess.IsUnknown() {
		vpnurlaction.Clientlessaccess = data.Clientlessaccess.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		vpnurlaction.Comment = data.Comment.ValueString()
	}
	if !data.Iconurl.IsNull() && !data.Iconurl.IsUnknown() {
		vpnurlaction.Iconurl = data.Iconurl.ValueString()
	}
	if !data.Linkname.IsNull() && !data.Linkname.IsUnknown() {
		vpnurlaction.Linkname = data.Linkname.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		vpnurlaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of
	// the add/update payload, so it is deliberately excluded here.
	if !data.Samlssoprofile.IsNull() && !data.Samlssoprofile.IsUnknown() {
		vpnurlaction.Samlssoprofile = data.Samlssoprofile.ValueString()
	}
	if !data.Ssotype.IsNull() && !data.Ssotype.IsUnknown() {
		vpnurlaction.Ssotype = data.Ssotype.ValueString()
	}
	if !data.Vservername.IsNull() && !data.Vservername.IsUnknown() {
		vpnurlaction.Vservername = data.Vservername.ValueString()
	}

	return vpnurlaction
}

// vpnurlactionSetAttrFromGet populates the RESOURCE model from a GET response.
// It preserves the user-facing key (name) and rename-only newname, and guards
// against the omit-on-default trap: when NITRO omits an attribute from GET, a
// known/configured value is preserved (only an unknown Computed value is nulled).
func vpnurlactionSetAttrFromGet(ctx context.Context, data *VpnurlactionResourceModel, getResponseData map[string]interface{}) *VpnurlactionResourceModel {
	tflog.Debug(ctx, "In vpnurlactionSetAttrFromGet Function")

	if val, ok := getResponseData["actualurl"]; ok && val != nil {
		data.Actualurl = types.StringValue(val.(string))
	} else if data.Actualurl.IsUnknown() {
		data.Actualurl = types.StringNull()
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
	// name is the user-facing key. After a rename (via newname) the live object
	// name (tracked by data.Id) diverges from the configured name, and GET returns
	// the live (new) name. Overwriting name from GET would clobber the user's
	// configured value and trigger a spurious RequiresReplace diff. Only adopt the
	// GET value when we don't already have one (e.g. on import, where state carries
	// only the ID); otherwise preserve the configured value.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
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
	if val, ok := getResponseData["vservername"]; ok && val != nil {
		data.Vservername = types.StringValue(val.(string))
	} else if data.Vservername.IsUnknown() {
		data.Vservername = types.StringNull()
	}

	// NOTE: data.Id is managed by the CRUD functions (Create sets it to name,
	// Update's rename branch sets it to newname). It is deliberately NOT reset here
	// so a rename is not undone on the read-back.

	return data
}

// vpnurlactionSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it must
// populate the model directly from the API response and set the ID itself.
func vpnurlactionSetAttrFromGetForDatasource(ctx context.Context, data *VpnurlactionResourceModel, getResponseData map[string]interface{}) *VpnurlactionResourceModel {
	tflog.Debug(ctx, "In vpnurlactionSetAttrFromGetForDatasource Function")

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
	// newname is rename-only and never echoed by GET.
	data.Newname = types.StringNull()
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

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
