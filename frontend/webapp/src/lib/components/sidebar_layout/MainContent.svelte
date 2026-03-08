<script lang="ts">
    import Footer from './Footer.svelte';
    let { children, copyright } = $props();
</script>

<div id="layoutSidenav_content" class="is-flex is-flex-direction-column">
    <main class="is-flex-grow-1">
        <div class="container is-fluid p-5">
            {@render children?.()}
        </div>
    </main>

    <Footer {copyright} />
</div>

<style>
    #layoutSidenav_content {
        flex-grow: 1;
        min-width: 0;
    }
</style>
