package transformaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/transform"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward
// and removal is a silent no-op. It does nothing when the config still carries a
// value, on create (no prior state), or when the prior value is already empty.
type unsetOnRemoveStringModifier struct{}

func (m unsetOnRemoveStringModifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-empty value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveStringModifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveStringModifier) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueString() != "" {
		resp.PlanValue = types.StringUnknown()
	}
}

// TransformactionResourceModel describes the resource data model.
type TransformactionResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Cookiedomainfrom types.String `tfsdk:"cookiedomainfrom"`
	Cookiedomaininto types.String `tfsdk:"cookiedomaininto"`
	Name             types.String `tfsdk:"name"`
	Priority         types.Int64  `tfsdk:"priority"`
	Profilename      types.String `tfsdk:"profilename"`
	Requrlfrom       types.String `tfsdk:"requrlfrom"`
	Requrlinto       types.String `tfsdk:"requrlinto"`
	Resurlfrom       types.String `tfsdk:"resurlfrom"`
	Resurlinto       types.String `tfsdk:"resurlinto"`
	State            types.String `tfsdk:"state"`
}

func (r *TransformactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the transformaction resource.",
			},
			"comment": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Any comments to preserve information about this URL Transformation action.",
			},
			"cookiedomainfrom": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Pattern that matches the domain to be transformed in Set-Cookie headers.",
			},
			"cookiedomaininto": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "PCRE-format regular expression that describes the transformation to be performed on cookie domains that match the cookieDomainFrom pattern. \nNOTE: The cookie domain to be transformed is extracted from the request.",
			},
			// SDK v2: Required + ForceNew
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the URL transformation action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the URL Transformation action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, my transform action or my transform action).",
			},
			// SDK v2: Optional + Computed (TypeInt). Metadata marks it required, but the
			// backward-compatible SDK v2 contract is Optional + Computed.
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Positive integer specifying the priority of the action within the profile. A lower number specifies a higher priority. Must be unique within the list of actions bound to the profile. Policies are evaluated in the order of their priority numbers, and the first policy that matches is applied.",
			},
			// SDK v2: Optional + Computed + ForceNew -> UseStateForUnknown() then
			// RequiresReplaceIfConfigured().
			"profilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name of the URL Transformation profile with which to associate this action.",
			},
			"requrlfrom": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "PCRE-format regular expression that describes the request URL pattern to be transformed.",
			},
			"requrlinto": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the reqUrlFrom pattern.",
			},
			"resurlfrom": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "PCRE-format regular expression that describes the response URL pattern to be transformed.",
			},
			"resurlinto": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "PCRE-format regular expression that describes the transformation to be performed on URLs that match the resUrlFrom pattern.",
			},
			// SDK v2: Optional + Computed. Default ENABLED (NITRO spec default) so
			// removing it from config produces a plan diff and the unset fires;
			// unset reverts the appliance to ENABLED, matching this default.
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("ENABLED"),
				Description: "Enable or disable this action.",
			},
		},
	}
}

// transformactionGetTheAddPayloadFromthePlan builds the payload for the NITRO add
// operation. Mirroring SDK v2, add only carries name/profilename/state/priority;
// the URL/cookie transformation patterns are applied via a follow-up update.
func transformactionGetTheAddPayloadFromthePlan(ctx context.Context, data *TransformactionResourceModel) transform.Transformaction {
	tflog.Debug(ctx, "In transformactionGetTheAddPayloadFromthePlan Function")

	transformaction := transform.Transformaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		transformaction.Name = data.Name.ValueString()
	}
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		transformaction.Profilename = data.Profilename.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		transformaction.State = data.State.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		transformaction.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}

	return transformaction
}

// transformactionGetTheUpdatablePayloadFromThePlan builds the payload for the NITRO
// update operation. Mirroring SDK v2, profilename is NOT included (it is set only at
// add time and is ForceNew).
func transformactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *TransformactionResourceModel) transform.Transformaction {
	tflog.Debug(ctx, "In transformactionGetTheUpdatablePayloadFromThePlan Function")

	transformaction := transform.Transformaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		transformaction.Name = data.Name.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		transformaction.Comment = data.Comment.ValueString()
	}
	if !data.Cookiedomainfrom.IsNull() && !data.Cookiedomainfrom.IsUnknown() {
		transformaction.Cookiedomainfrom = data.Cookiedomainfrom.ValueString()
	}
	if !data.Cookiedomaininto.IsNull() && !data.Cookiedomaininto.IsUnknown() {
		transformaction.Cookiedomaininto = data.Cookiedomaininto.ValueString()
	}
	if !data.Priority.IsNull() && !data.Priority.IsUnknown() {
		transformaction.Priority = utils.IntPtr(int(data.Priority.ValueInt64()))
	}
	if !data.Requrlfrom.IsNull() && !data.Requrlfrom.IsUnknown() {
		transformaction.Requrlfrom = data.Requrlfrom.ValueString()
	}
	if !data.Requrlinto.IsNull() && !data.Requrlinto.IsUnknown() {
		transformaction.Requrlinto = data.Requrlinto.ValueString()
	}
	if !data.Resurlfrom.IsNull() && !data.Resurlfrom.IsUnknown() {
		transformaction.Resurlfrom = data.Resurlfrom.ValueString()
	}
	if !data.Resurlinto.IsNull() && !data.Resurlinto.IsUnknown() {
		transformaction.Resurlinto = data.Resurlinto.ValueString()
	}
	if !data.State.IsNull() && !data.State.IsUnknown() {
		transformaction.State = data.State.ValueString()
	}

	return transformaction
}

func transformactionSetAttrFromGet(ctx context.Context, data *TransformactionResourceModel, getResponseData map[string]interface{}) *TransformactionResourceModel {
	tflog.Debug(ctx, "In transformactionSetAttrFromGet Function")

	// Convert API response to model.
	// The else-branch only nulls a value that is still Unknown (e.g. an omitted
	// Optional+Computed attribute right after create). A KNOWN configured/state value
	// that NITRO omits from GET (omit-on-default) is preserved, never clobbered.
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomainfrom"]; ok && val != nil {
		data.Cookiedomainfrom = types.StringValue(val.(string))
	} else if data.Cookiedomainfrom.IsUnknown() {
		data.Cookiedomainfrom = types.StringNull()
	}
	if val, ok := getResponseData["cookiedomaininto"]; ok && val != nil {
		data.Cookiedomaininto = types.StringValue(val.(string))
	} else if data.Cookiedomaininto.IsUnknown() {
		data.Cookiedomaininto = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["priority"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Priority = types.Int64Value(intVal)
		}
	} else if data.Priority.IsUnknown() {
		data.Priority = types.Int64Null()
	}
	if val, ok := getResponseData["profilename"]; ok && val != nil {
		data.Profilename = types.StringValue(val.(string))
	} else if data.Profilename.IsUnknown() {
		data.Profilename = types.StringNull()
	}
	if val, ok := getResponseData["requrlfrom"]; ok && val != nil {
		data.Requrlfrom = types.StringValue(val.(string))
	} else if data.Requrlfrom.IsUnknown() {
		data.Requrlfrom = types.StringNull()
	}
	if val, ok := getResponseData["requrlinto"]; ok && val != nil {
		data.Requrlinto = types.StringValue(val.(string))
	} else if data.Requrlinto.IsUnknown() {
		data.Requrlinto = types.StringNull()
	}
	if val, ok := getResponseData["resurlfrom"]; ok && val != nil {
		data.Resurlfrom = types.StringValue(val.(string))
	} else if data.Resurlfrom.IsUnknown() {
		data.Resurlfrom = types.StringNull()
	}
	if val, ok := getResponseData["resurlinto"]; ok && val != nil {
		data.Resurlinto = types.StringValue(val.(string))
	} else if data.Resurlinto.IsUnknown() {
		data.Resurlinto = types.StringNull()
	}
	if val, ok := getResponseData["state"]; ok && val != nil {
		data.State = types.StringValue(val.(string))
	} else if data.State.IsUnknown() {
		data.State = types.StringNull()
	}

	// Set ID for the resource: single unique attribute (name).
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
