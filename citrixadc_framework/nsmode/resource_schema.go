package nsmode

import (
	"context"

	"github.com/citrix/adc-nitro-go/resource/config/ns"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

func (r *NsmodeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version: 1,
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The ID of the nsmode resource.",
			},
			"fr": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Fast Ramp mode.",
			},
			"l2": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer 2 mode.",
			},
			"usip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Source IP mode.",
			},
			"cka": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Client Keep-Alive mode.",
			},
			"tcpb": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP Buffering mode.",
			},
			"mbf": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC-based forwarding mode.",
			},
			"edge": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Edge configuration mode.",
			},
			"usnip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Subnet IP mode.",
			},
			"l3": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer 3 mode.",
			},
			"pmtud": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Path MTU Discovery mode.",
			},
			"mediaclassification": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Media classification mode.",
			},
			"sradv": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Static route advertisement mode.",
			},
			"dradv": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Dynamic route advertisement mode.",
			},
			"iradv": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Intranet route advertisement mode.",
			},
			"sradv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 static route advertisement mode.",
			},
			"dradv6": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv6 dynamic route advertisement mode.",
			},
			"bridgebpdus": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bridge BPDUs mode.",
			},
			"ulfd": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use Layer 2 mode for IPv4 packets.",
			},
		},
	}
}

func nsmodeGetThePayloadFromtheConfig(ctx context.Context, data *NsmodeResourceModel) ns.Nsmode {
	tflog.Debug(ctx, "In nsmodeGetThePayloadFromtheConfig Function")

	// Create API request body from the model
	nsmode := ns.Nsmode{}

	return nsmode
}

func nsmodeSetAttrFromGet(ctx context.Context, data *NsmodeResourceModel, getResponseData map[string]interface{}) *NsmodeResourceModel {
	tflog.Debug(ctx, "In nsmodeSetAttrFromGet Function")

	// Convert API response to model
	if val, ok := getResponseData["fr"].(bool); ok {
		data.Fr = types.BoolValue(val)
	} else {
		data.Fr = types.BoolNull()
	}
	if val, ok := getResponseData["l2"].(bool); ok {
		data.L2 = types.BoolValue(val)
	} else {
		data.L2 = types.BoolNull()
	}
	if val, ok := getResponseData["usip"].(bool); ok {
		data.Usip = types.BoolValue(val)
	} else {
		data.Usip = types.BoolNull()
	}
	if val, ok := getResponseData["cka"].(bool); ok {
		data.Cka = types.BoolValue(val)
	} else {
		data.Cka = types.BoolNull()
	}
	if val, ok := getResponseData["tcpb"].(bool); ok {
		data.Tcpb = types.BoolValue(val)
	} else {
		data.Tcpb = types.BoolNull()
	}
	if val, ok := getResponseData["mbf"].(bool); ok {
		data.Mbf = types.BoolValue(val)
	} else {
		data.Mbf = types.BoolNull()
	}
	if val, ok := getResponseData["edge"].(bool); ok {
		data.Edge = types.BoolValue(val)
	} else {
		data.Edge = types.BoolNull()
	}
	if val, ok := getResponseData["usnip"].(bool); ok {
		data.Usnip = types.BoolValue(val)
	} else {
		data.Usnip = types.BoolNull()
	}
	if val, ok := getResponseData["l3"].(bool); ok {
		data.L3 = types.BoolValue(val)
	} else {
		data.L3 = types.BoolNull()
	}
	if val, ok := getResponseData["pmtud"].(bool); ok {
		data.Pmtud = types.BoolValue(val)
	} else {
		data.Pmtud = types.BoolNull()
	}
	if val, ok := getResponseData["mediaclassification"].(bool); ok {
		data.Mediaclassification = types.BoolValue(val)
	} else {
		data.Mediaclassification = types.BoolNull()
	}
	if val, ok := getResponseData["sradv"].(bool); ok {
		data.Sradv = types.BoolValue(val)
	} else {
		data.Sradv = types.BoolNull()
	}
	if val, ok := getResponseData["dradv"].(bool); ok {
		data.Dradv = types.BoolValue(val)
	} else {
		data.Dradv = types.BoolNull()
	}
	if val, ok := getResponseData["iradv"].(bool); ok {
		data.Iradv = types.BoolValue(val)
	} else {
		data.Iradv = types.BoolNull()
	}
	if val, ok := getResponseData["sradv6"].(bool); ok {
		data.Sradv6 = types.BoolValue(val)
	} else {
		data.Sradv6 = types.BoolNull()
	}
	if val, ok := getResponseData["dradv6"].(bool); ok {
		data.Dradv6 = types.BoolValue(val)
	} else {
		data.Dradv6 = types.BoolNull()
	}
	if val, ok := getResponseData["bridgebpdus"].(bool); ok {
		data.Bridgebpdus = types.BoolValue(val)
	} else {
		data.Bridgebpdus = types.BoolNull()
	}
	if val, ok := getResponseData["ulfd"].(bool); ok {
		data.Ulfd = types.BoolValue(val)
	} else {
		data.Ulfd = types.BoolNull()
	}

	// Set ID for the resource
	// Case 1: No unique attributes - static ID
	data.Id = types.StringValue("nsmode-config")

	return data
}
