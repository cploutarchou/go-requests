package examples

type Contact struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	DateAdded string `json:"dateAdded"`
	CompanyID string `json:"companyId"`
}

type Contacts struct {
	Data []Contact `json:"contacts"`
}

func GetContacts(url string) ([]Contact, error) {
	res, err := client.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var contacts Contacts
	err = res.Unmarshal(&contacts)
	if err != nil {
		return nil, err
	}
	return contacts.Data, nil
}
