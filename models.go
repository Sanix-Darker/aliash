package main

type Alias struct {
	ID          string `bson:"ID"`
	uid         string `bson:"uid"`
	url         string `bson:"url"`
	author      string `bson:"author"`
	description string `bson:"description"`
}
