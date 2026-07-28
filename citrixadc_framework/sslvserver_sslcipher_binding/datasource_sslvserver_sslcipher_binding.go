package sslvserver_sslcipher_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SslvserverSslcipherBindingDataSource)(nil)

func SSlvserverSslcipherBindingDataSource() datasource.DataSource {
	return &SslvserverSslcipherBindingDataSource{}
}

type SslvserverSslcipherBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslvserverSslcipherBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcipher_binding"
}

func (d *SslvserverSslcipherBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslvserverSslcipherBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslvserverSslcipherBindingDataSourceSchema()
}

func (d *SslvserverSslcipherBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslvserverSslcipherBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vservername_Name := data.Vservername.ValueString()
	ciphername_Name := data.Ciphername

	var dataArr []map[string]interface{}
	var err error

	// This binding subresource is not reflected over the typed GET on this firmware
	// (returns an empty {"message":"Done"} body); it is only reflected via the umbrella
	// sslvserver_binding endpoint under sslvserver_sslciphersuite_binding[]. Try the
	// typed GET first, then fall back to the umbrella. Mirror the resource Read exactly.
	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcipher_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	if !ciphername_Name.IsNull() && !ciphername_Name.IsUnknown() {
		findParams.FilterMap = map[string]string{"ciphername": ciphername_Name.ValueString()}
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcipher_binding, got error: %s", err))
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		if val, ok := v["ciphername"].(string); ok {
			if !ciphername_Name.IsNull() && val == ciphername_Name.ValueString() {
				foundIndex = i
				break
			}
		}
	}

	if foundIndex != -1 {
		sslvserver_sslcipher_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	} else {
		// Typed GET returned no matching row; fall back to the umbrella endpoint.
		row := findCipherBindingViaUmbrella(ctx, d.client, vservername_Name, ciphername_Name.ValueString(), &resp.Diagnostics)
		if resp.Diagnostics.HasError() {
			return
		}
		if row == nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslvserver_sslcipher_binding with ciphername %s not found", ciphername_Name.ValueString()))
			return
		}
		sslvserver_sslcipher_bindingSetAttrFromGet(ctx, &data, row)
	}

	// Datasource has no Create; set the composite ID here (Pattern 6).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("ciphername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Ciphername.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
