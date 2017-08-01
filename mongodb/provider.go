package mongodb

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("MONGODB_URL", ""),
				Description: "The MongoDB url.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"mongodb_user": resourceMongoDBUser(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		URL: d.Get("url").(string),
	}

	client, err := config.loadAndValidate()

	if err != nil {
		return nil, err
	}

	return client, nil
}
