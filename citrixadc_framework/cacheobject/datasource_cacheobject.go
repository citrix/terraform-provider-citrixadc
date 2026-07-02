package cacheobject

import (
	"context"
	"fmt"

	"github.com/citrix/adc-nitro-go/service"

	"github.com/citrix/terraform-provider-citrixadc/citrixadc_framework/utils"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*CacheobjectDataSource)(nil)

func CAcheobjectDataSource() datasource.DataSource {
	return &CacheobjectDataSource{}
}

type CacheobjectDataSource struct {
	client *service.NitroClient
}

func (d *CacheobjectDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cacheobject"
}

func (d *CacheobjectDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	d.client = *req.ProviderData.(**service.NitroClient)
}

func (d *CacheobjectDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = CacheobjectDataSourceSchema()
}

func (d *CacheobjectDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data CacheobjectResourceModel
	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// cacheobject has no get-by-name endpoint; only get(all). Fetch all cached
	// objects and filter locally on the supplied filter attributes.
	group_Name := data.Group
	groupname_Name := data.Groupname
	host_Name := data.Host
	httpmethod_Name := data.Httpmethod
	httpstatus_Name := data.Httpstatus
	ignoremarkerobjects_Name := data.Ignoremarkerobjects
	includenotreadyobjects_Name := data.Includenotreadyobjects
	locator_Name := data.Locator
	port_Name := data.Port
	url_Name := data.Url

	var dataArr []map[string]interface{}
	var err error

	findParams := service.FindParams{
		ResourceType:             service.Cacheobject.Type(),
		ResourceMissingErrorCode: 258,
	}
	dataArr, err = d.client.FindResourceArrayWithParams(findParams)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cacheobject, got error: %s", err))
		return
	}

	// Resource is missing
	if len(dataArr) == 0 {
		resp.Diagnostics.AddError("Client Error", "cacheobject returned empty array")
		return
	}

	// Iterate through results to find the one matching the supplied filters.
	foundIndex := -1
	for i, v := range dataArr {
		match := true

		// Check group
		if !group_Name.IsNull() {
			if val, ok := v["group"].(string); ok {
				if val != group_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check groupname
		if !groupname_Name.IsNull() {
			if val, ok := v["groupname"].(string); ok {
				if val != groupname_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check host
		if !host_Name.IsNull() {
			if val, ok := v["host"].(string); ok {
				if val != host_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check httpmethod
		if !httpmethod_Name.IsNull() {
			if val, ok := v["httpmethod"].(string); ok {
				if val != httpmethod_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check httpstatus
		if !httpstatus_Name.IsNull() {
			if val, ok := v["httpstatus"]; ok {
				intVal, _ := utils.ConvertToInt64(val)
				if intVal != httpstatus_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check ignoremarkerobjects
		if !ignoremarkerobjects_Name.IsNull() {
			if val, ok := v["ignoremarkerobjects"].(string); ok {
				if val != ignoremarkerobjects_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check includenotreadyobjects
		if !includenotreadyobjects_Name.IsNull() {
			if val, ok := v["includenotreadyobjects"].(string); ok {
				if val != includenotreadyobjects_Name.ValueString() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check locator
		if !locator_Name.IsNull() {
			if val, ok := v["locator"]; ok {
				intVal, _ := utils.ConvertToInt64(val)
				if intVal != locator_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check port
		if !port_Name.IsNull() {
			if val, ok := v["port"]; ok {
				intVal, _ := utils.ConvertToInt64(val)
				if intVal != port_Name.ValueInt64() {
					match = false
					continue
				}
			} else {
				match = false
				continue
			}
		}

		// Check url
		if !url_Name.IsNull() {
			if val, ok := v["url"].(string); ok {
				if val != url_Name.ValueString() {
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

	// Resource is missing
	if foundIndex == -1 {
		resp.Diagnostics.AddError("Client Error", "cacheobject not found matching the supplied filters")
		return
	}

	cacheobjectSetAttrFromGetForDatasource(ctx, &data, dataArr[foundIndex])

	// Datasource has no Create; compose the ID here. Prefer locator (unique int
	// ID); otherwise fall back to url.
	if !data.Locator.IsNull() {
		data.Id = types.StringValue(fmt.Sprintf("locator:%d", data.Locator.ValueInt64()))
	} else if !data.Url.IsNull() {
		data.Id = types.StringValue(fmt.Sprintf("url:%s", data.Url.ValueString()))
	} else {
		data.Id = types.StringValue("cacheobject")
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
