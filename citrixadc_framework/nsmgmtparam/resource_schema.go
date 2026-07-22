package nsmgmtparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsmgmtparamResourceModel describes the resource data model.
type NsmgmtparamResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Httpdmaxclients    types.Int64  `tfsdk:"httpdmaxclients"`
	Httpdmaxreqworkers types.Int64  `tfsdk:"httpdmaxreqworkers"`
	Mgmthttpport       types.Int64  `tfsdk:"mgmthttpport"`
	Mgmthttpsport      types.Int64  `tfsdk:"mgmthttpsport"`
}

func (r *NsmgmtparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmgmtparam resource.",
			},
			"httpdmaxclients": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This enables setting the HTTPD Max Clients value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255",
			},
			"httpdmaxreqworkers": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This enables setting the HTTPD Max Request Workers value in the httpd.conf file. You can configure either Max Clients or Max Request Workers. The allowable range is from a minimum of 1 to a maximum of 255",
			},
			"mgmthttpport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allow the configuration of management HTTP port.",
			},
			"mgmthttpsport": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "This allows the configuration of management HTTPS port.",
			},
		},
	}
}

func nsmgmtparamGetThePayloadFromthePlan(ctx context.Context, data *NsmgmtparamResourceModel) ns.Nsmgmtparam {
	tflog.Debug(ctx, "In nsmgmtparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsmgmtparam := ns.Nsmgmtparam{}
	if !data.Httpdmaxclients.IsNull() && !data.Httpdmaxclients.IsUnknown() {
		nsmgmtparam.Httpdmaxclients = utils.IntPtr(int(data.Httpdmaxclients.ValueInt64()))
	}
	if !data.Httpdmaxreqworkers.IsNull() && !data.Httpdmaxreqworkers.IsUnknown() {
		nsmgmtparam.Httpdmaxreqworkers = utils.IntPtr(int(data.Httpdmaxreqworkers.ValueInt64()))
	}
	if !data.Mgmthttpport.IsNull() && !data.Mgmthttpport.IsUnknown() {
		nsmgmtparam.Mgmthttpport = utils.IntPtr(int(data.Mgmthttpport.ValueInt64()))
	}
	if !data.Mgmthttpsport.IsNull() && !data.Mgmthttpsport.IsUnknown() {
		nsmgmtparam.Mgmthttpsport = utils.IntPtr(int(data.Mgmthttpsport.ValueInt64()))
	}

	return nsmgmtparam
}

// nsmgmtparamSetAttrFromGet populates the resource model from the GET response.
// This is a settable singleton: the attributes are Optional+Computed and the GET
// (get-all) response always echoes the server-applied values (or defaults), so we
// faithfully copy them when present. The synthetic ID is set exactly once in Create
// (Pattern 6), so it is NOT recomputed here.
func nsmgmtparamSetAttrFromGet(ctx context.Context, data *NsmgmtparamResourceModel, getResponseData map[string]interface{}) *NsmgmtparamResourceModel {
	tflog.Debug(ctx, "In nsmgmtparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["httpdmaxclients"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpdmaxclients = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["httpdmaxreqworkers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpdmaxreqworkers = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["mgmthttpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpport = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["mgmthttpsport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpsport = types.Int64Value(intVal)
		}
	}

	return data
}

// nsmgmtparamSetAttrFromGetForDatasource faithfully copies every field from the GET
// response and sets the synthetic ID, because the datasource never calls Create
// (Pattern 7).
func nsmgmtparamSetAttrFromGetForDatasource(ctx context.Context, data *NsmgmtparamResourceModel, getResponseData map[string]interface{}) *NsmgmtparamResourceModel {
	tflog.Debug(ctx, "In nsmgmtparamSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["httpdmaxclients"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpdmaxclients = types.Int64Value(intVal)
		}
	} else {
		data.Httpdmaxclients = types.Int64Null()
	}
	if val, ok := getResponseData["httpdmaxreqworkers"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Httpdmaxreqworkers = types.Int64Value(intVal)
		}
	} else {
		data.Httpdmaxreqworkers = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpport = types.Int64Value(intVal)
		}
	} else {
		data.Mgmthttpport = types.Int64Null()
	}
	if val, ok := getResponseData["mgmthttpsport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Mgmthttpsport = types.Int64Value(intVal)
		}
	} else {
		data.Mgmthttpsport = types.Int64Null()
	}

	// Datasource has no Create, so set the synthetic ID here.
	data.Id = types.StringValue("nsmgmtparam-config")

	return data
}
