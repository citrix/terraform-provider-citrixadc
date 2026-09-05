package sslparameter

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslparameterResourceModel describes the resource data model.
type SslparameterResourceModel struct {
	Id                       types.String `tfsdk:"id"`
	Crlmemorysizemb          types.Int64  `tfsdk:"crlmemorysizemb"`
	Cryptodevdisablelimit    types.Int64  `tfsdk:"cryptodevdisablelimit"`
	Defaultprofile           types.String `tfsdk:"defaultprofile"`
	Denysslreneg             types.String `tfsdk:"denysslreneg"`
	Dropreqwithnohostheader  types.String `tfsdk:"dropreqwithnohostheader"`
	Encrypttriggerpktcount   types.Int64  `tfsdk:"encrypttriggerpktcount"`
	Heterogeneoussslhw       types.String `tfsdk:"heterogeneoussslhw"`
	Hybridfipsmode           types.String `tfsdk:"hybridfipsmode"`
	Insertcertspace          types.String `tfsdk:"insertcertspace"`
	Insertionencoding        types.String `tfsdk:"insertionencoding"`
	Ndcppcompliancecertcheck types.String `tfsdk:"ndcppcompliancecertcheck"`
	Ocspcachesize            types.Int64  `tfsdk:"ocspcachesize"`
	Operationqueuelimit      types.Int64  `tfsdk:"operationqueuelimit"`
	Pushenctriggertimeout    types.Int64  `tfsdk:"pushenctriggertimeout"`
	Pushflag                 types.Int64  `tfsdk:"pushflag"`
	Quantumsize              types.String `tfsdk:"quantumsize"`
	Sendclosenotify          types.String `tfsdk:"sendclosenotify"`
	Sigdigesttype            types.List   `tfsdk:"sigdigesttype"`
	Snihttphostmatch         types.String `tfsdk:"snihttphostmatch"`
	Softwarecryptothreshold  types.Int64  `tfsdk:"softwarecryptothreshold"`
	Sslierrorcache           types.String `tfsdk:"sslierrorcache"`
	Sslimaxerrorcachemem     types.Int64  `tfsdk:"sslimaxerrorcachemem"`
	Ssltriggertimeout        types.Int64  `tfsdk:"ssltriggertimeout"`
	Strictcachecks           types.String `tfsdk:"strictcachecks"`
	Undefactioncontrol       types.String `tfsdk:"undefactioncontrol"`
	Undefactiondata          types.String `tfsdk:"undefactiondata"`
}

func (r *SslparameterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslparameter resource.",
			},
			"crlmemorysizemb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(256),
				Description: "Maximum memory size to use for certificate revocation lists (CRLs). This parameter reserves memory for a CRL but sets a limit to the maximum memory that the CRLs loaded on the appliance can consume.",
			},
			"cryptodevdisablelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit to the number of disabled SSL chips after which the ADC restarts. A value of zero implies that the ADC does not automatically restart.",
			},
			"defaultprofile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Global parameter used to enable default profile feature.",
			},
			"denysslreneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Deny renegotiation in specified circumstances. Available settings function as follows:\n* NO - Allow SSL renegotiation.\n* FRONTEND_CLIENT - Deny secure and nonsecure SSL renegotiation initiated by the client.\n* FRONTEND_CLIENTSERVER - Deny secure and nonsecure SSL renegotiation initiated by the client or the Citrix ADC during policy-based client authentication.\n* ALL - Deny all secure and nonsecure SSL renegotiation.\n* NONSECURE - Deny nonsecure SSL renegotiation. Allows only clients that support RFC 5746.",
			},
			"dropreqwithnohostheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Host header check for SNI enabled sessions. If this check is enabled and the HTTP request does not contain the host header for SNI enabled sessions(i.e vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension), the request is dropped.",
			},
			"encrypttriggerpktcount": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(45),
				Description: "Maximum number of queued packets after which encryption is triggered. Use this setting for SSL transactions that send small packets from server to Citrix ADC.",
			},
			"heterogeneoussslhw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To support both cavium and coleto based platforms in cluster environment, this mode has to be enabled.",
			},
			"hybridfipsmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When this mode is enabled, system will use additional crypto hardware to accelerate symmetric crypto operations.",
			},
			"insertcertspace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "To insert space between lines in the certificate header of request",
			},
			"insertionencoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("Unicode"),
				Description: "Encoding method used to insert the subject or issuer's name in HTTP requests to servers.",
			},
			"ndcppcompliancecertcheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Determines whether or not additional checks are carried out during a TLS handshake when validating an X.509 certificate received from the peer.",
			},
			"ocspcachesize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(10),
				Description: "Size, per packet engine, in megabytes, of the OCSP cache. A maximum of 10% of the packet engine memory can be assigned. Because the maximum allowed packet engine memory is 4GB, the maximum value that can be assigned to the OCSP cache is approximately 410 MB.",
			},
			"operationqueuelimit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Limit in percentage of capacity of the crypto operations queue beyond which new SSL connections are not accepted until the queue is reduced.",
			},
			"pushenctriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(1),
				Description: "PUSH encryption trigger timeout value. The timeout value is applied only if you set the Push Encryption Trigger parameter to Timer in the SSL virtual server settings.",
			},
			"pushflag": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Insert PUSH flag into decrypted, encrypted, or all records. If the PUSH flag is set to a value other than 0, the buffered records are forwarded on the basis of the value of the PUSH flag. Available settings function as follows:\n0 - Auto (PUSH flag is not set.)\n1 - Insert PUSH flag into every decrypted record.\n2 -Insert PUSH flag into every encrypted record.\n3 - Insert PUSH flag into every decrypted and encrypted record.",
			},
			"quantumsize": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("8192"),
				Description: "Amount of data to collect before the data is pushed to the crypto hardware for encryption. For large downloads, a larger quantum size better utilizes the crypto resources.",
			},
			"sendclosenotify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("YES"),
				Description: "Send an SSL Close-Notify message to the client at the end of a transaction.",
			},
			"sigdigesttype": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Signature Digest Algorithms that are supported by appliance. Default value is \"ALL\" and it will enable the following algorithms depending on the platform.\nOn VPX: ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512 RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 DSA-SHA1 DSA-SHA224 DSA-SHA256 DSA-SHA384 DSA-SHA512\nOn MPX with Nitrox-III and coleto cards: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512 ECDSA-SHA1 ECDSA-SHA224 ECDSA-SHA256 ECDSA-SHA384 ECDSA-SHA512\nOthers: RSA-SHA1 RSA-SHA224 RSA-SHA256 RSA-SHA384 RSA-SHA512.\nNote:ALL doesnot include RSA-MD5 for any platform.",
			},
			"snihttphostmatch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("CERT"),
				Description: "Controls how the HTTP 'Host' header value is validated. These checks are performed only if the session is SNI enabled (i.e when vserver or profile bound to vserver has SNI enabled and 'Client Hello' arrived with SNI extension) and HTTP request contains 'Host' header.",
			},
			"softwarecryptothreshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC CPU utilization threshold (in percentage) beyond which crypto operations are not done in software.\nA value of zero implies that CPU is not utilized for doing crypto in software.",
			},
			"sslierrorcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("DISABLED"),
				Description: "Enable or disable dynamically learning and caching the learned information to make the subsequent interception or bypass decision. When enabled, NS does the lookup of this cached data to do early bypass.",
			},
			"sslimaxerrorcachemem": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify the maximum memory that can be used for caching the learned data. This memory is used as a LRU cache so that the old entries gets replaced with new entry once the set memory limit is fully utilised. A value of 0 decides the limit automatically.",
			},
			"ssltriggertimeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Default:     int64default.StaticInt64(100),
				Description: "Time, in milliseconds, after which encryption is triggered for transactions that are not tracked on the Citrix ADC because their length is not known. There can be a delay of up to 10ms from the specified timeout value before the packet is pushed into the queue.",
			},
			"strictcachecks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NO"),
				Description: "Enable strict CA certificate checks on the appliance.",
			},
			"undefactioncontrol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("CLIENTAUTH"),
				Description: "Name of the undefined built-in control action: CLIENTAUTH, NOCLIENTAUTH, NOOP, RESET, or DROP.",
			},
			"undefactiondata": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("NOOP"),
				Description: "Name of the undefined built-in data action: NOOP, RESET or DROP.",
			},
		},
	}
}

func sslparameterGetThePayloadFromtheConfig(ctx context.Context, data *SslparameterResourceModel) ssl.Sslparameter {
	tflog.Debug(ctx, "In sslparameterGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslparameter := ssl.Sslparameter{}
	if !data.Crlmemorysizemb.IsNull() && !data.Crlmemorysizemb.IsUnknown() {
		sslparameter.Crlmemorysizemb = utils.IntPtr(int(data.Crlmemorysizemb.ValueInt64()))
	}
	if !data.Cryptodevdisablelimit.IsNull() && !data.Cryptodevdisablelimit.IsUnknown() {
		sslparameter.Cryptodevdisablelimit = utils.IntPtr(int(data.Cryptodevdisablelimit.ValueInt64()))
	}
	if !data.Defaultprofile.IsNull() && !data.Defaultprofile.IsUnknown() {
		sslparameter.Defaultprofile = data.Defaultprofile.ValueString()
	}
	if !data.Denysslreneg.IsNull() && !data.Denysslreneg.IsUnknown() {
		sslparameter.Denysslreneg = data.Denysslreneg.ValueString()
	}
	if !data.Dropreqwithnohostheader.IsNull() && !data.Dropreqwithnohostheader.IsUnknown() {
		sslparameter.Dropreqwithnohostheader = data.Dropreqwithnohostheader.ValueString()
	}
	if !data.Encrypttriggerpktcount.IsNull() && !data.Encrypttriggerpktcount.IsUnknown() {
		sslparameter.Encrypttriggerpktcount = utils.IntPtr(int(data.Encrypttriggerpktcount.ValueInt64()))
	}
	if !data.Heterogeneoussslhw.IsNull() && !data.Heterogeneoussslhw.IsUnknown() {
		sslparameter.Heterogeneoussslhw = data.Heterogeneoussslhw.ValueString()
	}
	if !data.Hybridfipsmode.IsNull() && !data.Hybridfipsmode.IsUnknown() {
		sslparameter.Hybridfipsmode = data.Hybridfipsmode.ValueString()
	}
	if !data.Insertcertspace.IsNull() && !data.Insertcertspace.IsUnknown() {
		sslparameter.Insertcertspace = data.Insertcertspace.ValueString()
	}
	if !data.Insertionencoding.IsNull() && !data.Insertionencoding.IsUnknown() {
		sslparameter.Insertionencoding = data.Insertionencoding.ValueString()
	}
	if !data.Ndcppcompliancecertcheck.IsNull() && !data.Ndcppcompliancecertcheck.IsUnknown() {
		sslparameter.Ndcppcompliancecertcheck = data.Ndcppcompliancecertcheck.ValueString()
	}
	if !data.Ocspcachesize.IsNull() && !data.Ocspcachesize.IsUnknown() {
		sslparameter.Ocspcachesize = utils.IntPtr(int(data.Ocspcachesize.ValueInt64()))
	}
	if !data.Operationqueuelimit.IsNull() && !data.Operationqueuelimit.IsUnknown() {
		sslparameter.Operationqueuelimit = utils.IntPtr(int(data.Operationqueuelimit.ValueInt64()))
	}
	if !data.Pushenctriggertimeout.IsNull() && !data.Pushenctriggertimeout.IsUnknown() {
		sslparameter.Pushenctriggertimeout = utils.IntPtr(int(data.Pushenctriggertimeout.ValueInt64()))
	}
	if !data.Pushflag.IsNull() && !data.Pushflag.IsUnknown() {
		sslparameter.Pushflag = utils.IntPtr(int(data.Pushflag.ValueInt64()))
	}
	if !data.Quantumsize.IsNull() && !data.Quantumsize.IsUnknown() {
		sslparameter.Quantumsize = data.Quantumsize.ValueString()
	}
	if !data.Sendclosenotify.IsNull() && !data.Sendclosenotify.IsUnknown() {
		sslparameter.Sendclosenotify = data.Sendclosenotify.ValueString()
	}
	if !data.Sigdigesttype.IsNull() && !data.Sigdigesttype.IsUnknown() {
		var sigdigesttypeList []string
		data.Sigdigesttype.ElementsAs(ctx, &sigdigesttypeList, false)
		sslparameter.Sigdigesttype = sigdigesttypeList
	}
	if !data.Snihttphostmatch.IsNull() && !data.Snihttphostmatch.IsUnknown() {
		sslparameter.Snihttphostmatch = data.Snihttphostmatch.ValueString()
	}
	if !data.Softwarecryptothreshold.IsNull() && !data.Softwarecryptothreshold.IsUnknown() {
		sslparameter.Softwarecryptothreshold = utils.IntPtr(int(data.Softwarecryptothreshold.ValueInt64()))
	}
	if !data.Sslierrorcache.IsNull() && !data.Sslierrorcache.IsUnknown() {
		sslparameter.Sslierrorcache = data.Sslierrorcache.ValueString()
	}
	if !data.Sslimaxerrorcachemem.IsNull() && !data.Sslimaxerrorcachemem.IsUnknown() {
		sslparameter.Sslimaxerrorcachemem = utils.IntPtr(int(data.Sslimaxerrorcachemem.ValueInt64()))
	}
	if !data.Ssltriggertimeout.IsNull() && !data.Ssltriggertimeout.IsUnknown() {
		sslparameter.Ssltriggertimeout = utils.IntPtr(int(data.Ssltriggertimeout.ValueInt64()))
	}
	if !data.Strictcachecks.IsNull() && !data.Strictcachecks.IsUnknown() {
		sslparameter.Strictcachecks = data.Strictcachecks.ValueString()
	}
	if !data.Undefactioncontrol.IsNull() && !data.Undefactioncontrol.IsUnknown() {
		sslparameter.Undefactioncontrol = data.Undefactioncontrol.ValueString()
	}
	if !data.Undefactiondata.IsNull() && !data.Undefactiondata.IsUnknown() {
		sslparameter.Undefactiondata = data.Undefactiondata.ValueString()
	}

	return sslparameter
}

func sslparameterSetAttrFromGet(ctx context.Context, data *SslparameterResourceModel, getResponseData map[string]interface{}) *SslparameterResourceModel {
	tflog.Debug(ctx, "In sslparameterSetAttrFromGet Function")

	// Convert API response to model.
	// NOTE (omit-on-default guard): NITRO omits attributes whose value equals the
	// device default from the GET response. Only null a missing attribute when the
	// current value is Unknown; never clobber a value the user explicitly configured.
	if val, ok := getResponseData["crlmemorysizemb"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Crlmemorysizemb = types.Int64Value(intVal)
		}
	} else if data.Crlmemorysizemb.IsUnknown() {
		data.Crlmemorysizemb = types.Int64Null()
	}
	if val, ok := getResponseData["cryptodevdisablelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cryptodevdisablelimit = types.Int64Value(intVal)
		}
	} else if data.Cryptodevdisablelimit.IsUnknown() {
		data.Cryptodevdisablelimit = types.Int64Null()
	}
	if val, ok := getResponseData["defaultprofile"]; ok && val != nil {
		data.Defaultprofile = types.StringValue(val.(string))
	} else if data.Defaultprofile.IsUnknown() {
		data.Defaultprofile = types.StringNull()
	}
	if val, ok := getResponseData["denysslreneg"]; ok && val != nil {
		data.Denysslreneg = types.StringValue(val.(string))
	} else if data.Denysslreneg.IsUnknown() {
		data.Denysslreneg = types.StringNull()
	}
	if val, ok := getResponseData["dropreqwithnohostheader"]; ok && val != nil {
		data.Dropreqwithnohostheader = types.StringValue(val.(string))
	} else if data.Dropreqwithnohostheader.IsUnknown() {
		data.Dropreqwithnohostheader = types.StringNull()
	}
	if val, ok := getResponseData["encrypttriggerpktcount"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Encrypttriggerpktcount = types.Int64Value(intVal)
		}
	} else if data.Encrypttriggerpktcount.IsUnknown() {
		data.Encrypttriggerpktcount = types.Int64Null()
	}
	if val, ok := getResponseData["heterogeneoussslhw"]; ok && val != nil {
		data.Heterogeneoussslhw = types.StringValue(val.(string))
	} else if data.Heterogeneoussslhw.IsUnknown() {
		data.Heterogeneoussslhw = types.StringNull()
	}
	if val, ok := getResponseData["hybridfipsmode"]; ok && val != nil {
		data.Hybridfipsmode = types.StringValue(val.(string))
	} else if data.Hybridfipsmode.IsUnknown() {
		data.Hybridfipsmode = types.StringNull()
	}
	if val, ok := getResponseData["insertcertspace"]; ok && val != nil {
		data.Insertcertspace = types.StringValue(val.(string))
	} else if data.Insertcertspace.IsUnknown() {
		data.Insertcertspace = types.StringNull()
	}
	if val, ok := getResponseData["insertionencoding"]; ok && val != nil {
		data.Insertionencoding = types.StringValue(val.(string))
	} else if data.Insertionencoding.IsUnknown() {
		data.Insertionencoding = types.StringNull()
	}
	if val, ok := getResponseData["ndcppcompliancecertcheck"]; ok && val != nil {
		data.Ndcppcompliancecertcheck = types.StringValue(val.(string))
	} else if data.Ndcppcompliancecertcheck.IsUnknown() {
		data.Ndcppcompliancecertcheck = types.StringNull()
	}
	if val, ok := getResponseData["ocspcachesize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ocspcachesize = types.Int64Value(intVal)
		}
	} else if data.Ocspcachesize.IsUnknown() {
		data.Ocspcachesize = types.Int64Null()
	}
	if val, ok := getResponseData["operationqueuelimit"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Operationqueuelimit = types.Int64Value(intVal)
		}
	} else if data.Operationqueuelimit.IsUnknown() {
		data.Operationqueuelimit = types.Int64Null()
	}
	if val, ok := getResponseData["pushenctriggertimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pushenctriggertimeout = types.Int64Value(intVal)
		}
	} else if data.Pushenctriggertimeout.IsUnknown() {
		data.Pushenctriggertimeout = types.Int64Null()
	}
	if val, ok := getResponseData["pushflag"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Pushflag = types.Int64Value(intVal)
		}
	} else if data.Pushflag.IsUnknown() {
		data.Pushflag = types.Int64Null()
	}
	if val, ok := getResponseData["quantumsize"]; ok && val != nil {
		data.Quantumsize = types.StringValue(val.(string))
	} else if data.Quantumsize.IsUnknown() {
		data.Quantumsize = types.StringNull()
	}
	if val, ok := getResponseData["sendclosenotify"]; ok && val != nil {
		data.Sendclosenotify = types.StringValue(val.(string))
	} else if data.Sendclosenotify.IsUnknown() {
		data.Sendclosenotify = types.StringNull()
	}
	if val, ok := getResponseData["sigdigesttype"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Sigdigesttype = listValue
		} else if data.Sigdigesttype.IsUnknown() {
			data.Sigdigesttype = types.ListNull(types.StringType)
		}
	} else if data.Sigdigesttype.IsUnknown() {
		data.Sigdigesttype = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["snihttphostmatch"]; ok && val != nil {
		data.Snihttphostmatch = types.StringValue(val.(string))
	} else if data.Snihttphostmatch.IsUnknown() {
		data.Snihttphostmatch = types.StringNull()
	}
	if val, ok := getResponseData["softwarecryptothreshold"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Softwarecryptothreshold = types.Int64Value(intVal)
		}
	} else if data.Softwarecryptothreshold.IsUnknown() {
		data.Softwarecryptothreshold = types.Int64Null()
	}
	if val, ok := getResponseData["sslierrorcache"]; ok && val != nil {
		data.Sslierrorcache = types.StringValue(val.(string))
	} else if data.Sslierrorcache.IsUnknown() {
		data.Sslierrorcache = types.StringNull()
	}
	if val, ok := getResponseData["sslimaxerrorcachemem"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Sslimaxerrorcachemem = types.Int64Value(intVal)
		}
	} else if data.Sslimaxerrorcachemem.IsUnknown() {
		data.Sslimaxerrorcachemem = types.Int64Null()
	}
	if val, ok := getResponseData["ssltriggertimeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Ssltriggertimeout = types.Int64Value(intVal)
		}
	} else if data.Ssltriggertimeout.IsUnknown() {
		data.Ssltriggertimeout = types.Int64Null()
	}
	if val, ok := getResponseData["strictcachecks"]; ok && val != nil {
		data.Strictcachecks = types.StringValue(val.(string))
	} else if data.Strictcachecks.IsUnknown() {
		data.Strictcachecks = types.StringNull()
	}
	if val, ok := getResponseData["undefactioncontrol"]; ok && val != nil {
		data.Undefactioncontrol = types.StringValue(val.(string))
	} else if data.Undefactioncontrol.IsUnknown() {
		data.Undefactioncontrol = types.StringNull()
	}
	if val, ok := getResponseData["undefactiondata"]; ok && val != nil {
		data.Undefactiondata = types.StringValue(val.(string))
	} else if data.Undefactiondata.IsUnknown() {
		data.Undefactiondata = types.StringNull()
	}

	// Set ID for the resource
	// Singleton resource - static ID
	data.Id = types.StringValue("sslparameter-config")

	return data
}
