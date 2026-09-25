package nslaslicense_offline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NSLASLicenseOfflineResourceModel describes the resource data model.
type NSLASLicenseOfflineResourceModel struct {
	Id              types.String `tfsdk:"id"`
	EntitlementName types.String `tfsdk:"entitlement_name"`
	RestrictedMode  types.Bool   `tfsdk:"restricted_mode"`
	IsFIPS          types.Bool   `tfsdk:"is_fips"`
	LASSecretsJson  types.String `tfsdk:"las_secrets_json"`
	SshHostPubkey   types.String `tfsdk:"ssh_host_pubkey"`
	LSGUID          types.String `tfsdk:"lsguid"`
	Version         types.String `tfsdk:"version"`
	Build           types.String `tfsdk:"build"`
	LicenseBlob     types.String `tfsdk:"license_blob_path"`
	Status          types.String `tfsdk:"status"`
	LastUpdated     types.String `tfsdk:"last_updated"`
}

func (r *NSLASLicenseOfflineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "NetScaler LAS Offline License resource. This resource generates and applies offline LAS licenses for NetScaler VPX and MPX appliances.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The system generated id for the nslaslicense_offline resource.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"entitlement_name": schema.StringAttribute{
				MarkdownDescription: "Entitlement name for the license (e.g., VPX 10000 Premium).",
				Required:            true,
			},
			"restricted_mode": schema.BoolAttribute{
				MarkdownDescription: "When enabled, uses restricted offline activation (JSON-based) instead of file upload.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"is_fips": schema.BoolAttribute{
				MarkdownDescription: "Whether this is a FIPS-enabled device",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"las_secrets_json": schema.StringAttribute{
				MarkdownDescription: "File path containing LAS secrets (ccid, client, password, las_endpoint, cc_endpoint)",
				Required:            true,
			},
			"ssh_host_pubkey": schema.StringAttribute{
				MarkdownDescription: "SSH host public key (authorized_keys format, e.g. \"ssh-rsa AAAA...\") used to verify the ADC host key for the SCP license transfer. Obtain it once via `ssh-keyscan -t rsa <adc-mgmt-ip>`. Required: the SCP connection is refused if it is empty or does not match the ADC host key, preventing man-in-the-middle disclosure of the nsroot credentials.",
				Required:            true,
			},
			"lsguid": schema.StringAttribute{
				MarkdownDescription: "License Server GUID (computed from device)",
				Computed:            true,
			},
			"version": schema.StringAttribute{
				MarkdownDescription: "Device software version (computed)",
				Computed:            true,
			},
			"build": schema.StringAttribute{
				MarkdownDescription: "Device build number (computed)",
				Computed:            true,
			},
			"license_blob_path": schema.StringAttribute{
				MarkdownDescription: "Path where license blob is saved locally",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "License application status",
				Computed:            true,
			},
			"last_updated": schema.StringAttribute{
				MarkdownDescription: "Timestamp of last update",
				Computed:            true,
			},
		},
	}
}
