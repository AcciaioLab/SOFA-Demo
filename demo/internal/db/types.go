package db

type Trailer struct {
	Trailer_ID           int
	VIN                  string
	IMEI                 int64
	Make                 string
	Model                string
	Active               bool
	Track_Width_T        float64
	Spring_Lat_Spacing   float64
	Roll_Centre_H        float64
	Kingpin_To_Axle_1    float64
	Axle_Spacing_1_2     float64
	Axle_Spacing_2_3     float64
	Unsprung_Mass_Axle_1 float64
	Unsprung_Mass_Axle_2 float64
	Unsprung_Mass_Axle_3 float64
	Sprung_Mass_Trailer  float64
	Total_Mass_Trailer   float64
	Suspension_Ratio     float64
	Sprung_CG_Height     float64
	CG_to_Axle_1         float64
	CG_to_Axle_2         float64
	CG_to_Axle_3         float64
}
