package responderaction

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward and
// removal is a silent no-op. It does nothing when the config still carries a
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

// unsetOnRemoveInt64Modifier is the Int64 counterpart of unsetOnRemoveStringModifier.
type unsetOnRemoveInt64Modifier struct{}

func (m unsetOnRemoveInt64Modifier) Description(_ context.Context) string {
	return "Marks the value unknown when removed from config while a prior non-zero value exists, so it is unset on the appliance."
}

func (m unsetOnRemoveInt64Modifier) MarkdownDescription(ctx context.Context) string {
	return m.Description(ctx)
}

func (m unsetOnRemoveInt64Modifier) PlanModifyInt64(_ context.Context, req planmodifier.Int64Request, resp *planmodifier.Int64Response) {
	if req.StateValue.IsNull() {
		return
	}
	if req.ConfigValue.IsNull() && req.StateValue.ValueInt64() != 0 {
		resp.PlanValue = types.Int64Unknown()
	}
}

// ResponderactionResourceModel describes the resource data model.
type ResponderactionResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Bypasssafetycheck  types.String `tfsdk:"bypasssafetycheck"`
	Comment            types.String `tfsdk:"comment"`
	Headers            types.List   `tfsdk:"headers"`
	Htmlpage           types.String `tfsdk:"htmlpage"`
	Name               types.String `tfsdk:"name"`
	Newname            types.String `tfsdk:"newname"`
	Reasonphrase       types.String `tfsdk:"reasonphrase"`
	Responsestatuscode types.Int64  `tfsdk:"responsestatuscode"`
	Target             types.String `tfsdk:"target"`
	Type               types.String `tfsdk:"type"`
}

func (r *ResponderactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the responderaction resource.",
			},
			"bypasssafetycheck": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead.",
			},
			"comment": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Comment. Any type of information about this responder action.",
			},
			"headers": schema.ListAttribute{
				// SDK v2: TypeList of strings, Optional+Computed, not ForceNew.
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more headers to insert into the HTTP response. Each header is specified as \"name(expr)\", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for a responder action.",
			},
			"htmlpage": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.",
			},
			"name": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew. Keep Optional+Computed (a name is
				// auto-generated in Create when omitted, mirroring the SDK v2 resource).
				// UseStateForUnknown keeps the (possibly generated) name stable across
				// refreshes; RequiresReplaceIfConfigured reproduces ForceNew only when the
				// user actually configured the name and changes it.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Name for the responder action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or 'my responder action').",
			},
			"newname": schema.StringAttribute{
				// newname is the rename trigger (NITRO ?action=rename). It is NOT present
				// in the SDK v2 schema, so it must not force replacement (auto-gen added a
				// spurious RequiresReplace that is removed here). It drives an in-place
				// rename via Update. Optional only: a pure user input, never echoed by GET.
				Optional:    true,
				Description: "New name for the responder action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or my responder action').",
			},
			"reasonphrase": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: \"Invalid URL: \" + HTTP.REQ.URL",
			},
			"responsestatuscode": schema.Int64Attribute{
				// SDK v2: TypeInt, Optional+Computed, not ForceNew.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{},
				},
				Description: "HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200",
			},
			"target": schema.StringAttribute{
				// SDK v2: Optional+Computed, not ForceNew.
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.  In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements.\n\nExamples:\n1) Respondwith expression that sends an HTTP 1.1 200 OK response:\n\"HTTP/1.1 200 OK\\r\\n\\r\\n\"\n\n2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect.\n\"http://backupsite2.com\" + HTTP.REQ.URL\n\n3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response:\n\"HTTP/1.1 404 Not Found\\r\\n\\r\\n\"+ \"HTTP.REQ.URL.HTTP_URL_SAFE\" + \"does not exist on the web server.\"\n\nThe following requirement applies only to the Citrix ADC CLI:\nEnclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.)",
			},
			"type": schema.StringAttribute{
				// SDK v2: Optional+Computed+ForceNew. Keep Optional+Computed (NITRO requires
				// it for add, but the SDK v2 contract was Optional). UseStateForUnknown +
				// RequiresReplaceIfConfigured reproduce the ForceNew semantics.
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Description: "Type of responder action. Available settings function as follows:\n* respondwith <target> - Respond to the request with the expression specified as the target.\n* respondwithhtmlpage - Respond to the request with the uploaded HTML page object specified as the target.\n* redirect - Redirect the request to the URL specified as the target.\n* sqlresponse_ok - Send an SQL OK response.\n* sqlresponse_error - Send an SQL ERROR response.",
			},
		},
	}
}

// responderactionGetThePayloadFromthePlan builds the full add (create) payload.
func responderactionGetThePayloadFromthePlan(ctx context.Context, data *ResponderactionResourceModel) responder.Responderaction {
	tflog.Debug(ctx, "In responderactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	responderaction := responder.Responderaction{}
	if !data.Bypasssafetycheck.IsNull() && !data.Bypasssafetycheck.IsUnknown() {
		responderaction.Bypasssafetycheck = data.Bypasssafetycheck.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		responderaction.Comment = data.Comment.ValueString()
	}
	if !data.Headers.IsNull() && !data.Headers.IsUnknown() {
		var headersList []string
		data.Headers.ElementsAs(ctx, &headersList, false)
		responderaction.Headers = headersList
	}
	if !data.Htmlpage.IsNull() && !data.Htmlpage.IsUnknown() {
		responderaction.Htmlpage = data.Htmlpage.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		responderaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only argument (NITRO ?action=rename). It is NOT part of the
	// add payload, so it is deliberately excluded from the create POST body.
	if !data.Reasonphrase.IsNull() && !data.Reasonphrase.IsUnknown() {
		responderaction.Reasonphrase = data.Reasonphrase.ValueString()
	}
	if !data.Responsestatuscode.IsNull() && !data.Responsestatuscode.IsUnknown() {
		responderaction.Responsestatuscode = utils.IntPtr(int(data.Responsestatuscode.ValueInt64()))
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		responderaction.Target = data.Target.ValueString()
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		responderaction.Type = data.Type.ValueString()
	}

	return responderaction
}

// responderactionGetTheUpdatablePayloadFromThePlan builds the PUT (update) payload,
// restricted to the NITRO-updatable fields. name and type are ForceNew (never reach
// Update); newname is rename-only (handled separately).
func responderactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *ResponderactionResourceModel) responder.Responderaction {
	tflog.Debug(ctx, "In responderactionGetTheUpdatablePayloadFromThePlan Function")

	responderaction := responder.Responderaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		responderaction.Name = data.Name.ValueString()
	}
	if !data.Bypasssafetycheck.IsNull() && !data.Bypasssafetycheck.IsUnknown() {
		responderaction.Bypasssafetycheck = data.Bypasssafetycheck.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		responderaction.Comment = data.Comment.ValueString()
	}
	if !data.Headers.IsNull() && !data.Headers.IsUnknown() {
		var headersList []string
		data.Headers.ElementsAs(ctx, &headersList, false)
		responderaction.Headers = headersList
	}
	if !data.Htmlpage.IsNull() && !data.Htmlpage.IsUnknown() {
		responderaction.Htmlpage = data.Htmlpage.ValueString()
	}
	if !data.Reasonphrase.IsNull() && !data.Reasonphrase.IsUnknown() {
		responderaction.Reasonphrase = data.Reasonphrase.ValueString()
	}
	if !data.Responsestatuscode.IsNull() && !data.Responsestatuscode.IsUnknown() {
		responderaction.Responsestatuscode = utils.IntPtr(int(data.Responsestatuscode.ValueInt64()))
	}
	if !data.Target.IsNull() && !data.Target.IsUnknown() {
		responderaction.Target = data.Target.ValueString()
	}

	return responderaction
}

// responderactionSetAttrFromGet populates the RESOURCE model from a GET response.
// It preserves configured/known values that NITRO omits from GET (omit-on-default
// trap), only nulling a field when the model value is still unknown.
func responderactionSetAttrFromGet(ctx context.Context, data *ResponderactionResourceModel, getResponseData map[string]interface{}) *ResponderactionResourceModel {
	tflog.Debug(ctx, "In responderactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bypasssafetycheck"]; ok && val != nil {
		data.Bypasssafetycheck = types.StringValue(val.(string))
	} else if data.Bypasssafetycheck.IsUnknown() {
		data.Bypasssafetycheck = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else if data.Comment.IsUnknown() {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["headers"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Headers = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Headers = listValue
		default:
			if data.Headers.IsUnknown() {
				data.Headers = types.ListNull(types.StringType)
			}
		}
	} else if data.Headers.IsUnknown() {
		data.Headers = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["htmlpage"]; ok && val != nil {
		data.Htmlpage = types.StringValue(val.(string))
	} else if data.Htmlpage.IsUnknown() {
		data.Htmlpage = types.StringNull()
	}
	// name is the user-facing key. After a rename (via newname) the live object name
	// (tracked by data.Id) diverges from the configured name, and GET returns the live
	// name. Overwriting name from GET would clobber the user's configured value and
	// trigger a spurious replace. Only adopt the GET value when the model has none
	// (e.g. on import, where state carries only the ID); otherwise preserve.
	if data.Name.IsNull() || data.Name.IsUnknown() || data.Name.ValueString() == "" {
		if val, ok := getResponseData["name"]; ok && val != nil {
			data.Name = types.StringValue(val.(string))
		}
	}
	// newname is rename-only and never echoed by GET; preserve plan/state value.
	if val, ok := getResponseData["reasonphrase"]; ok && val != nil {
		data.Reasonphrase = types.StringValue(val.(string))
	} else if data.Reasonphrase.IsUnknown() {
		data.Reasonphrase = types.StringNull()
	}
	if val, ok := getResponseData["responsestatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Responsestatuscode = types.Int64Value(intVal)
		} else if data.Responsestatuscode.IsUnknown() {
			data.Responsestatuscode = types.Int64Null()
		}
	} else if data.Responsestatuscode.IsUnknown() {
		data.Responsestatuscode = types.Int64Null()
	}
	if val, ok := getResponseData["target"]; ok && val != nil {
		data.Target = types.StringValue(val.(string))
	} else if data.Target.IsUnknown() {
		data.Target = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else if data.Type.IsUnknown() {
		data.Type = types.StringNull()
	}

	// Set ID for the resource (single unique attribute: name). Do not overwrite an
	// ID already tracking the live (possibly renamed) name.
	if data.Id.IsNull() || data.Id.IsUnknown() || data.Id.ValueString() == "" {
		data.Id = types.StringValue(data.Name.ValueString())
	}

	return data
}

// responderactionSetAttrFromGetForDatasource faithfully copies every field from the
// GET response. The datasource has no prior plan/state to preserve, so it populates
// the model directly from the API response and sets the ID itself.
func responderactionSetAttrFromGetForDatasource(ctx context.Context, data *ResponderactionResourceModel, getResponseData map[string]interface{}) *ResponderactionResourceModel {
	tflog.Debug(ctx, "In responderactionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["bypasssafetycheck"]; ok && val != nil {
		data.Bypasssafetycheck = types.StringValue(val.(string))
	} else {
		data.Bypasssafetycheck = types.StringNull()
	}
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["headers"]; ok && val != nil {
		switch v := val.(type) {
		case []interface{}:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, utils.ToStringList(v))
			data.Headers = listValue
		case string:
			listValue, _ := types.ListValueFrom(ctx, types.StringType, []string{v})
			data.Headers = listValue
		default:
			data.Headers = types.ListNull(types.StringType)
		}
	} else {
		data.Headers = types.ListNull(types.StringType)
	}
	if val, ok := getResponseData["htmlpage"]; ok && val != nil {
		data.Htmlpage = types.StringValue(val.(string))
	} else {
		data.Htmlpage = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// newname is rename-only and never echoed by GET.
	data.Newname = types.StringNull()
	if val, ok := getResponseData["reasonphrase"]; ok && val != nil {
		data.Reasonphrase = types.StringValue(val.(string))
	} else {
		data.Reasonphrase = types.StringNull()
	}
	if val, ok := getResponseData["responsestatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Responsestatuscode = types.Int64Value(intVal)
		} else {
			data.Responsestatuscode = types.Int64Null()
		}
	} else {
		data.Responsestatuscode = types.Int64Null()
	}
	if val, ok := getResponseData["target"]; ok && val != nil {
		data.Target = types.StringValue(val.(string))
	} else {
		data.Target = types.StringNull()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}

	// Single unique attribute - use plain value as ID.
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
