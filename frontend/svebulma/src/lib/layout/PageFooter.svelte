<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        copyright?: string;
        children?: Snippet;
    }

    let { copyright, children }: Props = $props();
</script>

<footer class="footer has-background-light">
    <div class="container is-fluid">
        <div class="is-flex is-justify-content-space-between is-size-7 has-text-grey">
            {#if children}
                {@render children()}
            {:else if copyright}
                <div>Copyright © {copyright}</div>
            {/if}
        </div>
    </div>
</footer>
