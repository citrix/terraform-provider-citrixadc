package aaaotpparameter

import (
	"context"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AaaotpparameterDataSourceModel is the data-source-specific model, decoupled
// from AaaotpparameterResourceModel. A data source is a pure read surface, so it
// can expose the FULL GET projection: the read/write attributes (as Computed
// outputs) AND the read-only attributes the resource deliberately omits.
type AaaotpparameterDataSourceModel struct {
	Id            types.String `tfsdk:"id"`
	Encryption    types.String `tfsdk:"encryption"`
	Maxotpdevices types.Int64  `tfsdk:"maxotpdevices"`
	Otptype       types.String `tfsdk:"otptype"`

	// Read-only (GET-only) attributes from the NITRO read-only set
	// (zion73x_readonly/aaaotpparameter.json). Never settable; populated from GET.
	Gwtestchallenge types.String `tfsdk:"gwtestchallenge"`
}

func AaaotpparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "To encrypt otp secret in AD or not. Default value is OFF",
			},
			"maxotpdevices": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of otp devices user can register. Default value is 4. Max value is 255",
			},
			"otptype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Input flag to generate OTP for the given type. Possible values = gwtest",
			},

			// Read-only (GET-only) attributes surfaced by the data source. All Computed.
			"gwtestchallenge": schema.StringAttribute{
				Computed:    true,
				Description: "Holds the generated OTP to access gwtest admin flow.",
			},
		},
	}
}

// aaaotpparameterDataSourceSetAttrFromGet projects a NITRO aaaotpparameter GET
// response onto the data-source model. Attributes are simply filled from the
// GET (or left Null when the GET omits them) via the shared utils.MapGet*
// helpers.
func aaaotpparameterDataSourceSetAttrFromGet(ctx context.Context, data *AaaotpparameterDataSourceModel, g map[string]interface{}) {
	tflog.Debug(ctx, "In aaaotpparameterDataSourceSetAttrFromGet Function")

	// Read/write attributes as read-back outputs.
	data.Encryption = utils.MapGetString(g, "encryption")
	data.Maxotpdevices = utils.MapGetInt64(g, "maxotpdevices")
	data.Otptype = utils.MapGetString(g, "otptype")

	// Read-only attributes.
	data.Gwtestchallenge = utils.MapGetString(g, "gwtestchallenge")

	// aaaotpparameter is a singleton; the ID is a fixed system-generated identifier.
	data.Id = types.StringValue("aaaotpparameter-config")
}
