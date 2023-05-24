<template>

	<div style="position: relative;">

		<FullCalendar ref="fullCalendar" v-if="userRole=='developer'" id="elem1_calendar" :options="options" class="schedule-block"/>
		<template id="elem2">
			<v-row justify="center">
				<v-dialog
						v-model="edit"
						persistent
						width="600"
				>
					<template v-slot:activator="{ props }">
						<v-btn
								color="primary"
								v-bind="props"
						>
							Открыть
						</v-btn>
					</template>
					<v-card class="rounded-card">
						<v-card-title>
							<span class="text-h5">Страница занятия</span>
							<hr>
						</v-card-title>
						<v-card-text>
							<v-container class="lesson_block">
								<v-row style="padding-top:15px;padding-left:15px; padding-bottom: 20px;">
									<v-col
											cols="12"
											sm="12"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Название урока</p>
										<p type="text"> {{ eventTitle }} </p>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Тема урока</p>
										<p>{{eventTopic}}</p>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Домашнее задание</p>
										<p>{{eventDescription}}</p>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Дата</p>
										<p> {{eventDate}} с {{eventStart}} до {{eventEnd}}</p>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Класс</p>
										<p> {{eventClass}}</p>
									</v-col>
								</v-row>
							</v-container>
							<v-container class="lesson_block">
								<v-row style="padding-top:15px;padding-left:15px; padding-bottom: 20px;">
									<v-col
											cols="12"
											sm="12"
											md="10"
											style="padding-bottom: 0px;"
									>
										<p class="lesson_user_info">Учитель</p>
										<p type="text"> {{ eventTeacher }} </p>
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>
						<v-card-actions style="text-align: center">
							<v-spacer></v-spacer>
							<v-btn
									v-if="userRole==0"
									style="margin: 0 auto;"
									color="cyan-darken-2"
									variant="elevated"
									@click="edit = false"
							>
								Закрыть
							</v-btn>
							<v-btn
									v-else
									style="margin: 0 auto;"
									color="cyan-darken-2"
									variant="elevated"
									@click="edit = false"
							>
								Сохранить и закрыть
							</v-btn>
						</v-card-actions>
					</v-card>
				</v-dialog>
			</v-row>
		</template>
	</div>

	
</template>
<script>
// import '@fullcalendar/core/vdom'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin from '@fullcalendar/interaction'
import listPlugin from '@fullcalendar/list'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendar from '@fullcalendar/vue3'
// import axios from "axios";
import {INITIAL_EVENTS} from "@/components/modules/events";

export default {
  components: {
    FullCalendar, // make the <FullCalendar> tag available

  },
  data() {
    return {
			datetime: '',
			edit: false,
			userRole: this.$store.getters.loadData.role,
			eventId: '',
			eventTitle: '',
			eventDate: '',
			eventTopic: '',
			eventStart: '',
			eventEnd: '',
			eventBack: '',
			eventClass: '',
			eventColor: '#5ba8ff',
			eventDescription: '',
			eventTeacher: '',
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
				editable: false,
				selectable: true,
				aspectRatio: 1,
				// height: window.screen.height-300,
				contentHeight: 700,
				weekends: true,
				headerToolbar: {
					start: 'title', // will normally be on the left. if RTL, will be on the right
					center: '',
					end: 'timeGridWeek timeGridDay today prev,next' // will normally be on the right. if RTL, will be on the left
				},
				firstDay: '1',
				timeZone: 'UTC',
				events: INITIAL_EVENTS,
				eventDidMount: function(info) {
					let t = info.el
					t.style.wordWrap = "break-word"
					let d = t.querySelector('.fc-event-title-container')
					let teacher = info.event.extendedProps.teacher
					teacher = teacher.split(' ')
					let shortTeacher = teacher[0]+" "+teacher[1][0]+"."+teacher[2][0]+"."
					d.innerHTML += "<div class='fc-event-title fc-sticky'>" + shortTeacher + "</div>"
					let m = t.querySelector('.fc-event-title')
					m.style.fontWeight = "bold"
				},
				eventClick: (info) => {
					this.edit = true
					this.eventTitle= info.event.title
					this.eventDate = info.event.start.toISOString().split("T")[0]
					let tempStart = info.event.start.toISOString().split("T")[1].split(".")[0]
					this.eventStart= tempStart.substring(0,tempStart.length-3)
					let tempEnd = info.event.end.toISOString().split("T")[1].split(".")[0]
					this.eventBack = info.event.start.toISOString().split("T")[1].split(".")[1]
					this.eventEnd = tempEnd.substring(0, tempEnd.length-3)
					this.eventClass = info.event.extendedProps.class
					this.eventDescription = info.event.extendedProps.homework
					this.eventTeacher = info.event.extendedProps.teacher
					this.eventId = info.event.id
					this.eventTopic = info.event.extendedProps.topic
				},
				eventReceive: (info) => {
					// info.event.setProp('id', this.events.length)
					console.log(info.relatedEvents)

				},
				buttonText: {
					today:    'Сегодня',
					month:    'Месяц',
					week:     'Неделя',
					day:      'День',
					list:     'Список'
				},
				locale: 'ru',
				themeSystem: 'bootstrap5'
      }
    }
  },
  methods: {
		// getAllEvents: function () {
		// 	let events
		// 		axios.post('http://localhost:9909/' + '/schedule/read')
		// 				.then(resp => {
		// 					console.log(resp)
		// 					events = resp.data
		// 				})
		// 				.catch(err => {
		// 					console.log(err)
		// 				})
		// 	return events
		// }
  },
	mounted() {

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
#elem1_calendar {
	padding:25px;
	font-size:15px;
	height:80vh;
}

#elem2 {
		width: 100%;
		height: 100%;
		position: absolute;
		top: 0;
		left: 0;
		z-index: 10;
}
.v-row input {
	width: 100%;
}
.v-card-title {
	text-align:center;
}
.lesson_user_info {
	font-size:16px;
	margin-bottom:0px;
	font-family: "Rubik Medium";
}
hr {
	border: none;
	height: 1px;
	/* Set the hr color */
	color: #333;  /* old IE */
	background-color: #333;
}
.lesson_block {
	padding-top: 0px;
	padding-left: 0px;
	padding-right: 0px;
	border-radius: 15px;
	background-color: #f7fbff;
	border-color: #cdcdcd;
	border-width: 1px;
	border-style:solid;
}
</style>