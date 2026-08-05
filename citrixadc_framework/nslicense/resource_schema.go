package nslicense

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// NslicenseResourceModel describes the resource data model.
//
// The nslicense RESOURCE is a custom resource (it is NOT a plain NITRO CRUD
// object). It uploads a license file to the ADC over SSH/SFTP and, optionally,
// reboots the appliance and polls until it is reachable again. Its user-facing
// contract mirrors the legacy SDK v2 resource
// (citrixadc/resource_citrixadc_nslicense.go) exactly so existing state and
// configs keep working after migration.
type NslicenseResourceModel struct {
	Id            types.String `tfsdk:"id"`
	LicenseFile   types.String `tfsdk:"license_file"`
	SshHost       types.String `tfsdk:"ssh_host"`
	SshUsername   types.String `tfsdk:"ssh_username"`
	SshPassword   types.String `tfsdk:"ssh_password"`
	SshPort       types.Int64  `tfsdk:"ssh_port"`
	SshHostPubkey types.String `tfsdk:"ssh_host_pubkey"`
	Reboot        types.Bool   `tfsdk:"reboot"`
	PollDelay     types.String `tfsdk:"poll_delay"`
	PollInterval  types.String `tfsdk:"poll_interval"`
	PollTimeout   types.String `tfsdk:"poll_timeout"`
}

func (r *NslicenseResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     1,
		Description: "Uploads a Citrix ADC license file to the appliance over SSH/SFTP and optionally reboots it.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nslicense resource (the license file name).",
			},
			// SDK v2: Required + ForceNew
			"license_file": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Description: "Name of the license file to upload to /nsconfig/license/ on the ADC.",
			},
			// SDK v2: Optional (no Computed, no Default)
			"ssh_host": schema.StringAttribute{
				Optional:    true,
				Description: "SSH host to connect to. Defaults to the host parsed from the NITRO endpoint.",
			},
			// SDK v2: Optional + Sensitive
			"ssh_username": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "SSH username. Defaults to the provider NITRO username.",
			},
			// SDK v2: Optional + Sensitive
			"ssh_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Description: "SSH password. Defaults to the provider NITRO password.",
			},
			// SDK v2: Optional + Computed (no Default)
			"ssh_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "SSH port. Defaults to 22.",
			},
			// SDK v2: Required
			"ssh_host_pubkey": schema.StringAttribute{
				Required:    true,
				Description: "SSH host public key used to verify the ADC host key.",
			},
			// SDK v2: Optional + Default(true)
			"reboot": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Whether to reboot the ADC after applying the license.",
			},
			// SDK v2: Optional + Default("60s")
			"poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "Delay before the first reachability poll after a reboot.",
			},
			// SDK v2: Optional + Default("60s")
			"poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "Interval between reachability polls after a reboot.",
			},
			// SDK v2: Optional + Default("10s")
			"poll_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10s"),
				Description: "Per-poll HTTP timeout used to test reachability after a reboot.",
			},
		},
	}
}
