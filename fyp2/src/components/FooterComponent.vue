<template>
    <section class="footer">
        <section class="footer-columns">
            <h2>Company</h2>
            <ul>
                <li>About</li>
                <li>Legal</li>
                <li>Blog</li>
            </ul>
        </section>
        <section class="footer-columns">
            <h2>Products</h2>
            <ul>
                <li>API</li>
                <li>Pricing</li>
                <li>Documentation</li>
            </ul>
        </section>
        <section class="footer-columns">
            <h2>Contact</h2>
            <ul>
                <li>Support</li>
                <li>Report</li>
                <li>Share</li>
            </ul>
        </section>
    </section>
</template>

<style scoped>
    .footer {
        display: grid;
        grid-template-columns: 33% 33% 33%;
        background-color: rgb(77, 76, 76);
        bottom: 0px;
    }

    .footer-columns {
        text-align: center;;
        border-top: solid 2px rgb(184, 255, 36);
        padding: 10px;
        padding-bottom: 30px;
    }

    .footer-columns h2 {
        color: rgb(184, 255, 36);
    }

    .footer-columns ul {
        padding: 0px;
    }

    .footer-columns ul li {
        list-style: none;
    }
</style>