package policyurlset

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

var _ datasource.DataSource = (*PolicyurlsetDataSource)(nil)

func POlicyurlsetDataSource() datasource.DataSource {
	return &PolicyurlsetDataSource{}
}

type PolicyurlsetDataSource struct {
	client *service.NitroClient
}

func (d *PolicyurlsetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_policyurlset"
}

func (d *PolicyurlsetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *PolicyurlsetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = PolicyurlsetDataSourceSchema()
}

func (d *PolicyurlsetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data PolicyurlsetResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// An imported urlset is only listed by the filtered GET
	// /policyurlset?args=imported:true (Pattern 15), not by a plain GET
	// /policyurlset/<name> (which returns errorcode 258). Fetch that list and
	// match by name.
	policyurlsetName := data.Name.ValueString()

	getResponseData, err := findImportedPolicyurlsetByName(d.client, policyurlsetName)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read policyurlset, got error: %s", err))
		return
	}

	// Resource is missing
	if getResponseData == nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("policyurlset %s not found", policyurlsetName))
		return
	}

	policyurlsetSetAttrFromGetForDatasource(ctx, &data, getResponseData)
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
