package ipsecalgsession

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*IpsecalgsessionDataSource)(nil)

func IPsecalgsessionDataSource() datasource.DataSource {
	return &IpsecalgsessionDataSource{}
}

type IpsecalgsessionDataSource struct {
	client *service.NitroClient
}

func (d *IpsecalgsessionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsecalgsession"
}

func (d *IpsecalgsessionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *IpsecalgsessionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = IpsecalgsessionDataSourceSchema()
}

func (d *IpsecalgsessionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpsecalgsessionResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// ipsecalgsession has no get-by-name endpoint; only get(all). Fetch the whole
	// session table and filter locally on the supplied filter attributes.
	destip_Name := data.Destip
	destipalg_Name := data.DestipAlg
	natip_Name := data.Natip
	natipalg_Name := data.NatipAlg
	sourceip_Name := data.Sourceip
	sourceipalg_Name := data.SourceipAlg

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Ipsecalgsession.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read ipsecalgsession, got error: %s", err))
		return
	}

	// An empty session table is valid (no active IPSec ALG sessions). Return an
	// empty result gracefully instead of hard-failing.
	if len(dataArr) == 0 {
		data.Id = types.StringValue("ipsecalgsession")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	// Iterate through results to find the one matching the supplied filters.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check destip
		if !destip_Name.IsNull() {
			if val, ok := v["destip"].(string); ok {
				if val != destip_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check destip_alg
		if !destipalg_Name.IsNull() {
			if val, ok := v["destip_alg"].(string); ok {
				if val != destipalg_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check natip
		if !natip_Name.IsNull() {
			if val, ok := v["natip"].(string); ok {
				if val != natip_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check natip_alg
		if !natipalg_Name.IsNull() {
			if val, ok := v["natip_alg"].(string); ok {
				if val != natipalg_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check sourceip
		if !sourceip_Name.IsNull() {
			if val, ok := v["sourceip"].(string); ok {
				if val != sourceip_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check sourceip_alg
		if !sourceipalg_Name.IsNull() {
			if val, ok := v["sourceip_alg"].(string); ok {
				if val != sourceipalg_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		if match {
			foundIndex = i
			break
		}
	}

	// No matching session for the supplied filters. Empty is valid; return without
	// hard-failing.
	if foundIndex == -1 {
		data.Id = types.StringValue("ipsecalgsession")
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
		return
	}

	ipsecalgsessionSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Datasource has no Create; compose a synthetic ID here.
	data.Id = types.StringValue("ipsecalgsession")

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
