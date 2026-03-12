package contentinspectionpolicylabel

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ContentinspectionpolicylabelDataSource)(nil)

func COntentinspectionpolicylabelDataSource() datasource.DataSource {
	return &ContentinspectionpolicylabelDataSource{}
}

type ContentinspectionpolicylabelDataSource struct {
	client *service.NitroClient
}

func (d *ContentinspectionpolicylabelDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contentinspectionpolicylabel"
}

func (d *ContentinspectionpolicylabelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ContentinspectionpolicylabelDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ContentinspectionpolicylabelDataSourceSchema()
}

func (d *ContentinspectionpolicylabelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ContentinspectionpolicylabelResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 2: Find with single ID attribute
	labelname_Name := data.Labelname.ValueString()

	var getResponseData map[string]interface{}
	var err error

	getResponseData, err = d.client.FindResource(service.Contentinspectionpolicylabel.Type(), labelname_Name)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read contentinspectionpolicylabel, got error: %s", err))
		return
	}

	contentinspectionpolicylabelSetAttrFromGet(ctx, &data, getResponseData)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
