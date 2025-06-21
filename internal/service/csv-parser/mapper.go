package csvparser

import (
	"fmt"
	"strings"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

func mapCSVToStruct(data [][]string, category string, weaponIdx int) (*types.Weapon, error) {
	if len(data) == 0 || weaponIdx >= len(data[0]) {
		return nil, fmt.Errorf("invalid data")
	}

	headers := make(map[string]int)
	for i, row := range data {
		if len(row) > 0 {
			headers[strings.TrimSpace(row[0])] = i
		}
	}

	getValue := func(header string) string {
		if rowIdx, ok := headers[header]; ok {
			if rowIdx < len(data) && weaponIdx < len(data[rowIdx]) {
				return data[rowIdx][weaponIdx]
			}
		}
		return ""
	}

	weapon := new(types.Weapon)

	weapon.Category = category
	weapon.Name = getValue("Name:")
	weapon.Mass = getValue("Mass: [kg]")
	weapon.Caliber = getValue("Calibre: [mm]")
	weapon.Length = getValue("Length: [m]")
	weapon.MassAtEndOfBoosterBurn = getValue("Mass at end of booster burn: [kg]")
	weapon.ForceExertedByBooster = getValue("Force exerted by booster: [N]")
	weapon.BurnTimeOfBooster = getValue("Burn time of booster: [s]")
	weapon.SpecificImpulseOfBooster = getValue("Specific impulse of booster: [s]")
	weapon.DeltaVOfBooster = getValue("ΔV of booster: [m/s]")
	weapon.BoosterStartDelay = getValue("Booster start delay: [s]")
	weapon.MassAtEndOfSustainerBurn = getValue("Mass at end of sustainer burn: [kg]")
	weapon.ForceExertedBySustainer = getValue("Force exerted by sustainer: [N]")
	weapon.BurnTimeOfSustainer = getValue("Burn time of sustainer: [s]")
	weapon.SpecificImpulseOfSustainer = getValue("Specific impulse of sustainer: [s]")
	weapon.DeltaVOfSustainer = getValue("ΔV of sustainer: [m/s]")
	weapon.TotalDeltaV = getValue("Total ΔV: [m/s]")
	weapon.ExplosiveMass = getValue("Explosive mass: [kg of TNT equivalent]")
	weapon.Warhead = getValue("Warhead:")
	weapon.Penetration = getValue("Penetration: [mm]")
	weapon.ProximityFuse = getValue("Proximity fuse:")
	weapon.ProximityFuseArmingDistance = getValue("Proximity fuse arming distance: [m]")
	weapon.ProximityFuseArmingDistanceFromTarget = getValue("Proximity fuse arming distance from target: [m]")
	weapon.ProximityFuseRange = getValue("Proximity fuse range: [m]")
	weapon.ProximityFuseShellDetection = getValue("Proximity fuse shell detection (80-200 mm):")
	weapon.ProximityFuseMinimumAltitude = getValue("Proximity fuse minimum altitude: [m]")
	weapon.ProximityFuseDelay = getValue("Proximity fuse delay: [s]")
	weapon.ImpactFuseSensitivity = getValue("Impact fuse sensitivity: [mm]")
	weapon.ImpactFuseDelay = getValue("Impact fuse delay: [m]")
	weapon.GuidanceType = getValue("Guidance type:")
	weapon.GuidanceStartDelay = getValue("Guidance start delay: [s]")
	weapon.GuidanceDuration = getValue("Guidance duration: [s]")
	weapon.GuidanceRange = getValue("Guidance range: [km]")
	weapon.GuidanceFOV = getValue("Guidance FOV: [degrees]")
	weapon.GuidanceMaxLead = getValue("Guidance max lead: [degrees]")
	weapon.GuidanceLaunchSector = getValue("Guidance launch sector: [degrees]")
	weapon.AimTrackingSensitivity = getValue("Aim tracking sensitivity:")
	weapon.SeekerWarmUpTime = getValue("Seeker warm up time: [s]")
	weapon.SeekerSearchDuration = getValue("Seeker search duration: [s]")
	weapon.SeekerRange = getValue("Seeker range: [km]")
	weapon.FieldOfView = getValue("Field of view: [degrees]")
	weapon.GimbalLimit = getValue("Gimbal limit: [degrees]")
	weapon.TrackRate = getValue("Track rate: [degrees/second]")
	weapon.UncagedSeekerBeforeLaunch = getValue("Uncaged seeker before launch:")
	weapon.MaxLockAngleBeforeLaunch = getValue("Maximum lock angle before launch: [degrees]")
	weapon.MinAngleOfIncidenceToSun = getValue("Minimum angle of incidence of the seeker to the Sun for it to not capture the Sun: [degrees]")
	weapon.BaselineLockRangeRear = getValue("Baseline lock range rear-aspect: [km]")
	weapon.BaselineLockRangeAll = getValue("Baseline lock range all-aspect: [km]")
	weapon.BaselineLockRangeGround = getValue("Baseline lock range (ground): [km]")
	weapon.BaselineLockRangeTarget = getValue("Baseline lock range (target): [km]")
	weapon.BaselineFlareDetection = getValue("Baseline flare detection range: [km]")
	weapon.BaselineIRCMDetection = getValue("Baseline IRCM detection range: [km]")
	weapon.BaselineDIRCMDetection = getValue("Baseline DIRCM detection range: [km]")
	weapon.BaselineLDIRCMDetection = getValue("Baseline LDIRCM detection: [km]")
	weapon.BaselineHeadOnLockRange = getValue("Baseline head-on lock range against afterburning target: [km]")
	weapon.MaxLockRangeHardLimit = getValue("Maximum lock range (hard limit): [km]")
	weapon.IRCCM = getValue("IRCCM:")
	weapon.IRCCMType = getValue("IRCCM type:")
	weapon.IRCCMFieldOfView = getValue("IRCCM field of view: [degrees]")
	weapon.IRCCMRejectionThreshold = getValue("IRCCM rejection threshold:")
	weapon.IRCCMReactionTime = getValue("IRCCM reaction time: [s]")
	weapon.MinTargetSize = getValue("Minimum target size: [m]")
	weapon.MaxBreakLockTime = getValue("Maximum break lock time: [s]")
	weapon.CanBeSlavedToRadar = getValue("Can be slaved to radar:")
	weapon.CanLockAfterLaunch = getValue("Can lock after launch:")
	weapon.Band = getValue("Band:")
	weapon.AngularSpeedRejectionThresh = getValue("Angular speed rejection threshold: [degrees/second]")
	weapon.AccelRejectionThreshRange = getValue("Acceleration rejection threshold range: [m/s^2]")
	weapon.InertialGuidanceDriftSpeed = getValue("Inertial guidance drift speed: [m/s]")
	weapon.Datalink = getValue("Datalink:")
	weapon.CanDatalinkReconnect = getValue("Can datalink reconnect:")
	weapon.SidelobeAttenuation = getValue("Sidelobe attenuation:")
	weapon.TransmitterPower = getValue("Transmitter power:")
	weapon.TransmitterHalfSensitivity = getValue("Transmitter angle of half sensitivity:")
	weapon.TransmitterSidelobeSens = getValue("Transmitter sidelobe sensitivity:")
	weapon.ReceiverHalfSensitivity = getValue("Receiver angle of half sensitivity:")
	weapon.ReceiverSidelobeSens = getValue("Receiver sidelobe sensitivity:")
	weapon.ProportionalNavMultiplier = getValue("Proportional navigation multiplier: (affects how far ahead it attempts to lead)")
	weapon.BaseIndicatedAirSpeed = getValue("Base indicated air speed: [m/s]")
	weapon.PIDProportionalTerm = getValue("PID proportional term:")
	weapon.PIDIntegralTerm = getValue("PID integral term:")
	weapon.PIDIntegralTermLimit = getValue("PID integral term limit:")
	weapon.PIDDerivativeTerm = getValue("PID derivative term:")
	weapon.RawAccelerationAtIgnition = getValue("Raw acceleration at ignition: [m/s²]")
	weapon.StartSpeed = getValue("Start speed: [m/s]")
	weapon.MaximumSpeed = getValue("Maximum speed: [m/s]")
	weapon.MinimumRange = getValue("Minimum range: [m]")
	weapon.FlightRangeLimit = getValue("Flight range limit: [km]")
	weapon.MaximumGLoad = getValue("Maximum G-load: [G]")
	weapon.MaximumFinAngleOfAttack = getValue("Maximum fin angle of attack: [degrees]")
	weapon.MaximumFinLateralAcceleration = getValue("Maximum fin lateral acceleration:")
	weapon.MaxLateralAcceleration = getValue("Max lateral acceleration:")
	weapon.MaximumLateralAcceleration = getValue("Maximum lateral acceleration:")
	weapon.WingAreaMultiplier = getValue("Wing area multiplier:")
	weapon.MaximumAOA = getValue("Maximum AOA: [degrees]")
	weapon.ThrustVectoring = getValue("Thrust vectoring:")
	weapon.ThrustVectoringAngle = getValue("Thrust vectoring angle: [degrees]")
	weapon.MaximumLaunchAngleHorizontalVertical = getValue("Maximum launch angle (horizontally / vertically): [degrees]")
	weapon.MaximumAxisValues = getValue("Maximum axis values:")
	weapon.StatcardSpeed = getValue("Maximum statcard (useless) speed: [m/s] or [Mach]")
	weapon.StatcardLaunchRange = getValue("Maximum statcard (useless) launch range: [km]")
	weapon.DragCoefficientMultiplier = getValue("Drag coefficient multiplier (this is not the only value affecting drag, just because it's higher than another missile's doesn't mean it actually has higher drag!!):")
	weapon.StatcardGuaranteedRange = getValue("Statcard (useless) guaranteed range: [km]")
	weapon.StatcardGLoad = getValue("Statcard (useless) max G-load: [G]")
	weapon.FlightTimeUntilGuidanceStarts = getValue("Flight time until guidance starts (delay): [s]")
	weapon.FlightTimeWhenPullLimitX = getValue("Flight time when pull limit reaches x%: [s/%]")
	weapon.FlightTimeWhenPullLimit100 = getValue("Flight time when pull limit reaches 100%: [s]")
	weapon.ETAtoImpactWhenPropMultiplier = getValue("ETA to impact when prop multiplier reaches x%: [s/%]")
	weapon.Loft = getValue("Loft:")
	weapon.LoftAngle = getValue("Loft angle: [degrees]")
	weapon.TargetElevation = getValue("Target elevation: [degrees]")
	weapon.MaximumTargetAngularChange = getValue("Maximum target angular change: [degrees/s]")
	weapon.HasTracerInTail = getValue("Has a tracer in its tail:")
	weapon.SeaSkimming = getValue("Sea skimming:")
	weapon.SkimAltitude = getValue("Skim altitude: [m]")
	weapon.AttackAltitude = getValue("Attack altitude: [m]")
	weapon.AdditionalNotes = getValue("Additional Notes:")

	return weapon, nil
}
