<script lang="ts">
    import type { Snippet } from 'svelte';

    interface Props {
        children?: Snippet;
        fluid?: boolean;
        padding?: string;
    }

    let { children, fluid = true, padding = 'p-5' }: Props = $props();
</script>

<div class="sb-main-content is-flex is-flex-direction-column is-flex-grow-1" style="min-width: 0;">
    <main class="is-flex-grow-1">
        <div class={['container', fluid ? 'is-fluid' : '', padding].filter(Boolean).join(' ')}>
            {@render children?.()}
        </div>
    </main>
</div>
