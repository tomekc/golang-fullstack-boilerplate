<script lang="ts">
    import type { Snippet } from 'svelte';

    type Color = 'primary' | 'link' | 'info' | 'success' | 'warning' | 'danger' | 'dark' | 'light' | 'black' | 'white';
    type Size = 'normal' | 'medium' | 'large';

    interface Props {
        color?: Color;
        size?: Size;
        rounded?: boolean;
        deletable?: boolean;
        ondelete?: () => void;
        children?: Snippet;
    }

    let { color, size, rounded = false, deletable = false, ondelete, children }: Props = $props();

    const classes = $derived([
        'tag',
        color ? `is-${color}` : '',
        size && size !== 'normal' ? `is-${size}` : '',
        rounded ? 'is-rounded' : '',
    ].filter(Boolean).join(' '));
</script>

<span class={classes}>
    {@render children?.()}
    {#if deletable}
        <button class="delete is-small" onclick={ondelete} aria-label="Remove tag"></button>
    {/if}
</span>
