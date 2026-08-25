package dnet

// ReadOnlyCtx is a wrapper/view for Ctx that provides read-only access to the exposed field values.
type ReadOnlyCtx struct {
	ctx *Ctx
}

// Get returns the value associated with the given key.
func (c *ReadOnlyCtx) Get(key string) (any, bool) {
	value, ok := c.ctx.values[key]

	return value, ok
}

// ID returns the ID of the user owning the connection.
func (c *ReadOnlyCtx) ID() string {
	return c.ctx.ID
}

// IsAuthed returns true if the connection is authenticated or not.
func (c *ReadOnlyCtx) IsAuthed() bool {
	return c.ctx.Authed
}

// Rec returns the id of the recipient.
func (c *ReadOnlyCtx) Rec() string {
	return c.ctx.Rec
}

// IP returns the IP address of the connection.
func (c *ReadOnlyCtx) IP() string {
	return c.ctx.IP
}

// NewReadOnlyCtx returns a new ReadOnlyCtx that wraps the given ctx.
func NewReadOnlyCtx(ctx *Ctx) *ReadOnlyCtx {
	return &ReadOnlyCtx{ctx: ctx}
}
