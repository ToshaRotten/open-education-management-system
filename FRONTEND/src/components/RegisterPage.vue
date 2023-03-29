<template>
    <main class="block">
		<a href="index.html"><img src="../assets/img/Logotype.png"></a><br>
		<div class="main_block_login">
			<p class="Rubik-Medium" style="font-size: 28px;">Регистрация аккаунта</p><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="login" id="login" placeholder="test@test.ru" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="email">Электронная почта</label>
			</div><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular tel" type="tel" id="tel" placeholder="Номер телефона" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="tel">Номер телефона</label>
			</div><br>
			<div class="hint_block Rubik-Regular">
				<h5>Пароль должен содержать:</h5>
				<ul class = "Rubik">
					<li>От 6 до 64 символов</li>
					<li>Минимум 1 заглавную и 1 строчную букву латинского алфавита (a-z, A-Z)</li>
					<li>Минимум 1 цифру</li>
				</ul>
				<hr class="password_divider">
				<p class="p_pass Rubik"> Допустимые специальные символы:<br> ! @ # $&nbsp;% ^ &amp; * ( ) _ + = ; : , ./ ?&nbsp;| ` ~ [ ] { }
				</p>
				
			</div>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="password" id="password" placeholder="Пароль" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="password">Пароль</label>
				<a href="#" class="password-control" onclick="return show_hide_password(this);"></a>
			</div><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="password" id="password_repeat" placeholder="Повторите пароль" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="password">Повторите пароль</label>
				<a href="#" class="password-control" onclick="return show_hide_password1(this);"></a>
			</div><br>
			<p class="Rubik-Medium" style="text-align: left; margin-bottom: 20px;">Личные данные</p>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="surname" id="surname" placeholder="Фамилия" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="surname">Фамилия</label>
			</div><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="name" id="name" placeholder="Имя" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="name">Имя</label>
			</div><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="name" id="midname" placeholder="Отчество" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="name">Отчество</label>
			</div><br>
			<div class="text-field text-field_floating">
				<input class="signup_input text-field__input Rubik-Regular" type="date" id="date" placeholder="ДР" onchange="toggleButton()">
				<label class="text-field__label Rubik" for="date">Дата рождения</label>
			</div><br>
		</div><br>
		<div class="terms_block_login">
			<div class = "term">
				<label class="checkbox bounce">
					<input type="checkbox">
					<svg viewBox="0 0 21 21">
						<polyline points="5 10.75 8.5 14.25 16 6"></polyline>
					</svg>
				</label>
				<p class="terms Rubik">Соглашаюсь на обработку персональных данных в соответствии с  <a href="#" class="signup_url">Политикой обработки персональных данных пользователей </a> и принимаю условия <a href="#" class="signup_url">Пользовательского соглашения</a></p>
			</div>
			<div class = "term">
				<label class="checkbox bounce">
					<input type="checkbox">
					<svg viewBox="0 0 21 21">
						<polyline points="5 10.75 8.5 14.25 16 6"></polyline>
					</svg>
				</label>
				<p class="terms Rubik">Хочу быть в курсе событий: получать новости о работе сайта <a href="#" class="signup_url">educont.ru</a>, анонсы и приглашения на полезные мероприятия — конференции, семинары, бесплатные вебинары и пр. (информационная и рекламная <a href="#" class="signup_url">рассылка</a>)</p>
			</div>
			</div><br>
			<form action="/personal/main.html">
				<button class="signup_button Rubik-Medium" id="submitButton" disabled>Зарегистрироваться</button>
			</form>
			
	</main>

	<footer>
		<div class="footer_block">
			<div class="first_footer_block">
				<p class="Rubik-Regular" style="font-size:25px;">Наши контакты</p>
				<p class="Rubik-Medium" style="font-size:40px;">8 800 111-11-11</p>
				<p class="Rubik">По общим вопросам (в будни с<br>7:00 до 18:00 по московскому времени)</p>
			</div>
			<div class="second_footer_block">
				<p class="Rubik-Medium" style="font-size:40px;">help@romanton.ru</p>
			</div>
			<div class="third_footer_block">
				<p class="Rubik-Regular" style="font-size:25px;">Подписаться на новости</p>
				<img src="../assets/img/tg.svg" style="width:150px;">
				<img src="../assets/img/vk.svg" style="width:150px;">
			</div>
		</div>
	</footer>
</template>
<script>

</script>
<style>
@import url(../assets/css/custom-library.css);
@import url(../assets/css/signup.css);
</style>