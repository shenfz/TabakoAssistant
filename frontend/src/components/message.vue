<script setup lang="ts">
import FactoryButton from "./slotContent/FactoryButton.vue";
</script>

<template>
  <div class="p-6 max-w-sm mx-auto bg-blue-500 hover:bg-orange-700 rounded-xl shadow-lg flex items-center space-x-4">
    <div class="shrink-0 hover:shrink">
      <img class="h-12 w-12 " src="../assets/images/img.svg" alt="ChitChat Logo">
    </div>
    <div>
      <div class="text-xl/8 font-medium text-black">ChitChat</div>
      <p class="text-slate-500">You have a new message!</p>
    </div>
  </div>

  <div class="flex w-full h-14 m-5">
           <FactoryButton>
             <template v-slot:btn>
               <!-- header 插槽的内容放这里 -->
               something
             </template>

           </FactoryButton>
  </div>
  <div class="p-8 max-w-sm mx-auto bg-white rounded-xl shadow-lg space-y-2 sm:py-4 sm:flex sm:items-center sm:space-y-0 sm:space-x-6">
    <img class="block mx-auto h-24 rounded-full sm:mx-0 sm:shrink-0" src="../assets/images/img.svg" alt="Woman's Face" />
    <div class="text-center space-y-2 sm:text-left">
      <div class="space-y-1">
        <p class="text-lg text-black font-semibold">
          Erin Lindford
        </p>
        <p class="text-slate-500 font-medium">
          Product Engineer
        </p>
      </div>
      <FactoryButton>Message5565</FactoryButton>
        </div>
  </div>

</template>

<style scoped>

</style>