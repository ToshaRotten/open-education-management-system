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

	<div style="position: relative;">
		<FullCalendar ref="fullCalendar" id="elem1" :options="options"/>

		<template id="elem2">
			<v-row justify="center">
				<v-dialog
						v-model="edit"
						persistent
						width="600"
						height="650"
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
							<span class="text-h5">Редактирование урока</span>
						</v-card-title>
						<v-card-text>
							<v-container>
								<v-row>
									<v-col
											cols="12"
											sm="12"
											md="10"
									>
										<p class="lesson_info">Название урока</p>
										<input
												type="text"
												v-model="eventTitle"
												placeholder="Введите название урока"
												required
										>
										<select style="width:100%;" v-model="eventTitle">
											<option v-for="lesson in lessons" :key="lesson"> {{lesson.title}}</option>
										</select>
									</v-col>

									<v-col
											cols="12"
											sm="6"
											md="10"
									>
										<p class="lesson_info">Описание урока</p>
										<input
												v-model="eventDescription"
												placeholder="Здесь находится описание занятия"
										>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
									>
										<p class="lesson_info">Начало урока</p>
										<input
												type="datetime-local"
												v-model="eventStart">
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
									>
										<p class="lesson_info">Конец урока</p>
										<input
												type="datetime-local"
												v-model="eventEnd">
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="10"
									>
										<p class="lesson_info">Выбор цвета</p>
										<input
												type="color"
												v-model="eventColor">
									</v-col>
								</v-row>
							</v-container>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn
									color="blue-darken-1"
									variant="text"
									@click="edit = false"
							>
								Закрыть
							</v-btn>
							<v-btn
									color="blue-darken-1"
									variant="text"
									@click="changeEventData"
							>
								Сохранить
							</v-btn>
							<v-btn
									color="red"
									@click="deleteEvent"
							>
								Удалить
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
import interactionPlugin, { Draggable } from '@fullcalendar/interaction'
import listPlugin from '@fullcalendar/list'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendar from '@fullcalendar/vue3'
import {INITIAL_EVENTS} from "./events";

export default {
  components: {
    FullCalendar, // make the <FullCalendar> tag available

  },
  data() {
    return {
			datetime: '',
			lessons: [
				{
					"title": "Пустой урок",
					"duration": "00:45:00",
					"backgroundColor": "#ffc14d"
				},
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
			edit: false,
			eventId: '',
			eventTitle: '',
			eventStart: '',
			eventEnd: '',
			eventColor: '#5ba8ff',
			eventDescription: '',
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
				aspectRatio: 1,
				weekends: true,
				firstDay: '1',
				timeZone: 'UTC',
				events: INITIAL_EVENTS,
				eventClick: (info) => {
					this.edit = true
					this.eventTitle= info.event.title
					this.eventStart= info.event.start.toISOString().split("Z")[0]
					this.eventEnd= info.event.end.toISOString().split("Z")[0]
					this.eventDescription = info.event.description
					this.eventId = info.event.id
					console.log(this.eventId)
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
				height: 700,
				locale: 'ru',
				themeSystem: 'bootstrap5'
      }
    }
  },
  methods: {
		changeEventData: function () {
			let calendarApi = this.$refs.fullCalendar.getApi()
			let event = calendarApi.getEventById(this.eventId)
			event.setProp('title', this.eventTitle)
			event.setDates(this.eventStart, this.eventEnd)
			event.setProp('backgroundColor', this.eventColor)
			this.edit=false
		},
		deleteEvent: function () {
			let calendarApi = this.$refs.fullCalendar.getApi()
			let event = calendarApi.getEventById(this.eventId)
			event.remove()
			this.edit=false
		},
  },
	created: function() {
		let array = this.lessons
		document.addEventListener('DOMContentLoaded', function() {
		let draggableEl = document.getElementsByClassName("drag");
		for(let i = 0; i < array.length; i++) {
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
#elem1,
#elem2 {
		width: 100%;
		height: 100%;
		position: absolute;
		top: 0;
		left: 0;
}
#elem2 {
	z-index: 10;
}
.v-row input {
	width: 100%;
}
.v-card-title {
	text-align:center;
}
.lesson_info {
	font-size:12px;
	margin-bottom:0px;
	margin-left:5px;
}

</style>