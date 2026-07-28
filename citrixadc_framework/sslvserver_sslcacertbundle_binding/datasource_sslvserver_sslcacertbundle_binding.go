package sslvserver_sslcacertbundle_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SslvserverSslcacertbundleBindingDataSource)(nil)

func SSlvserverSslcacertbundleBindingDataSource() datasource.DataSource {
	return &SslvserverSslcacertbundleBindingDataSource{}
}

type SslvserverSslcacertbundleBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslvserverSslcacertbundleBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcacertbundle_binding"
}

func (d *SslvserverSslcacertbundleBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslvserverSslcacertbundleBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslvserverSslcacertbundleBindingDataSourceSchema()
}

func (d *SslvserverSslcacertbundleBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslvserverSslcacertbundleBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vservername_Name := data.Vservername.ValueString()
	cacertbundlename_Name := data.Cacertbundlename

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcacertbundle_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcacertbundle_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslvserver_sslcacertbundle_binding returned empty array.")
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslvserver_sslcacertbundle_binding with cacertbundlename %s not found", cacertbundlename_Name))
		return
	}

	sslvserver_sslcacertbundle_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Datasource has no Create; set the composite ID here (Pattern 6).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("cacertbundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Cacertbundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
