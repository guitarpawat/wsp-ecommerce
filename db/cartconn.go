package db

import (
	"errors"

	"github.com/globalsign/mgo/bson"
	"github.com/guitarpawat/wsp-ecommerce/model/dbmodel"
)

var NonCart = errors.New("Carts Not Exist")

var mockUserID1 = bson.ObjectIdHex("ba2946f27d9d403ce895633b")
var mockUserID2 = bson.ObjectIdHex("f8f0b5922a47fef34a30327b")

func MockCart() {
	db, err := GetDB()
	if err != nil {
		panic("cannot connect to db")
	}
	defer db.Session.Close()

	cart1 := dbmodel.InitialCart(mockUserID1)
	cart2 := dbmodel.InitialCart(mockUserID2)

	db.C("Carts").RemoveAll(nil)

	RegisCart(cart1)
	RegisCart(cart2)
}

func RegisCart(cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	// Check case

	_, err = db.C("Carts").Upsert(bson.M{"userID": cart.UserID}, cart)
	if err != nil {
		return err
	}
	return nil
}

func SetCart(id bson.ObjectId, cart dbmodel.Cart) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	err = db.C("Carts").Update(bson.M{"_id": id}, cart)
	if err != nil {
		return err
	}
	return nil
}

func GetCart(userName string) (dbmodel.Cart, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.Cart{}, err
	}
	defer db.Session.Close()

	user, err := GetUserFromName(userName)
	err = CheckCartExist(user.ID)
	if err != nil {
		return dbmodel.Cart{}, err
	}
	if err != nil {
		return dbmodel.Cart{}, errors.New("Unable to find user")
	}
	return GetCartID(user.ID)
}

func GetCartID(id bson.ObjectId) (dbmodel.Cart, error) {
	db, err := GetDB()
	if err != nil {
		return dbmodel.Cart{}, err
	}
	defer db.Session.Close()

	var cart dbmodel.Cart
	err = CheckCartExist(id)
	if err != nil {
		return dbmodel.Cart{}, err
	}
	err = db.C("Carts").Find(bson.M{"userID": id}).One(&cart)
	if err != nil {
		return dbmodel.Cart{}, NonCart
	}
	return cart, nil
}

func UpdateCart(userID, meat bson.ObjectId, quantity int) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	cartMeat := dbmodel.CartMeats{
		ID:       meat,
		Quantity: quantity,
	}

	err = CheckCartExist(userID)
	if err != nil {
		return err
	}

	err = db.C("Carts").Update(bson.M{
		"userID": userID,
	},
		bson.M{
			"$pull": bson.M{
				"meats": bson.M{
					"meat": meat,
				},
			},
		})
	if err != nil {
		return err
	}

	_, err = db.C("Carts").Upsert(
		bson.M{
			"userID": userID,
		},
		bson.M{
			"$push": bson.M{"meats": cartMeat},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func CheckCartExist(id bson.ObjectId) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer db.Session.Close()

	count, err := db.C("Carts").Find(bson.M{
		"userID": id,
	}).Count()
	if err != nil {
		return err
	}
	if count == 0 {
		cart := dbmodel.InitialCart(id)
		RegisCart(cart)
	}
	return nil
}
