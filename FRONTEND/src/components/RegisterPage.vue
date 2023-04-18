<template>
	<main class="block">
		<a @click="$router.push('/')">
			<img src="../assets/img/Logotype.png" />
		</a><br />
		<div class="main_block_register">
			<p class="Rubik-Medium" style="font-size: 28px">Регистрация аккаунта</p>
			<br />
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="login" class="signup_input text-field__input Rubik-Regular" type="login" id="login" placeholder="test@test.ru" />
				<label class="text-field__label Rubik" for="email">Электронная почта</label>
			</div>
			<br />
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="tel" class="signup_input text-field__input Rubik-Regular tel" type="tel" id="tel" placeholder="Номер телефона" />
				<label class="text-field__label Rubik" for="tel">Номер телефона</label>
			</div>
			<br />
			<div class="hint_block Rubik-Regular">
				<h5>Пароль должен содержать:</h5>
				<ul class="Rubik">
					<li>От 6 до 64 символов</li>
					<li>Минимум 1 заглавную и 1 строчную букву латинского алфавита (a-z, A-Z)</li>
					<li>Минимум 1 цифру</li>
				</ul>
				<hr class="password_divider" />
				<p class="p_pass Rubik">
					Допустимые специальные символы:<br />
					! @ # $&nbsp;% ^ &amp; * ( ) _ + = ; : , ./ ?&nbsp;| ` ~ [ ] { }
				</p>
			</div>
			<div class="text-field text-field_floating">
				<input
					@input="checkReg"
					v-model="password"
					class="signup_input text-field__input Rubik-Regular"
					:type="passwordFieldType"
					id="password"
					placeholder="Пароль"
				/>
				<label class="text-field__label Rubik" for="password">Пароль</label>
				<p class="password-control" :class="{ view: isView }" @click="switchVisibility()"></p>
			</div>
			<br />
			<div class="text-field text-field_floating">
				<input
					@input="checkReg"
					class="signup_input text-field__input Rubik-Regular"
					:type="passwordFieldType"
					id="password_repeat"
					placeholder="Повторите пароль"
				/>
				<label class="text-field__label Rubik" for="password">Повторите пароль</label>
				<p class="password-control" :class="{ view: isView }" @click="switchVisibility()"></p>
			</div>
			<br />
			<p class="Rubik-Medium" style="text-align: left; margin-bottom: 20px">Личные данные</p>
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="surname" class="signup_input text-field__input Rubik-Regular" type="surname" id="surname" placeholder="Фамилия" />
				<label class="text-field__label Rubik" for="surname">Фамилия</label>
			</div>
			<br />
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="name" class="signup_input text-field__input Rubik-Regular" type="name" id="name" placeholder="Имя" />
				<label class="text-field__label Rubik" for="name">Имя</label>
			</div>
			<br />
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="midname" class="signup_input text-field__input Rubik-Regular" type="name" id="midname" placeholder="Отчество" />
				<label class="text-field__label Rubik" for="name">Отчество</label>
			</div>
			<br />
			<div class="text-field text-field_floating">
				<input @input="checkReg" v-model="date" class="signup_input text-field__input Rubik-Regular" type="date" id="date" placeholder="ДР" />
				<label class="text-field__label Rubik" for="date">Дата рождения</label>
			</div>
			<br />
		</div>
		<br />
		<div class="terms_block_login">
			<div class="term">
				<label class="checkbox bounce">
					<input @input="checkReg" type="checkbox" id="box" />
					<svg viewBox="0 0 21 21">
						<polyline points="5 10.75 8.5 14.25 16 6"></polyline>
					</svg>
				</label>
				<p class="terms Rubik">
					Соглашаюсь на обработку персональных данных в соответствии с
					<a href="#" class="signup_url">Политикой обработки персональных данных пользователей </a> и принимаю условия
					<a href="#" class="signup_url">Пользовательского соглашения</a>
				</p>
			</div>
			<div class="term">
				<label class="checkbox bounce">
					<input type="checkbox" />
					<svg viewBox="0 0 21 21">
						<polyline points="5 10.75 8.5 14.25 16 6"></polyline>
					</svg>
				</label>
				<p class="terms Rubik">
					Хочу быть в курсе событий: получать новости о работе <a href="#" class="signup_url">сайта</a>, анонсы и приглашения на полезные мероприятия —
					конференции, семинары, бесплатные вебинары и пр. (информационная и рекламная <a href="#" class="signup_url">рассылка</a>)
				</p>
			</div>
		</div>
		<br />
		<button class="signup_button Rubik-Medium" @click="register" id="submitButton" disabled>Зарегистрироваться</button>
	</main>

	<footer>
		<div class="footer_block">
			<div class="first_footer_block">
				<p class="Rubik-Regular" style="font-size: 25px; color: white;">Наши контакты</p>
				<p class="Rubik-Medium" style="font-size: 40px; color: white;">8 800 111-11-11</p>
				<p class="Rubik" style="color: white;">По общим вопросам (в будни с<br />7:00 до 18:00 по московскому времени)</p>
			</div>
			<div class="second_footer_block">
				<p class="Rubik-Medium" style="font-size: 40px; color: white;">help@romanton.ru</p>
			</div>
			<div class="third_footer_block">
				<p class="Rubik-Regular" style="font-size: 25px; color: white;">Подписаться на новости</p>
				<img src="../assets/img/tg.svg" style="width: 150px; color: white;" />
				<img src="../assets/img/vk.svg" style="width: 150px; color: white;" />
			</div>
		</div>
	</footer>
</template>
<script>
window.addEventListener('DOMContentLoaded', function () {
	[].forEach.call(document.querySelectorAll('.tel'), function (input) {
		var keyCode
		function mask(event) {
			event.keyCode && (keyCode = event.keyCode)
			var pos = this.selectionStart
			if (pos < 3) event.preventDefault()
			var matrix = '+7 (___)-___-____',
				i = 0,
				def = matrix.replace(/\D/g, ''),
				val = this.value.replace(/\D/g, ''),
				new_value = matrix.replace(/[_\d]/g, function (a) {
					return i < val.length ? val.charAt(i++) || def.charAt(i) : a
				})
			i = new_value.indexOf('_')
			if (i != -1) {
				i < 5 && (i = 3)
				new_value = new_value.slice(0, i)
			}
			var reg = matrix
				.substr(0, this.value.length)
				.replace(/_+/g, function (a) {
					return '\\d{1,' + a.length + '}'
				})
				.replace(/[+()]/g, '\\$&')
			reg = new RegExp('^' + reg + '$')
			if (!reg.test(this.value) || this.value.length < 5 || (keyCode > 47 && keyCode < 58)) this.value = new_value
			if (event.type == 'blur' && this.value.length < 5) this.value = ''
		}

		input.addEventListener('input', mask, false)
		input.addEventListener('focus', mask, false)
		input.addEventListener('blur', mask, false)
		input.addEventListener('keydown', mask, false)
	})
})
import Auth from '../services/auth.service'
export default {
	name: 'RegisterPage',
	data() {
		return {
			isView: false,
			passwordFieldType: 'password',
			login: '',
			tel: '',
			password: '',
			surname: '',
			name: '',
			midname: '',
			date: '',
		}
	},
	methods: {
		switchVisibility() {
			if (this.passwordFieldType == 'password') {
				this.passwordFieldType = 'text'
				this.isView = true
			} else {
				this.passwordFieldType = 'password'
				this.isView = false
			}
			//this.passwordFieldType = this.passwordFieldType === "password" ? "text" : "password";
		},
		checkReg() {
			if (
				document.getElementById('login').value &&
				document.getElementById('tel').value &&
				document.getElementById('password').value &&
				document.getElementById('password_repeat').value &&
				document.getElementById('surname').value &&
				document.getElementById('name').value &&
				document.getElementById('midname').value &&
				document.getElementById('date').value &&
				document.getElementById('box').checked
			) {
				document.getElementById('submitButton').disabled = false
				document.getElementById('submitButton').style = 'cursor: pointer;'
				document.getElementById('submitButton').style.background = '#73A2FF'
			} else {
				document.getElementById('submitButton').disabled = true
				document.getElementById('submitButton').style = 'cursor: default;'
				document.getElementById('submitButton').style.background = '#B5CEFF'
			}
		},
		register() {
			Auth.register(this.name, this.surname, this.midname, this.tel, this.login, this.date, '0', this.password)
		},
	},
}
</script>
<style>
* {
	box-sizing:content-box;
}
footer {
	color: white;
}
@import url(../assets/css/custom-library.css);
@import url(../assets/css/signup.css);
</style>
