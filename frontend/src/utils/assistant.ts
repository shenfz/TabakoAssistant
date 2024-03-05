import {onMounted, ref} from "vue"
import {GetAppName} from "../../wailsjs/go/mainApp/MainApp"
import {HandleOrder} from "../../wailsjs/go/handleSpeechApp/HandleSpeechApp"
import {EventsOn} from "../../wailsjs/runtime/runtime"

let appName = ref("");
GetAppName().then((name) =>{
    appName.value = name;
})

UseEvents()

export let respMsg = ref("")

export function handleSpeechOrder(instruction:string){
    respMsg.value = instruction
    if (instruction.indexOf(appName.value) != -1){
      let command = instruction.split(appName.value);
      if (command[1].length > 0 ){
          HandleOrder(command[1]).then(respStr =>{
              respMsg.value = respStr
          })
      }else {
          respMsg.value = "无法理解当前指令！"
      }
    }else{
        respMsg.value = "想问什么，请带上指令！"
    }
}

function UseEvents(){
    EventsOn("showSpark", callBack);
}


function callBack(...data:any){
    respMsg.value = data
}