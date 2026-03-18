package sslcipher

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// SslcipherResourceModel describes the resource data model.
type SslcipherResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Ciphergroupname types.String `tfsdk:"ciphergroupname"`
	Ciphername      types.String `tfsdk:"ciphername"`
	Cipherpriority  types.Int64  `tfsdk:"cipherpriority"`
	Ciphgrpalias    types.String `tfsdk:"ciphgrpalias"`
	Sslprofile      types.String `tfsdk:"sslprofile"`
}

func (r *SslcipherResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the sslcipher resource.",
			},
			"ciphergroupname": schema.StringAttribute{
				Required:    true,
				Description: "Name for the user-defined cipher group. Must begin with an ASCII alphanumeric or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters. Cannot be changed after the cipher group is created.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ciphergroup\" or 'my ciphergroup').",
			},
			"ciphername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cipher name.",
			},
			"cipherpriority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This indicates priority assigned to the particular cipher",
			},
			"ciphgrpalias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "The individual cipher name(s), a user-defined cipher group, or a system predefined cipher alias that will be added to the  predefined cipher alias that will be added to the group cipherGroupName.\nIf a cipher alias or a cipher group is specified, all the individual ciphers in the cipher alias or group will be added to the user-defined cipher group.",
			},
			"sslprofile": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which cipher is attached.",
			},
		},
	}
}

func sslcipherGetThePayloadFromtheConfig(ctx context.Context, data *SslcipherResourceModel) ssl.Sslcipher {
	tflog.Debug(ctx, "In sslcipherGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	sslcipher := ssl.Sslcipher{}
	if !data.Ciphergroupname.IsNull() {
		sslcipher.Ciphergroupname = data.Ciphergroupname.ValueString()
	}
	if !data.Ciphername.IsNull() {
		sslcipher.Ciphername = data.Ciphername.ValueString()
	}
	if !data.Cipherpriority.IsNull() {
		sslcipher.Cipherpriority = utils.IntPtr(int(data.Cipherpriority.ValueInt64()))
	}
	if !data.Ciphgrpalias.IsNull() {
		sslcipher.Ciphgrpalias = data.Ciphgrpalias.ValueString()
	}
	if !data.Sslprofile.IsNull() {
		sslcipher.Sslprofile = data.Sslprofile.ValueString()
	}

	return sslcipher
}

func sslcipherSetAttrFromGet(ctx context.Context, data *SslcipherResourceModel, getResponseData map[string]interface{}) *SslcipherResourceModel {
	tflog.Debug(ctx, "In sslcipherSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["ciphergroupname"]; ok && val != nil {
		data.Ciphergroupname = types.StringValue(val.(string))
	} else {
		data.Ciphergroupname = types.StringNull()
	}
	if val, ok := getResponseData["ciphername"]; ok && val != nil {
		data.Ciphername = types.StringValue(val.(string))
	} else {
		data.Ciphername = types.StringNull()
	}
	if val, ok := getResponseData["cipherpriority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Cipherpriority = types.Int64Value(intVal)
		}
	} else {
		data.Cipherpriority = types.Int64Null()
	}
	if val, ok := getResponseData["ciphgrpalias"]; ok && val != nil {
		data.Ciphgrpalias = types.StringValue(val.(string))
	} else {
		data.Ciphgrpalias = types.StringNull()
	}
	if val, ok := getResponseData["sslprofile"]; ok && val != nil {
		data.Sslprofile = types.StringValue(val.(string))
	} else {
		data.Sslprofile = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated
	data.Id = types.StringValue(fmt.Sprintf("%s", data.Ciphergroupname.ValueString()))

	return data
}
