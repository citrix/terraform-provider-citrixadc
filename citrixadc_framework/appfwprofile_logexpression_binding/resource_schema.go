package appfwprofile_logexpression_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/resource/config/appfw"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// AppfwprofileLogexpressionBindingResourceModel describes the resource data model.
type AppfwprofileLogexpressionBindingResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Alertonly       types.String `tfsdk:"alertonly"`
	AsLogexpression types.String `tfsdk:"as_logexpression"`
	Comment         types.String `tfsdk:"comment"`
	Isautodeployed  types.String `tfsdk:"isautodeployed"`
	Logexpression   types.String `tfsdk:"logexpression"`
	Name            types.String `tfsdk:"name"`
	Resourceid      types.String `tfsdk:"resourceid"`
	State           types.String `tfsdk:"state"`
	Ruletype        types.String `tfsdk:"ruletype"`
}

func (r *AppfwprofileLogexpressionBindingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the appfwprofile_logexpression_binding resource.",
			},
			"alertonly": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Send SNMP alert?",
			},
			"as_logexpression": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "LogExpression to log when violation happened on appfw profile",
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
			"logexpression": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of LogExpression object.",
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
			"state": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Enabled.",
			},
			"ruletype": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Specifies rule type of binding.",
			},
		},
	}
}

func appfwprofile_logexpression_bindingGetThePayloadFromthePlan(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel) appfw.Appfwprofilelogexpressionbinding {
	tflog.Debug(ctx, "In appfwprofile_logexpression_bindingGetThePayloadFromthePlan Function")

	// Create API request body from the model
	appfwprofile_logexpression_binding := appfw.Appfwprofilelogexpressionbinding{}
	if !data.Alertonly.IsNull() && !data.Alertonly.IsUnknown() {
		appfwprofile_logexpression_binding.Alertonly = data.Alertonly.ValueString()
	}
	if !data.AsLogexpression.IsNull() && !data.AsLogexpression.IsUnknown() {
		appfwprofile_logexpression_binding.Aslogexpression = data.AsLogexpression.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		appfwprofile_logexpression_binding.Comment = data.Comment.ValueString()
	}
	if !data.Isautodeployed.IsNull() && !data.Isautodeployed.IsUnknown() {
		appfwprofile_logexpression_binding.Isautodeployed = data.Isautodeployed.ValueString()
	}
	if !data.Logexpression.IsNull() && !data.Logexpression.IsUnknown() {
		appfwprofile_logexpression_binding.Logexpression = data.Logexpression.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		appfwprofile_logexpression_binding.Name = data.Name.ValueString()
	}
	if !data.Resourceid.IsNull() && !data.Resourceid.IsUnknown() {
		appfwprofile_logexpression_binding.Resourceid = data.Resourceid.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		appfwprofile_logexpression_binding.State = data.State.ValueString()
	}
	if !data.Ruletype.IsNull() && !data.Ruletype.IsUnknown() {
		appfwprofile_logexpression_binding.Ruletype = data.Ruletype.ValueString()
	}

	return appfwprofile_logexpression_binding
}

// appfwprofile_logexpression_bindingSetAttrFromGet is the RESOURCE-side setter.
// All attributes are RequiresReplace (no update endpoint) and the NITRO server may
// echo server-defaulted/normalized values for fields like alertonly and
// isautodeployed (the SDK v2 resource explicitly did NOT write those back). To avoid
// "inconsistent result after apply" we adopt the GET value only when the model field
// is currently null/unknown (e.g. import); otherwise we preserve the configured
// plan/state value. The ID is set once in Create and is preserved here.
func appfwprofile_logexpression_bindingSetAttrFromGet(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileLogexpressionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_logexpression_bindingSetAttrFromGet Function")

	adopt := func(cur types.String, key string) types.String {
		if !cur.IsNull() && !cur.IsUnknown() {
			return cur
		}
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = adopt(data.Alertonly, "alertonly")
	data.AsLogexpression = adopt(data.AsLogexpression, "as_logexpression")
	data.Comment = adopt(data.Comment, "comment")
	data.Isautodeployed = adopt(data.Isautodeployed, "isautodeployed")
	data.Logexpression = adopt(data.Logexpression, "logexpression")
	data.Name = adopt(data.Name, "name")
	data.Resourceid = adopt(data.Resourceid, "resourceid")
	data.State = adopt(data.State, "state")
	data.Ruletype = adopt(data.Ruletype, "ruletype")

	return data
}

// appfwprofile_logexpression_bindingSetAttrFromGetForDatasource is the
// DATASOURCE-side setter: it faithfully copies every field from the GET response
// (the datasource has no prior plan/state to preserve) and sets the composite ID.
func appfwprofile_logexpression_bindingSetAttrFromGetForDatasource(ctx context.Context, data *AppfwprofileLogexpressionBindingResourceModel, getResponseData map[string]interface{}) *AppfwprofileLogexpressionBindingResourceModel {
	tflog.Debug(ctx, "In appfwprofile_logexpression_bindingSetAttrFromGetForDatasource Function")

	copyField := func(key string) types.String {
		if val, ok := getResponseData[key]; ok && val != nil {
			return types.StringValue(val.(string))
		}
		return types.StringNull()
	}

	data.Alertonly = copyField("alertonly")
	data.AsLogexpression = copyField("as_logexpression")
	data.Comment = copyField("comment")
	data.Isautodeployed = copyField("isautodeployed")
	data.Logexpression = copyField("logexpression")
	data.Name = copyField("name")
	data.Resourceid = copyField("resourceid")
	data.State = copyField("state")
	data.Ruletype = copyField("ruletype")

	// Set ID for the datasource
	// Case 3: Multiple unique attributes - comma-separated key:UrlEncode(value) pairs
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("logexpression:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Logexpression.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("name:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Name.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	return data
}
