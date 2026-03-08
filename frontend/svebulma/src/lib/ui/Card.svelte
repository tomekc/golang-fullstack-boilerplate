<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        title?: string;
        header?: Snippet;
        footer?: Snippet;
        children?: Snippet;
    }

    let { title, header, footer, children }: Props = $props();
</script>

<div class="card">
    {#if header || title}
        <div class="card-header">
            {#if header}
                {@render header()}
            {:else}
                <p class="card-header-title">{title}</p>
            {/if}
        </div>
    {/if}

    <div class="card-content">
        {@render children?.()}
    </div>

    {#if footer}
        <div class="card-footer">
            {@render footer()}
        </div>
    {/if}
</div>
