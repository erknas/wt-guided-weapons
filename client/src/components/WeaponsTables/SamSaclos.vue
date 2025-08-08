<script setup>
const props = defineProps({
	weapons: {
		type: Array,
		required: true
	},
  highlightedWeapon: String 
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
            :data-weapon-name="weapon.name"
            :class="{'is-highlighted': highlightedWeapon === weapon.name }"
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
          <td class="sticky-col header-cell-param">Guidance range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.guidance_range || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Guidance FOV: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.guidance_fov || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Guidance max lead: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.guidance_max_lead || '-' }}
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
          <td class="sticky-col header-cell-param">Maximum launch angle: [degrees]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_launch_angle || '-' }}
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
          <td class="sticky-col header-cell-param">Maximum fin lateral acceleration:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.maximum_fin_lateral_acceleration || '-' }}
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
          <td class="sticky-col header-cell-param">Maximum statcard (useless) speed: [m/s]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.statcard_speed_ms || '-' }}
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
          <td class="sticky-col header-cell-param">Statcard (useless) guaranteed range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.statcard_guaranteed_range || '-' }}
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
          <td class="sticky-col header-cell-param">Flight time when pull limit reaches x%: [s/%]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.flight_time_when_pull_limit_x || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Thrust vectoring:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
            :class="{ 'text-right': !weapon.thrust_vectoring }"
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
          <td class="sticky-col header-cell-param">Has a tracer in its tail:</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell text-center"
          >
            {{ weapon.has_tracer_in_tail || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col add-notes">Additional Notes: </td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell add-notes-data"
          >
            {{ weapon.additional_notes }}
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
