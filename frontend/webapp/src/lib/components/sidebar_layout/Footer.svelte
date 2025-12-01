<script lang="ts">
    let { copyright = 'Golang-SvelteKit Boilerplate 2024' } = $props();
</script>

<footer class="py-3 bg-light mt-auto">
    <div class="container-fluid px-4">
        <div class="d-flex justify-content-between small text-muted">
            <div>Copyright © {copyright}</div>
            <div>
                Made with <a href="https://svelte.dev/docs/kit"><b>Svelte</b></a>
                and <a href="https://sveltestrap.js.org/?path=/docs/sveltestrap-overview--docs"><b>Sveltestrap</b></a>
                (Bootstrap CSS components).
            </div>
        </div>
    </div>
</footer>
