package sslvserver_sslcertkey_binding

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SslvserverSslcertkeyBindingDataSource)(nil)

func SSlvserverSslcertkeyBindingDataSource() datasource.DataSource {
	return &SslvserverSslcertkeyBindingDataSource{}
}

type SslvserverSslcertkeyBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslvserverSslcertkeyBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcertkey_binding"
}

func (d *SslvserverSslcertkeyBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslvserverSslcertkeyBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslvserverSslcertkeyBindingDataSourceSchema()
}

func (d *SslvserverSslcertkeyBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslvserverSslcertkeyBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vservername_Name := data.Vservername.ValueString()
	ca_Name := data.Ca
	certkeyname_Name := data.Certkeyname
	crlcheck_Name := data.Crlcheck
	ocspcheck_Name := data.Ocspcheck
	snicert_Name := data.Snicert

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcertkey_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcertkey_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslvserver_sslcertkey_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check ca
		if val, ok := v["ca"].(bool); ok {
			if ca_Name.IsNull() || val != ca_Name.ValueBool() {
				match = false
				continue
			}
		} else if !ca_Name.IsNull() {
			match = false
			continue
		}

		// Check certkeyname
		if val, ok := v["certkeyname"].(string); ok {
			if certkeyname_Name.IsNull() || val != certkeyname_Name.ValueString() {
				match = false
				continue
			}
		} else if !certkeyname_Name.IsNull() {
			match = false
			continue
		}

		// Check crlcheck
		if val, ok := v["crlcheck"].(string); ok {
			if crlcheck_Name.IsNull() || val != crlcheck_Name.ValueString() {
				match = false
				continue
			}
		} else if !crlcheck_Name.IsNull() {
			match = false
			continue
		}

		// Check ocspcheck
		if val, ok := v["ocspcheck"].(string); ok {
			if ocspcheck_Name.IsNull() || val != ocspcheck_Name.ValueString() {
				match = false
				continue
			}
		} else if !ocspcheck_Name.IsNull() {
			match = false
			continue
		}

		// Check snicert
		if val, ok := v["snicert"].(bool); ok {
			if snicert_Name.IsNull() || val != snicert_Name.ValueBool() {
				match = false
				continue
			}
		} else if !snicert_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslvserver_sslcertkey_binding with ca %s not found", ca_Name))
		return
	}

	sslvserver_sslcertkey_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
