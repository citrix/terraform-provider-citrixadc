package rdpclientprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/rdp"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// RdpclientprofileResourceModel describes the resource data model.
type RdpclientprofileResourceModel struct {
	Id                   types.String `tfsdk:"id"`
	Addusernameinrdpfile types.String `tfsdk:"addusernameinrdpfile"`
	Audiocapturemode     types.String `tfsdk:"audiocapturemode"`
	Keyboardhook         types.String `tfsdk:"keyboardhook"`
	Multimonitorsupport  types.String `tfsdk:"multimonitorsupport"`
	Name                 types.String `tfsdk:"name"`
	Psk                  types.String `tfsdk:"psk"`
	Randomizerdpfilename types.String `tfsdk:"randomizerdpfilename"`
	Rdpcookievalidity    types.Int64  `tfsdk:"rdpcookievalidity"`
	Rdpcustomparams      types.String `tfsdk:"rdpcustomparams"`
	Rdpfilename          types.String `tfsdk:"rdpfilename"`
	Rdphost              types.String `tfsdk:"rdphost"`
	Rdplinkattribute     types.String `tfsdk:"rdplinkattribute"`
	Rdplistener          types.String `tfsdk:"rdplistener"`
	Rdpurloverride       types.String `tfsdk:"rdpurloverride"`
	Rdpvalidateclientip  types.String `tfsdk:"rdpvalidateclientip"`
	Redirectclipboard    types.String `tfsdk:"redirectclipboard"`
	Redirectcomports     types.String `tfsdk:"redirectcomports"`
	Redirectdrives       types.String `tfsdk:"redirectdrives"`
	Redirectpnpdevices   types.String `tfsdk:"redirectpnpdevices"`
	Redirectprinters     types.String `tfsdk:"redirectprinters"`
	Videoplaybackmode    types.String `tfsdk:"videoplaybackmode"`
}

func (r *RdpclientprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the rdpclientprofile resource.",
			},
			"addusernameinrdpfile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Add username in rdp file.",
			},
			"audiocapturemode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "This setting corresponds to the selections in the Remote audio area on the Local Resources tab under Options in RDC.",
			},
			"keyboardhook": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("InFullScreenMode"),
				Description: "This setting corresponds to the selection in the Keyboard drop-down list on the Local Resources tab under Options in RDC.",
			},
			"multimonitorsupport": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "Enable/Disable Multiple Monitor Support for Remote Desktop Connection (RDC).",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "The name of the rdp profile",
			},
			"psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Pre shared key value",
			},
			"randomizerdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Will generate unique filename everytime rdp file is downloaded by appending output of time() function in the format <rdpfileName>_<time>.rdp. This tries to avoid the pop-up for replacement of existing rdp file during each rdp connection launch, hence providing better end-user experience.",
			},
			"rdpcookievalidity": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(60),
				Description: "RDP cookie validity period. RDP cookie validity time is applicable for new connection and also for any re-connection that might happen, mostly due to network disruption or during fail-over.",
			},
			"rdpcustomparams": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Option for RDP custom parameters settings (if any). Custom params needs to be separated by '&'",
			},
			"rdpfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RDP file name to be sent to End User",
			},
			"rdphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fully-qualified domain name (FQDN) of the RDP Listener.",
			},
			"rdplinkattribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix Gateway allows the configuration of rdpLinkAttribute parameter which can be used to fetch a list of RDP servers(IP/FQDN) that a user can access, from an Authentication server attribute(Example: LDAP, SAML). Based on the list received, the RDP links will be generated and displayed to the user.\n            Note: The Attribute mentioned in the rdpLinkAttribute should be fetched through corresponding authentication method.",
			},
			"rdplistener": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address (or) Fully-qualified domain name(FQDN) of the RDP Listener with the port in the format IP:Port (or) FQDN:Port",
			},
			"rdpurloverride": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "This setting determines whether the RDP parameters supplied in the vpn url override those specified in the RDP profile.",
			},
			"rdpvalidateclientip": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "This setting determines whether RDC launch is initiated by the valid client IP",
			},
			"redirectclipboard": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "This setting corresponds to the Clipboard check box on the Local Resources tab under Options in RDC.",
			},
			"redirectcomports": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "This setting corresponds to the selections for comports under More on the Local Resources tab under Options in RDC.",
			},
			"redirectdrives": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "This setting corresponds to the selections for Drives under More on the Local Resources tab under Options in RDC.",
			},
			"redirectpnpdevices": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLE"),
				Description: "This setting corresponds to the selections for pnpdevices under More on the Local Resources tab under Options in RDC.",
			},
			"redirectprinters": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "This setting corresponds to the selection in the Printers check box on the Local Resources tab under Options in RDC.",
			},
			"videoplaybackmode": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLE"),
				Description: "This setting determines if Remote Desktop Connection (RDC) will use RDP efficient multimedia streaming for video playback.",
			},
		},
	}
}

func rdpclientprofileGetThePayloadFromtheConfig(ctx context.Context, data *RdpclientprofileResourceModel) rdp.Rdpclientprofile {
	tflog.Debug(ctx, "In rdpclientprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	rdpclientprofile := rdp.Rdpclientprofile{}
	if !data.Addusernameinrdpfile.IsNull() {
		rdpclientprofile.Addusernameinrdpfile = data.Addusernameinrdpfile.ValueString()
	}
	if !data.Audiocapturemode.IsNull() {
		rdpclientprofile.Audiocapturemode = data.Audiocapturemode.ValueString()
	}
	if !data.Keyboardhook.IsNull() {
		rdpclientprofile.Keyboardhook = data.Keyboardhook.ValueString()
	}
	if !data.Multimonitorsupport.IsNull() {
		rdpclientprofile.Multimonitorsupport = data.Multimonitorsupport.ValueString()
	}
	if !data.Name.IsNull() {
		rdpclientprofile.Name = data.Name.ValueString()
	}
	if !data.Psk.IsNull() {
		rdpclientprofile.Psk = data.Psk.ValueString()
	}
	if !data.Randomizerdpfilename.IsNull() {
		rdpclientprofile.Randomizerdpfilename = data.Randomizerdpfilename.ValueString()
	}
	if !data.Rdpcookievalidity.IsNull() {
		rdpclientprofile.Rdpcookievalidity = utils.IntPtr(int(data.Rdpcookievalidity.ValueInt64()))
	}
	if !data.Rdpcustomparams.IsNull() {
		rdpclientprofile.Rdpcustomparams = data.Rdpcustomparams.ValueString()
	}
	if !data.Rdpfilename.IsNull() {
		rdpclientprofile.Rdpfilename = data.Rdpfilename.ValueString()
	}
	if !data.Rdphost.IsNull() {
		rdpclientprofile.Rdphost = data.Rdphost.ValueString()
	}
	if !data.Rdplinkattribute.IsNull() {
		rdpclientprofile.Rdplinkattribute = data.Rdplinkattribute.ValueString()
	}
	if !data.Rdplistener.IsNull() {
		rdpclientprofile.Rdplistener = data.Rdplistener.ValueString()
	}
	if !data.Rdpurloverride.IsNull() {
		rdpclientprofile.Rdpurloverride = data.Rdpurloverride.ValueString()
	}
	if !data.Rdpvalidateclientip.IsNull() {
		rdpclientprofile.Rdpvalidateclientip = data.Rdpvalidateclientip.ValueString()
	}
	if !data.Redirectclipboard.IsNull() {
		rdpclientprofile.Redirectclipboard = data.Redirectclipboard.ValueString()
	}
	if !data.Redirectcomports.IsNull() {
		rdpclientprofile.Redirectcomports = data.Redirectcomports.ValueString()
	}
	if !data.Redirectdrives.IsNull() {
		rdpclientprofile.Redirectdrives = data.Redirectdrives.ValueString()
	}
	if !data.Redirectpnpdevices.IsNull() {
		rdpclientprofile.Redirectpnpdevices = data.Redirectpnpdevices.ValueString()
	}
	if !data.Redirectprinters.IsNull() {
		rdpclientprofile.Redirectprinters = data.Redirectprinters.ValueString()
	}
	if !data.Videoplaybackmode.IsNull() {
		rdpclientprofile.Videoplaybackmode = data.Videoplaybackmode.ValueString()
	}

	return rdpclientprofile
}

func rdpclientprofileSetAttrFromGet(ctx context.Context, data *RdpclientprofileResourceModel, getResponseData map[string]interface{}) *RdpclientprofileResourceModel {
	tflog.Debug(ctx, "In rdpclientprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["addusernameinrdpfile"]; ok && val != nil {
		data.Addusernameinrdpfile = types.StringValue(val.(string))
	} else {
		data.Addusernameinrdpfile = types.StringNull()
	}
	if val, ok := getResponseData["audiocapturemode"]; ok && val != nil {
		data.Audiocapturemode = types.StringValue(val.(string))
	} else {
		data.Audiocapturemode = types.StringNull()
	}
	if val, ok := getResponseData["keyboardhook"]; ok && val != nil {
		data.Keyboardhook = types.StringValue(val.(string))
	} else {
		data.Keyboardhook = types.StringNull()
	}
	if val, ok := getResponseData["multimonitorsupport"]; ok && val != nil {
		data.Multimonitorsupport = types.StringValue(val.(string))
	} else {
		data.Multimonitorsupport = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["psk"]; ok && val != nil {
		data.Psk = types.StringValue(val.(string))
	} else {
		data.Psk = types.StringNull()
	}
	if val, ok := getResponseData["randomizerdpfilename"]; ok && val != nil {
		data.Randomizerdpfilename = types.StringValue(val.(string))
	} else {
		data.Randomizerdpfilename = types.StringNull()
	}
	if val, ok := getResponseData["rdpcookievalidity"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Rdpcookievalidity = types.Int64Value(intVal)
		}
	} else {
		data.Rdpcookievalidity = types.Int64Null()
	}
	if val, ok := getResponseData["rdpcustomparams"]; ok && val != nil {
		data.Rdpcustomparams = types.StringValue(val.(string))
	} else {
		data.Rdpcustomparams = types.StringNull()
	}
	if val, ok := getResponseData["rdpfilename"]; ok && val != nil {
		data.Rdpfilename = types.StringValue(val.(string))
	} else {
		data.Rdpfilename = types.StringNull()
	}
	if val, ok := getResponseData["rdphost"]; ok && val != nil {
		data.Rdphost = types.StringValue(val.(string))
	} else {
		data.Rdphost = types.StringNull()
	}
	if val, ok := getResponseData["rdplinkattribute"]; ok && val != nil {
		data.Rdplinkattribute = types.StringValue(val.(string))
	} else {
		data.Rdplinkattribute = types.StringNull()
	}
	if val, ok := getResponseData["rdplistener"]; ok && val != nil {
		data.Rdplistener = types.StringValue(val.(string))
	} else {
		data.Rdplistener = types.StringNull()
	}
	if val, ok := getResponseData["rdpurloverride"]; ok && val != nil {
		data.Rdpurloverride = types.StringValue(val.(string))
	} else {
		data.Rdpurloverride = types.StringNull()
	}
	if val, ok := getResponseData["rdpvalidateclientip"]; ok && val != nil {
		data.Rdpvalidateclientip = types.StringValue(val.(string))
	} else {
		data.Rdpvalidateclientip = types.StringNull()
	}
	if val, ok := getResponseData["redirectclipboard"]; ok && val != nil {
		data.Redirectclipboard = types.StringValue(val.(string))
	} else {
		data.Redirectclipboard = types.StringNull()
	}
	if val, ok := getResponseData["redirectcomports"]; ok && val != nil {
		data.Redirectcomports = types.StringValue(val.(string))
	} else {
		data.Redirectcomports = types.StringNull()
	}
	if val, ok := getResponseData["redirectdrives"]; ok && val != nil {
		data.Redirectdrives = types.StringValue(val.(string))
	} else {
		data.Redirectdrives = types.StringNull()
	}
	if val, ok := getResponseData["redirectpnpdevices"]; ok && val != nil {
		data.Redirectpnpdevices = types.StringValue(val.(string))
	} else {
		data.Redirectpnpdevices = types.StringNull()
	}
	if val, ok := getResponseData["redirectprinters"]; ok && val != nil {
		data.Redirectprinters = types.StringValue(val.(string))
	} else {
		data.Redirectprinters = types.StringNull()
	}
	if val, ok := getResponseData["videoplaybackmode"]; ok && val != nil {
		data.Videoplaybackmode = types.StringValue(val.(string))
	} else {
		data.Videoplaybackmode = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
