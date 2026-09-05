package nsmode

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// NsmodeResourceModel describes the resource data model.
type NsmodeResourceModel struct {
	Id                  types.String `tfsdk:"id"`
	Fr                  types.Bool   `tfsdk:"fr"`
	L2                  types.Bool   `tfsdk:"l2"`
	Usip                types.Bool   `tfsdk:"usip"`
	Cka                 types.Bool   `tfsdk:"cka"`
	Tcpb                types.Bool   `tfsdk:"tcpb"`
	Mbf                 types.Bool   `tfsdk:"mbf"`
	Edge                types.Bool   `tfsdk:"edge"`
	Usnip               types.Bool   `tfsdk:"usnip"`
	L3                  types.Bool   `tfsdk:"l3"`
	Pmtud               types.Bool   `tfsdk:"pmtud"`
	Mediaclassification types.Bool   `tfsdk:"mediaclassification"`
	Sradv               types.Bool   `tfsdk:"sradv"`
	Dradv               types.Bool   `tfsdk:"dradv"`
	Iradv               types.Bool   `tfsdk:"iradv"`
	Sradv6              types.Bool   `tfsdk:"sradv6"`
	Dradv6              types.Bool   `tfsdk:"dradv6"`
	Bridgebpdus         types.Bool   `tfsdk:"bridgebpdus"`
	Ulfd                types.Bool   `tfsdk:"ulfd"`
}

// modePlanModifiers returns the plan modifiers used for every nsmode mode
// attribute. Each mode was Optional+Computed+ForceNew in the SDK v2 resource,
// so a change to a *configured* mode must recreate (Delete is a no-op and
// Create re-applies the enable/disable actions). RequiresReplaceIfConfigured
// preserves that ForceNew contract without forcing a spurious replace on the
// Computed (unconfigured) modes whose values are only ever read back from the
// appliance.
func modePlanModifiers() []planmodifier.Bool {
	return []planmodifier.Bool{
		boolplanmodifier.RequiresReplaceIfConfigured(),
	}
}

func (r *NsmodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmode resource.",
			},
			"fr": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Fast Ramp mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"l2": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Layer 2 mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"usip": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Use Source IP mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"cka": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Client Keep-Alive mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"tcpb": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "TCP Buffering mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"mbf": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "MAC-based forwarding mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"edge": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Edge configuration mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"usnip": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Use Subnet IP mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"l3": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Layer 3 mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"pmtud": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Path MTU Discovery mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"mediaclassification": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Media classification mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"sradv": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Static route advertisement mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"dradv": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Dynamic route advertisement mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"iradv": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Intranet route advertisement mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"sradv6": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "IPv6 static route advertisement mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"dradv6": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "IPv6 dynamic route advertisement mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"bridgebpdus": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Bridge BPDUs mode.",
				PlanModifiers: modePlanModifiers(),
			},
			"ulfd": schema.BoolAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Use Layer 2 mode for IPv4 packets.",
				PlanModifiers: modePlanModifiers(),
			},
		},
	}
}

// nsmodeSetAttrFromGet copies the appliance state into the model. The nsmode
// GET always returns every mode as a JSON boolean, so the primary branch is
// what always fires. The else-branch is guarded to only null a value when it is
// still Unknown (so a first-create Unknown resolves cleanly); a known/configured
// value is never clobbered (omit-on-default guard).
func nsmodeSetAttrFromGet(ctx context.Context, data *NsmodeResourceModel, getResponseData map[string]interface{}) *NsmodeResourceModel {
	tflog.Debug(ctx, "In nsmodeSetAttrFromGet Function")

	if val, ok := getResponseData["fr"].(bool); ok {
		data.Fr = types.BoolValue(val)
	} else if data.Fr.IsUnknown() {
		data.Fr = types.BoolNull()
	}
	if val, ok := getResponseData["l2"].(bool); ok {
		data.L2 = types.BoolValue(val)
	} else if data.L2.IsUnknown() {
		data.L2 = types.BoolNull()
	}
	if val, ok := getResponseData["usip"].(bool); ok {
		data.Usip = types.BoolValue(val)
	} else if data.Usip.IsUnknown() {
		data.Usip = types.BoolNull()
	}
	if val, ok := getResponseData["cka"].(bool); ok {
		data.Cka = types.BoolValue(val)
	} else if data.Cka.IsUnknown() {
		data.Cka = types.BoolNull()
	}
	if val, ok := getResponseData["tcpb"].(bool); ok {
		data.Tcpb = types.BoolValue(val)
	} else if data.Tcpb.IsUnknown() {
		data.Tcpb = types.BoolNull()
	}
	if val, ok := getResponseData["mbf"].(bool); ok {
		data.Mbf = types.BoolValue(val)
	} else if data.Mbf.IsUnknown() {
		data.Mbf = types.BoolNull()
	}
	if val, ok := getResponseData["edge"].(bool); ok {
		data.Edge = types.BoolValue(val)
	} else if data.Edge.IsUnknown() {
		data.Edge = types.BoolNull()
	}
	if val, ok := getResponseData["usnip"].(bool); ok {
		data.Usnip = types.BoolValue(val)
	} else if data.Usnip.IsUnknown() {
		data.Usnip = types.BoolNull()
	}
	if val, ok := getResponseData["l3"].(bool); ok {
		data.L3 = types.BoolValue(val)
	} else if data.L3.IsUnknown() {
		data.L3 = types.BoolNull()
	}
	if val, ok := getResponseData["pmtud"].(bool); ok {
		data.Pmtud = types.BoolValue(val)
	} else if data.Pmtud.IsUnknown() {
		data.Pmtud = types.BoolNull()
	}
	if val, ok := getResponseData["mediaclassification"].(bool); ok {
		data.Mediaclassification = types.BoolValue(val)
	} else if data.Mediaclassification.IsUnknown() {
		data.Mediaclassification = types.BoolNull()
	}
	if val, ok := getResponseData["sradv"].(bool); ok {
		data.Sradv = types.BoolValue(val)
	} else if data.Sradv.IsUnknown() {
		data.Sradv = types.BoolNull()
	}
	if val, ok := getResponseData["dradv"].(bool); ok {
		data.Dradv = types.BoolValue(val)
	} else if data.Dradv.IsUnknown() {
		data.Dradv = types.BoolNull()
	}
	if val, ok := getResponseData["iradv"].(bool); ok {
		data.Iradv = types.BoolValue(val)
	} else if data.Iradv.IsUnknown() {
		data.Iradv = types.BoolNull()
	}
	if val, ok := getResponseData["sradv6"].(bool); ok {
		data.Sradv6 = types.BoolValue(val)
	} else if data.Sradv6.IsUnknown() {
		data.Sradv6 = types.BoolNull()
	}
	if val, ok := getResponseData["dradv6"].(bool); ok {
		data.Dradv6 = types.BoolValue(val)
	} else if data.Dradv6.IsUnknown() {
		data.Dradv6 = types.BoolNull()
	}
	if val, ok := getResponseData["bridgebpdus"].(bool); ok {
		data.Bridgebpdus = types.BoolValue(val)
	} else if data.Bridgebpdus.IsUnknown() {
		data.Bridgebpdus = types.BoolNull()
	}
	if val, ok := getResponseData["ulfd"].(bool); ok {
		data.Ulfd = types.BoolValue(val)
	} else if data.Ulfd.IsUnknown() {
		data.Ulfd = types.BoolNull()
	}

	// nsmode is a singleton (no unique attributes) - static ID.
	data.Id = types.StringValue("nsmode-config")

	return data
}
