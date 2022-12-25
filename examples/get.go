package examples

import "fmt"

type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	DateAdded string `json:"dateAdded"`
	CompanyID string `json:"companyId"`
}

type ContactsRes struct {
	Data []Contact `json:"contacts"`
}
type ContactRes struct {
	Data Contact `json:"contact"`
}

func GetContacts(url string) ([]Contact, error) {
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var contacts ContactsRes
	err = res.Unmarshal(&contacts)
	if err != nil {
		return nil, err
	}
	return contacts.Data, nil
}

func GetContactByID(url string, id int) (*Contact, error) {
	url = fmt.Sprintf("%s/%d", url, id)
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var contact ContactRes
	err = res.Unmarshal(&contact)
	if err != nil {
		return nil, err
	}
	return &contact.Data, nil
}
