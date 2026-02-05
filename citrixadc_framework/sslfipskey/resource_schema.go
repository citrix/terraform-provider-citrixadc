package sslfipskey

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslfipskeyResourceModel describes the resource data model.
type SslfipskeyResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Curve       types.String `tfsdk:"curve"`
	Exponent    types.String `tfsdk:"exponent"`
	Fipskeyname types.String `tfsdk:"fipskeyname"`
	Inform      types.String `tfsdk:"inform"`
	Iv          types.String `tfsdk:"iv"`
	Key         types.String `tfsdk:"key"`
	Keytype     types.String `tfsdk:"keytype"`
	Modulus     types.Int64  `tfsdk:"modulus"`
	Wrapkeyname types.String `tfsdk:"wrapkeyname"`
}

func (r *SslfipskeyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslfipskey resource.",
			},
			"curve": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("P_256"),
				Description: "Only p_256 (prime256v1) and P_384 (secp384r1) are supported.",
			},
			"exponent": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("3"),
				Description: "Exponent value for the FIPS key to be created. Available values function as follows:\n 3=3 (hexadecimal)\nF4=10001 (hexadecimal)",
			},
			"fipskeyname": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the FIPS key. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the FIPS key is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my fipskey\" or 'my fipskey').",
			},
			"inform": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("SIM"),
				Description: "Input format of the key file. Available formats are:\nSIM - Secure Information Management; select when importing a FIPS key. If the external FIPS key is encrypted, first decrypt it, and then import it.\nPEM - Privacy Enhanced Mail; select when importing a non-FIPS key.",
			},
			"iv": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Initialization Vector (IV) to use for importing the key. Required for importing a non-FIPS key.",
			},
			"key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the key file to be imported.\n /nsconfig/ssl/ is the default path.",
			},
			"keytype": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("RSA"),
				Description: "Only RSA key and ECDSA Key are supported.",
			},
			"modulus": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Modulus, in multiples of 64, of the FIPS key to be created.",
			},
			"wrapkeyname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the wrap key to use for importing the key. Required for importing a non-FIPS key.",
			},
		},
	}
}

func sslfipskeyGetThePayloadFromtheConfig(ctx context.Context, data *SslfipskeyResourceModel) ssl.Sslfipskey {
	tflog.Debug(ctx, "In sslfipskeyGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslfipskey := ssl.Sslfipskey{}
	if !data.Curve.IsNull() {
		sslfipskey.Curve = data.Curve.ValueString()
	}
	if !data.Exponent.IsNull() {
		sslfipskey.Exponent = data.Exponent.ValueString()
	}
	if !data.Fipskeyname.IsNull() {
		sslfipskey.Fipskeyname = data.Fipskeyname.ValueString()
	}
	if !data.Inform.IsNull() {
		sslfipskey.Inform = data.Inform.ValueString()
	}
	if !data.Iv.IsNull() {
		sslfipskey.Iv = data.Iv.ValueString()
	}
	if !data.Key.IsNull() {
		sslfipskey.Key = data.Key.ValueString()
	}
	if !data.Keytype.IsNull() {
		sslfipskey.Keytype = data.Keytype.ValueString()
	}
	if !data.Modulus.IsNull() {
		sslfipskey.Modulus = utils.IntPtr(int(data.Modulus.ValueInt64()))
	}
	if !data.Wrapkeyname.IsNull() {
		sslfipskey.Wrapkeyname = data.Wrapkeyname.ValueString()
	}

	return sslfipskey
}

func sslfipskeySetAttrFromGet(ctx context.Context, data *SslfipskeyResourceModel, getResponseData map[string]interface{}) *SslfipskeyResourceModel {
	tflog.Debug(ctx, "In sslfipskeySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["curve"]; ok && val != nil {
		data.Curve = types.StringValue(val.(string))
	} else {
		data.Curve = types.StringNull()
	}
	if val, ok := getResponseData["exponent"]; ok && val != nil {
		data.Exponent = types.StringValue(val.(string))
	} else {
		data.Exponent = types.StringNull()
	}
	if val, ok := getResponseData["fipskeyname"]; ok && val != nil {
		data.Fipskeyname = types.StringValue(val.(string))
	} else {
		data.Fipskeyname = types.StringNull()
	}
	if val, ok := getResponseData["inform"]; ok && val != nil {
		data.Inform = types.StringValue(val.(string))
	} else {
		data.Inform = types.StringNull()
	}
	if val, ok := getResponseData["iv"]; ok && val != nil {
		data.Iv = types.StringValue(val.(string))
	} else {
		data.Iv = types.StringNull()
	}
	if val, ok := getResponseData["key"]; ok && val != nil {
		data.Key = types.StringValue(val.(string))
	} else {
		data.Key = types.StringNull()
	}
	if val, ok := getResponseData["keytype"]; ok && val != nil {
		data.Keytype = types.StringValue(val.(string))
	} else {
		data.Keytype = types.StringNull()
	}
	if val, ok := getResponseData["modulus"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Modulus = types.Int64Value(intVal)
		}
	} else {
		data.Modulus = types.Int64Null()
	}
	if val, ok := getResponseData["wrapkeyname"]; ok && val != nil {
		data.Wrapkeyname = types.StringValue(val.(string))
	} else {
		data.Wrapkeyname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Fipskeyname.ValueString())

	return data
}
