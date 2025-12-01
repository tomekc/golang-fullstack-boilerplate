<script lang="ts">
    let { title, children } = $props();
</script>

<div class="sb-sidenav-menu-heading">{title}</div>
{@render children?.()}

<style>
    :global(.sb-sidenav-menu-heading) {
        padding: 1.75rem 1rem 0.75rem;
        font-size: 0.75rem;
        font-weight: 800;
        text-transform: uppercase;
        letter-spacing: 0.05em;
        color: rgba(255, 255, 255, 0.25);
    }
</style>
