package videooptimizationdetectionaction

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

func VideooptimizationdetectionactionDataSourceSchema() schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comment. Any type of information about this video optimization detection action.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name for the video optimization detection action. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"newname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "New name for the videooptimization detection action.\nMust begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of video optimization action. Available settings function as follows:\n* clear_text_pd - Cleartext PD type is detected.\n* clear_text_abr - Cleartext ABR is detected.\n* encrypted_abr - Encrypted ABR is detected.\n* trigger_enc_abr - Possible encrypted ABR is detected.\n* trigger_body_detection - Possible cleartext ABR is detected. Triggers body content detection.",
			},
		},
	}
}
