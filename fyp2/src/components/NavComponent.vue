<template>
    <section class="navbar" id="sticky">
        <section class="nav-content-container">
            <section class="nav-sections">
                <section class="logo-container">
                    <img src="../assets/logo.png" id="logo"/>
                </section>
            </section>
            <section class="nav-sections">
                <section class="nav-lists">
                    <section class="links"><b>Home</b></section>
                    <section class="links"><b>About</b></section>
                    <section class="links"><b>Contact</b></section>
                </section>
            </section>
            <section class="nav-sections">
                <section class="nav-buttons">
                    <ul>
                        <li><img src="../assets/login-icon.jpg" id="login-icon"/></li>
                        <li><button>Log In</button></li>
                    </ul>
                </section>
            </section>
        </section>
    </section>
</template>

<style scoped>
    .navbar {
        display: flex;
        flex-direction: row;
        background-color: rgb(184, 255, 36);
        color: rgb(80, 79, 79);
        font-size: 32px;
        max-height: 85px !important;
        border-bottom: solid rgb(80, 79, 79) 2px;
    }

    .nav-content-container {
        display: grid;
        grid-template-columns: 25% 50% 25%;
        background-color: rgb(184, 255, 36);
        width: 100%;
    }

    .nav-sections {
        text-align: center;
        text-shadow: 4px 4px 4px rgba(135, 135, 138, 0.8);
        display: flex;
        justify-content: center;
    }

    .nav-lists {
        display: grid;
        grid-template-columns: 33% 33% 33%;
        margin-left: 9%;
        margin-right: 9%;
    }

    #login-icon {
        height: 30px;
        width: 30px;
    }

    .links {
        margin-left: 10px;
        margin-right: 20px;
    }

    .links b:hover {
        color: white;
        cursor:pointer;
    }

    #logo {
        width: 100px;
        height: 75px;
    }

    .logo-container {
        padding-right: 10px;
        padding-left: 10;
        justify-content: center;
    }

    #sticky {
        position: sticky;
        top: 0;
        z-index: 1;
    }

    .nav-buttons {
        display: inline;
        justify-content: center;
    }

    .nav-buttons li {
        list-style: none;
        display: inline;
        justify-content: center;
        align-items: center;
    }

    .nav-buttons button {
        margin-right: 10px;
        padding: 3px;
        width: 60px;
        border: rgb(80, 79, 79) 2px solid;
        background-image: linear-gradient(120deg, rgb(255, 255, 255), rgb(184, 255, 36));
        color: rgb(61, 61, 61);
    }

    .nav-buttons button:hover {
        cursor: pointer;
    }


</style>

<script>


</script>