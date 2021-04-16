package window

// Kind is
type TKind string

const (
	AlwaysTrigger  TKind = "ALW"
	NeverTrigger   TKind = "NEV"
	DefaultTrigger TKind = "DEF"
	ElementTrigger TKind = "ELE"
)

// Trigger struct
type Trigger struct {
	Kind TKind
	//countElements int
	Count int
}

// type AfterPane struct {
//  countElements int
// }
// NewTrigger - Main func
func NewNeverTrigger() *Trigger {
	return &Trigger{Kind: NeverTrigger}
}

func NewDefaultTrigger() *Trigger {
	return &Trigger{Kind: DefaultTrigger}
}

func NewAlwaysTrigger() *Trigger {
	return &Trigger{Kind: AlwaysTrigger}
}

func NewElementCountTrigger(count int) *Trigger {
	return &Trigger{Kind: ElementTrigger, Count: count}
}
