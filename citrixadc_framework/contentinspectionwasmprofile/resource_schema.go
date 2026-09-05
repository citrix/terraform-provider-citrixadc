package contentinspectionwasmprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/contentinspection"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while the prior state
// holds a value other than the NITRO default. This makes Terraform detect a
// change (unknown != prior) and call Update, which issues the NITRO
// ?action=unset -- without it an Optional+Computed attribute is "sticky" and
// removal is a silent no-op. It does nothing once the prior state already equals
// the default, avoiding a perpetual post-unset plan diff (NITRO always echoes
// the default back on GET for these attributes).
type unsetOnRemoveStringModifier struct{ defaultValue string }

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while prior state differs from the default, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" && req.StateValue.ValueString() != m.defaultValue {
		resp.PlanValue = types.StringUnknown()
	}
}

// unsetOnRemoveInt64Modifier is the Int64 counterpart of unsetOnRemoveStringModifier.
type unsetOnRemoveInt64Modifier struct{ defaultValue int64 }

func (m unsetOnRemoveInt64Modifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while prior state differs from the default, so it is unset on the appliance."
}

func (m unsetOnRemoveInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveInt64Modifier) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueInt64() != m.defaultValue {
		resp.PlanValue = types.Int64Unknown()
	}
}

// ContentinspectionwasmprofileResourceModel describes the resource data model.
type ContentinspectionwasmprofileResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Anomalousdatasize types.Int64  `tfsdk:"anomalousdatasize"`
	Anomalousttfbtime types.Int64  `tfsdk:"anomalousttfbtime"`
	Maxbodylen        types.Int64  `tfsdk:"maxbodylen"`
	Name              types.String `tfsdk:"name"`
	Timeout           types.Int64  `tfsdk:"timeout"`
	Timeoutaction     types.String `tfsdk:"timeoutaction"`
	Wasmmodule        types.String `tfsdk:"wasmmodule"`
}

func (r *ContentinspectionwasmprofileResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionwasmprofile resource.",
			},
			"anomalousdatasize": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{defaultValue: 512},
				},
				Description: "Transaction data size (in KB) greater than which a transaction is considered as anomalous. Default is 512KB.",
			},
			"anomalousttfbtime": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{defaultValue: 1000},
				},
				Description: "Transaction time (in milliseconds) above which a transaction is considered as anomalous. Default is 1 seconds.",
			},
			"maxbodylen": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{defaultValue: 16},
				},
				Description: "Max data size (in KB) that will be sent to the CI Agent. Default is 16KB. Maximum value that can be configured is 32KB.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of CI WASM profile.",
			},
			"timeout": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{defaultValue: 1000},
				},
				Description: "Timeout (in milliseconds) for the connection with the CI WASM agent.",
			},
			"timeoutaction": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: "DROP"},
				},
				Description: "Timeout action for the connection with the CI agent. Either the original request can be bypassed i.e. request/response is forwarded to the endpoint or the transaction is dropped/reset. Possible values = BYPASS, DROP, RESET",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM Module.",
			},
		},
	}
}

func contentinspectionwasmprofileGetThePayloadFromthePlan(ctx context.Context, data *ContentinspectionwasmprofileResourceModel) contentinspection.Contentinspectionwasmprofile {
	tflog.Debug(ctx, "In contentinspectionwasmprofileGetThePayloadFromthePlan Function")

	// Create API request body from the model
	contentinspectionwasmprofile := contentinspection.Contentinspectionwasmprofile{}
	if !data.Anomalousdatasize.IsNull() && !data.Anomalousdatasize.IsUnknown() {
		contentinspectionwasmprofile.Anomalousdatasize = utils.IntPtr(int(data.Anomalousdatasize.ValueInt64()))
	}
	if !data.Anomalousttfbtime.IsNull() && !data.Anomalousttfbtime.IsUnknown() {
		contentinspectionwasmprofile.Anomalousttfbtime = utils.IntPtr(int(data.Anomalousttfbtime.ValueInt64()))
	}
	if !data.Maxbodylen.IsNull() && !data.Maxbodylen.IsUnknown() {
		contentinspectionwasmprofile.Maxbodylen = utils.IntPtr(int(data.Maxbodylen.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		contentinspectionwasmprofile.Name = data.Name.ValueString()
	}
	if !data.Timeout.IsNull() && !data.Timeout.IsUnknown() {
		contentinspectionwasmprofile.Timeout = utils.IntPtr(int(data.Timeout.ValueInt64()))
	}
	if !data.Timeoutaction.IsNull() && !data.Timeoutaction.IsUnknown() {
		contentinspectionwasmprofile.Timeoutaction = data.Timeoutaction.ValueString()
	}
	if !data.Wasmmodule.IsNull() && !data.Wasmmodule.IsUnknown() {
		contentinspectionwasmprofile.Wasmmodule = data.Wasmmodule.ValueString()
	}

	return contentinspectionwasmprofile
}

func contentinspectionwasmprofileSetAttrFromGet(ctx context.Context, data *ContentinspectionwasmprofileResourceModel, getResponseData map[string]interface{}) *ContentinspectionwasmprofileResourceModel {
	tflog.Debug(ctx, "In contentinspectionwasmprofileSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["anomalousdatasize"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anomalousdatasize = types.Int64Value(intVal)
		}
	} else {
		data.Anomalousdatasize = types.Int64Null()
	}
	if val, ok := getResponseData["anomalousttfbtime"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Anomalousttfbtime = types.Int64Value(intVal)
		}
	} else {
		data.Anomalousttfbtime = types.Int64Null()
	}
	if val, ok := getResponseData["maxbodylen"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxbodylen = types.Int64Value(intVal)
		}
	} else {
		data.Maxbodylen = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["timeout"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Timeout = types.Int64Value(intVal)
		}
	} else {
		data.Timeout = types.Int64Null()
	}
	if val, ok := getResponseData["timeoutaction"]; ok && val != nil {
		data.Timeoutaction = types.StringValue(val.(string))
	} else {
		data.Timeoutaction = types.StringNull()
	}
	if val, ok := getResponseData["wasmmodule"]; ok && val != nil {
		data.Wasmmodule = types.StringValue(val.(string))
	} else {
		data.Wasmmodule = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
