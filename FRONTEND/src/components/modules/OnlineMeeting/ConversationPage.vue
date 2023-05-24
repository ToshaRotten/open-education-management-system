<template id="test">
        <JaaSMeeting
            :app-id="111"
            :room-name="room"
            lang="ru"
            :width="width"
            :height="height"
            :userInfo="user"
            :interface-config-overwrite="config"
            :config-overwrite="{prejoinPageEnabled: false, disableProfile: true, whiteboard: {enabled: true, collabServerBaseUrl: 'https://excalidraw-backend.example.com'}}"
            @on-ready-to-close="babe"
            :use-staging="true"
        />
</template>

<script>
// import {JitsiMeeting} from "@jitsi/vue-sdk";
import {JaaSMeeting} from "@jitsi/vue-sdk";
export default {
    name: "ConversationPage",
    computed: {

    },
    components: {
        'JaaSMeeting': JaaSMeeting
    },
    data() {
        return {
            room: this.$route.params.roomId,
            user: {
                displayName: this.$store.getters.loadData.firstName + " " + this.$store.getters.loadData.lastName,
                email: this.$store.getters.loadData.email
            },
            config: {
                SHOW_CHROME_EXTENSION_BANNER: true,
                SHOW_DEEP_LINKING_IMAGE: false,
                SHOW_JITSI_WATERMARK: false,
                HIDE_DEEP_LINKING_LOGO: false,
                    SHOW_POWERED_BY: false,
                    SHOW_BRAND_WATERMARK: false,
                    SHOW_PROMOTIONAL_CLOSE_PAGE: false,
                    BRAND_WATERMARK_LINK: '',
                    DEFAULT_BACKGROUND: '#ffffff',
                    RANDOM_AVATAR_URL_PREFIX: 'https://i.imgur.com/FN3kSip.jpg'
            },
            username: this.$store.getters.loadData.firstName + " " + this.$store.getters.loadData.lastName,
            width: window.innerWidth,
            height: window.innerHeight - 7
        }
    },
    methods: {
        babe: function () {
            this.$router.push({name:'meet'})
        },
        resize: function () {
            this.width = window.innerWidth;
            this.height = window.innerHeight;
        }
    },
    created() {
        window.addEventListener("resize", this.resize);
    },
    unmounted() {
        window.removeEventListener("resize", this.resize);
    },
}
</script>

<style>
div.watermark{
    visibility: hidden;
}
</style>