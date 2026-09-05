package nsversion

import (
	"context"
	"fmt"
	"strconv"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Backward-compatible migration of the legacy SDKv2 `citrixadc_nsversion` data
// source. The data-source type name and every attribute (name/type/optionality)
// are preserved exactly, so existing configurations continue to work unchanged.

var _ datasource.DataSource = (*nsversionDataSource)(nil)
var _ datasource.DataSourceWithConfigure = (*nsversionDataSource)(nil)

func NsversionDataSource() datasource.DataSource {
	return &nsversionDataSource{}
}

type nsversionDataSource struct {
	client *service.NitroClient
}

type NsversionDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Installedversion types.Bool   `tfsdk:"installedversion"`
	Version          types.String `tfsdk:"version"`
	Mode             types.Int64  `tfsdk:"mode"`
}

func (d *nsversionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsversion"
}

func (d *nsversionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *nsversionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source to retrieve the NetScaler firmware version (nsversion).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "The firmware version string.",
			},
			"installedversion": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set, requests the installed (not running) version from the appliance.",
			},
			"version": schema.StringAttribute{
				Computed:    true,
				Description: "The NetScaler version string.",
			},
			"mode": schema.Int64Attribute{
				Computed:    true,
				Description: "The NetScaler mode.",
			},
		},
	}
}

func (d *nsversionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "In NsversionDataSource Read")
	var data NsversionDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	findParams := service.FindParams{ResourceType: "nsversion"}
	// Mirror SDKv2 d.GetOkExists("installedversion"): only pass the arg when set.
	if !data.Installedversion.IsNull() && !data.Installedversion.IsUnknown() {
		findParams.ArgsMap = map[string]string{
			"installedversion": fmt.Sprintf("%v", data.Installedversion.ValueBool()),
		}
	}

	dataArr, err := d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsversion, got error: %s", err))
		return
	}
	if len(dataArr) != 1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unexpected length of nsversion response: %v", dataArr))
		return
	}
	result := dataArr[0]

	version, _ := result["version"].(string)
	data.Id = types.StringValue(version)
	data.Version = types.StringValue(version)

	if val, ok := result["mode"]; ok && val != nil {
		intVal, err := strconv.Atoi(fmt.Sprintf("%v", val))
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error during Atoi for mode: %s", err))
			return
		}
		data.Mode = types.Int64Value(int64(intVal))
	} else {
		data.Mode = types.Int64Null()
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
