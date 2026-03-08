<script lang="ts">
    let { title, children } = $props();
</script>

<p class="menu-label">{title}</p>
<ul class="menu-list">
    {@render children?.()}
</ul>

<style>
    :global(.menu-label) {
        padding: 1.75rem 1rem 0.75rem;
        font-size: 0.75rem;
        font-weight: 800;
        text-transform: uppercase;
        letter-spacing: 0.05em;
        color: rgba(255, 255, 255, 0.25);
    }
</style>
