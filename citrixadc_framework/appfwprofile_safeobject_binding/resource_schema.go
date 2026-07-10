package appfwprofile_safeobject_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileSafeobjectBindingResourceModel describes the resource data model.
type AppfwprofileSafeobjectBindingResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Action         types.List   `tfsdk:"action"`
	Alertonly      types.String `tfsdk:"alertonly"`
	AsExpression   types.String `tfsdk:"as_expression"`
	Comment        types.String `tfsdk:"comment"`
	Isautodeployed types.String `tfsdk:"isautodeployed"`
	Maxmatchlength types.Int64  `tfsdk:"maxmatchlength"`
	Name           types.String `tfsdk:"name"`
	Resourceid     types.String `tfsdk:"resourceid"`
	Ruletype       types.String `tfsdk:"ruletype"`
	Safeobject     types.String `tfsdk:"safeobject"`
	State          types.String `tfsdk:"state"`
}

func (r *AppfwprofileSafeobjectBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_safeobject_binding resource.",
			},
			"action": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Description: "Safe Object action types. (BLOCK | LOG | STATS | NONE)",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_expression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A regular expression that defines the Safe Object.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Any comments about the purpose of profile, or other useful information about the profile.",
			},
			"isautodeployed": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Is the rule auto deployed by dynamic profile ?",
			},
			"maxmatchlength": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Maximum match length for a Safe Object expression.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the profile to which to bind an exemption or rule.",
			},
			"resourceid": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "A \"id\" that identifies the rule.",
			},
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
			"safeobject": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the Safe Object.",
			},
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
		},
	}
}

func appfwprofile_safeobject_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileSafeobjectBindingResourceModel) appfw.Appfwprofilesafeobjectbinding {
	tflog.Debug(ctx, "In appfwprofile_safeobject_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_safeobject_binding := appfw.Appfwprofilesafeobjectbinding{}
	if !data.Action.IsNull() && !data.Action.IsUnknown() {
		var actionList []string
		data.Action.ElementsAs(ctx, &actionList, false)
		appfwprofile_safeobject_binding.Action = actionList
	}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_safeobject_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsExpression.IsNull() && !data.AsExpression.IsUnknown() {
		appfwprofile_safeobject_binding.Asexpression = data.AsExpression.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_safeobject_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_safeobject_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Maxmatchlength.IsNull() && !data.Maxmatchlength.IsUnknown() {
		appfwprofile_safeobject_binding.Maxmatchlength = utils.IntPtr(int(data.Maxmatchlength.ValueInt64()))
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_safeobject_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_safeobject_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_safeobject_binding.Ruletype = data.Ruletype.ValueString()
	}
	if !data.Safeobject.IsNull() && !data.Safeobject.IsUnknown() {
		appfwprofile_safeobject_binding.Safeobject = data.Safeobject.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_safeobject_binding.State = data.State.ValueString()
	}

	return appfwprofile_safeobject_binding
}

// Resource-side setter. Mirrors SDK v2 read behavior: alertonly and isautodeployed
// are server-overridden inputs that the NITRO GET does not faithfully echo back, so
// the SDK v2 resource deliberately did NOT set them from the GET response (the
// d.Set calls were commented out). Preserve the prior plan/state values for those
// two attributes here to avoid "inconsistent result after apply" diffs (Pattern 13 / e).
func appfwprofile_safeobject_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileSafeobjectBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileSafeobjectBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_safeobject_bindingSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Action = listValue
		} else {
			data.Action = types.ListNull(types.StringType)
		}
	} else {
		data.Action = types.ListNull(types.StringType)
	}
	// alertonly: server-overridden, not reliably echoed by GET. Preserve existing value.
	if val, ok := getResponseData["as_expression"]; ok && val != nil {
		data.AsExpression = types.StringValue(val.(string))
	} else {
		data.AsExpression = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	// isautodeployed: server-overridden, not reliably echoed by GET. Preserve existing value.
	if val, ok := getResponseData["maxmatchlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxmatchlength = types.Int64Value(intVal)
		}
	} else {
		data.Maxmatchlength = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["safeobject"]; ok && val != nil {
		data.Safeobject = types.StringValue(val.(string))
	} else {
		data.Safeobject = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	// Set ID for the resource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("safeobject:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Safeobject.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}

// Datasource-side setter (Pattern 7 split). Faithfully copies every field from the
// GET response, including alertonly and isautodeployed, and sets data.Id (datasource
// has no Create).
func appfwprofile_safeobject_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileSafeobjectBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileSafeobjectBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_safeobject_bindingSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["action"]; ok && val != nil {
		if sliceVal, ok := val.([]interface{}); ok {
			stringList := utils.ToStringList(sliceVal)
			listValue, _ := types.ListValueFrom(ctx, types.StringType, stringList)
			data.Action = listValue
		} else {
			data.Action = types.ListNull(types.StringType)
		}
	} else {
		data.Action = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["alertonly"]; ok && val != nil {
		data.Alertonly = types.StringValue(val.(string))
	} else {
		data.Alertonly = types.StringNull()
	}
	if val, ok := getResponseData["as_expression"]; ok && val != nil {
		data.AsExpression = types.StringValue(val.(string))
	} else {
		data.AsExpression = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["isautodeployed"]; ok && val != nil {
		data.Isautodeployed = types.StringValue(val.(string))
	} else {
		data.Isautodeployed = types.StringNull()
	}
	if val, ok := getResponseData["maxmatchlength"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Maxmatchlength = types.Int64Value(intVal)
		}
	} else {
		data.Maxmatchlength = types.Int64Null()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["resourceid"]; ok && val != nil {
		data.Resourceid = types.StringValue(val.(string))
	} else {
		data.Resourceid = types.StringNull()
	}
	if val, ok := getResponseData["ruletype"]; ok && val != nil {
		data.Ruletype = types.StringValue(val.(string))
	} else {
		data.Ruletype = types.StringNull()
	}
	if val, ok := getResponseData["safeobject"]; ok && val != nil {
		data.Safeobject = types.StringValue(val.(string))
	} else {
		data.Safeobject = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else {
		data.State = types.StringNull()
	}

	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("safeobject:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Safeobject.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
