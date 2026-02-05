package nscapacity

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// NscapacityResourceModel describes the resource data model.
type NscapacityResourceModel struct {
	Id        types.String `tfsdk:"id"`
	Bandwidth types.Int64  `tfsdk:"bandwidth"`
	Edition   types.String `tfsdk:"edition"`
	Nodeid    types.Int64  `tfsdk:"nodeid"`
	Password  types.String `tfsdk:"password"`
	Platform  types.String `tfsdk:"platform"`
	Unit      types.String `tfsdk:"unit"`
	Username  types.String `tfsdk:"username"`
	Vcpu      types.Bool   `tfsdk:"vcpu"`
}

func (r *NscapacityResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nscapacity resource.",
			},
			"bandwidth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "System bandwidth limit.",
			},
			"edition": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Product edition.",
			},
			"nodeid": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Password to use when authenticating with ADM Agent for LAS licensing.",
			},
			"platform": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "appliance platform type.",
			},
			"unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bandwidth unit.",
			},
			"username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username to authenticate with ADM Agent for LAS licensing. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"vcpu": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "licensed using vcpu pool.",
			},
		},
	}
}

func nscapacityGetThePayloadFromtheConfig(ctx context.Context, data *NscapacityResourceModel) ns.Nscapacity {
	tflog.Debug(ctx, "In nscapacityGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nscapacity := ns.Nscapacity{}
	if !data.Bandwidth.IsNull() {
		nscapacity.Bandwidth = utils.IntPtr(int(data.Bandwidth.ValueInt64()))
	}
	if !data.Edition.IsNull() {
		nscapacity.Edition = data.Edition.ValueString()
	}
	if !data.Nodeid.IsNull() {
		nscapacity.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Password.IsNull() {
		nscapacity.Password = data.Password.ValueString()
	}
	if !data.Platform.IsNull() {
		nscapacity.Platform = data.Platform.ValueString()
	}
	if !data.Unit.IsNull() {
		nscapacity.Unit = data.Unit.ValueString()
	}
	if !data.Username.IsNull() {
		nscapacity.Username = data.Username.ValueString()
	}
	if !data.Vcpu.IsNull() {
		nscapacity.Vcpu = data.Vcpu.ValueBool()
	}

	return nscapacity
}

func nscapacitySetAttrFromGet(ctx context.Context, data *NscapacityResourceModel, getResponseData map[string]interface{}) *NscapacityResourceModel {
	tflog.Debug(ctx, "In nscapacitySetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bandwidth = types.Int64Value(intVal)
		}
	} else {
		data.Bandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["edition"]; ok && val != nil {
		data.Edition = types.StringValue(val.(string))
	} else {
		data.Edition = types.StringNull()
	}
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else {
		data.Nodeid = types.Int64Null()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else {
		data.Password = types.StringNull()
	}
	if val, ok := getResponseData["platform"]; ok && val != nil {
		data.Platform = types.StringValue(val.(string))
	} else {
		data.Platform = types.StringNull()
	}
	if val, ok := getResponseData["unit"]; ok && val != nil {
		data.Unit = types.StringValue(val.(string))
	} else {
		data.Unit = types.StringNull()
	}
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else {
		data.Username = types.StringNull()
	}
	if val, ok := getResponseData["vcpu"]; ok && val != nil {
		data.Vcpu = types.BoolValue(val.(bool))
	} else {
		data.Vcpu = types.BoolNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nscapacity-config")

	return data
}
