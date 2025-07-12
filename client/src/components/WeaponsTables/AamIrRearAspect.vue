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
  <div class="table-container">
    <table class="table">
      <thead>
        <tr>
          <th class="sticky-col sticky-row header-cell-name">Name</th>
          <th 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="header-cell"
          >
            {{ weapon.name }}
          </th>
        </tr>
        <tr class="physical-props-row">
          <td class="sticky-col physical-props">
            Physical properties
          </td>
          <td colspan="100%" class="physical-props-fill"></td>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="sticky-col header-cell-param">Mass [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Mass at end of booster burn [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass_at_end_of_booster_burn || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Mass at end of sustainer burn [kg]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.mass_at_end_of_sustainer_burn || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Calibre [mm]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.caliber || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Length [m]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.length || '-' }}
          </td>
        </tr>
        <tr class="engine-props-row">
          <td class="sticky-col engine-props">
            Engine properties
          </td>
          <td colspan="100%" class="engine-props-fill"></td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Force exerted by booster [N]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.force_exerted_by_booster || '-' }}
          </td>
        </tr>
        <tr>
          <td class="sticky-col header-cell-param">Burn time of booster [s]</td>
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
        <tr class="fuse-warhead-props-row">
          <td class="sticky-col fuse-warhead-props">
            Fuse and warhead properties
          </td>
          <td colspan="100%" class="fuse-warhead-props-fill"></td>
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
        <tr class="guidance-props-row">
          <td class="sticky-col guidance-props">
            Guidance properties
          </td>
          <td colspan="100%" class="guidance-props-fill"></td>
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
          <td class="sticky-col header-cell-param">Baseline flare and baseline IRCM detection range: [km]</td>
          <td 
            v-for="weapon in weapons" 
            :key="weapon.id"
            class="data-cell"
          >
            {{ weapon.baseline_flare_and_ircm_detection_range || '-' }}
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
        <tr class="flight-props-row">
          <td class="sticky-col flight-props">
            Flight properties
          </td>
          <td colspan="100%" class="flight-props-fill"></td>
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
          <td class="sticky-col header-cell-param">Maximum fin lateral acceleration: </td>
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
</div>
</template>

<style scoped>
.table-container-wrapper {
  width: calc(100vw - 0px);
  height: calc(100vh - 120px);
  overflow: auto;
  position: fixed;
  top: 120px;
  left: 0px;
  right: 0px;
  bottom: 0px;
}

.table-container {
  width: max-content;  
  min-width: 100%;
}

.table {
  border-collapse: separate;
  border-spacing: 0;
  width: auto;
  color: black;
  background-color: whitesmoke;
  font-family: Arial, Helvetica, sans-serif;
}

th:not(.sticky-col),
td:not(.sticky-col) {
  border-right: 1px solid #bababa;
  border-bottom: 1px solid #bababa;
}

.sticky-col.sticky-row {
  position: sticky;
  left: 0;
  top: 0;
  z-index: 12;
  border-right: 5px solid #bababa;
  border-bottom: 5px solid #bababa;
  text-align: left;
  padding-left: 2px;
}

thead th {
  position: sticky;
  top: 0;
  z-index: 3;
  background: #f5f5f5;
  border-bottom: 1px solid #bababa;
  border-top: 1px solid #bababa;
  font-size: large;
}

.sticky-col {
  position: sticky;
  left: 0;
  z-index: 2;
  background: white;
  border-right: 5px solid #bababa;
  border-left: 1px solid #bababa;
}

.table .header-cell {
  border-bottom: 5px solid #bababa;
  background-color: whitesmoke;
  font-size: large;
  min-width: 150px;
  max-width: 150px;
}

.table .data-cell {
  text-align: right;
  padding-right: 5px;
  padding-top: 0;
  padding-bottom: 0;
  z-index: 0;
  font-size: 14px;
  line-height: 1.3;
  position: relative;
}

.header-cell-param {
  border-bottom: 1px solid #bababa;
  text-align: left;
  word-wrap: break-word;
  white-space: normal;
  padding: 0px 5px 0 2px;
  font-size: 14px;
  line-height: 1.3;
  max-width: 295px;
}

.engine-props-row {
  position: relative;
}

.engine-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #e46a61;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
}

.engine-props-row .engine-props-fill{
  background-color: #e46a61;
}

.physical-props-row {
  position: relative;
}

.physical-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #8dcbff;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
  padding: 0px 5px 0 2px;
}

.physical-props-row .physical-props-fill {
  background-color: #8dcbff;
}


.fuse-warhead-props-row {
  position: relative;
}

.fuse-warhead-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #94bb96;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
  padding: 0px 5px 0 2px;
}

.fuse-warhead-props-row .fuse-warhead-props-fill {
  background-color: #94bb96;
}


.guidance-props-row {
  position: relative;
}

.guidance-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #f1dc51;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
  padding: 0px 5px 0 2px;
}

.guidance-props-row .guidance-props-fill{
  background-color: #f1dc51;
}

.flight-props-row {
  position: relative;
}

.flight-props.sticky-col {
  display: block;
  text-align: left;
  background-color: #c2638f;
  font-size: large;
  font-weight: bold;
  border-bottom: 1px solid #bababa;
  padding: 0px 5px 0 2px;
}

.flight-props-row .flight-props-fill{
  background-color: #c2638f;
}

.add-notes {
  padding: 50px 2px;
  font-size: large;
  font-weight: bold;
  text-align: left;
}

.data-cell.add-notes-data {
  width: 150px;
  text-align: center;
}

.data-cell.text-center{
  text-align: center;
}


th:nth-child(even), td:nth-child(even) {
  background-color: #dbdbdb
}

.header-cell:nth-child(even) {
  background-color: #dbdbdb
}
</style>
