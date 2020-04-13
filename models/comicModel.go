package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ComicModel = struct of comic model
type ComicModel struct {
	ID    primitive.ObjectID `json:"ID" bson:"_id"`
	Alias string             `json:"Alias"`
	Name  string             `json:"Name"`
	Title string             `json:"Title"`
	Files []ComicFileModel   `json:"Files"`
}

// ComicFileModel = struct of comic file model
type ComicFileModel struct {
	ID       primitive.ObjectID `json:"ID" form:"ID" query:"ID" bson:"_id"`
	Image    string             `json:"Image" form:"Image" query:"Image" bson:"Image"`
	ImageAlt string             `json:"ImageAlt" form:"ImageAlt" query:"ImageAlt" bson:"ImageAlt"`
}

// {
// 	ID: lwkwwkwkw,
// 	Alias: DC,
// 	Name: Detective Conan,
// 	Title: File 4 - Cerobong Asap Ke Enam
// }

// NewComicModel = create new comic model
func NewComicModel() *ComicModel {
	c := new(ComicModel)
	c.ID = primitive.NewObjectID()

	return c
}

// NewComicFileModel = create new comic file model
func NewComicFileModel() *ComicFileModel {
	cf := new(ComicFileModel)
	cf.ID = primitive.NewObjectID()

	return cf
}

// TableName = table name for comic
func (c *ComicModel) TableName() string {
	return "comic"
}
