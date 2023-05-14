<template>
  <div class="main">
    <div class="main-block">
        <img @click="$router.push('/')" src="../assets/img/Logotype.png">
          <div class="profile">
              <img src="../assets/img/ava.png">
              <p class="profile">{{ firstName }} {{ lastName }}</p>
          </div>
          <div class="modules">
              <div class="module" @click="$router.push('/dashboard')">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-home" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                      <path d="M5 12l-2 0l9 -9l9 9l-2 0m-14 0v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7m-10 9v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                   </svg>
                  <p class = "module-name">Главная</p>
              </div>
              <div class="module selected-module">
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
        <div class='chooseText'>
            <p>Выберите доступный модуль</p>
        </div>
        <div>
            <button class='module_list'
                    v-for='module in modules'
                    :key='module'
                    @click="$router.push(module.path)" style="cursor:pointer;">
               {{ module.moduleName }}
            </button>
        </div>
      </div>
  </div>
</template>
<script>

export default {
    name: 'ModulesPage',
    data() {
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
            role: this.$store.getters.loadData.role,
            modules: [
                {
                    moduleId: '1',
                    moduleName: 'Расписание',
                    path: '/modules/generate',
                    role: '0'
                },
                {
                    moduleId: '2',
                    moduleName: 'Дневник',
                    path: '/modules/journal',
                    role: '0'
                },
                {
                    moduleId: '3',
                    moduleName: 'Онлайн встреча',
                    path: '/modules/meet',
                    role: '0'
                },
                {
                    moduleId: '4',
                    moduleName: 'Что это?',
                    path: '/modules/journal',
                    role: '0'
                },
            ]
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
<style>
.module_list {
    border-radius: 30px;
    border-width: 0px;
    border-color: black;
    filter: drop-shadow(0px 0px 3px black);
    font-size: 20px;
    background-color: #FFFFFF;
    width: 15em;
    height: 3em;
    margin-inline: 25px;
    margin-bottom: 2em;
}
.chooseText {
    height:100%;
    width:100%;
    margin-bottom: 2em;
    color: black;
    font-family: 'TT Commons';
    font-size:24px;
}

</style>