package domain

type Contact struct {
	Id   string `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name,omitempty"`
}

type ContactService interface {
	GetById(string) (*Contact, error)
	GetAll() ([]*Contact, error)
	Create(*Contact) (*Contact, error)
}
