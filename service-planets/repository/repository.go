package repository

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	InsertNewPlanet(newInsertPlanet NewInsertPlanet) ([]byte, error)
	CounsultPlanetByName(name string) ([]byte, error)
	CounsultPlanetByID(name string) ([]byte, error)
	CounsultAllPlanets() ([]byte, error)
	DeletePlanetById(id string) error
}

func NewRepository(client *mongo.Client) (Repository, error) {
	repository := MongoConnect{Client: client}
	return &repository, nil
}

func NewMongoConnect() (Repository, error) {
	clientStruct := &MongoConnect{}
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUri))
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Erro ao criar o novo client mongo")
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Erro ao conectar ao conectar ao mongo")
	}

	quickstartDatabase := client.Database(DataBaseName)
	quickstartDatabase.CreateCollection(ctx, CollectionName)

	clientStruct.Client = client
	log.Println("Database Conectado")
	res, err := NewRepository(clientStruct.Client)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Erro ao conectar ao conectar ao mongo")
	}
	return res, nil
}

func ConvertBson(getPlanets GetPlanets) ([]byte, error) {
	result := NewPlanet{}

	ws, err := hex.DecodeString(getPlanets.ObjId.Hex())
	if err != nil {
		log.Println("Erro hex.DecodeString() CounsultAllPlanets:", err)
		return nil, fmt.Errorf("Erro hex.DecodeString() CounsultAllPlanets")
	}

	result.Id = string(ws)
	result.Clima = getPlanets.Clima
	result.FilmsCount = getPlanets.FilmsCount
	result.Nome = getPlanets.Nome
	result.Terreno = getPlanets.Terreno

	convertBsonToJson, err := json.Marshal(result)
	if err != nil {
		log.Println("Erro json.Marshal() CounsultAllPlanets : ", err)
		return nil, fmt.Errorf("Erro json.Marshal() CounsultAllPlanets")
	}

	return convertBsonToJson, nil
}

func (c *MongoConnect) InsertNewPlanet(newInsertPlanet NewInsertPlanet) ([]byte, error) {

	_, err := c.Client.Database(DataBaseName).Collection(CollectionName).InsertOne(context.TODO(), &newInsertPlanet)
	if err != nil {
		log.Println("Collection.InsertOne() erro no InsertNewPlanet: ", err)
		return nil, fmt.Errorf("Erro InsertOne() InsertNewPlanet")
	}

	return []byte("Sucesso ao Inserir um Novo Planete"), nil
}

func (c *MongoConnect) CounsultPlanetByName(name string) ([]byte, error) {
	primitiveBson := GetPlanets{}

	result := c.Client.Database(DataBaseName).Collection(CollectionName).FindOne(context.TODO(), bson.D{{"nome", name}})
	err := result.Decode(&primitiveBson)
	if err != nil {
		log.Println("Erro Decode() CounsultPlanetByName: ", err)
		return []byte("Nenhum resultado encontrado"), nil
	}
	planet, err := ConvertBson(primitiveBson)
	if err != nil {
		log.Println("Erro ConvertBson() CounsultPlanetByName: ", err)
		return nil, fmt.Errorf("Erro ConvertBson() CounsultPlanetByName")
	}

	return planet, nil
}
func (c *MongoConnect) CounsultPlanetByID(id string) ([]byte, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Erro ObjectIDFromHex() ConsultPlanetByID: ", err)
		return nil, fmt.Errorf("Erro ObjectIDFromHex() ConsultPlanetByID")
	}

	primitiveBson := GetPlanets{}
	result := c.Client.Database(DataBaseName).Collection(CollectionName).FindOne(context.TODO(), bson.M{"_id": objID})

	err = result.Decode(&primitiveBson)
	if err != nil {
		log.Println("Erro Decode() ConsultPlanetByID: ", err)
		return []byte("Nenhum resultado encontrado"), nil
	}

	planet, err := ConvertBson(primitiveBson)
	if err != nil {
		log.Println("Erro ConvertBson() ConsultPlanetByID: ", err)
		return nil, fmt.Errorf("Erro ConvertBson() ConsultPlanetByID")
	}

	return planet, nil
}
func (c *MongoConnect) CounsultAllPlanets() ([]byte, error) {
	ctx := context.TODO()

	results := []NewPlanet{}
	cursor, err := c.Client.Database(DataBaseName).Collection(CollectionName).Find(ctx, bson.D{})
	if err != nil {
		log.Println("Erro Cursor CounsultAllPlanets(): ", err)
		return nil, fmt.Errorf("Erro no cursor CounsultAllPlanets")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		result := NewPlanet{}
		getPlanets := GetPlanets{}
		err := cursor.Decode(&getPlanets)
		if err != nil {
			log.Println("Erro cursor.Next() CounsultAllPlanets: ", err)
			return []byte("Nenhum resultado encontrado"), nil
		}

		ws, err := hex.DecodeString(getPlanets.ObjId.Hex())
		if err != nil {
			log.Println("Error hex.DecodeString() CounsultAllPlanets: ", err)
			return nil, fmt.Errorf("Erro hex.DecodeString() CounsultAllPlanets")
		}

		result.Id = string(ws)
		result.Clima = getPlanets.Clima
		result.FilmsCount = getPlanets.FilmsCount
		result.Nome = getPlanets.Nome
		result.Terreno = getPlanets.Terreno

		results = append(results, result)
	}

	js, err := json.Marshal(results)
	if err != nil {
		log.Println("Erro json.Marshal() CounsultAllPlanets: ", err)
		return nil, fmt.Errorf("Erro json.Marshal() CounsultAllPlanets")
	}

	return js, nil
}

func (c *MongoConnect) DeletePlanetById(id string) error {
	ctx := context.TODO()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Erro ObjectIDFromHex() DeletePlanetById: ", err)
		return fmt.Errorf("Erro ObjectIDFromHex() DeletePlanetById")
	}

	_, err = c.Client.Database(DataBaseName).Collection(CollectionName).DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Println("Erro DeleteOne() DeletePlanetById: ", err)
		return fmt.Errorf("Erro DeleteOne() DeletePlanetById")
	}

	return nil
}
