package subscriberprofile

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*SubscriberprofileDataSource)(nil)

func SUbscriberprofileDataSource() datasource.DataSource {
	return &SubscriberprofileDataSource{}
}

type SubscriberprofileDataSource struct {
	client *service.NitroClient
}

func (d *SubscriberprofileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_subscriberprofile"
}

func (d *SubscriberprofileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *SubscriberprofileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = SubscriberprofileDataSourceSchema()
}

func (d *SubscriberprofileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SubscriberprofileResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Case 3: Array filter without parent ID

	ip_Name := data.Ip.ValueString()

	vlan_Name := data.Vlan.ValueInt64()

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             "subscriberprofile",
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read subscriberprofile, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "subscriberprofile returned empty array")
		return
	}

	// Iterate through results to find the one with the right id
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check IP match
		if v["ip"].(string) != ip_Name {
			match = false
		}

		// Check VLAN match
		if vlan_Name != 0 {
			vlanVal, ok := v["vlan"]
			if !ok {
				// vlan field doesn't exist in response
				match = false
			} else if vlanVal.(string) != fmt.Sprintf("%d", vlan_Name) {
				// vlan exists but doesn't match
				match = false
			}
		} else {
			// If vlan_Name is 0, only match entries without vlan field
			if _, ok := v["vlan"]; ok {
				match = false
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("subscriberprofile with ip %s not found", ip_Name))
		return
	}

	subscriberprofileSetAttrFromGet(ctx, &data, dataArr[foundIndex])

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
