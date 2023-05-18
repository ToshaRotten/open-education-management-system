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
            <v-table class="journal-table">
                <thead>
                <tr>
                    <th v-for="name in headers" :key="name">
                        {{ name.text }}
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr
                    v-for="item in journal"
                    :key="item.name"
                >
                    <td>{{ item.name }}</td>
                    <td>
                        <div class="graduates">
                            <div class="graduate" v-for="grad in item.grades" :key="grad">
                                {{ grad }}
                            </div>
                        </div>
                    </td>
                    <td>
                        {{ averageGraduate(item.grades) }}
                    </td>
                </tr>
                </tbody>
            </v-table>
        </div>
    </div>

</template>

<script>

export default {
    name: "JournalPage",
    data() {
        return {
            firstName: this.$store.getters.loadData.firstName,
            lastName: this.$store.getters.loadData.lastName,
            headers: [
                {text: "Урок"},
                {text: "Оценки"},
                {text: "Средний балл"}
            ],
            journal: [
                {
                    name: "Русский язык",
                    grades: [5,4,3,5,2,5,4]
                },
                {
                    name: "Математика",
                    grades: [5,5]
                },
                {
                    name: "Информатика",
                    grades: [5,5]
                },
                {
                    name: "Элементы мат. логики",
                    grades: [5,5,5,4,3,5,2,5,4,"н"]
                },
                {
                    name: "Литература",
                    grades: [2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2]
                },
                {
                    name: "Основы алгоритмизации",
                    grades: [5,5]
                }
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
        averageGraduate: function(arr) {
            let sum = 0;
            for(let i = 0; i < arr.length; i++)
            {
                if (typeof arr[i] == 'number')
                sum = sum + arr[i]
            }
            return (sum / arr.length).toFixed(2)
        },
        colorizeGrads:function () {
            let grads = document.getElementsByClassName('graduate')
            console.log(grads)
            for(let i = 0; i < grads.length; i++) {
                if (grads[i].innerText == '5') grads[i].classList.add('mG')
                else if (grads[i].innerText == '4') grads[i].classList.add('mLG')
                else if (grads[i].innerText == '3') grads[i].classList.add('mY')
                else if (grads[i].innerText == '2') grads[i].classList.add('mR')
                else grads[i].classList.add('mR')
            }
        }
    },
    mounted() {
        this.colorizeGrads()
    }
}
</script>

<style>
.graduates {
    display:inline-block;
}
.graduate {
    float:left;
    margin: 5px;
    background-color: #9dd1ff;
    color: white;
    padding: 0px 5px 0px 5px;
    border-radius: 7px;
}
.journal-table {
    text-align:left;
    border-radius: 5px;
    height:100%;
    filter:drop-shadow(0px 0px 2px black);
}
.mG {
    background-color: #01e230;
}
.mLG {
    background-color: #01e230;
}
.mY {
    background-color: #e2af00;
}
.mR {
    background-color: #e2210b;
}
</style>