<template>
    <div class="main">
        <div class="main-block">
            <img @click="$router.push('/')" src="../../assets/img/Logotype.png">
            <div class="profile">
                <img src="../../assets/img/ava.png">
                <p class="profile">{{ firstName }} {{ lastName }}</p>
            </div>
            <div class="modules">
                <div class="module" @click="$router.push({ name: 'dashboard'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-home" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M5 12l-2 0l9 -9l9 9l-2 0m-14 0v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7m-10 9v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                    </svg>
                    <p class = "module-name">Главная</p>
                </div>
                <div class="module" @click="$router.push({ name: 'modules'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-apps" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M4 4m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 9m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm10 -1m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 -8l6 0m-3 -3l0 6"></path>
                    </svg>
                    <p class = "module-name">Модули</p>
                </div>
                <div class="module" @click="$router.push({ name: 'help'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-bulb" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M3 12h1m8 -9v1m8 8h1m-15.4 -6.4l.7 .7m12.1 -.7l-.7 .7m-8.7 9.7a5 5 0 1 1 6 0a3.5 3.5 0 0 0 -1 3a2 2 0 0 1 -4 0a3.5 3.5 0 0 0 -1 -3m.7 1l4.6 0"></path>
                    </svg>
                    <p class = "module-name">Помощь</p>
                </div>
                <div class="module selected-module">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-settings" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M10.325 4.317c.426 -1.756 2.924 -1.756 3.35 0a1.724 1.724 0 0 0 2.573 1.066c1.543 -.94 3.31 .826 2.37 2.37a1.724 1.724 0 0 0 1.065 2.572c1.756 .426 1.756 2.924 0 3.35a1.724 1.724 0 0 0 -1.066 2.573c.94 1.543 -.826 3.31 -2.37 2.37a1.724 1.724 0 0 0 -2.572 1.065c-.426 1.756 -2.924 1.756 -3.35 0a1.724 1.724 0 0 0 -2.573 -1.066c-1.543 .94 -3.31 -.826 -2.37 -2.37a1.724 1.724 0 0 0 -1.065 -2.572c-1.756 -.426 -1.756 -2.924 0 -3.35a1.724 1.724 0 0 0 1.066 -2.573c-.94 -1.543 .826 -3.31 2.37 -2.37c1 .608 2.296 .07 2.572 -1.065zm1.675 7.683m-3 0a3 3 0 1 0 6 0a3 3 0 1 0 -6 0"></path>
                    </svg>
                    <p class = "module-name selected-module">Профиль</p>
                </div>
                <div class="logout">
                    <p class="" @click="logout">Выход</p>
                </div>
            </div>
        </div>
        <div class="alternative-block">
            <div class="some_settings">
                <div class="back_settings">
                    <div class="settings security">
                        <p style="font-size: 20px;">Данные аккаунта</p>
                        <v-dialog
                            v-model="edit"
                            width="auto"
                        >
                            <template v-slot:activator="{ props }">
                                <v-btn
                                    color="#83b1ff"
                                    v-bind="props"
                                    class="text-none"
                                    prepend-icon="mdi-email"
                                    style="color: white; width:90%; margin-bottom: 10px;"
                                >
                                    Смена email
                                </v-btn>
                            </template>
                            <v-card width="600">
                                <v-card-text style="text-align: center;">
                                    <v-icon icon="mdi-at" size="150px"></v-icon>
                                    <p class="Rubik-Medium" style="font-size:22px; text-align:center;">Смена электронной почты</p>
                                    <p style="text-align: center; padding:0 50px 0 50px;">Введите новый адрес электронной почты, на который будут приходить уведомления</p>
                                    <v-text-field :rules="[rules.required, rules.email]" label="Новый email" variant="solo" v-model="user.email"></v-text-field>
                                </v-card-text>
                                <v-card-actions style="text-align: center;">
                                    <v-btn class="text-none text-subtitle-1" color="#83b1ff" variant="flat" style="color:white;">Привязать почту</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-dialog>
                        <v-dialog
                            v-model="edit1"
                            width="auto"
                        >
                            <template v-slot:activator="{ props }">
                                <v-btn
                                    color="#83b1ff"
                                    v-bind="props"
                                    class="text-none"
                                    prepend-icon="mdi-phone"
                                    style="color: white; width:90%;"
                                >
                                    Смена телефона
                                </v-btn>
                            </template>
                            <v-card width="600">
                                <v-card-text style="text-align: center;">
                                    <v-icon icon="mdi-phone" size="150px"></v-icon>
                                    <p class="Rubik-Medium" style="font-size:22px; text-align:center;">Смена телефона</p>
                                    <p style="text-align: center; padding:0 50px 0 50px;">Введите новый номер телефона, который будет привязан к вашему аккаунту</p>
                                    <v-text-field v-mask="['+7 (###)-###-##-##']" type="tel" label="Телефон" variant="solo" v-model="user.phone"></v-text-field>
                                </v-card-text>
                                <v-card-actions style="text-align: center;">
                                    <v-btn class="text-none text-subtitle-1" color="#83b1ff" variant="flat" style="color:white;">Получить код</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-dialog>
                    </div>
                </div><br>
                <div class="back_settings">
                    <div class="settings security">
                        <p style="font-size: 20px;">Подтверждение входа</p>
                        <v-dialog
                            v-model="edit2"
                            width="auto"
                        >
                            <template v-slot:activator="{ props }">
                                <v-btn
                                    color="#83b1ff"
                                    v-bind="props"
                                    class="text-none text-left"
                                    prepend-icon="mdi-key-variant"
                                    style="color: white; margin-bottom:20px; width:90%;"

                                >
                                    Смена пароля
                                </v-btn>
                            </template>
                            <v-card width="600">
                                <v-card-text>
                                    <p style="font-size:22px; text-align:center;">Смена пароля</p>
                                    <v-text-field :rules="[rules.required]" label="Текущий пароль" variant="solo" clearable></v-text-field>
                                    <v-text-field :rules="[rules.required]" label="Новый пароль" variant="solo" clearable></v-text-field>
                                    <v-text-field  :rules="[rules.required]" label="Повторите новый пароль" variant="solo" clearable></v-text-field>
                                </v-card-text>
                                <v-card-actions style="text-align: center;">
                                    <v-btn class="text-none text-subtitle-1" color="#83b1ff" variant="flat" style="color:white;">СОХРАНИТЬ</v-btn>
                                </v-card-actions>
                            </v-card>
                        </v-dialog>
                    </div>
                </div>
            </div>


            <div class="navigation_settings">
                <div @click="$router.push('/settings')" class="nav_setting">Аккаунт и личная информация</div>
                <div class="nav_setting selected_nav_setting">Безопасность</div>
                <div @click="$router.push('/settings/notifications')" class="nav_setting">Уведомления</div>
            </div>

        </div>
    </div>

</template>

<script>
import {mask} from 'vue-the-mask'
export default {
    name: "SecurityPage",
    directives:{mask},
    data() {
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
            edit: false,
            edit1:false,
            edit2: false,
            user: {
                email: "",
            },
            rules: {
                required: value => !!value || 'Это поле обязательно',
                email: value => {
                    const pattern = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
                    return pattern.test(value) || 'Введите верный e-mail.'
                },
            },

        }
    },
    methods: {
        logout: function () {
            console.log(this.$store.state.user)
            this.$store.commit('logout')
            console.log(this.$store.state.user)
            this.$router.push('/auth')
        },
    }
}
</script>

<style scoped>
.some_settings {
    display:inline-block;
}
</style>