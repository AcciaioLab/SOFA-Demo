package calc

import (
	"math"
	"strconv"
	"time"
)

func (ti *TrailerInput) UnmarshalRecord(r []string) {
	timeLayout := "2006-01-02 15:04:05.000"
	ti.Timestamp, _ = time.Parse(timeLayout, r[0])
	ti.VIN = r[93]
	ti.ABPresLHFront, _ = strconv.ParseFloat(r[99], 64)
	ti.ABPresRHFront, _ = strconv.ParseFloat(r[100], 64)
	ti.ABPresLHMID, _ = strconv.ParseFloat(r[98], 64)
	ti.ABPresRHMID, _ = strconv.ParseFloat(r[101], 64)
	ti.ABPresLHRear, _ = strconv.ParseFloat(r[97], 64)
	ti.ABPresRHRear, _ = strconv.ParseFloat(r[102], 64)
	ti.AxleLoadLHRear, _ = strconv.ParseFloat(r[103], 64)
	ti.AxleLoadLHFront, _ = strconv.ParseFloat(r[104], 64)
	ti.AxleLoadRHFront, _ = strconv.ParseFloat(r[105], 64)
	ti.AxleLoadRHRear, _ = strconv.ParseFloat(r[106], 64)
	ti.AxleLoadLHCentre, _ = strconv.ParseFloat(r[107], 64)
	ti.AxleLoadRHCentre, _ = strconv.ParseFloat(r[108], 64)
	ti.DLAxle1LAZ, _ = strconv.ParseFloat(r[15], 64)
	ti.DLAxle1RAZ, _ = strconv.ParseFloat(r[24], 64)
	ti.DLAxle2LAZ, _ = strconv.ParseFloat(r[33], 64)
	ti.DLAxle2RAZ, _ = strconv.ParseFloat(r[42], 64)
	ti.SpeedSat, _ = strconv.ParseFloat(r[7], 64)
	ti.Speed, _ = strconv.ParseFloat(r[6], 64)
	ti.DLKingpinGZ, _ = strconv.ParseFloat(r[54], 64)
	ti.DLKingpinAY, _ = strconv.ParseFloat(r[50], 64)
	ti.DLKingpinAZ, _ = strconv.ParseFloat(r[51], 64)
	ti.BrakeApplicationPressureH, _ = strconv.ParseFloat(r[91], 64)
}

// Calculate trailer loads
func (ti *TrailerInput) CalcLoadsFromInput(tl *TrailerLoad, last *TrailerLoad, t *Trailer) error {
	tl.Timestamp = ti.Timestamp
	tl.VIN = ti.VIN

	tl.DeltaT = float64(tl.Timestamp.UnixMilli()-last.Timestamp.UnixMilli()) / 1000.0

	tl.StatusLoaded = calcTrailerLoaded((ti.ABPresLHMID + ti.ABPresRHMID) / 2)

	// Air spring forces
	tl.FAirAxlex1LN = calcAirSpringForce(ti.ABPresLHFront, t.Suspension_Ratio) //
	tl.FAirAxlex1RN = calcAirSpringForce(ti.ABPresRHFront, t.Suspension_Ratio) //
	tl.FAirAxlex2LN = calcAirSpringForce(ti.ABPresLHMID, t.Suspension_Ratio)   //
	tl.FAirAxlex2RN = calcAirSpringForce(ti.ABPresRHMID, t.Suspension_Ratio)   //
	tl.FAirAxlex3LN = calcAirSpringForce(ti.ABPresLHRear, t.Suspension_Ratio)  //
	tl.FAirAxlex3RN = calcAirSpringForce(ti.ABPresRHRear, t.Suspension_Ratio)  //

	// Spring stiffness
	tl.KAx1LNPerm = lookupAirSpringCurve(ti.ABPresLHFront)
	tl.KAx1RNPerm = lookupAirSpringCurve(ti.ABPresRHFront)
	tl.KAx2LNPerm = lookupAirSpringCurve(ti.ABPresLHMID)
	tl.KAx2RNPerm = lookupAirSpringCurve(ti.ABPresRHMID)
	tl.KAx3LNPerm = lookupAirSpringCurve(ti.ABPresLHRear)
	tl.KAx3RNPerm = lookupAirSpringCurve(ti.ABPresRHRear)

	// Inertia forces (Z)
	tl.FInertiaAx1LN = calcInertiaForce(ti.AxleLoadLHFront, ti.DLAxle1LAZ)
	tl.FInertiaAx1RN = calcInertiaForce(ti.AxleLoadRHFront, ti.DLAxle1RAZ)
	tl.FInertiaAx2LN = calcInertiaForce(ti.AxleLoadLHCentre, ((ti.DLAxle1LAZ + ti.DLAxle2LAZ) / 2))
	tl.FInertiaAx2RN = calcInertiaForce(ti.AxleLoadRHCentre, ((ti.DLAxle1RAZ + ti.DLAxle2RAZ) / 2))
	tl.FInertiaAx3LN = calcInertiaForce(ti.AxleLoadLHRear, ti.DLAxle2LAZ)
	tl.FInertiaAx3RN = calcInertiaForce(ti.AxleLoadRHRear, ti.DLAxle2RAZ)

	// Suspension deflections
	tl.XAx1Lm = calcSuspensionDeflection(tl.FAirAxlex1LN, tl.FInertiaAx1LN, tl.KAx1LNPerm)
	tl.XAx1Rm = calcSuspensionDeflection(tl.FAirAxlex1RN, tl.FInertiaAx1RN, tl.KAx1RNPerm)
	tl.XAx2Lm = calcSuspensionDeflection(tl.FAirAxlex2LN, tl.FInertiaAx2LN, tl.KAx2LNPerm)
	tl.XAx2Rm = calcSuspensionDeflection(tl.FAirAxlex2RN, tl.FInertiaAx2RN, tl.KAx2RNPerm)
	tl.XAx3Lm = calcSuspensionDeflection(tl.FAirAxlex3LN, tl.FInertiaAx3LN, tl.KAx3LNPerm)
	tl.XAx3Rm = calcSuspensionDeflection(tl.FAirAxlex3RN, tl.FInertiaAx3RN, tl.KAx3RNPerm)

	// Velocities
	tl.XDotAx1LMps = calcDeltaOverT(tl.XAx1Lm, last.XAx1Lm, tl.DeltaT)
	tl.XDotAx1RMps = calcDeltaOverT(tl.XAx1Rm, last.XAx1Rm, tl.DeltaT)
	tl.XDotAx2LMps = calcDeltaOverT(tl.XAx2Lm, last.XAx2Lm, tl.DeltaT)
	tl.XDotAx2RMps = calcDeltaOverT(tl.XAx2Rm, last.XAx2Rm, tl.DeltaT)
	tl.XDotAx3LMps = calcDeltaOverT(tl.XAx3Lm, last.XAx3Lm, tl.DeltaT)
	tl.XDotAx3RMps = calcDeltaOverT(tl.XAx3Rm, last.XAx3Rm, tl.DeltaT)

	tl.XDotFAx1LMps = calcDeltaOverT(tl.XDotAx1LMps, last.XDotAx1LMps, tl.DeltaT)
	tl.XDotFAx1RMps = calcDeltaOverT(tl.XDotAx1RMps, last.XDotAx1RMps, tl.DeltaT)
	tl.XDotFAx2LMps = calcDeltaOverT(tl.XDotAx2LMps, last.XDotAx2LMps, tl.DeltaT)
	tl.XDotFAx2RMps = calcDeltaOverT(tl.XDotAx2RMps, last.XDotAx2RMps, tl.DeltaT)
	tl.XDotFAx3LMps = calcDeltaOverT(tl.XDotAx3LMps, last.XDotAx3LMps, tl.DeltaT)
	tl.XDotFAx3RMps = calcDeltaOverT(tl.XDotAx3RMps, last.XDotAx3RMps, tl.DeltaT)

	// Damper forces
	tl.FdAx1LN = calcDamperForce(tl.XDotAx1LMps)
	tl.FdAx1RN = calcDamperForce(tl.XDotAx1RMps)
	tl.FdAx2LN = calcDamperForce(tl.XDotAx2LMps)
	tl.FdAx2RN = calcDamperForce(tl.XDotAx2RMps)
	tl.FdAx3LN = calcDamperForce(tl.XDotAx3LMps)
	tl.FdAx3RN = calcDamperForce(tl.XDotAx3RMps)

	// Vertical reactions
	tl.FzAx1LN = calcVerticalReactions(tl.FAirAxlex1LN, tl.FdAx1LN, tl.FInertiaAx1LN)
	tl.FzAx1RN = calcVerticalReactions(tl.FAirAxlex1RN, tl.FdAx1RN, tl.FInertiaAx1RN)
	tl.FzAx2LN = calcVerticalReactions(tl.FAirAxlex2LN, tl.FdAx2LN, tl.FInertiaAx2LN)
	tl.FzAx2RN = calcVerticalReactions(tl.FAirAxlex2RN, tl.FdAx2RN, tl.FInertiaAx2RN)
	tl.FzAx3LN = calcVerticalReactions(tl.FAirAxlex3LN, tl.FdAx3LN, tl.FInertiaAx3LN)
	tl.FzAx3RN = calcVerticalReactions(tl.FAirAxlex3RN, tl.FdAx3RN, tl.FInertiaAx3RN)

	// Total Axle Loads
	tl.FzTotalN = tl.FzAx1LN + tl.FzAx2LN + tl.FzAx3LN + tl.FzAx1RN + tl.FzAx2RN + tl.FzAx3LN
	tl.FzLTotalN = tl.FzAx1LN + tl.FzAx2LN + tl.FzAx3LN
	tl.FzRTotalN = tl.FzAx1RN + tl.FzAx2RN + tl.FzAx3LN
	tl.FzAx1TotalN = tl.FzAx1LN + tl.FzAx1RN
	tl.FzAx2TotalN = tl.FzAx2LN + tl.FzAx2RN
	tl.FzAx3TotalN = tl.FzAx3LN + tl.FzAx3RN

	// Per hanger absolute
	tl.FzAx1AbsLN = calcPerHangerAbs(ti.AxleLoadLHFront, tl.FzAx1LN)
	tl.FzAx1AbsRN = calcPerHangerAbs(ti.AxleLoadRHFront, tl.FzAx1RN)
	tl.FzAx2AbsLN = calcPerHangerAbs(ti.AxleLoadLHCentre, tl.FzAx2LN)
	tl.FzAx2AbsRN = calcPerHangerAbs(ti.AxleLoadRHCentre, tl.FzAx2RN)
	tl.FzAx3AbsLN = calcPerHangerAbs(ti.AxleLoadLHRear, tl.FzAx3LN)
	tl.FzAx3AbsRN = calcPerHangerAbs(ti.AxleLoadRHRear, tl.FzAx3RN)

	// Load transfer
	tl.LTRGroup = (tl.FzLTotalN - tl.FzRTotalN) / tl.FzTotalN
	tl.DeltaFzLatAx1 = tl.FzAx1LN - tl.FzAx1RN
	tl.DeltaFzLatAx2 = tl.FzAx2LN - tl.FzAx2RN
	tl.DeltaFzLatAx3 = tl.FzAx3LN - tl.FzAx3RN

	// Vehicle states
	tl.UMps = ((ti.SpeedSat + ti.Speed) / 2) / 3.6 // Until WheelBasedVehSpeed can be used
	tl.RKpRps = ti.DLKingpinGZ
	tl.AyG = ti.DLKingpinAY

	// Roll moment
	tl.MRollNm = (tl.FzLTotalN + tl.FzRTotalN) * (t.Track_Width_T / 2)

	// Brake pressure and force
	tl.BPAxBar = convertKpaToBar(ti.BrakeApplicationPressureH)

	// Braking pressure per side
	tl.BPAx1LBar = tl.BPAxBar * tl.FzAx1LN / (tl.FzAx1LN + tl.FzAx1RN)
	tl.BPAx1RBar = tl.BPAxBar * tl.FzAx1RN / (tl.FzAx1LN + tl.FzAx1RN)
	tl.BPAx2LBar = tl.BPAxBar * tl.FzAx2LN / (tl.FzAx2LN + tl.FzAx2RN)
	tl.BPAx2RBar = tl.BPAxBar * tl.FzAx2RN / (tl.FzAx2LN + tl.FzAx2RN)
	tl.BPAx3LBar = tl.BPAxBar * tl.FzAx3LN / (tl.FzAx3LN + tl.FzAx3RN)
	tl.BPAx3RBar = tl.BPAxBar * tl.FzAx3RN / (tl.FzAx3LN + tl.FzAx3RN)

	// Wheel brake torque
	tl.TWheelAx1L = calcWheelBrakeTorque(tl.BPAx1LBar)
	tl.TWheelAx1R = calcWheelBrakeTorque(tl.BPAx1RBar)
	tl.TWheelAx2L = calcWheelBrakeTorque(tl.BPAx2LBar)
	tl.TWheelAx2R = calcWheelBrakeTorque(tl.BPAx2RBar)
	tl.TWheelAx3L = calcWheelBrakeTorque(tl.BPAx3LBar)
	tl.TWheelAx3R = calcWheelBrakeTorque(tl.BPAx3RBar)

	// Per-side force
	tl.FxBrakeAx1L = calcPerSideForce(tl.TWheelAx1L, tl.FzAx1LN)
	tl.FxBrakeAx1R = calcPerSideForce(tl.TWheelAx1R, tl.FzAx1RN)
	tl.FxBrakeAx2L = calcPerSideForce(tl.TWheelAx2L, tl.FzAx2LN)
	tl.FxBrakeAx2R = calcPerSideForce(tl.TWheelAx2R, tl.FzAx2RN)
	tl.FxBrakeAx3L = calcPerSideForce(tl.TWheelAx3L, tl.FzAx3LN)
	tl.FxBrakeAx3R = calcPerSideForce(tl.TWheelAx3R, tl.FzAx3RN)

	// Per-axle
	tl.FxBrakeAx1Total = tl.FxBrakeAx1L + tl.FxBrakeAx1R
	tl.FxBrakeAx2Total = tl.FxBrakeAx2L + tl.FxBrakeAx2R
	tl.FxBrakeAx3Total = tl.FxBrakeAx3L + tl.FxBrakeAx3R

	// Total
	tl.FxBrakeTotal = tl.FxBrakeAx1Total + tl.FxBrakeAx2Total + tl.FxBrakeAx3Total

	// Slip angles
	tl.BetaBodyRad = math.Atan(ti.DLKingpinAY * g)
	tl.AlphaAx1LRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_1, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)
	tl.AlphaAx1RRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_1, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)
	tl.AlphaAx2LRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_2, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)
	tl.AlphaAx2RRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_2, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)
	tl.AlphaAx3LRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_3, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)
	tl.AlphaAx3RRad = calcSlipAngle(tl.BetaBodyRad, t.CG_to_Axle_3, ti.DLKingpinAY, ti.Speed, t.Track_Width_T)

	// Lateral tyre forces - linear cornering stiffness
	tl.FyAx1LN = calcLateralTyreForces(tl.AlphaAx1LRad, tl.FzAx1LN)
	tl.FyAx1RN = calcLateralTyreForces(tl.AlphaAx1RRad, tl.FzAx1RN)
	tl.FyAx2LN = calcLateralTyreForces(tl.AlphaAx2LRad, tl.FzAx2LN)
	tl.FyAx2RN = calcLateralTyreForces(tl.AlphaAx2RRad, tl.FzAx2RN)
	tl.FyAx3LN = calcLateralTyreForces(tl.AlphaAx3LRad, tl.FzAx3LN)
	tl.FyAx3RN = calcLateralTyreForces(tl.AlphaAx3RRad, tl.FzAx3RN)
	tl.FyTotalN = tl.FyAx1LN + tl.FyAx1RN + tl.FyAx2LN + tl.FyAx2RN + tl.FyAx3LN + tl.FyAx3RN

	// Yaw moment
	tl.MzYawNm = (tl.FyAx1LN+tl.FyAx1RN)*t.CG_to_Axle_1 + (tl.FyAx2LN+tl.FyAx2RN)*t.CG_to_Axle_2 + (tl.FyAx3LN+tl.FyAx3RN)*t.CG_to_Axle_3

	// Bias
	tl.MassEstRearkg = ti.AxleLoadLHFront + ti.AxleLoadRHFront + ti.AxleLoadLHCentre + ti.AxleLoadRHCentre + ti.AxleLoadLHRear + ti.AxleLoadRHRear
	tl.RollTransferGain = (tl.FzLTotalN + tl.FzRTotalN) / ayGZeroThresh
	tl.DeflectBiasAx1 = (tl.XAx1Lm + tl.XAx1Rm) / 2
	tl.DeflectBiasAx2 = (tl.XAx2Lm + tl.XAx2Rm) / 2
	tl.DeflectBiasAx3 = (tl.XAx3Lm + tl.XAx3Rm) / 2

	// Kingpin forces
	tl.MassEffective = estimateEffectiveMass(t.Total_Mass_Trailer, tl.MassEstRearkg)
	tl.FxKingpinN = tl.FxBrakeTotal
	tl.FyKingpinN = tl.FyTotalN - (0.5 * dragForce * (ti.Speed * ti.Speed))
	tl.FzKingpinN = tl.MassEffective*(g+ti.DLKingpinAZ*g) - tl.FzTotalN

	return nil
}

// Determine if the trailer is loaded.
func calcTrailerLoaded(abp float64) string {
	if abp > 1.0 {
		staticPressure = 3.0
		return "Loaded"
	}
	staticPressure = 0.65
	return "Unloaded"
}

// Calculate the air spring forces
func calcAirSpringForce(abp float64, ratio float64) float64 {
	return ((abp - staticPressure) * math.Pow10(5) * airbagEffectiveArea * ratio)
}

// Lookup the air spring value from the curve map
func lookupAirSpringCurve(p float64) float64 {
	if p < 0.5 {
		// Below lower limit
		return AirSpring[0.5]
	} else if p > 9.0 {
		// Above upper limit
		return AirSpring[9.0]
	} else {
		// Round to 1 decimal
		pres := math.Round(p*10) / 10
		return AirSpring[float32(pres)]
	}
}

// Calculate the inertia forces
func calcInertiaForce(axleload, accelz float64) float64 {
	return (axleload * (accelz * g))
}

func calcSuspensionDeflection(air, inertia, kaxm float64) float64 {
	return ((air - inertia) / kaxm)
}

func calcDeltaOverT(now, last, dt float64) float64 {
	return ((now - last) / dt)
}

func calcDamperForce(xdot float64) float64 {
	return (damperCoeffAxlePerSide * xdot)
}

func calcVerticalReactions(fair, fd, fintertia float64) float64 {
	return (fair + fd - fintertia)
}

func calcPerHangerAbs(al, fz float64) float64 {
	return ((al * g) + fz)
}

// Convert a pressure in kPa to bar
func convertKpaToBar(in float64) float64 {
	return ((in / 100) * 0.3333)
}

func calcWheelBrakeTorque(bp float64) float64 {
	return (permittedTorque * (bp / pRef))
}

func calcPerSideForce(tw, fz float64) float64 {
	a := tw / rollingRadius
	b := uBrake * fz
	return (math.Min(a, b))
}

func calcSlipAngle(bbr, cg, dlk, speed, track float64) float64 {
	return (-1 * (bbr + (cg*dlk)/speed + (track/2)*dlk/speed))
}

func calcLateralTyreForces(alpha, fz float64) float64 {
	test := uBrake * fz
	result := cAlphaRef * alpha
	// Force can't go above uBrake * FzAx
	if result > test {
		return test
	}
	return result
}

func estimateEffectiveMass(trailer, est float64) float64 {
	if est < 10000 {
		return trailer
	}
	return (2 * est)
}
