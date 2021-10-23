package schemas

var Products []*Product
var Users []*User

func AddRequiredData(){
	adduser()
	addProduct()
}
func adduser()  {
	user1 := &User{
		Name:    "Arnob kumar saha",
		Id:      10,
		Contact: Contact{
			PhoneNumber: 123123123,
			Address:     "uttara",
		},
	}
	user2 := &User{
		Name:    "Tasdidur rahman",
		Id:      15,
		Contact: Contact{
			PhoneNumber: 4556132465,
			Address:     "banani",
		},
	}
	user3 := &User{
		Name:    "Rakibul hossain",
		Id:      12,
		Contact: Contact{
			PhoneNumber: 54132133,
			Address:     "dhanmondi",
		},
	}
	Users = append(Users, user1, user2, user3)
}
func addProduct()  {
	pr1 := &Product{
		Title:   "samsung j7",
		Price:   123,
		Type:    "phone",
		Id:      2,
		OwnerId: 12,
	}
	pr2 := &Product{
		Title:   "asus 3453",
		Price:   6200,
		Type:    "laptop",
		Id:      3,
		OwnerId: 10,
	}
	pr3 := &Product{
		Title:   "redme note4",
		Price:   10023,
		Type:    "phone",
		Id:      6,
		OwnerId: 15,
	}
	Products = append(Products, pr1, pr2, pr3)
}