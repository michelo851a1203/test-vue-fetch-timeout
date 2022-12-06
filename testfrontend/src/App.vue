<script setup lang="ts">
const mainClick = async () => {
  try {
    const controller = new AbortController();
    const timeoutID = setTimeout(() => {
      controller.abort();
      clearTimeout(timeoutID);
    }, 500)

    const response = await fetch('http://localhost:8080', {
      method: 'GET',
      redirect: 'follow',
      signal: controller.signal,
    })
    if (response.status !== 200) return;
    const mainData = await response.json();
    console.log(mainData);
  } catch(err) {
    console.log('this is error', err);
  }
}
</script>

<template>
  <button
    class="
      px-3 py-2 rounded-full bg-blue-400 border-none
      hover:bg-blue-800 hover:text-white
    "
    @click="mainClick"
  >
    testing
  </button>
</template>

<style scoped></style>
