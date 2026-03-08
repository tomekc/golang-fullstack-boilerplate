<script lang="ts">
    import type { Snippet } from 'svelte';

    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger' | 'dark' | 'light';

    interface Props {
        color?: Color;
        dismissible?: boolean;
        children?: Snippet;
        ondismiss?: () => void;
    }

    let { color, dismissible = false, children, ondismiss }: Props = $props();
    let visible = $state(true);

    function dismiss() {
        visible = false;
        ondismiss?.();
    }
</script>

{#if visible}
    <div class={['notification', color ? `is-${color}` : ''].filter(Boolean).join(' ')}>
        {#if dismissible}
            <button class="delete" onclick={dismiss} aria-label="Dismiss"></button>
        {/if}
        {@render children?.()}
    </div>
{/if}
