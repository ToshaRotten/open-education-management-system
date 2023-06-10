<template>
    <router-view/>
  <div class="main">
      <div class="main-block">
          <img @click="$router.push('/')" src="../assets/img/Logotype.png">
          <div class="profile">
              <img src="../assets/img/ava.png">
              <p class="profile">{{ firstName }} {{ lastName }}</p>
          </div>
          <div class="modules">
              <div class="module selected-module">
                  <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-home" width="23" height="23" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                      <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                      <path d="M5 12l-2 0l9 -9l9 9l-2 0m-14 0v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-7m-10 9v-6a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v6"></path>
                   </svg>
                  <p class = "module-name selected-module">Главная</p>
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
            <div style="display:flex; float:left;">
                <SchedulePage/>
            </div>
            <div class="last-marks">
                <p class="Rubik-Regular" style="font-size: 22px;">Последние оценки</p> <hr>
                <div class="grade" :id = "grade.grade" v-for="grade in lastGrades" :key="grade">
                    <p style="font-size: 12px;"> {{ grade.lesson }}</p>
                    <p style="font-size: 25px;"> {{ grade.grade }}</p>
                    <p style="font-size: 12px;"> {{ grade.date }}</p>
                </div>
            </div>
            <div class="news-block">
                <p align="center" class="Rubik-Regular" style="font-size: 22px;">Последние новости</p><hr>
                <div class="list-news">
                    <div class="news" v-for="news in lastNews" :key="news">
                        <p style="margin:0;"> {{ news.date }} </p>
                        <p style="font-weight: bold; margin:0;"> {{ news.name }} </p>
                        <p style="font-size:14px;"> {{ news.description }} </p>
                    </div>
                </div>

            </div>
        </div>

  </div>
  
</template>

<script>
import SchedulePage from './modules/SchedulePage.vue'

export default {
  name: 'MainPage',
  data () {
      return {
          firstName: this.$store.getters.loadData.firstName,
          lastName: this.$store.getters.loadData.lastName,
          lastGrades: [],
          lastNews: []
      }
  },
  components: {
    'SchedulePage': SchedulePage
  },
  created() {
      this.getGrades()
      this.getLastNews()
  },
  methods: {
    logout: function () {
        console.log(this.$store.state.user)
        this.$store.commit('logout')
        console.log(this.$store.state.user)
        this.$router.push('/auth')
    },
      getGrades: function () {
          let grades = [
              {
                  date: '24.05.2023',
                  lesson: 'Русский язык',
                  grade: '5'
              },
              {
                  date: '22.05.2023',
                  lesson: 'Математика',
                  grade: '2'
              },
              {
                  date: '22.05.2023',
                  lesson: 'Математика',
                  grade: '4'
              },
              {
                  date: '22.05.2023',
                  lesson: 'Математика',
                  grade: '3'
              }
          ]
          this.lastGrades = grades
      },
      getLastNews: function () {
          let lastNews = [
              {
                  date: '24.05.2023',
                  name: 'Новый чайник',
                  description: 'Корабли лавировали, лавировали, да не вылавировали.'
              },
              {
                  date: '22.05.2023',
                  name: 'Новый трамвай',
                  description: 'У четырех черепашек четыре черепашонка.У четырех черепашек четыре черепашонка.У четырех черепашек четыре черепашонка.У четырех черепашек четыре черепашонка.'
              },
              {
                  date: '22.05.2023',
                  name: 'Новый учебник',
                  description: 'Купили учебник'
              },
              {
                  date: '22.05.2023',
                  name: 'Новый учитель',
                  description: 'Купили учителя'
              }
          ]
          this.lastNews = lastNews
      },
      colorizeGrads:function () {
          let grads = document.getElementsByClassName('grade')
          for(let i = 0; i < grads.length; i++) {
              if (grads[i].id == '5') grads[i].classList.add('mG')
              else if (grads[i].id == '4') grads[i].classList.add('mLG')
              else if (grads[i].id == '3') grads[i].classList.add('mY')
              else if (grads[i].id == '2') grads[i].classList.add('mR')
              else grads[i].classList.add('mR')
          }
      }
  },
    mounted() {
      this.colorizeGrads();
      console.log(this.$store.state.user)
    }
}
</script>
<style>
#elem1 {
    padding:25px;
}
.schedule-block {
    background-color: white;
    border-radius: 15px;
  box-shadow: 0 2px 4px black;
    width:50vw;
}
.last-marks {
    display:inline-block;
    border-style: solid;
    border-radius: 15px;
    width: 30%;
    border-width: 0px;
    padding:15px;
    background-color: white;
    filter: drop-shadow(0px 1px 4px #3c3c3c);
}

.grade {
    border-radius: 15px;
    border-style: solid;
    background-color: #0A8CFF;
    color: white;
    width:20%;
    display: inline-block;
}
.grade p {
    margin: 0;
}
.news-block {
    display:inline-block;
    border-style: solid;
    border-radius: 15px;
    width: 30%;
    border-width: 0px;
    padding:15px;
    margin-top: 50px;

    height: 50vh;
    background-color: white;
    text-align: left;
    filter: drop-shadow(0px 1px 4px #3c3c3c);
}
.list-news {
    height: 88%;
    overflow-y: scroll;
}
.list-news::-webkit-scrollbar {
    width:7px;
}
.list-news::-webkit-scrollbar-thumb {
    background-color: #223c50;
}
.news {
    border-radius: 15px;
    border-style: solid;
    background-color: #48aaff;
    color: white;
    padding: 15px;
}
@import url("../assets/css/style.css");
</style>