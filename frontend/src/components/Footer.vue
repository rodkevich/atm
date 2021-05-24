<template>
    <footer class="footer">
        <span class="footer-text">Belarus Gomel 2021</span>
    </footer>
</template>

<script>
export default {
    name: "Footer"
}
</script>

<style lang="scss" scoped>
    .footer {
        height: 100px;
        display: flex;
        justify-content: flex-end;
        align-items: center;
        background-color: #41B883;

        .footer-text {
            font-size: 1rem;
            margin-right: 20px;
            font-weight: bold;
        }
    }
</style>
