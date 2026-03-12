package responderaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/responder"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

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
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check, allowing potentially unsafe expressions. An unsafe expression in a response is one that contains references to request elements that might not be present in all requests. If a response refers to a missing request element, an empty string is used instead.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this responder action.",
			},
			"headers": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "One or more headers to insert into the HTTP response. Each header is specified as \"name(expr)\", where expr is an expression that is evaluated at runtime to provide the value for the named header. You can configure a maximum of eight headers for a responder action.",
			},
			"htmlpage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For respondwithhtmlpage policies, name of the HTML page object to use as the response. You must first import the page object.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the responder action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Can be changed after the responder policy is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or 'my responder action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the responder action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my responder action\" or my responder action').",
			},
			"reasonphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying the reason phrase of the HTTP response. The reason phrase may be a string literal with quotes or a PI expression. For example: \"Invalid URL: \" + HTTP.REQ.URL",
			},
			"responsestatuscode": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "HTTP response status code, for example 200, 302, 404, etc. The default value for the redirect action type is 302 and for respondwithhtmlpage is 200",
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expression specifying what to respond with. Typically a URL for redirect policies or a default-syntax expression.  In addition to Citrix ADC default-syntax expressions that refer to information in the request, a stringbuilder expression can contain text and HTML, and simple escape codes that define new lines and paragraphs. Enclose each stringbuilder expression element (either a Citrix ADC default-syntax expression or a string) in double quotation marks. Use the plus (+) character to join the elements.\n\nExamples:\n1) Respondwith expression that sends an HTTP 1.1 200 OK response:\n\"HTTP/1.1 200 OK\\r\\n\\r\\n\"\n\n2) Redirect expression that redirects user to the specified web host and appends the request URL to the redirect.\n\"http://backupsite2.com\" + HTTP.REQ.URL\n\n3) Respondwith expression that sends an HTTP 1.1 404 Not Found response with the request URL included in the response:\n\"HTTP/1.1 404 Not Found\\r\\n\\r\\n\"+ \"HTTP.REQ.URL.HTTP_URL_SAFE\" + \"does not exist on the web server.\"\n\nThe following requirement applies only to the Citrix ADC CLI:\nEnclose the entire expression in single quotation marks. (Citrix ADC expression elements should be included inside the single quotation marks for the entire expression, but do not need to be enclosed in double quotation marks.)",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of responder action. Available settings function as follows:\n* respondwith <target> - Respond to the request with the expression specified as the target.\n* respondwithhtmlpage - Respond to the request with the uploaded HTML page object specified as the target.\n* redirect - Redirect the request to the URL specified as the target.\n* sqlresponse_ok - Send an SQL OK response.\n* sqlresponse_error - Send an SQL ERROR response.",
			},
		},
	}
}

func responderactionGetThePayloadFromtheConfig(ctx context.Context, data *ResponderactionResourceModel) responder.Responderaction {
	tflog.Debug(ctx, "In responderactionGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	responderaction := responder.Responderaction{}
	if !data.Bypasssafetycheck.IsNull() {
		responderaction.Bypasssafetycheck = data.Bypasssafetycheck.ValueString()
	}
	if !data.Comment.IsNull() {
		responderaction.Comment = data.Comment.ValueString()
	}
	if !data.Htmlpage.IsNull() {
		responderaction.Htmlpage = data.Htmlpage.ValueString()
	}
	if !data.Name.IsNull() {
		responderaction.Name = data.Name.ValueString()
	}
	if !data.Newname.IsNull() {
		responderaction.Newname = data.Newname.ValueString()
	}
	if !data.Reasonphrase.IsNull() {
		responderaction.Reasonphrase = data.Reasonphrase.ValueString()
	}
	if !data.Responsestatuscode.IsNull() {
		responderaction.Responsestatuscode = utils.IntPtr(int(data.Responsestatuscode.ValueInt64()))
	}
	if !data.Target.IsNull() {
		responderaction.Target = data.Target.ValueString()
	}
	if !data.Type.IsNull() {
		responderaction.Type = data.Type.ValueString()
	}

	return responderaction
}

func responderactionSetAttrFromGet(ctx context.Context, data *ResponderactionResourceModel, getResponseData map[string]interface{}) *ResponderactionResourceModel {
	tflog.Debug(ctx, "In responderactionSetAttrFromGet Function")

	// Convert API response to model
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
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else {
		data.Newname = types.StringNull()
	}
	if val, ok := getResponseData["reasonphrase"]; ok && val != nil {
		data.Reasonphrase = types.StringValue(val.(string))
	} else {
		data.Reasonphrase = types.StringNull()
	}
	if val, ok := getResponseData["responsestatuscode"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Responsestatuscode = types.Int64Value(intVal)
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

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
