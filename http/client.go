package http

type client struct{}

type Client interface {
	Get()
	Post()
	Put()
	Patch()
	Delete()
	Head()
}

func NewClient() Client {
	return &client{}
}

func (c *client) Get()    {}
func (c *client) Post()   {}
func (c *client) Put()    {}
func (c *client) Delete() {}
func (c *client) Patch()  {}
func (c *client) Head()   {}
