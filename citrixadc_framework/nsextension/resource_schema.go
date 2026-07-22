package nsextension

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsextensionResourceModel describes the resource data model.
type NsextensionResourceModel struct {
	Id             types.String `tfsdk:"id"`
	Comment        types.String `tfsdk:"comment"`
	Name           types.String `tfsdk:"name"`
	Overwrite      types.Bool   `tfsdk:"overwrite"`
	Src            types.String `tfsdk:"src"`
	Trace          types.String `tfsdk:"trace"`
	Tracefunctions types.String `tfsdk:"tracefunctions"`
	Tracevariables types.String `tfsdk:"tracevariables"`
}

func (r *NsextensionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsextension resource.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Any comments to preserve information about the extension object.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name to assign to the extension object on the Citrix ADC.",
			},
			"overwrite": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Overwrites the existing file",
			},
			"src": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Local path to and name of, or URL (protocol, host, path, and file name) for, the file in which to store the imported extension.\nNOTE: The import fails if the object to be imported is on an HTTPS server that requires client certificate authentication for access, and the issuer certificate of the HTTPS server is not present in the specific path on NetScaler to authenticate the HTTPS server.",
			},
			"trace": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables tracing to the NS log file of extension execution:\n   off   - turns off tracing (equivalent to unset ns extension <extension-name> -trace)\n   calls - traces extension function calls with arguments and function returns with the first return value\n   lines - traces the above plus line numbers for executed extension lines\n   all   - traces the above plus local variables changed by executed extension lines\nNote that the DEBUG log level must be enabled to see extension tracing.\nThis can be done by set audit syslogParams -loglevel ALL or -loglevel DEBUG.",
			},
			"tracefunctions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of extension functions to trace. By default, all extension functions are traced.",
			},
			"tracevariables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of variables (in traced extension functions) to trace. By default, all variables are traced.",
			},
		},
	}
}

// Import (create) payload. The ?action=Import endpoint accepts only
// src/name/comment/overwrite (per the NITRO doc). detail is a GET-only filter
// and the trace* fields belong to the update (PUT) endpoint, so they are excluded.
func nsextensionGetThePayloadFromthePlan(ctx context.Context, data *NsextensionResourceModel) ns.Nsextension {
	tflog.Debug(ctx, "In nsextensionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	nsextension := ns.Nsextension{}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nsextension.Comment = data.Comment.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsextension.Name = data.Name.ValueString()
	}
	if !data.Overwrite.IsNull() && !data.Overwrite.IsUnknown() {
		nsextension.Overwrite = data.Overwrite.ValueBool()
	}
	if !data.Src.IsNull() && !data.Src.IsUnknown() {
		nsextension.Src = data.Src.ValueString()
	}

	return nsextension
}

// Update (PUT) payload. The update endpoint accepts a different field set than
// Import: name/trace/tracefunctions/tracevariables/comment only.
func nsextensionGetTheUpdatePayloadFromthePlan(ctx context.Context, data *NsextensionResourceModel) ns.Nsextension {
	tflog.Debug(ctx, "In nsextensionGetTheUpdatePayloadFromthePlan Function")

	nsextension := ns.Nsextension{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		nsextension.Name = data.Name.ValueString()
	}
	if !data.Comment.IsNull() && !data.Comment.IsUnknown() {
		nsextension.Comment = data.Comment.ValueString()
	}
	if !data.Trace.IsNull() && !data.Trace.IsUnknown() {
		nsextension.Trace = data.Trace.ValueString()
	}
	if !data.Tracefunctions.IsNull() && !data.Tracefunctions.IsUnknown() {
		nsextension.Tracefunctions = data.Tracefunctions.ValueString()
	}
	if !data.Tracevariables.IsNull() && !data.Tracevariables.IsUnknown() {
		nsextension.Tracevariables = data.Tracevariables.ValueString()
	}

	return nsextension
}

// Resource-side state setter. src and overwrite are write-only import inputs
// (src is normalized by the server, overwrite is not meaningfully echoed), so
// they are preserved from the existing plan/state rather than copied from GET.
// The ID is set once in Create; it is not recomputed here.
func nsextensionSetAttrFromGet(ctx context.Context, data *NsextensionResourceModel, getResponseData map[string]interface{}) *NsextensionResourceModel {
	tflog.Debug(ctx, "In nsextensionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// src / overwrite: preserve existing plan/state value (write-only import inputs).
	if val, ok := getResponseData["trace"]; ok && val != nil {
		data.Trace = types.StringValue(val.(string))
	} else {
		data.Trace = types.StringNull()
	}
	if val, ok := getResponseData["tracefunctions"]; ok && val != nil {
		data.Tracefunctions = types.StringValue(val.(string))
	} else {
		data.Tracefunctions = types.StringNull()
	}
	if val, ok := getResponseData["tracevariables"]; ok && val != nil {
		data.Tracevariables = types.StringValue(val.(string))
	} else {
		data.Tracevariables = types.StringNull()
	}

	return data
}

// Datasource-side state setter. Faithfully copies every field from the GET
// response (the datasource has no prior plan/state to preserve) and sets the ID.
func nsextensionSetAttrFromGetForDatasource(ctx context.Context, data *NsextensionResourceModel, getResponseData map[string]interface{}) *NsextensionResourceModel {
	tflog.Debug(ctx, "In nsextensionSetAttrFromGetForDatasource Function")

	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["overwrite"]; ok && val != nil {
		if b, ok := val.(bool); ok {
			data.Overwrite = types.BoolValue(b)
		} else {
			data.Overwrite = types.BoolNull()
		}
	} else {
		data.Overwrite = types.BoolNull()
	}
	if val, ok := getResponseData["src"]; ok && val != nil {
		data.Src = types.StringValue(val.(string))
	} else {
		data.Src = types.StringNull()
	}
	if val, ok := getResponseData["trace"]; ok && val != nil {
		data.Trace = types.StringValue(val.(string))
	} else {
		data.Trace = types.StringNull()
	}
	if val, ok := getResponseData["tracefunctions"]; ok && val != nil {
		data.Tracefunctions = types.StringValue(val.(string))
	} else {
		data.Tracefunctions = types.StringNull()
	}
	if val, ok := getResponseData["tracevariables"]; ok && val != nil {
		data.Tracevariables = types.StringValue(val.(string))
	} else {
		data.Tracevariables = types.StringNull()
	}

	// Datasource has no Create; set the ID from the name key.
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
