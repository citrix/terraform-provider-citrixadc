package sslcipher_sslvserver_bindings

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Backward-compatible migration of the legacy SDKv2
// `citrixadc_sslcipher_sslvserver_bindings` data source. It returns the list of
// SSL vservers (as a comma-separated string) that have the given cipher bound.
// The data-source type name and attributes (ciphername Required, bound_sslvservers
// Computed) are preserved exactly.

var _ datasource.DataSource = (*sslcipherSslvserverBindingsDataSource)(nil)
var _ datasource.DataSourceWithConfigure = (*sslcipherSslvserverBindingsDataSource)(nil)

func SslcipherSslvserverBindingsDataSource() datasource.DataSource {
	return &sslcipherSslvserverBindingsDataSource{}
}

type sslcipherSslvserverBindingsDataSource struct {
	client *service.NitroClient
}

type SslcipherSslvserverBindingsDataSourceModel struct {
	Id               types.String `tfsdk:"id"`
	Ciphername       types.String `tfsdk:"ciphername"`
	BoundSslvservers types.String `tfsdk:"bound_sslvservers"`
}

func (d *sslcipherSslvserverBindingsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslcipher_sslvserver_bindings"
}

func (d *sslcipherSslvserverBindingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *sslcipherSslvserverBindingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Data source that lists the SSL vservers which have a given SSL cipher bound.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Synthetic ID for the data source (equals the cipher name).",
			},
			"ciphername": schema.StringAttribute{
				Required:    true,
				Description: "Name of the SSL cipher / cipher group to search for.",
			},
			"bound_sslvservers": schema.StringAttribute{
				Computed:    true,
				Description: "Comma-separated list of SSL vserver names that have the cipher bound.",
			},
		},
	}
}

func (d *sslcipherSslvserverBindingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "In SslcipherSslvserverBindingsDataSource Read")
	var data SslcipherSslvserverBindingsDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	ciphername := data.Ciphername.ValueString()

	sslvserverArr, err := d.client.FindResourceArrayWithParams(service.FindParams{ResourceType: "sslvserver"})
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver list, got error: %s", err))
		return
	}

	boundSslvservers := make([]string, 0)
	for _, sslvserver := range sslvserverArr {
		vservername, _ := sslvserver["vservername"].(string)
		bindingArr, err := d.client.FindResourceArrayWithParams(service.FindParams{
			ResourceType:             "sslvserver_sslciphersuite_binding",
			ResourceName:             vservername,
			ResourceMissingErrorCode: 461,
		})
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Error during FindResourceArrayWithParams: %s", err))
			return
		}
		for _, v := range bindingArr {
			if name, ok := v["ciphername"].(string); ok && name == ciphername {
				boundSslvservers = append(boundSslvservers, vservername)
			}
		}
	}

	data.Id = types.StringValue(ciphername)
	data.BoundSslvservers = types.StringValue(strings.Join(boundSslvservers, ","))

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
