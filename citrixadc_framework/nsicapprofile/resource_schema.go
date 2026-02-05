package nsicapprofile

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsicapprofileResourceModel describes the resource data model.
type NsicapprofileResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Allow204            types.String `tfsdk:"allow204"`
	Connectionkeepalive types.String `tfsdk:"connectionkeepalive"`
	Hostheader          types.String `tfsdk:"hostheader"`
	Inserthttprequest   types.String `tfsdk:"inserthttprequest"`
	Inserticapheaders   types.String `tfsdk:"inserticapheaders"`
	Logaction           types.String `tfsdk:"logaction"`
	Mode                types.String `tfsdk:"mode"`
	Name                types.String `tfsdk:"name"`
	Preview             types.String `tfsdk:"preview"`
	Previewlength       types.Int64  `tfsdk:"previewlength"`
	Queryparams         types.String `tfsdk:"queryparams"`
	Reqtimeout          types.Int64  `tfsdk:"reqtimeout"`
	Reqtimeoutaction    types.String `tfsdk:"reqtimeoutaction"`
	Uri                 types.String `tfsdk:"uri"`
	Useragent           types.String `tfsdk:"useragent"`
}

func (r *NsicapprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsicapprofile resource.",
			},
			"allow204": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or Disable sending Allow: 204 header in ICAP request.",
			},
			"connectionkeepalive": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "If enabled, Citrix ADC keeps the ICAP connection alive after a transaction to reuse it to send next ICAP request.",
			},
			"hostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICAP Host Header",
			},
			"inserthttprequest": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Exact HTTP request, in the form of an expression, which the Citrix ADC encapsulates and sends to the ICAP server. If you set this parameter, the ICAP request is sent using only this header. This can be used when the HTTP header is not available to send or ICAP server only needs part of the incoming HTTP request. The request expression is constrained by the feature for which it is used.\nThe Citrix ADC does not check the validity of this request. You must manually validate the request.",
			},
			"inserticapheaders": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert custom ICAP headers in the ICAP request to send to ICAP server. The headers can be static or can be dynamically constructed using PI Policy Expression. For example, to send static user agent and Client's IP address, the expression can be specified as \"User-Agent: NS-ICAP-Client/V1.0\\r\\nX-Client-IP: \"+CLIENT.IP.SRC+\"\\r\\n\".\nThe Citrix ADC does not check the validity of the specified header name-value. You must manually validate the specified header syntax.",
			},
			"logaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the audit message action which would be evaluated on receiving the ICAP response to emit the logs.",
			},
			"mode": schema.StringAttribute{
				Required:    true,
				Description: "ICAP Mode of operation. It is a mandatory argument while creating an icapprofile.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for an ICAP profile. Must begin with a letter, number, or the underscore \\(_\\) character. Other characters allowed, after the first character, are the hyphen \\(-\\), period \\(.\\), hash \\(\\#\\), space \\( \\), at \\(@\\), colon \\(:\\), and equal \\(=\\) characters. The name of a ICAP profile cannot be changed after it is created.\n\nCLI Users: If the name includes one or more spaces, enclose the name in double or single quotation marks \\(for example, \"my icap profile\" or 'my icap profile'\\).",
			},
			"preview": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or Disable preview header with ICAP request. This feature allows an ICAP server to see the beginning of a transaction, then decide if it wants to opt-out of the transaction early instead of receiving the remainder of the request message.",
			},
			"previewlength": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(4096),
				Description: "Value of Preview Header field. Citrix ADC uses the minimum of this set value and the preview size received on OPTIONS response.",
			},
			"queryparams": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Query parameters to be included with ICAP request URI. Entered values should be in arg=value format. For more than one parameters, add & separated values. e.g.: arg1=val1&arg2=val2.",
			},
			"reqtimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Time, in seconds, within which the remote server should respond to the ICAP-request. If the Netscaler does not receive full response with this time, the specified request timeout action is performed. Zero value disables this timeout functionality.",
			},
			"reqtimeoutaction": schema.StringAttribute{
				Optional:    true,
				Default:     stringdefault.StaticString("RESET"),
				Description: "Name of the action to perform if the Vserver/Server representing the remote service does not respond with any response within the timeout value configured. The Supported actions are\n* BYPASS - This Ignores the remote server response and sends the request/response to Client/Server.\n           * If the ICAP response with Encapsulated headers is not received within the request-timeout value configured, this Ignores the remote ICAP server response and sends the Full request/response to Server/Client.\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.",
			},
			"uri": schema.StringAttribute{
				Required:    true,
				Description: "URI representing icap service. It is a mandatory argument while creating an icapprofile.",
			},
			"useragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ICAP User Agent Header String",
			},
		},
	}
}

func nsicapprofileGetThePayloadFromtheConfig(ctx context.Context, data *NsicapprofileResourceModel) ns.Nsicapprofile {
	tflog.Debug(ctx, "In nsicapprofileGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsicapprofile := ns.Nsicapprofile{}
	if !data.Allow204.IsNull() {
		nsicapprofile.Allow204 = data.Allow204.ValueString()
	}
	if !data.Connectionkeepalive.IsNull() {
		nsicapprofile.Connectionkeepalive = data.Connectionkeepalive.ValueString()
	}
	if !data.Hostheader.IsNull() {
		nsicapprofile.Hostheader = data.Hostheader.ValueString()
	}
	if !data.Inserthttprequest.IsNull() {
		nsicapprofile.Inserthttprequest = data.Inserthttprequest.ValueString()
	}
	if !data.Inserticapheaders.IsNull() {
		nsicapprofile.Inserticapheaders = data.Inserticapheaders.ValueString()
	}
	if !data.Logaction.IsNull() {
		nsicapprofile.Logaction = data.Logaction.ValueString()
	}
	if !data.Mode.IsNull() {
		nsicapprofile.Mode = data.Mode.ValueString()
	}
	if !data.Name.IsNull() {
		nsicapprofile.Name = data.Name.ValueString()
	}
	if !data.Preview.IsNull() {
		nsicapprofile.Preview = data.Preview.ValueString()
	}
	if !data.Previewlength.IsNull() {
		nsicapprofile.Previewlength = utils.IntPtr(int(data.Previewlength.ValueInt64()))
	}
	if !data.Queryparams.IsNull() {
		nsicapprofile.Queryparams = data.Queryparams.ValueString()
	}
	if !data.Reqtimeout.IsNull() {
		nsicapprofile.Reqtimeout = utils.IntPtr(int(data.Reqtimeout.ValueInt64()))
	}
	if !data.Reqtimeoutaction.IsNull() {
		nsicapprofile.Reqtimeoutaction = data.Reqtimeoutaction.ValueString()
	}
	if !data.Uri.IsNull() {
		nsicapprofile.Uri = data.Uri.ValueString()
	}
	if !data.Useragent.IsNull() {
		nsicapprofile.Useragent = data.Useragent.ValueString()
	}

	return nsicapprofile
}

func nsicapprofileSetAttrFromGet(ctx context.Context, data *NsicapprofileResourceModel, getResponseData map[string]interface{}) *NsicapprofileResourceModel {
	tflog.Debug(ctx, "In nsicapprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["allow204"]; ok && val != nil {
		data.Allow204 = types.StringValue(val.(string))
	} else {
		data.Allow204 = types.StringNull()
	}
	if val, ok := getResponseData["connectionkeepalive"]; ok && val != nil {
		data.Connectionkeepalive = types.StringValue(val.(string))
	} else {
		data.Connectionkeepalive = types.StringNull()
	}
	if val, ok := getResponseData["hostheader"]; ok && val != nil {
		data.Hostheader = types.StringValue(val.(string))
	} else {
		data.Hostheader = types.StringNull()
	}
	if val, ok := getResponseData["inserthttprequest"]; ok && val != nil {
		data.Inserthttprequest = types.StringValue(val.(string))
	} else {
		data.Inserthttprequest = types.StringNull()
	}
	if val, ok := getResponseData["inserticapheaders"]; ok && val != nil {
		data.Inserticapheaders = types.StringValue(val.(string))
	} else {
		data.Inserticapheaders = types.StringNull()
	}
	if val, ok := getResponseData["logaction"]; ok && val != nil {
		data.Logaction = types.StringValue(val.(string))
	} else {
		data.Logaction = types.StringNull()
	}
	if val, ok := getResponseData["mode"]; ok && val != nil {
		data.Mode = types.StringValue(val.(string))
	} else {
		data.Mode = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["preview"]; ok && val != nil {
		data.Preview = types.StringValue(val.(string))
	} else {
		data.Preview = types.StringNull()
	}
	if val, ok := getResponseData["previewlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Previewlength = types.Int64Value(intVal)
		}
	} else {
		data.Previewlength = types.Int64Null()
	}
	if val, ok := getResponseData["queryparams"]; ok && val != nil {
		data.Queryparams = types.StringValue(val.(string))
	} else {
		data.Queryparams = types.StringNull()
	}
	if val, ok := getResponseData["reqtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Reqtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Reqtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["reqtimeoutaction"]; ok && val != nil {
		data.Reqtimeoutaction = types.StringValue(val.(string))
	} else {
		data.Reqtimeoutaction = types.StringNull()
	}
	if val, ok := getResponseData["uri"]; ok && val != nil {
		data.Uri = types.StringValue(val.(string))
	} else {
		data.Uri = types.StringNull()
	}
	if val, ok := getResponseData["useragent"]; ok && val != nil {
		data.Useragent = types.StringValue(val.(string))
	} else {
		data.Useragent = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
