package sslvserver_sslcertkeybundle_binding

import (
	"context"
	"fmt"
	"strings"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*SslvserverSslcertkeybundleBindingDataSource)(nil)

func SSlvserverSslcertkeybundleBindingDataSource() datasource.DataSource {
	return &SslvserverSslcertkeybundleBindingDataSource{}
}

type SslvserverSslcertkeybundleBindingDataSource struct {
	client *service.NitroClient
}

func (d *SslvserverSslcertkeybundleBindingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sslvserver_sslcertkeybundle_binding"
}

func (d *SslvserverSslcertkeybundleBindingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SslvserverSslcertkeybundleBindingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SslvserverSslcertkeybundleBindingDataSourceSchema()
}

func (d *SslvserverSslcertkeybundleBindingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SslvserverSslcertkeybundleBindingResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Case 4: Array filter with parent ID
	vservername_Name := data.Vservername.ValueString()
	certkeybundlename_Name := data.Certkeybundlename
	snicertkeybundle_Name := data.Snicertkeybundle

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Sslvserver_sslcertkeybundle_binding.Type(),
		ResourceName:             vservername_Name,
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read sslvserver_sslcertkeybundle_binding, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "sslvserver_sslcertkeybundle_binding returned empty array.")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check certkeybundlename
		if val, ok := v["certkeybundlename"].(string); ok {
			if certkeybundlename_Name.IsNull() || val != certkeybundlename_Name.ValueString() {
				match = false
				continue
			}
		} else if !certkeybundlename_Name.IsNull() {
			match = false
			continue
		}

		// Check snicertkeybundle
		if val, ok := v["snicertkeybundle"].(bool); ok {
			if snicertkeybundle_Name.IsNull() || val != snicertkeybundle_Name.ValueBool() {
				match = false
				continue
			}
		} else if !snicertkeybundle_Name.IsNull() {
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
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("sslvserver_sslcertkeybundle_binding with certkeybundlename %s not found", certkeybundlename_Name))
		return
	}

	sslvserver_sslcertkeybundle_bindingSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Datasource has no Create; set the composite ID here (Pattern 6).
	// Composite ID is vservername,certkeybundlename (snicertkeybundle is not part of the ID).
	idParts := []string{}
	idParts = append(idParts, fmt.Sprintf("certkeybundlename:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Certkeybundlename.ValueString()))))
	idParts = append(idParts, fmt.Sprintf("vservername:%s", utils.UrlEncode(fmt.Sprintf("%v", data.Vservername.ValueString()))))
	data.Id = types.StringValue(strings.Join(idParts, ","))

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
