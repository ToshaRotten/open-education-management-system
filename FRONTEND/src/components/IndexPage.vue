<template>
<header>
        <div class="header-area ">
            <div id="sticky-header" class="main-header-area">
                <div class="container-fluid">
                    <div class="header_bottom_border">
                        <div class="row align-items-center">
                            <div class="col-xl-3 col-lg-6">
                                <div class="logo">
                                    <a @click="$router.push('/')">
                                        <img src="img/Logotype.png" alt="">
                                    </a>
                                </div>
                            </div>
                            <!--<div class="col-xl-2 col-lg-2">
                                <div class="logo">
                                    <a href="index.html">
                                        <img src="img/cinsu.png" alt="">
                                    </a>
                                </div>
                            </div>-->
                            <div class="col-xl-6 col-lg-7">
                                <div class="main-menu  d-none d-lg-block">
                                    <nav>
                                        <ul id="navigation" style="font-family: 'Rubik Light';">
                                            <li><a href="#">Главная</a></li>
                                            <li><a href="#">О системе</a></li>
                                            <li><a href="#">Кому это нужно?</a></li>
                                            <li><a href="#">Каталог модулей</a></li>
                                            <li><a href="#">Контакты</a></li>
                                        </ul>
                                    </nav>
                                </div>
                            </div>
                            <div class="col-xl-3 col-lg-3 d-none d-lg-block">
                                <div class="Appointment">

                                    <div class="d-none d-lg-block">
                                        <a class="boxed-btn3 Rubik" @click="$router.push('/auth')">Войти</a>
                                    </div>
                                </div>
                            </div>
                            <div class="col-12">
                                <div class="mobile_menu d-block d-lg-none"></div>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    </header>
		<div class="slider_area">
        <div class="single_slider  d-flex align-items-center slider_bg_1">
            <div class="container">
                <div class="row align-items-center">
                    <div class="col-lg-7 col-md-6">
                        <div class="slider_text">
                            <h3 class="wow fadeInLeft text-1" data-wow-duration="1s" data-wow-delay=".3s">Совершенно
                                новая информационная система</h3><br>
                            <p class="wow fadeInLeft text-2" data-wow-duration="1s" data-wow-delay=".4s">Удобный и
                                практичный инструмент для<br> образовательных организаций</p>
                            <div class="row align-items-center">
                                <div class="col-xl-3 col-lg-6 sldier_btn wow fadeInLeft" data-wow-duration="1s"
                                    data-wow-delay=".5s">
                                    <a @click="$router.push('/auth')" class="boxed-btn4">Войти</a>
                                </div>
                                <div class="col-xl-3 col-lg-6 sldier_btn wow fadeInLeft" data-wow-duration="1s"
                                    data-wow-delay=".5s">
                                    <a @click="$router.push('/register')" class="boxed-btn5">Зарегистрироваться</a>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="ilstration_img wow fadeInRight d-none d-lg-block text-right" data-wow-duration="1s"
            data-wow-delay=".2s">
            <img src="img/table.png" alt="">
        </div>
    </div>
</template>

<script>
export default {
    name: 'IndexPage',
    props: {},
    computed: {},
    methods: {},
};
</script>

<style scoped>
* {
	box-sizing:border-box;
}
  @import '../../public/css/animate.min.css';
	@import '../../public/css/bootstrap.min.css';
	@import '../../public/css/custom-library.css';
	@import '../../public/css/slicknav.css';
	@import '../../public/css/magnific-popup.css';
	@import '../../public/css/style.css';
</style>

