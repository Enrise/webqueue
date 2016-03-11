package webqueue

import (
	"gopkg.in/mgo.v2"
	"time"
)

var AppMongoConfig MongoConfig

type MessageResultOut struct {
	Request   string
	Response  string
	Status    int
	Id        interface{} "_id"
	Timestamp time.Time
}

type DateMessageCount struct {
	Timestamp time.Time
	count     int
}

func GetLatestMessages() (results []MessageResultOut) {
	session, err := getMongoSession()
	if err != nil {
		return results
	}
	collection := session.DB(AppMongoConfig.Database).C("messagelog")

	err = collection.Find(nil).Sort("-_id").Limit(10).All(&results)
	if err != nil {
		Log.Fatal("Invalid query for fetching messages: %s", err)
		return results
	}
	return results
}

func GetProcessedMessagesHistogram() (results []MessageResultOut) {
	session, err := getMongoSession()
	if err != nil {
		return results
	}
	collection := session.DB(AppMongoConfig.Database).C("map_reduce_example")

	err = collection.Find(nil).Sort("-_id").All(&results)
	if err != nil {
		Log.Fatal("Invalid query for fetching messages: %s", err)
		return results
	}
	return results
}

func getMongoSession() (*mgo.Session, error) {
	timeout := time.Duration(AppMongoConfig.Timeout) * time.Second
	session, err := mgo.DialWithTimeout(AppMongoConfig.Host, time.Duration(timeout))
	if err != nil {
		Log.Fatal("Could not get messages from MongoDB: %s", err)
		return nil, err
	}
	return session, nil
}
