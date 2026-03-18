package ssldhparam

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

// SsldhparamResourceModel describes the resource data model.
type SsldhparamResourceModel struct {
	Id     types.String `tfsdk:"id"`
	Bits   types.Int64  `tfsdk:"bits"`
	Dhfile types.String `tfsdk:"dhfile"`
	Gen    types.String `tfsdk:"gen"`
}

func (r *SsldhparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the ssldhparam resource.",
			},
			"bits": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Size, in bits, of the DH key being generated.",
			},
			"dhfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the DH key file. /nsconfig/ssl/ is the default path.",
			},
			"gen": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Default:     stringdefault.StaticString("2"),
				Description: "Random number required for generating the DH key. Required as part of the DH key generation algorithm.",
			},
		},
	}
}

func ssldhparamGetThePayloadFromtheConfig(ctx context.Context, data *SsldhparamResourceModel) ssl.Ssldhparam {
	tflog.Debug(ctx, "In ssldhparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ssldhparam := ssl.Ssldhparam{}
	if !data.Bits.IsNull() {
		ssldhparam.Bits = utils.IntPtr(int(data.Bits.ValueInt64()))
	}
	if !data.Dhfile.IsNull() {
		ssldhparam.Dhfile = data.Dhfile.ValueString()
	}
	if !data.Gen.IsNull() {
		ssldhparam.Gen = data.Gen.ValueString()
	}

	return ssldhparam
}

func ssldhparamSetAttrFromGet(ctx context.Context, data *SsldhparamResourceModel, getResponseData map[string]interface{}) *SsldhparamResourceModel {
	tflog.Debug(ctx, "In ssldhparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bits"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bits = types.Int64Value(intVal)
		}
	} else {
		data.Bits = types.Int64Null()
	}
	if val, ok := getResponseData["dhfile"]; ok && val != nil {
		data.Dhfile = types.StringValue(val.(string))
	} else {
		data.Dhfile = types.StringNull()
	}
	if val, ok := getResponseData["gen"]; ok && val != nil {
		data.Gen = types.StringValue(val.(string))
	} else {
		data.Gen = types.StringNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("ssldhparam-config")

	return data
}
