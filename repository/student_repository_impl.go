package repository

import (
	"go-clean/config"
	"go-clean/entity"
	"go-clean/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewStudentRepository(database *mongo.Database) StudentRepository {
	return &studentRepositoryImpl{
		Collection: database.Collection("students"),
	}
}

type studentRepositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *studentRepositoryImpl) UpdateById(id string, student entity.Student) bool {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{
			"name": student.Name,
			"email": student.Email,
		},
	}

	result,err := repository.Collection.UpdateOne(ctx, filter, update)
	exception.PanicIfNeeded(err)
	if result.ModifiedCount == 1 {
		return true
	}
	return false
}

func (repository *studentRepositoryImpl) GetById(id string) entity.Student {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var student entity.Student
	err := repository.Collection.FindOne(ctx, bson.M{
		"_id": id,
	}).Decode(&student)

	exception.PanicIfNeeded(err)
	return student
}

type countDeleted struct {
	DeletedCount int64 `bson:"n"`
}

func (repository *studentRepositoryImpl) Delete(id string) string {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	result, _ := repository.Collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	if result.DeletedCount == 1 {
		return "DELETED"
	}
	return "ID NOT FOUND"
}

func (repository *studentRepositoryImpl) Insert(student entity.Student) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		"_id":      	student.Id,
		"identifier":   student.Identifier,
		"name":    		student.Name,
		"email": 		student.Email,
	})
	exception.PanicIfNeeded(err)
}

func (repository *studentRepositoryImpl) FindAll() (students []entity.Student) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)
	exception.PanicIfNeeded(err)

	for _, document := range documents {
		students = append(students, entity.Student{
			Id:       	document["_id"].(string),
			Identifier: document["identifier"].(string),
			Name:    	document["name"].(string),
			Email: 		document["email"].(string),
		})
	}

	return students
}

func (repository *studentRepositoryImpl) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.DeleteMany(ctx, bson.M{})
	exception.PanicIfNeeded(err)
}