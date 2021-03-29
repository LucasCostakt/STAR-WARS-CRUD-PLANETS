package repository

import (
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoConnect struct {
	Client *mongo.Client
}

type NewPlanet struct {
	Id         string `json:"id"`
	Nome       string `json:"nome"`
	Clima      string `json:"clima"`
	Terreno    string `json:"terreno"`
	FilmsCount int64  `json:"films"`
}

type GetPlanets struct {
	ObjId      bson.ObjectId `bson:"_id" json:"id"`
	Nome       string        `bson:"nome"`
	Clima      string        `bson:"clima"`
	Terreno    string        `bson:"terreno"`
	FilmsCount int64         `bson:"films"`
}
type NewInsertPlanet struct {
	Nome       string `bson:"nome"`
	Clima      string `bson:"clima"`
	Terreno    string `bson:"terreno"`
	FilmsCount int64  `bson:"films"`
}

type Filters struct {
	Name string `json:"nome"`
	Id   string `json:"id"`
}

const MongoUri = "mongodb://localhost:27017"
const DataBaseName = "admin"
const CollectionName = "planets"
const UrlPlanets = "https://swapi.dev/api/planets/?search="
