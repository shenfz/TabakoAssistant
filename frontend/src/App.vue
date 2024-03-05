<script lang="ts" setup>
import Head  from "./components/Head.vue"
import SpeechPanel from "./components/SpeechPanel.vue";
import Message from "./components/message.vue"
import {WindowSetAlwaysOnTop} from "../wailsjs/runtime"
import {ref} from "vue";
let toggleShow = ref(true);
function windowToggleShowTop() {
  toggleShow.value = !toggleShow.value;
  WindowSetAlwaysOnTop(!toggleShow.value)
};
</script>

<template>
  <div class="flex flex-col text-amber-50">
    <Head class="flex h-[35px] w-full items-center bg-gray-800/50" v-show="toggleShow"></Head>
    <SpeechPanel class="flex items-center h-[80px]"></SpeechPanel>
    <Message class="flex items-center"></Message>
    <el-icon class="text-cyan-50/50 hover:text-cyan-50 hover:cursor-pointer absolute right-1 bottom-1.5"
    size="30"
     @click="windowToggleShowTop"
    >
      <i-ep-HelpFilled />
    </el-icon>
  </div>

</template>

<style scoped>

</style>
