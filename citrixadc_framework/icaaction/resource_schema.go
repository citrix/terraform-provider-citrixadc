package icaaction

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ica"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// IcaactionResourceModel describes the resource data model.
type IcaactionResourceModel struct {
	Id                 types.String `tfsdk:"id"`
	Accessprofilename  types.String `tfsdk:"accessprofilename"`
	Latencyprofilename types.String `tfsdk:"latencyprofilename"`
	Name               types.String `tfsdk:"name"`
	Newname            types.String `tfsdk:"newname"`
}

// unsetOnRemoveStringModifier forces the planned value to unknown when the user
// removes a previously-set attribute from configuration while a non-empty value
// still exists in prior state. This makes Terraform detect a change (unknown !=
// prior) and call Update, which issues the NITRO ?action=unset. Without it an
// Optional+Computed attribute is "sticky": the prior value is carried forward and
// removal is a silent no-op.
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

func (r *IcaactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the icaaction resource.",
			},
			"accessprofilename": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{},
				},
				Description: "Name of the ica accessprofile to be associated with this action.",
			},
			"latencyprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ica latencyprofile to be associated with this action.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name for the ICA action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) pound (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the ICA action is added.\n\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks (for example, \"my ica action\" or 'my ica action').",
			},
			"newname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "New name for the ICA action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#),period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.\nThe following requirement applies only to the Citrix ADC CLI:\nIf the name includes one or more spaces, enclose the name in double or single quotation marks ( for example, \"my ica action\" or 'my ica action').",
			},
		},
	}
}

func icaactionGetThePayloadFromthePlan(ctx context.Context, data *IcaactionResourceModel) ica.Icaaction {
	tflog.Debug(ctx, "In icaactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	icaaction := ica.Icaaction{}
	if !data.Accessprofilename.IsNull() && !data.Accessprofilename.IsUnknown() {
		icaaction.Accessprofilename = data.Accessprofilename.ValueString()
	}
	if !data.Latencyprofilename.IsNull() && !data.Latencyprofilename.IsUnknown() {
		icaaction.Latencyprofilename = data.Latencyprofilename.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		icaaction.Name = data.Name.ValueString()
	}
	// newname is a rename-only parameter (NITRO ?action=rename) - excluded from the add payload

	return icaaction
}

func icaactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *IcaactionResourceModel) ica.Icaaction {
	tflog.Debug(ctx, "In icaactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	// name is always carried because the update is a PUT to /config/icaaction (unnamed).
	icaaction := ica.Icaaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		icaaction.Name = data.Name.ValueString()
	}
	if !data.Accessprofilename.IsNull() && !data.Accessprofilename.IsUnknown() {
		icaaction.Accessprofilename = data.Accessprofilename.ValueString()
	}
	if !data.Latencyprofilename.IsNull() && !data.Latencyprofilename.IsUnknown() {
		icaaction.Latencyprofilename = data.Latencyprofilename.ValueString()
	}
	// newname is a rename-only parameter - excluded from the update payload

	return icaaction
}

func icaactionSetAttrFromGet(ctx context.Context, data *IcaactionResourceModel, getResponseData map[string]interface{}) *IcaactionResourceModel {
	tflog.Debug(ctx, "In icaactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["accessprofilename"]; ok && val != nil {
		data.Accessprofilename = types.StringValue(val.(string))
	} else {
		data.Accessprofilename = types.StringNull()
	}
	if val, ok := getResponseData["latencyprofilename"]; ok && val != nil {
		data.Latencyprofilename = types.StringValue(val.(string))
	} else {
		data.Latencyprofilename = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	// newname is a rename-only parameter and is never returned by the NITRO GET.
	// Preserve a configured value; only resolve an unknown (Computed) value to null
	// so a configured newname does not trigger "inconsistent result after apply".
	if val, ok := getResponseData["newname"]; ok && val != nil {
		data.Newname = types.StringValue(val.(string))
	} else if data.Newname.IsUnknown() {
		data.Newname = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(data.Name.ValueString())

	return data
}
