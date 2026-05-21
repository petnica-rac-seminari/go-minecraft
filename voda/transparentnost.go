package voda

func transparentnost() {
	if Water != nil {
		
		switch {
		case jezero:
			Water.(0.3)
		case reka:
			Water.(0.1)
		case more:
			Water.(0.5)
		case bara:
			Water.(0.7)
		}
	}
}
