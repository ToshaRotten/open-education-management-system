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
		<FullCalendar id="elem1" :options="options"/>
		<template id="elem2">
			<v-row justify="center">
				<v-dialog
						v-model="edit"
						persistent
						width="600"
						height="512"
				>
					<template v-slot:activator="{ props }">
						<v-btn
								color="primary"
								v-bind="props"
						>
							Open Dialog
						</v-btn>
					</template>
					<v-card>
						<v-card-title>
							<span class="text-h5">Редактирование урока</span>
						</v-card-title>
						<v-card-text>
							<v-container>
								<v-row>
									<v-col
											cols="12"
											sm="6"
											md="4"
									>
										<v-text-field
												label="Название"
												required
										></v-text-field>
									</v-col>
									<v-col
											cols="12"
											sm="6"
											md="4"
									>
										<v-text-field
												label="Описание"
												hint="example of helper text only on focus"
										></v-text-field>
									</v-col>
									<v-col
											cols="12"
											sm="6"
									>
										<v-datetime-picker label="Select Datetime" v-model="datetime"> </v-datetime-picker>
									</v-col>
								</v-row>
							</v-container>
							<small>*indicates required field</small>
						</v-card-text>
						<v-card-actions>
							<v-spacer></v-spacer>
							<v-btn
									color="blue-darken-1"
									variant="text"
									@click="edit = false"
							>
								Close
							</v-btn>
							<v-btn
									color="blue-darken-1"
									variant="text"
									@click="edit = false"
							>
								Сохранить
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
				events: INITIAL_EVENTS,
				eventClick: (info) => {
					console.log(this.edit)
					this.edit = true
					console.log(this.edit)
					console.log('View: ' + info.event.title)
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
    handleDateClick: function(arg) {
      console.log('date click! ' + arg.dateStr)
    },
		showModal() {
			return this.edit = !this.edit
		}
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
</style>