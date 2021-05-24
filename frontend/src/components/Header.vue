<template>
    <header class="header">
        <nav class="nav">
            <ul class="nav-list">
                <li v-for="item in navItems"
                    v-bind:item="navItems"
                    :key="item"
                    class="list"
                    ><a class="list-link" href="#">{{ item.title }}</a></li>
            </ul>
        </nav>
    </header>
</template>

<script>
export default {
    name: "Header",
    data() {
        return {
            navItems: [
                {title: "Main"},
                {title: "Gallery"},
                {title: "Forum"}
            ]
        }
    }
}
</script>

<style lang="scss" scoped>
    .header {
        height: 100px;
        margin-bottom: 30px;
        display: flex;
        justify-content: flex-end;
        align-items: center;
        background-color: #41B883;
    }

    .nav-list {
        display: flex;
        justify-content: flex-end;
        align-items: center;
        list-style-type: none;

        .list {
            width: 100px;
            border: 2px solid #41B883;

            .list-link {
                text-decoration: none;
                font-size: 1.3rem;
                font-weight: 700;
                color: #35495E;
            }
        }
    }
</style>
