package contentinspectionaction

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
// removal is a silent no-op.
//
// It intentionally does nothing on create (no prior state) so the value is not
// forced into the create payload -- important here because ifserverdown/serverport
// are rejected by NITRO on a NOINSPECTION action. It also does nothing once the
// prior state already equals the default, avoiding a perpetual post-unset plan
// diff (NITRO always echoes the default back on GET for these attributes).
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
	if req.ConfigValue.IsNull() && req.StateValue.ValueInt64() != 0 && req.StateValue.ValueInt64() != m.defaultValue {
		resp.PlanValue = types.Int64Unknown()
	}
}

// ContentinspectionactionResourceModel describes the resource data model.
type ContentinspectionactionResourceModel struct {
	Id              types.String `tfsdk:"id"`
	Icapprofilename types.String `tfsdk:"icapprofilename"`
	Ifserverdown    types.String `tfsdk:"ifserverdown"`
	Name            types.String `tfsdk:"name"`
	Serverip        types.String `tfsdk:"serverip"`
	Servername      types.String `tfsdk:"servername"`
	Serverport      types.Int64  `tfsdk:"serverport"`
	Type            types.String `tfsdk:"type"`
	Wasmprofilename types.String `tfsdk:"wasmprofilename"`
}

func (r *ContentinspectionactionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the contentinspectionaction resource.",
			},
			"icapprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the ICAP profile to be attached to the contentInspection action.",
			},
			"ifserverdown": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					unsetOnRemoveStringModifier{defaultValue: "RESET"},
				},
				Description: "Name of the action to perform if the Vserver representing the remote service is not UP. This is not supported for NOINSPECTION Type. The Supported actions are:\n* RESET - Reset the client connection by closing it. The client program, such as a browser, will handle this and may inform the user. The client may then resend the request if desired.\n* DROP - Drop the request without sending a response to the user.\n* CONTINUE - It bypasses the ContentIsnpection and Continues/resumes the Traffic-Flow to Client/Server.",
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the remote service action. Must begin with an ASCII alphabetic or underscore (_) character, and must contain only ASCII alphanumeric, underscore, hash (#), period (.), space, colon (:), at (@), equals (=), and hyphen (-) characters.",
			},
			"serverip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remoteService",
			},
			"servername": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the LB vserver or service",
			},
			"serverport": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					unsetOnRemoveInt64Modifier{defaultValue: 1344},
				},
				Description: "Port of remoteService",
			},
			"type": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Type of operation this action is going to perform. following actions are available to configure:\n* ICAP - forward the incoming request or response to an ICAP server for modification.\n* INLINEINSPECTION - forward the incoming or outgoing packets to IPS server for Intrusion Prevention.\n* MIRROR - Forwards cloned packets for Intrusion Detection.\n* NOINSPECTION - This does not forward incoming and outgoing packets to the Inspection device.\n* NSTRACE - capture current and further incoming packets on this transaction.",
			},
			"wasmprofilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the CI WASM profile to be attached to the contentInspection action.",
			},
		},
	}
}

func contentinspectionactionGetThePayloadFromthePlan(ctx context.Context, data *ContentinspectionactionResourceModel) contentinspection.Contentinspectionaction {
	tflog.Debug(ctx, "In contentinspectionactionGetThePayloadFromthePlan Function")

	// Create API request body from the model
	contentinspectionaction := contentinspection.Contentinspectionaction{}
	if !data.Icapprofilename.IsNull() && !data.Icapprofilename.IsUnknown() {
		contentinspectionaction.Icapprofilename = data.Icapprofilename.ValueString()
	}
	if !data.Ifserverdown.IsNull() && !data.Ifserverdown.IsUnknown() {
		contentinspectionaction.Ifserverdown = data.Ifserverdown.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		contentinspectionaction.Name = data.Name.ValueString()
	}
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() {
		contentinspectionaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		contentinspectionaction.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() && !data.Serverport.IsUnknown() {
		contentinspectionaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Type.IsNull() && !data.Type.IsUnknown() {
		contentinspectionaction.Type = data.Type.ValueString()
	}
	if !data.Wasmprofilename.IsNull() && !data.Wasmprofilename.IsUnknown() {
		contentinspectionaction.Wasmprofilename = data.Wasmprofilename.ValueString()
	}

	return contentinspectionaction
}

func contentinspectionactionGetTheUpdatablePayloadFromThePlan(ctx context.Context, data *ContentinspectionactionResourceModel) contentinspection.Contentinspectionaction {
	tflog.Debug(ctx, "In contentinspectionactionGetTheUpdatablePayloadFromThePlan Function")

	// Create API request body from the model, restricted to NITRO-updatable fields.
	// "name" is the key (always sent); "type" is not updatable (ForceNew/RequiresReplace).
	contentinspectionaction := contentinspection.Contentinspectionaction{}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		contentinspectionaction.Name = data.Name.ValueString()
	}
	if !data.Icapprofilename.IsNull() && !data.Icapprofilename.IsUnknown() {
		contentinspectionaction.Icapprofilename = data.Icapprofilename.ValueString()
	}
	if !data.Ifserverdown.IsNull() && !data.Ifserverdown.IsUnknown() {
		contentinspectionaction.Ifserverdown = data.Ifserverdown.ValueString()
	}
	if !data.Serverip.IsNull() && !data.Serverip.IsUnknown() {
		contentinspectionaction.Serverip = data.Serverip.ValueString()
	}
	if !data.Servername.IsNull() && !data.Servername.IsUnknown() {
		contentinspectionaction.Servername = data.Servername.ValueString()
	}
	if !data.Serverport.IsNull() && !data.Serverport.IsUnknown() {
		contentinspectionaction.Serverport = utils.IntPtr(int(data.Serverport.ValueInt64()))
	}
	if !data.Wasmprofilename.IsNull() && !data.Wasmprofilename.IsUnknown() {
		contentinspectionaction.Wasmprofilename = data.Wasmprofilename.ValueString()
	}

	return contentinspectionaction
}

func contentinspectionactionSetAttrFromGet(ctx context.Context, data *ContentinspectionactionResourceModel, getResponseData map[string]interface{}) *ContentinspectionactionResourceModel {
	tflog.Debug(ctx, "In contentinspectionactionSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["icapprofilename"]; ok && val != nil {
		data.Icapprofilename = types.StringValue(val.(string))
	} else {
		data.Icapprofilename = types.StringNull()
	}
	if val, ok := getResponseData["ifserverdown"]; ok && val != nil {
		data.Ifserverdown = types.StringValue(val.(string))
	} else {
		data.Ifserverdown = types.StringNull()
	}
	if val, ok := getResponseData["name"]; ok && val != nil {
		data.Name = types.StringValue(val.(string))
	} else {
		data.Name = types.StringNull()
	}
	if val, ok := getResponseData["serverip"]; ok && val != nil {
		data.Serverip = types.StringValue(val.(string))
	} else {
		data.Serverip = types.StringNull()
	}
	if val, ok := getResponseData["servername"]; ok && val != nil {
		data.Servername = types.StringValue(val.(string))
	} else {
		data.Servername = types.StringNull()
	}
	if val, ok := getResponseData["serverport"]; ok && val != nil {
		if intVal, err := utils.ConvertToInt64(val); err == nil {
			data.Serverport = types.Int64Value(intVal)
		}
	} else if data.Serverport.IsUnknown() {
		// Preserve a configured value; only null the port when it was never set.
		data.Serverport = types.Int64Null()
	}
	if val, ok := getResponseData["type"]; ok && val != nil {
		data.Type = types.StringValue(val.(string))
	} else {
		data.Type = types.StringNull()
	}
	if val, ok := getResponseData["wasmprofilename"]; ok && val != nil {
		data.Wasmprofilename = types.StringValue(val.(string))
	} else {
		data.Wasmprofilename = types.StringNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute - use plain value as ID
	data.Id = types.StringValue(fmt.Sprintf("%v", data.Name.ValueString()))

	return data
}
