<template>
   <div class="playground-page" :style="{visibility:visible ? 'visible':'hidden'}">
      <transition appear name="slide-fade">
         <div class="content-bg" ref="content">
            <iframe
               :src="`${localUrl}/hub/external/login?next=/hub/&username=${user.name}&token=${user.user_token}`">
            </iframe>
         </div>
      </transition>
   </div>
</template>

<script>
import { mapState } from "vuex";
export default {
   name: "playground",
   props: {
      visible: {
         type: Boolean,
         default: false,
      },
   },
   computed: {
      ...mapState({
         user: (state) => state.user,
      }),
      localUrl() {
         return window.BASE_API;
      },
   },
};
</script>

<style lang="stylus" scoped>
.playground-page {
   background: black;
   height: 100%;
   width:100%
   position absolute
   overflow hidden
   top 0px
   -webkit-overflow-scrolling: touch;
   iframe {
      width:100%
      height:100%
      border:none
   }
}

.upload-button {
   color: #607185;
   background: #dfe5ec;
   border: 1px solid #607185;
}

.content-bg {
   height: 100%;
   background: black;
   border: 1px #e4e4e4 solid;
}
</style>
