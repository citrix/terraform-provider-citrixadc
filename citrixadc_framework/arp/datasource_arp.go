package arp

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*ArpDataSource)(nil)

func ARpDataSource() datasource.DataSource {
	return &ArpDataSource{}
}

type ArpDataSource struct {
	client *service.NitroClient
}

func (d *ArpDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_arp"
}

func (d *ArpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *ArpDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ArpDataSourceSchema()
}

func (d *ArpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ArpResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	ipaddress_Name := data.Ipaddress.ValueString()

	var ownernode_Name int64
	if !data.Ownernode.IsNull() {
		ownernode_Name = data.Ownernode.ValueInt64()
	} else {
		ownernode_Name = 0
	}

	var td_Name int64
	if !data.Td.IsNull() {
		td_Name = data.Td.ValueInt64()
	} else {
		td_Name = 0
	}

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Arp.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read arp, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "arp returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {

		match := true

		if v["ipaddress"].(string) != ipaddress_Name {
			match = false
		}

		// Check ownernode with nil handling
		ownernodeVal := int64(0)
		if v["ownernode"] != nil {
			ownernodeVal, err = utils.ConvertToInt64(v["ownernode"])
		}
		if ownernodeVal != ownernode_Name {
			match = false
		}

		// Check td with nil handling
		tdVal := int64(0)
		if v["td"] != nil {
			tdVal, err = utils.ConvertToInt64(v["td"])
		}
		if tdVal != td_Name {
			match = false
		}

		if match {
			foundIndex = i
			break
		}

	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("arp with ipaddress %s not found", ipaddress_Name))
		return
	}

	arpSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
