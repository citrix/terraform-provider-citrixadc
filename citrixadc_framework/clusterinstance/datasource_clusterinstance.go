package clusterinstance

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*ClusterinstanceDataSource)(nil)

func CLusterinstanceDataSource() datasource.DataSource {
	return &ClusterinstanceDataSource{}
}

type ClusterinstanceDataSource struct {
	client *service.NitroClient
}

func (d *ClusterinstanceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_clusterinstance"
}

func (d *ClusterinstanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ClusterinstanceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ClusterinstanceDataSourceSchema()
}

func (d *ClusterinstanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ClusterinstanceResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	clid_Name := types.StringValue(fmt.Sprintf("%d", data.Clid.ValueInt64()))

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Clusterinstance.Type(), clid_Name.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read clusterinstance, got error: %s", err))
		return
	}

	clusterinstanceSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
