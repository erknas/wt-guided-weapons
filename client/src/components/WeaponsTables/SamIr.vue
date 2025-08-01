<script setup>
defineProps({
	weapons: {
		type: Array,
		required: true
	}
})
</script>

<template>
<div class="table-container-wrapper">
    <table class="table">
      <thead>
        <tr>
          <th class="sticky-col sticky-row header-cell-name">Name:</th>
          <th 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="header-cell"
          >
            {{ weapon.name }}
          </th>
        </tr>
        <tr class="physical-props-row no-hover">
          <td class="sticky-col physical-props no-hover">
            Physical properties
          </td>
          <td :colspan="weapons.length" class="physical-props-fill no-hover"></td>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="sticky-col header-cell-param">Mass: [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Mass at end of booster burn: [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass_at_end_of_booster_burn || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Mass at end of sustainer burn: [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass_at_end_of_sustainer_burn || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Calibre: [mm]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.caliber || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Length: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.length || '-' }}
          </td>
        </tr>
        <tr class="engine-props-row no-hover">
          <td class="sticky-col engine-props no-hover">
            Engine properties
          </td>
          <td :colspan="weapons.length" class="engine-props-fill no-hover"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Force exerted by booster: [N]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.force_exerted_by_booster || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Burn time of booster: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.burn_time_of_booster || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Raw acceleration at ignition: [m/s²]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.raw_acceleration_at_ignition || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Specific impulse of booster: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.specific_impulse_of_booster || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">ΔV of booster: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.delta_v_of_booster || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Booster start delay: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.booster_start_delay || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Force exerted by sustainer: [N]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.force_exerted_by_sustainer || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Burn time of sustainer: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.burn_time_of_sustainer || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Specific impulse of sustainer: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.specific_impulse_of_sustainer || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">ΔV of sustainer: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.delta_v_of_sustainer || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Total ΔV: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.total_delta_v || '-' }}
          </td>
        </tr>
        <tr class="fuse-warhead-props-row no-hover">
          <td class="sticky-col fuse-warhead-props no-hover">
            Fuse and warhead properties
          </td>
          <td :colspan="weapons.length" class="fuse-warhead-props-fill no-hover"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Explosive mass: [kg of TNT equivalent]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.explosive_mass || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Warhead:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.warhead || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Penetration: [mm]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.penetration || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.proximity_fuse || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse range: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.proximity_fuse_range || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse arming distance: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.proximity_fuse_arming_distance || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse shell detection (80-200 mm):</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.proximity_fuse_shell_detection || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse minimum altitude: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.proximity_fuse_minimum_altitude || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proximity fuse delay: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.proximity_fuse_delay || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Impact fuse sensitivity: [mm]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.impact_fuse_sensitivity || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Impact fuse delay: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.impact_fuse_delay|| '-' }}
          </td>
        </tr>
        <tr class="guidance-props-row no-hover">
          <td class="sticky-col guidance-props no-hover">
            Guidance properties
          </td>
          <td :colspan="weapons.length" class="guidance-props-fill no-hover"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Guidance type:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.guidance_type || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Guidance start delay: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.guidance_start_delay || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Guidance duration: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.guidance_duration || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Seeker warm up time: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.seeker_warm_up_time || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Seeker search duration: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.seeker_search_duration || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Field of view: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.field_of_view || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Gimbal limit: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.gimbal_limit || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Track rate: [degrees/second]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.track_rate || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Uncaged seeker before launch:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.uncaged_seeker_before_launch || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum lock angle before launch: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.max_lock_angle_before_launch|| '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Minimum angle of incidence of the seeker to the Sun for it to not capture the Sun: [degrees] </td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.min_angle_of_incidence_to_sun || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline lock range from rear-aspect: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_lock_range_rear || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline lock range from all-aspect: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_lock_range_all|| '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline flare detection range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_flare_detection || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline IRCM detection range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_ircm_detection || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline DIRCM detection range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_dircm_detection || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline LDIRCM detection range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_ldircm_detection || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Baseline head-on lock range against afterburning target: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_head_on_lock_range|| '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum lock range (hard limit): [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.max_lock_range || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">IRCCM:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.irccm || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">IRCCM type:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.irccm_type || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">IRCCM field of view: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.irccm_field_of_view || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">IRCCM rejection treshold: </td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.irccm_rejection_threshold || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">IRCCM reaction time: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.irccm_reaction_time || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum break lock time: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.max_break_lock_time || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Can be slaved to radar:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.can_be_slaved_to_radar || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Can lock after launch:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.can_lock_after_launch || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Inertial guidance drift speed: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.inertial_guidance_drift_ms || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Datalink:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.datalink || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Can datalink reconnect:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.can_datalink_reconnect || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Proportional navigation multiplier: (affects how far ahead it attempts to lead)</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.proportional_nav_multiplier|| '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Base indicated air speed: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.base_air_speed || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">PID proportional term:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.pid_proportional || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">PID integral term:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.pid_integral || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">PID integral term limit:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.pid_integral_limit || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">PID derivative term:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.pid_derivative || '-' }}
          </td>
        </tr>
        <tr class="orienting-props-row no-hover">
          <td class="sticky-col orienting-props no-hover">
            Orienting properties
          </td>
          <td :colspan="weapons.length" class="orienting-props-fill no-hover"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Orienting phase:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.orienting_phase || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Orienting start delay: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.orienting_start_delay || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Orienting control time: [s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.orienting_control_time || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Orienting elevation addition: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.orienting_elevation_addition || '-' }}
          </td>
        </tr>
        <tr class="flight-props-row no-hover">
          <td class="sticky-col flight-props no-hover">
            Flight properties
          </td>
          <td :colspan="weapons.length" class="flight-props-fill no-hover"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Drag coefficient multiplier (this is not the only value affecting drag, just because it's higher than another missile's doesn't mean it actually has higher drag!!):</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.drag_coefficient_multiplier || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum fin angle of attack: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_fin_angle_of_attack || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Fins lateral acceleration:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.fins_lateral_acceleration || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Wing area multiplier:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.wing_area_multiplier || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Start speed: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.start_speed || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum speed: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_speed || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum statcard (useless) speed: [Mach]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.statcard_speed_mach || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum statcard (useless) launch range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.statcard_launch_range || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Minimum range: [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.minimum_range || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Flight range limit: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.flight_range_limit || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum G-load: [G]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_g_load || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum statcard (useless) G-load: [G]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_statcard_g_load || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Thrust vectoring:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.thrust_vectoring || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Thrust vectoring angle: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.thrust_vectoring_angle || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Loft:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.loft || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Loft angle: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.loft_angle || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Target elevation: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.target_elevation || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Maximum target angular change: [degrees/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_target_angular_change || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col add-notes">Additional Notes: </td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell add-notes-data"
          >
            {{ weapon.additional_notes || '-' }}
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style scoped>
.orienting-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #e9a252;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
  padding: 5px 2px;
}

.orienting-props-row .orienting-props-fill{
  background-color: #e9a252;
}
</style>
