package systembackup

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystembackupResourceModel describes the resource data model.
type SystembackupResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Filename         types.String `tfsdk:"filename"`
	Includekernel    types.String `tfsdk:"includekernel"`
	Level            types.String `tfsdk:"level"`
	Skipbackup       types.Bool   `tfsdk:"skipbackup"`
	Uselocaltimezone types.Bool   `tfsdk:"uselocaltimezone"`
}

func (r *SystembackupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systembackup resource.",
			},
			// filename is Required + ForceNew in SDK v2. It is the resource identity
			// (is_get_id/is_delete_id/x-unique-attr). Not Computed.
			"filename": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the backup file(*.tgz) to be restored.",
			},
			// The remaining attributes are create-only args (is_updateable=false).
			// Read is a no-op (SDK v2 schema.Noop), so they must NOT be Computed —
			// an Optional+Computed attribute can never be resolved by a no-op Read and
			// would fail apply with "still indicated an unknown value" (Pattern 13).
			"comment": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Comment specified at the time of creation of the backup file(*.tgz).",
			},
			"includekernel": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to add kernel in the backup file",
			},
			// level had an invalid Default (Default without Computed) and SDK v2
			// defined no default; dropped both flags and left it Optional.
			"level": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Level of data to be backed up.",
			},
			"skipbackup": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "Use this option to skip taking backup during restore operation",
			},
			"uselocaltimezone": schema.BoolAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Description: "This option will create backup file with local timezone timestamp",
			},
		},
	}
}

func systembackupGetThePayloadFromtheConfig(ctx context.Context, data *SystembackupResourceModel) system.Systembackup {
	tflog.Debug(ctx, "In systembackupGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	systembackup := system.Systembackup{}
	if !data.Comment.IsNull() {
		systembackup.Comment = data.Comment.ValueString()
	}
	if !data.Filename.IsNull() {
		systembackup.Filename = data.Filename.ValueString()
	}
	if !data.Includekernel.IsNull() {
		systembackup.Includekernel = data.Includekernel.ValueString()
	}
	if !data.Level.IsNull() {
		systembackup.Level = data.Level.ValueString()
	}
	if !data.Skipbackup.IsNull() {
		systembackup.Skipbackup = data.Skipbackup.ValueBool()
	}
	if !data.Uselocaltimezone.IsNull() {
		systembackup.Uselocaltimezone = data.Uselocaltimezone.ValueBool()
	}

	return systembackup
}

func systembackupSetAttrFromGet(ctx context.Context, data *SystembackupResourceModel, getResponseData map[string]interface{}) *SystembackupResourceModel {
	tflog.Debug(ctx, "In systembackupSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["comment"]; ok && val != nil {
		data.Comment = types.StringValue(val.(string))
	} else {
		data.Comment = types.StringNull()
	}
	if val, ok := getResponseData["filename"]; ok && val != nil {
		data.Filename = types.StringValue(val.(string))
	} else {
		data.Filename = types.StringNull()
	}
	if val, ok := getResponseData["includekernel"]; ok && val != nil {
		data.Includekernel = types.StringValue(val.(string))
	} else {
		data.Includekernel = types.StringNull()
	}
	if val, ok := getResponseData["level"]; ok && val != nil {
		data.Level = types.StringValue(val.(string))
	} else {
		data.Level = types.StringNull()
	}
	if val, ok := getResponseData["skipbackup"]; ok && val != nil {
		data.Skipbackup = types.BoolValue(val.(bool))
	} else {
		data.Skipbackup = types.BoolNull()
	}
	if val, ok := getResponseData["uselocaltimezone"]; ok && val != nil {
		data.Uselocaltimezone = types.BoolValue(val.(bool))
	} else {
		data.Uselocaltimezone = types.BoolNull()
	}

	// Set ID for the resource
	// Case 2: Single unique attribute
	data.Id = types.StringValue(data.Filename.ValueString())

	return data
}
