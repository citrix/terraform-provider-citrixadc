package systemextramgmtcpu

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/system"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// SystemextramgmtcpuResourceModel describes the resource data model.
// It preserves the exact SDK v2 user-facing contract (attribute names/types).
type SystemextramgmtcpuResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	Reboot                types.Bool   `tfsdk:"reboot"`
	ReachableTimeout      types.String `tfsdk:"reachable_timeout"`
	ReachablePollDelay    types.String `tfsdk:"reachable_poll_delay"`
	ReachablePollInterval types.String `tfsdk:"reachable_poll_interval"`
	ReachablePollTimeout  types.String `tfsdk:"reachable_poll_timeout"`
}

func (r *SystemextramgmtcpuResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the systemextramgmtcpu resource.",
			},
			// SDK v2: Required + ForceNew -> Required + RequiresReplace().
			// Drives the enable/disable action performed at Create time.
			"enabled": schema.BoolAttribute{
				Required:    true,
				Description: "Boolean value indicating the effective state of the extra management CPU. Set to true to enable, false to disable.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			// SDK v2: Optional, Default true, NOT ForceNew.
			// Whether to reboot the ADC (and wait until reachable) after the action.
			"reboot": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
				Description: "Reboot the ADC instance after applying the extra management CPU configuration.",
			},
			// SDK v2: Optional, Default "10m", ForceNew.
			"reachable_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("10m"),
				Description: "Maximum duration to wait for the ADC to become reachable after reboot.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// SDK v2: Optional, Default "60s", ForceNew.
			"reachable_poll_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "Initial delay before polling the ADC for reachability after reboot.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// SDK v2: Optional, Default "60s", ForceNew.
			"reachable_poll_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("60s"),
				Description: "Interval between reachability polls of the ADC after reboot.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			// SDK v2: Optional, Default "20s", ForceNew.
			"reachable_poll_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Default:     stringdefault.StaticString("20s"),
				Description: "Per-request timeout for each reachability poll of the ADC after reboot.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func systemextramgmtcpuGetThePayloadFromtheConfig(ctx context.Context, data *SystemextramgmtcpuResourceModel) system.Systemextramgmtcpu {
	tflog.Debug(ctx, "In systemextramgmtcpuGetThePayloadFromtheConfig Function")

	// The enable/disable action does not carry any body fields (mirrors SDK v2,
	// which passed an empty system.Systemextramgmtcpu{}). configuredstate /
	// effectivestate are read-only and nodeid is unused here.
	systemextramgmtcpu := system.Systemextramgmtcpu{}

	return systemextramgmtcpu
}

func systemextramgmtcpuSetAttrFromGet(ctx context.Context, data *SystemextramgmtcpuResourceModel, getResponseData map[string]interface{}) *SystemextramgmtcpuResourceModel {
	tflog.Debug(ctx, "In systemextramgmtcpuSetAttrFromGet Function")

	// effectivestate is always present in the GET response, so mapping it to the
	// bool `enabled` is safe drift detection (this mirrors SDK v2 readFunc).
	if val, ok := getResponseData["effectivestate"]; ok && val != nil && val == "ENABLED" {
		data.Enabled = types.BoolValue(true)
	} else {
		data.Enabled = types.BoolValue(false)
	}

	// Singleton resource -> static ID (matches the datasource contract:
	// "systemextramgmtcpu-config").
	data.Id = types.StringValue("systemextramgmtcpu-config")

	// reboot / reachable_* are Terraform-only knobs never returned by NITRO;
	// they are intentionally left untouched so prior state/plan values survive.

	return data
}
