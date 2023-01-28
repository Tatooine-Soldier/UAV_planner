<template>
    <label for="Speed">Speed: </label>
    <input type="number" id="sspeed" name="speed" min="1" max="120" ref="selectedspeed" value="30" @click="handleThis()">
</template>

<script>
export default {
    data() {
      return {
        speed: 30
      }
    },
    methods: {
        handleThis() {
            this.speed = this.$refs.selectedspeed.value;
            console.log("this.speed:",this.speed)
            this.$emit('speedEvent', this.speed)
        }
    }

}
</script>