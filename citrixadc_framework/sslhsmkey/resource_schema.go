package sslhsmkey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SslhsmkeyResourceModel describes the resource data model.
type SslhsmkeyResourceModel struct {
	Id         types.String `tfsdk:"id"`
	Hsmkeyname types.String `tfsdk:"hsmkeyname"`
	Hsmtype    types.String `tfsdk:"hsmtype"`
	Key        types.String `tfsdk:"key"`
	Keystore   types.String `tfsdk:"keystore"`
	Password   types.String `tfsdk:"password"`
	Serialnum  types.String `tfsdk:"serialnum"`
}

func (r *SslhsmkeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslhsmkey resource.",
			},
			"hsmkeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "0",
			},
			"hsmtype": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("THALES"),
				Description: "Type of HSM.",
			},
			"key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the key. optionally, for Thales, path to the HSM key file; /var/opt/nfast/kmdata/local/ is the default path. Applies when HSMTYPE is THALES or KEYVAULT.",
			},
			"keystore": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of keystore object representing HSM where key is stored. For example, name of keyvault object or azurekeyvault authentication object. Applies only to KEYVAULT type HSM.",
			},
			"password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Password for a partition. Applies only to SafeNet HSM.",
			},
			"serialnum": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Serial number of the partition on which the key is present. Applies only to SafeNet HSM.",
			},
		},
	}
}

func sslhsmkeyGetThePayloadFromtheConfig(ctx context.Context, data *SslhsmkeyResourceModel) ssl.Sslhsmkey {
	tflog.Debug(ctx, "In sslhsmkeyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslhsmkey := ssl.Sslhsmkey{}
	if !data.Hsmkeyname.IsNull() {
		sslhsmkey.Hsmkeyname = data.Hsmkeyname.ValueString()
	}
	if !data.Hsmtype.IsNull() {
		sslhsmkey.Hsmtype = data.Hsmtype.ValueString()
	}
	if !data.Key.IsNull() {
		sslhsmkey.Key = data.Key.ValueString()
	}
	if !data.Keystore.IsNull() {
		sslhsmkey.Keystore = data.Keystore.ValueString()
	}
	if !data.Password.IsNull() {
		sslhsmkey.Password = data.Password.ValueString()
	}
	if !data.Serialnum.IsNull() {
		sslhsmkey.Serialnum = data.Serialnum.ValueString()
	}

	return sslhsmkey
}

func sslhsmkeySetAttrFromGet(ctx context.Context, data *SslhsmkeyResourceModel, getResponseData map[string]interface{}) *SslhsmkeyResourceModel {
	tflog.Debug(ctx, "In sslhsmkeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["hsmkeyname"]; ok && val != nil {
		data.Hsmkeyname = types.StringValue(val.(string))
	} else {
		data.Hsmkeyname = types.StringNull()
	}
	if val, ok := getResponseData["hsmtype"]; ok && val != nil {
		data.Hsmtype = types.StringValue(val.(string))
	} else {
		data.Hsmtype = types.StringNull()
	}
	if val, ok := getResponseData["key"]; ok && val != nil {
		data.Key = types.StringValue(val.(string))
	} else {
		data.Key = types.StringNull()
	}
	if val, ok := getResponseData["keystore"]; ok && val != nil {
		data.Keystore = types.StringValue(val.(string))
	} else {
		data.Keystore = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["serialnum"]; ok && val != nil {
		data.Serialnum = types.StringValue(val.(string))
	} else {
		data.Serialnum = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Hsmkeyname.ValueString())

	return data
}
