<template>
    <main class="block">
		<a href="index.html">
            <img src="../assets/img/Logotype.png">
        </a>
		<div class="main_block_login">
			<p class="Rubik-Medium" style="font-size: 28px;">Вход по логину<br> и паролю</p>
			<div class="text-field text-field_floating">
				<input v-model="login" class="login_input text-field__input Rubik-Regular" type="login" id="login" placeholder="test@test.ru">
				<label class="text-field__label Rubik" for="email">Логин/электронная почта</label>
			</div>
			<div class="text-field text-field_floating">
				<input v-model="pass" class="login_input text-field__input Rubik" id="password" :type="passwordFieldType" placeholder="test@test.ru">
				<label class="text-field__label Rubik" for="password">Пароль</label>
				<a href="#" class="password-control" :class="{view: isView}" @click="switchVisibility()"></a>
			</div>
			<form @click="$router.push('/dashboard')">
				<button class="btn_login Rubik-Medium" id="submitButton">Войти</button>
			</form>
			
		</div>
		<a href="reset.html"><p class = "p_cl Rubik-Regular" style="margin-bottom:10px;">Я забыл пароль</p></a>
		<a href="/signup.html"><p class = "p_cl Rubik-Regular">У меня нет аккаунта</p></a>
	</main>
</template>
<script>
//import signin from '@/assets/js/signin'
export default {
  name: 'MainPage',
  data () {
      return {
        passwordFieldType: "password",
        isView: false,
        login: "",
        pass: ""
      }
  },
  mounted() {
    
  },
  computed() {

  },
  methods: {
    switchVisibility() {
        if (this.passwordFieldType == "password") {
            this.passwordFieldType = "text"
            this.isView = true

        }
        else {
            this.passwordFieldType = "password"
            this.isView = false
        }
        //this.passwordFieldType = this.passwordFieldType === "password" ? "text" : "password";
    },
    checkFill() {
        if (this.login != "" && this.pass != "") {
            console.log(1)
        }
    }
  }
}
</script>
<style>
@import url(../assets/css/login.css);
@import url(../assets/css/custom-library.css);

</style>