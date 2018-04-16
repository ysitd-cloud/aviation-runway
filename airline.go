package tower

type Airline interface {
	Flyer
	// Marker function
	Airline()
}
