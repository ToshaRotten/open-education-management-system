<template>
<!--    <header>-->
<!--        <div class="navigation">-->
<!--            <div class="third-nav">-->
<!--                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-door-exit" width="24" height="24" viewBox="0 0 24 24" stroke-width="1" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">-->
<!--                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>-->
<!--                    <path d="M13 12v.01m-10 8.99h18m-16 0v-16a2 2 0 0 1 2 -2h7.5m2.5 10.5v7.5m-3 -14h7m-3 -3l3 3l-3 3"></path>-->
<!--                </svg>-->
<!--                <a class="" @click="$router.push('/auth')">Выйти</a>-->
<!--            </div>-->
<!--            <div class="second-nav">-->
<!--                <p class="" style="font-size:26px;">Расписание</p>-->
<!--            </div>-->
<!--        </div>-->
<!--    </header>-->

    <div class="main">
        <div class="main-block">
            <img @click="$router.push('/')" src="../../assets/img/Logotype.png">
            <div class="profile">
                <img src="../../assets/img/ava.png">
                <p class="profile">{{ firstName }} {{ lastName }}</p>
                <p style="font-size:17px;">{{ role }}</p>
            </div>
            <div class="modules">
                <div class="module" @click="$router.push({ name: 'dashboard'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-home" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M5 12l-2 0l9 -9l9 9l-2 0m-14 0v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7m-10 9v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                    </svg>
                    <p class = "module-name">Главная</p>
                </div>
                <div class="module selected-module" @click="$router.push({ name: 'modules'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-apps" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M4 4m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 9m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm10 -1m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 -8l6 0m-3 -3l0 6"></path>
                    </svg>
                    <p class = "module-name selected-module">Модули</p>
                </div>
                <div class="module" @click="$router.push({ name: 'help'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-bulb" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M3 12h1m8 -9v1m8 8h1m-15.4 -6.4l.7 .7m12.1 -.7l-.7 .7m-8.7 9.7a5 5 0 1 1 6 0a3.5 3.5 0 0 0 -1 3a2 2 0 0 1 -4 0a3.5 3.5 0 0 0 -1 -3m.7 1l4.6 0"></path>
                    </svg>
                    <p class = "module-name">Помощь</p>
                </div>
                <div class="module" @click="$router.push({ name: 'settings'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-settings" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065zm1.675 7.683m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0"></path>
                    </svg>
                    <p class = "module-name">Профиль</p>
                </div>
                <div class="logout">
                    <p class="" @click="logout">Выход</p>
                </div>
            </div>
        </div>
        <div class="alternative-block">
            <div>
                <div class='lessons_list'>

                    <div class='drag' v-for="lesson in lessons" :key='lesson'>
                        <v-icon icon="mdi-drag-variant"></v-icon>
                            {{ lesson.title}}
                    </div>
                </div>
            </div>
            <div style="position: relative;margin-top:10px;">
                <FullCalendar ref="fullCalendar" id="elem1" :options="options" class="schedule-block"/>
                <template id="elem2">
                    <v-row justify="center">
                        <v-dialog
                            v-model="edit"
                            persistent
                            width="600"
                            height="1000"
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
                                    <hr>
                                </v-card-title>
                                <v-card-text>
                                    <v-container>
                                        <v-row class="vrow">
                                            <v-col
                                                cols="12"
                                                sm="12"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Название урока</p>
<!--                                                <input-->
<!--                                                    type="text"-->
<!--                                                    v-model="eventTitle"-->
<!--                                                    placeholder="Введите название урока"-->
<!--                                                    required-->
<!--                                                >-->
                                                <select style="width:100%;" v-model="eventTitle">
                                                    <option v-for="lesson in lessons" :key="lesson"> {{lesson.title}}</option>
                                                </select>
                                            </v-col>
                                            <v-col
                                                cols="12"
                                                sm="6"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Тема урока</p>
                                                <input
                                                    v-model="eventTopic"
                                                    placeholder="Тема урока"
                                                >
                                            </v-col>
                                            <v-col
                                                cols="12"
                                                sm="6"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Описание урока</p>
                                                <input
                                                    v-model="eventDescription"
                                                    placeholder="Здесь находится описание занятия"
                                                >
                                            </v-col>
                                            <v-col
                                                cols="12"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Начало урока</p>
                                                <input
                                                    type="date"
                                                    v-model="eventDate">
                                                <div style="margin: 0 auto; justify-content: center; display:flex; text-align:center;">
                                                    <div
                                                        style="width:100%; margin-right:2em;"
                                                        class="float-left"
                                                    >
                                                        <p style="margin-bottom: 0px;">с</p>
                                                        <input
                                                            type="time"
                                                            v-model="eventStart"
                                                        >
                                                    </div>
                                                    <div

                                                        class="float-left"
                                                        style="width:100%;"
                                                    >
                                                        <p style="margin-bottom: 0px;">до</p>
                                                        <input
                                                            type="time"
                                                            v-model="eventEnd">
                                                    </div>
                                                </div>

                                            </v-col>
                                            <v-col
                                                cols="12"
                                                sm="6"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Выбор педагога</p>
                                                <select style="width:100%;" v-model="eventTeacher">
                                                    <option v-for="teacher in teachers" :key="teacher"> {{teacher.name}}</option>
                                                </select>
                                            </v-col>
                                            <v-col
                                                cols="12"
                                                sm="6"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Выбор класса</p>
                                                <select style="width:100%;" v-model="eventClass">
                                                    <option v-for="clas in classes" :key="clas"> {{clas}}</option>
                                                </select>
                                            </v-col>

                                            <v-col
                                                cols="12"
                                                sm="6"
                                                md="10"
                                                style="padding-bottom: 0px;"
                                            >
                                                <p class="lesson_info">Выбор цвета</p>
                                                <input
                                                    class="color_picker"
                                                    type="color"
                                                    v-model="eventColor"
                                                >
                                            </v-col>
                                        </v-row>
                                    </v-container>
                                </v-card-text>
                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn
                                        color="transparent"
                                        variant="elevated"
                                        class="float-right"
                                        @click="edit = false"
                                    >
                                        Закрыть
                                    </v-btn>
                                    <v-btn
                                        color="teal-lighten-2"
                                        variant="elevated"
                                        class="float-left"
                                        @click="changeEventData"
                                    >
                                        Сохранить
                                    </v-btn>
                                    <v-btn
                                        color="red-accent-2"
                                        variant="outlined"
                                        @click="deleteEvent"
                                        class=""
                                    >
                                        Удалить
                                    </v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-dialog>
                    </v-row>
                </template>
            </div>
        </div>
    </div>





</template>
<script>
// import '@fullcalendar/core/vdom'
import dayGridPlugin from '@fullcalendar/daygrid'
import interactionPlugin, { Draggable } from '@fullcalendar/interaction'
import listPlugin from '@fullcalendar/list'
import timeGridPlugin from '@fullcalendar/timegrid'
import FullCalendar from '@fullcalendar/vue3'
import { useToast } from "vue-toastification"
import {INITIAL_EVENTS} from "@/components/modules/events";

export default {
    components: {
        FullCalendar, // make the <FullCalendar> tag available

    },

    data() {
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
            datetime: '',
            lessons: [
                {
                    "title": "Пустой урок",
                    "duration": "00:45:00",
                    "backgroundColor": "#de4dff",
                    "extendedProps": {
                        "description": "",
                        "teacher": this.firstName + this.lastName
                    }
                },
                {
                    "title": "Русский язык",
                    "duration": "00:45:00",
                    "backgroundColor": "#FF6B6B",
                    "extendedProps": {
                        "description": "",
                        "teacher": this.firstName + this.lastName
                    }
                },
                {
                    "title": "Математика",
                    "duration": "00:45:00",
                    "backgroundColor": "#3FBAFF",
                    "extendedProps": {
                        "description": "",
                        "teacher": ""
                    }
                },
                {
                    "title": "Информатика",
                    "duration": "00:45:00",
                    "backgroundColor": "#653fff",
                    "extendedProps": {
                        "description": "",
                        "teacher": ""
                    }
                },
                {
                    "title": "Астрономия",
                    "duration": "00:45:00",
                    "backgroundColor": "#ff8c3f",
                    "extendedProps": {
                        "description": "",
                        "teacher": ""
                    }
                }
            ],
            teachers: [
                {name:"Баринов Виктор Петрович"},
                {name:"Семёнов Павел Пальчикович"},
                {name:"Павлов Павел Павлович"},
                {name:"Козлов Антон Дмитриевич"},
                {name:"Свиридонов Роман Александрович"},
                {name:"Романов Антон Дмитриевич"},
                {name:"Антонов Роман Александрович"},
            ],
            classes: ["1А","1Б","1В","1Г",
                "2А","2Б","2В","2Г",
                "3А","3Б","3В","3Г",
                "4А","4Б","4В","4Г",
                "5А","5Б","5В","5Г",
                "6А","6Б","6В","6Г",
                "7А","7Б","7В","7Г",
                "8А","8Б","8В","8Г",
                "9А","9Б","9В","9Г",
                "10А","10Б","10В","10Г",
                "11А","11Б","11В","11Г"],
            edit: false,
            eventId: '',
            eventTitle: '',
            eventTeacher: '',
            eventDate: '',
            eventStart: '',
            eventEnd: '',
            eventBack: '',
            eventClass: '',
            eventTopic: '',
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
                    this.eventDate = info.event.start.toISOString().split("T")[0]
                    let tempStart = info.event.start.toISOString().split("T")[1].split(".")[0]
                    this.eventStart= tempStart.substring(0,tempStart.length-3)
                    let tempEnd = info.event.end.toISOString().split("T")[1].split(".")[0]
                    this.eventBack = info.event.start.toISOString().split("T")[1].split(".")[1]
                    this.eventEnd = tempEnd.substring(0, tempEnd.length-3)
                    this.eventClass = info.event.extendedProps.class
                    this.eventTeacher = info.event.extendedProps.teacher
                    this.eventDescription = info.event.extendedProps.homework
                    this.eventId = info.event.id
                    this.eventTopic = info.event.extendedProps.topic
                    console.log(this.eventId)
                },
                eventReceive: (info) => {
                    let calendar = this.$refs.fullCalendar.getApi()
                    let events = calendar.getEvents()
                    info.event.setProp('id', events.length)
                    console.log(events.length)
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
    setup() {
        const toast = useToast();
        return { toast }
    },
    watch: {
        $route() {
        }
    },
    methods: {
        logout: function () {
            console.log(this.$store.state.user)
            this.$store.commit('logout')
            console.log(this.$store.state.user)
            this.$router.push('/auth')
        },
        changeEventData: function () {
            let calendarApi = this.$refs.fullCalendar.getApi()
            let event = calendarApi.getEventById(this.eventId)
            event.setProp('title', this.eventTitle)
            let startDate = this.eventDate + "T" + this.eventStart + ":00" + "." + this.eventBack
            let endDate = this.eventDate + "T" + this.eventEnd + ":00" + "." + this.eventBack
            event.setDates(startDate, endDate)
            event.setProp('backgroundColor', this.eventColor)
            event.setExtendedProp('description', this.eventDescription)
            event.setExtendedProp('teacher', this.eventTeacher)
            event.setExtendedProp('class', this.eventClass)
            event.setExtendedProp('topic', this.eventTopic)
            this.edit=false
            this.toast.success("Урок успешно изменён!")
            this.updateEvents()
        },
        deleteEvent: function() {
            let calendarApi = this.$refs.fullCalendar.getApi()
            let event = calendarApi.getEventById(this.eventId)
            event.remove()
            this.edit=false
            this.toast.success("Урок удалён!")
        },
        makeDraggable: function () {
            let array = this.lessons
            let draggableEl = document.getElementsByClassName("drag");
            for(let i = 0; i < array.length; i++) {
                new Draggable(draggableEl[i], {
                    eventData: array[i]
                });
            }
        },
        // getAllEvents: function () {
        //     let events
        //     axios.post('http://185.21.142.92:9909' + '/schedule/read')
        //         .then(resp => {
        //             console.log(resp)
        //             events = resp.data
        //         })
        //         .catch(err => {
        //             console.log(err)
        //         })
        //     return events
        // },
        // updateEvents: function () {
        //     let calendarApi = this.$refs.fullCalendar.getApi()
        //     let events = calendarApi.getEvents()
        //     axios.post('http://185.21.142.92:9909' +'/schedule/create', events)
        //         .then(resp => {
        //             console.log(resp)
        //         })
        //         .catch(err => {
        //             console.log(err)
        //         })
        // }
    },
    mounted() {
        this.makeDraggable()
        // this.getAllEvents()
    },
    created() {

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
    background-color: #b864d9;
    color: white;
    width: 10vw;
    height:5vh;
    text-align: center;
    vertical-align: middle;
    line-height: 5vh;
    border-radius: 10px;
    border-color: white;
}
#elem1,
#elem2 {
    width: 100%;
    height: 100%;
    font-size: 16px;
    position: absolute;
    top: 0;
    left: 0;
}
#elem2 {
    z-index: 10;
}
.v-row input {
    width: 100%;
    padding-left:15px;
    padding-right:15px;
    box-sizing: border-box;
    border-width: 1px;
    border-radius: 15px;
    height:2em;
}
.v-row select {
    padding-left: 15px;
    box-sizing: border-box;
    border-radius: 15px;
    height:2em;
}
.vrow {
    justify-content: center;
}
.v-container {
    padding:0px;
}
.v-card-title {
    text-align:center;
}
.lesson_info {
    font-size:12px;
    margin-bottom:0px;
    margin-left:5px;
}
.lessons_list {
    margin: 0 auto;
    display: inline-block;
    border-radius: 15px;
    background-color: #ffffff;
    filter: drop-shadow(0px 0px 3px #9e9eff);
}
.drag {
    display:inline-table;
    margin: 15px;

}
input[type="date"] {
    border-width: 1px;
}
.color_picker {
    border:none;
    cursor: pointer;
    background-color: transparent;
    padding-left: 0px;
    padding-right: 0px;
    height:30px;

}
.color_picker::-webkit-color-swatch {
    border-radius: 15px;
    border:none;
}
.color_picker::-moz-color-swatch {
    border-radius: 15px;
    border:none;
}
button, input, select, textarea {
    border-style: solid;
}
.v-card-actions {
    display:inline-block;
    padding-left: 50px;
    padding-right: 50px;
}
</style>