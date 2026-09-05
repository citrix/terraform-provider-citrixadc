package icaaccessprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// IcaaccessprofileResourceModel describes the resource data model.
type IcaaccessprofileResourceModel struct {
	Id                           types.String `tfsdk:"id"`
	Clientaudioredirection       types.String `tfsdk:"clientaudioredirection"`
	Clientclipboardredirection   types.String `tfsdk:"clientclipboardredirection"`
	Clientcomportredirection     types.String `tfsdk:"clientcomportredirection"`
	Clientdriveredirection       types.String `tfsdk:"clientdriveredirection"`
	Clientprinterredirection     types.String `tfsdk:"clientprinterredirection"`
	Clienttwaindeviceredirection types.String `tfsdk:"clienttwaindeviceredirection"`
	Clientusbdriveredirection    types.String `tfsdk:"clientusbdriveredirection"`
	Connectclientlptports        types.String `tfsdk:"connectclientlptports"`
	Draganddrop                  types.String `tfsdk:"draganddrop"`
	Fido2redirection             types.String `tfsdk:"fido2redirection"`
	Localremotedatasharing       types.String `tfsdk:"localremotedatasharing"`
	Multistream                  types.String `tfsdk:"multistream"`
	Name                         types.String `tfsdk:"name"`
	Smartcardredirection         types.String `tfsdk:"smartcardredirection"`
	Wiaredirection               types.String `tfsdk:"wiaredirection"`
}

func (r *IcaaccessprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icaaccessprofile resource.",
			},
			"clientaudioredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable applications hosted on the server to play sounds through a sound device installed on the client computer, also allows or prevents users to record audio input",
			},
			"clientclipboardredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable the clipboard on the client device to be mapped to the clipboard on the server",
			},
			"clientcomportredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable COM port redirection to and from the client",
			},
			"clientdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disables drive redirection to and from the client",
			},
			"clientprinterredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable client printers to be mapped to a server when a user logs on to a session",
			},
			"clienttwaindeviceredirection": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Allow default access or disable TWAIN devices, such as digital cameras or scanners, on the client device from published image processing applications",
			},
			"clientusbdriveredirection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable the redirection of USB devices to and from the client",
			},
			"connectclientlptports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable automatic connection of LPT ports from the client when the user logs on",
			},
			"draganddrop": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Allow default access or disable drag and drop between client and remote applications and desktops",
			},
			"fido2redirection": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Allow default access or disable FIDO2 redirection",
			},
			"localremotedatasharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable file/data sharing via the Receiver for HTML5",
			},
			"multistream": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Allow Default access/Disable the multistream feature for the specified users",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the ICA accessprofile. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and\nthe hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA accessprofile is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica accessprofile\" or 'my ica accessprofile').\n\nEach of the features can be configured as DEFAULT/DISABLED.\nHere, DISABLED means that the policy settings on the backend XenApp/XenDesktop server are overridden and the Citrix ADC makes the decision to deny access. Whereas DEFAULT means that the Citrix ADC allows the request to reach the XenApp/XenDesktop that takes the decision to allow/deny access based on the policy configured on it. For example, if ClientAudioRedirection is enabled on the backend XenApp/XenDesktop server, and the configured profile has ClientAudioRedirection as DISABLED, the Citrix ADC makes the decision to deny the request irrespective of the configuration on the backend. If the configured profile has ClientAudioRedirection as DEFAULT, then the Citrix ADC forwards the requests to the backend XenApp/XenDesktop server.It then makes the decision to allow/deny access based on the policy configured on it.",
			},
			"smartcardredirection": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Allow default access or disable smart card redirection. Smart card virtual channel is always allowed in CVAD",
			},
			"wiaredirection": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					utils.UnsetOnRemoveOrKeepDefaultString{DefaultValue: "DISABLED"},
				},
				Description: "Allow default access or disable WIA scanner redirection",
			},
		},
	}
}

func icaaccessprofileGetThePayloadFromthePlan(ctx context.Context, data *IcaaccessprofileResourceModel) ica.Icaaccessprofile {
	tflog.Debug(ctx, "In icaaccessprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	icaaccessprofile := ica.Icaaccessprofile{}
	if !data.Clientaudioredirection.IsNull() && !data.Clientaudioredirection.IsUnknown() {
		icaaccessprofile.Clientaudioredirection = data.Clientaudioredirection.ValueString()
	}
	if !data.Clientclipboardredirection.IsNull() && !data.Clientclipboardredirection.IsUnknown() {
		icaaccessprofile.Clientclipboardredirection = data.Clientclipboardredirection.ValueString()
	}
	if !data.Clientcomportredirection.IsNull() && !data.Clientcomportredirection.IsUnknown() {
		icaaccessprofile.Clientcomportredirection = data.Clientcomportredirection.ValueString()
	}
	if !data.Clientdriveredirection.IsNull() && !data.Clientdriveredirection.IsUnknown() {
		icaaccessprofile.Clientdriveredirection = data.Clientdriveredirection.ValueString()
	}
	if !data.Clientprinterredirection.IsNull() && !data.Clientprinterredirection.IsUnknown() {
		icaaccessprofile.Clientprinterredirection = data.Clientprinterredirection.ValueString()
	}
	if !data.Clienttwaindeviceredirection.IsNull() && !data.Clienttwaindeviceredirection.IsUnknown() {
		icaaccessprofile.Clienttwaindeviceredirection = data.Clienttwaindeviceredirection.ValueString()
	}
	if !data.Clientusbdriveredirection.IsNull() && !data.Clientusbdriveredirection.IsUnknown() {
		icaaccessprofile.Clientusbdriveredirection = data.Clientusbdriveredirection.ValueString()
	}
	if !data.Connectclientlptports.IsNull() && !data.Connectclientlptports.IsUnknown() {
		icaaccessprofile.Connectclientlptports = data.Connectclientlptports.ValueString()
	}
	if !data.Draganddrop.IsNull() && !data.Draganddrop.IsUnknown() {
		icaaccessprofile.Draganddrop = data.Draganddrop.ValueString()
	}
	if !data.Fido2redirection.IsNull() && !data.Fido2redirection.IsUnknown() {
		icaaccessprofile.Fido2redirection = data.Fido2redirection.ValueString()
	}
	if !data.Localremotedatasharing.IsNull() && !data.Localremotedatasharing.IsUnknown() {
		icaaccessprofile.Localremotedatasharing = data.Localremotedatasharing.ValueString()
	}
	if !data.Multistream.IsNull() && !data.Multistream.IsUnknown() {
		icaaccessprofile.Multistream = data.Multistream.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		icaaccessprofile.Name = data.Name.ValueString()
	}
	if !data.Smartcardredirection.IsNull() && !data.Smartcardredirection.IsUnknown() {
		icaaccessprofile.Smartcardredirection = data.Smartcardredirection.ValueString()
	}
	if !data.Wiaredirection.IsNull() && !data.Wiaredirection.IsUnknown() {
		icaaccessprofile.Wiaredirection = data.Wiaredirection.ValueString()
	}

	return icaaccessprofile
}

func icaaccessprofileSetAttrFromGet(ctx context.Context, data *IcaaccessprofileResourceModel, getResponseData map[string]interface{}) *IcaaccessprofileResourceModel {
	tflog.Debug(ctx, "In icaaccessprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clientaudioredirection"]; ok && val != nil {
		data.Clientaudioredirection = types.StringValue(val.(string))
	} else {
		data.Clientaudioredirection = types.StringNull()
	}
	if val, ok := getResponseData["clientclipboardredirection"]; ok && val != nil {
		data.Clientclipboardredirection = types.StringValue(val.(string))
	} else {
		data.Clientclipboardredirection = types.StringNull()
	}
	if val, ok := getResponseData["clientcomportredirection"]; ok && val != nil {
		data.Clientcomportredirection = types.StringValue(val.(string))
	} else {
		data.Clientcomportredirection = types.StringNull()
	}
	if val, ok := getResponseData["clientdriveredirection"]; ok && val != nil {
		data.Clientdriveredirection = types.StringValue(val.(string))
	} else {
		data.Clientdriveredirection = types.StringNull()
	}
	if val, ok := getResponseData["clientprinterredirection"]; ok && val != nil {
		data.Clientprinterredirection = types.StringValue(val.(string))
	} else {
		data.Clientprinterredirection = types.StringNull()
	}
	if val, ok := getResponseData["clienttwaindeviceredirection"]; ok && val != nil {
		data.Clienttwaindeviceredirection = types.StringValue(val.(string))
	} else {
		data.Clienttwaindeviceredirection = types.StringNull()
	}
	if val, ok := getResponseData["clientusbdriveredirection"]; ok && val != nil {
		data.Clientusbdriveredirection = types.StringValue(val.(string))
	} else {
		data.Clientusbdriveredirection = types.StringNull()
	}
	if val, ok := getResponseData["connectclientlptports"]; ok && val != nil {
		data.Connectclientlptports = types.StringValue(val.(string))
	} else {
		data.Connectclientlptports = types.StringNull()
	}
	if val, ok := getResponseData["draganddrop"]; ok && val != nil {
		data.Draganddrop = types.StringValue(val.(string))
	} else {
		data.Draganddrop = types.StringNull()
	}
	if val, ok := getResponseData["fido2redirection"]; ok && val != nil {
		data.Fido2redirection = types.StringValue(val.(string))
	} else {
		data.Fido2redirection = types.StringNull()
	}
	if val, ok := getResponseData["localremotedatasharing"]; ok && val != nil {
		data.Localremotedatasharing = types.StringValue(val.(string))
	} else {
		data.Localremotedatasharing = types.StringNull()
	}
	if val, ok := getResponseData["multistream"]; ok && val != nil {
		data.Multistream = types.StringValue(val.(string))
	} else {
		data.Multistream = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["smartcardredirection"]; ok && val != nil {
		data.Smartcardredirection = types.StringValue(val.(string))
	} else {
		data.Smartcardredirection = types.StringNull()
	}
	if val, ok := getResponseData["wiaredirection"]; ok && val != nil {
		data.Wiaredirection = types.StringValue(val.(string))
	} else {
		data.Wiaredirection = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
