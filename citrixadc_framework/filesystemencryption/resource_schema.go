package filesystemencryption

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/utility"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// FilesystemencryptionResourceModel describes the resource data model.
type FilesystemencryptionResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Nodeid              types.Int64  `tfsdk:"nodeid"`
	Ntimes0flash        types.Int64  `tfsdk:"ntimes0flash"`
	Ntimes0var          types.Int64  `tfsdk:"ntimes0var"`
	Passphrase          types.String `tfsdk:"passphrase"`
	PassphraseWo        types.String `tfsdk:"passphrase_wo"`
	PassphraseWoVersion types.Int64  `tfsdk:"passphrase_wo_version"`
	Supportedstate      types.String `tfsdk:"supportedstate"`
	Effectivestate      types.String `tfsdk:"effectivestate"`
}

func (r *FilesystemencryptionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the filesystemencryption resource.",
			},
			// GET-only filter argument (args=nodeid:<...>) - not an enable/disable payload field (Pattern 15).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			// CLI-mandatory for enable/disable action - Required (no Default allowed on Required attrs).
			"ntimes0flash": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of times /flash directory has to be written with 0s.",
			},
			// CLI-mandatory for enable/disable action - Required.
			"ntimes0var": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of times /var directory has to be written with 0s.",
			},
			// CLI-mandatory secret for enable/disable action. ValidateConfig enforces
			// at-least-one-of passphrase / passphrase_wo (Pattern 17).
			"passphrase": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Encryption Passphrase.",
			},
			"passphrase_wo": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
				WriteOnly: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Encryption Passphrase.",
			},
			"passphrase_wo_version": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Increment this version to signal a passphrase_wo update.",
			},
			// Read-only capability/state attributes from GET.
			"supportedstate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the supported state of File System Encryption. Possible values = DISABLED, ENABLED, UNKNOWN.",
			},
			"effectivestate": schema.StringAttribute{
				Computed:    true,
				Description: "Get the current encrypted state of the File System. Possible values = ENABLED, DISABLED.",
			},
		},
	}
}

// filesystemencryptionGetThePayloadFromthePlan builds the enable/disable action payload.
// nodeid is intentionally excluded (GET-only filter, Pattern 15).
func filesystemencryptionGetThePayloadFromthePlan(ctx context.Context, data *FilesystemencryptionResourceModel) utility.Filesystemencryption {
	tflog.Debug(ctx, "In filesystemencryptionGetThePayloadFromthePlan Function")

	filesystemencryption := utility.Filesystemencryption{}
	if !data.Ntimes0flash.IsNull() && !data.Ntimes0flash.IsUnknown() {
		filesystemencryption.Ntimes0flash = utils.IntPtr(int(data.Ntimes0flash.ValueInt64()))
	}
	if !data.Ntimes0var.IsNull() && !data.Ntimes0var.IsUnknown() {
		filesystemencryption.Ntimes0var = utils.IntPtr(int(data.Ntimes0var.ValueInt64()))
	}
	if !data.Passphrase.IsNull() && !data.Passphrase.IsUnknown() {
		filesystemencryption.Passphrase = data.Passphrase.ValueString()
	}
	// Skip write-only attribute: passphrase_wo (applied from config)
	// Skip version tracker attribute: passphrase_wo_version
	// Skip GET-only filter attribute: nodeid

	return filesystemencryption
}

func filesystemencryptionGetThePayloadFromtheConfig(ctx context.Context, data *FilesystemencryptionResourceModel, payload *utility.Filesystemencryption) {
	tflog.Debug(ctx, "In filesystemencryptionGetThePayloadFromtheConfig Function")

	// Handle write-only secret attribute: passphrase_wo -> passphrase
	if !data.PassphraseWo.IsNull() {
		passphraseWo := data.PassphraseWo.ValueString()
		if passphraseWo != "" {
			payload.Passphrase = passphraseWo
		}
	}
}

// filesystemencryptionSetAttrFromGet populates read-only state from the nameless singleton GET.
// It preserves the write-only/action-input attributes (never echoed back) and never
// refreshes passphrase (write-only secret).
func filesystemencryptionSetAttrFromGet(ctx context.Context, data *FilesystemencryptionResourceModel, getResponseData map[string]interface{}) *FilesystemencryptionResourceModel {
	tflog.Debug(ctx, "In filesystemencryptionSetAttrFromGet Function")

	// Read-only state fields from GET.
	if val, ok := getResponseData["supportedstate"]; ok && val != nil {
		data.Supportedstate = types.StringValue(val.(string))
	}
	if val, ok := getResponseData["effectivestate"]; ok && val != nil {
		data.Effectivestate = types.StringValue(val.(string))
	}
	// nodeid, ntimes0flash, ntimes0var, passphrase are action inputs / write-only:
	// not refreshed from GET, preserved from plan/state.

	// Static synthetic ID.
	data.Id = types.StringValue("filesystemencryption-config")

	return data
}
