package mongodb

import (
	"os"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var testAccMongoDBClient *mgo.Session

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"mongodb": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	v := os.Getenv("MONGODB_URL")
	if v == "" {
		t.Fatal("MONGODB_URL must be set for acceptance tests")
	}

	if testAccMongoDBClient == nil {
		config := Config{
			URL: os.Getenv("MONGO_URL"),
		}

		if client, err := config.loadAndValidate(); err != nil {
			t.Fatalf("could not load MongoDB Client: %s", err)
		} else {
			testAccMongoDBClient = client
		}
	}
}
