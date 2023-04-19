<template>
	<div>
		<p>Список уроков</p>
		<div class='lessons_list'>
			<div class='drag' v-for="lesson in lessons" :key='lesson'>
				{{ lesson.title}}
			</div>
			<div class='drag'></div>
		</div>
	</div>
	<FullCalendar :options="options"/>
	<EditLesson/>
	
</template>
<script>
// import '@fullcalendar/core/vdom'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin, { Draggable } from '@fullcalendar/interaction'
import listPlugin from '@fullcalendar/list'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendar from '@fullcalendar/vue3'
import EditLesson from './EditLesson.vue'

export default {
  components: {
    FullCalendar, // make the <FullCalendar> tag available
		EditLesson
  },
  data() {
    return {
			lessons: [
				{ 
					"title": "Русский язык", 
					"duration": "00:45:00", 
					"backgroundColor": "#FF6B6B" 
				},
				{ 
					"title": "Математика", 
					"duration": "00:45:00", 
					"backgroundColor": "#3FBAFF" 
				}
			],
      options: {
        plugins: [dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin],
				initialView: 'timeGridWeek',
				slotLabelFormat: {
					hour: 'numeric',
					minute: '2-digit',
					omitZeroMinute: false,
					meridiem: 'short'
				},
				slotMinTime: "07:00:00",
				slotMaxTime: "19:00:00",
				slotDuration: "00:15:00",
				allDaySlot:false,
				editable: true,
				selectable: true,
				dateClick: this.handleDateClick,
				weekends: true,
				firstDay: '1',
				events: [
					{
						id: 0,
						title: 'Русский язык',
						start: '2023-04-15T10:30:00',
						end: '2023-04-15T12:30:00'
					},
					{
						id: 1,
						title: 'Литература',
						start: '2023-04-19T10:00:00',
						end: '2023-04-19T10:45:00'
					},
				],
				// eventDrop: function(info) {
				// 	alert(info.event.title + " was dropped on " + info.event.start.toISOString() + "and old event is" + info.oldEvent.start.toISOString());

				// 	if (!confirm("Are you sure about this change?")) {
				// 		info.revert();
				// 	}
				// },
				buttonText: {
					today:    'Сегодня',
					month:    'Месяц',
					week:     'Неделя',
					day:      'День',
					list:     'Список'
				},
				height: 700,
				locale: 'ru',
				themeSystem: 'bootstrap4'
      }
    }
  },
  methods: {
    handleDateClick: function(arg) {
      console.log('date click! ' + arg.dateStr)
    }
  },
	created: function() {
		let array = this.lessons
		document.addEventListener('DOMContentLoaded', function() {
		let draggableEl = document.getElementsByClassName("drag");
		for(let i = 0; i < array.length; i++) {
			console.log(array[i])
			new Draggable(draggableEl[i], {
				eventData: array[i]
			});
		}
	});
	}
}
</script>

<style>
.drag {
	-webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Old versions of Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, 		currently
                                  supported by Chrome, Edge, Opera and Firefox */
	background-color: #FF6B6B;
	color: white;
	width: 10vw;
	text-align: center;
	border-radius: 10px;
	border-color: white;
}

</style>