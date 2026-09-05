package ssldhparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ssl"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
			// SDK v2: Required + ForceNew (TypeInt)
			"bits": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Size, in bits, of the DH key being generated.",
			},
			// SDK v2: Required + ForceNew (TypeString). This value is the resource ID.
			"dhfile": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of and, optionally, path to the DH key file. /nsconfig/ssl/ is the default path.",
			},
			// SDK v2: Optional + Computed + ForceNew (no Default declared in SDK v2).
			"gen": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Random number required for generating the DH key. Required as part of the DH key generation algorithm.",
			},
		},
	}
}

func ssldhparamGetThePayloadFromtheConfig(ctx context.Context, data *SsldhparamResourceModel) ssl.Ssldhparam {
	tflog.Debug(ctx, "In ssldhparamGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	ssldhparam := ssl.Ssldhparam{}
	if !data.Bits.IsNull() && !data.Bits.IsUnknown() {
		ssldhparam.Bits = utils.IntPtr(int(data.Bits.ValueInt64()))
	}
	if !data.Dhfile.IsNull() && !data.Dhfile.IsUnknown() {
		ssldhparam.Dhfile = data.Dhfile.ValueString()
	}
	if !data.Gen.IsNull() && !data.Gen.IsUnknown() {
		ssldhparam.Gen = data.Gen.ValueString()
	}

	return ssldhparam
}
