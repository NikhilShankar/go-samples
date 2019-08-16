package codeg

//go:generate stringer -type=Jobs
type Jobs int

const  (

	Director Jobs = iota
	Producer
	Cinematographer
	Choreographer
	DanceMaster = Choreographer

)



