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
                    <p class = "module-name">Главная</p>
                </div>
                <div class="module" @click="$router.push({ name: 'modules'})">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-apps" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M4 4m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 9m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm10 -1m0 1a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v4a1 1 0 0 1 -1 1h-4a1 1 0 0 1 -1 -1zm0 -8l6 0m-3 -3l0 6"></path>
                    </svg>
                    <p class = "module-name">Модули</p>
                </div>
                <div class="module selected-module">
                    <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-bulb" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                        <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                        <path d="M3 12h1m8 -9v1m8 8h1m-15.4 -6.4l.7 .7m12.1 -.7l-.7 .7m-8.7 9.7a5 5 0 1 1 6 0a3.5 3.5 0 0 0 -1 3a2 2 0 0 1 -4 0a3.5 3.5 0 0 0 -1 -3m.7 1l4.6 0"></path>
                    </svg>
                    <p class = "module-name selected-module">Помощь</p>
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
<!--            Форму тикета, контакты и гугл мапсы-->
            <div class="back_help">
                <div class="help">
                    <v-card-title>
                        <span class="text-h5">Оставить заявку об ошибке</span>
                    </v-card-title>
                    <v-textarea label="Имя" variant="solo" :model-value="firstName" auto-grow rows="1"></v-textarea>
                    <v-textarea rows="1" label="Фамилия" variant="solo" :model-value="lastName" auto-grow></v-textarea>
                    <v-textarea no-resize label="Описание проблемы" rows="3" variant="solo"></v-textarea>
                    <div class="help_button">
                        Отправить
                    </div>
                </div>
            </div>
            <div class="contacts">
                <div class="contacts_info">
                    <p class="Rubik-Regular" style="font-size: 25px; color: black;">Наши контакты</p>
                    <p class="Rubik-Medium" style="font-size: 40px; color: black;">8 800 111-11-11</p>
                    <p class="Rubik-Regular" style="font-size: 25px; color: black;">help@romanton.ru</p>
                    <p class="Rubik" style="color: black;">По общим вопросам (в будни с<br />7:00 до 18:00 по московскому времени)</p>
                </div>

                <div class="google_maps">

                </div>
            </div>


        </div>
    </div>

</template>

<script>
export default {
    name: "HelpPage",
    data(){
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
        }
    }
}
</script>

<style>
.back_help {
    background-color: white;
    width: 35em;
    border-radius: 7px;
    display: inline-flex;
    filter: drop-shadow(0px 1px 4px #3c3c3c);
    padding-top:15px;
    padding-bottom: 50px;
}
.help {
    text-align: center;
    margin: 0 auto;
    width:25em;
}
.v-input {
    display:inline !important;
}
.v-field__input {
    border-style:none;
}
.v-text-field input {
    padding-left: 10px;
}
.help_button {
    border-style: solid;
    border-width: 1px;
    width:15em;
    display:inline-flex;
    border-radius: 15px;
    height: 2em;
    justify-content: center;
    align-items: center;
    cursor: pointer;
}
.contacts {
    display: inline-flex;
    margin-left: 50px;
    justify-content: center;
    margin-top:50px;
}
.contacts_info {
    border-style: solid;
    border-width: 1px;
    padding: 10px 50px 10px 50px;
}
.google_maps {
    display: inline-flex;
}
</style>