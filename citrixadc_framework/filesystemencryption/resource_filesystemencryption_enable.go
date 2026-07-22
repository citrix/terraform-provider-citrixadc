package filesystemencryption

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/resource/config/utility"
	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FilesystemencryptionEnableResource{}
var _ resource.ResourceWithConfigure = (*FilesystemencryptionEnableResource)(nil)

// passphrase is a CLI-mandatory secret enforced at plan time via ValidateConfig
// (Pattern 17); this assertion registers the extra interface.
var _ resource.ResourceWithValidateConfig = (*FilesystemencryptionEnableResource)(nil)

func NewFilesystemencryptionEnableResource() resource.Resource {
	return &FilesystemencryptionEnableResource{}
}

// FilesystemencryptionEnableResource models the NITRO filesystemencryption
// `?action=enable` action (POST). enable is a one-shot side-effect action with no
// GET endpoint and no inverse API, so Read/Update/Delete are no-ops. The read-only
// supportedstate/effectivestate fields are intentionally omitted (they belong to
// the citrixadc_filesystemencryption datasource).
type FilesystemencryptionEnableResource struct {
	client *service.NitroClient
}

// FilesystemencryptionEnableResourceModel describes the resource data model.
type FilesystemencryptionEnableResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Nodeid              types.Int64  `tfsdk:"nodeid"`
	Ntimes0flash        types.Int64  `tfsdk:"ntimes0flash"`
	Ntimes0var          types.Int64  `tfsdk:"ntimes0var"`
	Passphrase          types.String `tfsdk:"passphrase"`
	PassphraseWo        types.String `tfsdk:"passphrase_wo"`
	PassphraseWoVersion types.Int64  `tfsdk:"passphrase_wo_version"`
}

func (r *FilesystemencryptionEnableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_filesystemencryption_enable"
}

func (r *FilesystemencryptionEnableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}
	// Set the client for the resource.
	r.client = *req.ProviderData.(**service.NitroClient)
}

// ValidateConfig enforces the CLI-mandatory passphrase secret: at least one of
// passphrase / passphrase_wo must be set (Pattern 17).
func (r *FilesystemencryptionEnableResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var data FilesystemencryptionEnableResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if data.Passphrase.IsNull() && data.PassphraseWo.IsNull() {
		resp.Diagnostics.AddError(
			"Missing required attribute",
			"At least one of \"passphrase\" or \"passphrase_wo\" must be set for filesystemencryption_enable.",
		)
	}
}

func (r *FilesystemencryptionEnableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the filesystemencryption_enable resource.",
			},
			// GET-only filter argument (args=nodeid:<...>) - not an enable payload field (Pattern 15).
			"nodeid": schema.Int64Attribute{
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Unique number that identifies the cluster node.",
			},
			// CLI-mandatory for enable action - Required (no Default allowed on Required attrs).
			"ntimes0flash": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of times /flash directory has to be written with 0s.",
			},
			// CLI-mandatory for enable action - Required.
			"ntimes0var": schema.Int64Attribute{
				Required: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Description: "Number of times /var directory has to be written with 0s.",
			},
			// CLI-mandatory secret for enable action. ValidateConfig enforces
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
		},
	}
}

func (r *FilesystemencryptionEnableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data, config FilesystemencryptionEnableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	// Read write-only attributes from config (they are nullified in plan)
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Enabling filesystemencryption (action-only resource)")
	// Get payload from plan (regular attributes)
	payload := filesystemencryption_enableGetThePayloadFromthePlan(ctx, &data)
	// Add write-only attributes from config to the payload
	filesystemencryption_enableGetThePayloadFromtheConfig(ctx, &config, &payload)

	// NITRO exposes enable as POST ?action=enable. Verb casing matches the URL.
	err := r.client.ActOnResource(service.Filesystemencryption.Type(), &payload, "enable")
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to enable filesystemencryption, got error: %s", err))
		return
	}

	tflog.Trace(ctx, "Enabled filesystemencryption")

	// Synthetic ID for the action-only resource; keeps Read/Delete no-ops
	// addressable by Terraform.
	data.Id = types.StringValue("filesystemencryption_enable")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilesystemencryptionEnableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// enable is a one-shot action. NITRO has no GET endpoint that reports the
	// enable action state, so Read is a pure preserve-state no-op.
	var data FilesystemencryptionEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, "Read is a no-op for filesystemencryption_enable; NITRO has no query endpoint for the action state")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilesystemencryptionEnableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NITRO has no update endpoint for enable; every schema attribute is
	// RequiresReplace, so Terraform never invokes Update for a real change.
	var data, state FilesystemencryptionEnableResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	data.Id = state.Id
	tflog.Debug(ctx, "Update is a no-op for filesystemencryption_enable; NITRO has no update endpoint and all attributes are RequiresReplace")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FilesystemencryptionEnableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// enable is a one-shot side-effect action. There is no inverse NITRO API
	// invoked here (disable is a separate resource). Delete simply removes the
	// resource from Terraform state.
	tflog.Debug(ctx, "Delete is a no-op for filesystemencryption_enable; disable is modelled as a separate resource")
}

// filesystemencryption_enableGetThePayloadFromthePlan builds the enable action
// payload. nodeid is intentionally excluded (GET-only filter, Pattern 15).
func filesystemencryption_enableGetThePayloadFromthePlan(ctx context.Context, data *FilesystemencryptionEnableResourceModel) utility.Filesystemencryption {
	tflog.Debug(ctx, "In filesystemencryption_enableGetThePayloadFromthePlan Function")

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
	// Skip GET-only filter attribute: nodeid (Pattern 15)

	return filesystemencryption
}

func filesystemencryption_enableGetThePayloadFromtheConfig(ctx context.Context, data *FilesystemencryptionEnableResourceModel, payload *utility.Filesystemencryption) {
	tflog.Debug(ctx, "In filesystemencryption_enableGetThePayloadFromtheConfig Function")

	// Handle write-only secret attribute: passphrase_wo -> passphrase
	if !data.PassphraseWo.IsNull() {
		passphraseWo := data.PassphraseWo.ValueString()
		if passphraseWo != "" {
			payload.Passphrase = passphraseWo
		}
	}
}
