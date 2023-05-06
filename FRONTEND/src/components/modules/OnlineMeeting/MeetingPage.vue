<template>
    <div class="meeting_header">
        <div class="navigation">
            <div class="third-nav">
                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-door-exit" width="24" height="24" viewBox="0 0 24 24" stroke-width="1" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none"></path>
                    <path d="M13 12v.01m-10 8.99h18m-16 0v-16a2 2 0 0 1 2 -2h7.5m2.5 10.5v7.5m-3 -14h7m-3 -3l3 3l-3 3"></path>
                </svg>
                <a class="" @click="$router.push({name: 'modules'})">Выйти</a>
            </div>
            <div class="second-nav">
                <p class="" style="font-size:25px;">Онлайн обучение</p>
            </div>
        </div>
    </div>
    <div class="meeting_main module-name">
        <div class="active_rooms">
            Доступные комнаты
        </div>
        <div>
            <button class="rooms_list" v-for="room in rooms" :key="room" @click="$router.push({ name: 'conversation', params: { roomId: room.roomId } })"
            > {{room.title}}</button>
        </div>
    </div>

</template>

<script>
export default {
    name: "MeetingPage",
    data() {
        return {
            rooms: [
                {
                    roomId: 'Русский_Язык_28.04.23',
                    title: 'Русский язык'
                },
                {
                    roomId: '2',
                    title: 'Математика'
                },
            ]
        }
    }
}
</script>

<style>
.meeting_header {
    padding: 1px 0px 1px 0px;
    background-color: #e4f4ff;
}
.active_rooms {
    margin: 15px;
}
.rooms_list {
    border-radius: 30px;
    border-width: 0px;
    filter: drop-shadow(0px 0px 3px black);
    font-size: 24px;
    background-color: #FFFFFF;
    padding: 15px 15px 10px 15px;
    margin-inline: 20px;
    margin-bottom: 2em;
}
</style>