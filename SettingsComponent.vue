<template>
    <section class="settings-container">
        <form>
            <section class="settings-form">
                <div class="settings-row">
                    <label for="font-size">Select Font Size: </label>
                    <input type="number" id="font-size" name="font-size" ref="font">
                </div>
                <div class="settings-row">
                    <label for="theme">Select Theme: </label>
                    <select name="theme" id="theme" ref="theme">
                        <option value="light">Light</option>
                        <option value="dark">Dark</option>
                    </select>
                </div>
                <div class="settings-row">
                    <label for="units">Select Measurement Units: </label>
                    <select name="units" id="units" ref="units">
                        <option value="km">Kilometres</option>
                        <option value="m">Metres</option>
                    </select>
                </div>
            </section>
            <button type="button" @click="handleSubmit()">Apply</button>
        </form>
    </section>
</template>

<script>
export default {
    data() {
        return {
            size: 0,
            theme: "",
            units: ""
        }
    }, methods: {
        handleSubmit() {
            this.size =  this.$refs.font.value
            this.theme =  this.$refs.theme.value
            this.units = this.$refs.units.value
            console.log(this.$refs.units.value, this.$refs.theme.value, this.$refs.font.value)
        }
    }
}
</script>

<style>
    .settings-container {
        background-color: rgb(117, 115, 115);
        color: white;
        padding: 5px;
    }

    .settings-form {
        display: grid;
        grid-template-rows: 33% 33% 33%;
    }

    .settings-row {
        padding: 5px;
    }

    .settings-row input {
        float: right;
    }

    .settings-row select {
        float: right;
    }

    select { width: 35%;}

    .settings-container button {
        margin-left: 40%;
        margin-right: 20%;
    }

</style>