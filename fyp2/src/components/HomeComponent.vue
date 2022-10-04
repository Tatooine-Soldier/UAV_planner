<template>
    <section class="homepage">
        <section class="homepageOne">
            <div class="text-over-drone">
                <h1><b class="highlight-letter">U</b>AV <b class="highlight-letter">F</b>LIGHT <b class="highlight-letter">P</b>LANNER</h1>
            </div>
        </section>
        <section class="homepageTwo">Seeking to plan the best route for your next? Do you wan to get from point A to B as fast as possible? The system uses a series of nodes to calulate the fastest flight path from a given source to the target destination</section>
        <section class="homepageThree">.</section>
    </section>
</template>

<style scoped>
    .homepage {
        display: grid;
        grid-template-rows: 50% 25% 25%;
        font-family: Verdana, Geneva, Tahoma, sans-serif;
        border-left: solid 2px rgb(184, 255, 36);
        border-right: solid 2px rgb(184, 255, 36);
        background-color: rgb(57, 56, 56);
    }

    .homepageOne {
        background-image: url("../assets/greyDrone.png");
        background-position: bottom;
        background-size: cover;
        height: 580px;
        color: white;
        padding: 10px;
    }

    .text-over-drone {
        padding-top: 2%;
        position: absolute;
        text-align: center;
        width: 100%;
        font-weight: 400;
    }

    .text-over-drone h1{
        font-size: 56px;
        margin-top: 7.5%;
    }

    .highlight-letter {
        color: rgb(184, 255, 36);
    }

    .homepageTwo {
        display: grid;
        grid-template-columns: 50% 50%;
        column-gap: 2%;
        font-size: 20px;
        font-weight: 300;
        padding: 10px;
        color: white;
        text-align: center;;
        background-color: rgb(101, 100, 100);
    }

    .homepageThree {
        font-size: 20px;
        font-weight: 300;
        padding: 15px;
        color: white;
        text-align: center;;
        background-color: rgb(101, 100, 100);
    }


</style>

<script>

</script>