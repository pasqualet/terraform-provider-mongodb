package mongodb

import (
	"bytes"

	"github.com/hashicorp/terraform/helper/schema"
	"gopkg.in/mgo.v2"
)

func resourceMongoDBUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceMongoDBUserCreate,
		Update: resourceMongoDBUserUpdate,
		Read:   resourceMongoDBUserRead,
		Exists: resourceMongoDBUserExists,
		Delete: resourceMongoDBUserDelete,
		Schema: map[string]*schema.Schema{
			"database": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"roles": &schema.Schema{
				Type:     schema.TypeSet,
				Required: false,
				ForceNew: false,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
		},
	}
}

func readMongoDBUser(d *schema.ResourceData, meta interface{}) error {
	dbname := d.Get("database").(string)
	username := d.Get("username").(string)

	var id bytes.Buffer
	id.WriteString(dbname)
	id.WriteString(".")
	id.WriteString(username)

	d.SetId(id.String())

	return nil
}

func resourceMongoDBUserRead(d *schema.ResourceData, meta interface{}) error {
	database := d.Get("database").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var id bytes.Buffer
	id.WriteString(database)
	id.WriteString(".")
	id.WriteString(username)

	d.SetId(id.String())
	d.Set("username", username)
	d.Set("password", password)

	return nil
}

func resourceMongoDBUserCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceMongoDBUserUpdate(d, meta)
}

func resourceMongoDBUserUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*mgo.Session)

	dbname := d.Get("database").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	roles := d.Get("roles").(*schema.Set)
	mongodb_roles := getMongoDBUserRoles(roles)

	user := mgo.User{
		Username: username,
		Password: password,
		Roles:    mongodb_roles,
	}

	db := client.DB(dbname)
	err := db.UpsertUser(&user)
	if err != nil {
		return err
	}

	return readMongoDBUser(d, meta)
}

func resourceMongoDBUserDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*mgo.Session)

	database := d.Get("database").(string)
	username := d.Get("username").(string)

	db := client.DB(database)
	err := db.RemoveUser(username)

	if err != nil && err.Error() != "not found" {
		return err
	}

	return nil
}

func resourceMongoDBUserExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	client := meta.(*mgo.Session)

	database := d.Get("database").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	db := client.DB(database)
	err := db.Login(username, password)
	db.Logout()

	return err == nil, nil
}

func getMongoDBUserRoles(roles *schema.Set) []mgo.Role {
	mongodb_roles := []mgo.Role{}

	for _, role := range roles.List() {
		mrole := mgo.Role(role.(string))
		mongodb_roles = append(mongodb_roles, mrole)
	}
	return mongodb_roles
}
