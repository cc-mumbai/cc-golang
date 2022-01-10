package main

import (
	"log"
	"time"

	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type todoModel struct {
	ID        bson.ObjectId `bson:"_id, omitempty"`
	Title     string        `bson:"title"`
	Completed bool          `bson:"completed"`
	CreatedAt time.Time     `bson:"createdAt"`
}

type todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Completed string    `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

func init() {
	rnd = renderer.New()
	sess, err := mgo.Dial(hostName)
	checkErr(err)
	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
