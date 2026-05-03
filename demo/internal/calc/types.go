package calc

import "time"

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

type TrailerInput struct {
	Timestamp                 time.Time // Timestamp (ts)
	VIN                       string    //
	ABPresLHFront             float64   // Calculated, from ABPresLHRear
	ABPresRHFront             float64   // Calculated, from ABPresLHFront
	ABPresLHMID               float64   // Calculated, from ABPresRHFront
	ABPresRHMID               float64   // Calculated, from ABPresRHRear
	ABPresLHRear              float64   // Calculated, from ABPresLHFront and ABPresLHRear
	ABPresRHRear              float64   // Calculated, from ABPresRHFront and ABPresRHRear
	AxleLoadLHRear            float64   //
	AxleLoadLHFront           float64   //
	AxleLoadRHFront           float64   //
	AxleLoadRHRear            float64   //
	AxleLoadLHCentre          float64   //
	AxleLoadRHCentre          float64   //
	DLAxle1LAZ                float64   //
	DLAxle1RAZ                float64   //
	DLAxle2LAZ                float64   //
	DLAxle2RAZ                float64   //
	SpeedSat                  float64   //
	Speed                     float64   //
	DLKingpinGZ               float64   //
	DLKingpinAY               float64   //
	DLKingpinAZ               float64   //
	BrakeApplicationPressureH float64   //
}

type TrailerLoad struct {
	Timestamp        time.Time // Timestamp (ts)
	VIN              string    //
	StatusLoaded     string    //
	FAirAxlex1LN     float64   //
	FAirAxlex1RN     float64   //
	FAirAxlex2LN     float64   //
	FAirAxlex2RN     float64   //
	FAirAxlex3LN     float64   //
	FAirAxlex3RN     float64   //
	KAx1LNPerm       float64   //
	KAx1RNPerm       float64   //
	KAx2LNPerm       float64   //
	KAx2RNPerm       float64   //
	KAx3LNPerm       float64   //
	KAx3RNPerm       float64   //
	FInertiaAx1LN    float64   //
	FInertiaAx1RN    float64   //
	FInertiaAx2LN    float64   //
	FInertiaAx2RN    float64   //
	FInertiaAx3LN    float64   //
	FInertiaAx3RN    float64   //
	XAx1Lm           float64   //
	XAx1Rm           float64   //
	XAx2Lm           float64   //
	XAx2Rm           float64   //
	XAx3Lm           float64   //
	XAx3Rm           float64   //
	DeltaT           float64   //
	XDotAx1LMps      float64   //
	XDotAx1RMps      float64   //
	XDotAx2LMps      float64   //
	XDotAx2RMps      float64   //
	XDotAx3LMps      float64   //
	XDotAx3RMps      float64   //
	XDotFAx1LMps     float64   //
	XDotFAx1RMps     float64   //
	XDotFAx2LMps     float64   //
	XDotFAx2RMps     float64   //
	XDotFAx3LMps     float64   //
	XDotFAx3RMps     float64   //
	FdAx1LN          float64   //
	FdAx1RN          float64   //
	FdAx2LN          float64   //
	FdAx2RN          float64   //
	FdAx3LN          float64   //
	FdAx3RN          float64   //
	FzAx1LN          float64   //
	FzAx1RN          float64   //
	FzAx2LN          float64   //
	FzAx2RN          float64   //
	FzAx3LN          float64   //
	FzAx3RN          float64   //
	FzTotalN         float64   //
	FzLTotalN        float64   //
	FzRTotalN        float64   //
	FzAx1TotalN      float64   //
	FzAx2TotalN      float64   //
	FzAx3TotalN      float64   //
	FzAx1AbsLN       float64   //
	FzAx1AbsRN       float64   //
	FzAx2AbsLN       float64   //
	FzAx2AbsRN       float64   //
	FzAx3AbsLN       float64   //
	FzAx3AbsRN       float64   //
	LTRGroup         float64   //
	DeltaFzLatAx1    float64   //
	DeltaFzLatAx2    float64   //
	DeltaFzLatAx3    float64   //
	UMps             float64   //
	RKpRps           float64   //
	AyG              float64   //
	MRollNm          float64   //
	BPAxBar          float64   //
	BPAx1LBar        float64   //
	BPAx1RBar        float64   //
	BPAx2LBar        float64   //
	BPAx2RBar        float64   //
	BPAx3LBar        float64   //
	BPAx3RBar        float64   //
	TWheelAx1L       float64   //
	TWheelAx1R       float64   //
	TWheelAx2L       float64   //
	TWheelAx2R       float64   //
	TWheelAx3L       float64   //
	TWheelAx3R       float64   //
	FxBrakeAx1L      float64   //
	FxBrakeAx1R      float64   //
	FxBrakeAx2L      float64   //
	FxBrakeAx2R      float64   //
	FxBrakeAx3L      float64   //
	FxBrakeAx3R      float64   //
	FxBrakeAx1Total  float64   //
	FxBrakeAx2Total  float64   //
	FxBrakeAx3Total  float64   //
	FxBrakeTotal     float64   //
	BetaBodyRad      float64   //
	AlphaAx1LRad     float64   //
	AlphaAx1RRad     float64   //
	AlphaAx2LRad     float64   //
	AlphaAx2RRad     float64   //
	AlphaAx3LRad     float64   //
	AlphaAx3RRad     float64   //
	FyAx1LN          float64   //
	FyAx1RN          float64   //
	FyAx2LN          float64   //
	FyAx2RN          float64   //
	FyAx3LN          float64   //
	FyAx3RN          float64   //
	FyTotalN         float64   //
	MzYawNm          float64   //
	MassEffective    float64   //
	FxKingpinN       float64   //
	FyKingpinN       float64   //
	FzKingpinN       float64   //
	MassEstRearkg    float64   //
	RollTransferGain float64   //
	DeflectBiasAx1   float64   //
	DeflectBiasAx2   float64   //
	DeflectBiasAx3   float64   //
}
