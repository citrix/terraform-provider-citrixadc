package aaatacacsparams

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/aaa"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AaatacacsparamsResourceModel describes the resource data model.
type AaatacacsparamsResourceModel struct {
	Id                         types.String `tfsdk:"id"`
	Accounting                 types.String `tfsdk:"accounting"`
	Auditfailedcmds            types.String `tfsdk:"auditfailedcmds"`
	Authorization              types.String `tfsdk:"authorization"`
	Authtimeout                types.Int64  `tfsdk:"authtimeout"`
	Defaultauthenticationgroup types.String `tfsdk:"defaultauthenticationgroup"`
	Groupattrname              types.String `tfsdk:"groupattrname"`
	Serverip                   types.String `tfsdk:"serverip"`
	Serverport                 types.Int64  `tfsdk:"serverport"`
	Tacacssecret               types.String `tfsdk:"tacacssecret"`
}

func (r *AaatacacsparamsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the aaatacacsparams resource.",
			},
			"accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send accounting messages to the TACACS+ server.",
			},
			"auditfailedcmds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The option for sending accounting messages to the TACACS+ server.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use streaming authorization on the TACACS+ server.",
			},
			"authtimeout": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(3),
				Description: "Maximum number of seconds that the Citrix ADC waits for a response from the TACACS+ server.",
			},
			"defaultauthenticationgroup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the default group that is chosen when the authentication succeeds in addition to extracted groups.",
			},
			"groupattrname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TACACS+ group attribute name.Used for group extraction on the TACACS+ server.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of your TACACS+ server.",
			},
			"serverport": schema.Int64Attribute{
				Optional:    true,
				Default:     int64default.StaticInt64(49),
				Description: "Port number on which the TACACS+ server listens for connections.",
			},
			"tacacssecret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Key shared between the TACACS+ server and clients. Required for allowing the Citrix ADC to communicate with the TACACS+ server.",
			},
		},
	}
}

func aaatacacsparamsGetThePayloadFromtheConfig(ctx context.Context, data *AaatacacsparamsResourceModel) aaa.Aaatacacsparams {
	tflog.Debug(ctx, "In aaatacacsparamsGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	aaatacacsparams := aaa.Aaatacacsparams{}
	if !data.Accounting.IsNull() {
		aaatacacsparams.Accounting = data.Accounting.ValueString()
	}
	if !data.Auditfailedcmds.IsNull() {
		aaatacacsparams.Auditfailedcmds = data.Auditfailedcmds.ValueString()
	}
	if !data.Authorization.IsNull() {
		aaatacacsparams.Authorization = data.Authorization.ValueString()
	}
	if !data.Authtimeout.IsNull() {
		aaatacacsparams.Authtimeout = utils.IntPtr(int(data.Authtimeout.ValueInt64()))
	}
	if !data.Defaultauthenticationgroup.IsNull() {
		aaatacacsparams.Defaultauthenticationgroup = data.Defaultauthenticationgroup.ValueString()
	}
	if !data.Groupattrname.IsNull() {
		aaatacacsparams.Groupattrname = data.Groupattrname.ValueString()
	}
	if !data.Serverip.IsNull() {
		aaatacacsparams.Serverip = data.Serverip.ValueString()
	}
	if !data.Serverport.IsNull() {
		aaatacacsparams.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Tacacssecret.IsNull() {
		aaatacacsparams.Tacacssecret = data.Tacacssecret.ValueString()
	}

	return aaatacacsparams
}

func aaatacacsparamsSetAttrFromGet(ctx context.Context, data *AaatacacsparamsResourceModel, getResponseData map[string]interface{}) *AaatacacsparamsResourceModel {
	tflog.Debug(ctx, "In aaatacacsparamsSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accounting"]; ok && val != nil {
		data.Accounting = types.StringValue(val.(string))
	} else {
		data.Accounting = types.StringNull()
	}
	if val, ok := getResponseData["auditfailedcmds"]; ok && val != nil {
		data.Auditfailedcmds = types.StringValue(val.(string))
	} else {
		data.Auditfailedcmds = types.StringNull()
	}
	if val, ok := getResponseData["authorization"]; ok && val != nil {
		data.Authorization = types.StringValue(val.(string))
	} else {
		data.Authorization = types.StringNull()
	}
	if val, ok := getResponseData["authtimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Authtimeout = types.Int64Value(intVal)
		}
	} else {
		data.Authtimeout = types.Int64Null()
	}
	if val, ok := getResponseData["defaultauthenticationgroup"]; ok && val != nil {
		data.Defaultauthenticationgroup = types.StringValue(val.(string))
	} else {
		data.Defaultauthenticationgroup = types.StringNull()
	}
	if val, ok := getResponseData["groupattrname"]; ok && val != nil {
		data.Groupattrname = types.StringValue(val.(string))
	} else {
		data.Groupattrname = types.StringNull()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else {
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["tacacssecret"]; ok && val != nil {
		data.Tacacssecret = types.StringValue(val.(string))
	} else {
		data.Tacacssecret = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("aaatacacsparams-config")

	return data
}
