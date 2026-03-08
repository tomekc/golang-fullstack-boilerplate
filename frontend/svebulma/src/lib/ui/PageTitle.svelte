<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        title: string;
        subtitle?: string;
        actions?: Snippet;
    }

    let { title, subtitle, actions }: Props = $props();
</script>

<div class="is-flex is-justify-content-space-between is-align-items-flex-start mb-5">
    <div>
        <h1 class="title">{title}</h1>
        {#if subtitle}
            <p class="subtitle">{subtitle}</p>
        {/if}
    </div>
    {#if actions}
        <div class="buttons">
            {@render actions()}
        </div>
    {/if}
</div>
