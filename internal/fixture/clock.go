package fixture

type FixedClock struct{ Value string }

func (c FixedClock) Today() string { return c.Value }

func DefaultClock() FixedClock { return FixedClock{Value: SessionDate} }
