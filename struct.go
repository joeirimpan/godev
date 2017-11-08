package main

import "fmt"

type geopoint struct {
  latitude, longitude float64
}

func modify_geopoints(g *geopoint) *geopoint {
  g.latitude, g.longitude = g.longitude, g.latitude
  return g
}

func (g *geopoint) modify_geopoints() *geopoint {
  g.latitude, g.longitude = g.longitude, g.latitude
  return g
}


func main() {
  point1 := geopoint{latitude: 9.253936, longitude: 7.558594}
  fmt.Println("Lat: ", point1.latitude)
  fmt.Println("Long: ", point1.longitude)

  point2 := modify_geopoints(&point1)
  fmt.Println("Lat: ", point2.latitude)
  fmt.Println("Long: ", point2.longitude)

  point3 := point1.modify_geopoints()
  fmt.Println("Lat: ", point3.latitude)
  fmt.Println("Long: ", point3.longitude)
}
