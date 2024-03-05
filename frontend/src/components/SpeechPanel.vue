<script setup lang="ts">
import { useSpeechRecognition } from '@vueuse/core'
import {ref, onMounted, watch} from "vue";
import {GetAppName} from "../../wailsjs/go/mainApp/MainApp"
import {GetSpeechAPIEventName} from "../../wailsjs/go/systemApp/SystemApp";
import {handleSpeechOrder,respMsg} from "../utils/assistant"

let appName = ref("");
let enableSpeechAPIEventName = ref("");
GetAppName().then((name) =>{
  appName.value = name;
})
GetSpeechAPIEventName().then((eventName) =>{
  enableSpeechAPIEventName.value = eventName;
})
let resultText = ref();
const {
  isSupported,
  isListening,
  isFinal,
  result,
  start,
  stop,
} = useSpeechRecognition({
  lang: 'zh-CN',
  interimResults: true,
  continuous: true,
});

onMounted(()=>{
  if (!isSupported) {
    alert("你的浏览器不支持语音识别")
  }else {
    start()
  }
})

watch(result,(newVal)=>{
  let reStr = /\.|\?|。|？+/
  let flag = reStr.test(newVal[newVal.length-1]);
  let strings  = newVal.split(reStr);
   if (flag) {
      resultText.value = strings[strings.length-2]
     handleSpeechOrder(resultText.value)
   } else{
     resultText.value =  strings[strings.length-1]
   }
})

watch(respMsg,(newVal) =>{
  resultText.value = newVal
})

</script>

<template>
  <div class="flex items-center justify-center h-full w-full text-3xl font-bold">
   {{resultText}}
  </div>
</template>

<style scoped>

</style>