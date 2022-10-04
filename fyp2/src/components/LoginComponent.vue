<template>
    <section class="login-page">
        <section class="login-container">
            <p>Login</p>
        </section>
        <section class="login-input-container">
            <p>Login to UFP</p>
            <section class="form-container">
                <form>
                    <section class="form-elements">
                        <label>Username:</label>
                        <input type="text" placeholder="username..." maxlength="30"/>
                    </section>    
                    <section class="form-elements">
                        <label>Password:</label>
                        <input type="password" placeholder="password..." maxlength="30"/>
                    </section>
                </form>
            </section>
        </section>
    </section>
</template>
<style scoped>
    .login-input-container {
        margin: 10%;
        width: 70%;
        border: solid rgb(184, 255, 36) 2px;
        background-color: white;
        padding: 10px;
        text-align: center;
    }
    .login-input-container p {font-size: 22px;}
    .login-page {padding: 10px; display: flex; justify-content: center; color: rgb(57, 56, 56); background-image: url("../assets/output-onlinepngtools.png"); background-size: cover;}
    .form-container {display: grid; grid-template-rows: 50% 50%;}
    .form-elements{padding: 10px; display: grid; grid-template-columns: 50% 50%;}

    input {
        border: solid 1px rgb(57, 56, 56);
        border-radius: 10px;
        padding: 5px;
    }

</style>