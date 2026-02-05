package nsfeature

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*NsfeatureDataSource)(nil)

func NSfeatureDataSource() datasource.DataSource {
	return &NsfeatureDataSource{}
}

type NsfeatureDataSource struct {
	client *service.NitroClient
}

func (d *NsfeatureDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_nsfeature"
}

func (d *NsfeatureDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *NsfeatureDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = NsfeatureDataSourceSchema()
}

func (d *NsfeatureDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data NsfeatureResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Use ListEnabledFeatures to match v2 implementation
	featuresData, err := d.client.ListEnabledFeatures()
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read nsfeature, got error: %s", err))
		return
	}

	// Convert features to lowercase for comparison
	enabledFeatures := make([]string, len(featuresData))
	for i, val := range featuresData {
		enabledFeatures[i] = strings.ToLower(val)
	}

	nsfeatureSetAttrFromGet(ctx, &data, enabledFeatures)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
