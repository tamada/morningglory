package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"cloud.google.com/go/datastore"
)

type datastoreDB struct {
	client *datastore.Client
}

const ProjectID = "morningglory"

var db *datastoreDB

func InitDatastore() error {
	var ctx = context.Background()
	client, err := datastore.NewClient(ctx, ProjectID)
	if err != nil {
		return err
	}
	t, err := client.NewTransaction(ctx)
	if err != nil {
		return fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	if err := t.Rollback(); err != nil {
		return fmt.Errorf("datastoredb: could not connect: %v", err)
	}
	db = &datastoreDB{client: client}
	return nil
}

// Close closes the database.
func (db *datastoreDB) Close() {
	// No op.
}

type Point struct {
	User       string    `json:"-"`
	Repository string    `json:"repository"`
	Action     string    `json:"action"`
	Point      int       `json:"point"`
	RefURL     string    `json:"ref_url"`
	Date       time.Time `json:"-"`
}

type User struct {
	Name      string
	KeyPhrase string
}

/*
Md5sum converts given str to md5 hex string.
*/
func Md5sum(str string) string {
	var hash = md5.New()
	var md5 = hash.Sum([]byte(str))
	return hex.EncodeToString(md5)
}

func RegisterPoint(point *Point) error {
	var ctx = context.Background()
	var key = datastore.IncompleteKey("Points", nil)
	var _, err = db.client.Put(ctx, key, point)
	return err
}

func RegisterUser(userName, md5KeyPhrase string) error {
	var ctx = context.Background()
	var _, err = findUser(ctx, userName)
	if err != nil {
		return err
	}
	var key = datastore.NameKey("Users", userName, nil)
	var _, err2 = db.client.Put(ctx, key, &User{
		KeyPhrase: md5KeyPhrase,
	})
	return err2
}

func DeleteUser(userName string) error {
	var ctx = context.Background()
	var key = datastore.NameKey("Users", userName, nil)
	return db.client.Delete(ctx, key)
}

func UpdateKeyPhrase(userName, md5KeyPhrase string) error {
	var ctx = context.Background()
	var _, err = findUser(ctx, userName)
	if err != nil {
		return err
	}
	var key = datastore.NameKey("Users", userName, nil)
	var _, err2 = db.client.Put(ctx, key, &User{
		Name:      userName,
		KeyPhrase: md5KeyPhrase,
	})
	return err2
}

/*
Authenticate finds userName and match md5KeyPhrase.
*/
func Authenticate(userName, md5KeyPhrase string) error {
	var ctx = context.Background()
	var found, err = findUser(ctx, userName)
	if err != nil {
		return fmt.Errorf("%s: user not found", userName)
	}
	if found.KeyPhrase != md5KeyPhrase {
		return fmt.Errorf("%s: authenticate failed", userName)
	}
	return nil
}

func FindUser(userName string) *User {
	var ctx = context.Background()
	var user, _ = findUser(ctx, userName)
	return user
}

func findUser(ctx context.Context, userName string) (*User, error) {
	var key = datastore.NameKey("Users", userName, nil)
	var found = User{}
	var err = db.client.Get(ctx, key, &found)
	if err != nil {
		return nil, err
	}
	return &found, nil
}
