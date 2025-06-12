package types

type WeaponParams struct {
	Name                                  string `csv:"Name" json:"name"`
	Mass                                  string `csv:"Mass" json:"mass_kg"`
	MassAtEndOfBoosterBurn                string `csv:"Mass at end of booster burn" json:"mass_end_booster_burn_kg"`
	MassAtEndOfSustainerBurn              string `csv:"Mass at end of sustainer burn" json:"mass_end_sustainer_burn_kg"`
	Caliber                               string `csv:"Calibre" json:"caliber_mm"`
	Length                                string `csv:"Length" json:"length_m"`
	ForceExertedByBooster                 string `csv:"Force exerted by booster" json:"force_exerted_by_booster_N"`
	BurnTimeOfBooster                     string `csv:"Burn time of booster" json:"burn_time_of_booster_s"`
	RawAccelerationAtIgnition             string `csv:"Raw acceleration at ignition" json:"raw_acceleration_at_ignition_ms2"`
	SpecificImpulseOfBooster              string `csv:"Specific impulse of booster" json:"specific_impulse_of_booster_s"`
	DeltaVOfBooster                       string `csv:"ΔV of booster" json:"delta_v_of_booster_ms"`
	BoosterStartDelay                     string `csv:"Booster start delay" json:"booster_start_delay_s"`
	ForceExertedBySustainer               string `csv:"Force exerted by sustainer" json:"force_exerted_by_sustainer_N"`
	BurnTimeOfSustainer                   string `csv:"Burn time of sustainer" json:"burn_time_of_sustainer_s"`
	SpecificImpulseOfSustainer            string `csv:"Specific impulse of sustainer" json:"specific_impulse_of_sustainer_s"`
	DeltaVOfSustainer                     string `csv:"ΔV of sustainer" json:"delta_v_of_sustainer_ms"`
	TotalDeltaV                           string `csv:"Total ΔV" json:"total_delta_v_ms"`
	ExplosiveMass                         string `csv:"Explosive mass" json:"explosive_mass_kg_tnt"`
	Warhead                               string `csv:"Warhead:" json:"warhead"`
	Penetration                           string `csv:"Penetration" json:"penetration_mm"`
	ProximityFuse                         string `csv:"Proximity fuse:" json:"proximity_fuse"`
	ProximityFuseArmingDistance           string `csv:"Proximity fuse arming distance" json:"proximity_fuse_arming_distance"`
	ProximityFuseArmingDistanceFromTarget string `csv:"Proximity fuse arming distance from target" json:"proximity_fuse_arming_distance_from_target"`
	ProximityFuseRange                    string `csv:"Proximity fuse range" json:"proximity_fuse_range_m"`
	ProximityFuseShellDetection           string `csv:"Proximity fuse shell detection (80-200 mm):" json:"proximity_fuse_shell_detection"`
	ProximityFuseMinimumAltitude          string `csv:"Proximity fuse minimum altitude" json:"proximity_fuse_minimum_atitude"`
	ProximityFuseDelay                    string `csv:"Proximity fuse delay" json:"proximity_fuse_delay_s"`
	ImpactFuseSensitivity                 string `csv:"Impact fuse sensitivity" json:"impact_fuse_sensitivity_mm"`
	ImpactFuseDelay                       string `csv:"Impact fuse delay" json:"impact_fuse_delay_m"`
	GuidanceType                          string `csv:"Guidance type" json:"guidance_type,omitempty"`
	GuidanceStartDelay                    string `csv:"Guidance start delay" json:"guidance_start_delay_s,omitempty"`
	GuidanceDuration                      string `csv:"Guidance duration" json:"guidance_duration_s,omitempty"`
	GuidanceRange                         string `csv:"Guidance range" json:"guidance_range_km,omitempty"`
	GuidanceFOV                           string `csv:"Guidance FOV" json:"guidance_fov_deg,omitempty"`
	GuidanceMaxLead                       string `csv:"Guidance max lead" json:"guidance_max_lead_deg,omitempty"`
	GuidanceLaunchSector                  string `csv:"Guidance launch sector" json:"guidance_launch_sector_deg,omitempty"`
	AimTrackingSensitivity                string `csv:"Aim tracking sensitivity" json:"aim_tracking_sensitivity,omitempty"`
	SeekerWarmUpTime                      string `csv:"Seeker warm up time" json:"seeker_warm_up_time_s,omitempty"`
	SeekerSearchDuration                  string `csv:"Seeker search duration" json:"seeker_search_duration_s,omitempty"`
	SeekerRange                           string `csv:"Seeker range" json:"seeker_range_km,omitempty"`
	FieldOfView                           string `csv:"Field of view" json:"field_of_view_deg,omitempty"`
	GimbalLimit                           string `csv:"Gimbal limit" json:"gimbal_limit_deg,omitempty"`
	TrackRate                             string `csv:"Track rate" json:"track_rate_deg_sec,omitempty"`
	UncagedSeekerBeforeLaunch             string `csv:"Uncaged seeker before launch" json:"uncaged_seeker_before_launch,omitempty"`
	MaxLockAngleBeforeLaunch              string `csv:"Maximum lock angle before launch" json:"max_lock_angle_before_launch_deg,omitempty"`
	MinAngleOfIncidenceToSun              string `csv:"Minimum angle of incidence to Sun" json:"min_angle_of_incidence_to_sun_deg,omitempty"`
	BaselineLockRangeRear                 string `csv:"Baseline lock range rear-aspect" json:"baseline_lock_range_rear_km,omitempty"`
	BaselineLockRangeAll                  string `csv:"Baseline lock range all-aspect" json:"baseline_lock_range_all_km,omitempty"`
	BaselineLockRangeGround               string `csv:"Baseline lock range (ground)" json:"baseline_lock_range_ground_km,omitempty"`
	BaselineLockRangeTarget               string `csv:"Baseline lock range (target)" json:"baseline_lock_range_target_km,omitempty"`
	BaselineFlareDetection                string `csv:"Baseline flare detection" json:"baseline_flare_detection_km,omitempty"`
	BaselineIRCMDetection                 string `csv:"Baseline IRCM detection" json:"baseline_ircm_detection_km,omitempty"`
	BaselineDIRCMDetection                string `csv:"Baseline DIRCM detection" json:"baseline_dircm_detection_km,omitempty"`
	BaselineLDIRCMDetection               string `csv:"Baseline LDIRCM detection" json:"baseline_ldircm_detection_km,omitempty"`
	BaselineHeadOnLockRange               string `csv:"Baseline head-on lock range" json:"baseline_head_on_lock_range_km,omitempty"`
	MaxLockRangeHardLimit                 string `csv:"Maximum lock range" json:"max_lock_range_km,omitempty"`
	IRCCM                                 string `csv:"IRCCM" json:"irccm,omitempty"`
	IRCCMType                             string `csv:"IRCCM type" json:"irccm_type,omitempty"`
	IRCCMFieldOfView                      string `csv:"IRCCM field of view" json:"irccm_field_of_view_deg,omitempty"`
	IRCCMRejectionThreshold               string `csv:"IRCCM rejection threshold" json:"irccm_rejection_threshold,omitempty"`
	IRCCMReactionTime                     string `csv:"IRCCM reaction time" json:"irccm_reaction_time_s,omitempty"`
	MinTargetSize                         string `csv:"Minimum target size" json:"min_target_size_m,omitempty"`
	MaxBreakLockTime                      string `csv:"Maximum break lock time" json:"max_break_lock_time_s,omitempty"`
	CanBeSlavedToRadar                    string `csv:"Can be slaved to radar" json:"can_be_slaved_to_radar,omitempty"`
	CanLockAfterLaunch                    string `csv:"Can lock after launch" json:"can_lock_after_launch,omitempty"`
	Band                                  string `csv:"Band" json:"band,omitempty"`
	AngularSpeedRejectionThresh           string `csv:"Angular speed rejection" json:"angular_speed_rejection_deg_s,omitempty"`
	AccelRejectionThreshRange             string `csv:"Acceleration rejection" json:"accel_rejection_m_s2,omitempty"`
	InertialGuidanceDriftSpeed            string `csv:"Inertial guidance drift" json:"inertial_guidance_drift_m_s,omitempty"`
	Datalink                              string `csv:"Datalink" json:"datalink,omitempty"`
	CanDatalinkReconnect                  string `csv:"Can datalink reconnect" json:"can_datalink_reconnect,omitempty"`
	SidelobeAttenuation                   string `csv:"Sidelobe attenuation" json:"sidelobe_attenuation,omitempty"`
	TransmitterPower                      string `csv:"Transmitter power" json:"transmitter_power,omitempty"`
	TransmitterHalfSensitivity            string `csv:"Transmitter half sensitivity" json:"transmitter_half_sensitivity,omitempty"`
	TransmitterSidelobeSens               string `csv:"Transmitter sidelobe sensitivity" json:"transmitter_sidelobe_sensitivity,omitempty"`
	ReceiverHalfSensitivity               string `csv:"Receiver half sensitivity" json:"receiver_half_sensitivity,omitempty"`
	ReceiverSidelobeSens                  string `csv:"Receiver sidelobe sensitivity" json:"receiver_sidelobe_sensitivity,omitempty"`
	DistanceMinValue                      string `csv:"Distance min" json:"distance_min_m,omitempty"`
	DistanceMaxValue                      string `csv:"Distance max" json:"distance_max_km,omitempty"`
	DistanceWidth                         string `csv:"Distance width" json:"distance_width_m,omitempty"`
	DistanceRefWidth                      string `csv:"Distance ref width" json:"distance_ref_width_m,omitempty"`
	DistanceMinSignalGate                 string `csv:"Distance min signal gate" json:"distance_min_signal_gate_m,omitempty"`
	DistanceGateSearchRange               string `csv:"Distance gate search" json:"distance_gate_search_m,omitempty"`
	DistanceGateAlphaFilter               string `csv:"Distance gate alpha" json:"distance_gate_alpha,omitempty"`
	DistanceGateBetaFilter                string `csv:"Distance gate beta" json:"distance_gate_beta,omitempty"`
	DopplerSpeedMinValue                  string `csv:"Doppler speed min" json:"doppler_speed_min_m_s,omitempty"`
	DopplerSpeedMaxValue                  string `csv:"Doppler speed max" json:"doppler_speed_max_m_s,omitempty"`
	DopplerSpeedWidth                     string `csv:"Doppler speed width" json:"doppler_speed_width_m_s,omitempty"`
	DopplerSpeedRefWidth                  string `csv:"Doppler speed ref width" json:"doppler_speed_ref_width_m_s,omitempty"`
	DopplerSpeedMinSignalGate             string `csv:"Doppler speed min gate" json:"doppler_speed_min_gate_m_s,omitempty"`
	DopplerSpeedGateSearch                string `csv:"Doppler speed gate search" json:"doppler_speed_gate_search_m_s,omitempty"`
	DopplerSpeedGateAlpha                 string `csv:"Doppler speed gate alpha" json:"doppler_speed_gate_alpha,omitempty"`
	DopplerSpeedGateBeta                  string `csv:"Doppler speed gate beta" json:"doppler_speed_gate_beta,omitempty"`
	ProportionalNavMultiplier             string `csv:"Proportional nav multiplier" json:"proportional_nav_multiplier,omitempty"`
	BaseIndicatedAirSpeed                 string `csv:"Base air speed" json:"base_air_speed_m_s,omitempty"`
	PIDProportionalTerm                   string `csv:"PID proportional" json:"pid_proportional,omitempty"`
	PIDIntegralTerm                       string `csv:"PID integral" json:"pid_integral,omitempty"`
	PIDIntegralTermLimit                  string `csv:"PID integral term limit" json:"pid_integral_limit,omitempty"`
	PIDDerivativeTerm                     string `csv:"PID derivative" json:"pid_derivative,omitempty"`
	OrientingPhase                        string `csv:"Orienting phase" json:"orienting_phase,omitempty"`
	OrientingStartDelay                   string `csv:"Orienting start delay" json:"orienting_start_delay,omitempty"`
	OrientingControlTime                  string `csv:"Orienting control time" json:"orienting_control_time,omitempty"`
	OrientingElevationAddition            string `csv:"Orienting elevation addition" json:"orienting_elevation_addition,omitempty"`
	DragCoefficientMultiplier             string `csv:"Drag coefficient multiplier" json:"drag_coefficient_multiplier,omitempty"`
	WingAreaMultiplier                    string `csv:"Wing area multiplier" json:"wing_area_multiplier,omitempty"`
	StartSpeed                            string `csv:"Start speed" json:"start_speed,omitempty"`
	MaximumSpeed                          string `csv:"Maximum speed" json:"maximum_speed,omitempty"`
	MinimumRange                          string `csv:"Minimum range" json:"minimum_range,omitempty"`
	FlightRangeLimit                      string `csv:"Flight range limit" json:"flight_range_limit,omitempty"`
	MaximumGLoad                          string `csv:"Maximum G-load" json:"maximum_g_load,omitempty"`
	MaximumFinAngleOfAttack               string `csv:"Maximum fin angle of attack" json:"maximum_fin_angle_of_attack,omitempty"`
	MaximumFinLateralAcceleration         string `csv:"Maximum fin lateral acceleration" json:"maximum_fin_lateral_acceleration,omitempty"`
	MaximumLateralAcceleration            string `csv:"Maximum lateral acceleration" json:"maximum_lateral_acceleration,omitempty"`
	MaximumAOA                            string `csv:"Maximum AOA" json:"maximum_aoa,omitempty"`
	ThrustVectoring                       string `csv:"Thrust vectoring" json:"thrust_vectoring,omitempty"`
	ThrustVectoringAngle                  string `csv:"Thrust vectoring angle" json:"thrust_vectoring_angle,omitempty"`
	MaximumLaunchAngleHorizontal          string `csv:"Maximum launch angle (horizontally)" json:"maximum_launch_angle_horizontal,omitempty"`
	MaximumLaunchAngleVertical            string `csv:"Maximum launch angle (vertically)" json:"maximum_launch_angle_vertical,omitempty"`
	MaximumAxisValues                     string `csv:"Maximum axis values" json:"maximum_axis_values,omitempty"`
	StatcardSpeed                         string `csv:"Maximum statcard (useless) speed" json:"statcard_speed,omitempty"`
	StatcardLaunchRange                   string `csv:"Maximum statcard (useless) launch range" json:"statcard_launch_range,omitempty"`
	StatcardGuaranteedRange               string `csv:"Statcard (useless) guaranteed range" json:"statcard_guaranteed_range,omitempty"`
	StatcardGLoad                         string `csv:"Maximum statcard (useless) G-load" json:"statcard_g_load,omitempty"`
	FlightTimeUntilGuidanceStarts         string `csv:"Flight time until guidance starts (delay)" json:"flight_time_until_guidance_starts,omitempty"`
	FlightTimeWhenPullLimitX              string `csv:"Flight time when pull limit reaches x%" json:"flight_time_when_pull_limit_x,omitempty"`
	FlightTimeWhenPullLimit100            string `csv:"Flight time when pull limit reaches 100%" json:"flight_time_when_pull_limit_100,omitempty"`
	ETAtoImpactWhenPropMultiplier         string `csv:"ETA to impact when prop multiplier reaches x%" json:"eta_to_impact_when_prop_multiplier,omitempty"`
	Loft                                  string `csv:"Loft" json:"loft,omitempty"`
	LoftAngle                             string `csv:"Loft angle" json:"loft_angle,omitempty"`
	TargetElevation                       string `csv:"Target elevation" json:"target_elevation,omitempty"`
	MaximumTargetAngularChange            string `csv:"Maximum target angular change" json:"maximum_target_angular_change,omitempty"`
	HasTracerInTail                       string `csv:"Has a tracer in its tail" json:"has_tracer_in_tail,omitempty"`
	SeaSkimming                           string `csv:"Sea skimming" json:"sea_skimming,omitempty"`
	SkimAltitude                          string `csv:"Skim altitude" json:"skim_altitude,omitempty"`
	AttackAltitude                        string `csv:"Attack altitude" json:"attack_altitude,omitempty"`
	AdditionalNotes                       string `csv:"Additional Notes:" json:"additional_notes"`
}
