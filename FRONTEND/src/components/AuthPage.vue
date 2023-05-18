<template>
	<main class="block">
		<a @click="$router.push('/')">
			<img src="../assets/img/Logotype.png" />
		</a>
		<div class="main_block_login">
			<p class="Rubik-Medium" style="font-size: 28px">
				Вход по логину<br />
				и паролю
			</p>
			<div class="text-field text-field_floating">
				<input v-model="login" @input="Prikol" class="login_input text-field__input Rubik-Regular" type="login" id="login" placeholder="test@test.ru" />
				<label class="text-field__label Rubik" for="email">Логин/электронная почта</label>
			</div>
			<div class="text-field text-field_floating">
				<input v-model="pass" @input="Prikol" class="login_input text-field__input Rubik" id="password" :type="passwordFieldType" placeholder="test@test.ru" />
				<label class="text-field__label Rubik" for="password">Пароль</label>
				<a href="#" class="password-control" :class="{ view: isView }" @click="switchVisibility()"></a>
			</div>
			<div class="authButton">
				<button disabled @click="checkAuth" class="btn_login Rubik-Medium" id="submitButton">Войти</button>
			</div>
		</div>
		<a href="reset.html"><p class="p_cl Rubik-Regular" style="margin-bottom: 10px">Я забыл пароль</p></a>
		<a @click="$router.push('/register')" ><p class="p_cl Rubik-Regular">У меня нет аккаунта</p></a>
	</main>
</template>
<script>
export default {
	name: 'AuthPage',
	data() {
		return {
			passwordFieldType: 'password',
			isView: false,
			login: '',
			pass: '',
		}
	},
	mounted() {},
	computed() {},
	methods: {
		switchVisibility() {
			if (this.passwordFieldType == 'password') {
				this.passwordFieldType = 'text'
				this.isView = true
			} else {
				this.passwordFieldType = 'password'
				this.isView = false
			}
		},
		checkAuth() {
			// let firstName = this.login
			// let hash = this.pass
			let user = {users: [
					{
						email: this.login,
						hash: this.pass
					}
				]
			}
			//Auth.login(this.login, this.pass)
			this.$store.dispatch('login', user)
		},
		Prikol() {
			if (this.login != '' && this.pass != '') {
				document.getElementById('submitButton').disabled = false
				document.getElementById('submitButton').style = 'cursor: pointer;'
				document.getElementById('submitButton').style.background = '#73A2FF'
			} else {
				document.getElementById('submitButton').disabled = true
				document.getElementById('submitButton').style = 'cursor: default;'
				document.getElementById('submitButton').style.background = '#B5CEFF'
			}
		},
	},
	created: function() {

	}
}
</script>
<style>
a {
	cursor: pointer;
}
@import url(../assets/css/custom-library.css);
@import url(../assets/css/login.css);
</style>
