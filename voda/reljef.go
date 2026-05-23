package voda

func postaviReku() {
	x
	y
	for _,i :=range WorldReke.potoci{
		for _,j := range i.deo{
			if (j.start.x<=x*16+15 && j.end.x>=x*16) || (j.start.x*16>=x&&j.end.x<=x*16+15){
				if (j.start.z<=z*16+15&&j.end.z>=z*16) || (j.start.z<=z*16&&j.end.z>=z*16+15){
					TrazenoX1 := j.start.x
					TrazenoY1 := j.start.y
					if j.start.x < x*16 {
						TrazenoX1 = x*16
					} else if j.start.x > x*16+15{
						TrazenoX1 = x*16+15
					}
					if j.start.y < y*16 {
						TrazenoY1 = y*16
					} else if j.start.y > y*16+15{
						TrazenoY1 = y*16+15
					}
					TrazenoX2 := j.end.x
					TrazenoY2 := j.end.y
					if j.end.y < y*16 {
						TrazenoY2 = y*16
					} else j.end.y > y*16+15{
						TrazenoY2 = y*16+15
					}
					if j.end.x < x*16 {
						TrazenoX2 = x*16
					} else if j.end.y > x*16+15 {
						TrazenoX2 = x*16 +15
					}
				}
			}
		}
	}
}