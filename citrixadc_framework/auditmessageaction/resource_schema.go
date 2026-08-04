package auditmessageaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/audit"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// AuditmessageactionResourceModel describes the resource data model.
type AuditmessageactionResourceModel struct {
	Id                types.String `tfsdk:"id"`
	Bypasssafetycheck types.String `tfsdk:"bypasssafetycheck"`
	Loglevel          types.String `tfsdk:"loglevel"`
	Logtonewnslog     types.String `tfsdk:"logtonewnslog"`
	Name              types.String `tfsdk:"name"`
	Stringbuilderexpr types.String `tfsdk:"stringbuilderexpr"`
}

func (r *AuditmessageactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the auditmessageaction resource.",
			},
			"bypasssafetycheck": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bypass the safety check and allow unsafe expressions.",
			},
			"loglevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Audit log level, which specifies the severity level of the log message being generated..\nThe following loglevels are valid:\n* EMERGENCY - Events that indicate an immediate crisis on the server.\n* ALERT - Events that might require action.\n* CRITICAL - Events that indicate an imminent server crisis.\n* ERROR - Events that indicate some type of error.\n* WARNING - Events that require action in the near future.\n* NOTICE - Events that the administrator should know about.\n* INFORMATIONAL - All but low-level events.\n* DEBUG - All events, in extreme detail.",
			},
			"logtonewnslog": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Send the message to the new nslog.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the audit message action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the message action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my message action\" or 'my message action').",
			},
			"stringbuilderexpr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default-syntax expression that defines the format and content of the log message.",
			},
		},
	}
}

func auditmessageactionGetThePayloadFromthePlan(ctx context.Context, data *AuditmessageactionResourceModel) audit.Auditmessageaction {
	tflog.Debug(ctx, "In auditmessageactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	auditmessageaction := audit.Auditmessageaction{}
	if !data.Bypasssafetycheck.IsNull() && !data.Bypasssafetycheck.IsUnknown() {
		auditmessageaction.Bypasssafetycheck = data.Bypasssafetycheck.ValueString()
	}
	if !data.Loglevel.IsNull() && !data.Loglevel.IsUnknown() {
		auditmessageaction.Loglevel = data.Loglevel.ValueString()
	}
	if !data.Logtonewnslog.IsNull() && !data.Logtonewnslog.IsUnknown() {
		auditmessageaction.Logtonewnslog = data.Logtonewnslog.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		auditmessageaction.Name = data.Name.ValueString()
	}
	if !data.Stringbuilderexpr.IsNull() && !data.Stringbuilderexpr.IsUnknown() {
		auditmessageaction.Stringbuilderexpr = data.Stringbuilderexpr.ValueString()
	}

	return auditmessageaction
}

func auditmessageactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *AuditmessageactionResourceModel) audit.Auditmessageaction {
	tflog.Debug(ctx, "In auditmessageactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields
	auditmessageaction := audit.Auditmessageaction{}
	if !data.Bypasssafetycheck.IsNull() && !data.Bypasssafetycheck.IsUnknown() {
		auditmessageaction.Bypasssafetycheck = data.Bypasssafetycheck.ValueString()
	}
	if !data.Loglevel.IsNull() && !data.Loglevel.IsUnknown() {
		auditmessageaction.Loglevel = data.Loglevel.ValueString()
	}
	if !data.Logtonewnslog.IsNull() && !data.Logtonewnslog.IsUnknown() {
		auditmessageaction.Logtonewnslog = data.Logtonewnslog.ValueString()
	}
	// name is the primary key and is required in the update payload
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		auditmessageaction.Name = data.Name.ValueString()
	}
	if !data.Stringbuilderexpr.IsNull() && !data.Stringbuilderexpr.IsUnknown() {
		auditmessageaction.Stringbuilderexpr = data.Stringbuilderexpr.ValueString()
	}

	return auditmessageaction
}

func auditmessageactionSetAttrFromGet(ctx context.Context, data *AuditmessageactionResourceModel, getResponseData map[string]interface{}) *AuditmessageactionResourceModel {
	tflog.Debug(ctx, "In auditmessageactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["bypasssafetycheck"]; ok && val != nil {
		data.Bypasssafetycheck = types.StringValue(val.(string))
	} else {
		data.Bypasssafetycheck = types.StringNull()
	}
	// Note: NITRO GET returns the resolved log level in the read-only "loglevel1"
	// field. SDK v2 mapped state "loglevel" from "loglevel1"; preserved here for
	// backward compatibility.
	if val, ok := getResponseData["loglevel1"]; ok && val != nil {
		data.Loglevel = types.StringValue(val.(string))
	} else {
		data.Loglevel = types.StringNull()
	}
	if val, ok := getResponseData["logtonewnslog"]; ok && val != nil {
		data.Logtonewnslog = types.StringValue(val.(string))
	} else {
		data.Logtonewnslog = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["stringbuilderexpr"]; ok && val != nil {
		data.Stringbuilderexpr = types.StringValue(val.(string))
	} else {
		data.Stringbuilderexpr = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
