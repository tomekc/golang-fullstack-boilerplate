<script lang="ts">
    import Footer from './Footer.svelte';
    let { children, copyright } = $props();
</script>

<div id="layoutSidenav_content" class="d-flex flex-column">
    <main class="flex-grow-1">
        <div class="container-fluid px-4 py-4">
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
