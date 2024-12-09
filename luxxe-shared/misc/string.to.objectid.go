package misc

import "go.mongodb.org/mongo-driver/bson/primitive"

func StringToObjectID(idStr string) primitive.ObjectID {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return primitive.NilObjectID
	}
	return id
}
