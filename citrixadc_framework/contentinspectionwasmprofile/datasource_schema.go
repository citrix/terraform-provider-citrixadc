package contentinspectionwasmprofile

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func ContentinspectionwasmprofileDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"anomalousdatasize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Transaction data size (in KB) greater than which a transaction is considered as anomalous. Default is 512KB.",
			},
			"anomalousttfbtime": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Transaction time (in milliseconds) above which a transaction is considered as anomalous. Default is 1 seconds.",
			},
			"maxbodylen": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Max data size (in KB) that will be sent to the CI Agent. Default is 16KB. Maximum value that can be configured is 32KB.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of CI WASM profile.",
			},
			"timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout (in milliseconds) for the connection with the CI WASM agent.",
			},
			"timeoutaction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout action for the connection with the CI agent. Either the original request can be bypassed i.e. request/response is forwarded to the endpoint or the transaction is dropped/reset. Possible values = BYPASS, DROP, RESET",
			},
			"wasmmodule": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the WASM Module.",
			},
		},
	}
}
