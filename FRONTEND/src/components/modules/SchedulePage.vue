<template>
	<FullCalendar v-bind:options="options"/>
</template>
<script setup>
// import '@fullcalendar/core/vdom'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from '@fullcalendar/interaction'
import listPlugin from '@fullcalendar/list'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendar from '@fullcalendar/vue3'
import { reactive } from 'vue'

const options = reactive({
	plugins: [dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin],
	initialView: 'dayGridMonth',
	editable: true,
	selectable: true,
	weekends: true
	// select: (arg) => {
		
	// }
})

</script>