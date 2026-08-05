package nscapacity

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
				Optional:    true,
				Computed:    true,
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

	// Create API request body from the model.
	// Guard against IsUnknown(): during Create the Optional+Computed attributes
	// that the user did not configure are UNKNOWN (not null), and reading a value
	// out of an unknown yields a zero value (0 / "" / false) which must NOT be
	// pushed to the ADC. SDK v2 achieved the same via GetRawConfig() null checks.
	nscapacity := ns.Nscapacity{}
	if !data.Bandwidth.IsNull() && !data.Bandwidth.IsUnknown() {
		nscapacity.Bandwidth = utils.IntPtr(int(data.Bandwidth.ValueInt64()))
	}
	if !data.Edition.IsNull() && !data.Edition.IsUnknown() {
		nscapacity.Edition = data.Edition.ValueString()
	}
	if !data.Nodeid.IsNull() && !data.Nodeid.IsUnknown() {
		nscapacity.Nodeid = utils.IntPtr(int(data.Nodeid.ValueInt64()))
	}
	if !data.Password.IsNull() && !data.Password.IsUnknown() {
		nscapacity.Password = data.Password.ValueString()
	}
	if !data.Platform.IsNull() && !data.Platform.IsUnknown() {
		nscapacity.Platform = data.Platform.ValueString()
	}
	if !data.Unit.IsNull() && !data.Unit.IsUnknown() {
		nscapacity.Unit = data.Unit.ValueString()
	}
	if !data.Username.IsNull() && !data.Username.IsUnknown() {
		nscapacity.Username = data.Username.ValueString()
	}
	if !data.Vcpu.IsNull() && !data.Vcpu.IsUnknown() {
		nscapacity.Vcpu = data.Vcpu.ValueBool()
	}

	return nscapacity
}

func nscapacitySetAttrFromGet(ctx context.Context, data *NscapacityResourceModel, getResponseData map[string]interface{}) *NscapacityResourceModel {
	tflog.Debug(ctx, "In nscapacitySetAttrFromGet Function")

	// Convert API response to model.
	//
	// This mirrors the SDK v2 read semantics (createNscapacityFunc/readNscapacityFunc):
	//   - "Pooled" license: bandwidth/edition/unit come back together;
	//   - "CICO" license: platform;
	//   - "vCPU" license: derived from the presence of the read-only "vcpucount" key
	//     (the ADC never returns a "vcpu" field on GET).
	//
	// Every else-branch is guarded with IsUnknown() (the omit-on-default trap): when
	// NITRO omits a key from GET we only reset the attribute to a concrete value when
	// the incoming value is still unknown (i.e. it was Computed and not user-set).
	// A known, user-configured value is preserved so we never produce a
	// "provider produced inconsistent result after apply" error, and never clobber
	// a configured 0/false/"" that the ADC simply does not echo back.

	// Pooled license: bandwidth / edition / unit
	if val, ok := getResponseData["bandwidth"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Bandwidth = types.Int64Value(intVal)
		}
	} else if data.Bandwidth.IsUnknown() {
		data.Bandwidth = types.Int64Null()
	}
	if val, ok := getResponseData["edition"]; ok && val != nil {
		data.Edition = types.StringValue(val.(string))
	} else if data.Edition.IsUnknown() {
		data.Edition = types.StringNull()
	}
	if val, ok := getResponseData["unit"]; ok && val != nil {
		data.Unit = types.StringValue(val.(string))
	} else if data.Unit.IsUnknown() {
		data.Unit = types.StringNull()
	}

	// CICO license: platform
	if val, ok := getResponseData["platform"]; ok && val != nil {
		data.Platform = types.StringValue(val.(string))
	} else if data.Platform.IsUnknown() {
		data.Platform = types.StringNull()
	}

	// vCPU license: derived from read-only "vcpucount" presence
	if _, ok := getResponseData["vcpucount"]; ok {
		data.Vcpu = types.BoolValue(true)
	} else if data.Vcpu.IsUnknown() {
		data.Vcpu = types.BoolValue(false)
	}

	// nodeid (cluster node id)
	if val, ok := getResponseData["nodeid"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Nodeid = types.Int64Value(intVal)
		}
	} else if data.Nodeid.IsUnknown() {
		data.Nodeid = types.Int64Null()
	}

	// username / password (ADM Agent credentials; not echoed back by GET)
	if val, ok := getResponseData["username"]; ok && val != nil {
		data.Username = types.StringValue(val.(string))
	} else if data.Username.IsUnknown() {
		data.Username = types.StringNull()
	}
	if val, ok := getResponseData["password"]; ok && val != nil {
		data.Password = types.StringValue(val.(string))
	} else if data.Password.IsUnknown() {
		data.Password = types.StringNull()
	}

	// Set ID for the resource
	// Singleton resource - static ID
	data.Id = types.StringValue("nscapacity-config")

	return data
}
