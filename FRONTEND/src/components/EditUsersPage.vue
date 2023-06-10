<template>
<div class="main">
    <div class="main-block">
        <img @click="$router.push('/')" src="../assets/img/Logotype.png">
        <div class="profile">
            <img src="../assets/img/ava.png">
            <p class="profile">{{ firstName }} {{ lastName }}</p>
        </div>
        <div class="modules">
            <div class="module" @click="$router.push({ name: 'dashboard'})">
                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-home" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                    <path d="M5 12l-2 0l9 -9l9 9l-2 0m-14 0v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7m-10 9v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                </svg>
                <p class = "module-name ">Главная</p>
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
        <div style="width: 50%; margin: 0 auto;">
            <v-text-field
                v-model="search"
                append-icon="mdi-magnify"
                label="Поиск по имени"
                single-line
                hide-details
            ></v-text-field>
            <v-data-table
                v-model:items-per-page="itemsPerPage"
                v-model:page="page"
                :headers="headers"
                :items="users"
                :search="search"
                item-value="lastName"
                @click:row="rowClick"
                single-select
                hide-default-footer
                :items-per-page="itemsPerPage"
                class="elevation-6"
            style="text-align: left">
                <template v-slot:bottom>
                    <div class="text-center pt-2">
                        <v-pagination
                            v-model="page"
                            :length="pageCount"
                        ></v-pagination>
                    </div>
                </template>
            </v-data-table>
        </div>
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
                            <span class="text-h5">Редактирование пользователя</span>
                            <hr>
                        </v-card-title>
                        <v-card-text>
                            <v-container class="lesson_block" style="font-size:15px;">
                                <v-row style="padding-top:15px;padding-left:15px; padding-bottom: 20px;">
                                    <v-col
                                        cols="12"
                                        sm="12"
                                        md="10"
                                        style="padding-bottom: 0px;"
                                    >
                                        <p class="lesson_user_info">Фамилия</p>
                                        <v-text-field v-model="test.lastName" density="compact" hide-details=true></v-text-field>
                                    </v-col>
                                    <v-col
                                        cols="12"
                                        sm="6"
                                        md="10"
                                        style="padding-bottom: 0px;"
                                    >
                                        <p class="lesson_user_info">Имя</p>
                                        <v-text-field v-model="test.firstName" density="compact" hide-details=true></v-text-field>
                                    </v-col>
                                    <v-col
                                        cols="12"
                                        sm="6"
                                        md="10"
                                        style="padding-bottom: 0px;"
                                    >
                                        <p class="lesson_user_info">Отчество</p>
                                        <v-text-field v-model="test.thirdName" density="compact" hide-details=true></v-text-field>
                                    </v-col>
                                    <v-col
                                        cols="12"
                                        sm="6"
                                        md="10"
                                        style="padding-bottom: 0px;"
                                    >
                                        <p class="lesson_user_info">Дата рождения</p>
                                        <v-text-field
                                            type="date"
                                            v-model="test.DOB"
                                            density="compact"
                                            hide-details=true
                                        ></v-text-field>
                                    </v-col>
                                    <v-col
                                        cols="12"
                                        sm="6"
                                        md="10"
                                        style="padding-bottom: 0px;"
                                    >
                                        <p class="lesson_user_info">Телефон</p>
                                        <v-text-field v-model="test.phone" density="compact" hide-details=true></v-text-field>
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
                                        <p class="lesson_user_info">Роль</p>
                                        <v-combobox
                                            label="Роль"
                                            :items="['developer', 'student', 'teacher']"
                                            v-model="test.role"
                                            density="compact"
                                        ></v-combobox>
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
</div>
</template>

<script>
import { VDataTable } from 'vuetify/labs/VDataTable'
import axios from "axios";
import config from '@/config/config'
export default {
    name: "EditUsersPage",
    components: {
        VDataTable,
    },
    data() {
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
            edit: false,
            test: this.$store.getters.loadData,
            // headers: ["Пользователь"],
            search: '',
            headers: [
                {
                    title: 'Фамилия',
                    align: 'start',
                    key: 'lastName',
                },
                {
                    title: 'Имя',
                    align: 'start',
                    key: 'firstName',
                },
                {
                    title: 'Отчество',
                    align: 'start',
                    key: 'thirdName',
                },
                {
                    title: 'Почта',
                    align: 'start',
                    key: 'email',
                }
            ],
            page: 1,
            itemsPerPage: 10,
            users: [
                {
                    lastName: "Пупкин",
                    firstName: "Василий",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Асафьев",
                    firstName: "Стас",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Стасов",
                    firstName: "Никита",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Козлов",
                    firstName: "Василий",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Матов",
                    firstName: "Максим",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Кружкин",
                    firstName: "Василий",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Дробин",
                    firstName: "Стас",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Норс",
                    firstName: "Никита",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Петрин",
                    firstName: "Василий",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Комаров",
                    firstName: "Пётр",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Грудинин",
                    firstName: "Иван",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Петрин",
                    firstName: "Алексей",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Петрин",
                    firstName: "Денис",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Петрин",
                    firstName: "Максим",
                    thirdName: "Иванович"
                },
                {
                    lastName: "Петрин",
                    firstName: "Василий",
                    thirdName: "Иванович"
                },
            ]
        }
    },
    mounted() {
        this.teeest()
    },
    computed: {
        pageCount () {
            return Math.ceil(this.users.length / this.itemsPerPage)
        },
    },
    methods: {
        rowClick: function (item, row) {
            console.log(row)
            this.edit = true
            this.test = row.item.props.title
            console.log(this.test)
            //item.name - selected id
        },
        teeest: function () {
            this.users.push(this.test)
            axios.post(config.AUTHSERVICE_URL+"/user/read")
                .then(response => {
                    console.log(response)
                })
        }
    }
}
</script>

<style scoped>
tr.v-data-table__selected {
    background: #7d92f5 !important;
}
.v-input__details {
    visibility: hidden;
}
</style>