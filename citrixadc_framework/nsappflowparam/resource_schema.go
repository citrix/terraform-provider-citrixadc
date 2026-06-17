package nsappflowparam

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NsappflowparamResourceModel describes the resource data model.
type NsappflowparamResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Clienttrafficonly types.String `tfsdk:"clienttrafficonly"`
	Httpcookie        types.String `tfsdk:"httpcookie"`
	Httphost          types.String `tfsdk:"httphost"`
	Httpmethod        types.String `tfsdk:"httpmethod"`
	Httpreferer       types.String `tfsdk:"httpreferer"`
	Httpurl           types.String `tfsdk:"httpurl"`
	Httpuseragent     types.String `tfsdk:"httpuseragent"`
	Templaterefresh   types.Int64  `tfsdk:"templaterefresh"`
	Udppmtu           types.Int64  `tfsdk:"udppmtu"`
}

func (r *NsappflowparamResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsappflowparam resource.",
			},
			"clienttrafficonly": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control whether AppFlow records should be generated only for client-side traffic.",
			},
			"httpcookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP cookie logging.",
			},
			"httphost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP host logging.",
			},
			"httpmethod": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP method logging.",
			},
			"httpreferer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP referer logging.",
			},
			"httpurl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP URL logging.",
			},
			"httpuseragent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable AppFlow HTTP user-agent logging.",
			},
			"templaterefresh": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "IPFIX template refresh interval (in seconds).",
			},
			"udppmtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU to be used for IPFIX UDP packets.",
			},
		},
	}
}

func nsappflowparamGetThePayloadFromthePlan(ctx context.Context, data *NsappflowparamResourceModel) ns.Nsappflowparam {
	tflog.Debug(ctx, "In nsappflowparamGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsappflowparam := ns.Nsappflowparam{}
	if !data.Clienttrafficonly.IsNull() && !data.Clienttrafficonly.IsUnknown() {
		nsappflowparam.Clienttrafficonly = data.Clienttrafficonly.ValueString()
	}
	if !data.Httpcookie.IsNull() && !data.Httpcookie.IsUnknown() {
		nsappflowparam.Httpcookie = data.Httpcookie.ValueString()
	}
	if !data.Httphost.IsNull() && !data.Httphost.IsUnknown() {
		nsappflowparam.Httphost = data.Httphost.ValueString()
	}
	if !data.Httpmethod.IsNull() && !data.Httpmethod.IsUnknown() {
		nsappflowparam.Httpmethod = data.Httpmethod.ValueString()
	}
	if !data.Httpreferer.IsNull() && !data.Httpreferer.IsUnknown() {
		nsappflowparam.Httpreferer = data.Httpreferer.ValueString()
	}
	if !data.Httpurl.IsNull() && !data.Httpurl.IsUnknown() {
		nsappflowparam.Httpurl = data.Httpurl.ValueString()
	}
	if !data.Httpuseragent.IsNull() && !data.Httpuseragent.IsUnknown() {
		nsappflowparam.Httpuseragent = data.Httpuseragent.ValueString()
	}
	if !data.Templaterefresh.IsNull() && !data.Templaterefresh.IsUnknown() {
		nsappflowparam.Templaterefresh = utils.IntPtr(int(data.Templaterefresh.ValueInt64()))
	}
	if !data.Udppmtu.IsNull() && !data.Udppmtu.IsUnknown() {
		nsappflowparam.Udppmtu = utils.IntPtr(int(data.Udppmtu.ValueInt64()))
	}

	return nsappflowparam
}

// nsappflowparamSetAttrFromGet populates the resource model from the GET response.
// This is a settable singleton: the attributes are Optional+Computed and the GET
// (get-all) response always echoes the server-applied values (or defaults), so we
// faithfully copy them when present. The synthetic ID is set exactly once in Create
// (Pattern 6), so it is NOT recomputed here.
func nsappflowparamSetAttrFromGet(ctx context.Context, data *NsappflowparamResourceModel, getResponseData map[string]interface{}) *NsappflowparamResourceModel {
	tflog.Debug(ctx, "In nsappflowparamSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["clienttrafficonly"]; ok && val != nil {
		data.Clienttrafficonly = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httpcookie"]; ok && val != nil {
		data.Httpcookie = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httphost"]; ok && val != nil {
		data.Httphost = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httpreferer"]; ok && val != nil {
		data.Httpreferer = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httpurl"]; ok && val != nil {
		data.Httpurl = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["httpuseragent"]; ok && val != nil {
		data.Httpuseragent = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["templaterefresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Templaterefresh = types.Int64Value(intVal)
		}
	}
	if val, ok := getResponseData["udppmtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Udppmtu = types.Int64Value(intVal)
		}
	}

	return data
}

// nsappflowparamSetAttrFromGetForDatasource faithfully copies every field from the GET
// response and sets the synthetic ID, because the datasource never calls Create
// (Pattern 7).
func nsappflowparamSetAttrFromGetForDatasource(ctx context.Context, data *NsappflowparamResourceModel, getResponseData map[string]interface{}) *NsappflowparamResourceModel {
	tflog.Debug(ctx, "In nsappflowparamSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["clienttrafficonly"]; ok && val != nil {
		data.Clienttrafficonly = types.StringValue(val.(string))
	} else {
		data.Clienttrafficonly = types.StringNull()
	}
	if val, ok := getResponseData["httpcookie"]; ok && val != nil {
		data.Httpcookie = types.StringValue(val.(string))
	} else {
		data.Httpcookie = types.StringNull()
	}
	if val, ok := getResponseData["httphost"]; ok && val != nil {
		data.Httphost = types.StringValue(val.(string))
	} else {
		data.Httphost = types.StringNull()
	}
	if val, ok := getResponseData["httpmethod"]; ok && val != nil {
		data.Httpmethod = types.StringValue(val.(string))
	} else {
		data.Httpmethod = types.StringNull()
	}
	if val, ok := getResponseData["httpreferer"]; ok && val != nil {
		data.Httpreferer = types.StringValue(val.(string))
	} else {
		data.Httpreferer = types.StringNull()
	}
	if val, ok := getResponseData["httpurl"]; ok && val != nil {
		data.Httpurl = types.StringValue(val.(string))
	} else {
		data.Httpurl = types.StringNull()
	}
	if val, ok := getResponseData["httpuseragent"]; ok && val != nil {
		data.Httpuseragent = types.StringValue(val.(string))
	} else {
		data.Httpuseragent = types.StringNull()
	}
	if val, ok := getResponseData["templaterefresh"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Templaterefresh = types.Int64Value(intVal)
		}
	} else {
		data.Templaterefresh = types.Int64Null()
	}
	if val, ok := getResponseData["udppmtu"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Udppmtu = types.Int64Value(intVal)
		}
	} else {
		data.Udppmtu = types.Int64Null()
	}

	// Datasource has no Create, so set the synthetic ID here.
	data.Id = types.StringValue("nsappflowparam-config")

	return data
}
