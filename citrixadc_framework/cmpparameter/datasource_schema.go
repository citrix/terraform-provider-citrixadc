package cmpparameter

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func CmpparameterDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"addvaryheader": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control insertion of the Vary header in HTTP responses compressed by Citrix ADC. Intermediate caches store different versions of the response for different values of the headers present in the Vary response header.",
			},
			"cmpbypasspct": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC CPU threshold after which compression is not performed. Range: 0 - 100",
			},
			"cmplevel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify a compression level. Available settings function as follows:\n * Optimal - Corresponds to a gzip GZIP level of 5-7.\n * Best speed - Corresponds to a gzip level of 1.\n * Best compression - Corresponds to a gzip level of 9.",
			},
			"cmponpush": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Citrix ADC does not wait for the quantum to be filled before starting to compress data. Upon receipt of a packet with a PUSH flag, the appliance immediately begins compression of the accumulated packets.",
			},
			"externalcache": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable insertion of  Cache-Control: private response directive to indicate response message is intended for a single user and must not be cached by a shared or proxy cache.",
			},
			"heurexpiry": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Heuristic basefile expiry.",
			},
			"heurexpiryhistwt": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "For heuristic basefile expiry, weightage to be given to historical delta compression ratio, specified as percentage.  For example, to give 25% weightage to historical ratio (and therefore 75% weightage to the ratio for current delta compression transaction), specify 25.",
			},
			"heurexpirythres": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Threshold compression ratio for heuristic basefile expiry, multiplied by 100. For example, to set the threshold ratio to 1.25, specify 125.",
			},
			"minressize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Smallest response size, in bytes, to be compressed.",
			},
			"policytype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the policy. The only possible value is ADVANCED",
			},
			"quantumsize": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum quantum of data to be filled before compression begins.",
			},
			"randomgzipfilename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Control the addition of a random filename of random length in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenamemaxlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"randomgzipfilenameminlength": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum length of the random filename to be added in the GZIP header to apply the Heal-the-BREACH mitigation for the BREACH attack.",
			},
			"servercmp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow the server to send compressed data to the Citrix ADC. With the default setting, the Citrix ADC appliance handles all compression.",
			},
			"varyheadervalue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The value of the HTTP Vary header for compressed responses. If this argument is not specified, a default value of \"Accept-Encoding\" will be used.",
			},
		},
	}
}
