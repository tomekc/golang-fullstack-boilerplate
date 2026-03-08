<script lang="ts">
    let { copyright = 'Golang-SvelteKit Boilerplate 2024' } = $props();
</script>

<footer class="footer has-background-light">
    <div class="container is-fluid">
        <div class="is-flex is-justify-content-space-between is-size-7 has-text-grey">
            <div>Copyright © {copyright}</div>
            <div>
                Made with <a href="https://svelte.dev/docs/kit"><strong>Svelte</strong></a>
                and <a href="https://bulma.io"><strong>Bulma</strong></a>
                CSS framework.
            </div>
        </div>
    </div>
</footer>
