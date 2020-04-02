<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>
  import {initWebSocket} from "@/websocket/websocket";

  export default {
    mounted() {
      initWebSocket()
    }
  }
</script>

<style>
#app {
}
</style>
