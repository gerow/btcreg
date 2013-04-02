package btcreg

type Address struct {
  Email string
  Address string
}

type AddRequest struct {
  UUID string
  Email string
}

type DeleteRequest struct {
  UUID string
  Email string
}

//func LoadAddressByEmail(email string) (Address, error) {
//}

//func InsertAddress(address Address) (error) {
//}

//func LoadAddRequestByUUID(uuid string) (AddRequest, error) {
//}

//func LoadDeleteRequestByUUID(uuid string) (DeleteRequest, error) {
//}
