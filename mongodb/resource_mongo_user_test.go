package mongodb

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

var testAccMongoDBUserConfig = fmt.Sprintf(`
resource "mongodb_user" "user" {
	database = "testing"
    username = "user"
    password = "pass"
    roles = ["read", "dbAdmin", "userAdmin"]
}
`)

func TestAccMongoDBUser_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccMongoDBUserDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccMongoDBUserConfig,
				Check: resource.ComposeTestCheckFunc(
					testCheckMongoDBUserExists("mongodb_user.user", t),
				),
			},
		},
	})
}

func testCheckMongoDBUserExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*mgo.Session)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Username is set")
		}

		database := rs.Primary.Attributes["database"]
		username := rs.Primary.Attributes["username"]
		password := rs.Primary.Attributes["password"]

		if database == "" {
			return fmt.Errorf("No Database is set")
		}

		if username == "" {
			return fmt.Errorf("No Username is set")
		}

		if password == "" {
			return fmt.Errorf("No Password is set")
		}

		db := client.DB(database)
		err := db.Login(username, password)
		if err != nil {
			return err
		}
		db.Logout()

		return nil
	}
}

func testAccMongoDBUserDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*mgo.Session)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "mongodb_user" {
			continue
		}

		username := rs.Primary.Attributes["username"]
		database := rs.Primary.Attributes["database"]
		password := rs.Primary.Attributes["password"]

		db := client.DB(database)
		err := db.Login(username, password)
		if err != nil {
			return nil
		}

		return nil
	}

	return nil
}
