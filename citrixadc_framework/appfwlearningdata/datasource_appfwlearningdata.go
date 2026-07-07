package appfwlearningdata

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// appfwlearningdata datasource is a BEST-EFFORT reader for the App-Firewall
// learned-data table. NITRO exposes get(all) only (there is no per-object GET);
// this datasource reads the table and surfaces the FIRST matching entry. When
// profilename (and optionally securitycheck) are supplied they are passed as
// NITRO args to scope the lookup; otherwise the whole table is read.
var _ datasource.DataSource = (*AppfwlearningdataDataSource)(nil)

func APpfwlearningdataDataSource() datasource.DataSource {
	return &AppfwlearningdataDataSource{}
}

type AppfwlearningdataDataSource struct {
	client *service.NitroClient
}

func (d *AppfwlearningdataDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_appfwlearningdata"
}

func (d *AppfwlearningdataDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *AppfwlearningdataDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = AppfwlearningdataDataSourceSchema()
}

func (d *AppfwlearningdataDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data AppfwlearningdataDataSourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var results []map[string]interface{}
	var err error

	// get(all) requires profilename/securitycheck as args on the ADC. When set,
	// scope the lookup; otherwise fall back to reading the full table (best-effort).
	if !data.Profilename.IsNull() && !data.Profilename.IsUnknown() {
		argsMap := map[string]string{
			"profilename": data.Profilename.ValueString(),
		}
		if !data.Securitycheck.IsNull() && !data.Securitycheck.IsUnknown() {
			argsMap["securitycheck"] = data.Securitycheck.ValueString()
		}
		findParams := service.FindParams{
			ResourceType: service.Appfwlearningdata.Type(),
			ArgsMap:      argsMap,
		}
		results, err = d.client.FindResourceArrayWithParams(findParams)
	} else {
		results, err = d.client.FindAllResources(service.Appfwlearningdata.Type())
	}

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read appfwlearningdata, got error: %s", err))
		return
	}

	if len(results) == 0 {
		resp.Diagnostics.AddError("Not Found", "No appfwlearningdata entries were returned by the ADC for the given lookup")
		return
	}

	// Best-effort: expose the first learned-data entry.
	appfwlearningdataSetAttrFromGet(ctx, &data, results[0])

	// Synthetic ID for the datasource read.
	data.Id = types.StringValue("appfwlearningdata-config")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
