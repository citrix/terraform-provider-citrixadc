package sslservicegroup_sslcacertbundle_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslservicegroupSslcacertbundleBindingDataSource)(nil)

func SSlservicegroupSslcacertbundleBindingDataSource() datasource.DataSource {
	return &SslservicegroupSslcacertbundleBindingDataSource{}
}

type SslservicegroupSslcacertbundleBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslservicegroupSslcacertbundleBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslservicegroup_sslcacertbundle_binding"
}

func (d *SslservicegroupSslcacertbundleBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslservicegroupSslcacertbundleBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslservicegroupSslcacertbundleBindingDataSourceSchema()
}

func (d *SslservicegroupSslcacertbundleBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslservicegroupSslcacertbundleBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	servicegroupname_Name := data.Servicegroupname.ValueString()
	cacertbundlename_Name := data.Cacertbundlename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslservicegroup_sslcacertbundle_binding.Type(),
		ResourceName:             servicegroupname_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslservicegroup_sslcacertbundle_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslservicegroup_sslcacertbundle_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check cacertbundlename
		if val, ok := v["cacertbundlename"].(string); ok {
			if cacertbundlename_Name.IsNull() || val != cacertbundlename_Name.ValueString() {
				match = false
				continue
			}
		} else if !cacertbundlename_Name.IsNull() {
			match = false
			continue
		}
		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslservicegroup_sslcacertbundle_binding with cacertbundlename %s not found", cacertbundlename_Name))
		return
	}

	sslservicegroup_sslcacertbundle_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
